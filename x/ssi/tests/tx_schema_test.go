package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"

	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	testssi "github.com/hypersign-protocol/hid-node/x/ssi/tests/ssi"
)

func TestSchemaTC1(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("FAIL: Alice creates a DID, deactivates it and attempts to create a schema")

	t.Log("Create Alice's DID")
	alice_kp := testcrypto.GenerateEd25519KeyPair()
	alice_didDoc := testssi.GenerateDidDoc(alice_kp)
	alice_didDoc.Controller = append(alice_didDoc.Controller, alice_didDoc.Id)
	t.Logf("Alice's DID Id: %s", alice_didDoc.Id)
	alice_kp.VerificationMethodId = alice_didDoc.VerificationMethod[0].Id
	didDocTx := testssi.GetRegisterDidDocumentRPC(alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err := msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Alice attempts to deactivate DID Document")
	deactivateDidElements := testssi.GetDeactivateDidDocumentRPC(k, ctx, alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.DeactivateDID(goCtx, deactivateDidElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Alice creates a schema with her deactivated DID Document")
	credentialSchema := testssi.GenerateSchema(alice_kp, alice_didDoc.Id)
	schemaRPCElements := testssi.GenerateSchemaRPCElements(alice_kp, credentialSchema, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialSchema(goCtx, schemaRPCElements)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}
}

func TestSchemaTC2(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("PASS: Bob creates a DID and attempts to register a Schema")

	t.Log("Create Bob's DID")
	bob_kp := testcrypto.GenerateEd25519KeyPair()
	bob_didDoc := testssi.GenerateDidDoc(bob_kp)
	bob_didDoc.Controller = append(bob_didDoc.Controller, bob_didDoc.Id)
	t.Logf("bob's DID Id: %s", bob_didDoc.Id)
	bob_kp.VerificationMethodId = bob_didDoc.VerificationMethod[0].Id
	didDocTx := testssi.GetRegisterDidDocumentRPC(bob_didDoc, []testcrypto.IKeyPair{bob_kp})
	_, err := msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Bob creates a schema")
	credentialSchema := testssi.GenerateSchema(bob_kp, bob_didDoc.Id)
	schemaRPCElements := testssi.GenerateSchemaRPCElements(bob_kp, credentialSchema, bob_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialSchema(goCtx, schemaRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestSchemaTC3(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("FAIL: Bob creates a Schema where the name field is not in Pascal Case")

	t.Log("Create Bob's DID")
	bob_kp := testcrypto.GenerateEd25519KeyPair()
	bob_didDoc := testssi.GenerateDidDoc(bob_kp)
	bob_didDoc.Controller = append(bob_didDoc.Controller, bob_didDoc.Id)
	t.Logf("bob's DID Id: %s", bob_didDoc.Id)
	bob_kp.VerificationMethodId = bob_didDoc.VerificationMethod[0].Id
	didDocTx := testssi.GetRegisterDidDocumentRPC(bob_didDoc, []testcrypto.IKeyPair{bob_kp})
	_, err := msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	credentialSchema := testssi.GenerateSchema(bob_kp, bob_didDoc.Id)
	credentialSchema.Name = "Day Pass"
	schemaRPCElements := testssi.GenerateSchemaRPCElements(bob_kp, credentialSchema, bob_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialSchema(goCtx, schemaRPCElements)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}
}

func TestSchemaTC4(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("FAIL: Bob creates a Schema where the property field is some random string")

	t.Log("Create Bob's DID")
	bob_kp := testcrypto.GenerateEd25519KeyPair()
	bob_didDoc := testssi.GenerateDidDoc(bob_kp)
	bob_didDoc.Controller = append(bob_didDoc.Controller, bob_didDoc.Id)
	t.Logf("bob's DID Id: %s", bob_didDoc.Id)
	bob_kp.VerificationMethodId = bob_didDoc.VerificationMethod[0].Id
	didDocTx := testssi.GetRegisterDidDocumentRPC(bob_didDoc, []testcrypto.IKeyPair{bob_kp})
	_, err := msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	credentialSchema := testssi.GenerateSchema(bob_kp, bob_didDoc.Id)
	credentialSchema.Schema.Properties = "someString"
	schemaRPCElements := testssi.GenerateSchemaRPCElements(bob_kp, credentialSchema, bob_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialSchema(goCtx, schemaRPCElements)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}
}

func TestSchemaTC5(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("FAIL: Bob creates a Schema where the property field is a valid JSON, but one of the attributes has an invalid sub-attribute")

	t.Log("Create Bob's DID")
	bob_kp := testcrypto.GenerateEd25519KeyPair()
	bob_didDoc := testssi.GenerateDidDoc(bob_kp)
	bob_didDoc.Controller = append(bob_didDoc.Controller, bob_didDoc.Id)
	t.Logf("bob's DID Id: %s", bob_didDoc.Id)
	bob_kp.VerificationMethodId = bob_didDoc.VerificationMethod[0].Id
	didDocTx := testssi.GetRegisterDidDocumentRPC(bob_didDoc, []testcrypto.IKeyPair{bob_kp})
	_, err := msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	credentialSchema := testssi.GenerateSchema(bob_kp, bob_didDoc.Id)
	credentialSchema.Schema.Properties = "{\"fullName\":{\"type\":\"string\",\"sda\":\"string\"},\"companyName\":{\"type\":\"string\"},\"center\":{\"type\":\"string\"},\"invoiceNumber\":{\"type\":\"string\"}}"
	schemaRPCElements := testssi.GenerateSchemaRPCElements(bob_kp, credentialSchema, bob_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialSchema(goCtx, schemaRPCElements)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}
}

func TestSchemaTC6(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("FAIL: Bob creates a Schema where the property field is a valid JSON, but `type` sub-attribute is missing")

	t.Log("Create Bob's DID")
	bob_kp := testcrypto.GenerateEd25519KeyPair()
	bob_didDoc := testssi.GenerateDidDoc(bob_kp)
	bob_didDoc.Controller = append(bob_didDoc.Controller, bob_didDoc.Id)
	t.Logf("bob's DID Id: %s", bob_didDoc.Id)
	bob_kp.VerificationMethodId = bob_didDoc.VerificationMethod[0].Id
	didDocTx := testssi.GetRegisterDidDocumentRPC(bob_didDoc, []testcrypto.IKeyPair{bob_kp})
	_, err := msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	credentialSchema := testssi.GenerateSchema(bob_kp, bob_didDoc.Id)
	credentialSchema.Schema.Properties = "{\"fullName\":{\"format\":\"string\"},\"companyName\":{\"type\":\"string\"},\"center\":{\"type\":\"string\"},\"invoiceNumber\":{\"type\":\"string\"}}"
	schemaRPCElements := testssi.GenerateSchemaRPCElements(bob_kp, credentialSchema, bob_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialSchema(goCtx, schemaRPCElements)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}
}