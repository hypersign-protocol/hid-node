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
	rpcElements := GenerateDidDocumentRPCElements(keyPair1, []DidSigningElements{})
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

func TestCreateSchemaWithMultiControllerDid(t *testing.T) {
	t.Log("Running test for Valid Create Schmea Tx")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

	keyPair1 := GenerateSecp256k1KeyPair()
	rpcElements1 := GenerateDidDocumentRPCElements(keyPair1, []DidSigningElements{})
	t.Logf("Registering Employee 1 with DID Id: %s", rpcElements1.DidDocument.GetId())

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: rpcElements1.DidDocument,
		Signatures:   rpcElements1.Signatures,
		Creator:      rpcElements1.Creator,
	}

	_, err := msgServer.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Error(err)
		t.FailNow()
	}
	t.Log("Employee 1 DID Registered Successfully")

	keyPair2 := GenerateSecp256k1KeyPair()
	rpcElements2 := GenerateDidDocumentRPCElements(keyPair2, []DidSigningElements{})
	t.Logf("Registering Employee 2 with DID Id: %s", rpcElements2.DidDocument.GetId())

	msgCreateDID = &types.MsgCreateDID{
		DidDocString: rpcElements2.DidDocument,
		Signatures:   rpcElements2.Signatures,
		Creator:      rpcElements2.Creator,
	}

	_, err = msgServer.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Error(err)
		t.FailNow()
	}
	t.Log("Employee 2 DID Registered Successfully")

	keyPairOrg := GenerateSecp256k1KeyPair()
	singingElements := []DidSigningElements{
		DidSigningElements{
			keyPair: keyPair1,
			vmId:    rpcElements1.DidDocument.VerificationMethod[0].Id,
		},
		DidSigningElements{
			keyPair: keyPair2,
			vmId:    rpcElements2.DidDocument.VerificationMethod[0].Id,
		},
	}
	rpcElementsOrg := GenerateDidDocumentRPCElements(keyPairOrg, singingElements)
	t.Logf("Registering Org with DID Id: %s", rpcElementsOrg.DidDocument.GetId())

	msgCreateDID = &types.MsgCreateDID{
		DidDocString: rpcElementsOrg.DidDocument,
		Signatures:   rpcElementsOrg.Signatures,
		Creator:      rpcElementsOrg.Creator,
	}

	_, err = msgServer.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Error(err)
		t.FailNow()
	}
	t.Log("Org DID Registered Successfully")

	t.Log("Registering Schema")
	schemaRpcElements := GenerateSchemaDocumentRPCElements(
		keyPair1,
		rpcElementsOrg.DidDocument.Id,
		rpcElements1.DidDocument.AssertionMethod[0],
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
