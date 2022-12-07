package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	docVerify "github.com/hypersign-protocol/vid-node/x/ssi/document_verification"

	sigVerify "github.com/hypersign-protocol/vid-node/x/ssi/signature"
	"github.com/hypersign-protocol/vid-node/x/ssi/types"
)

func (k msgServer) RegisterCredentialStatus(goCtx context.Context, msg *types.MsgRegisterCredentialStatus) (*types.MsgRegisterCredentialStatusResponse, error) {
	var id uint64

	ctx := sdk.UnwrapSDKContext(goCtx)

	credMsg := msg.GetCredentialStatus()
	credProof := msg.GetProof()

	credId := credMsg.GetClaim().GetId()

	chainNamespace := k.GetChainNamespace(&ctx)

	// Check the format of Credential ID
	err := docVerify.IsValidID(credId, chainNamespace, "credDocument")
	if err != nil {
		return nil, err
	}

	// Check if the credential already exist in the store
	if !k.HasCredential(ctx, credId) {
		// Check for the correct credential status
		credStatus := credMsg.GetClaim().GetCurrentStatus()
		if credStatus != "Live" {
			return nil, sdkerrors.Wrap(types.ErrInvalidCredentialStatus, fmt.Sprintf("expected credential status to be `Live`, got %s", credStatus))
		}

		// Check if the DID of the issuer exists
		issuerId := credMsg.GetIssuer()
		if !k.HasDid(ctx, issuerId) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Issuer`s DID %s doesnt exists", issuerId))
		}

		// Check if issuer's DID is deactivated
		if issuerDidDocument, _ := k.GetDid(&ctx, issuerId); issuerDidDocument.DidDocumentMetadata.Deactivated {
			return nil, sdkerrors.Wrap(types.ErrDidDocDeactivated, fmt.Sprintf("%s is deactivated and cannot used be used to register credential status", issuerDidDocument.DidDocument.Id))
		}

		// Check if the expiration date is not before the issuance date.
		expirationDate := credMsg.GetExpirationDate()
		expirationDateParsed, err := time.Parse(time.RFC3339, expirationDate)
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid expiration date format: %s", expirationDate))
		}

		issuanceDate := credMsg.GetIssuanceDate()
		issuanceDateParsed, err := time.Parse(time.RFC3339, issuanceDate)
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid issuance date format: %s", issuanceDate))
		}

		if err := docVerify.VerifyCredentialStatusDates(issuanceDateParsed, expirationDateParsed); err != nil {
			return nil, err
		}

		// Check if updated date iss imilar to created date
		if err := docVerify.VerifyCredentialProofDates(credProof, true); err != nil {
			return nil, err
		}

		// Check if the created date lies between issuance and expiration
		currentDate, _ := time.Parse(time.RFC3339, credProof.Created)
		if currentDate.After(expirationDateParsed) || currentDate.Before(issuanceDateParsed) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDate, "credential registeration is happening on a date which doesn`t lie between issuance date and expiration date")
		}

		// Check the hash type of credentialHash
		isValidCredentialHash := docVerify.VerifyCredentialHash(credMsg.GetCredentialHash())
		if !isValidCredentialHash {
			return nil, sdkerrors.Wrapf(types.ErrInvalidCredentialHash, "supported hashing algorithms: sha256")
		}

		// Verify the Signature
		didDocument, err := k.GetDid(&ctx, issuerId)
		if err != nil {
			return nil, err
		}

		did := didDocument.GetDidDocument()

		signature := &types.SignInfo{
			VerificationMethodId: credProof.GetVerificationMethod(),
			Signature:            credProof.GetProofValue(),
		}
		signatures := []*types.SignInfo{signature}
		signers := did.GetSigners()
		signersWithVM, err := k.GetVMForSigners(&ctx, signers)
		if err != nil {
			return nil, err
		}

		// Proof Type Check
		err = sigVerify.DocumentProofTypeCheck(credProof.Type, signersWithVM, credProof.VerificationMethod)
		if err != nil {
			return nil, err
		}

		// Verify Signature
		err = sigVerify.VerifyDocumentSignature(&ctx, credMsg, signersWithVM, signatures)
		if err != nil {
			return nil, err
		}

		cred := &types.Credential{
			Claim:          credMsg.GetClaim(),
			Issuer:         credMsg.GetIssuer(),
			IssuanceDate:   credMsg.GetIssuanceDate(),
			ExpirationDate: credMsg.GetExpirationDate(),
			CredentialHash: credMsg.GetCredentialHash(),
			Proof:          credProof,
		}

		id = k.RegisterCred(ctx, cred)

	} else {
		cred, err := k.updateCredentialStatus(ctx, credMsg, credProof)
		if err != nil {
			return nil, err
		}
		id = k.RegisterCred(ctx, cred)
	}

	return &types.MsgRegisterCredentialStatusResponse{Id: id}, nil
}

