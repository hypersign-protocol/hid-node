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

func verifyAll(extendedVmList []*types.ExtendedVerificationMethod, ssiMsg types.SsiMsg) error {
	for _, extendedVm := range extendedVmList {
		err := verify(extendedVm, ssiMsg)
		if err != nil {
			return err
		}
	}
	return nil
}

func verifyAny(extendedVmList []*types.ExtendedVerificationMethod, ssiMsg types.SsiMsg) bool {
	found := false

	for _, extendedVm := range extendedVmList {
		err := verify(extendedVm, ssiMsg)
		if err == nil {
			found = true
			break
		}
	}

	return found
}

func verify(extendedVm *types.ExtendedVerificationMethod, ssiMsg types.SsiMsg) error {
	docBytes, err := getDocBytesByClientSpec(ssiMsg, extendedVm)
	if err != nil {
		return err
	}

	switch extendedVm.Type {
	case types.Ed25519VerificationKey2020:
		return verifyEd25519VerificationKey2020Key(extendedVm, docBytes)
	case types.EcdsaSecp256k1VerificationKey2019:
		return verifyEcdsaSecp256k1VerificationKey2019Key(extendedVm, docBytes)
	case types.EcdsaSecp256k1RecoveryMethod2020:
		return verifyEcdsaSecp256k1RecoveryMethod2020Key(extendedVm, docBytes)
	default:
		return fmt.Errorf("unsupported verification method: %s", extendedVm.Type)
	}
}

// verifyEcdsaSecp256k1RecoveryMethod2020Key verifies the verification key for verification method type EcdsaSecp256k1RecoveryMethod2020
func verifyEcdsaSecp256k1RecoveryMethod2020Key(extendedVm *types.ExtendedVerificationMethod, documentBytes []byte) error {
	extractedCAIP10Prefix, err := getCAIP10Prefix(extendedVm.BlockchainAccountId)
	if err != nil {
		return err
	}

	switch extractedCAIP10Prefix {
	case types.EthereumCAIP10Prefix:
		return verifyEthereumBlockchainAccountId(
			extendedVm,
			documentBytes,
		)
	default:
		return fmt.Errorf(
			"unsupported CAIP-10 prefix: '%v', supported CAIP-10 prefixes for verification method type %v: %v",
			extractedCAIP10Prefix,
			extendedVm.Type,
			types.CAIP10PrefixForEcdsaSecp256k1RecoveryMethod2020,
		)
	}
}

// verifyEd25519VerificationKey2020Key verifies the verification key for verification method type Ed25519VerificationKey2020
func verifyEd25519VerificationKey2020Key(extendedVm *types.ExtendedVerificationMethod, documentBytes []byte) error {
	// Decode Public Key
	_, publicKeyBytes, err := multibase.Decode(extendedVm.PublicKeyMultibase)
	if err != nil {
		return fmt.Errorf(
			"cannot decode Ed25519 public key %s",
			extendedVm.PublicKeyMultibase,
		)
	}

	// Decode Signatures
	signatureBytes, err := base64.StdEncoding.DecodeString(extendedVm.Signature)
	if err != nil {
		return err
	}

	if !ed25519.Verify(publicKeyBytes, documentBytes, signatureBytes) {
		return fmt.Errorf("signature could not be verified for verificationMethodId: %v", extendedVm.Id)
	} else {
		return nil
	}
}

