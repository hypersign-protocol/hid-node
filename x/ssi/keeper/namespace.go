package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// namespaceValidation validates the namespace in document id
func namespaceValidation(k msgServer, ctx sdk.Context, docId string) error {
	genesisNamespace := k.GetChainNamespace(&ctx)

	docIdElements := strings.Split(docId, ":")
	docNamespaceIndex := 2

	if genesisNamespace == "" {
		if len(docIdElements) != 3 {
			return fmt.Errorf(
				"expected number of did id elements for mainnet to be 3, got %s for id: %s",
				fmt.Sprint(len(docIdElements)),
				docId,
			)
		}
	} else {
		docNamespace := docIdElements[docNamespaceIndex]
		if genesisNamespace != docNamespace {
			return fmt.Errorf(
				"expected namespace for id %s to be %s, got %s",
				docId, genesisNamespace, docNamespace)
		}
	}

	return nil
}

// didNamespaceValidation validates the namespace in Did Document Id
func didNamespaceValidation(k msgServer, ctx sdk.Context, didDoc *types.Did) error {
	// Subject ID check
	if err := namespaceValidation(k, ctx, didDoc.Id); err != nil {
		return err
	}

	// Controllers check
	for _, controller := range didDoc.Controller {
		if err := namespaceValidation(k, ctx, controller); err != nil {
			return err
		}
	}

	// Verification Method ID checks
	for _, vm := range didDoc.VerificationMethod {
		didId, _ := types.GetElementsFromDidUrl(vm.Id)
		if err := namespaceValidation(k, ctx, didId); err != nil {
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
			if err := namespaceValidation(k, ctx, id); err != nil {
				return err
			}
		}
	}

	return nil
}
