package cli

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"

	secp256k1 "github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
	"golang.org/x/crypto/ripemd160" // nolint: staticcheck

	etheraccounts "github.com/ethereum/go-ethereum/accounts"
	etherhexutil "github.com/ethereum/go-ethereum/common/hexutil"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"

	bbs "github.com/hyperledger/aries-framework-go/component/kmscrypto/crypto/primitive/bbs12381g2pub"

	"github.com/iden3/go-iden3-crypto/babyjub"
)

// GetBbsBlsSignature2020 signs a message and returns a BBS signature
func GetBbsBlsSignature2020(privateKey string, message []byte) (string, error) {
	privKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		panic(err)
	}

	bbsObj := bbs.New()

	signatureBytes, err := bbsObj.Sign([][]byte{message}, privKeyBytes)
	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}

// Get BabyJubJub Signature
func GetBJJSignature2021(privateKey string, message []byte) (string, error) {
	// Decode private key from hex
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		panic(err)
	}
	var privateKeyBytes32 [32]byte
	copy(privateKeyBytes32[:], privateKeyBytes)

	var privKeyObj babyjub.PrivateKey = privateKeyBytes32

	msgBigInt := new(big.Int).SetBytes(message)

	// Get Signature
	signatureObj := privKeyObj.SignPoseidon(msgBigInt)
	signatureHex := signatureObj.Compress().String()

	// Convert Signature to multibase base58
	signatureBytes, err := hex.DecodeString(signatureHex)
	if err != nil {
		return "", err
	}

	signature, err := multibase.Encode(multibase.Base58BTC, signatureBytes)
	if err != nil {
		return "", err
	}

	return signature, nil
}

func GetEcdsaSecp256k1RecoverySignature2020(privateKey string, message []byte) (string, error) {
	// Decode key into bytes
	privKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		panic(err)
	}
	privKeyObject, err := ethercrypto.ToECDSA(privKeyBytes)
	if err != nil {
		panic(err)
	}

	// Hash the message
	msgHash := etheraccounts.TextHash(message)

	// Sign Message
	sigBytes, err := ethercrypto.Sign(msgHash, privKeyObject)
	if err != nil {
		panic(err)
	}

	return etherhexutil.Encode(sigBytes), nil
}

func GetEcdsaSecp256k1Signature2019(privateKey string, message []byte) (string, error) {
	// Decode key into bytes
	privKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	// Convert private key string to Secp256k1 object
	var privKeyObject secp256k1.PrivKey = privKeyBytes

	// Sign Message
	signature, err := privKeyObject.Sign(message)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

func GetEd25519Signature2020(privateKey string, message []byte) (string, error) {
	// Decode key into bytes
	privKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	// Sign Message
	signatureBytes := ed25519.Sign(privKeyBytes, message)

	return multibase.Encode(multibase.Base58BTC, signatureBytes)
}

func getDocumentProofs(ctx client.Context, proofStrings []string) ([]*types.DocumentProof, error) {
	var documentProofs []*types.DocumentProof

	for i := 0; i < len(proofStrings); i++ {
		var documentProof types.DocumentProof
		err := ctx.Codec.UnmarshalJSON([]byte(proofStrings[i]), &documentProof)
		if err != nil {
			return nil, fmt.Errorf("unable to process the proof: %v", proofStrings[i])
		}

		// Get the VM Ids
		documentProofs = append(documentProofs, &documentProof)
	}

	return documentProofs, nil
}

// validateDidAliasSignerAddress checks if the signer address provided in the --from flag matches
// the address extracted from the publicKeyMultibase
func validateDidAliasSignerAddress(fromSignerAddress, publicKeyMultibase string) error {
	// Decode public key
	_, publicKeyBytes, err := multibase.Decode(publicKeyMultibase)
	if err != nil {
		return err
	}

	// Throw error if the length of secp256k1 publicKey is not 33
	if len(publicKeyBytes) != 33 {
		return fmt.Errorf("invalid secp256k1 public key length %v", len(publicKeyBytes))
	}

	// Hash pubKeyBytes as: RIPEMD160(SHA256(public_key_bytes))
	pubKeySha256Hash := sha256.Sum256(publicKeyBytes)
	ripemd160hash := ripemd160.New()
	ripemd160hash.Write(pubKeySha256Hash[:])
	addressBytes := ripemd160hash.Sum(nil)

	// Convert addressBytes to bech32 encoded address
	convertedAddress, err := bech32.ConvertAndEncode("hid", addressBytes)
	if err != nil {
		return err
	}

	if fromSignerAddress != convertedAddress {
		return fmt.Errorf("transaction signer address is not the author of DID Document alias")
	}

	return nil
}
