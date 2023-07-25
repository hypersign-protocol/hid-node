package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetFeeParam(ctx sdk.Context, fee sdk.Coin, identityfeeParamStoreKey []byte) {
	k.paramSpace.Set(ctx, identityfeeParamStoreKey, fee)
}

func (k Keeper) GetFeeParams(ctx sdk.Context, identityfeeParamStoreKey []byte) sdk.Coin {
	var fee sdk.Coin
	k.paramSpace.Get(ctx, identityfeeParamStoreKey, &fee)
	return fee
}
