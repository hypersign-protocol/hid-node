package keeper

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"reflect"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
)

func VerifyIdentitySignature(signer types.Signer, signatures []*types.SignInfo, signingInput []byte) (bool, error) {
	result := false
	matchFound := false

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

			result = ed25519.Verify(pubKey, signingInput, signature)
			matchFound = true
		}
	}

	if !matchFound {
		return false, fmt.Errorf("signature for %s not found", signer.Signer)
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
			signers = AppendSignerIfNeed(signers, oldVM.Controller, newDIDDoc)
			continue
		}

		// Verification Method has been changed
		if !reflect.DeepEqual(oldVM, newVM) {
			signers = AppendSignerIfNeed(signers, newVM.Controller, newDIDDoc)
		}

		// Verification Method Controller has been changed, need to add old controller
		if newVM.Controller != oldVM.Controller {
			signers = AppendSignerIfNeed(signers, oldVM.Controller, newDIDDoc)
		}
	}

	if err := k.VerifySignature(ctx, newDIDDoc, signers, signatures); err != nil {
		return err
	}

	return nil
}

func AppendSignerIfNeed(signers []types.Signer, controller string, msg *types.Did) []types.Signer {
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

func (k *Keeper) VerifySignature(ctx *sdk.Context, msg *types.Did, signers []types.Signer, signatures []*types.SignInfo) error {
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

			signer.Authentication = didDoc.Did.Authentication
			signer.VerificationMethod = didDoc.Did.VerificationMethod
		}

		valid, err := VerifyIdentitySignature(signer, signatures, signingInput)
		if err != nil {
			return sdkerrors.Wrap(types.ErrInvalidSignature, err.Error())
		}

		if !valid {
			return types.ErrInvalidSignature
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

func (k *Keeper) ValidateController(ctx *sdk.Context, id string, controller string) error {
	if id == controller {
		return nil
	}
	didDoc, err := k.GetDid(ctx, controller)
	if err != nil {
		return types.ErrDidDocNotFound.Wrap(controller)
	}
	if len(didDoc.Did.Authentication) == 0 {
		return types.ErrBadRequestInvalidVerMethod.Wrap(
			fmt.Sprintf("Verificatition method controller %s doesn't have an authentication keys", controller))
	}
	return nil
}

func (k msgServer) ValidateDidControllers(ctx *sdk.Context, id string, controllers []string, verMethods []*types.VerificationMethod) error {

	for _, verificationMethod := range verMethods {
		if err := k.ValidateController(ctx, id, verificationMethod.Controller); err != nil {
			return err
		}
	}

	for _, didController := range controllers {
		if err := k.ValidateController(ctx, id, didController); err != nil {
			return err
		}
	}
	return nil
}

// Check the Deactivate status of DID
func VerifyDidDeactivate(metadata *types.Metadata, id string) error {
	if metadata.GetDeactivated() {
		return sdkerrors.Wrap(types.ErrDidDocDeactivated, fmt.Sprintf("DidDoc ID: %s", id))
	}
	return nil
}

// Verify Credential Signature
func (k msgServer) VerifyCredentialSignature(msg *types.CredentialStatus, didDoc *types.Did, signature string, verificationMethod string) error {
	signingInput := msg.GetSignBytes()

	signer := types.Signer{
		Signer:             didDoc.GetId(),
		Authentication:     didDoc.GetAuthentication(),
		VerificationMethod: didDoc.GetVerificationMethod(),
	}

	signingInfo := &types.SignInfo{
		VerificationMethodId: verificationMethod,
		Signature:            signature,
	}

	signingInfoList := []*types.SignInfo{
		signingInfo,
	}

	valid, err := VerifyIdentitySignature(signer, signingInfoList, signingInput)
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidSignature, err.Error())
	}

	if !valid {
		return sdkerrors.Wrap(types.ErrInvalidSignature, signer.Signer)
	}
	return nil
}

func VerifyCredentialStatusDates(issuanceDate time.Time, expirationDate time.Time) error {
	var dateDiff int64 = int64(expirationDate.Sub(issuanceDate)) / 1e9 // converting nanoseconds to seconds
	if dateDiff < 0 {
		return sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("expiration date %s cannot be less than issuance date %s", expirationDate, issuanceDate))
	}

	return nil
}

func VerifyCredentialProofDates(credProof *types.CredentialProof, credRegistration bool) error {
	var dateDiff int64

	proofCreatedDate := credProof.GetCreated()
	proofCreatedDateParsed, err := time.Parse(time.RFC3339, proofCreatedDate)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid created date format: %s", proofCreatedDate))
	}

	proofUpdatedDate := credProof.GetUpdated()
	proofUpdatedDateParsed, err := time.Parse(time.RFC3339, proofUpdatedDate)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid created date format: %s", proofUpdatedDate))
	}

	// If credRegistration is True, check for equity of updated and created dates will proceeed
	// Else, check for updated date being greater than created date will proceeed
	if credRegistration {
		if !proofUpdatedDateParsed.Equal(proofCreatedDateParsed) {
			return sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("updated date %s should be similar to created date %s", proofUpdatedDate, proofCreatedDate))
		}
	} else {
		dateDiff = int64(proofUpdatedDateParsed.Sub(proofCreatedDateParsed)) / 1e9 // converting nanoseconds to seconds
		if dateDiff <= 0 {
			return sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("update date %s cannot be less than or equal to created date %s in case of credential status update", proofUpdatedDate, proofCreatedDate))
		}
	}

	return nil
}
