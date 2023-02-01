package keeper

import (
	"context"
	saodidparser "github.com/SaoNetwork/sao-did/parser"
	"strings"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdatePaymentAddress(goCtx context.Context, msg *types.MsgUpdatePaymentAddress) (*types.MsgUpdatePaymentAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	logger := k.Logger(ctx)

	if !k.CheckCreator(ctx, msg.Creator, msg.Did) {
		logger.Error("invalid Creator", "creator", msg.Creator, "did", msg.Did)
		return nil, types.ErrInvalidCreator
	}

	accId := msg.GetAccountId()
	accIdSplits := strings.Split(accId, ":")

	OldAddr, found := k.GetPaymentAddress(ctx, msg.Did)
	if found {
		if OldAddr.Address == accIdSplits[2] {
			logger.Error("try to update the same address as the old one", "paymentAddress", OldAddr)
			return nil, types.ErrSamePayAddr
		}
	}

	if len(accIdSplits) == 3 && accIdSplits[0] == "cosmos" && accIdSplits[1] == ctx.ChainID() {
		// err if did is a key did or empty, which means update payment address for sid
		did, err := saodidparser.Parse(msg.Did)
		if err != nil {
			logger.Error("failed to parse did", "did", msg.Did)
			return nil, types.ErrInvalidDid
		}
		switch did.Method {
		case "sid":
			storedDid, found := k.GetDid(ctx, accId)
			if !found {
				logger.Error("account id has not bound to did yet", "accountId", msg.AccountId, "did", msg.Did)
				return nil, types.ErrBindingNotFound
			}
			if msg.Did != storedDid.Did {
				logger.Error("did in binding proof is different to did in message", "Did(stored)", storedDid.Did, "Did(msg)", msg.Did)
				return nil, types.ErrInconsistentDid
			}

			paymentAddress := types.PaymentAddress{
				Did:     storedDid.Did,
				Address: accIdSplits[2],
			}
			k.SetPaymentAddress(ctx, paymentAddress)
		case "key":
			paymentAddress := types.PaymentAddress{
				Did:     msg.Did,
				Address: accIdSplits[2],
			}
			k.SetPaymentAddress(ctx, paymentAddress)
		default:
			logger.Error("unsupported did", "did", msg.Did)
			return nil, types.ErrUnsupportedDid
		}
	} else {
		logger.Error("invalid chain address", "expectedChainId", ctx.ChainID(), "accountId", msg.AccountId)
		return nil, types.ErrInvalidAccountId
	}

	return &types.MsgUpdatePaymentAddressResponse{}, nil
}
