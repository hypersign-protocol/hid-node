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

func getDidDocumentIdentifier(didId string) (string, error) {
	didIdElements := strings.Split(didId, ":")

	if len(didIdElements) < 1 {
		return "", fmt.Errorf("invalid DID Id: %v", didId)
	}

	return didIdElements[0], nil
}

func getDidDocumentMethod(didId string) (string, error) {
	didIdElements := strings.Split(didId, ":")

	if len(didIdElements) < 2 {
		return "", fmt.Errorf("invalid DID Id: %v", didId)
	}

	return didIdElements[1], nil
}

// GetMethodSpecificIdAndType extracts the method-specific-id from DID Id and which of following
// types it belongs to:
//
// 1. CAIP-10 Blockchain Account ID
//
// 2. String consisting of Alphanumeric, dot (.) and hyphen (-) characters only
func GetMethodSpecificIdAndType(didId string) (string, string, error) {
	didIdElements := strings.Split(didId, ":")
	var methodSpecificId string
	var methodSpecificIdCondition string

	if getMSIBlockchainAccountIdCondition(didId) {
		methodSpecificId = strings.Join(didIdElements[(len(didIdElements)-3):], ":")
		methodSpecificIdCondition = MSIBlockchainAccountId
	} else if getMSINonBlockchainAccountIdCondition(didId) {
		methodSpecificId = didIdElements[len(didIdElements)-1]
		methodSpecificIdCondition = MSINonBlockchainAccountId
	} else {
		return "", "", fmt.Errorf(
			"unable to retrieve method-specific-id from DID Id: %v. It should either be a CAIP-10 Blockchain Account ID or a string consisting of alphanumeric, dot(.) and hyphen(-) characters only", didId)
	}

	return methodSpecificId, methodSpecificIdCondition, nil
}

// getMSINonBlockchainAccountIdCondition asserts if the Method Specific Id is a CAIP-10 Blockchain Account Id
func getMSINonBlockchainAccountIdCondition(didId string) bool {
	didIdElements := strings.Split(didId, ":")
	return (len(didIdElements) == 3 || len(didIdElements) == 4)
}

// getMSIMultibaseCondition asserts if the Method Specific Id is a Multibase encoded string
func getMSIBlockchainAccountIdCondition(didId string) bool {
	didIdElements := strings.Split(didId, ":")
	return (len(didIdElements) == 5 || len(didIdElements) == 6)
}
