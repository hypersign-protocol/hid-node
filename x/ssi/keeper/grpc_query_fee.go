package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// QuerySSIFee fetches fees for all SSI based transactions
func (k Keeper) QuerySSIFee(goCtx context.Context, _ *types.QuerySSIFeeRequest) (*types.QuerySSIFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	registerDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyRegisterDidFee)
	updateDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyUpdateDidFee)
	deactivateDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyDeactivateDidFee)
	registerCredentialSchemaFee := k.GetFeeParams(ctx, types.ParamStoreKeyRegisterCredentialSchemaFee)
	updateCredentialSchemaFee := k.GetFeeParams(ctx, types.ParamStoreKeyUpdateCredentialSchemaFee)
	registerCredentialStatusFee := k.GetFeeParams(ctx, types.ParamStoreKeyRegisterCredentialStatusFee)
	updateCredentialStatusFee := k.GetFeeParams(ctx, types.ParamStoreKeyUpdateCredentialStatusFee)

	return &types.QuerySSIFeeResponse{
		RegisterDidFee:              &registerDidFee,
		UpdateDidFee:                &updateDidFee,
		DeactivateDidFee:            &deactivateDidFee,
		RegisterCredentialSchemaFee: &registerCredentialSchemaFee,
		UpdateCredentialSchemaFee:   &updateCredentialSchemaFee,
		RegisterCredentialStatusFee: &registerCredentialStatusFee,
		UpdateCredentialStatusFee:   &updateCredentialStatusFee,
	}, nil
}
