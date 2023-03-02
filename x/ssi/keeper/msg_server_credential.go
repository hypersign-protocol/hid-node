package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/verification"
)

func (k msgServer) RegisterCredentialStatus(goCtx context.Context, msg *types.MsgRegisterCredentialStatus) (*types.MsgRegisterCredentialStatusResponse, error) {
	var id uint64

	ctx := sdk.UnwrapSDKContext(goCtx)

	msgCredStatus := msg.GetCredentialStatus()
	msgCredProof := msg.GetProof()

	credId := msgCredStatus.GetClaim().GetId()

	chainNamespace := k.GetChainNamespace(&ctx)

	// Check the format of Credential ID
	err := verification.IsValidID(credId, chainNamespace, "credDocument")
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidSchemaID, err.Error())
	}

	// Check if the credential already exist in the store
	if !k.HasCredential(ctx, credId) {
		// Check for the correct credential status
		credStatus := msgCredStatus.GetClaim().GetCurrentStatus()
		if credStatus != "Live" {
			return nil, sdkerrors.Wrap(types.ErrInvalidCredentialStatus, fmt.Sprintf("expected credential status to be `Live`, got %s", credStatus))
		}

		// Check if the DID of the issuer exists
		issuerId := msgCredStatus.GetIssuer()
		if !k.HasDid(ctx, issuerId) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Issuer`s DID %s doesnt exists", issuerId))
		}

		// Check if issuer's DID is deactivated
		issuerDidDocument, err := k.GetDidDocumentState(&ctx, issuerId)
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
		}
		if issuerDidDocument.DidDocumentMetadata.Deactivated {
			return nil, sdkerrors.Wrap(types.ErrDidDocDeactivated, fmt.Sprintf("%s is deactivated and cannot used be used to register credential status", issuerDidDocument.DidDocument.Id))
		}

		// Check if the expiration date is not before the issuance date.
		expirationDate := msgCredStatus.GetExpirationDate()
		expirationDateParsed, err := time.Parse(time.RFC3339, expirationDate)
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid expiration date format: %s", expirationDate))
		}

		issuanceDate := msgCredStatus.GetIssuanceDate()
		issuanceDateParsed, err := time.Parse(time.RFC3339, issuanceDate)
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid issuance date format: %s", issuanceDate))
		}

		if err := verification.VerifyCredentialStatusDates(issuanceDateParsed, expirationDateParsed); err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidCredentialField, err.Error())
		}

		// Check if updated date is similar to created date
		if err := verification.VerifyCredentialProofDates(msgCredProof, true); err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidCredentialField, err.Error())
		}

		// Check if the created date lies between issuance and expiration
		currentDate, _ := time.Parse(time.RFC3339, msgCredProof.Created)
		if currentDate.After(expirationDateParsed) || currentDate.Before(issuanceDateParsed) {
			return nil, sdkerrors.Wrapf(types.ErrInvalidDate, "credential registeration is happening on a date which doesn`t lie between issuance date and expiration date")
		}

		// Check the hash type of credentialHash
		isValidCredentialHash := verification.VerifyCredentialHash(msgCredStatus.GetCredentialHash())
		if !isValidCredentialHash {
			return nil, sdkerrors.Wrapf(types.ErrInvalidCredentialHash, "supported hashing algorithms: sha256")
		}

		// ClientSpec check
		clientSpecOpts := types.ClientSpecOpts{
			ClientSpecType: msg.ClientSpec,
			SSIDoc:         msgCredStatus,
			SignerAddress:  msg.Creator,
		}

		credDocBytes, err := getClientSpecDocBytes(clientSpecOpts)
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidClientSpecType, err.Error())
		}

		// Verify Signature
		err = k.VerifyDocumentProof(ctx, credDocBytes, msgCredProof)
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrInvalidSignature, err.Error())
		}

		cred := &types.Credential{
			Claim:          msgCredStatus.GetClaim(),
			Issuer:         msgCredStatus.GetIssuer(),
			IssuanceDate:   msgCredStatus.GetIssuanceDate(),
			ExpirationDate: msgCredStatus.GetExpirationDate(),
			CredentialHash: msgCredStatus.GetCredentialHash(),
			Proof:          msgCredProof,
		}

		id = k.RegisterCredentialStatusInState(ctx, cred)

	} else {
		cred, err := k.updateCredentialStatus(ctx, msg)
		if err != nil {
			return nil, err
		}
		id = k.RegisterCredentialStatusInState(ctx, cred)
	}

	return &types.MsgRegisterCredentialStatusResponse{Id: id}, nil
}

