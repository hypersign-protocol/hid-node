package keeper_test

import (
	"testing"

	testkeeper "github.com/hypersign-protocol/hid-node/testutil/keeper"
	"github.com/hypersign-protocol/hid-node/x/did/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DidKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
