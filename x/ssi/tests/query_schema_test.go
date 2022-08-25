package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/stretchr/testify/assert"
)

func TestGetSchema(t *testing.T) {
	t.Log("Running test for GetSchema (Query)")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetDidNamespace(&ctx, "devnet")

	keyPair1 := GeneratePublicPrivateKeyPair()
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

	t.Log("Registering Schema")
	schemaRpcElements := GenerateSchemaDocumentRPCElements(
		keyPair1, 
		rpcElements.DidDocument.Id,
		rpcElements.DidDocument.AssertionMethod[0],
	)

	msgCreateSchema := &types.MsgCreateSchema{
		SchemaDoc: schemaRpcElements.SchemaDocument,
		SchemaProof: schemaRpcElements.SchemaProof,
		Creator: schemaRpcElements.Creator,
	}

	_, errCreateSchema := msgServer.CreateSchema(goCtx, msgCreateSchema)
	if errCreateSchema != nil {
		t.Error("Schema Registeration Failed")
		t.Error(errCreateSchema)
		t.FailNow()
	} 
	t.Log("Schema Registered Successfully")


	t.Log("Querying the Schema from store")

	req := &types.QueryGetSchemaRequest{
		SchemaId: schemaRpcElements.SchemaDocument.GetId(),
	}

	res, errResponse := k.GetSchema(goCtx, req)
	if errResponse != nil {
		t.Error("Schema Resolve Failed")
		t.Error(errResponse)
		t.FailNow()
	}
	t.Log("Querying successful")

	// To check if queried Schema Document is not nil
	assert.NotNil(t, res.Schema)

	t.Log("GetSchema Test Completed")
}

func TestSchemaParam(t *testing.T) {
	t.Log("Running test for SchemaParam (Query)")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetDidNamespace(&ctx, "devnet")
	
	keyPair1 := GeneratePublicPrivateKeyPair()
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
		t.Error(errCreateDID)
		t.FailNow()
	}
	
	t.Log("Did Registeration Successful")

	t.Log("Registering Schema")
	schemaRpcElements := GenerateSchemaDocumentRPCElements(
		keyPair1, 
		rpcElements.DidDocument.Id,
		rpcElements.DidDocument.AssertionMethod[0],
	)

	msgCreateSchema := &types.MsgCreateSchema{
		SchemaDoc: schemaRpcElements.SchemaDocument,
		SchemaProof: schemaRpcElements.SchemaProof,
		Creator: schemaRpcElements.Creator,
	}

	_, errCreateSchema := msgServer.CreateSchema(goCtx, msgCreateSchema)
	if errCreateSchema != nil {
		t.Error("Schema Registeration Failed")
		t.Error(errCreateSchema)
		t.FailNow()
	} 
	t.Log("Schema Registered Successfully")

	t.Log("Querying the list of Schema Documents")

	req := &types.QuerySchemaParamRequest{}

	res, errResponse := k.SchemaParam(goCtx, req)
	if errResponse != nil {
		t.Error("Schema List Resolve Failed")
		t.Error(errResponse)
		t.FailNow()
	}

	t.Log("Querying successful")

	// Schema Document Count should't be zero 
	assert.NotEqual(t, "0", res.TotalCount)
	// List should be populated with a single Schema Document
	assert.Equal(t, 1, len(res.SchemaList))
	// Schema Document shouldnt be nil
	assert.NotNil(t, res.SchemaList[0])

	t.Log("SchemaParam Test Completed")
}