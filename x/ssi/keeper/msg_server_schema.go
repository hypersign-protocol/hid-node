package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	docVerify "github.com/hypersign-protocol/hid-node/x/ssi/keeper/document_verification"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k msgServer) CreateSchema(goCtx context.Context, msg *types.MsgCreateSchema) (*types.MsgCreateSchemaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	schemaDoc := msg.GetSchemaDoc()
	schemaProof := msg.GetSchemaProof()
	schemaID := schemaDoc.GetId()

	chainNamespace := k.GetChainNamespace(&ctx)

	// Get the Did Document of Schema's Author
	authorDidDocument, err := k.GetDid(&ctx, schemaDoc.GetAuthor())
	if err != nil {
		return nil, sdkerrors.Wrap(err, fmt.Sprintf("The DID %s is not available", schemaDoc.GetAuthor()))
	}

	// Check if author's DID is deactivated
	if authorDidDocument.DidDocumentMetadata.Deactivated {
		return nil, sdkerrors.Wrap(types.ErrDidDocDeactivated, fmt.Sprintf("%s is deactivated and cannot be used to create schema", authorDidDocument.DidDocument.Id))
	}

	// Check if Schema ID is valid
	err = docVerify.IsValidID(schemaID, chainNamespace, "schemaDocument")
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidSchemaID, err.Error())
	}

	// Check if Schema already exists
	if k.HasSchema(ctx, schemaID) {
		return nil, sdkerrors.Wrap(types.ErrSchemaExists, fmt.Sprintf("Schema ID:  %s", schemaID))
	}

	// Check proper date syntax for `authored` and `created`
	blockTime := ctx.BlockTime()

	authoredDate := schemaDoc.GetAuthored()
	authoredDateParsed, err := time.Parse(time.RFC3339, authoredDate)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid authored date format: %s", authoredDateParsed))
	}
	if authoredDateParsed.After(blockTime) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, "authored date provided shouldn't be greater than the current block time")
	}

	createdDate := schemaProof.GetCreated()
	createdDateParsed, err := time.Parse(time.RFC3339, createdDate)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid created date format: %s", createdDateParsed))
	}
	if createdDateParsed.After(blockTime) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, "created date provided shouldn't be greater than the current block time")
	}

	// Signature check
	signature := &types.SignInfo{
		VerificationMethodId: schemaProof.VerificationMethod,
		Signature:            schemaProof.ProofValue,
	}
	signatures := []*types.SignInfo{signature}
	if err := k.VerifyDocumentSignature(&ctx, schemaDoc, authorDidDocument.DidDocument, signatures); err != nil {
		return nil, err
	}

	var schema = types.Schema{
		Type:         schemaDoc.GetType(),
		ModelVersion: schemaDoc.GetModelVersion(),
		Id:           schemaDoc.GetId(),
		Name:         schemaDoc.GetName(),
		Author:       schemaDoc.GetAuthor(),
		Authored:     schemaDoc.GetAuthored(),
		Schema:       schemaDoc.GetSchema(),
		Proof:        schemaProof,
	}

	id := k.AppendSchema(ctx, schema)

	return &types.MsgCreateSchemaResponse{Id: id}, nil
}
