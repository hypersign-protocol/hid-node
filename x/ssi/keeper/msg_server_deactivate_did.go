package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/verification"
)

// RPC controller for deactivating an existing DID document registered on hid-node
func (k msgServer) DeactivateDID(goCtx context.Context, msg *types.MsgDeactivateDID) (*types.MsgDeactivateDIDResponse, error) {
	// Unwrap Go Context to Cosmos SDK Context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the RPC inputs
	msgDidId := msg.DidId
	msgSignatures := msg.Signatures

	// Checks if the Did Document is already registered
	if !k.HasDid(ctx, msgDidId) {
		return nil, sdkerrors.Wrap(types.ErrDidDocNotFound, msgDidId)
	}

	// Get the DID Document from state
	didDocumentState, err := k.GetDidDocumentState(&ctx, msgDidId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrDidDocNotFound, err.Error())
	}
	didDocument := didDocumentState.DidDocument
	didDocumentMetadata := didDocumentState.DidDocumentMetadata

	// Check if Did Document is deactivated
	if didDocumentMetadata.Deactivated {
		return nil, sdkerrors.Wrapf(types.ErrDidDocDeactivated, "DID Document %s is already deactivated", msgDidId)
	}

	// Check if the version id of existing Did Document matches with the current one
	existingDidDocVersionId := didDocumentMetadata.VersionId
	incomingDidDocVersionId := msg.VersionId
	if existingDidDocVersionId != incomingDidDocVersionId {
		errMsg := fmt.Sprintf(
			"Expected %s with version %s. Got version %s",
			didDocument.Id, existingDidDocVersionId, incomingDidDocVersionId)
		return nil, sdkerrors.Wrap(types.ErrUnexpectedDidVersion, errMsg)
	}

	// Gather controllers
	controllers := getControllersForDeactivateDID(didDocument)
	if err := k.checkControllerPresenceInState(ctx, controllers, didDocument.Id); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	signMap := makeSignatureMap(msgSignatures)

	// Get controller VM map
	controllerMap, err := k.formAnyControllerVmListMap(ctx, controllers,
		didDocument.VerificationMethod, signMap)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	// Signature Verification
	err = verification.VerifySignatureOfAnyController(didDocument, controllerMap)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSignature, err.Error())
	}

	// Create updated metadata
	updatedMetadata := types.CreateNewMetadata(ctx)
	updatedMetadata.Created = didDocumentState.GetDidDocumentMetadata().GetCreated()
	updatedMetadata.Deactivated = true

	// Form the updated DID Document
	updatedDidDocumentState := types.DidDocumentState{
		DidDocument:         didDocument,
		DidDocumentMetadata: &updatedMetadata,
	}

	// Update the DID Document in Store
	if err := k.UpdateDidDocumentInStore(ctx, updatedDidDocumentState); err != nil {
		return nil, err
	}

	// Remove the BlockchainAccountId from BlockchainAddressStore
	for _, vm := range didDocumentState.DidDocument.VerificationMethod {
		if vm.BlockchainAccountId != "" {
			k.RemoveBlockchainAddressInStore(&ctx, vm.BlockchainAccountId)
		}
	}

	return &types.MsgDeactivateDIDResponse{Id: 1}, nil
}

// getControllersForDeactivateDID returns a list of controllers required for Deactivate DID Operation
func getControllersForDeactivateDID(didDocument *types.Did) []string {
	var controllers []string = []string{}

	// If the controller list is empty, DID Subject is assumed to be the sole controller of DID Document
	if len(didDocument.Controller) == 0 {
		controllers = append(controllers, didDocument.Id)
	} else {
		controllers = didDocument.Controller
	}

	return controllers
}
