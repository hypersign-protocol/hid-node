package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/verification"
)

// CreateDID is a RPC method for registration of a DID Document
func (k msgServer) CreateDID(goCtx context.Context, msg *types.MsgCreateDID) (*types.MsgCreateDIDResponse, error) {
	// Unwrap Go Context to Cosmos SDK Context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the RPC inputs
	msgDidDocument := msg.DidDocString
	msgSignatures := msg.Signatures

	// Validate DID Document
	if err := msgDidDocument.ValidateDidDocument(); err != nil {
		return nil, err
	}

	// Validate namespace in DID Document
	if err := didNamespaceValidation(k, ctx, msgDidDocument); err != nil {
		return nil, err
	}

	// Checks if the Did Document is already registered
	if k.HasDid(ctx, msgDidDocument.Id) {
		return nil, sdkerrors.Wrap(types.ErrDidDocExists, msgDidDocument.Id)
	}

	// Get the list of controllers and check the non-subject controller's existance in the state
	controllerList := getControllersForCreateDID(msgDidDocument)

	if err := k.checkControllerPresenceInState(ctx, controllerList, msgDidDocument.Id); err != nil {
		return nil, err
	}

	// Collect necessary Verification Methods which are needed to be valid
	requiredVMs, err := getVerificationMethodsForCreateDID(msgDidDocument)
	if err != nil {
		return nil, err
	}

	// Associate Signatures
	signMap := makeSignatureMap(msgSignatures)

	requiredVmMap, err := k.formMustControllerVmListMap(ctx, controllerList, requiredVMs, signMap)
	if err != nil {
		return nil, err
	}

	// ClientSpec check
	clientSpecOpts := types.ClientSpecOpts{
		ClientSpecType: msg.ClientSpec,
		SSIDoc:         msgDidDocument,
		SignerAddress:  msg.Creator,
	}
	var didDocBytes []byte
	didDocBytes, err = getClientSpecDocBytes(clientSpecOpts)
	if err != nil {
		return nil, err
	}

	// Verify Signatures
	err = verification.VerifySignatureOfEveryController(didDocBytes, requiredVmMap)
	if err != nil {
		return nil, err
	}

	// Formt DID Document Metadata
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

// getControllersForCreateDID returns a list of controller DIDs
// from controller and verification method attributes
func getControllersForCreateDID(didDocument *types.Did) []string {
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
func getVerificationMethodsForCreateDID(didDocument *types.Did) ([]*types.VerificationMethod, error) {
	var mustHaveVerificaitonMethods []*types.VerificationMethod = []*types.VerificationMethod{}
	var foundAtleastOneSubjectVM bool = false

	for _, vm := range didDocument.VerificationMethod {
		// set foundAtleastOneSubjectVM to true if atleast one Subject DID's VM is found
		if vm.Controller == didDocument.Id {
			foundAtleastOneSubjectVM = true
		}
		mustHaveVerificaitonMethods = append(mustHaveVerificaitonMethods, vm)
	}

	if !foundAtleastOneSubjectVM && types.FindInSlice(didDocument.Controller, didDocument.Id) {
		return nil, fmt.Errorf(
			"there should be atleast one verification method from subject DID controller %v", didDocument.Id)
	}

	return mustHaveVerificaitonMethods, nil
}
