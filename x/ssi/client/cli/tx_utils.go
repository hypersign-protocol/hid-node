package cli

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
	ecdsa "github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"

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

func getSecp256k1Signature(privateKey string, message []byte) (string, error) {
	// Decode key into bytes
	privKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	// Convert private key string to Secp256k1 object
	privKeyObject := secp256k1.PrivKeyFromBytes(privKeyBytes)

	// Secp256k1 requires the message
	dataHash := sha256.Sum256(message)

	// Sign Message
	signature := ecdsa.Sign(privKeyObject, dataHash[:])
	serialisedSignature := signature.Serialize()

	return base64.StdEncoding.EncodeToString(serialisedSignature), nil
}

func getEd25519Signature(privateKey string, message []byte) (string, error) {
	// Decode key into bytes
	privKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	// Sign Message
	signatureBytes := ed25519.Sign(privKeyBytes, message)

	return base64.StdEncoding.EncodeToString(signatureBytes), nil
}

func getSignatures(cmd *cobra.Command, message []byte, cmdArgs []string) ([]*types.SignInfo, error) {
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
		case "secp256k1":
			signInfoList[i].Signature, err = getSecp256k1Signature(didSigningElementsList[i].SignKey, message)
			if err != nil {
				return nil, err
			}
		case "ed25519":
			signInfoList[i].Signature, err = getEd25519Signature(didSigningElementsList[i].SignKey, message)
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Unsupported signing algorithm %s", didSigningElementsList[i].SignAlgo)
		}
	}

	return signInfoList, nil
}
