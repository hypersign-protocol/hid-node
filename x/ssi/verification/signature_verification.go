package verification

import (
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// VerifySignatureOfEveryController verifies every required verification method of every controller
func VerifySignatureOfEveryController(
	didDoc []byte, VmMap map[string][]*types.ExtendedVerificationMethod,
) error {
	for controller, vmList := range VmMap {
		if len(vmList) == 0 {
			return fmt.Errorf("require atleast one valid signature for controller %s", controller)
		}
		err := verifyAll(vmList, didDoc)
		if err != nil {
			return fmt.Errorf("%s: need every signature for controller %s to be valid", err.Error(), controller)
		}
	}
	return nil
}

// VerifySignatureOfEveryController verifies any required0 verification8
func VerifySignatureOfAnyController(
	didDoc []byte, VmMap map[string][]*types.ExtendedVerificationMethod,
) error {
	found := false
	for _, vmList := range VmMap {
		found = verifyAny(vmList, didDoc)
		if found {
			break
		}
	}

	if !found {
		return fmt.Errorf(
			"need atleast one valid signature from any of the existing controllers in the registered didDoc")
	}

	return nil
}

// VerifyDocumentProofSignature verfies the proof of the SSI Document such as Credential Schema and Credential Status
func VerifyDocumentProofSignature(singedData []byte, vm *types.VerificationMethod, singature string) error {
	vmExtended := types.CreateExtendedVerificationMethod(vm, singature)
	if err := verify(vmExtended, singedData); err != nil {
		return err
	}
	return nil
}
