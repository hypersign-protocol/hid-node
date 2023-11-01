package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
	"github.com/hypersign-protocol/hid-node/x/ssi/verification"
)

// RegisterDID is a RPC method for registration of a DID Document
func (k msgServer) RegisterDID(goCtx context.Context, msg *types.MsgRegisterDID) (*types.MsgRegisterDIDResponse, error) {
	// Unwrap Go Context to Cosmos SDK Context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the RPC inputs
	msgDidDocument := msg.DidDocument
	msgDidDocumentProofs := msg.DidDocumentProofs

	// Validate DID Document
	if err := msgDidDocument.ValidateDidDocument(); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	// Validate namespace in DID Document
	chainNamespace := k.GetChainNamespace(&ctx)
	if err := types.DidChainNamespaceValidation(msgDidDocument, chainNamespace); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	// Validate ownership of method specific id
	if err := checkMethodSpecificIdOwnership(msgDidDocument.VerificationMethod, msgDidDocument.Id); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	// Checks if the Did Document is already registered
	if k.hasDidDocument(ctx, msgDidDocument.Id) {
		return nil, sdkerrors.Wrap(types.ErrDidDocExists, msgDidDocument.Id)
	}

	// Validate Document Proofs
	for _, proof := range msgDidDocumentProofs {
		if err := proof.Validate(); err != nil {
			return nil, err
		}
	}

	// Check if any of the blockchainAccountId is present in any registered DID Document. If so, throw error
	for _, vm := range msgDidDocument.VerificationMethod {
		if vm.BlockchainAccountId != "" {
			if existingDidDocId := k.getBlockchainAddressFromStore(&ctx, vm.BlockchainAccountId); len(existingDidDocId) != 0 {
				return nil, sdkerrors.Wrapf(
					types.ErrInvalidDidDoc,
					"blockchainAccountId %v of verification method %v is already part of DID Document %v",
					vm.BlockchainAccountId,
					vm.Id,
					string(existingDidDocId),
				)
			}
		}
	}

	// Get the list of controllers and check the non-subject controller's existance in the state
	controllerList := getControllersForCreateDID(msgDidDocument)

	if err := k.checkControllerPresenceInState(ctx, controllerList, msgDidDocument.Id); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	// Collect necessary Verification Methods which are needed to be valid
	requiredVMs, err := getVerificationMethodsForCreateDID(msgDidDocument)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrVerificationMethodNotFound, err.Error())
	}

	// Associate Signatures
	signMap := makeSignatureMap(msgDidDocumentProofs)

	requiredVmMap, err := k.formMustControllerVmListMap(ctx, controllerList, requiredVMs, signMap)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}

	// Verify Signatures
	err = verification.VerifySignatureOfEveryController(msgDidDocument, requiredVmMap)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSignature, err.Error())
	}

	// Formt DID Document Metadata
	metadata := types.CreateNewMetadata(ctx)

	// Form the Completet DID Document
	didDocumentState := types.DidDocumentState{
		DidDocument:         msgDidDocument,
		DidDocumentMetadata: &metadata,
	}

	// Register DID Document in Store once all validation checks are passed
	// and increment the DID Document count
	k.setDidDocumentInStore(ctx, &didDocumentState)
	k.incrementDidCount(ctx)

	// After successful registration of the DID Document, every blockchainAccountIds
	// can be added to the store
	for _, vm := range didDocumentState.DidDocument.VerificationMethod {
		if vm.BlockchainAccountId != "" {
			k.setBlockchainAddressInStore(&ctx, vm.BlockchainAccountId, vm.Controller)
		}
	}

	// Emit a successful DID Document Registration event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("create_did", sdk.NewAttribute("tx_author", msg.GetTxAuthor())),
	)

	return &types.MsgRegisterDIDResponse{}, nil
}

// checkMethodSpecificIdOwnership validates the ownership of blockchain account id passed in the method specific
// identifier of DID Document. This ensures that a DID ID (containing a blockchain address) is being created by someone
// who owns the blockchain address.
func checkMethodSpecificIdOwnership(verificationMethods []*types.VerificationMethod, didId string) error {
	inputMSI, inputMSIType, err := types.GetMethodSpecificIdAndType(didId)
	if err != nil {
		return err
	}

	if inputMSIType == types.MSIBlockchainAccountId {
		foundMSIinAnyVM := false
		for _, vm := range verificationMethods {
			if vm.BlockchainAccountId == inputMSI {
				foundMSIinAnyVM = true
				break
			}
		}

		if !foundMSIinAnyVM {
			return fmt.Errorf(
				"proof of ownership for method-specific-id in %v must be provided",
				didId,
			)
		} else {
			return nil
		}
	} else {
		return nil
	}
}

// getControllersForCreateDID returns a list of controller DIDs
// from controller and verification method attributes
func getControllersForCreateDID(didDocument *types.DidDocument) []string {
	var controllerList []string

	// DID Subject is assumed to be the DID Controller if the controller list is empty
	if len(didDocument.Controller) == 0 {
		controllerList = append(controllerList, didDocument.Id)
	}

	controllerList = append(controllerList, didDocument.Controller...)

	for _, vm := range didDocument.VerificationMethod {
		controllerList = append(controllerList, vm.Controller)
	}

	return types.GetUniqueElements(controllerList)
}

// getVerificationMethodsForCreateDID fetches all the Verification Methods needed to be verified
func getVerificationMethodsForCreateDID(didDocument *types.DidDocument) ([]*types.VerificationMethod, error) {
	var mustHaveVerificaitonMethods []*types.VerificationMethod = []*types.VerificationMethod{}
	var foundAtleastOneSubjectVM bool = false

	for _, vm := range didDocument.VerificationMethod {
		// set foundAtleastOneSubjectVM to true if atleast one Subject DID's VM is found
		if vm.Controller == didDocument.Id {
			foundAtleastOneSubjectVM = true
		}

		// Skip X25519KeyAgreementKey2020 or X25519KeyAgreementKey2020 because these
		// are not allowed for Authentication and Assertion purposes
		if (vm.Type == types.X25519KeyAgreementKey2020) || (vm.Type == types.X25519KeyAgreementKeyEIP5630) {
			continue
		}

		mustHaveVerificaitonMethods = append(mustHaveVerificaitonMethods, vm)
	}

	if !foundAtleastOneSubjectVM && utils.FindInSlice(didDocument.Controller, didDocument.Id) {
		return nil, fmt.Errorf(
			"there should be atleast one verification method of DID Subject %v", didDocument.Id)
	}

	return mustHaveVerificaitonMethods, nil
}
