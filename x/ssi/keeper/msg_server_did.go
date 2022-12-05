package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	docVerify "github.com/hypersign-protocol/hid-node/x/ssi/document_verification"
	"github.com/hypersign-protocol/hid-node/x/ssi/signature"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// RPC controller for registering DID document on hid-node
func (k msgServer) CreateDID(goCtx context.Context, msg *types.MsgCreateDID) (*types.MsgCreateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	didMsg := msg.GetDidDocString()
	didId := didMsg.GetId()
	chainNamespace := k.GetChainNamespace(&ctx)
	didSigners := didMsg.GetSigners()
	didSignersWithVM, err := k.GetVMForSigners(&ctx, didSigners)
	if err != nil {
		return nil, err
	}
	signatures := msg.GetSignatures()

	// Checks if the Did Document is valid
	err = docVerify.ValidateDidDocument(msg.DidDocString, chainNamespace)
	if err != nil {
		return nil, err
	}

	// Checks if the Did Document is already registered
	if k.HasDid(ctx, didId) {
		return nil, sdkerrors.Wrap(types.ErrDidDocExists, fmt.Sprintf("DID already exists %s", didId))
	}

	// Checks if the Controllers are valid
	didController := didMsg.GetController()
	didVerificationMethod := didMsg.GetVerificationMethod()
	if k.ValidateDidControllers(&ctx, didId, didController, didVerificationMethod) != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, "DID controller is not valid")
	}

	// Verification of Did Document Signature
	if err := signature.VerifyDidSignature(&ctx, didMsg, didSignersWithVM, signatures); err != nil {
		return nil, err
	}

	// Create the Metadata
	metadata := types.CreateNewMetadata(ctx)

	// Form the Completet DID Document
	didDocument := types.DidDocumentState{
		DidDocument:         didMsg,
		DidDocumentMetadata: &metadata,
	}

	// Set the DID Document to KVStore
	id := k.CreateDidDocumentInStore(ctx, &didDocument)

	return &types.MsgCreateDIDResponse{Id: id}, nil
}

// RPC controller for updating an existing DID document registered on hid-node
func (k msgServer) UpdateDID(goCtx context.Context, msg *types.MsgUpdateDID) (*types.MsgUpdateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	didMsg := msg.GetDidDocString()
	versionId := msg.GetVersionId()
	didId := didMsg.GetId()
	chainNamespace := k.GetChainNamespace(&ctx)

	// Check if the input DID Document is valid
	err := docVerify.ValidateDidDocument(didMsg, chainNamespace)
	if err != nil {
		return nil, err
	}

	// Checks whether the DID Document exists in the store
	if !k.HasDid(ctx, didId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("DID doesnt exists %s", didId))
	}

	// Retrieve existing Did Document from store
	oldDIDDoc, err := k.GetDid(&ctx, didId)
	if err != nil {
		return nil, err
	}
	oldDid := oldDIDDoc.GetDidDocument()
	oldMetaData := oldDIDDoc.GetDidDocumentMetadata()

	// Check if the status of DID Document is deactivated
	if err := docVerify.VerifyDidDeactivate(oldMetaData, didId); err != nil {
		return nil, err
	}

	// Check if the version id of existing Did Document matches with the current one
	if oldMetaData.VersionId != versionId {
		errMsg := fmt.Sprintf("Expected %s with version %s. Got version %s", didMsg.Id, oldMetaData.VersionId, versionId)
		return nil, sdkerrors.Wrap(types.ErrUnexpectedDidVersion, errMsg)
	}

	// Check if the controllers are valid
	if k.ValidateDidControllers(&ctx, didId, didMsg.GetController(), didMsg.GetVerificationMethod()) != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, "DID controller is not valid")
	}

	// Validate Signatures
	signers := GetUpdatedSigners(&ctx, oldDid, didMsg, msg.Signatures)
	signersWithVm, err := k.GetVMForSigners(&ctx, signers)
	if err != nil {
		return nil, err
	}
	if err := signature.VerifyDidSignature(&ctx, didMsg, signersWithVm, msg.Signatures); err != nil {
		return nil, err
	}

	// Create the Metadata
	metadata := types.CreateNewMetadata(ctx)
	// Assign `created` and `deactivated` to previous DIDDoc's metadata values
	metadata.Created = oldDIDDoc.GetDidDocumentMetadata().GetCreated()
	metadata.Deactivated = oldDIDDoc.GetDidDocumentMetadata().GetDeactivated()

	// Form the DID Document
	didDoc := types.DidDocumentState{
		DidDocument:         didMsg,
		DidDocumentMetadata: &metadata,
	}

	// Update the DID Document in store
	if err := k.UpdateDidDocumentInStore(ctx, didDoc); err != nil {
		return nil, err
	}

	return &types.MsgUpdateDIDResponse{UpdateId: didId}, nil
}

// RPC controller for deactivating an existing DID document registered on hid-node
func (k msgServer) DeactivateDID(goCtx context.Context, msg *types.MsgDeactivateDID) (*types.MsgDeactivateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	didId := msg.GetDidId()
	versionId := msg.GetVersionId()
	chainNamespace := k.GetChainNamespace(&ctx)

	// Check if the Did id format is valid
	err := docVerify.IsValidID(didId, chainNamespace, "didDocument")
	if err != nil {
		return nil, err
	}

	// Checks whether the DID Document exists in the store
	if !k.HasDid(ctx, didId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("DID doesnt exists %s", didId))
	}

	// Retrieve the DID Document from store
	didDocumentState, err := k.GetDid(&ctx, didId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch DID Document with Did Id %s from store", didId)
	}
	didDoc := didDocumentState.GetDidDocument()
	metadata := didDocumentState.GetDidDocumentMetadata()
	oldVersionId := metadata.GetVersionId()

	// Check if the DID is already deactivated
	if err := docVerify.VerifyDidDeactivate(metadata, didId); err != nil {
		return nil, err
	}

	// Check if the versionId passed is the same as the one in the Latest DID Document in store
	if oldVersionId != versionId {
		errMsg := fmt.Sprintf("Expected %s with version %s. Got version %s", didId, oldVersionId, versionId)
		return nil, sdkerrors.Wrap(types.ErrUnexpectedDidVersion, errMsg)
	}

	// Check the validity of DID Controllers
	didController := didDoc.GetController()
	didVerificationMethod := didDoc.GetVerificationMethod()
	if k.ValidateDidControllers(&ctx, didId, didController, didVerificationMethod) != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, "DID controller is not valid")
	}

	// Signature Verification
	signers := didDoc.GetSigners()
	signersWithVM, err := k.GetVMForSigners(&ctx, signers)
	if err != nil {
		return nil, err
	}
	signatures := msg.Signatures
	if err := signature.VerifyDidSignature(&ctx, didDoc, signersWithVM, signatures); err != nil {
		return nil, err
	}

	// Create updated metadata
	updatedMetadata := types.CreateNewMetadata(ctx)
	updatedMetadata.Created = didDocumentState.GetDidDocumentMetadata().GetCreated()
	updatedMetadata.Deactivated = true

	// Form the updated DID Document
	updatedDidDocument := types.DidDocumentState{
		DidDocument:         didDoc,
		DidDocumentMetadata: &updatedMetadata,
	}

	// Update the DID Document in Store
	if err := k.UpdateDidDocumentInStore(ctx, updatedDidDocument); err != nil {
		return nil, err
	}

	return &types.MsgDeactivateDIDResponse{Id: 1}, nil
}
