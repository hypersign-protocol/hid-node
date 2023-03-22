package types

import (
	"fmt"
	"regexp"
	"strings"
)

// isValidDidDoc checks if the DID Id is valid
func isValidDidDocId(id string) error {
	// check the number of elements in DID Document
	idElements := strings.Split(id, ":")
	if !(len(idElements) == 3 || len(idElements) == 4) {
		return fmt.Errorf(
			"number of elements in DID Id %s should be either 3 or 4",
			id,
		)
	}

	// check if the first element is valid document identifier
	if idElements[0] != DocumentIdentifierDid {
		return fmt.Errorf(
			"document identifier should be %s",
			DocumentIdentifierDid,
		)
	}

	// check if the second element is the correct DID method
	if idElements[1] != DidMethod {
		return fmt.Errorf(
			"DID method should be %s",
			DidMethod,
		)
	}

	// check proper method specific id
	// TODO: need to define a specification for method-specific-id
	methodSpecificId := idElements[len(idElements)-1]
	isProperMethodSpecificId, err := regexp.MatchString(
		"^[a-zA-Z0-9]{32,}$",
		methodSpecificId,
	)
	if err != nil {
		return fmt.Errorf("error in parsing regular expression for method-specific-id: %s", err.Error())
	}
	if !isProperMethodSpecificId {
		return fmt.Errorf(
			"method-specific-id should be an alphanumeric string with minimum of 32 characters, received: %s",
			methodSpecificId,
		)
	}

	return nil
}

// isDidUrl checks if the input is a DID Url or not.
// More on DID Url syntax: https://www.w3.org/TR/did-core/#did-url-syntax
// TODO: add check for path and query
func isDidUrl(id string) error {
	didId, fragment := GetElementsFromDidUrl(id)

	// check for fragment
	if didId == "" || fragment == "" {
		return fmt.Errorf(
			"invalid didUrl %s",
			id,
		)
	}

	// check for DID Id
	err := isValidDidDocId(didId)
	if err != nil {
		return err
	}
	return nil
}

func isSupportedVmType(typ string) error {
	if _, supported := VerificationKeySignatureMap[typ]; !supported {
		return fmt.Errorf(
			"%v verification method type not supported, supported verification method types: %v ",
			typ,
			supportedVerificationMethodTypes,
		)
	}
	return nil
}

// verificationKeyCheck validates of publicKeyMultibase and blockchainAccountId.
func verificationKeyCheck(vm *VerificationMethod) error {
	switch vm.Type {
	case EcdsaSecp256k1VerificationKey2019:
		if vm.GetPublicKeyMultibase() == "" {
			return fmt.Errorf(
				"publicKeyMultibase cannot be empty for verification method %s as it is type %s",
				vm.Id,
				vm.Type,
			)
		}
	case EcdsaSecp256k1RecoveryMethod2020:
		if vm.GetBlockchainAccountId() == "" {
			return fmt.Errorf(
				"blockchainAccountId cannot be empty for verification method %s as it is of type %s",
				vm.Id,
				vm.Type,
			)
		}

		if vm.GetPublicKeyMultibase() != "" {
			return fmt.Errorf(
				"publicKeyMultibase should not be provided for verification method %s as it is of type %s",
				vm.Id,
				vm.Type,
			)
		}
	case Ed25519VerificationKey2020:
		if vm.GetBlockchainAccountId() != "" {
			return fmt.Errorf(
				"blockchainAccountId is currently not supported for verification method %s as it is of type %s",
				vm.Id,
				vm.Type,
			)
		}
		if vm.GetPublicKeyMultibase() == "" {
			return fmt.Errorf(
				"publicKeyMultibase cannot be empty for verification method %s as it is of type %s",
				vm.Id,
				vm.Type,
			)
		}
	default:
		return fmt.Errorf("unsupported verification method type: %v", supportedVerificationMethodTypes)
	}

	// validate blockchainAccountId
	if vm.BlockchainAccountId != "" {
		err := validateBlockchainAccountId(vm.BlockchainAccountId)
		if err != nil {
			return fmt.Errorf("invalid blockchainAccount Id %v: "+err.Error(), vm.BlockchainAccountId)
		}
	}

	return nil
}

// checkDuplicateId return a duplicate Id from the list, if found
func checkDuplicateId(list []string) string {
	presentMap := map[string]bool{}
	for idx := range list {
		if _, present := presentMap[list[idx]]; !present {
			presentMap[list[idx]] = true
		} else {
			return list[idx]
		}
	}
	return ""
}

