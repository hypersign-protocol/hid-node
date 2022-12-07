package verification

import (
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/vid-node/x/ssi/types"
	utils "github.com/hypersign-protocol/vid-node/x/ssi/utils"
)

// Cheks whether the Service is valid
func ValidateServices(services []*types.Service, method string, namespace string) error {
	for idx, service := range services {
		if !IsValidDidFragment(service.Id, method, namespace) {
			return types.ErrInvalidService.Wrapf("Service ID %s is Invalid", service.Id)
		}

		if !IsValidDidServiceType(service.Type) {
			return types.ErrInvalidService.Wrapf("Service Type %s is Invalid", service.Type)
		}

		if DuplicateServiceExists(service.Id, services[idx+1:]) {
			return types.ErrInvalidService.Wrapf("Service with Id: %s is duplicate", service.Id)
		}
	}
	return nil
}

// Check for valid DID fragment
func IsValidDidFragment(didUrl string, method string, namespace string) bool {
	if !strings.Contains(didUrl, "#") {
		return false
	}

	did, _ := utils.SplitDidUrlIntoDid(didUrl)
	err := IsValidID(did, namespace, "didDocument")
	return err == nil
}

// Check Valid DID service type
func IsValidDidServiceType(sType string) bool {
	for _, val := range ServiceTypes {
		if val == sType {
			return true
		}
	}
	return false
}

func DuplicateServiceExists(serviceId string, services []*types.Service) bool {
	_, fragment := utils.SplitDidUrlIntoDid(serviceId)
	for _, s := range services {
		_, sFragment := utils.SplitDidUrlIntoDid(s.Id)
		if fragment == sFragment {
			return true
		}
	}
	return false
}

// Checks whether the DidDoc string is valid
func ValidateDidDocument(didDoc *types.Did, genesisNamespace string) error {
	didArrayMap := map[string][]string{
		"authentication":       didDoc.GetAuthentication(),
		"assertionMethod":      didDoc.GetAssertionMethod(),
		"keyAgreement":         didDoc.GetKeyAgreement(),
		"capabilityInvocation": didDoc.GetCapabilityInvocation(),
		"capabilityDelegation": didDoc.GetCapabilityDelegation(),
	}

	nonEmptyFields := map[string]string{
		"id": didDoc.GetId(),
	}

	// Format Check for Did Id
	err := IsValidID(didDoc.GetId(), genesisNamespace, "didDocument")
	if err != nil {
		return err
	}

	// Verification Method Relationships Check
	for field, didArray := range didArrayMap {
		for _, elem := range didArray {
			if !IsValidDidFragment(elem, DidMethod, genesisNamespace) {
				return sdkerrors.Wrap(types.ErrInvalidDidDoc, fmt.Sprintf("The field %s is an invalid DID Array", field))
			}
		}
	}

	// Empty Field check
	for field, value := range nonEmptyFields {
		if value == "" {
			return sdkerrors.Wrap(types.ErrInvalidDidDoc, fmt.Sprintf("The field %s must have a value", field))
		}
	}

	// Valid Services Check
	err = ValidateServices(didDoc.GetService(), DidMethod, genesisNamespace)
	if err != nil {
		return err
	}

	return nil
}

// Check the Deactivate status of DID
func VerifyDidDeactivate(metadata *types.Metadata, id string) error {
	if metadata.GetDeactivated() {
		return sdkerrors.Wrap(types.ErrDidDocDeactivated, fmt.Sprintf("DidDoc ID: %s", id))
	}
	return nil
}
