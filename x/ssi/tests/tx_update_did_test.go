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
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

	keyPair1 := GenerateEd25519KeyPair()
	rpcElements := GenerateDidDocumentRPCElements(keyPair1, []DidSigningElements{})
	t.Logf("Registering DID with DID Id: %s", rpcElements.DidDocument.GetId())

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: rpcElements.DidDocument,
		Signatures:   rpcElements.Signatures,
		Creator:      rpcElements.Creator,
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
	resolvedDidDocument, errResolve := k.GetDidDocumentState(&ctx, rpcElements.DidDocument.GetId())
	if errResolve != nil {
		t.Error("Error in retrieving registered did document")
		t.Error(errResolve)
		t.FailNow()
	}
	versionId := resolvedDidDocument.GetDidDocumentMetadata().GetVersionId()

	// Updated the existing DID by appending a link
	resolvedDidDocument.DidDocument.AlsoKnownAs = append(resolvedDidDocument.DidDocument.AlsoKnownAs, "http://www.example.com")

	updateRpcElements := GetModifiedDidDocumentSignature(
		resolvedDidDocument.DidDocument,
		keyPair1,
		resolvedDidDocument.DidDocument.VerificationMethod[0].Id,
	)
	t.Logf("Updating context field of the registered did with Id: %s", updateRpcElements.DidDocument.Id)

	msgUpdateDID := &types.MsgUpdateDID{
		DidDocString: updateRpcElements.DidDocument,
		Signatures:   updateRpcElements.Signatures,
		VersionId:    versionId,
		Creator:      updateRpcElements.Creator,
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
