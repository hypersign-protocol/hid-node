package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) QueryCredential(goCtx context.Context, req *types.QueryCredentialRequest) (*types.QueryCredentialResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	cred, err := k.GetCredential(&ctx, req.CredId)
	if err != nil {
		return nil, err
	}

	return &types.QueryCredentialResponse{CredStatus: cred}, nil
}