func (k msgServer) updateCredentialStatus(ctx sdk.Context, msg *types.MsgRegisterCredentialStatus) (*types.Credential, error) {
	msgNewCredStatus := msg.CredentialStatus
	msgNewCredProof := msg.Proof

	credId := msgNewCredStatus.GetClaim().GetId()

	// Get Credential from store
	oldCredStatus, err := k.GetCredentialStatusFromState(&ctx, credId)
	if err != nil {
		return nil, err
	}

	// Check if the DID of the issuer exists
	issuerId := msgNewCredStatus.GetIssuer()
	if !k.HasDid(ctx, issuerId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Issuer`s DID %s doesnt exists", issuerId))
	}

	// Check if issuer's DID is deactivated
	issuerDidDocument, err := k.GetDidDocumentState(&ctx, issuerId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrDidDocNotFound, err.Error())
	}
	if issuerDidDocument.DidDocumentMetadata.Deactivated {
		return nil, sdkerrors.Wrap(types.ErrDidDocDeactivated, fmt.Sprintf("%s is deactivated and cannot used be used to register credential status", issuerDidDocument.DidDocument.Id))
	}

	// Check if the provided isser Id is the one who issued the VC
	if issuerId != oldCredStatus.GetIssuer() {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCredentialField,
			fmt.Sprintf("issuer id %s is not issuer of verifiable credential id %s", issuerId, credId))
	}

	// Check if the new expiration date and issuance date are same as old one.
	newExpirationDate := msgNewCredStatus.GetExpirationDate()
	newExpirationDateParsed, err := time.Parse(time.RFC3339, newExpirationDate)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid expiration date format: %s", newExpirationDate))
	}

	newIssuanceDate := msgNewCredStatus.GetIssuanceDate()
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
	if err := verification.VerifyCredentialStatusDates(newIssuanceDateParsed, newExpirationDateParsed); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidCredentialField, err.Error())
	}

	// Check if updated date iss imilar to created date
	if err := verification.VerifyCredentialProofDates(msgNewCredProof, false); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidCredentialField, err.Error())
	}

	// Check if the created date lies between issuance and expiration
	currentDate, _ := time.Parse(time.RFC3339, msgNewCredProof.Created)
	if currentDate.After(newExpirationDateParsed) || currentDate.Before(newIssuanceDateParsed) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, "credential update is happening on a date which doesn`t lie between issuance date and expiration date")
	}

	// Check for the correct credential status
	oldClaimStatus := oldCredStatus.GetClaim().GetCurrentStatus()
	newClaimStatus := msgNewCredStatus.GetClaim().GetCurrentStatus()
	statusFound := 0
	for _, acceptablestatus := range verification.GetAcceptedCredentialClaimStatuses() {
		if newClaimStatus == acceptablestatus {
			statusFound = 1
		}
	}
	if statusFound == 0 {
		return nil, sdkerrors.Wrap(
			types.ErrInvalidCredentialStatus,
			fmt.Sprintf("unsupported credential claim status %s", newClaimStatus))
	}

	// Old and New Credential Claim Check

	// Reject manual status change to Expired
	if newClaimStatus == verification.ClaimStatus_expired {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidCredentialStatus,
			"claim status cannot be manually changed to Expired")
	}

	switch oldClaimStatus {
	case verification.ClaimStatus_live:
		switch newClaimStatus {
		case verification.ClaimStatus_live:
			return nil, sdkerrors.Wrapf(
				types.ErrInvalidCredentialStatus,
				fmt.Sprintf("credential claim status is already %s", newClaimStatus))
		case verification.ClaimStatus_suspended, verification.ClaimStatus_revoked:
			newStatusReason := msgNewCredStatus.GetClaim().GetStatusReason()
			if len(newStatusReason) == 0 {
				return nil, sdkerrors.Wrapf(
					types.ErrInvalidCredentialStatus,
					"claim status reason cannot be empty",
				)
			}
		}

	case verification.ClaimStatus_suspended:
		switch newClaimStatus {
		case verification.ClaimStatus_live, verification.ClaimStatus_revoked:
			newStatusReason := msgNewCredStatus.GetClaim().GetStatusReason()
			if len(newStatusReason) == 0 {
				return nil, sdkerrors.Wrapf(
					types.ErrInvalidCredentialStatus,
					"claim status reason cannot be empty",
				)
			}
		case verification.ClaimStatus_suspended:
			return nil, sdkerrors.Wrapf(
				types.ErrInvalidCredentialStatus,
				fmt.Sprintf("credential claim status is already %s", newClaimStatus))
		}

	case verification.ClaimStatus_revoked, verification.ClaimStatus_expired:
		if newClaimStatus == oldClaimStatus {
			return nil, sdkerrors.Wrapf(
				types.ErrInvalidCredentialStatus,
				fmt.Sprintf("credential claim status is already %s", newClaimStatus))
		}

		if newClaimStatus != oldClaimStatus {
			return nil, sdkerrors.Wrapf(
				types.ErrInvalidCredentialStatus,
				fmt.Sprintf("credential claim status cannot be updated from %s to %s", oldClaimStatus, newClaimStatus))
		}
	}

	// ClientSpec check
	clientSpecOpts := types.ClientSpecOpts{
		ClientSpecType: msg.ClientSpec,
		SSIDoc:         msgNewCredStatus,
		SignerAddress:  msg.Creator,
	}

	credDocBytes, err := getClientSpecDocBytes(clientSpecOpts)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidClientSpecType, err.Error())
	}

	// Verify Signature
	err = k.VerifyDocumentProof(ctx, credDocBytes, msgNewCredProof)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidSignature, err.Error())
	}

	cred := types.Credential{
		Claim:          msgNewCredStatus.GetClaim(),
		Issuer:         msgNewCredStatus.GetIssuer(),
		IssuanceDate:   msgNewCredStatus.GetIssuanceDate(),
		ExpirationDate: msgNewCredStatus.GetExpirationDate(),
		CredentialHash: msgNewCredStatus.GetCredentialHash(),
		Proof:          msgNewCredProof,
	}

	return &cred, nil
}
