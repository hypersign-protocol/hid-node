package types

import (
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
