package ssifee

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/identityfee/keeper"
	"github.com/hypersign-protocol/hid-node/x/identityfee/types"
)

// InitGenesis initializes the identityfee module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetFeeParam(ctx, *genState.CreateDidFee, types.ParamStoreKeyCreateDidFee)
	k.SetFeeParam(ctx, *genState.UpdateDidFee, types.ParamStoreKeyUpdateDidFee)
	k.SetFeeParam(ctx, *genState.DeactivateDidFee, types.ParamStoreKeyDeactivateDidFee)
	k.SetFeeParam(ctx, *genState.CreateSchemaFee, types.ParamStoreKeyCreateSchemaFee)
	k.SetFeeParam(ctx, *genState.RegisterCredentialStatusFee, types.ParamStoreKeyRegisterCredentialStatusFee)
}

// ExportGenesis returns the identity module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	createDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyCreateDidFee)
	updateDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyUpdateDidFee)
	deactivateDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyDeactivateDidFee)
	createSchemaFee := k.GetFeeParams(ctx, types.ParamStoreKeyCreateSchemaFee)
	registerCredentialStatusFee := k.GetFeeParams(ctx, types.ParamStoreKeyRegisterCredentialStatusFee)

	genesis.CreateDidFee = &createDidFee
	genesis.UpdateDidFee = &updateDidFee
	genesis.DeactivateDidFee = &deactivateDidFee
	genesis.CreateSchemaFee = &createSchemaFee
	genesis.RegisterCredentialStatusFee = &registerCredentialStatusFee

	return genesis
}
