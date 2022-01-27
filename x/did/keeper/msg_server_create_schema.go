package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/did/types"
)

func (k msgServer) CreateSchema(goCtx context.Context, msg *types.MsgCreateSchema) (*types.MsgCreateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var schema = types.Schema{
		Creator:   msg.Creator,
		SchemaID:  msg.SchemaID,
		SchemaStr: msg.SchemaStr,
	}
	// Add a Schema to the store and get back the ID
	id := k.AppendSchema(ctx, schema)

	// Return the Id of the Schema
	return &types.MsgCreateSchemaResponse{Id: id}, nil
}
