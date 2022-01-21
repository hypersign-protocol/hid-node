package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/utils"
	"github.com/hypersign-protocol/hid-node/x/did/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Dids(goCtx context.Context, req *types.QueryDidsRequest) (*types.QueryDidsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)
	// Get the part of the store that keeps posts (using post key, which is "Post-value-")
	didStore := prefix.NewStore(store, []byte(types.DidKey))
	bz := didStore.Get(utils.UnsafeStrToBytes(req.Did))

	// Return a struct containing a list of posts and pagination info
	return &types.QueryDidsResponse{DidDoc: utils.UnsafeBytesToStr(bz)}, nil
}
