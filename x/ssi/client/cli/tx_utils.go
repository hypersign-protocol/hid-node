package cli

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/types/bech32"
	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
	secp256k1 "github.com/tendermint/tendermint/crypto/secp256k1"
	"golang.org/x/crypto/ripemd160" // nolint: staticcheck
	"golang.org/x/crypto/sha3"

	etheraccounts "github.com/ethereum/go-ethereum/accounts"
	etherhexutil "github.com/ethereum/go-ethereum/common/hexutil"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"

	bbs "github.com/hyperledger/aries-framework-go/component/kmscrypto/crypto/primitive/bbs12381g2pub"

	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/spf13/cobra"
)

// Extract Verification Method Ids and their respective signatures from Arguments
// NOTE: Only Verificaiton Method Ids, Signatures and Signing Algorithms are supposed to be passed,
// and the sequence of arguments are to be preserved.
func extractDIDSigningElements(cmdArgs []string) ([]DIDSigningElements, error) {
	// Since, a trio of VM Id, Siganature and Signing Algorithm are expected, an error should be thrown
	// if the number of argumens isn't a multiple of 3
	nArgs := len(cmdArgs)
	if nArgs%3 != 0 {
		return nil, fmt.Errorf("unexpected number of arguments recieved")
	}

	var didSigningElementsList []DIDSigningElements

	for i := 0; i < nArgs; i += 3 {
		didSigningElementsList = append(
			didSigningElementsList,
			DIDSigningElements{
				VerificationMethodId: cmdArgs[i],
				SignKey:              cmdArgs[i+1],
				SignAlgo:             cmdArgs[i+2],
			},
		)
	}

	return didSigningElementsList, nil
}

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
func GetBJJSignature(privateKey string, message []byte) (string, error) {
	// Decode private key from hex
	privateKeyBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		panic(err)
	}
	var privateKeyBytes32 [32]byte
	copy(privateKeyBytes32[:], privateKeyBytes)

	var privKeyObj babyjub.PrivateKey = privateKeyBytes32

	// SHA-224 the message and convert it to BigInt
	msgHash := sha3.Sum224(message)
	msgBigInt := new(big.Int).SetBytes(msgHash[:])

	// Get Signature
	signatureObj := privKeyObj.SignPoseidon(msgBigInt)
	signatureHex := signatureObj.Compress().String()

	return signatureHex, nil
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

func getSignatures(cmd *cobra.Command, didDoc *types.Did, cmdArgs []string) ([]*types.SignInfo, error) {
	var signInfoList []*types.SignInfo

	didSigningElementsList, err := extractDIDSigningElements(cmdArgs)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(didSigningElementsList); i++ {
		// Get the VM Ids
		signInfoList = append(signInfoList, &types.SignInfo{
			VerificationMethodId: didSigningElementsList[i].VerificationMethodId,
		})

		// Sign based on the Signing Algorithm
		switch didSigningElementsList[i].SignAlgo {
		case types.Ed25519Signature2020:
			// Perform EdDSACryptoSuite2020 Normalization
			didDocBytes, err := ldcontext.EdDSACryptoSuite2020Canonize(didDoc)
			if err != nil {
				return nil, err
			}

			signInfoList[i].Signature, err = GetEd25519Signature2020(didSigningElementsList[i].SignKey, didDocBytes[:])
			if err != nil {
				return nil, err
			}
		case types.EcdsaSecp256k1Signature2019:
			didDocBytes, err := ldcontext.EcdsaSecp256k1Signature2019Canonize(didDoc)
			if err != nil {
				return nil, err
			}

			signInfoList[i].Signature, err = GetEcdsaSecp256k1Signature2019(didSigningElementsList[i].SignKey, didDocBytes)
			if err != nil {
				return nil, err
			}
		case types.EcdsaSecp256k1RecoverySignature2020:
			didDocBytes, err := ldcontext.EcdsaSecp256k1RecoverySignature2020Canonize(didDoc)
			if err != nil {
				return nil, err
			}

			signInfoList[i].Signature, err = GetEcdsaSecp256k1RecoverySignature2020(didSigningElementsList[i].SignKey, didDocBytes)
			if err != nil {
				return nil, err
			}
		case types.BbsBlsSignature2020:
			didDocBytes, err := ldcontext.BbsBlsSignature2020Canonize(didDoc)
			if err != nil {
				return nil, err
			}

			signInfoList[i].Signature, err = GetBbsBlsSignature2020(didSigningElementsList[i].SignKey, didDocBytes)
			if err != nil {
				return nil, err
			}
		case "bjj":
			signInfoList[i].Signature, err = GetBJJSignature(didSigningElementsList[i].SignKey, didDoc.GetSignBytes())
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf(
				"unsupported signing algorithm %s, supported signing algorithms: ['Ed25519Signature2020', 'secp256k1', 'recover-eth', 'bbs', 'bjj']",
				didSigningElementsList[i].SignAlgo,
			)
		}
	}

	return signInfoList, nil
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
