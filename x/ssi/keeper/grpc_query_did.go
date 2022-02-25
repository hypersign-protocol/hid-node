package keeper

import (
	"context"
	"time"

	//"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DidParam(goCtx context.Context, req *types.QueryDidParamRequest) (*types.QueryDidParamResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))

	var didResolveList []*types.DidResolutionResponse
	_, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var (
			didResolve types.DidResolutionResponse
			didDoc types.DidDocument
		)
		if err := k.cdc.Unmarshal(value, &didDoc); err != nil {
			return err
		}

		didResolve.DidDocument = didDoc.Did
		didResolve.DidDocumentMetadata = didDoc.Metadata
		didResolve.DidResolutionMetadata = &types.DidResolveMeta{
			Error: "",
			Retrieved: ctx.BlockTime().Format(time.RFC3339),
		}

		didResolveList = append(didResolveList, &didResolve)
		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var didDocCount uint64 = k.GetDidCount(ctx)
	if req.Count {
		return &types.QueryDidParamResponse{TotalDidCount: didDocCount}, nil
	}
	return &types.QueryDidParamResponse{TotalDidCount: didDocCount, DidDocList: didResolveList}, nil
}

// Ref: https://w3c-ccg.github.io/did-resolution/#resolving-algorithm
func (k Keeper) ResolveDid(goCtx context.Context, req *types.QueryGetDidDocByIdRequest) (*types.QueryGetDidDocByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Invalid DID Check
	err := utils.IsValidDid(req.DidId)
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
	didDoc, err := k.GetDid(&ctx, req.DidId)
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
