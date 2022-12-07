package tests

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/vid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/vid-node/x/ssi/types"
	"github.com/stretchr/testify/assert"
)

func TestQueryCredential(t *testing.T) {
	t.Log("Running test for QueryCredential (Query)")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

	keyPair1 := GenerateEd25519KeyPair()
	didRpcElements := GenerateDidDocumentRPCElements(keyPair1)
	didId := didRpcElements.DidDocument.GetId()
	t.Logf("Registering DID with DID Id: %s", didId)

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: didRpcElements.DidDocument,
		Signatures:   didRpcElements.Signatures,
		Creator:      didRpcElements.Creator,
	}

	_, errCreateDID := msgServer.CreateDID(goCtx, msgCreateDID)
	if errCreateDID != nil {
		t.Error("DID Registeration Failed")
		t.Log(didRpcElements.DidDocument.Id)
		t.Error(errCreateDID)
		t.FailNow()
	}

	t.Log("Did Registeration Successful")

	credRpcElements := GenerateCredStatusRPCElements(
		keyPair1,
		didRpcElements.DidDocument.Id,
		didRpcElements.DidDocument.VerificationMethod[0],
	)

	msgRegisterCredentialStatus := &types.MsgRegisterCredentialStatus{
		CredentialStatus: credRpcElements.Status,
		Proof:            credRpcElements.Proof,
		Creator:          Creator,
	}
	credId := credRpcElements.Status.GetClaim().GetId()
	t.Logf("Registering Credential Status with Id: %s", credId)

	t.Logf("Block Time: %s", ctx.BlockTime().Format(time.RFC3339))

	_, errCredStatus := msgServer.RegisterCredentialStatus(goCtx, msgRegisterCredentialStatus)
	if errCredStatus != nil {
		t.Error("Credential Status Registeration Failed")
		t.Error(errCredStatus)
		t.FailNow()
	}
	t.Log("Credential Status Registeration Successful")

	t.Log("Querying the Credential Status from store")

	req := &types.QueryCredentialRequest{
		CredId: credId,
	}

	res, errResponse := k.QueryCredential(goCtx, req)
	if errResponse != nil {
		t.Error("Credential Status Resolve Failed")
		t.Error(errResponse)
		t.FailNow()
	}
	t.Log("Querying successful")

	// To check if queried Schema Document is not nil
	assert.NotNil(t, res.CredStatus)

	t.Log("QueryCredential Test Completed")
}

func TestQueryCredentials(t *testing.T) {
	t.Log("Running test for QueryCredentials (Query)")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

	keyPair1 := GenerateEd25519KeyPair()
	didRpcElements := GenerateDidDocumentRPCElements(keyPair1)
	didId := didRpcElements.DidDocument.GetId()
	t.Logf("Registering DID with DID Id: %s", didId)

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: didRpcElements.DidDocument,
		Signatures:   didRpcElements.Signatures,
		Creator:      didRpcElements.Creator,
	}

	_, errCreateDID := msgServer.CreateDID(goCtx, msgCreateDID)
	if errCreateDID != nil {
		t.Error("DID Registeration Failed")
		t.Error(errCreateDID)
		t.FailNow()
	}

	t.Log("Did Registeration Successful")

	credRpcElements := GenerateCredStatusRPCElements(
		keyPair1,
		didRpcElements.DidDocument.Id,
		didRpcElements.DidDocument.VerificationMethod[0],
	)

	msgRegisterCredentialStatus := &types.MsgRegisterCredentialStatus{
		CredentialStatus: credRpcElements.Status,
		Proof:            credRpcElements.Proof,
		Creator:          Creator,
	}
	credId := credRpcElements.Status.GetClaim().GetId()
	t.Logf("Registering Credential Status with Id: %s", credId)

	t.Logf("Block Time: %s", ctx.BlockTime().Format(time.RFC3339))

	_, errCredStatus := msgServer.RegisterCredentialStatus(goCtx, msgRegisterCredentialStatus)
	if errCredStatus != nil {
		t.Error("Credential Status Registeration Failed")
		t.Error(errCredStatus)
		t.FailNow()
	}
	t.Log("Credential Status Registeration Successful")

	t.Log("Querying list of Credential Statuses from store")

	req := &types.QueryCredentialsRequest{}

	res, errResponse := k.QueryCredentials(goCtx, req)
	if errResponse != nil {
		t.Error("Credential Status Resolve Failed")
		t.Error(errResponse)
		t.FailNow()
	}
	t.Log("Querying successful")

	// Schema Document Count should't be zero
	assert.NotEqual(t, "0", res.TotalCount)
	// List should be populated with a single Schema Document
	assert.Equal(t, 1, len(res.Credentials))
	// Schema Document shouldnt be nil
	assert.NotNil(t, res.Credentials[0])

	t.Log("QueryCredentials Test Completed")
}
