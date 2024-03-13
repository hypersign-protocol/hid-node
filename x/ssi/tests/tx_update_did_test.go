package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"

	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	testssi "github.com/hypersign-protocol/hid-node/x/ssi/tests/ssi"
)

func TestUpdateDidTC(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("1 FAIL: Alice creates an Org DID where alice is the controller, and Bob's VM is added to its VM List only. Bob attempts to update Org DID by sending his signature.")

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
	org_didDoc.Controller = []string{alice_didDoc.Id}
	org_didDoc.VerificationMethod = []*types.VerificationMethod{bob_didDoc.VerificationMethod[0]}
	didDocTx = testssi.GetRegisterDidDocumentRPC(org_didDoc, []testcrypto.IKeyPair{alice_kp, bob_kp})
	_, err = msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Bob not being controller, attempts to change Organization DID")
	org_didDoc.CapabilityDelegation = []string{org_didDoc.VerificationMethod[0].Id}
	updateDidDocTx := testssi.GetUpdateDidDocumentRPC(k, ctx, org_didDoc, []testcrypto.IKeyPair{bob_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}

	t.Log("2 PASS: Alice being the controller, attempts to update Org DID by sending their signature.")
	updateDidDocTx = testssi.GetUpdateDidDocumentRPC(k, ctx, org_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("3 FAIL: George tries to add himself as a controller in Org DID by only passing his signature")
	george_kp := testcrypto.GenerateEd25519KeyPair()
	george_didDoc := testssi.GenerateDidDoc(george_kp)
	george_didDoc.Controller = append(george_didDoc.Controller, george_didDoc.Id)
	t.Logf("george's DID Id: %s", george_didDoc.Id)
	george_kp.VerificationMethodId = george_didDoc.VerificationMethod[0].Id
	didDocTx = testssi.GetRegisterDidDocumentRPC(george_didDoc, []testcrypto.IKeyPair{george_kp})
	_, err = msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	org_didDoc.Controller = append(org_didDoc.Controller, george_didDoc.Id)
	updateDidDocTx = testssi.GetUpdateDidDocumentRPC(k, ctx, org_didDoc, []testcrypto.IKeyPair{george_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}

	t.Log("4 FAIL: Alice attemps to add George as controller, by only sender her signature")
	updateDidDocTx = testssi.GetUpdateDidDocumentRPC(k, ctx, org_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}

	t.Log("5 PASS: Both Alice and George's signature are passed")
	updateDidDocTx = testssi.GetUpdateDidDocumentRPC(k, ctx, org_didDoc, []testcrypto.IKeyPair{alice_kp, george_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("6 PASS: Alice attempts to remove George as a cotroller by passing only her signature")
	org_didDoc.Controller = []string{alice_didDoc.Id}
	updateDidDocTx = testssi.GetUpdateDidDocumentRPC(k, ctx, org_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("7 PASS: Addition of George as a controller and simultaneous removal of Alice as a controller. Both alice's and george's signature are passed")
	org_didDoc.Controller = []string{george_didDoc.Id}
	updateDidDocTx = testssi.GetUpdateDidDocumentRPC(k, ctx, org_didDoc, []testcrypto.IKeyPair{alice_kp, george_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("8 FAIL: Alice attempts to perform UpdateDID by not changing any property of DID Document")
	updateDidDocTx = testssi.GetUpdateDidDocumentRPC(k, ctx, org_didDoc, []testcrypto.IKeyPair{alice_kp, george_kp})
	_, err = msgServer.UpdateDID(goCtx, updateDidDocTx)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}
}
