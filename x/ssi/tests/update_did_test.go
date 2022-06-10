package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/hypersign-protocol/hid-node/testutil/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func TestUpdateDID(t *testing.T) {
	t.Log("Running test for Valid Update DID Tx")
	k, ctx := testkeeper.SsiKeeper(t)
	msgServ := keeper.NewMsgServerImpl(*k)
	goCtx :=  sdk.WrapSDKContext(ctx)
	
	t.Logf("Registering DID with DID Id: %s", ValidDidDocument.GetId())
	msgCreateDID := &types.MsgCreateDID{
		DidDocString: ValidDidDocument,
		Signatures: DidDocumentValidSignInfo,
		Creator: Creator,
	}

	_, err := msgServ.CreateDID(goCtx, msgCreateDID)
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

	t.Logf("Updating context field of the registered did with Id: %s", UpdatedValidDidDocument.GetId()) 
	msgUpdateDID := &types.MsgUpdateDID{
		DidDocString: UpdatedValidDidDocument,
		Signatures: UpdatedDidDocumentValidSignInfo,
		VersionId: versionId,
		Creator: Creator,
	}

	_, errUpdateDID := msgServ.UpdateDID(goCtx, msgUpdateDID)
	if errUpdateDID != nil {
		t.Error("DID Update Failed")
		t.Error(errUpdateDID)
		t.FailNow()
	} 
	t.Log("Did Update Successful")

	t.Log("Update DID Tx Test Completed")
}