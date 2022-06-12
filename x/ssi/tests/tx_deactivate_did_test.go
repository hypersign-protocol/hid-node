package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func TestDeactivateDID(t *testing.T) {
	t.Log("Running test for Valid Deactivate DID Tx")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx :=  sdk.WrapSDKContext(ctx)
	
	t.Logf("Registering DID with DID Id: %s", ValidDidDocument.GetId())
	msgCreateDID := &types.MsgCreateDID{
		DidDocString: ValidDidDocument,
		Signatures: DidDocumentValidSignInfo,
		Creator: Creator,
	}

	_, err := msgServer.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Error(err)
		t.FailNow()
	} else {
		t.Log("Did Registeration Successful")
	}

	// Querying registered did document to get the version ID
	resolvedDidDocument, errResolve := k.GetDid(&ctx, ValidDidDocument.Id)
	if errResolve != nil {
		t.Error("Error in retrieving registered did document")
		t.Error(errResolve)
		t.FailNow()
	}
	versionId := resolvedDidDocument.GetMetadata().GetVersionId()

	t.Logf("Deactivating DID with Id: %s", ValidDidDocument.GetId()) 
	msgDeactivateDID := &types.MsgDeactivateDID{
		DidId: ValidDidDocument.GetId(),
		Signatures: DidDocumentValidSignInfo,
		VersionId: versionId,
		Creator: Creator,
	}

	_, errDeactivateDID := msgServer.DeactivateDID(goCtx, msgDeactivateDID)
	if errDeactivateDID != nil {
		t.Error("DID Deactivation Failed")
		t.Error(errDeactivateDID)
		t.FailNow()
	} 
	t.Log("DID Deactivation Successful")

	t.Log("Deactivate DID Tx Test Completed")
}