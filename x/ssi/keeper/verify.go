package keeper

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)



func VerifyIdentitySignature(signer types.Signer, signatures []*types.SignInfo, signingInput []byte) (bool, error) {
	result := true
	foundOne := false

	for _, info := range signatures {
		did, _ := utils.SplitDidUrlIntoDid(info.VerificationMethodId)
		if did == signer.Signer {
			pubKey, err := utils.FindPublicKey(signer, info.VerificationMethodId)
			if err != nil {
				return false, err
			}

			signature, err := base64.StdEncoding.DecodeString(info.Signature)
			if err != nil {
				return false, err
			}

			result = result && ed25519.Verify(pubKey, signingInput, signature)
			foundOne = true
		}
	}

	if !foundOne {
		return false, fmt.Errorf("signature %s not found", signer.Signer)
	}

	return result, nil
}

func (k *Keeper) VerifySignature(ctx *sdk.Context, msg types.IdentityMsg, signers []types.Signer, signatures []*types.SignInfo) error {
	if len(signers) == 0 {
		return types.ErrInvalidSignature.Wrap("At least one signer should be present")
	}

	if len(signatures) == 0 {
		return types.ErrInvalidSignature.Wrap("At least one signature should be present")
	}

	signingInput := msg.GetSignBytes()

	for _, signer := range signers {
		// TODO: Understand stateValue part of it on Cheqd
		// if signer.PublicKeyStruct == nil {
		// 	state, err := k.GetDid(ctx, signer.Signer)
		// 	if err != nil {
		// 		return types.ErrDidDocNotFound.Wrap(signer.Signer)
		// 	}

		// 	didDoc, err := state.UnpackDataAsDid()
		// 	if err != nil {
		// 		return types.ErrDidDocNotFound.Wrap(signer.Signer)
		// 	}

		// 	signer.Authentication = didDoc.Authentication
		// 	signer.PublicKeyStruct = didDoc.PublicKeyStruct
		// }

		valid, err := VerifyIdentitySignature(signer, signatures, signingInput)
		if err != nil {
			return sdkerrors.Wrap(types.ErrInvalidSignature, err.Error())
		}

		if !valid {
			return sdkerrors.Wrap(types.ErrInvalidSignature, signer.Signer)
		}
	}

	return nil
}