func (k msgServer) updateCredentialStatus(ctx sdk.Context, newCredStatus *types.CredentialStatus, newCredProof *types.CredentialProof) (*types.Credential, error) {
	credId := newCredStatus.GetClaim().GetId()

	// Get Credential from store
	oldCredStatus, err := k.GetCredential(&ctx, credId)
	if err != nil {
		return nil, err
	}
	oldClaimStatus := oldCredStatus.GetClaim().GetCurrentStatus()
	// Check for the correct credential status
	newClaimStatus := newCredStatus.GetClaim().GetCurrentStatus()
	statusFound := 0
	for _, acceptablestatus := range docVerify.AcceptedCredStatuses {
		if newClaimStatus == acceptablestatus {
			statusFound = 1
		}
	}
	if statusFound == 0 {
		return nil, sdkerrors.Wrap(types.ErrInvalidCredentialStatus, fmt.Sprintf("expected credential status to be either of Revoke, Suspend or Expired, got %s", newClaimStatus))
	}

	// Check if the DID of the issuer exists
	issuerId := newCredStatus.GetIssuer()
	if !k.HasDid(ctx, issuerId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Issuer`s DID %s doesnt exists", issuerId))
	}

	// Check if issuer's DID is deactivated
	if issuerDidDocument, _ := k.GetDid(&ctx, issuerId); issuerDidDocument.DidDocumentMetadata.Deactivated {
		return nil, sdkerrors.Wrap(types.ErrDidDocDeactivated, fmt.Sprintf("%s is deactivated and cannot used be used to register credential status", issuerDidDocument.DidDocument.Id))
	}

	// Check if the provided isser Id is the one who issued the VC
	if issuerId != oldCredStatus.GetIssuer() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCredentialField,
			fmt.Sprintf("Isser ID %s is not issuer of verifiable credential id %s", issuerId, credId))
	}

	// Check if the new expiration date and issuance date are same as old one.
	newExpirationDate := newCredStatus.GetExpirationDate()
	newExpirationDateParsed, err := time.Parse(time.RFC3339, newExpirationDate)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid expiration date format: %s", newExpirationDate))
	}

	newIssuanceDate := newCredStatus.GetIssuanceDate()
	newIssuanceDateParsed, err := time.Parse(time.RFC3339, newIssuanceDate)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid issuance date format: %s", newIssuanceDate))
	}

	oldExpirationDateParsed, err := time.Parse(time.RFC3339, oldCredStatus.GetExpirationDate())
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid existing expiration date format: %s", oldCredStatus.GetExpirationDate()))
	}

	oldIssuanceDateParsed, err := time.Parse(time.RFC3339, oldCredStatus.GetIssuanceDate())
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid existing issuance date format: %s", oldCredStatus.GetIssuanceDate()))
	}

	if !newIssuanceDateParsed.Equal(oldIssuanceDateParsed) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("issuance date should be same, new issuance date provided : %s", newIssuanceDate))
	}

	if !newExpirationDateParsed.Equal(oldExpirationDateParsed) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("expiration date should be same, new expiration date provided : %s", newExpirationDate))
	}

	// Check if new expiration date isn't less than new issuance date
	if err := docVerify.VerifyCredentialStatusDates(newIssuanceDateParsed, newExpirationDateParsed); err != nil {
		return nil, err
	}

	// Check if updated date iss imilar to created date
	if err := docVerify.VerifyCredentialProofDates(newCredProof, false); err != nil {
		return nil, err
	}

	// Check if the created date lies between issuance and expiration
	currentDate, _ := time.Parse(time.RFC3339, newCredProof.Created)
	if currentDate.After(newExpirationDateParsed) || currentDate.Before(newIssuanceDateParsed) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, "credential update is happening on a date which doesn`t lie between issuance date and expiration date")
	}

	// Old and New Credential Claim Check
	newStatusReason := newCredStatus.GetClaim().GetStatusReason()

	switch oldClaimStatus {
	case "Live":
		if newClaimStatus == oldClaimStatus {
			return nil, sdkerrors.Wrapf(
				types.ErrInvalidCredentialStatus,
				fmt.Sprintf("credential is already %s", newClaimStatus))
		}

		if newClaimStatus == "Revoked" || newClaimStatus == "Suspended" || newClaimStatus == "Expired" {
			if len(newStatusReason) == 0 {
				return nil, sdkerrors.Wrapf(
					types.ErrInvalidCredentialField,
					fmt.Sprintf("claim status reason cannot be empty for claim status %s", newClaimStatus))
			}
		}

	case "Suspended":
		if newClaimStatus == oldClaimStatus {
			return nil, sdkerrors.Wrapf(
				types.ErrInvalidCredentialStatus,
				fmt.Sprintf("credential is already %s", newClaimStatus))
		}

	case "Revoked", "Expired":
		if newClaimStatus == oldClaimStatus {
			return nil, sdkerrors.Wrapf(
				types.ErrInvalidCredentialStatus,
				fmt.Sprintf("credential is already %s", newClaimStatus))
		}

		if newClaimStatus != oldClaimStatus {
			return nil, sdkerrors.Wrapf(
				types.ErrInvalidCredentialStatus,
				fmt.Sprintf("credential cannot be updated from %s to %s", oldClaimStatus, newClaimStatus))
		}

	default:
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidCredentialField,
			fmt.Sprintf("invalid Credential Status present in existing credential %s", oldClaimStatus))
	}

	// Verify the Signature
	didDocument, err := k.GetDid(&ctx, issuerId)
	if err != nil {
		return nil, err
	}

	did := didDocument.GetDidDocument()

	signature := &types.SignInfo{
		VerificationMethodId: newCredProof.GetVerificationMethod(),
		Signature:            newCredProof.GetProofValue(),
	}
	signatures := []*types.SignInfo{signature}
	signers := did.GetSigners()
	signersWithVM, err := k.GetVMForSigners(&ctx, signers)
	if err != nil {
		return nil, err
	}

	// Proof Type Check
	err = sigVerify.DocumentProofTypeCheck(newCredProof.Type, signersWithVM, newCredProof.VerificationMethod)
	if err != nil {
		return nil, err
	}

	// Verify Signature
	err = sigVerify.VerifyDocumentSignature(&ctx, newCredStatus, signersWithVM, signatures)
	if err != nil {
		return nil, err
	}

	cred := types.Credential{
		Claim:          newCredStatus.GetClaim(),
		Issuer:         newCredStatus.GetIssuer(),
		IssuanceDate:   newCredStatus.GetIssuanceDate(),
		ExpirationDate: newCredStatus.GetExpirationDate(),
		CredentialHash: newCredStatus.GetCredentialHash(),
		Proof:          newCredProof,
	}

	return &cred, nil
}
