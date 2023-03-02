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

func (k Keeper) QueryCredential(goCtx context.Context, req *types.QueryCredentialRequest) (*types.QueryCredentialResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	cred, err := k.GetCredentialStatusFromState(&ctx, req.CredId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCredentialStatusNotFound, err.Error())
	}

	return &types.QueryCredentialResponse{CredStatus: cred}, nil
}

func (k Keeper) QueryCredentials(goCtx context.Context, req *types.QueryCredentialsRequest) (*types.QueryCredentialsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Define a variable that will store a list of credentials
	var credentials []*types.Credential
	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Get the key-value module store using the store key
	store := ctx.KVStore(k.storeKey)
	// Get the part of the store that keeps credential statuses
	credStore := prefix.NewStore(store, []byte(types.CredKey))
	// Paginate the credential store based on PageRequest
	_, err := query.Paginate(credStore, req.Pagination, func(key []byte, value []byte) error {
		var credential types.Credential
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
	return &types.QueryCredentialsResponse{Credentials: credentials, TotalCount: k.GetCredentialStatusCount(ctx)}, nil
}
