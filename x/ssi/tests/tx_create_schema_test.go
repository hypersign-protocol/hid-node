package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func TestCreateSchema(t *testing.T) {
	t.Log("Running test for Valid Create Schmea Tx")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

	keyPair1 := GenerateSecp256k1KeyPair()
	rpcElements := GenerateDidDocumentRPCElements(keyPair1)
	t.Logf("Registering DID with DID Id: %s", rpcElements.DidDocument.GetId())

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: rpcElements.DidDocument,
		Signatures:   rpcElements.Signatures,
		Creator:      rpcElements.Creator,
	}

	_, err := msgServer.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Error(err)
		t.FailNow()
	}
	t.Log("DID Registered Successfully")

	t.Log("Registering Schema")
	schemaRpcElements := GenerateSchemaDocumentRPCElements(
		keyPair1,
		rpcElements.DidDocument.Id,
		rpcElements.DidDocument.AssertionMethod[0],
	)

	msgCreateSchema := &types.MsgCreateSchema{
		SchemaDoc:   schemaRpcElements.SchemaDocument,
		SchemaProof: schemaRpcElements.SchemaProof,
		Creator:     schemaRpcElements.Creator,
	}

	_, errCreateSchema := msgServer.CreateSchema(goCtx, msgCreateSchema)
	if errCreateSchema != nil {
		t.Error("Schema Registeration Failed")
		t.Error(errCreateSchema)
		t.FailNow()
	}
	t.Log("Schema Registered Successfully")

	t.Log("Create Schema Tx Test Completed")
}