// verifyEcdsaSecp256k1VerificationKey2019Key verifies the verification key for verification method type EcdsaSecp256k1VerificationKey2019
func verifyEcdsaSecp256k1VerificationKey2019Key(extendedVm *types.ExtendedVerificationMethod, documentBytes []byte) error {
	// Decode and Parse Signature
	signatureBytes, err := base64.StdEncoding.DecodeString(extendedVm.Signature)
	if err != nil {
		return err
	}

	// Decode Public Key
	_, publicKeyBytes, err := multibase.Decode(extendedVm.PublicKeyMultibase)
	if err != nil {
		return err
	}
	var pubKeyObj secp256k1.PubKey = publicKeyBytes

	// Check if the signature is valid for given publicKeyMultibase
	if !pubKeyObj.VerifySignature(documentBytes, signatureBytes) {
		return fmt.Errorf("signature could not be verified for verificationMethodId: %v", extendedVm.Id)
	}

	// Check if blockchainAccountId is passed
	if extendedVm.BlockchainAccountId != "" {
		extractedCAIP10Prefix, err := getCAIP10Prefix(extendedVm.BlockchainAccountId)
		if err != nil {
			return err
		}

		switch extractedCAIP10Prefix {
		case types.CosmosCAIP10Prefix:
			return verifyCosmosBlockchainAccountId(
				extendedVm.BlockchainAccountId,
				extendedVm.PublicKeyMultibase,
			)
		default:
			return fmt.Errorf(
				"unsupported CAIP-10 prefix: '%v', supported CAIP-10 prefixes for verification method type %v: %v",
				extractedCAIP10Prefix,
				extendedVm.Type,
				types.CAIP10PrefixForEcdsaSecp256k1VerificationKey2019,
			)
		}
	}

	return nil
}

// verifyCosmosBlockchainAccountId verifies Cosmos Ecosystem based blockchain address. The verified
// publicKeyMultibase is converted to a bech32 encoded blockchain address which is then compared with the
// user provided blockchain address. If they do not match, error is returned.
func verifyCosmosBlockchainAccountId(blockchainAccountId, publicKeyMultibase string) error {
	// Check if the blockchainAccountId prefix is valid
	extractedCAIP10Prefix, err := getCAIP10Prefix(blockchainAccountId)
	if err != nil {
		return err
	}
	if extractedCAIP10Prefix != types.CosmosCAIP10Prefix {
		return fmt.Errorf(
			"expected CAIP-10 prefix to be '%v', got '%v'",
			types.CosmosCAIP10Prefix,
			extractedCAIP10Prefix,
		)
	}

	// Decode public key
	_, publicKeyBytes, err := multibase.Decode(publicKeyMultibase)
	if err != nil {
		return err
	}

	// Convert publicKeyMultibase to bech32 encoded blockchain address
	chainId, err := getChainIdFromBlockchainAccountId(blockchainAccountId)
	if err != nil {
		return err
	}
	validAddressPrefix := types.CosmosCAIP10ChainIdBech32PrefixMap[chainId]
	convertedAddress, err := publicKeyToCosmosBech32Address(validAddressPrefix, publicKeyBytes)
	if err != nil {
		return err
	}

	// Compare converted blockchain address with user provided blockchain address
	inputAddress, err := getBlockchainAddress(blockchainAccountId)
	if err != nil {
		return err
	}
	if convertedAddress != inputAddress {
		return fmt.Errorf(
			"blockchain address provided in blockchainAccountId '%v' is unexpected",
			blockchainAccountId,
		)
	} else {
		return nil
	}
}

// verifyEthereumBlockchainAccountId verifies Ethereum Ecosystem based blockchain address. A secp256k1 based
// publicKey is extracted from the recoverable Secp256k1 signature. It is converted into a hex encoded based
// blockchain address, and matched with user provided blockchain address. If they do not match, error is returned.
func verifyEthereumBlockchainAccountId(extendedVm *types.ExtendedVerificationMethod, documentBytes []byte) error {
	// Extract blockchain address from blockchain account id
	blockchainAddress, err := getBlockchainAddress(extendedVm.BlockchainAccountId)
	if err != nil {
		return err
	}

	// Convert message bytes to hash
	// More info on the `personal_sign` here: https://docs.metamask.io/guide/signing-data.html#personal-sign
	msgHash := etheraccounts.TextHash(documentBytes)

	// Decode hex-encoded signature string to bytes
	signatureBytes, err := etherhexutil.Decode(extendedVm.Signature)
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

	// Convert public key to b-encoded address
	recoveredBlockchainAddress := ethercrypto.PubkeyToAddress(*recoveredPublicKey).Hex()

	// Match the recovered address against user provided address
	if recoveredBlockchainAddress != blockchainAddress {
		return fmt.Errorf("eth-recovery-method-secp256k1: signature could not be verified")
	} else {
		return nil
	}
}
