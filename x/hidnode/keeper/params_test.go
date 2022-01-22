package keeper_test

import (
	"testing"

	testkeeper "github.com/hypersign-protocol/hid-node/testutil/keeper"
	"github.com/hypersign-protocol/hid-node/x/hidnode/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HidnodeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
