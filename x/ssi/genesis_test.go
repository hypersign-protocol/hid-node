package ssi_test

import (
	"testing"

	keepertest "github.com/hypersign-protocol/hid-node/testutil/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.DefaultGenesis()

	k, ctx := keepertest.SsiKeeper(t)
	ssi.InitGenesis(ctx, *k, *genesisState)
	exportedGenesisState := ssi.ExportGenesis(ctx, *k)

	ExpectedDidMethod := "hs"
	ExpectedDidNamespace := ""

	require.NotNil(t, exportedGenesisState)
	require.Equal(t, ExpectedDidMethod, genesisState.DidMethod)
	require.Equal(t, ExpectedDidNamespace, genesisState.DidNamespace)
}
