package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
)

func (k msgServer) CreateSchema(goCtx context.Context, msg *types.MsgCreateSchema) (*types.MsgCreateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schemaMsg := msg.GetSchema()
	schemaID := schemaMsg.GetId()

	if err := utils.IsValidSchemaID(schemaID); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidSchemaID, err.Error())
	}

	if k.HasSchema(ctx, schemaID) {
		return nil, sdkerrors.Wrap(types.ErrSchemaExists, fmt.Sprintf("Schema ID:  %s", schemaID))
	}

	//Get the DID of SChema's Author
	authorDID, err := k.GetDid(&ctx, schemaMsg.GetAuthor())
	if err != nil {
		return nil, sdkerrors.Wrap(err, fmt.Sprintf("The DID %s is not available", schemaMsg.GetAuthor()))
	}

	// Signature check
	if err := k.VerifySignatureOnCreateSchema(&ctx, schemaMsg, authorDID.GetSigners(), msg.GetSignatures()); err != nil {
		return nil, err
	}

	var schema = types.Schema{
		Type:         schemaMsg.GetType(),
		ModelVersion: schemaMsg.GetModelVersion(),
		Id:           schemaMsg.GetId(),
		Name:         schemaMsg.GetName(),
		Author:       schemaMsg.GetAuthor(),
		Authored:     schemaMsg.GetAuthored(),
		Schema:       schemaMsg.GetSchema(),
	}

	id := k.AppendSchema(ctx, schema)

	return &types.MsgCreateSchemaResponse{Id: id}, nil
}
