package tests

import (
	"crypto/ed25519"
)

func SignGeneric(keyPair GenericKeyPair, data []byte) []byte {
	switch kp := keyPair.(type) {
	case ed25519KeyPair:
		signature := ed25519.Sign(kp.privateKey, data)
		return signature
	default:
		panic("Unsupported Key Pair Type")
	}
}

func GetPublicKeyGeneric(keyPair GenericKeyPair) string {
	switch kp := keyPair.(type) {
	case ed25519KeyPair:
		return kp.publicKey
	default:
		panic("Unsupported Key Pair Type")
	}
}
