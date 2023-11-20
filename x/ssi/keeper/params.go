package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetFeeParam(ctx sdk.Context, fee sdk.Coin, ssiParamStoreKey []byte) {
	k.paramSpace.Set(ctx, ssiParamStoreKey, fee)
}

func (k Keeper) GetFeeParams(ctx sdk.Context, ssiParamStoreKey []byte) sdk.Coin {
	var fee sdk.Coin
	k.paramSpace.Get(ctx, ssiParamStoreKey, &fee)
	return fee
}
