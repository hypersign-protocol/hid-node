package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"

	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	testssi "github.com/hypersign-protocol/hid-node/x/ssi/tests/ssi"
)

// Note: Ed25519VerificationKey2020 tests are skipped as it is being
// used in other tests.

func TestEcdsaSecp256k1VerificationKey2019(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	alice_kp := testcrypto.GenerateSecp256k1KeyPair()

	t.Log("Register DID Document")

	alice_didDoc := testssi.GenerateDidDoc(alice_kp)
	alice_didDoc.Controller = append(alice_didDoc.Controller, alice_didDoc.Id)

	alice_kp.VerificationMethodId = alice_didDoc.VerificationMethod[0].Id

	didDocTx := testssi.GetRegisterDidDocumentRPC(alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err := msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Update DID Document")
	alice_didDoc.CapabilityDelegation = []string{alice_didDoc.VerificationMethod[0].Id}
	updateDidDocTx := testssi.GetUpdateDidDocumentRPC(k, ctx, alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Register Credential Schema Document")
	credentialSchema := testssi.GenerateSchema(alice_kp, alice_didDoc.Id)
	schemaRPCElements := testssi.GenerateSchemaRPCElements(alice_kp, credentialSchema, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialSchema(goCtx, schemaRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Register Credential Status Document")
	credentialStatus := testssi.GenerateCredentialStatus(alice_kp, alice_didDoc.Id)
	credentialStatusRPCElements := testssi.GenerateRegisterCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialStatus(goCtx, credentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Deactivate DID Document")
	deactivateDidElements := testssi.GetDeactivateDidDocumentRPC(k, ctx, alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.DeactivateDID(goCtx, deactivateDidElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestEcdsaSecp256k1RecoveryMethod2020(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	alice_kp := testcrypto.GenerateSecp256k1RecoveryKeyPair()

	t.Log("Register DID Document")

	alice_didDoc := testssi.GenerateDidDoc(alice_kp)
	alice_didDoc.Controller = append(alice_didDoc.Controller, alice_didDoc.Id)

	alice_kp.VerificationMethodId = alice_didDoc.VerificationMethod[0].Id

	didDocTx := testssi.GetRegisterDidDocumentRPC(alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err := msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Update DID Document")
	alice_didDoc.CapabilityDelegation = []string{alice_didDoc.VerificationMethod[0].Id}
	updateDidDocTx := testssi.GetUpdateDidDocumentRPC(k, ctx, alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Register Credential Schema Document")
	credentialSchema := testssi.GenerateSchema(alice_kp, alice_didDoc.Id)
	schemaRPCElements := testssi.GenerateSchemaRPCElements(alice_kp, credentialSchema, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialSchema(goCtx, schemaRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Register Credential Status Document")
	credentialStatus := testssi.GenerateCredentialStatus(alice_kp, alice_didDoc.Id)
	credentialStatusRPCElements := testssi.GenerateRegisterCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialStatus(goCtx, credentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Deactivate DID Document")
	deactivateDidElements := testssi.GetDeactivateDidDocumentRPC(k, ctx, alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.DeactivateDID(goCtx, deactivateDidElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestBabyJubJubKey2021(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	alice_kp := testcrypto.GenerateBabyJubJubKeyPair()

	t.Log("Register DID Document")

	alice_didDoc := testssi.GenerateDidDoc(alice_kp)
	alice_didDoc.Controller = append(alice_didDoc.Controller, alice_didDoc.Id)

	alice_kp.VerificationMethodId = alice_didDoc.VerificationMethod[0].Id

	didDocTx := testssi.GetRegisterDidDocumentRPC(alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err := msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Update DID Document")
	alice_didDoc.CapabilityDelegation = []string{alice_didDoc.VerificationMethod[0].Id}
	updateDidDocTx := testssi.GetUpdateDidDocumentRPC(k, ctx, alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Register Credential Schema Document")
	credentialSchema := testssi.GenerateSchema(alice_kp, alice_didDoc.Id)
	schemaRPCElements := testssi.GenerateSchemaRPCElements(alice_kp, credentialSchema, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialSchema(goCtx, schemaRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Register Credential Status Document")
	credentialStatus := testssi.GenerateCredentialStatus(alice_kp, alice_didDoc.Id)
	credentialStatusRPCElements := testssi.GenerateRegisterCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialStatus(goCtx, credentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Deactivate DID Document")
	deactivateDidElements := testssi.GetDeactivateDidDocumentRPC(k, ctx, alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.DeactivateDID(goCtx, deactivateDidElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestBbsBls(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	alice_kp := testcrypto.GenerateBbsBlsKeyPair()

	t.Log("Register DID Document")

	alice_didDoc := testssi.GenerateDidDoc(alice_kp)
	alice_didDoc.Controller = append(alice_didDoc.Controller, alice_didDoc.Id)

	alice_kp.VerificationMethodId = alice_didDoc.VerificationMethod[0].Id

	didDocTx := testssi.GetRegisterDidDocumentRPC(alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err := msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Update DID Document")
	alice_didDoc.CapabilityDelegation = []string{alice_didDoc.VerificationMethod[0].Id}
	updateDidDocTx := testssi.GetUpdateDidDocumentRPC(k, ctx, alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Register Credential Schema Document")
	credentialSchema := testssi.GenerateSchema(alice_kp, alice_didDoc.Id)
	schemaRPCElements := testssi.GenerateSchemaRPCElements(alice_kp, credentialSchema, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialSchema(goCtx, schemaRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Register Credential Status Document")
	credentialStatus := testssi.GenerateCredentialStatus(alice_kp, alice_didDoc.Id)
	credentialStatusRPCElements := testssi.GenerateRegisterCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialStatus(goCtx, credentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Deactivate DID Document")
	deactivateDidElements := testssi.GetDeactivateDidDocumentRPC(k, ctx, alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.DeactivateDID(goCtx, deactivateDidElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
