package types

import (
	fmt "fmt"
	"regexp"
)

// DefaultGenesis returns the default ssi genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ChainNamespace: "",
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	namespace := gs.ChainNamespace

	regexPattern, _ := regexp.Compile("^[a-zA-Z0-9-]*$") // Matches string containing whitespaces and tabs
	maxChainNamespaceLength := 10

	if len(namespace) > maxChainNamespaceLength {
		return fmt.Errorf("chain namespace shouldn't shouldn't be more than 10, namespace recieved %s", namespace)
	}

	if !regexPattern.MatchString(namespace) && len(namespace) != 0 {
		return fmt.Errorf("chain namespace should be in alphanumeric format, namespace recieved %s", namespace)
	}

	return nil
}
