package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/did/types"
)

func (k Keeper) DidCount(goCtx context.Context, req *types.QueryDidCountRequest) (*types.QueryDidCountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var count uint64 = k.GetDidCount(ctx)

	return &types.QueryDidCountResponse{Count: count}, nil
}
