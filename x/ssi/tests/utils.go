package tests

import (
	"crypto/ed25519"

	sha256 "crypto/sha256"
	ecdsa "github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

func SignGeneric(keyPair GenericKeyPair, data []byte) []byte {
	switch kp := keyPair.(type) {
	case ed25519KeyPair:
		signature := ed25519.Sign(kp.privateKey, data)
		return signature
	case secp256k1KeyPair:
		dataHash := sha256.Sum256(data)
		secp256kSignature := ecdsa.Sign(kp.privateKey, dataHash[:])
		signature := secp256kSignature.Serialize()
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
