package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DidDocCount(goCtx context.Context, req *types.QueryDidDocCountRequest) (*types.QueryDidDocCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var didDocCount uint64 = k.GetDidCount(ctx)

	return &types.QueryDidDocCountResponse{Count: didDocCount}, nil
}
