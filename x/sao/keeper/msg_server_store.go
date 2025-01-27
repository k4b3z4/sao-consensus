package keeper

import (
	"context"
	"strings"

	nodetypes "github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ipfs/go-cid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) Store(goCtx context.Context, msg *types.MsgStore) (*types.MsgStoreResponse, error) {
	var sigDid string
	var err error
	ctx := sdk.UnwrapSDKContext(goCtx)
	proposal := &msg.Proposal
	if proposal == nil {
		return nil, status.Errorf(codes.InvalidArgument, "proposal is required")
	}

	if proposal.Owner != "all" {
		sigDid, err = k.verifySignature(ctx, proposal.Owner, proposal, msg.JwsSignature)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, sdkerrors.Wrap(types.ErrorNoPermission, "No permission to update the open data model")
	}

	var metadata ordertypes.Metadata
	var node nodetypes.Node

	if proposal.CommitId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid commitId")
	}

	if proposal.DataId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid dataId")
	}

	// check cid
	_, err = cid.Decode(proposal.Cid)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCid, "invalid cid: %s", proposal.Cid)
	}

	if !strings.Contains(proposal.CommitId, proposal.DataId) {
		// validate the permission for all update operations
		meta, isFound := k.Keeper.model.GetMetadata(ctx, proposal.DataId)
		if !isFound {
			return nil, status.Errorf(codes.NotFound, "dataId Operation:%d not found", proposal.Operation)
		}

		isValid := meta.Owner == sigDid
		if !isValid {
			for _, readwriteDid := range meta.ReadwriteDids {
				if readwriteDid == sigDid {
					isValid = true
					break
				}
			}

			if !isValid {
				return nil, sdkerrors.Wrap(types.ErrorNoPermission, "No permission to update the model")
			}
		}
	}

	// check provider
	node, found := k.node.GetNode(ctx, proposal.Provider)
	if !found {
		return nil, sdkerrors.Wrapf(nodetypes.ErrNodeNotFound, "%s does not register yet", node.Creator)
	}

	commitId := proposal.CommitId
	lastCommitId := proposal.CommitId
	if strings.Contains(proposal.CommitId, "|") {
		// extract the base version from the proposal
		lastCommitId = strings.Split(proposal.CommitId, "|")[0]
		commitId = strings.Split(proposal.CommitId, "|")[1]
	}

	metadata = ordertypes.Metadata{
		DataId:     proposal.DataId,
		Owner:      proposal.Owner,
		Alias:      proposal.Alias,
		GroupId:    proposal.GroupId,
		Tags:       proposal.Tags,
		Cid:        proposal.Cid,
		Commit:     commitId,
		ExtendInfo: proposal.ExtendInfo,
		Rule:       proposal.Rule,
	}

	if proposal.Size_ == 0 {
		proposal.Size_ = 1
	}

	var order = ordertypes.Order{
		Creator:   msg.Creator,
		Owner:     proposal.Owner,
		Cid:       proposal.Cid,
		Expire:    proposal.Timeout + int32(ctx.BlockHeight()),
		Duration:  proposal.Duration,
		Status:    types.OrderPending,
		Replica:   proposal.Replica,
		Metadata:  &metadata,
		Operation: proposal.Operation,
		Size_:     proposal.Size_,
	}

	if node.Creator != "" {
		order.Provider = node.Creator
	}

	var sps []nodetypes.Node

	if order.Provider == msg.Creator {
		if order.Operation == 1 {
			sps = k.node.RandomSP(ctx, int(order.Replica))
			if order.Replica <= 0 || int(order.Replica) > len(sps) {
				return nil, sdkerrors.Wrapf(types.ErrInvalidReplica, "replica should > 0 and <= %d", len(sps))
			}
		} else if order.Operation > 1 {
			sps = k.FindSPByDataId(ctx, proposal.DataId)
		}
	}

	if order.Size_ == 0 {
		order.Size_ = 1
	}

	price := sdk.NewDecWithPrec(1, 6)

	ownerAddress, err := k.did.GetCosmosPaymentAddress(ctx, proposal.Owner)
	if err != nil {
		return nil, err
	}

	denom := k.staking.BondDenom(ctx)
	amount, _ := sdk.NewDecCoinFromDec(denom, price.MulInt64(int64(order.Size_)).MulInt64(int64(order.Replica)).MulInt64(int64(order.Duration))).TruncateDecimal()
	balance := k.bank.GetBalance(ctx, ownerAddress, denom)

	if balance.IsLT(amount) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientCoin, "insuffcient coin: need %d", amount.Amount.Int64())
	}

	order.Amount = amount

	spCreators := make([]string, 0)
	for _, sp := range sps {
		spCreators = append(spCreators, sp.Creator)
	}

	orderId, err := k.order.NewOrder(ctx, &order, spCreators)
	if err != nil {
		return nil, err
	}

	// avoid version conflicts
	meta, found := k.model.GetMetadata(ctx, proposal.DataId)
	if found {
		if meta.OrderId > orderId {
			// report error if order id is less than the latest version
			return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidCommitId, "invalid commitId: %s, detected version conficts with order: %d", commitId, meta.OrderId)
		}

		lastOrder, isFound := k.order.GetOrder(ctx, meta.OrderId)
		if isFound {
			if lastOrder.Status == ordertypes.OrderPending || lastOrder.Status == ordertypes.OrderInProgress || lastOrder.Status == ordertypes.OrderDataReady {
				return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidLastOrder, "unexpected last order: %s, status: %d", meta.OrderId, lastOrder.Status)
			}
		} else {
			return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidLastOrder, "invalid last order: %s", meta.OrderId)
		}

		if !strings.Contains(meta.Commit, lastCommitId) {
			// report error if base version is not the latest version
			return nil, sdkerrors.Wrapf(nodetypes.ErrInvalidCommitId, "invalid commitId: %s, detected version conficts, should be %s", lastCommitId, meta.Commit[:36])
		}
	}

	if order.Provider == msg.Creator {
		shards := make(map[string]*types.ShardMeta, 0)
		for p, shard := range order.Shards {
			node, node_found := k.node.GetNode(ctx, p)
			if !node_found {
				continue
			}
			meta := types.ShardMeta{
				ShardId:  shard.Id,
				Peer:     node.Peer,
				Cid:      shard.Cid,
				Provider: order.Provider,
			}
			shards[p] = &meta
		}

		return &types.MsgStoreResponse{
			OrderId: orderId,
			Shards:  shards,
		}, nil
	} else {
		return &types.MsgStoreResponse{
			OrderId: orderId,
		}, nil
	}

}
