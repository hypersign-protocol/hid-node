package verification

import (
	"fmt"
	"strings"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	utils "github.com/hypersign-protocol/hid-node/x/ssi/utils"
)

// Check if the DID Method is valid
func IsValidDidMethod(method string) bool {
	return method == didMethod
}

// Checks whether the given string is a valid DID
func IsValidDid(did string) error {
	didElements := strings.Split(did, ":")

	if (didElements[0] != "did") || len(didElements) != didIdElements {
		return types.ErrInvalidDidElements
	}
	if !IsValidDidMethod(didElements[1]) {
		return types.ErrInvalidDidMethod
	}

	return nil
}

// Checks whether the ID in the DidDoc is a valid string
func IsValidDidDocID(didDoc *types.Did) bool {
	if err := IsValidDid(didDoc.GetId()); err != nil {
		return false
	}
	return true
}

// Cheks whether the Service is valid
func ValidateServices(services []*types.Service) error {
	for idx, service := range services {
		if !IsValidDidFragment(service.Id) {
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
func IsValidDidFragment(didUrl string) bool {
	if !strings.Contains(didUrl, "#") {
		return false
	}

	did, _ := utils.SplitDidUrlIntoDid(didUrl)
	err := IsValidDid(did)
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
	did, _ := utils.SplitDidUrlIntoDid(serviceId)
	for _, s := range services {
		sDid, _ := utils.SplitDidUrlIntoDid(s.Id)
		if did == sDid {
			return true
		}
	}
	return false
}

// Check whether the fields whose values are array of DIDs are valid DID
func IsValidDIDArray(didArray []string) bool {
	for _, did := range didArray {
		if err := IsValidDid(did); err != nil {
			return false
		}
	}
	return true
}

// Checks whether the DidDoc string is valid
func IsValidDidDoc(didDoc *types.Did) string {
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

	// Invalid ID check
	if !IsValidDidDocID(didDoc) {
		return fmt.Sprintf("The DidDoc ID %s is invalid", didDoc.GetId())
	}

	// Did Array Check
	for field, didArray := range didArrayMap {
		if !IsValidDIDArray(didArray) {
			return fmt.Sprintf("The field %s is an invalid DID Array", field)
		}
	}

	// Empty Field check
	for field, value := range nonEmptyFields {
		if value == "" {
			return fmt.Sprintf("The field %s must have a value", field)
		}
	}

	// Valid Services Check
	err := ValidateServices(didDoc.GetService())
	if err != nil {
		return fmt.Sprint(err)
	}

	return ""

}