package keeper

import (
	"crypto/ed25519"
	sha256 "crypto/sha256"
	"encoding/base64"
	"fmt"

	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
	ecdsa "github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
)

func verify(verificationMethodType string, verificationKey string, signature string, data []byte) (bool, error) {
	switch verificationMethodType {
	case EcdsaSecp256k1VerificationKey2019:
		return verifySecp256k1(verificationKey, signature, data)
	case Ed25519VerificationKey2020:
		return verifyEd25519(verificationKey, signature, data)
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
	pubKeyObject, err := secp256k1.ParsePubKey(publicKeyBytes)
	if err != nil {
		return false, err
	}

	// Decode and Parse Signature
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, err
	}
	signatureSecpObject, err := ecdsa.ParseDERSignature(signatureBytes)
	if err != nil {
		return false, err
	}

	// Hash Document Bytes with SHA-256
	documentHash := sha256.Sum256(documentBytes)

	isValidSignature := signatureSecpObject.Verify(documentHash[:], pubKeyObject)
	return isValidSignature, nil
}
