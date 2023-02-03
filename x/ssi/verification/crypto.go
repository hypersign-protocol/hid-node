package verification

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/common"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"

	secp256k1 "github.com/tendermint/tendermint/crypto/secp256k1"

	// Ethereum based libraries
	etheraccounts "github.com/ethereum/go-ethereum/accounts"
	etherhexutil "github.com/ethereum/go-ethereum/common/hexutil"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"
)

func verify(verificationMethodType string, verificationKey string, signature string, data []byte) (bool, error) {
	switch verificationMethodType {
	case common.Ed25519VerificationKey2020:
		return verifyEd25519(verificationKey, signature, data)
	case common.EcdsaSecp256k1VerificationKey2019:
		return verifySecp256k1(verificationKey, signature, data)
	case common.EcdsaSecp256k1RecoveryMethod2020:
		chain := getCAIP10Chain(verificationKey)
		// Check for supported chains
		switch chain {
		// Ethereum based chains
		case common.EIP155:
			return recoverEthPublicKey(verificationKey, signature, data)
		default:
			return false, fmt.Errorf("unsupported blockchain address: %s", verificationKey)
		}
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

// Support for EIP155 based blockchain address
func recoverEthPublicKey(blockchainAccountId string, signature string, documentBytes []byte) (bool, error) {
	var isValidSignature bool

	// Extract blockchain address from blockchain account id
	blockchainAddress := getBlockchainAddress(blockchainAccountId)

	// Convert message bytes to hash
	// More info on the `personal_sign` here: https://docs.metamask.io/guide/signing-data.html#personal-sign
	msgHash := etheraccounts.TextHash(documentBytes)

	// Decode hex-encoded signature string to bytes
	signatureBytes, err := etherhexutil.Decode(signature)
	if err != nil {
		return false, err
	}

	// Handle the signature recieved from web3-js client package by subtracting 27 from the recovery byte
	if signatureBytes[ethercrypto.RecoveryIDOffset] == 27 || signatureBytes[ethercrypto.RecoveryIDOffset] == 28 {
		signatureBytes[ethercrypto.RecoveryIDOffset] -= 27
	} 

	// Recover public key from signature
	recoveredPublicKey, err := ethercrypto.SigToPub(msgHash, signatureBytes)
	if err != nil {
		return false, err
	}

	// Convert public key to hex-encoded address
	recoveredBlockchainAddress := ethercrypto.PubkeyToAddress(*recoveredPublicKey).Hex()

	// Match the recovered address against user provided address
	isValidSignature = recoveredBlockchainAddress == blockchainAddress
	return isValidSignature, nil
}
