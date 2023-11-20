package types

import (
	"fmt"
	"regexp"
)

// isValidDidDoc checks if the DID Id is valid
func isValidDidDocId(id string) error {
	inputDocumentIdentifier, err := getDocumentIdentifier(id)
	if err != nil {
		return err
	}

	inputDidMethod, err := getDocumentMethod(id)
	if err != nil {
		return err
	}

	inputMSI, inputMSIType, err := GetMethodSpecificIdAndType(id)
	if err != nil {
		return err
	}

	// Validate Document Identifier
	if inputDocumentIdentifier != DocumentIdentifierDid {
		return fmt.Errorf(
			"document identifier should be %s",
			DocumentIdentifierDid,
		)
	}

	// Validate DID Method
	if inputDidMethod != DidMethod {
		return fmt.Errorf(
			"DID method should be %s",
			DidMethod,
		)
	}

	// Validate Method Specific ID
	switch inputMSIType {
	case MSIBlockchainAccountId:
		if err := validateBlockchainAccountId(inputMSI); err != nil {
			return err
		}
	case MSINonBlockchainAccountId:
		// Non Blockchain Account ID should be a string that supports alphanumeric characters,
		// and dot (.) and hypen (-). The first character MUST NOT be dot (.) or hyphen (-).
		isValidMSI, err := regexp.MatchString(
			"^[a-zA-Z0-9][a-zA-Z0-9.-]*$",
			inputMSI,
		)
		if err != nil {
			return err
		}
		if !isValidMSI {
			return fmt.Errorf(
				"method-specific-id of non BlockchainAccountId type %v should only contain alphanumeric, dot (.) and hyphen (-)",
				inputMSI,
			)
		}
	default:
		return fmt.Errorf("invalid method specific id type: %v", inputMSIType)
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
	case X25519KeyAgreementKey2020:
		if vm.GetPublicKeyMultibase() == "" {
			return fmt.Errorf(
				"publicKeyMultibase cannot be empty for verification method %s as it is of type %s",
				vm.Id,
				vm.Type,
			)
		}
		if vm.GetBlockchainAccountId() != "" {
			return fmt.Errorf(
				"blockchainAccountId must be empty for verification method %s as it is of type %s",
				vm.Id,
				vm.Type,
			)
		}
	case X25519KeyAgreementKeyEIP5630:
		if vm.GetPublicKeyMultibase() == "" {
			return fmt.Errorf(
				"publicKeyMultibase cannot be empty for verification method %s as it is of type %s",
				vm.Id,
				vm.Type,
			)
		}
		if vm.GetBlockchainAccountId() != "" {
			return fmt.Errorf(
				"blockchainAccountId must be empty for verification method %s as it is of type %s",
				vm.Id,
				vm.Type,
			)
		}
	case Bls12381G2Key2020:
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
	case BabyJubJubVerificationKey2023:
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
		return fmt.Errorf("unsupported verification method type: %v. Supported verification method types are: %v", vm.Type, supportedVerificationMethodTypes)
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
	if duplicateId := checkDuplicateItems(serviceIdList); duplicateId != "" {
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

	var pubKeyMultibaseBlockchainAccIdMap map[string]bool = map[string]bool{}

	for _, vm := range vms {
		vmIdList = append(vmIdList, vm.Id)

		if vm.Type == EcdsaSecp256k1VerificationKey2019 {
			if _, present := pubKeyMultibaseBlockchainAccIdMap[vm.PublicKeyMultibase]; present {
				// TODO: Following is a temporary measure, where we will be allowing duplicate publicKeyMultibase values
				// for type EcdsaSecp256k1VerificationKey2019, provided if blockchainAccountId field is populated. This is done since
				// one secp256k1 key pair from Keplr wallet can have multiple blockchain addresses depending upon the bech32
				// prefix. This will eventually be removed once the successful recovery of public key from the `signEthereum()`
				// generated signature is figured out.
				if vm.BlockchainAccountId == "" {
					return fmt.Errorf(
						"duplicate publicKeyMultibase of type EcdsaSecp256k1VerificationKey2019 without blockchainAccountId is not allowed: %s ",
						vm.PublicKeyMultibase,
					)
				}
			} else {
				pubKeyMultibaseBlockchainAccIdMap[vm.PublicKeyMultibase] = true
			}
		} else {
			publicKeyMultibaseList = append(publicKeyMultibaseList, vm.PublicKeyMultibase)
		}

		blockchainAccountIdList = append(blockchainAccountIdList, vm.BlockchainAccountId)
	}

	if duplicateId := checkDuplicateItems(vmIdList); duplicateId != "" {
		return fmt.Errorf("duplicate verification method Id found: %s ", duplicateId)
	}
	if duplicateKey := checkDuplicateItems(publicKeyMultibaseList); duplicateKey != "" {
		return fmt.Errorf("duplicate publicKeyMultibase found: %s ", duplicateKey)
	}
	if duplicateKey := checkDuplicateItems(blockchainAccountIdList); duplicateKey != "" {
		return fmt.Errorf("duplicate blockchainAccountId found: %s ", duplicateKey)
	}

	return nil
}

func validateVmRelationships(didDoc *DidDocument) error {
	// make verificationMethodType map between VM Id and VM type
	vmTypeMap := map[string]string{}
	for _, vm := range didDoc.VerificationMethod {
		vmTypeMap[vm.Id] = vm.Type
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
		for _, vmId := range vmRelationship {
			err := isDidUrl(vmId)
			if err != nil {
				return fmt.Errorf("%s: %s", field, err)
			}

			if _, found := vmTypeMap[vmId]; !found {
				return fmt.Errorf(
					"%s: verification method id %s not found in verificationMethod list",
					field,
					vmId,
				)
			}

			// keyAgreement field should harbour only those Verification Methods whose type is either X25519KeyAgreementKey2020
			// or X25519KeyAgreementKeyEIP5630
			if (vmTypeMap[vmId] == X25519KeyAgreementKey2020) || (vmTypeMap[vmId] == X25519KeyAgreementKeyEIP5630) {
				if field != "keyAgreement" {
					return fmt.Errorf(
						"verification method id %v is of type %v which is not allowed in '%v' attribute",
						vmId,
						vmTypeMap[vmId],
						field,
					)
				}
			} else {
				if field == "keyAgreement" {
					return fmt.Errorf(
						"verification method id %v provided in '%v' attribute must be of type X25519KeyAgreementKey2020 or X25519KeyAgreementKeyEIP5630",
						vmId,
						field,
					)
				}
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
func (didDoc *DidDocument) ValidateDidDocument() error {
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
