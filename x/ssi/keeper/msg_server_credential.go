package keeper

import (
	"context"
	"fmt"

	//"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k msgServer) RegisterCredentialStatus(goCtx context.Context, msg *types.MsgRegisterCredentialStatus) (*types.MsgRegisterCredentialStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	credMsg := msg.GetCredentialStatus()
	credProof := msg.GetProof()

	// Check if the credential already exist in the store
	credId := credMsg.GetClaim().GetId()
	if k.HasCredential(ctx, credId) {
		return nil, sdkerrors.Wrap(types.ErrCredentialExists, fmt.Sprintf("Credential ID: %s ", credId))
	}

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

	// Check if the expiration date is not before the issuance date.
	if err := VerifyCredentialStatusDates(credMsg); err != nil {
		return nil, err
	}

	// Check if updated date iss imilar to created date
	if err := VerifyCredentialProofDates(credProof, true); err != nil {
		return nil, err
	}

	// Verify the Signature
	didDocument, err := k.GetDid(&ctx, issuerId)
	if err != nil {
		return nil, err
	}

	did := didDocument.GetDid()
	signature := credProof.GetProofValue()
	verificationMethod := credProof.GetVerificationMethod()

	err = k.VerifyCredentialSignature(credMsg, did, signature, verificationMethod)
	if err != nil {
		return nil, err
	}

	cred := types.Credential{
		Claim:          credMsg.GetClaim(),
		Issuer:         credMsg.GetIssuer(),
		IssuanceDate:   credMsg.GetIssuanceDate(),
		ExpirationDate: credMsg.GetExpirationDate(),
		CredentialHash: credMsg.GetCredentialHash(),
		Proof:          credProof,
	}

	id := k.RegisterCred(ctx, &cred)

	return &types.MsgRegisterCredentialStatusResponse{Id: id}, nil
}
