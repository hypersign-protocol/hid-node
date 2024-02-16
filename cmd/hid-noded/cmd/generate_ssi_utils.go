package cmd

import (
	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
)

func formDidId(didNamespace string, publicKeyMultibase string) string {
	if didNamespace != "" {
		return types.DocumentIdentifierDid + ":" + types.DidMethod + ":" + didNamespace + ":" + publicKeyMultibase
	} else {
		return types.DocumentIdentifierDid + ":" + types.DidMethod + ":" + publicKeyMultibase
	}
}

func generateDidDoc(didNamespace string, publicKeyMultibase string, userAddress string) *types.DidDocument {
	didId := formDidId(didNamespace, publicKeyMultibase)

	_, pubKeyBytes, _ := multibase.Decode(publicKeyMultibase)

	return &types.DidDocument{
		Context: []string{
			ldcontext.DidContext,
			ldcontext.Secp256k12019Context,
		},
		Id:         didId,
		Controller: []string{didId},
		VerificationMethod: []*types.VerificationMethod{
			{
				Id:                  didId + "#k1",
				Type:                types.EcdsaSecp256k1VerificationKey2019,
				Controller:          didId,
				PublicKeyMultibase:  publicKeyMultibase,
				// REVERT
				BlockchainAccountId: types.CosmosCAIP10Prefix + ":osmosis-1:" + publicKeyToBech32Address("osmo", pubKeyBytes),
			},
			{
				Id:                  didId + "#k2",
				Type:                types.EcdsaSecp256k1VerificationKey2019,
				Controller:          didId,
				PublicKeyMultibase:  publicKeyMultibase,
				// REVERT
				BlockchainAccountId: types.CosmosCAIP10Prefix + ":prajna:" + publicKeyToBech32Address("hid", pubKeyBytes),
			},
		},
	}
}
