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
	chainNamespace := k.GetChainNamespace(&ctx)
	if err := types.DidChainNamespaceValidation(msgDidDocument, chainNamespace); err != nil {
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

	// Check if the DID Document is already deactivated
	if existingDidDocumentState.DidDocumentMetadata.Deactivated {
		return nil, sdkerrors.Wrapf(types.ErrDidDocDeactivated, "cannot update didDocument %v as it is deactivated", existingDidDocument.Id)
	}

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

	signMap := makeSignatureMap(msgSignatures)

	// Check if there is any change in controllers
	var optionalVmMap map[string][]*types.ExtendedVerificationMethod
	var requiredVmMap map[string][]*types.ExtendedVerificationMethod
	var vmMapErr error

	existingDidDocumentControllers := existingDidDocument.Controller
	incomingDidDocumentControllers := msgDidDocument.Controller

	// Assume DID Subject as controller if the controller array is empty
	if len(existingDidDocumentControllers) == 0 {
		existingDidDocumentControllers = append(existingDidDocumentControllers, existingDidDocument.Id)
	}
	if len(incomingDidDocumentControllers) == 0 {
		incomingDidDocumentControllers = append(incomingDidDocumentControllers, msgDidDocument.Id)
	}

	vmsToBeRemoved := []*types.VerificationMethod{}
	vmsToBeAdded := []*types.VerificationMethod{}

	// check if both controller arrays are equal
	if reflect.DeepEqual(existingDidDocumentControllers, incomingDidDocumentControllers) {
		commonController := existingDidDocumentControllers
		if err := k.checkControllerPresenceInState(ctx, commonController, msgDidDocument.Id); err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
		}

		// Check if verification Methods are similar
		if reflect.DeepEqual(existingDidDocument.VerificationMethod, msgDidDocument.VerificationMethod) {
			optionalVmMap, vmMapErr = k.formAnyControllerVmListMap(ctx, commonController, existingDidDocument.VerificationMethod, signMap)
			if vmMapErr != nil {
				return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, vmMapErr.Error())
			}
		} else {
			// Get a list of Verification Methods having a populated `blockchainAccountId` field, which are being newly added
			// and/or removed from the DID Document
			vmsToBeAdded, vmsToBeRemoved, err = processBlockchainAccountIdForUpdateDID(k, ctx, existingDidDocument.VerificationMethod, msgDidDocument.VerificationMethod)
			if err != nil {
				return nil, err
			}

			// if Vms are not similar
			// Get the distinct VMs (new)
			updatedVms := getVerificationMethodsForUpdateDID(existingDidDocument.VerificationMethod, msgDidDocument.VerificationMethod)

			for _, vm := range updatedVms {
				if _, signInfoProvided := signMap[vm.Id]; !signInfoProvided {
					return nil, sdkerrors.Wrapf(
						types.ErrInvalidSignature,
						"signature must be provided for verification method id %v",
						vm.Id,
					)
				}
				vmExtended := types.CreateExtendedVerificationMethod(vm, signMap[vm.Id])
				if requiredVmMap == nil {
					requiredVmMap = map[string][]*types.ExtendedVerificationMethod{}
				}
				requiredVmMap[vm.Controller] = append(requiredVmMap[vm.Controller], vmExtended)
				delete(signMap, vm.Id)
			}

			optionalVmMap, vmMapErr = k.formAnyControllerVmListMap(ctx, commonController, existingDidDocument.VerificationMethod, signMap)
			if err != vmMapErr {
				return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, vmMapErr.Error())
			}
		}
	} else {
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

		requiredVmMap, vmMapErr = k.formMustControllerVmListMap(ctx, mandatoryControllers, updatedVms, signMap)
		if vmMapErr != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, vmMapErr.Error())
		}

		optionalVmMap, vmMapErr = k.formAnyControllerVmListMap(ctx, anyControllers, existingDidDocument.VerificationMethod, signMap)
		if err != vmMapErr {
			return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, vmMapErr.Error())
		}
	}

	// Signature Verification
	if err := verification.VerifySignatureOfEveryController(msgDidDocument, requiredVmMap); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSignature, err.Error())
	}

	if err := verification.VerifySignatureOfAnyController(msgDidDocument, optionalVmMap); err != nil {
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

	// Iterate through the removed Verification Methods having `blockchainAccountId` populated
	// and remove them from store
	for _, vm := range vmsToBeRemoved {
		k.RemoveBlockchainAddressInStore(&ctx, vm.BlockchainAccountId)
	}

	// Iterate through the added Verification Methods having `blockchainAccountId` populated
	// and add them to store
	for _, vm := range vmsToBeAdded {
		k.SetBlockchainAddressInStore(&ctx, vm.BlockchainAccountId, vm.Controller)
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
		// Skip X25519KeyAgreementKey2020 or X25519KeyAgreementKey2020 because these
		// are not allowed for Authentication and Assertion purposes
		if ((vm.Type == types.X25519KeyAgreementKey2020) || (vm.Type == types.X25519KeyAgreementKeyEIP5630)) {
			continue
		}
		existingVmMap[vm.Id] = vm
	}

	// Make map of incoming VMs
	incomingVmMap := map[string]*types.VerificationMethod{}
	for _, vm := range incomingVMs {
		incomingVmMap[vm.Id] = vm
		// Check if VM is present in existing VM map.
		// If it's not present, the VM is being added to existing Did Document.
		// Add the VM to "required" group
		if _, present := existingVmMap[vm.Id]; !present && ((vm.Type != types.X25519KeyAgreementKey2020) && (vm.Type != types.X25519KeyAgreementKeyEIP5630))  {
			updatedVms = append(
				updatedVms,
				vm,
			)
		}
	}

	return updatedVms
}

func processBlockchainAccountIdForUpdateDID(k msgServer, ctx sdk.Context, existingVMs []*types.VerificationMethod, incomingVMs []*types.VerificationMethod) ([]*types.VerificationMethod, []*types.VerificationMethod, error) {
	newVms := []*types.VerificationMethod{}
	deletedVms := []*types.VerificationMethod{}

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
			if vm.BlockchainAccountId != "" {
				if didIdBytes := k.GetBlockchainAddressFromStore(&ctx, vm.BlockchainAccountId); len(didIdBytes) != 0 {
					return nil, nil, fmt.Errorf(
						"blockchainAccountId %v of verification method %v is already part of DID Document %v",
						vm.BlockchainAccountId,
						vm.Id,
						string(didIdBytes),
					)
				} else {
					newVms = append(newVms, vm)
				}
			}
		}
	}

	// Get the list of VMs that are being removed
	for _, vm := range existingVMs {
		if _, present := incomingVmMap[vm.Id]; !present {
			if vm.BlockchainAccountId != "" {
				deletedVms = append(deletedVms, vm)
			}
		}
	}

	return newVms, deletedVms, nil
}
