package verification

import (
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
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

// Checks whether the ID in the DidDoc is a valid string
func IsValidID(Id string, namespace string, docType string) error {
	var docIdentifier string = documentIdentifier(docType)

	docElements := strings.Split(Id, ":")

	docIdentifierIndex := 0
	docMethodIndex := 1
	docNamespaceIndex := 2
	docMethodSpecificId := 3

	// Document Identifier check
	if docElements[docIdentifierIndex] != docIdentifier {
		return sdkerrors.Wrap(types.ErrInvalidDidDoc, fmt.Sprintf("expected document identifier to be %s, got %s", docIdentifier, docElements[docIdentifierIndex]))
	}

	// did method check
	inputDidMethod := docElements[docMethodIndex]
	if inputDidMethod != DidMethod {
		return sdkerrors.Wrap(types.ErrInvalidDidMethod, fmt.Sprintf("expected did method %s, got %s", DidMethod, inputDidMethod))
	}

	// Mainnet namespace check
	if namespace == "mainnet" {
		if len(docElements) != 3 {
			return sdkerrors.Wrap(types.ErrInvalidDidNamespace, fmt.Sprintf("expected number of did id elements for mainnet to be 3, got %s", fmt.Sprint(len(docElements))))
		}
		docMethodSpecificId = 2
	} else {
		docNamespace := docElements[docNamespaceIndex]
		if namespace != docNamespace {
			return sdkerrors.Wrap(types.ErrInvalidDidNamespace, fmt.Sprintf("expected did namespace %s, got %s", namespace, docNamespace))
		}
	}

	// Check if method-specific-id follows multibase format
	_, _, err := multibase.Decode(docElements[docMethodSpecificId])
	if err != nil || len(docElements[docMethodSpecificId]) != 45 {
		return types.ErrInvalidMethodSpecificId
	}
	return nil
}
