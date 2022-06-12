package tests

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func TestRegisterCredentialStatus(t *testing.T) {
	t.Log("Running test for Valid Register Cred Tx")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Logf("Registering DID with DID Id: %s", ValidDidDocument.GetId())
	msgCreateDID := &types.MsgCreateDID{
		DidDocString: ValidDidDocument,
		Signatures:   DidDocumentValidSignInfo,
		Creator:      Creator,
	}

	_, err := msgServer.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Error(err)
		t.FailNow()
	}
	t.Log("Did Registeration Successful")

	t.Logf("Registering Credential Status with Id: %s", ValidCredentialStatus.GetClaim().GetId())
	msgRegisterCredentialStatus := &types.MsgRegisterCredentialStatus{
		CredentialStatus: ValidCredentialStatus,
		Proof:            ValidCredentialProof,
		Creator:          Creator,
	}

	t.Logf("Block Time: %s", ctx.BlockTime().Format(time.RFC3339))

	_, errCredStatus := msgServer.RegisterCredentialStatus(goCtx, msgRegisterCredentialStatus)
	if errCredStatus != nil {
		t.Error("Credential Status Registeration Failed")
		t.Error(errCredStatus)
		t.FailNow()
	}
	t.Log("Credential Status Registeration Successful")

	t.Log("Valid Register Cred Tx Test Completed")
}

func TestUpdateCredentialStatus(t *testing.T) {
	t.Log("Running test for updating status of registered credential status")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Logf("Registering DID with DID Id: %s", ValidDidDocument.GetId())
	msgCreateDID := &types.MsgCreateDID{
		DidDocString: ValidDidDocument,
		Signatures:   DidDocumentValidSignInfo,
		Creator:      Creator,
	}

	_, err := msgServer.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Error(err)
		t.FailNow()
	}
	t.Log("Did Registeration Successful")

	t.Logf("Registering Credential Status with Id: %s", ValidCredentialStatus.GetClaim().GetId())
	msgRegisterCredentialStatus := &types.MsgRegisterCredentialStatus{
		CredentialStatus: ValidCredentialStatus,
		Proof:            ValidCredentialProof,
		Creator:          Creator,
	}

	t.Logf("Block Time: %s", ctx.BlockTime().Format(time.RFC3339))

	_, errCredStatus := msgServer.RegisterCredentialStatus(goCtx, msgRegisterCredentialStatus)
	if errCredStatus != nil {
		t.Error("Credential Status Registeration Failed")
		t.Error(errCredStatus)
		t.FailNow()
	}
	t.Log("Credential Status Registeration Successful")

	t.Logf("Updating Credential Status (Id: %s) from Live to Revoked", ValidCredentialStatus.GetClaim().GetId())
	
	updatedCredentialStatus, updatedCredentialProof := updateCredStatus("Revoked", *ValidCredentialStatus, *ValidCredentialProof)
	
	msgUpdateCredentialStatus := &types.MsgRegisterCredentialStatus{
		CredentialStatus: &updatedCredentialStatus,
		Proof:            &updatedCredentialProof,
		Creator:          Creator,
	}

	_, errUpdatedCredStatus := msgServer.RegisterCredentialStatus(goCtx, msgUpdateCredentialStatus)
	if errUpdatedCredStatus != nil {
		t.Error("Credential Status Update Failed")
		t.Error(errUpdatedCredStatus)
		t.FailNow()
	}
	t.Log("Credential Status Update Successful")
}
