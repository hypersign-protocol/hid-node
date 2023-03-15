package keeper

import (
	"context"
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/verification"
)

// RPC controller for updating an existing DID document registered on hid-node
func (k msgServer) UpdateDID(goCtx context.Context, msg *types.MsgUpdateDID) (*types.MsgUpdateDIDResponse, error) {
	// Unwrap Go Context to Cosmos SDK Context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the RPC inputs
	msgDidDocument := msg.DidDocString
	msgSignatures := msg.Signatures

	// Validate DID Document
	if err := msgDidDocument.ValidateDidDocument(); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	// Validate namespace in DID Document
	if err := didDocNamespaceValidation(k, ctx, msgDidDocument); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	// Checks if the Did Document is already registered
	if !k.HasDid(ctx, msgDidDocument.Id) {
		return nil, sdkerrors.Wrap(types.ErrDidDocNotFound, msgDidDocument.Id)
	}

	// Fetch registered Did Document from state
	existingDidDocumentState, err := k.GetDidDocumentState(&ctx, msgDidDocument.Id)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrDidDocNotFound, err.Error())
	}
	existingDidDocument := existingDidDocumentState.DidDocument

	// Check if the incoming DID Document has any changes. If not, throw an error.
	if reflect.DeepEqual(existingDidDocument, msgDidDocument) {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, "incoming DID Document does not have any changes")
	}

	// Check if the version id of existing Did Document matches with the current one
	existingDidDocVersionId := existingDidDocumentState.DidDocumentMetadata.VersionId
	incomingDidDocVersionId := msg.VersionId
	if existingDidDocVersionId != incomingDidDocVersionId {
		errMsg := fmt.Sprintf(
			"Expected %s with version %s. Got version %s",
			msgDidDocument.Id, existingDidDocVersionId, incomingDidDocVersionId)
		return nil, sdkerrors.Wrap(types.ErrUnexpectedDidVersion, errMsg)
	}

	// Look for any change in Controller array
	mandatoryControllers, anyControllers := getControllersForUpdateDID(existingDidDocument, msgDidDocument)
	if err := k.checkControllerPresenceInState(ctx, mandatoryControllers, msgDidDocument.Id); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}
	if err := k.checkControllerPresenceInState(ctx, anyControllers, msgDidDocument.Id); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	// Gather Verification Methods
	updatedVms := getVerificationMethodsForUpdateDID(existingDidDocument.VerificationMethod, msgDidDocument.VerificationMethod)

	signMap := makeSignatureMap(msgSignatures)

	requiredVmMap, err := k.formMustControllerVmListMap(ctx, mandatoryControllers, updatedVms, signMap)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	optionalVmMap, err := k.formAnyControllerVmListMap(ctx,
		anyControllers, existingDidDocument.VerificationMethod, signMap)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	// ClientSpec Opts
	clientSpecOpts := types.ClientSpecOpts{
		ClientSpecType: msg.ClientSpec,
		SSIDoc:         msgDidDocument,
		SignerAddress:  msg.Creator,
	}
	var didDocBytes []byte
	didDocBytes, err = getClientSpecDocBytes(clientSpecOpts)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidClientSpecType, err.Error())
	}

	// Signature Verification
	if err := verification.VerifySignatureOfEveryController(didDocBytes, requiredVmMap); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSignature, err.Error())
	}

	if err := verification.VerifySignatureOfAnyController(didDocBytes, optionalVmMap); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSignature, err.Error())
	}

	// Create the Metadata and assign `created` and `deactivated` to previous DIDDoc's metadata values
	metadata := types.CreateNewMetadata(ctx)
	metadata.Created = existingDidDocumentState.GetDidDocumentMetadata().GetCreated()
	metadata.Deactivated = existingDidDocumentState.GetDidDocumentMetadata().GetDeactivated()

	// Form the DID Document
	didDocumentState := types.DidDocumentState{
		DidDocument:         msgDidDocument,
		DidDocumentMetadata: &metadata,
	}

	// Update the DID Document in store
	if err := k.UpdateDidDocumentInStore(ctx, didDocumentState); err != nil {
		return nil, err
	}

	return &types.MsgUpdateDIDResponse{UpdateId: msgDidDocument.Id}, nil
}

// getControllersForUpdateDID returns two lists of controllers. The first parameter is a list of controllers,
// where every controller needs to be verified. The second parameter is a list of controllers, where atleast one
// of the controllers require verification.
func getControllersForUpdateDID(existingDidDoc *types.Did, incomingDidDoc *types.Did) ([]string, []string) {
	var mandatoryControllers []string = []string{}
	var anyControllers []string = []string{}

	// Make map of existing controllers
	existingControllersMap := map[string]bool{}
	if len(existingDidDoc.Controller) == 0 {
		// If the controller list is empty, DID Subject is assumed to be the sole controller of DID Document
		existingControllersMap[existingDidDoc.Id] = true
	} else {
		for _, controller := range existingDidDoc.Controller {
			existingControllersMap[controller] = true
		}
	}

	// Make map of incoming controllers
	incomingControllersMap := map[string]bool{}
	// If the controller list is empty, DID Subject is assumed to be the sole controller of DID Document
	if len(incomingDidDoc.Controller) == 0 {
		incomingControllersMap[incomingDidDoc.Id] = true
		if _, present := existingControllersMap[incomingDidDoc.Id]; !present {
			mandatoryControllers = append(
				mandatoryControllers,
				incomingDidDoc.Id,
			)
		} else {
			anyControllers = append(
				anyControllers,
				incomingDidDoc.Id,
			)
		}
	} else {
		for _, controller := range incomingDidDoc.Controller {
			incomingControllersMap[controller] = true
			// Check if controller present in existing controller map.
			// If it's not present, the controller is being added to existing Did Document.
			// Add the controller to "required" group
			if _, present := existingControllersMap[controller]; !present {
				mandatoryControllers = append(
					mandatoryControllers,
					controller,
				)
			} else {
				anyControllers = append(
					anyControllers,
					controller,
				)
			}
		}
	}

	// Check if any controllers are deleted
	// If so, add them to the "optional" group
	for controller := range existingControllersMap {
		if _, present := incomingControllersMap[controller]; !present {
			anyControllers = append(
				anyControllers,
				controller,
			)
		}
	}

	return mandatoryControllers, anyControllers
}

// getVerificationMethodsForUpdateDID returns a map highlighting inclusion of new Verification methods
// and/or removal of any existing Verification method.
func getVerificationMethodsForUpdateDID(existingVMs []*types.VerificationMethod, incomingVMs []*types.VerificationMethod) []*types.VerificationMethod {
	updatedVms := []*types.VerificationMethod{}

	// Make map of existing VMs
	existingVmMap := map[string]*types.VerificationMethod{}
	for _, vm := range existingVMs {
		existingVmMap[vm.Id] = vm
	}

	// Make map of incoming VMs
	incomingVmMap := map[string]*types.VerificationMethod{}
	for _, vm := range incomingVMs {
		incomingVmMap[vm.Id] = vm
		// Check if VM is present in existing VM map.
		// If it's not present, the VM is being added to existing Did Document.
		// Add the VM to "required" group
		if _, present := existingVmMap[vm.Id]; !present {
			updatedVms = append(
				updatedVms,
				vm,
			)
		}
	}

	return updatedVms
}
