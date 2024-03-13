package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"

	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	testssi "github.com/hypersign-protocol/hid-node/x/ssi/tests/ssi"
)

func TestDeactivateDidTC1(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("PASS: Alice creates an Org DID with herself and Bob being the Controller. Bob attempts to deactivate it")

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

	t.Log("Create Bob's DID")
	bob_kp := testcrypto.GenerateEd25519KeyPair()
	bob_didDoc := testssi.GenerateDidDoc(bob_kp)
	bob_didDoc.Controller = append(bob_didDoc.Controller, bob_didDoc.Id)
	t.Logf("Bob's DID Id: %s", bob_didDoc.Id)
	bob_kp.VerificationMethodId = bob_didDoc.VerificationMethod[0].Id
	didDocTx = testssi.GetRegisterDidDocumentRPC(bob_didDoc, []testcrypto.IKeyPair{bob_kp})
	_, err = msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Create Organization DID")
	org_kp := testcrypto.GenerateEd25519KeyPair()
	org_didDoc := testssi.GenerateDidDoc(org_kp)
	org_didDoc.Controller = []string{alice_didDoc.Id, bob_didDoc.Id}
	org_didDoc.VerificationMethod = []*types.VerificationMethod{}
	didDocTx = testssi.GetRegisterDidDocumentRPC(org_didDoc, []testcrypto.IKeyPair{alice_kp, bob_kp})
	_, err = msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Bob attempts to deactivate DID Document")
	rpcElements := testssi.GetDeactivateDidDocumentRPC(k, ctx, org_didDoc, []testcrypto.IKeyPair{bob_kp})
	_, err = msgServer.DeactivateDID(goCtx, rpcElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestDeactivateDidTC2(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("FAIL: Alice creates a DID for herself. She deactivates it and attempts to update it.")

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

	t.Log("Alice attempts to update it")
	alice_didDoc.CapabilityDelegation = append(alice_didDoc.CapabilityDelegation, alice_didDoc.VerificationMethod[0].Id)
	updateDidElements := testssi.GetUpdateDidDocumentRPC(k, ctx, alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidElements)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}
}
