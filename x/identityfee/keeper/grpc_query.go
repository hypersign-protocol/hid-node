package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/identityfee/types"
)

// QuerySSIFee fetches fees for all SSI based transactions
func (k Keeper) QuerySSIFee(goCtx context.Context, _ *types.QuerySSIFeeRequest) (*types.QuerySSIFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	createDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyCreateDidFee)
	updateDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyUpdateDidFee)
	deactivateDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyDeactivateDidFee)
	createSchemaFee := k.GetFeeParams(ctx, types.ParamStoreKeyCreateSchemaFee)
	registerCredentialStatusFee := k.GetFeeParams(ctx, types.ParamStoreKeyRegisterCredentialStatusFee)

	return &types.QuerySSIFeeResponse{
		CreateDidFee:                &createDidFee,
		UpdateDidFee:                &updateDidFee,
		DeactivateDidFee:            &deactivateDidFee,
		CreateSchemaFee:             &createSchemaFee,
		RegisterCredentialStatusFee: &registerCredentialStatusFee,
	}, nil
}
