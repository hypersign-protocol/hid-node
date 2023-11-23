package cmd

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/ripemd160" //nolint: staticcheck

	bech32 "github.com/cosmos/cosmos-sdk/types/bech32"
	hidnodecli "github.com/hypersign-protocol/hid-node/x/ssi/client/cli"
	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// publicKeyToBech32Address converts publicKey byteArray to Bech32 encoded blockchain address
func publicKeyToBech32Address(addressPrefix string, pubKeyBytes []byte) string {
	// Throw error if the length of secp256k1 publicKey is not 33
	if len(pubKeyBytes) != 33 {
		panic(fmt.Sprintf("invalid secp256k1 public key length %v", len(pubKeyBytes)))
	}

	// Hash pubKeyBytes as: RIPEMD160(SHA256(public_key_bytes))
	pubKeySha256Hash := sha256.Sum256(pubKeyBytes)
	ripemd160hash := ripemd160.New()
	ripemd160hash.Write(pubKeySha256Hash[:])
	addressBytes := ripemd160hash.Sum(nil)

	// Convert addressBytes to bech32 encoded address
	address, err := bech32.ConvertAndEncode(addressPrefix, addressBytes)
	if err != nil {
		panic(err)
	}
	return address
}

// getDocumentSignature returns signature for the input SSI Document
func getDocumentSignature(doc types.SsiMsg, docProof *types.DocumentProof, privateKey string) (string, error) {
	var signature string

	switch docProof.Type {
	case types.Ed25519Signature2020:
		var docBytes []byte
		docBytes, err := ldcontext.Ed25519Signature2020Normalize(doc, docProof)
		if err != nil {
			return "", err
		}

		signature, err = hidnodecli.GetEd25519Signature2020(privateKey, docBytes)
		if err != nil {
			return "", err
		}
	case types.EcdsaSecp256k1Signature2019:
		var docBytes []byte
		docBytes, err := ldcontext.EcdsaSecp256k1Signature2019Normalize(doc, docProof)
		if err != nil {
			return "", err
		}

		signature, err = hidnodecli.GetEcdsaSecp256k1Signature2019(privateKey, docBytes)
		if err != nil {
			return "", err
		}
	case types.EcdsaSecp256k1RecoverySignature2020:
		var docBytes []byte
		docBytes, err := ldcontext.EcdsaSecp256k1RecoverySignature2020Normalize(doc, docProof)
		if err != nil {
			return "", err
		}

		signature, err = hidnodecli.GetEcdsaSecp256k1RecoverySignature2020(privateKey, docBytes)
		if err != nil {
			return "", err
		}
	case types.BbsBlsSignature2020:
		var docBytes []byte
		docBytes, err := ldcontext.BbsBlsSignature2020Normalize(doc, docProof)
		if err != nil {
			return "", err
		}

		signature, err = hidnodecli.GetBbsBlsSignature2020(privateKey, docBytes)
		if err != nil {
			return "", err
		}
	case types.BJJSignature2021:
		var docBytes []byte
		docBytes, err := ldcontext.BJJSignature2021Normalize(doc)
		if err != nil {
			return "", err
		}

		signature, err = hidnodecli.GetBJJSignature2021(privateKey, docBytes)
		if err != nil {
			return "", err
		}
	default:
		panic("recieved unsupported signing-algo. Supported algorithms are: [Ed25519Signature2020, EcdsaSecp256k1Signature2019, EcdsaSecp256k1RecoverySignature2020, BbsBlsSignature2020, BJJSignature2021]")
	}

	return signature, nil
}
