package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/verification"
)

// RPC controller for registering DID document on hid-node
func (k msgServer) CreateDID(goCtx context.Context, msg *types.MsgCreateDID) (*types.MsgCreateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chainNamespace := k.GetChainNamespace(&ctx)

	msgDidDocument := msg.GetDidDocString()
	didId := msgDidDocument.GetId()
	didSigners := msgDidDocument.GetSigners()
	// Get the Verification Method for controller DIDs
	didSignersWithVM, err := k.GetVMForSigners(&ctx, didSigners)
	if err != nil {
		return nil, err
	}

	msgSignatures := msg.GetSignatures()
	signerAddress := msg.GetCreator()

	// Checks if the Did Document is valid
	err = verification.ValidateDidDocument(msgDidDocument, chainNamespace)
	if err != nil {
		return nil, err
	}

	// Checks if the Did Document is already registered
	if k.HasDid(ctx, didId) {
		return nil, sdkerrors.Wrap(types.ErrDidDocExists, fmt.Sprintf("DID already exists %s", didId))
	}

	// Checks if the Controllers are valid
	didControllers := msgDidDocument.GetController()
	didVerificationMethod := msgDidDocument.GetVerificationMethod()
	if k.ValidateDidControllers(&ctx, didId, didControllers, didVerificationMethod) != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, "DID controller is not valid")
	}

	// Verification of Did Document Signature

	// ClientSpec check
	var didDocBytes []byte
	didDocBytes = msgDidDocument.GetSignBytes()

	msgClientSpecType := msg.GetClientSpec()
	clientSpecOpts := types.ClientSpecOpts{
		SSIDocBytes:   didDocBytes,
		SignerAddress: signerAddress,
	}

	didDocBytes, err = getClientSpecDocBytes(msgClientSpecType, clientSpecOpts)
	if err != nil {
		return nil, err
	}

	if err := verification.VerifyDidSignature(&ctx, didDocBytes, didSignersWithVM, msgSignatures); err != nil {
		return nil, err
	}

	// Create the Metadata
	metadata := types.CreateNewMetadata(ctx)

	// Form the Completet DID Document
	didDocumentState := types.DidDocumentState{
		DidDocument:         msgDidDocument,
		DidDocumentMetadata: &metadata,
	}

	// Register DID Document in Store once all validation checks are passed
	id := k.RegisterDidDocumentInStore(ctx, &didDocumentState)

	return &types.MsgCreateDIDResponse{Id: id}, nil
}

// RPC controller for updating an existing DID document registered on hid-node
func (k msgServer) UpdateDID(goCtx context.Context, msg *types.MsgUpdateDID) (*types.MsgUpdateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	chainNamespace := k.GetChainNamespace(&ctx)

	msgDidDocument := msg.GetDidDocString()
	didId := msgDidDocument.GetId()

	msgVersionId := msg.GetVersionId()

	// Check if the input DID Document is valid
	err := verification.ValidateDidDocument(msgDidDocument, chainNamespace)
	if err != nil {
		return nil, err
	}

	// Checks whether the DID Document exists in the store
	if !k.HasDid(ctx, didId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("DID doesnt exists %s", didId))
	}

	// Get the registered Did Document from store
	oldDidDocumentState, err := k.GetDidDocumentState(&ctx, didId)
	if err != nil {
		return nil, err
	}
	oldDidDocument := oldDidDocumentState.GetDidDocument()
	oldMetaData := oldDidDocumentState.GetDidDocumentMetadata()

	// Check if the status of DID Document is deactivated
	if err := verification.VerifyDidDeactivate(oldMetaData, didId); err != nil {
		return nil, err
	}

	// Check if the version id of existing Did Document matches with the current one
	if oldMetaData.VersionId != msgVersionId {
		errMsg := fmt.Sprintf("Expected %s with version %s. Got version %s", msgDidDocument.Id, oldMetaData.VersionId, msgVersionId)
		return nil, sdkerrors.Wrap(types.ErrUnexpectedDidVersion, errMsg)
	}

	// Check if the controllers are valid
	didControllers := msgDidDocument.GetController()
	didVerificationMethod := msgDidDocument.GetVerificationMethod()
	if k.ValidateDidControllers(&ctx, didId, didControllers, didVerificationMethod) != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, "DID controller is not valid")
	}

	// Validate Signatures
	signatures := msg.GetSignatures()
	signers := GetUpdatedSigners(&ctx, oldDidDocument, msgDidDocument, signatures)
	signersWithVm, err := k.GetVMForSigners(&ctx, signers)
	if err != nil {
		return nil, err
	}

	// ClientSpec check
	var didDocBytes []byte
	didDocBytes = msgDidDocument.GetSignBytes()

	clientSpecType := msg.ClientSpec
	signerAddress := msg.GetCreator()
	clientSpecOpts := types.ClientSpecOpts{
		SSIDocBytes:   didDocBytes,
		SignerAddress: signerAddress,
	}

	didDocBytes, err = getClientSpecDocBytes(clientSpecType, clientSpecOpts)
	if err != nil {
		return nil, err
	}

	if err := verification.VerifyDidSignature(&ctx, didDocBytes, signersWithVm, msg.Signatures); err != nil {
		return nil, err
	}

	// Create the Metadata and assign `created` and `deactivated` to previous DIDDoc's metadata values
	metadata := types.CreateNewMetadata(ctx)
	metadata.Created = oldDidDocumentState.GetDidDocumentMetadata().GetCreated()
	metadata.Deactivated = oldDidDocumentState.GetDidDocumentMetadata().GetDeactivated()

	// Form the DID Document
	didDocumentState := types.DidDocumentState{
		DidDocument:         msgDidDocument,
		DidDocumentMetadata: &metadata,
	}

	// Update the DID Document in store
	if err := k.UpdateDidDocumentInStore(ctx, didDocumentState); err != nil {
		return nil, err
	}

	return &types.MsgUpdateDIDResponse{UpdateId: didId}, nil
}

