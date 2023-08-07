package types

import (
	"fmt"
)

func chainNamespaceValidation(docId string, validChainNamespace string) error {
	documentChainNamespace := getDocumentChainNamespace(docId)
	if documentChainNamespace != validChainNamespace {
		return fmt.Errorf(
			"expected namespace for id %v to be %v, got %v",
			docId,
			validChainNamespace,
			documentChainNamespace,
		)
	} else {
		return nil
	}
}

// DidDocNamespaceValidation validates the namespace in Did Document
func DidChainNamespaceValidation(didDoc *Did, validChainNamespace string) error {
	// Subject ID check
	if err := chainNamespaceValidation(didDoc.Id, validChainNamespace); err != nil {
		return err
	}

	// Controllers check
	for _, controller := range didDoc.Controller {
		if err := chainNamespaceValidation(controller, validChainNamespace); err != nil {
			return err
		}
	}

	// Verification Method ID checks
	for _, vm := range didDoc.VerificationMethod {
		didId, _ := GetElementsFromDidUrl(vm.Id)
		if err := chainNamespaceValidation(didId, validChainNamespace); err != nil {
			return err
		}
	}

	// Verification Relationships check
	vmRelationshipList := [][]string{
		didDoc.Authentication,
		didDoc.AssertionMethod,
		didDoc.KeyAgreement,
		didDoc.CapabilityDelegation,
		didDoc.CapabilityInvocation,
	}

	for _, vmRelationship := range vmRelationshipList {
		// didUrl check and presence in verification methods
		for _, id := range vmRelationship {
			if err := chainNamespaceValidation(id, validChainNamespace); err != nil {
				return err
			}
		}
	}

	return nil
}
