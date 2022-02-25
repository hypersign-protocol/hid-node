package keeper

import (
	"context"
	"time"

	//"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
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

// Ref: https://w3c-ccg.github.io/did-resolution/#resolving-algorithm
func (k Keeper) Resolve(goCtx context.Context, req *types.QueryGetDidDocByIdRequest) (*types.QueryGetDidDocByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Invalid DID Check
	err := utils.IsValidDid(req.DidDocId)
	switch err {
	case types.ErrInvalidDidElements:
		return &types.QueryGetDidDocByIdResponse{
			DidDocument:         nil,
			DidDocumentMetadata: nil,
			DidResolutionMetadata: &types.DidResolveMeta{
				Error: "invalidDid",
			},
		}, nil
	case types.ErrInvalidDidMethod:
		return &types.QueryGetDidDocByIdResponse{
			DidDocument:         nil,
			DidDocumentMetadata: nil,
			DidResolutionMetadata: &types.DidResolveMeta{
				Error: "methodNotSupported",
			},
		}, nil
	}

	// Check if DID Document exists
	didDoc, err := k.GetDid(&ctx, req.DidDocId)
	if err != nil {
		return &types.QueryGetDidDocByIdResponse{
			DidDocument:         nil,
			DidDocumentMetadata: nil,
			DidResolutionMetadata: &types.DidResolveMeta{
				Error: "notFound",
			},
		}, nil
	}

	// Check if DID Document is deactivated
	if didDoc.GetMetadata().GetDeactivated() {
		return &types.QueryGetDidDocByIdResponse{
			DidDocument:         nil,
			DidDocumentMetadata: didDoc.GetMetadata(),
			DidResolutionMetadata: &types.DidResolveMeta{
				Retrieved: ctx.BlockTime().Format(time.RFC3339),
			},
		}, nil
	}

	return &types.QueryGetDidDocByIdResponse{
		DidDocument:         didDoc.GetDid(),
		DidDocumentMetadata: didDoc.GetMetadata(),
		DidResolutionMetadata: &types.DidResolveMeta{
			Retrieved: ctx.BlockTime().Format(time.RFC3339),
		},
	}, nil
}
