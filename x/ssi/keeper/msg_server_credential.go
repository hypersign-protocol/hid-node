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

var acceptableCredentialStatuses = []string{
	"Live",
	"Revoked",
}

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
	statusFound := 0
	for _, elem := range acceptableCredentialStatuses {
		if elem == credStatus {
			statusFound = 1
		}
	}
	if statusFound != 1 {
		return nil, sdkerrors.Wrap(types.ErrInvalidCredentialStatus, fmt.Sprintf("expected credential status to be either of %v, got %s", acceptableCredentialStatuses, credStatus))
	}

	// Check if the DID of the issuer exists
	issuerId := credMsg.GetIssuer()
	if !k.HasDid(ctx, issuerId) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Issuer`s DID %s doesnt exists", issuerId))
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
		Claim:  credMsg.GetClaim(),
		Issuer: credMsg.GetIssuer(),
		Issued: credMsg.GetIssuanceDate(),
		Proof:  credProof,
	}

	id := k.RegisterCred(ctx, &cred)

	return &types.MsgRegisterCredentialStatusResponse{Id: id}, nil
}
