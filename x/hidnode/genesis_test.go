package hidnode_test

import (
	"testing"

	keepertest "github.com/hypersign-protocol/hid-node/testutil/keeper"
	"github.com/hypersign-protocol/hid-node/testutil/nullify"
	"github.com/hypersign-protocol/hid-node/x/hidnode"
	"github.com/hypersign-protocol/hid-node/x/hidnode/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HidnodeKeeper(t)
	hidnode.InitGenesis(ctx, *k, genesisState)
	got := hidnode.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
