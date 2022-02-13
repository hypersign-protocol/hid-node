package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/utils"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetSchema(goCtx context.Context, req *types.QueryGetSchemaRequest) (*types.QueryGetSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)
	// Get the part of the store that keeps posts (using post key, which is "Schema-value-")
	schemaStore := prefix.NewStore(store, []byte(types.SchemaKey))
	bz := schemaStore.Get(utils.UnsafeStrToBytes(req.SchemaId))
	if bz == nil {
		return nil, status.Error(codes.NotFound, "schema not found")
	}

	var schema types.Schema
	k.cdc.MustUnmarshal(bz, &schema)

	return &types.QueryGetSchemaResponse{
		Schema: &schema,
	}, nil
}
