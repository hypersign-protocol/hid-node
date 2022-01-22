package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/hypersign-protocol/hid-node/testutil/keeper"
	"github.com/hypersign-protocol/hid-node/x/hidnode/keeper"
	"github.com/hypersign-protocol/hid-node/x/hidnode/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.HidnodeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
