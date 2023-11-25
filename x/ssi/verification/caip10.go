package verification

import (
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// Extracts the blockchain address from blockchainAccountId
func getBlockchainAddress(blockchainAccountId string) (string, error) {
	bid, err := types.NewBlockchainId(blockchainAccountId)
	if err != nil {
		return "", err
	}

	return bid.BlockchainAddress, nil
}

// getChainIdFromBlockchainAccountId extracts chain-id from blockchainAccountId
func getChainIdFromBlockchainAccountId(blockchainAccountId string) (string, error) {
	bid, err := types.NewBlockchainId(blockchainAccountId)
	if err != nil {
		return "", err
	}

	return bid.ChainId, nil
}

// Extracts the CAIP-10 prefix from blockchainAccountId and returns the chain spec
func getCAIP10Prefix(blockchainAccountId string) (string, error) {
	bid, err := types.NewBlockchainId(blockchainAccountId)
	if err != nil {
		return "", err
	}

	return bid.CAIP10Prefix, nil
}
