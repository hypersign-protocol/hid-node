package verification

import (
	"fmt"
	"strings"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// Extracts the blockchain address from blockchainAccountId
func getBlockchainAddress(blockchainAccountId string) string {
	blockchainAccountIdElements := strings.Split(blockchainAccountId, ":")
	blockchainAddress := blockchainAccountIdElements[len(blockchainAccountIdElements)-1]
	return blockchainAddress
}

// getChainIdFromBlockchainAccountId extracts chain from blockchainAccountId
func getChainIdFromBlockchainAccountId(blockchainAccountId string) string {
	chainIdIdx := 1

	blockchainAccountIdElements := strings.Split(blockchainAccountId, ":")
	blockchainChainId := blockchainAccountIdElements[chainIdIdx]
	return blockchainChainId
}

// Extracts the CAIP-10 prefix from blockchainAccountId and returns the chain spec
func getCAIP10Prefix(blockchainAccountId string) (string, error) {
	segments := strings.Split(blockchainAccountId, ":")

	// Validate blockchainAccountId
	if len(segments) != 3 {
		return "", fmt.Errorf(
			"invalid CAIP-10 format for blockchainAccountId '%v'. Please refer: https://github.com/ChainAgnostic/CAIPs/blob/master/CAIPs/caip-10.md",
			blockchainAccountId,
		)
	}

	// Prefix check
	if segments[0] == types.EthereumCAIP10Prefix {
		return types.EthereumCAIP10Prefix, nil
	} else if segments[0] == types.CosmosCAIP10Prefix {
		return types.CosmosCAIP10Prefix, nil
	} else {
		return "", fmt.Errorf(
			"unsupported CAIP-10 prefix in blockchainAccountId '%v'. Supported CAIP-10 prefixes: %v",
			blockchainAccountId,
			types.CAIP10PrefixForEcdsaSecp256k1RecoveryMethod2020,
		)
	}
}
