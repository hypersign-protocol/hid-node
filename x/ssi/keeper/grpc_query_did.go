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

func (k Keeper) QueryDidDocuments(goCtx context.Context, req *types.QueryDidDocumentsRequest) (*types.QueryDidDocumentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))

	var didResolveList []*types.QueryDidDocumentResponse
	_, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var (
			didResolve types.QueryDidDocumentResponse
			didDoc     types.DidDocumentState
		)
		if err := k.cdc.Unmarshal(value, &didDoc); err != nil {
			return err
		}

		didResolve.DidDocument = didDoc.DidDocument
		didResolve.DidDocumentMetadata = didDoc.DidDocumentMetadata

		didResolveList = append(didResolveList, &didResolve)
		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var didDocCount uint64 = k.GetDidCount(ctx)
	if req.Count {
		return &types.QueryDidDocumentsResponse{TotalDidCount: didDocCount}, nil
	}
	return &types.QueryDidDocumentsResponse{TotalDidCount: didDocCount, DidDocList: didResolveList}, nil
}

// Ref: https://w3c-ccg.github.io/did-resolution/#resolving-algorithm
func (k Keeper) QueryDidDocument(goCtx context.Context, req *types.QueryDidDocumentRequest) (*types.QueryDidDocumentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if DID Document exists
	didDoc, err := k.GetDidDocumentState(&ctx, req.DidId)
	if err != nil {
		return nil, err
	}

	return &types.QueryDidDocumentResponse{
		DidDocument:         didDoc.GetDidDocument(),
		DidDocumentMetadata: didDoc.GetDidDocumentMetadata(),
	}, nil
}
