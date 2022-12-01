package keeper

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	docVerify "github.com/hypersign-protocol/hid-node/x/ssi/keeper/document_verification"
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

func (k msgServer) VerifySignatureOnDidUpdate(ctx *sdk.Context, oldDIDDoc *types.Did, newDIDDoc *types.Did, signatures []*types.SignInfo) error {
	var signers []types.Signer

	oldController := oldDIDDoc.Controller
	if len(oldController) == 0 {
		oldController = []string{oldDIDDoc.Id}
	}

	for _, controller := range oldController {
		signers = append(signers, types.Signer{Signer: controller})
	}

	for _, oldVM := range oldDIDDoc.VerificationMethod {
		newVM := utils.FindVerificationMethod(newDIDDoc.VerificationMethod, oldVM.Id)

		// Verification Method has been deleted
		if newVM == nil {
			signers = appendSignerIfNeed(signers, oldVM.Controller, newDIDDoc)
			continue
		}

		// Verification Method has been changed
		if !reflect.DeepEqual(oldVM, newVM) {
			signers = appendSignerIfNeed(signers, newVM.Controller, newDIDDoc)
		}

		// Verification Method Controller has been changed, need to add old controller
		if newVM.Controller != oldVM.Controller {
			signers = appendSignerIfNeed(signers, oldVM.Controller, newDIDDoc)
		}
	}

	if err := k.VerifyDidSignature(ctx, newDIDDoc, signers, signatures); err != nil {
		return err
	}

	return nil
}

func (k *Keeper) VerifyDidSignature(ctx *sdk.Context, msg *types.Did, signers []types.Signer, signatures []*types.SignInfo) error {
	var validArr []types.ValidDid

	if len(signers) == 0 {
		return types.ErrInvalidSignature.Wrap("At least one signer should be present")
	}

	if len(signatures) == 0 {
		return types.ErrInvalidSignature.Wrap("At least one signature should be present")
	}

	signingInput := msg.GetSignBytes()

	for _, signer := range signers {
		if signer.VerificationMethod == nil {
			didDoc, err := k.GetDid(ctx, signer.Signer)
			if err != nil {
				return types.ErrDidDocNotFound.Wrap(signer.Signer)
			}

			signer.Authentication = didDoc.DidDocument.Authentication
			signer.VerificationMethod = didDoc.DidDocument.VerificationMethod
		}

		valid, err := VerifyIdentitySignature(signer, signatures, signingInput)
		if err != nil {
			return sdkerrors.Wrap(types.ErrInvalidSignature, err.Error())
		}
		validArr = append(validArr, types.ValidDid{DidId: signer.Signer, IsValid: valid})
	}

	validDid := docVerify.HasAtleastOneTrueSigner(validArr)

	if validDid == (types.ValidDid{}) {
		return sdkerrors.Wrap(types.ErrInvalidSignature, validDid.DidId)
	}

	return nil
}

// Verify Signature for Credential Schema and Credential Status Documents
func (k *Keeper) VerifyDocumentSignature(ctx *sdk.Context, msg types.IdentityMsg, didDoc *types.Did, signatures []*types.SignInfo) error {
	var validArr []types.ValidDid
	signingInput := msg.GetSignBytes()
	signers := didDoc.GetSigners()

	for _, signer := range signers {
		if signer.VerificationMethod == nil {
			fetchedDidDoc, err := k.GetDid(ctx, signer.Signer)
			if err != nil {
				return types.ErrDidDocNotFound.Wrap(signer.Signer)
			}

			signer.Authentication = fetchedDidDoc.DidDocument.Authentication
			signer.VerificationMethod = fetchedDidDoc.DidDocument.VerificationMethod
		}

		valid, err := VerifyIdentitySignature(signer, signatures, signingInput)
		if err != nil {
			return sdkerrors.Wrap(types.ErrInvalidSignature, err.Error())
		}
		validArr = append(validArr, types.ValidDid{DidId: signer.Signer, IsValid: valid})
	}

	validDid := docVerify.HasAtleastOneTrueSigner(validArr)

	if validDid == (types.ValidDid{}) {
		return sdkerrors.Wrap(types.ErrInvalidSignature, validDid.DidId)
	}

	return nil
}

func (k msgServer) ValidateDidControllers(ctx *sdk.Context, id string, controllers []string, verMethods []*types.VerificationMethod) error {

	for _, verificationMethod := range verMethods {
		if err := k.validateController(ctx, id, verificationMethod.Controller); err != nil {
			return err
		}
	}

	for _, didController := range controllers {
		if err := k.validateController(ctx, id, didController); err != nil {
			return err
		}
	}
	return nil
}

func (k *Keeper) validateController(ctx *sdk.Context, id string, controller string) error {
	if id == controller {
		return nil
	}
	didDoc, err := k.GetDid(ctx, controller)
	if err != nil {
		return types.ErrDidDocNotFound.Wrap(controller)
	}
	if len(didDoc.DidDocument.Authentication) == 0 {
		return types.ErrBadRequestInvalidVerMethod.Wrap(
			fmt.Sprintf("Verificatition method controller %s doesn't have an authentication keys", controller))
	}
	return nil
}

func appendSignerIfNeed(signers []types.Signer, controller string, msg *types.Did) []types.Signer {
	for _, signer := range signers {
		if signer.Signer == controller {
			return signers
		}
	}

	signer := types.Signer{
		Signer: controller,
	}

	if controller == msg.Id {
		signer.VerificationMethod = msg.VerificationMethod
		signer.Authentication = msg.Authentication
	}

	return append(signers, signer)
}
