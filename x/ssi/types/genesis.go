package types

import (
	fmt "fmt"
	"regexp"
)

// DefaultGenesis returns the default ssi genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// TODO: Once did method spec has been confirmed, did_method should be removed
		DidMethod: "hs",
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	namespace := gs.DidNamespace

	regexPattern, _ := regexp.Compile("^[a-zA-Z0-9-]*$") // Matches string containing whitespaces and tabs
	maxDidNamespaceLength := 10

	if len(namespace) > maxDidNamespaceLength  {
		return fmt.Errorf("Did Namespace shouldn't shouldn't be more than 10, namespace recieved %s", namespace)
	}

	if !regexPattern.MatchString(namespace) && len(namespace) != 0 {
		return fmt.Errorf("Did Namespace should be in alphanumeric format, namespace recieved %s", namespace)
	}

	return nil
}
