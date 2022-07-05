package verification

import (
	"fmt"
	"strings"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	utils "github.com/hypersign-protocol/hid-node/x/ssi/utils"
	"github.com/multiformats/go-multibase"
)


// Checks whether the ID in the DidDoc is a valid string
func IsValidDidDocID(Id string) string {
	didElements := strings.Split(Id, ":")

	if len(didElements) != didIdElements {
		return types.ErrInvalidDidElements.Error()
	}

	// Check if method-specific-id follows multibase format
	_, _, err := multibase.Decode(didElements[3])
	if err != nil || len(didElements[3]) != 45 {
		return types.ErrInvalidMethodSpecificId.Error()
	}
	return ""
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
	err := IsValidDidDocID(did)
	if err != "" {
		return false
	}
	return true
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

// Check whether the fields whose values are array of DIDs are valid DID
func IsValidDIDArray(didArray []string) bool {
	for _, did := range didArray {
		if err := IsValidDidDocID(did); err != "" {
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
	if IsValidDidDocID(didDoc.GetId()) != "" {
		return fmt.Sprintf("The DidDoc ID %s is invalid", didDoc.GetId())
	}

	// Did Array Check
	for field, didArray := range didArrayMap {
		for _, elem := range didArray{
			if !IsValidDidFragment(elem) {
				return fmt.Sprintf("The field %s is an invalid DID Array", field)
			}
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
