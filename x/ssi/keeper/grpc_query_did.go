package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DidDocuments(goCtx context.Context, req *types.QueryDidDocumentsRequest) (*types.QueryDidDocumentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))

	var didDocuments []*types.DidDocumentState
	_, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		didDoc, err := k.getDidDocumentState(&ctx, string(key))
		if err != nil {
			return sdkerrors.Wrap(types.ErrDidDocNotFound, err.Error())
		}

		didDocuments = append(didDocuments, didDoc)
		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var didDocCount uint64 = k.getDidDocumentCount(ctx)

	return &types.QueryDidDocumentsResponse{DidDocuments: didDocuments, Count: didDocCount}, nil
}

func (k Keeper) DidDocumentByID(goCtx context.Context, req *types.QueryDidDocumentRequest) (*types.QueryDidDocumentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if DID Document exists
	didDoc, err := k.getDidDocumentState(&ctx, req.DidId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrDidDocNotFound, err.Error())
	}

	return &types.QueryDidDocumentResponse{
		DidDocument:         didDoc.GetDidDocument(),
		DidDocumentMetadata: didDoc.GetDidDocumentMetadata(),
	}, nil
}
