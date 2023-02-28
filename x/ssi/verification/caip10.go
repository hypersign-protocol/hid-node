package verification

import (
	"strings"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// Extracts the blockchain address from blockchainAccountId
func getBlockchainAddress(blockchainAccountId string) string {
	blockchainAccountIdElements := strings.Split(blockchainAccountId, ":")
	blockchainAddress := blockchainAccountIdElements[len(blockchainAccountIdElements)-1]
	return blockchainAddress
}

// Extracts the CAIP-10 prefix from blockchainAccountId and returns the chain spec
func getCAIP10Chain(blockchainAccountId string) string {
	segments := strings.Split(blockchainAccountId, ":")
	userPrefix := strings.Join(segments[0:len(segments)-1], ":")

	// Ethereum based chain (EIP-155) check
	if strings.HasPrefix(userPrefix, types.EIP155) {
		return types.EIP155
	} else {
		return ""
	}
}
