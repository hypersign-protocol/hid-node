package keeper

import (
	"context"
	"fmt"
	//"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (k Keeper) GetDidDocById(goCtx context.Context, req *types.QueryGetDidDocByIdRequest) (*types.QueryGetDidDocByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	didDoc, err := k.GetDid(&ctx, req.DidDocId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "DidDoc not found")
	}

	if didDoc.GetMetadata().GetDeactivated() {
		return nil, sdkerrors.Wrap(types.ErrDidDocDeactivated, fmt.Sprintf("DidDoc ID: %s", req.DidDocId))
	}

	return &types.QueryGetDidDocByIdResponse{
		DidDoc: didDoc,
	}, nil
}
