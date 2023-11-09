package keeper

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/verification"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// checkControllerPresenceInState checks if VM of subject is present
func (k msgServer) checkControllerPresenceInState(
	ctx sdk.Context, controllers []string,
	didSubject string,
) error {
	for _, controller := range controllers {
		// Skip if controller is the DID subject
		if controller == didSubject {
			continue
		}

		didDoc, err := k.getDidDocumentState(&ctx, controller)
		if err != nil {
			return err
		}

		// Check if they are deactivated
		if didDoc.DidDocumentMetadata.Deactivated {
			return fmt.Errorf("DID document %s is deactivated", didDoc.DidDocument.Id)
		}
	}
	return nil
}

// formMustControllerVmListMap returns a map between controller and a list of Extended Verification Methods, where
// every verification method of every controller needs to be valid
func (k msgServer) formMustControllerVmListMap(ctx sdk.Context,
	controllers []string, verificationMethods []*types.VerificationMethod,
	inputSignMap map[string]*types.DocumentProof,
) (map[string][]*types.ExtendedVerificationMethod, error) {
	var controllerMap map[string][]*types.ExtendedVerificationMethod = map[string][]*types.ExtendedVerificationMethod{}
	var vmMap map[string]*types.VerificationMethod = map[string]*types.VerificationMethod{}

	// Make controller map
	for _, controller := range controllers {
		controllerMap[controller] = []*types.ExtendedVerificationMethod{}
	}

	// Make Verification method map
	for _, vm := range verificationMethods {
		vmMap[vm.Id] = vm
	}

	for _, vm := range verificationMethods {
		// Check if the required verification method is present in SignMap
		if _, present := inputSignMap[vm.Id]; !present {
			return nil, fmt.Errorf("signature required for verification method %s", vm.Id)
		} else {
			if _, presentInControllerMap := controllerMap[vm.Controller]; presentInControllerMap {
				vmExtended := types.CreateExtendedVerificationMethod(vm, inputSignMap[vm.Id])
				controllerMap[vm.Controller] = append(controllerMap[vm.Controller], vmExtended)
			}
		}
	}

	// Append only those signatures whose controller are present in the controllerMap
	for vmId, sign := range inputSignMap {
		controller, _ := types.SplitDidUrl(vmId)

		if _, present := controllerMap[controller]; present {
			// Check if the VM is present in Subject DID
			_, presentInSubjectDidDoc := vmMap[vmId]

			if presentInSubjectDidDoc {
				_, presentInControllerMap := controllerMap[vmMap[vmId].Controller]
				if presentInControllerMap {
					vmExtended := types.CreateExtendedVerificationMethod(vmMap[vmId], sign)
					controllerMap[controller] = append(controllerMap[controller], vmExtended)
					delete(inputSignMap, vmId)
				}
				// Check for VM from the respective controller's DID Doc
			} else {
				vmState, err := k.getControllerVmFromState(ctx, vmId)
				if err != nil {
					return nil, err
				}
				_, presentInControllerMap := controllerMap[vmState.Controller]
				if presentInControllerMap {
					// Skip X25519KeyAgreementKey2020 or X25519KeyAgreementKey2020 because these
					// are not allowed for Authentication and Assertion purposes
					if (vmState.Type != types.X25519KeyAgreementKey2020) && (vmState.Type != types.X25519KeyAgreementKeyEIP5630) {
						vmExtended := types.CreateExtendedVerificationMethod(vmState, sign)
						controllerMap[controller] = append(controllerMap[controller], vmExtended)
					}
					delete(inputSignMap, vmId)
				}
			}
		}
	}
	return controllerMap, nil
}

// formAnyControllerVmListMap returns a map between controller and a list of Extended Verification Methods, where
// atleast one verification method of any controller needs to be valid
func (k msgServer) formAnyControllerVmListMap(ctx sdk.Context,
	controllers []string, verificationMethods []*types.VerificationMethod,
	inputSignMap map[string]*types.DocumentProof,
) (map[string][]*types.ExtendedVerificationMethod, error) {
	var controllerMap map[string][]*types.ExtendedVerificationMethod = map[string][]*types.ExtendedVerificationMethod{}
	var vmMap map[string]*types.VerificationMethod = map[string]*types.VerificationMethod{}

	// Make controller map
	for _, controller := range controllers {
		controllerMap[controller] = []*types.ExtendedVerificationMethod{}
	}

	// Make Verification method map
	for _, vm := range verificationMethods {
		vmMap[vm.Id] = vm
	}

	// Append only those signatures whose controller are present in the controllerMap
	for vmId, sign := range inputSignMap {
		controller, _ := types.SplitDidUrl(vmId)

		if _, present := controllerMap[controller]; present {
			// Check if the VM is present in Subject DID
			_, presentInSubjectDidDoc := vmMap[vmId]

			if presentInSubjectDidDoc {
				_, presentInControllerMap := controllerMap[vmMap[vmId].Controller]
				if presentInControllerMap {
					vmExtended := types.CreateExtendedVerificationMethod(vmMap[vmId], sign)
					controllerMap[controller] = append(controllerMap[controller], vmExtended)
				}
				// Check for VM from the respective controller's DID Doc
			} else {
				vmState, err := k.getControllerVmFromState(ctx, vmId)
				if err != nil {
					return nil, err
				}
				_, presentInControllerMap := controllerMap[vmState.Controller]
				if presentInControllerMap {
					// Skip X25519KeyAgreementKey2020 or X25519KeyAgreementKey2020 because these
					// are not allowed for Authentication and Assertion purposes
					if (vmState.Type != types.X25519KeyAgreementKey2020) && (vmState.Type != types.X25519KeyAgreementKeyEIP5630) {
						vmExtended := types.CreateExtendedVerificationMethod(vmState, sign)
						controllerMap[controller] = append(controllerMap[controller], vmExtended)
					}
				}
			}
		}
	}
	return controllerMap, nil
}

