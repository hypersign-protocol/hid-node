package verification

import (
	"fmt"
	"strings"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	utils "github.com/hypersign-protocol/hid-node/x/ssi/utils"
	"github.com/multiformats/go-multibase"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


// Checks whether the ID in the DidDoc is a valid string
func IsValidDidDocID(Id string, namespace string) error {
	didElements := strings.Split(Id, ":")

	didStringIndex := 0
	didMethodIndex := 1
	didNamespaceIndex := 2
	didMethodSpecificId := 3 

	// `did` string check
	if didElements[didStringIndex] != "did" {
		return sdkerrors.Wrap(types.ErrInvalidDidDoc, "first element of Id should be `did`")
	}

	// did method check
	inputDidMethod := didElements[didMethodIndex]
	if inputDidMethod != DidMethod {
		return sdkerrors.Wrap(types.ErrInvalidDidMethod, fmt.Sprintf("expected did method %s, got %s", DidMethod, inputDidMethod))
	}

	// Mainnet namespace check
	if namespace == "mainnet" {
		if len(didElements) != 3 {
			return sdkerrors.Wrap(types.ErrInvalidDidNamespace, fmt.Sprintf("expected number of did id elements for mainnet to be 3"))
		}
		didMethodSpecificId = 2
	} else {
		didNamespace := didElements[didNamespaceIndex]
		if namespace != didNamespace {
			return sdkerrors.Wrap(types.ErrInvalidDidNamespace, fmt.Sprintf("expected did namespace %s, got %s", namespace, didNamespace))
		}
	}

	// Check if method-specific-id follows multibase format
	_, _, err := multibase.Decode(didElements[didMethodSpecificId])
	if err != nil || len(didElements[didMethodSpecificId]) != 45 {
		return types.ErrInvalidMethodSpecificId
	}
	return nil
}

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
	err := IsValidDidDocID(did, namespace)
	if err != nil {
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


// Checks whether the DidDoc string is valid
func IsValidDidDoc(didDoc *types.Did, genesisNamespace string) error {
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

	// Did Id Format Check
	err := IsValidDidDocID(didDoc.GetId(), genesisNamespace)
	if err != nil {
		return err
	}

	// Did Array Check
	for field, didArray := range didArrayMap {
		for _, elem := range didArray{
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
