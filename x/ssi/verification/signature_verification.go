package verification

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
)

// Verify signatures against signer's public keys
// If atleast one of the signatures is valid, return true
func VerifyIdentitySignature(signer types.Signer, signatures []*types.SignInfo, signingInput []byte) (bool, error) {
	result := false

	for _, signature := range signatures {
		did, _ := utils.SplitDidUrlIntoDid(signature.VerificationMethodId)
		if did == signer.Signer {
			pubKey, vmType, err := utils.FindPublicKeyAndVerificationMethodType(signer, signature.VerificationMethodId)
			if err != nil {
				return false, err
			}

			result, err = verify(vmType, pubKey, signature.Signature, signingInput)
			if err != nil {
				return false, err
			}
		}
	}

	return result, nil
}

func VerifyDidSignature(ctx *sdk.Context, didDocBytes []byte, signers []types.Signer, signatures []*types.SignInfo) error {
	var validArr []types.ValidDid

	if len(signers) == 0 {
		return types.ErrInvalidSignature.Wrap("At least one signer should be present")
	}

	if len(signatures) == 0 {
		return types.ErrInvalidSignature.Wrap("At least one signature should be present")
	}

	for _, signer := range signers {
		valid, err := VerifyIdentitySignature(signer, signatures, didDocBytes)
		if err != nil {
			return sdkerrors.Wrap(types.ErrInvalidSignature, err.Error())
		}
		validArr = append(validArr, types.ValidDid{DidId: signer.Signer, IsValid: valid})
	}

	validDid := HasAtleastOneTrueSigner(validArr)

	if validDid == (types.ValidDid{}) {
		return sdkerrors.Wrap(types.ErrInvalidSignature, validDid.DidId)
	}

	return nil
}

func DocumentProofTypeCheck(inputProofType string, signers []types.Signer, vmId string) error {
	var vmType string
	var expectedProofType string
	
	for i := 0; i < len(signers); i++ {
		if signers[i].VerificationMethod[0].Id == vmId {
			vmType = signers[i].VerificationMethod[0].Type
			break
		}
	}

	if vmType == "" {
		return types.ErrVerificationMethodNotFound.Wrap(vmId)
	}

	expectedProofType = VerificationKeySignatureMap[vmType]
	if inputProofType != expectedProofType {
		return fmt.Errorf(
			"expected document proof type for verification method type %s to be '%s', recieved '%s'",
			vmType,
			expectedProofType,
			inputProofType,
		)
	}
	return nil
}

// Verify Signature for Credential Schema and Credential Status Documents
func VerifyDocumentSignature(ctx *sdk.Context, docBytes []byte, signers []types.Signer, signatures []*types.SignInfo) error {
	var validArr []types.ValidDid

	for _, signer := range signers {
		valid, err := VerifyIdentitySignature(signer, signatures, docBytes)
		if err != nil {
			return sdkerrors.Wrap(types.ErrInvalidSignature, err.Error())
		}
		validArr = append(validArr, types.ValidDid{DidId: signer.Signer, IsValid: valid})
	}

	validDid := HasAtleastOneTrueSigner(validArr)

	if validDid == (types.ValidDid{}) {
		return sdkerrors.Wrap(types.ErrInvalidSignature, validDid.DidId)
	}

	return nil
}