// validateServices validates the Service attribute of DID Document
func validateServices(services []*Service) error {
	for _, service := range services {
		var err error

		// validate service Id
		if err = isDidUrl(service.Id); err != nil {
			return fmt.Errorf("service ID %s is Invalid", service.Id)
		}

		// validate service Type
		foundType := false
		for _, sType := range SupportedServiceTypes {
			if service.Type == sType {
				foundType = true
			}
		}
		if !foundType {
			return fmt.Errorf("service Type %s is Invalid", service.Type)
		}
	}

	// check if any duplicate service id exists
	serviceIdList := []string{}
	for _, service := range services {
		serviceIdList = append(serviceIdList, service.Id)
	}
	if duplicateId := checkDuplicateId(serviceIdList); duplicateId != "" {
		return fmt.Errorf("duplicate service found with Id: %s ", duplicateId)
	}

	return nil
}

// validateVerificationMethods validates all the verification methods present in DID Document
func validateVerificationMethods(vms []*VerificationMethod) error {
	for _, vm := range vms {
		var err error

		// Vm Id check
		err = isDidUrl(vm.Id)
		if err != nil {
			return err
		}

		// Vm Type check
		err = isSupportedVmType(vm.Type)
		if err != nil {
			return err
		}

		// Controller check
		err = isValidDidDocId(vm.Controller)
		if err != nil {
			return err
		}

		// blockchainAccountId and publicKeyMultibase check
		err = verificationKeyCheck(vm)
		if err != nil {
			return err
		}
	}

	// check duplicate Vm Ids, publicKeyMultibase and blockchainAccountId
	vmIdList := []string{}
	publicKeyMultibaseList := []string{}
	blockchainAccountIdList := []string{}

	for _, vm := range vms {
		vmIdList = append(vmIdList, vm.Id)
		publicKeyMultibaseList = append(publicKeyMultibaseList, vm.PublicKeyMultibase)
		blockchainAccountIdList = append(blockchainAccountIdList, vm.BlockchainAccountId)
	}
	if duplicateId := checkDuplicateId(vmIdList); duplicateId != "" {
		return fmt.Errorf("duplicate verification method Id found: %s ", duplicateId)
	}
	if duplicateKey := checkDuplicateId(publicKeyMultibaseList); duplicateKey != "" {
		return fmt.Errorf("duplicate publicKeyMultibase found: %s ", duplicateKey)
	}
	if duplicateKey := checkDuplicateId(blockchainAccountIdList); duplicateKey != "" {
		return fmt.Errorf("duplicate blockchainAccountId found: %s ", duplicateKey)
	}

	return nil
}

func validateVmRelationships(didDoc *Did) error {
	// make verificationMethods map
	vmMap := map[string]bool{}
	for _, vm := range didDoc.VerificationMethod {
		vmMap[vm.Id] = true
	}

	vmRelationshipList := map[string][]string{
		"authentication":       didDoc.Authentication,
		"assertionMethod":      didDoc.AssertionMethod,
		"keyAgreement":         didDoc.KeyAgreement,
		"capabilityDelegation": didDoc.CapabilityDelegation,
		"capabilityInvocation": didDoc.CapabilityInvocation,
	}

	for field, vmRelationship := range vmRelationshipList {
		// didUrl check and presence in verification methods
		for _, element := range vmRelationship {
			err := isDidUrl(element)
			if err != nil {
				return fmt.Errorf("%s: %s", field, err)
			}
			if _, found := vmMap[element]; !found {
				return fmt.Errorf(
					"%s: verification method id %s not found in verificationMethod list",
					field,
					element,
				)
			}
		}
	}

	return nil
}

func validateBlockchainAccountId(blockchainAccountId string) error {
	blockchainId, err := NewBlockchainId(blockchainAccountId)
	if err != nil {
		return err
	}

	var validationErr error

	// Check for supported CAIP-10 prefix
	validationErr = blockchainId.ValidateSupportedCAIP10Prefix()
	if validationErr != nil {
		return validationErr
	}

	// Check for supported CAIP-10 chain-ids
	validationErr = blockchainId.ValidateSupportChainId()
	if validationErr != nil {
		return validationErr
	}

	// Check for supported CAIP-10 bech32 prefix. Perform this validation
	// only when the CAIP-10 prefix is "cosmos"
	if blockchainId.CAIP10Prefix == CosmosCAIP10Prefix {
		validationErr = blockchainId.ValidateSupportedBech32Prefix()
		if validationErr != nil {
			return validationErr
		}
	}

	return nil
}

// ValidateDidDocument validates the DID Document
func (didDoc *Did) ValidateDidDocument() error {
	// Id check
	err := isValidDidDocId(didDoc.Id)
	if err != nil {
		return err
	}

	// Controller check
	for _, controller := range didDoc.Controller {
		err := isValidDidDocId(controller)
		if err != nil {
			return err
		}
	}

	// VerificationMethod check
	err = validateVerificationMethods(didDoc.VerificationMethod)
	if err != nil {
		return err
	}

	// Services check
	err = validateServices(didDoc.Service)
	if err != nil {
		return err
	}

	// Verification Method Relationships check
	err = validateVmRelationships(didDoc)
	if err != nil {
		return err
	}

	return nil
}