// RPC controller for deactivating an existing DID document registered on hid-node
func (k msgServer) DeactivateDID(goCtx context.Context, msg *types.MsgDeactivateDID) (*types.MsgDeactivateDIDResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chainNamespace := k.GetChainNamespace(&ctx)

	msgDidId := msg.GetDidId()
	msgVersionId := msg.GetVersionId()

	// Check if the Did id format is valid
	err := verification.IsValidID(msgDidId, chainNamespace, "didDocument")
	if err != nil {
		return nil, err
	}

	// Checks whether the DID Document exists in the store
	if !k.HasDid(ctx, msgDidId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("DID doesnt exists %s", msgDidId))
	}

	// Retrieve the DID Document from store
	didDocumentState, err := k.GetDidDocumentState(&ctx, msgDidId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch DID Document with Did Id %s from store", msgDidId)
	}
	didDoc := didDocumentState.GetDidDocument()
	didDocMetadata := didDocumentState.GetDidDocumentMetadata()
	oldVersionId := didDocMetadata.GetVersionId()

	// Check if the DID is already deactivated
	if err := verification.VerifyDidDeactivate(didDocMetadata, msgDidId); err != nil {
		return nil, err
	}

	// Check if the input version id is similar to registered DID Document's version id
	if oldVersionId != msgVersionId {
		errMsg := fmt.Sprintf("Expected %s with version %s. Got version %s", msgDidId, oldVersionId, msgVersionId)
		return nil, sdkerrors.Wrap(types.ErrUnexpectedDidVersion, errMsg)
	}

	// Check the validity of DID Controllers
	didControllers := didDoc.GetController()
	didVerificationMethod := didDoc.GetVerificationMethod()
	if k.ValidateDidControllers(&ctx, msgDidId, didControllers, didVerificationMethod) != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, "DID controller is not valid")
	}

	// Signature Verification
	signers := didDoc.GetSigners()
	// Get the Verification Method for controller DIDs
	signersWithVM, err := k.GetVMForSigners(&ctx, signers)
	if err != nil {
		return nil, err
	}
	signatures := msg.GetSignatures()

	// ClientSpec check
	var didDocBytes []byte
	didDocBytes = didDoc.GetSignBytes()
	msgClientSpecType := msg.GetClientSpec()
	msgSignerAddress := msg.GetCreator()

	clientSpecOpts := types.ClientSpecOpts{
		SSIDocBytes:   didDocBytes,
		SignerAddress: msgSignerAddress,
	}

	didDocBytes, err = getClientSpecDocBytes(msgClientSpecType, clientSpecOpts)
	if err != nil {
		return nil, err
	}

	if err := verification.VerifyDidSignature(&ctx, didDocBytes, signersWithVM, signatures); err != nil {
		return nil, err
	}

	// Create updated metadata
	updatedMetadata := types.CreateNewMetadata(ctx)
	updatedMetadata.Created = didDocumentState.GetDidDocumentMetadata().GetCreated()
	updatedMetadata.Deactivated = true

	// Form the updated DID Document
	updatedDidDocumentState := types.DidDocumentState{
		DidDocument:         didDoc,
		DidDocumentMetadata: &updatedMetadata,
	}

	// Update the DID Document in Store
	if err := k.UpdateDidDocumentInStore(ctx, updatedDidDocumentState); err != nil {
		return nil, err
	}

	return &types.MsgDeactivateDIDResponse{Id: 1}, nil
}
