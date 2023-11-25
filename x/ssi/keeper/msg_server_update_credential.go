package keeper

import (
	"context"
	"fmt"
	"reflect" /* #nosec G702 */
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k msgServer) UpdateCredentialStatus(goCtx context.Context, msg *types.MsgUpdateCredentialStatus) (*types.MsgUpdateCredentialStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	msgNewCredStatus := msg.GetCredentialStatusDocument()
	msgNewCredProof := msg.GetCredentialStatusProof()

	credId := msgNewCredStatus.GetId()

	// Get Credential from store
	oldCredStatusState, err := k.getCredentialStatusFromState(&ctx, credId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCredentialStatusNotFound, err.Error())
	}
	oldCredStatus := oldCredStatusState.CredentialStatusDocument

	// Check if the incoming Credential Status is equal to registered Credential Status
	if reflect.DeepEqual(oldCredStatus, msgNewCredStatus) {
		return nil, sdkerrors.Wrap(types.ErrInvalidCredentialStatus, "incoming Credential Status Document does not have any changes")
	}

	// Check if the DID of the issuer exists
	issuerId := msgNewCredStatus.GetIssuer()
	if !k.hasDidDocument(ctx, issuerId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Issuer`s DID %s doesnt exists", issuerId))
	}

	// Check if issuer's DID is deactivated
	issuerDidDocument, err := k.getDidDocumentState(&ctx, issuerId)
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

	// Check if the credential is already revoked
	if oldCredStatus.Revoked {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCredentialStatus, "credential status %v could not be updated since it is revoked", oldCredStatus.Id)
	}

	// Check if the new issuance date are same as old one.
	newIssuanceDate := msgNewCredStatus.GetIssuanceDate()
	newIssuanceDateParsed, err := time.Parse(time.RFC3339, newIssuanceDate)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid issuance date format: %s", newIssuanceDate))
	}

	oldIssuanceDateParsed, err := time.Parse(time.RFC3339, oldCredStatus.GetIssuanceDate())
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid existing issuance date format: %s", oldCredStatus.GetIssuanceDate()))
	}

	if !newIssuanceDateParsed.Equal(oldIssuanceDateParsed) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("issuance date should be same, new issuance date provided : %s", newIssuanceDate))
	}

	// Validate Merkle Root Hash
	if err := verifyCredentialMerkleRootHash(msgNewCredStatus.GetCredentialMerkleRootHash()); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCredentialMerkleRootHash, err.Error())
	}

	// Check if input Merkle Root Hash is different. The Credential Merkle Root Hash MUST NEVER change.
	if msgNewCredStatus.CredentialMerkleRootHash != oldCredStatus.CredentialMerkleRootHash {
		return nil, sdkerrors.Wrapf(
			types.ErrInvalidCredentialMerkleRootHash,
			"recieved credential merkle root hash '%v' is different from the merkle root hash of registered credential status document '%v'",
			msgNewCredStatus.CredentialMerkleRootHash,
			oldCredStatus.CredentialMerkleRootHash,
		)
	}

	// Check if the created date before issuance date
	currentDate, err := time.Parse(time.RFC3339, msgNewCredProof.Created)
	if err != nil {
		return nil, err
	}
	if currentDate.Before(newIssuanceDateParsed) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, "proof attached has a creation date before issuance date")
	}

	// Validate Document Proof
	if err := msgNewCredProof.Validate(); err != nil {
		return nil, err
	}

	// Verify Signature
	err = k.VerifyDocumentProof(ctx, msgNewCredStatus, msgNewCredProof)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidSignature, err.Error())
	}

	cred := types.CredentialStatusState{
		CredentialStatusDocument: msgNewCredStatus,
		CredentialStatusProof:    msgNewCredProof,
	}

	k.setCredentialStatusInState(ctx, &cred)

	return &types.MsgUpdateCredentialStatusResponse{}, nil
}
