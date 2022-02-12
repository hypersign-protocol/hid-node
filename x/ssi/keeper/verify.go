package keeper

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
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

func (k msgServer) VerifySignatureOnDidUpdate(ctx *sdk.Context, oldDIDDoc *types.DidDocStructCreateDID, newDIDDoc *types.DidDocStructUpdateDID, signatures []*types.SignInfo) error {
	var signers []types.Signer

	// TODO: Implement when controller feild is used.
	// Get Old DID Doc controller if it's nil then assign self
	// oldController := oldDIDDoc.Controller
	// if len(oldController) == 0 {
	// 	oldController = []string{oldDIDDoc.Id}
	// }

	// for _, controller := range oldController {
	// 	signers = append(signers, types.Signer{Signer: controller})
	// }

	// TODO: Implement this when `controller` field is added
	// for _, oldVM := range oldDIDDoc.PublicKey {
	// 	newVM := utils.FindVerificationMethod(newDIDDoc.PublicKey, oldVM.Id)

	// 	// Verification Method has been deleted
	// 	if newVM == nil {
	// 		signers = AppendSignerIfNeed(signers, oldVM.Controller, newDIDDoc)
	// 		continue
	// 	}

	// 	// Verification Method has been changed
	// 	if !reflect.DeepEqual(oldVM, newVM) {
	// 		signers = AppendSignerIfNeed(signers, newVM.Controller, newDIDDoc)
	// 	}

	// 	// Verification Method Controller has been changed, need to add old controller
	// 	if newVM.Controller != oldVM.Controller {
	// 		signers = AppendSignerIfNeed(signers, oldVM.Controller, newDIDDoc)
	// 	}
	// }

	if err := k.VerifySignature(ctx, newDIDDoc, signers, signatures); err != nil {
		return err
	}

	return nil
}

// TODO: Implement this when `controller` field is added
// func AppendSignerIfNeed(signers []types.Signer, controller string, msg *types.MsgUpdateDidPayload) []types.Signer {
// 	for _, signer := range signers {
// 		if signer.Signer == controller {
// 			return signers
// 		}
// 	}

// 	signer := types.Signer{
// 		Signer: controller,
// 	}

// 	if controller == msg.Id {
// 		signer.VerificationMethod = msg.VerificationMethod
// 		signer.Authentication = msg.Authentication
// 	}

// 	return append(signers, signer)
// }

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

func (k *Keeper) VerifySignatureOnCreateSchema(ctx *sdk.Context, msg *types.Schema, signers []types.Signer, signatures []*types.SignInfo) error {
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
			// return sdkerrors.Wrap(types.ErrInvalidSignature, signer.Signer)
			return sdkerrors.Wrap(types.ErrInvalidSignature, string(signingInput))
		}
	}

	return nil
}
