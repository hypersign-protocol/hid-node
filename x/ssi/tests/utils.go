package tests

import (
	"crypto/ed25519"
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

func GetPublicKeyGeneric(keyPair GenericKeyPair) string {
	switch kp := keyPair.(type) {
	case ed25519KeyPair:
		return kp.publicKey
	case secp256k1KeyPair:
		return kp.publicKey
	default:
		panic("Unsupported Key Pair Type")
	}
}
