package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) OrderPledge(ctx sdk.Context, sp sdk.AccAddress, order *ordertypes.Order) error {

	pledge, foundPledge := k.GetPledge(ctx, sp.String())

	if !foundPledge {
		pledge = types.Pledge{
			Creator:             sp.String(),
			TotalOrderPledged:   sdk.NewInt64Coin(sdk.DefaultBondDenom, 0),
			TotalStoragePledged: sdk.NewInt64Coin(sdk.DefaultBondDenom, 0),
			Reward:              sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 0),
			RewardDebt:          sdk.NewInt64DecCoin(sdk.DefaultBondDenom, 0),
			TotalStorage:        0,
			LastRewardAt:        ctx.BlockTime().Unix(),
		}
	}

	pool, foundPool := k.GetPool(ctx)

	if !foundPool {
		return sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	pledge.TotalOrderPledged = pledge.TotalOrderPledged.Add(order.Amount)

	params := k.GetParams(ctx)

	transfer_coins := sdk.Coins{order.Amount}

	if pledge.TotalStorage > 0 {
		pending := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage).Sub(pledge.Reward.Amount)
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
		pledge.TotalStorage += int64(order.Size_)
		pledge.RewardDebt.Amount = pool.AccPledgePerByte.Amount.MulInt64(pledge.TotalStorage)
	}

	rewardPerByte := sdk.NewInt64DecCoin(params.BlockReward.Denom, 0)
	if !params.BlockReward.Amount.IsZero() {
		rewardPerByte.Amount = sdk.NewDecCoinFromCoin(params.BlockReward).Amount.QuoInt64(pool.TotalStorage)

		storage_dec_pledge := sdk.NewInt64DecCoin(params.BlockReward.Denom, 0)
		storage_dec_pledge.Amount = storage_dec_pledge.Amount.MulInt64(int64(order.Shards[sp.String()].Size_ * order.Duration))
		storage_pledge_coin, _ := storage_dec_pledge.TruncateDecimal()

		transfer_coins.Add(storage_pledge_coin)
		transfer_coins = transfer_coins.Add(storage_pledge_coin)

		pledge.TotalStoragePledged.Amount = pledge.TotalStoragePledged.Amount.Add(storage_pledge_coin.Amount)

	}

	order_pledge_err := k.bank.SendCoinsFromAccountToModule(ctx, sp, types.ModuleName, transfer_coins)

	if order_pledge_err != nil {
		return order_pledge_err
	}

	pool.TotalStorage += int64(order.Shards[sp.String()].Size_)

	order.Shards[sp.String()].PledgePerByte = rewardPerByte

	k.SetPledge(ctx, pledge)

	k.SetPool(ctx, pool)
	return nil
}

func (k Keeper) OrderRelease(ctx sdk.Context, sp sdk.AccAddress, order *ordertypes.Order) error {

	pledge, found_pledge := k.GetPledge(ctx, sp.String())
	if !found_pledge {
		return sdkerrors.Wrap(types.ErrPledgeNotFound, "")
	}

	pool, found_pool := k.GetPool(ctx)

	if !found_pool {
		return sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	if pledge.TotalStorage > 0 {
		pending := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage).Sub(pledge.Reward.Amount)
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
		pledge.RewardDebt.Amount = pool.AccPledgePerByte.Amount.MulInt64(pledge.TotalStorage)
	}

	if order != nil {
		pledge.TotalStorage -= int64(order.Size_)
		params := k.GetParams(ctx)
		transfer_coins := sdk.Coins{order.Amount}

		storage_dec_pledge := sdk.NewInt64DecCoin(params.BlockReward.Denom, 0)
		storage_dec_pledge.Amount = storage_dec_pledge.Amount.MulInt64(int64(order.Shards[sp.String()].Size_ * order.Duration))
		storage_pledge_coin, _ := storage_dec_pledge.TruncateDecimal()

		if !storage_pledge_coin.IsZero() {
			transfer_coins = transfer_coins.Add(storage_pledge_coin)
		}

		transfer_err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sp, transfer_coins)

		if transfer_err != nil {
			return transfer_err
		}

		pledge.TotalStoragePledged = pledge.TotalStoragePledged.Sub(storage_pledge_coin)

		pledge.TotalOrderPledged = pledge.TotalOrderPledged.Sub(order.Amount)

		pool.TotalStorage -= int64(order.Shards[sp.String()].Size_)
	}

	k.SetPledge(ctx, pledge)

	k.SetPool(ctx, pool)

	return nil
}
