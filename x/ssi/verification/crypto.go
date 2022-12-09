package verification

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
	secp256k1 "github.com/tendermint/tendermint/crypto/secp256k1"
)

func verify(verificationMethodType string, verificationKey string, signature string, data []byte) (bool, error) {
	switch verificationMethodType {
	case Ed25519VerificationKey2020:
		return verifyEd25519(verificationKey, signature, data)
	case EcdsaSecp256k1VerificationKey2019:
		return verifySecp256k1(verificationKey, signature, data)
	default:
		return false, fmt.Errorf("unsupported verification method: %s", verificationMethodType)
	}
}

func verifyEd25519(publicKey string, signature string, documentBytes []byte) (bool, error) {
	// Decode Public Key
	_, publicKeyBytes, err := multibase.Decode(publicKey)
	if err != nil {
		return false, types.ErrInvalidPublicKey.Wrapf(
			"Cannot decode Ed25519 public key %s",
			publicKey,
		)
	}

	// Decode Signatures
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	isValidSignature := ed25519.Verify(publicKeyBytes, documentBytes, signatureBytes)
	return isValidSignature, nil
}

func verifySecp256k1(publicKey string, signature string, documentBytes []byte) (bool, error) {
	// Decode Public Key
	_, publicKeyBytes, err := multibase.Decode(publicKey)
	if err != nil {
		return false, err
	}
	var pubKeyObj secp256k1.PubKey = publicKeyBytes

	// Decode and Parse Signature
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}

	isValidSignature := pubKeyObj.VerifySignature(documentBytes, signatureBytes)
	return isValidSignature, nil
}
