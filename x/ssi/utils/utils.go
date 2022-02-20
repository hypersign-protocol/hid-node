package utils

import (
	"crypto/ed25519"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// Checks whether the given string is a valid DID
func IsValidDid(did string) bool {
	return strings.HasPrefix(did, DidPrefix)
}

// Checks whether the ID in the DidDoc is a valid string
func IsValidDidDocID(didDoc *types.Did) bool {
	return strings.HasPrefix(didDoc.GetId(), DidPrefix)
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

	did, _ := SplitDidUrlIntoDid(didUrl)
	hasPrefix := IsValidDid(did)
	return hasPrefix
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
	did, _ := SplitDidUrlIntoDid(serviceId)
	for _, s := range services {
		sDid, _ := SplitDidUrlIntoDid(s.Id)
		if did == sDid {
			return true
		}
	}
	return false
}

// Check whether the fields whose values are array of DIDs are valid DID
func IsValidDIDArray(didArray []string) bool {
	for _, did := range didArray {
		if !IsValidDid(did) {
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

func SplitDidUrlIntoDid(didUrl string) (string, string) {
	segments := strings.Split(didUrl, "#")
	return segments[0], segments[1]
}

func FindPublicKey(signer types.Signer, id string) (ed25519.PublicKey, error) {
	for _, authentication := range signer.Authentication {
		if authentication == id {
			vm := FindVerificationMethod(signer.VerificationMethod, id)
			if vm == nil {
				return nil, types.ErrVerificationMethodNotFound.Wrap(id)
			}
			return vm.GetPublicKey()
		}
	}

	return nil, types.ErrVerificationMethodNotFound.Wrap(id)
}

func FindVerificationMethod(vms []*types.VerificationMethod, id string) *types.VerificationMethod {
	for _, vm := range vms {
		if vm.Id == id {
			return vm
		}
	}

	return nil
}

func IsValidSchemaID(schemaId string) error {
	IdComponents := strings.Split(schemaId, ";")
	if len(IdComponents) != 3 {
		return errors.New("Expected 3 components in schema ID after being seperated by `;`, got " + fmt.Sprint(len(IdComponents)) + " components. The Schema ID is `" + schemaId + "` ")
	}

	//Checking the prefix
	if !strings.HasPrefix(IdComponents[0], "did:hs:") {
		return errors.New("Expected did:hs as prefix in schema ID, The Schema ID is " + schemaId)
	}

	//Checking the type of version
	versionNumber := strings.Split(IdComponents[2], "=")[1]
	// TODO: The regex pattern should be configurable to match the version format.
	// Currently it's set for floating point validation
	regexPatternForVersion := regexp.MustCompile(`^(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)$`)
	if !regexPatternForVersion.MatchString(versionNumber) {
		return fmt.Errorf("input version Id -> `%s`, is an invalid format", versionNumber)
	}
	return nil
}
