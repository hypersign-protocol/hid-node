package ssi_test

import (
	"testing"

	keepertest "github.com/hypersign-protocol/vid-node/testutil/keeper"
	"github.com/hypersign-protocol/vid-node/x/ssi"
	"github.com/hypersign-protocol/vid-node/x/ssi/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.DefaultGenesis()

	k, ctx := keepertest.SsiKeeper(t)
	ssi.InitGenesis(ctx, *k, *genesisState)
	exportedGenesisState := ssi.ExportGenesis(ctx, *k)

	ExpectedChainNamespace := ""

	require.NotNil(t, exportedGenesisState)
	require.Equal(t, ExpectedChainNamespace, genesisState.ChainNamespace)
}
