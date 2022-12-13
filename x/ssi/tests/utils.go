package tests

import (
	"crypto/ed25519"
	"strings"
)

func SignGeneric(keyPair GenericKeyPair, data []byte) []byte {
	switch kp := keyPair.(type) {
	case ed25519KeyPair:
		signature := ed25519.Sign(kp.privateKey, data)
		return signature
	case secp256k1KeyPair:
		signature, err := kp.privateKey.Sign(data)
		if err != nil {
			panic(err)
		}
		return signature
	default:
		panic("Unsupported Key Pair Type")
	}
}

func GetPublicKeyAndOptionalID(keyPair GenericKeyPair) (string, string) {
	switch kp := keyPair.(type) {
	case ed25519KeyPair:
		return kp.publicKey, kp.optionalID
	case secp256k1KeyPair:
		return kp.publicKey, kp.optionalID
	default:
		panic("Unsupported Key Pair Type")
	}
}

func stripDidFromVerificationMethod(vmId string) string {
	segments := strings.Split(vmId, "#")
	return segments[0]
}
