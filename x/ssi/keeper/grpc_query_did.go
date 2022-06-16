package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
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
			didDoc     types.DidDocument
		)
		if err := k.cdc.Unmarshal(value, &didDoc); err != nil {
			return err
		}

		didResolve.DidDocument = didDoc.Did
		didResolve.DidDocumentMetadata = didDoc.Metadata

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
func (k Keeper) ResolveDid(goCtx context.Context, req *types.QueryGetDidDocByIdRequest) (*types.DidResolutionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if DID Document exists
	didDoc, err := k.GetDid(&ctx, req.DidId)
	if err != nil {
		return nil, err
	}

	return &types.DidResolutionResponse{
		DidDocument:         didDoc.GetDid(),
		DidDocumentMetadata: didDoc.GetMetadata(),
	}, nil
}
