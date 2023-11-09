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

func (k Keeper) CredentialStatusByID(
	goCtx context.Context,
	req *types.QueryCredentialStatusRequest,
) (*types.QueryCredentialStatusResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	cred, err := k.getCredentialStatusFromState(&ctx, req.CredId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCredentialStatusNotFound, err.Error())
	}

	return &types.QueryCredentialStatusResponse{CredentialStatus: cred}, nil
}

func (k Keeper) CredentialStatuses(
	goCtx context.Context,
	req *types.QueryCredentialStatusesRequest,
) (*types.QueryCredentialStatusesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Define a variable that will store a list of credentials
	var credentials []*types.CredentialStatusState
	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Get the key-value module store using the store key
	store := ctx.KVStore(k.storeKey)
	// Get the part of the store that keeps credential statuses
	credStore := prefix.NewStore(store, []byte(types.CredKey))
	// Paginate the credential store based on PageRequest
	_, err := query.Paginate(credStore, req.Pagination, func(key []byte, value []byte) error {
		var credential types.CredentialStatusState
		if err := k.cdc.Unmarshal(value, &credential); err != nil {
			return err
		}
		credentials = append(credentials, &credential)
		return nil
	})
	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryCredentialStatusesResponse{
		CredentialStatuses: credentials,
		Count:              k.getCredentialStatusCount(ctx),
	}, nil
}
