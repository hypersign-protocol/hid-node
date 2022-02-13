package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SchemaCount(goCtx context.Context, req *types.QuerySchemaCountRequest) (*types.QuerySchemaCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var count uint64 = k.GetSchemaCount(ctx)

	return &types.QuerySchemaCountResponse{Count: count}, nil
}
