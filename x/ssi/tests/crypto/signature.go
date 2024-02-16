package crypto

import (
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/client/cli"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func SignGeneric(keyPair *KeyPair, doc types.SsiMsg, docProof *types.DocumentProof) string {
	switch keyPair.Type {
	case Ed25519KeyPair:
		docProof.Type = types.Ed25519Signature2020

		signature, err := cli.GetDocumentSignature(doc, docProof, keyPair.PrivateKey)
		if err != nil {
			panic(err)
		}
		return signature
	default:
		panic(fmt.Sprintf("Unsupported KeyPair Type : %v", keyPair.Type))
	}
}

func GetPublicKeyAndOptionalID(keyPair *KeyPair) (string, string) {
	switch keyPair.Type {
	case Ed25519KeyPair:
		return keyPair.PublicKey, keyPair.OptionalID
	case Secp256k1Pair:
		return keyPair.PublicKey, keyPair.OptionalID
	default:
		panic(fmt.Sprintf("Unsupported KeyPair Type : %v", keyPair.Type))
	}
}
