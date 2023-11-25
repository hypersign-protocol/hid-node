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
	ctx := sdk.UnwrapSDKContext(goCtx)

	msgCredStatus := msg.GetCredentialStatusDocument()
	msgCredProof := msg.GetCredentialStatusProof()

	credId := msgCredStatus.GetId()

	chainNamespace := k.GetChainNamespace(&ctx)

	// Check the format of Credential Status ID
	err := verification.IsValidID(credId, chainNamespace, "credDocument")
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidCredentialStatusID, err.Error())
	}

	// Check if the credential already exist in the store
	if k.hasCredential(ctx, credId) {
		return nil, types.ErrCredentialStatusExists
	}

	// Check if the DID of the issuer exists
	issuerId := msgCredStatus.GetIssuer()
	if !k.hasDidDocument(ctx, issuerId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Issuer`s DID %s doesnt exists", issuerId))
	}

	// Check if issuer's DID is deactivated
	issuerDidDocument, err := k.getDidDocumentState(&ctx, issuerId)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDidDoc, err.Error())
	}
	if issuerDidDocument.DidDocumentMetadata.Deactivated {
		return nil, sdkerrors.Wrap(types.ErrDidDocDeactivated, fmt.Sprintf("%s is deactivated and cannot used be used to register credential status", issuerDidDocument.DidDocument.Id))
	}

	issuanceDate := msgCredStatus.GetIssuanceDate()
	issuanceDateParsed, err := time.Parse(time.RFC3339, issuanceDate)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid issuance date format: %s", issuanceDate))
	}

	// Check if the created date before issuance date
	currentDate, err := time.Parse(time.RFC3339, msgCredProof.Created)
	if err != nil {
		return nil, err
	}
	if currentDate.Before(issuanceDateParsed) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidDate, "proof attached has a creation date before issuance date")
	}

	// Validate Merkle Root Hash
	if err := verifyCredentialMerkleRootHash(msgCredStatus.GetCredentialMerkleRootHash()); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidCredentialMerkleRootHash, err.Error())
	}

	// Validate Document Proof
	if err := msgCredProof.Validate(); err != nil {
		return nil, err
	}

	// Verify Signature
	err = k.VerifyDocumentProof(ctx, msgCredStatus, msgCredProof)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidSignature, err.Error())
	}

	cred := &types.CredentialStatusState{
		CredentialStatusDocument: msgCredStatus,
		CredentialStatusProof:    msgCredProof,
	}

	k.setCredentialStatusInState(ctx, cred)

	// Emit a successful Credential Status Registration event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent("create_credential_status", sdk.NewAttribute("tx_author", msg.GetTxAuthor())),
	)

	return &types.MsgRegisterCredentialStatusResponse{}, nil
}
