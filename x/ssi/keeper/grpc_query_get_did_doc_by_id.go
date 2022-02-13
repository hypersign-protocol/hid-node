package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetDidDocById(goCtx context.Context, req *types.QueryGetDidDocByIdRequest) (*types.QueryGetDidDocByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	didDoc, err := k.GetDid(&ctx, req.DidDocId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "DidDoc not found")
	}

	return &types.QueryGetDidDocByIdResponse{
		DidDoc: didDoc,
	}, nil
}