func (k msgServer) getControllerVmFromState(ctx sdk.Context, verificationMethodId string) (*types.VerificationMethod, error) {
	didId, _ := types.SplitDidUrl(verificationMethodId)

	didDocumentState, err := k.getDidDocumentState(&ctx, didId)
	if err != nil {
		return nil, err
	}

	for _, vm := range didDocumentState.DidDocument.VerificationMethod {
		if vm.Id == verificationMethodId {
			return vm, nil
		}
	}

	return nil, fmt.Errorf("verification method %v not found in controller %v", verificationMethodId, didId)
}

// VerifyDocumentProof verifies the proof of a SSI Document
func (k msgServer) VerifyDocumentProof(ctx sdk.Context, ssiMsg types.SsiMsg, inputDocProof types.SSIProofInterface) error {
	// Get DID Document from State
	docProofVmId := inputDocProof.GetVerificationMethod()
	didId, _ := types.SplitDidUrl(docProofVmId)
	didDocumentState, err := k.getDidDocumentState(&ctx, didId)
	if err != nil {
		return err
	}
	didDoc := didDocumentState.DidDocument

	// Search for Verification Method in DID Document
	var docVm *types.VerificationMethod = nil
	for _, vm := range didDoc.VerificationMethod {
		if vm.Id == docProofVmId {
			docVm = vm
			break
		}
	}
	if docVm == nil {
		return fmt.Errorf("verificationMethod %s is not present in DID document %s", docProofVmId, didId)
	}

	// VerificationKeySignatureMap has X25519KeyAgreementKey2020 and X25519KeyAgreementKeyEIP5630 as supported Verification Type.
	// However, they are not allowed to be used for Authentication or Assertion purposes. Since, their corresponding values in the map
	// are empty string, the following check is in place.
	if types.VerificationKeySignatureMap[docVm.Type] == "" {
		return fmt.Errorf("proof type must be specified")
	}

	// Check if the Proof Type is correct
	if types.VerificationKeySignatureMap[docVm.Type] != inputDocProof.GetType() {
		return fmt.Errorf(
			"expected proof type to be %v as the verificationMethod type of %v is %v, recieved %v",
			types.VerificationKeySignatureMap[docVm.Type],
			docVm.Id,
			docVm.Type,
			inputDocProof.GetType(),
		)
	}

	// Verify signature
	documentProof := &types.DocumentProof{
		VerificationMethod: inputDocProof.GetVerificationMethod(),
		ProofValue:         inputDocProof.GetProofValue(),
		ClientSpecType:     inputDocProof.GetClientSpecType(),
	}
	err = verification.VerifyDocumentProofSignature(ssiMsg, docVm, documentProof)
	if err != nil {
		return err
	}

	return nil
}

// makeSignatureMap converts []SignInfo to map
func makeSignatureMap(inputSignatures []*types.DocumentProof) map[string]*types.DocumentProof {
	var signMap map[string]*types.DocumentProof = map[string]*types.DocumentProof{}

	for _, sign := range inputSignatures {
		signMap[sign.VerificationMethod] = sign
	}

	return signMap
}

// verifyCredentialMerkleRootHash verifies the MerkleRootHash of Credential Status
func verifyCredentialMerkleRootHash(credHash string) error {
	decodedCredentialHash, err := hex.DecodeString(credHash)
	if err != nil {
		return fmt.Errorf("merkle root hash %v is an invalid hex string", credHash)
	}

	if len(decodedCredentialHash) != 32 {
		return fmt.Errorf(
			"expected merkle root hash %v to be of length %v, got %v",
			credHash,
			32,
			len(credHash),
		)
	}

	return nil
}
