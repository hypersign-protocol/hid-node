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

	k.SetFeeParam(ctx, *genState.Params.CreateDidFee, types.ParamStoreKeyCreateDidFee)
	k.SetFeeParam(ctx, *genState.Params.UpdateDidFee, types.ParamStoreKeyUpdateDidFee)
	k.SetFeeParam(ctx, *genState.Params.DeactivateDidFee, types.ParamStoreKeyDeactivateDidFee)
	k.SetFeeParam(ctx, *genState.Params.CreateSchemaFee, types.ParamStoreKeyCreateSchemaFee)
	k.SetFeeParam(ctx, *genState.Params.RegisterCredentialStatusFee, types.ParamStoreKeyRegisterCredentialStatusFee)
}

// ExportGenesis returns the ssi module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.ChainNamespace = k.GetChainNamespace(&ctx)

	createDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyCreateDidFee)
	updateDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyUpdateDidFee)
	deactivateDidFee := k.GetFeeParams(ctx, types.ParamStoreKeyDeactivateDidFee)
	createSchemaFee := k.GetFeeParams(ctx, types.ParamStoreKeyCreateSchemaFee)
	registerCredentialStatusFee := k.GetFeeParams(ctx, types.ParamStoreKeyRegisterCredentialStatusFee)

	genesis.Params.CreateDidFee = &createDidFee
	genesis.Params.UpdateDidFee = &updateDidFee
	genesis.Params.DeactivateDidFee = &deactivateDidFee
	genesis.Params.CreateSchemaFee = &createSchemaFee
	genesis.Params.RegisterCredentialStatusFee = &registerCredentialStatusFee

	return genesis
}
