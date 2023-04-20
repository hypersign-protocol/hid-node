package types

import (
	"fmt"
	"strings"
)

// GetDidFromDidUrl returns didId from didURL.
// TODO: need to handle query and path
func GetElementsFromDidUrl(didUrl string) (string, string) {
	didId, fragment, _ := strings.Cut(didUrl, "#")
	return didId, fragment
}

// splitDidUrl returns the elements of a DID URL. It returns the element in order:
// didId, fragment
func SplitDidUrl(didUrl string) (string, string) {
	didUrlElements := strings.Split(didUrl, "#")

	didId := didUrlElements[0]
	fragment := didUrlElements[1]

	return didId, fragment
}

// GetUniqueElements returns a list with unique elements
func GetUniqueElements(list []string) []string {
	var uniqueList []string
	var elementMap map[string]bool = make(map[string]bool)

	for _, element := range list {
		if !elementMap[element] {
			elementMap[element] = true
			uniqueList = append(uniqueList, element)
		}
	}

	return uniqueList
}

// FindInSlice checks if an element is present in the list
func FindInSlice(list []string, element string) bool {
	for _, x := range list {
		if x == element {
			return true
		}
	}
	return false
}

// checkDuplicateItems return a duplicate Id from the list, if found
func checkDuplicateItems(list []string) string {
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

func getDocumentChainNamespace(docId string) string {
	docIdElements := strings.Split(docId, ":")

	// Non Blockchain Account ID MSI
	if len(docIdElements) == 4 || len(docIdElements) == 6 {
		return docIdElements[2]
	} else {
		return ""
	}
}

func getDocumentIdentifier(docId string) (string, error) {
	docIdElements := strings.Split(docId, ":")

	if len(docIdElements) < 1 {
		return "", fmt.Errorf("invalid document Id: %v", docId)
	}

	return docIdElements[0], nil
}

func getDocumentMethod(docId string) (string, error) {
	docIdElements := strings.Split(docId, ":")

	if len(docIdElements) < 2 {
		return "", fmt.Errorf("invalid document Id: %v", docId)
	}

	return docIdElements[1], nil
}

// GetMethodSpecificIdAndType extracts the method-specific-id from Document Id and which of following
// types it belongs to:
//
// 1. CAIP-10 Blockchain Account ID
//
// 2. String consisting of Alphanumeric, dot (.) and hyphen (-) characters only
func GetMethodSpecificIdAndType(didId string) (string, string, error) {
	docIdElements := strings.Split(didId, ":")
	var methodSpecificId string
	var methodSpecificIdCondition string

	if isMSIBlockchainAccountId(didId) {
		methodSpecificId = strings.Join(docIdElements[(len(docIdElements)-3):], ":")
		methodSpecificIdCondition = MSIBlockchainAccountId
	} else if isMSINonBlockchainAccountId(didId) {
		methodSpecificId = docIdElements[len(docIdElements)-1]
		methodSpecificIdCondition = MSINonBlockchainAccountId
	} else {
		return "", "", fmt.Errorf(
			"unable to retrieve method-specific-id from DID Id: %v. It should either be a CAIP-10 Blockchain Account ID or a string consisting of alphanumeric, dot(.) and hyphen(-) characters only", didId)
	}

	return methodSpecificId, methodSpecificIdCondition, nil
}

// isMSINonBlockchainAccountId asserts if the Method Specific Id is a CAIP-10 Blockchain Account Id
func isMSINonBlockchainAccountId(didId string) bool {
	didIdElements := strings.Split(didId, ":")
	return (len(didIdElements) == 3 || len(didIdElements) == 4)
}

// isMSIBlockchainAccountId asserts if the Method Specific Id is a string containing alphanumeric, dot (.) and hyphen (-) characters
func isMSIBlockchainAccountId(didId string) bool {
	didIdElements := strings.Split(didId, ":")
	return (len(didIdElements) == 5 || len(didIdElements) == 6)
}
