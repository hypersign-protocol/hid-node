package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/did/types"
)

func (k Keeper) SchemaCount(goCtx context.Context, req *types.QuerySchemaCountRequest) (*types.QuerySchemaCountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var count uint64 = k.GetSchemaCount(ctx)

	return &types.QuerySchemaCountResponse{Count: count}, nil
}
