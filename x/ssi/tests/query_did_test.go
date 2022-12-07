package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/vid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/vid-node/x/ssi/types"
	"github.com/stretchr/testify/assert"
)

func TestQueryDidDocument(t *testing.T) {
	t.Log("Running test for QueryDidDocument (Query)")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

	keyPair1 := GenerateEd25519KeyPair()
	rpcElements := GenerateDidDocumentRPCElements(keyPair1)
	didId := rpcElements.DidDocument.GetId()
	t.Logf("Registering DID with DID Id: %s", didId)

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: rpcElements.DidDocument,
		Signatures:   rpcElements.Signatures,
		Creator:      rpcElements.Creator,
	}

	_, errCreateDID := msgServer.CreateDID(goCtx, msgCreateDID)
	if errCreateDID != nil {
		t.Error("DID Registeration Failed")
		t.Log(rpcElements.DidDocument.Id)
		t.Error(errCreateDID)
		t.FailNow()
	}

	t.Log("Did Registeration Successful")
	t.Log("Querying the DID from store")

	req := &types.QueryDidDocumentRequest{
		DidId: didId,
	}

	res, errResponse := k.QueryDidDocument(goCtx, req)
	if errResponse != nil {
		t.Error("Did Resolve Failed")
		t.Error(errResponse)
		t.FailNow()
	}
	t.Log("Querying successful")
	// To check if queried Did Document is not nil
	assert.NotNil(t, res.DidDocument)
	t.Log("Did Resolve Test Completed")
}

func TestQueryDidDocuments(t *testing.T) {
	t.Log("Running test for QueryDocuments (Query)")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

	keyPair1 := GenerateEd25519KeyPair()
	rpcElements := GenerateDidDocumentRPCElements(keyPair1)
	didId := rpcElements.DidDocument.GetId()
	t.Logf("Registering DID with DID Id: %s", didId)

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: rpcElements.DidDocument,
		Signatures:   rpcElements.Signatures,
		Creator:      rpcElements.Creator,
	}

	_, errCreateDID := msgServer.CreateDID(goCtx, msgCreateDID)
	if errCreateDID != nil {
		t.Error("DID Registeration Failed")
		t.Log(rpcElements.DidDocument.Id)
		t.Error(errCreateDID)
		t.FailNow()
	}

	t.Log("Did Registeration Successful")
	t.Log("Querying the list of Did Documents")

	req := &types.QueryDidDocumentsRequest{}

	res, errResponse := k.QueryDidDocuments(goCtx, req)
	if errResponse != nil {
		t.Error("Did Resolve Failed")
		t.Error(errResponse)
		t.FailNow()
	}

	t.Log("Querying successful")

	// Did Document Count should't be zero
	assert.NotEqual(t, "0", res.TotalDidCount)
	// List should be populated with a single Did Document
	assert.Equal(t, 1, len(res.DidDocList))
	// Did Document shouldnt be nil
	assert.NotNil(t, res.DidDocList[0].DidDocument)

	t.Log("Did Param Test Completed")
}
