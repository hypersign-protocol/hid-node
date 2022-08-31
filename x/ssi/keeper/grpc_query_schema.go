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

func (k Keeper) GetSchema(goCtx context.Context, req *types.QueryGetSchemaRequest) (*types.QueryGetSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var schema []*types.Schema = k.GetSchemaFromStore(ctx, req.SchemaId)

	return &types.QueryGetSchemaResponse{
		Schema: schema,
	}, nil
}

func (k Keeper) SchemaParam(goCtx context.Context, req *types.QuerySchemaParamRequest) (*types.QuerySchemaParamResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Define a variable that will store a list of schemas
	var schemas []*types.Schema
	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)
	// Get the part of the store that keeps schema (using post key, which is "Schema-value-")
	schemaStore := prefix.NewStore(store, []byte(types.SchemaKey))
	// Paginate the schema store based on PageRequest
	_, err := query.Paginate(schemaStore, req.Pagination, func(key []byte, value []byte) error {
		var schema types.Schema
		if err := k.cdc.Unmarshal(value, &schema); err != nil {
			return err
		}
		schemas = append(schemas, &schema)
		return nil
	})
	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QuerySchemaParamResponse{SchemaList: schemas, TotalCount: k.GetSchemaCount(ctx)}, nil
}
