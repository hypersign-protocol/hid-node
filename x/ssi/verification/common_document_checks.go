package verification

import (
	"fmt"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func documentIdentifier(docType string) string {
	if docType == "didDocument" {
		return "did"
	}
	if docType == "schemaDocument" {
		return "sch"
	}
	if docType == "credDocument" {
		return "vc"
	}
	return ""
}

func returnVersionNumIdx(namespace string) int {
	if namespace == "mainnet" {
		return 3
	} else {
		return 4
	}
}

func schemaVersionNumberFormatCheck(docElementsList []string, namespace string) error {
	var verNumIdx int = returnVersionNumIdx(namespace)

	if (len(docElementsList) - 1) != verNumIdx {
		return fmt.Errorf("schema version number is not present in schema id")
	}

	versionNum := docElementsList[verNumIdx]
	versionNumPattern := regexp.MustCompile(`^(\d+\.)?(\d+)$`)
	if !versionNumPattern.MatchString(versionNum) {
		return fmt.Errorf("input version id: %s is invalid", versionNum)
	}

	return nil
}

// Checks whether the ID in the DidDoc is a valid string
func IsValidID(Id string, namespace string, docType string) error {
	var docIdentifier string = documentIdentifier(docType)

	docElements := strings.Split(Id, ":")

	if len(docElements) == 1 {
		return fmt.Errorf("%s id cannot be blank", docType)
	}

	docIdentifierIndex := 0
	docMethodIndex := 1
	docNamespaceIndex := 2
	docMethodSpecificId := 3

	// Document Identifier check
	if docElements[docIdentifierIndex] != docIdentifier {
		return fmt.Errorf("expected document identifier to be %s, got %s", docIdentifier, docElements[docIdentifierIndex])
	}

	// Document method check
	inputDidMethod := docElements[docMethodIndex]
	if inputDidMethod != DidMethod {
		return fmt.Errorf("expected did method %s, got %s", DidMethod, inputDidMethod)
	}

	// Mainnet Chain namespace check. If the document is registered on the mainnet chain,
	// the namespace should be empty
	if namespace == "" {
		if len(docElements) != 3 {
			return fmt.Errorf("expected number of did id elements for mainnet to be 3, got %s", fmt.Sprint(len(docElements)))
		}
		docMethodSpecificId = 2
	} else {
		docNamespace := docElements[docNamespaceIndex]
		if namespace != docNamespace {
			return fmt.Errorf("expected did namespace %s, got %s", namespace, docNamespace)
		}
	}

	// Check if method-specific-id string is alphanumeric and
	// has the minimum required character length of 32
	isProperMethodSpecificId, err := regexp.MatchString(
		"^[a-zA-Z0-9]{32,}$",
		docElements[docMethodSpecificId],
	)
	if err != nil {
		return fmt.Errorf("error in parsing regular expression for method-specific-id: %s", err.Error())
	}
	if !isProperMethodSpecificId {
		return sdkerrors.Wrap(
			types.ErrInvalidMethodSpecificId,
			fmt.Sprintf(
				"method-specific-id should be an alphanumeric string with minimum 32 characters, recieved: %s",
				docElements[docMethodSpecificId],
			),
		)
	}

	// Check for Schema Version Number
	if docType == "schemaDocument" {
		err := schemaVersionNumberFormatCheck(docElements, namespace)
		if err != nil {
			return err
		}
	}

	return nil
}
