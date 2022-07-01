package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	verify "github.com/hypersign-protocol/hid-node/x/ssi/keeper/document_verification"
)

func (k msgServer) CreateSchema(goCtx context.Context, msg *types.MsgCreateSchema) (*types.MsgCreateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schemaMsg := msg.GetSchema()
	schemaID := schemaMsg.GetId()

	// Get the Did Document of Schema's Author
	authorDidDocument, err := k.GetDid(&ctx, schemaMsg.GetAuthor())
	if err != nil {
		return nil, sdkerrors.Wrap(err, fmt.Sprintf("The DID %s is not available", schemaMsg.GetAuthor()))
	}

	// Check if author's DID is deactivated
	if authorDidDocument.Metadata.Deactivated {
		return nil, sdkerrors.Wrap(types.ErrDidDocDeactivated, fmt.Sprintf("%s is deactivated and cannot used be used to create schema", authorDidDocument.Did.Id))
	}

	// Check if Schema ID is valid
	authorDid := authorDidDocument.GetDid().GetId()
	if err := verify.IsValidSchemaID(schemaID, authorDid); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidSchemaID, err.Error())
	}

	// Check if Schema already exists
	if k.HasSchema(ctx, schemaID) {
		return nil, sdkerrors.Wrap(types.ErrSchemaExists, fmt.Sprintf("Schema ID:  %s", schemaID))
	}

	// Signature check
	didSigners := authorDidDocument.GetDid().GetSigners()
	if err := k.VerifySignatureOnCreateSchema(&ctx, schemaMsg, didSigners, msg.GetSignatures()); err != nil {
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
