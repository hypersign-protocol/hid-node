package ssi

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// InitGenesis initializes the ssi module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetChainNamespace(&ctx, genState.ChainNamespace)

	k.SetFeeParam(ctx, *genState.Params.RegisterDidFee, types.ParamStoreKeyRegisterDidFee)
	k.SetFeeParam(ctx, *genState.Params.UpdateDidFee, types.ParamStoreKeyUpdateDidFee)
	k.SetFeeParam(ctx, *genState.Params.DeactivateDidFee, types.ParamStoreKeyDeactivateDidFee)
	k.SetFeeParam(ctx, *genState.Params.RegisterCredentialSchemaFee, types.ParamStoreKeyRegisterCredentialSchemaFee)
	k.SetFeeParam(ctx, *genState.Params.UpdateCredentialSchemaFee, types.ParamStoreKeyUpdateCredentialSchemaFee)
	k.SetFeeParam(ctx, *genState.Params.RegisterCredentialStatusFee, types.ParamStoreKeyRegisterCredentialStatusFee)
	k.SetFeeParam(ctx, *genState.Params.UpdateCredentialStatusFee, types.ParamStoreKeyUpdateCredentialStatusFee)
}

// ExportGenesis returns the ssi module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.ChainNamespace = k.GetChainNamespace(&ctx)

	registerDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyRegisterDidFee)
	updateDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyUpdateDidFee)
	deactivateDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyDeactivateDidFee)
	registerCredentialSchemaFee := k.GetFeeParams(ctx, types.ParamStoreKeyRegisterCredentialSchemaFee)
	updateCredentialSchemaFee := k.GetFeeParams(ctx, types.ParamStoreKeyUpdateCredentialSchemaFee)
	registerCredentialStatusFee := k.GetFeeParams(ctx, types.ParamStoreKeyRegisterCredentialStatusFee)
	updateCredentialStatusFee := k.GetFeeParams(ctx, types.ParamStoreKeyUpdateCredentialStatusFee)

	genesis.Params.RegisterDidFee = &registerDidFee
	genesis.Params.UpdateDidFee = &updateDidFee
	genesis.Params.DeactivateDidFee = &deactivateDidFee
	genesis.Params.RegisterCredentialSchemaFee = &registerCredentialSchemaFee
	genesis.Params.UpdateCredentialSchemaFee = &updateCredentialSchemaFee
	genesis.Params.RegisterCredentialStatusFee = &registerCredentialStatusFee
	genesis.Params.UpdateCredentialStatusFee = &updateCredentialStatusFee

	return genesis
}
