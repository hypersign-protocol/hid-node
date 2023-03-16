package verification

import (
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// VerifySignatureOfEveryController verifies every required verification method of every controller
func VerifySignatureOfEveryController(
	didDocMsg types.SsiMsg, VmMap map[string][]*types.ExtendedVerificationMethod,
) error {
	for controller, vmList := range VmMap {
		if len(vmList) == 0 {
			return fmt.Errorf("require atleast one valid signature for controller %s", controller)
		}
		err := verifyAll(vmList, didDocMsg)
		if err != nil {
			return fmt.Errorf("%s: need every signature for controller %s to be valid", err.Error(), controller)
		}
	}
	return nil
}

// VerifySignatureOfEveryController verifies any required0 verification8
func VerifySignatureOfAnyController(
	didDocMsg types.SsiMsg, VmMap map[string][]*types.ExtendedVerificationMethod,
) error {
	found := false
	for _, vmList := range VmMap {
		found = verifyAny(vmList, didDocMsg)
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
func VerifyDocumentProofSignature(ssiMsg types.SsiMsg, vm *types.VerificationMethod, signInfo *types.SignInfo) error {
	vmExtended := types.CreateExtendedVerificationMethod(vm, signInfo)
	if err := verify(vmExtended, ssiMsg); err != nil {
		return err
	}
	return nil
}
