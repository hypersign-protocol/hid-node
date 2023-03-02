package verification

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"

	secp256k1 "github.com/tendermint/tendermint/crypto/secp256k1"

	// Ethereum based libraries
	etheraccounts "github.com/ethereum/go-ethereum/accounts"
	etherhexutil "github.com/ethereum/go-ethereum/common/hexutil"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"
)

func verifyAll(extendedVmList []*types.ExtendedVerificationMethod, data []byte) error {
	for _, extendedVm := range extendedVmList {
		err := verify(extendedVm, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func verifyAny(extendedVmList []*types.ExtendedVerificationMethod, data []byte) bool {
	found := false

	for _, extendedVm := range extendedVmList {
		err := verify(extendedVm, data)
		if err == nil {
			found = true
			break
		}
	}

	return found
}

func verify(extendedVm *types.ExtendedVerificationMethod, data []byte) error {
	var verificationKey string
	var signature string = extendedVm.Signature

	if extendedVm.PublicKeyMultibase != "" {
		verificationKey = extendedVm.PublicKeyMultibase
	} else if extendedVm.BlockchainAccountId != "" {
		verificationKey = extendedVm.BlockchainAccountId
	} else {
		return fmt.Errorf("either publicKeyMultibase or BlockchainAccountId must be present")
	}

	switch extendedVm.Type {
	case types.Ed25519VerificationKey2020:
		return verifyEd25519(verificationKey, signature, data)
	case types.EcdsaSecp256k1VerificationKey2019:
		return verifySecp256k1(verificationKey, signature, data)
	case types.EcdsaSecp256k1RecoveryMethod2020:
		chain := getCAIP10Chain(verificationKey)
		// Check for supported chains
		switch chain {
		// Ethereum based chains
		case types.EIP155:
			return recoverEthPublicKey(verificationKey, signature, data)
		default:
			return fmt.Errorf("unsupported blockchain address: %s", verificationKey)
		}
	default:
		return fmt.Errorf("unsupported verification method: %s", extendedVm.Type)
	}
}

func verifyEd25519(publicKey string, signature string, documentBytes []byte) error {
	// Decode Public Key
	_, publicKeyBytes, err := multibase.Decode(publicKey)
	if err != nil {
		return fmt.Errorf(
			"cannot decode Ed25519 public key %s",
			publicKey,
		)
	}

	// Decode Signatures
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	if !ed25519.Verify(publicKeyBytes, documentBytes, signatureBytes) {
		return fmt.Errorf("ed25519: signature could not be verified")
	} else {
		return nil
	}
}

func verifySecp256k1(publicKey string, signature string, documentBytes []byte) error {
	// Decode Public Key
	_, publicKeyBytes, err := multibase.Decode(publicKey)
	if err != nil {
		return err
	}
	var pubKeyObj secp256k1.PubKey = publicKeyBytes

	// Decode and Parse Signature
	signatureBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	if !pubKeyObj.VerifySignature(documentBytes, signatureBytes) {
		return fmt.Errorf("secp256k1: signature could not be verified")
	} else {
		return nil
	}
}

// Support for EIP155 based blockchain address
func recoverEthPublicKey(blockchainAccountId string, signature string, documentBytes []byte) error {
	// Extract blockchain address from blockchain account id
	blockchainAddress := getBlockchainAddress(blockchainAccountId)

	// Convert message bytes to hash
	// More info on the `personal_sign` here: https://docs.metamask.io/guide/signing-data.html#personal-sign
	msgHash := etheraccounts.TextHash(documentBytes)

	// Decode hex-encoded signature string to bytes
	signatureBytes, err := etherhexutil.Decode(signature)
	if err != nil {
		return err
	}

	// Handle the signature recieved from web3-js client package by subtracting 27 from the recovery byte
	if signatureBytes[ethercrypto.RecoveryIDOffset] == 27 || signatureBytes[ethercrypto.RecoveryIDOffset] == 28 {
		signatureBytes[ethercrypto.RecoveryIDOffset] -= 27
	}

	// Recover public key from signature
	recoveredPublicKey, err := ethercrypto.SigToPub(msgHash, signatureBytes)
	if err != nil {
		return err
	}

	// Convert public key to hex-encoded address
	recoveredBlockchainAddress := ethercrypto.PubkeyToAddress(*recoveredPublicKey).Hex()

	// Match the recovered address against user provided address
	if recoveredBlockchainAddress != blockchainAddress {
		return fmt.Errorf("eth-recovery-method-secp256k1: signature could not be verified")
	} else {
		return nil
	}
}
