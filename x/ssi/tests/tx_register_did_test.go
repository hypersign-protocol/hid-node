package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"

	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	testssi "github.com/hypersign-protocol/hid-node/x/ssi/tests/ssi"
)

func TestCreateDidTC1(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	var didDocTx *types.MsgRegisterDID
	var err error
	t.Log("1. Alice has a registered DID Document where Alice is the controller. Bob tries to register their DID Document by keeping both Alice and Bob as controllers.")
	t.Log("Create Alice's DID")
	
	alice_kp := testcrypto.GenerateEd25519KeyPair()
	alice_didDoc := testssi.GenerateDidDoc(alice_kp)
	alice_didDoc.Controller = append(alice_didDoc.Controller, alice_didDoc.Id)

	t.Logf("Alice's DID Id: %s", alice_didDoc.Id)
	alice_kp.VerificationMethodId = alice_didDoc.VerificationMethod[0].Id

	didDocTx = testssi.GetRegisterDidDocumentRPC(alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("1.1 FAIL: Only Bob's signature is sent")
	t.Log("Create Bob's DID")
	bob_kp := testcrypto.GenerateEd25519KeyPair()
	bob_didDoc := testssi.GenerateDidDoc(bob_kp)
	bob_didDoc.Controller = []string{alice_didDoc.Id, bob_didDoc.Id}
	t.Logf("Bob's DID Id: %s", bob_didDoc.Id)
	bob_kp.VerificationMethodId = bob_didDoc.VerificationMethod[0].Id

	didDocTx = testssi.GetRegisterDidDocumentRPC(bob_didDoc, []testcrypto.IKeyPair{bob_kp})
	_, err = msgServer.RegisterDID(goCtx, didDocTx)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}

	t.Log("1.2 PASS: Both Alice's and Bob's signatures are sent")
	didDocTx = testssi.GetRegisterDidDocumentRPC(bob_didDoc, []testcrypto.IKeyPair{alice_kp, bob_kp})
	_, err = msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestCreateDidTC2(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("2. PASS: Alice has a registered DID Document where they are the controller. They attempt to create an organization DID, in which they are the only controller and the verification method field is empty.")

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

	t.Log("Create Organization DID")
	org_kp := testcrypto.GenerateEd25519KeyPair()
	org_didDoc := testssi.GenerateDidDoc(org_kp)
	org_didDoc.Controller = []string{alice_didDoc.Id}
	org_didDoc.VerificationMethod = []*types.VerificationMethod{}

	t.Logf("Organization DID Id: %s", org_didDoc.Id)
	didDocTx = testssi.GetRegisterDidDocumentRPC(org_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err = msgServer.RegisterDID(goCtx, didDocTx)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestCreateDidTC3(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("3. Alice has a registered DID Document where they are the controller. Alice tries to register an Org DID Document where they are the sole controller, and there are two verification Methods, of type EcdsaSecp256k1RecoveryMethod2020, and Alice is the controller for each one of them.")

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

	wallet1_kp := testcrypto.GenerateEd25519KeyPair()
	wallet1_didDoc := testssi.GenerateDidDoc(wallet1_kp)
	wallet1_didDoc.VerificationMethod[0].Controller = alice_didDoc.Id
	t.Logf("Wallet 1 DID Id: %s", wallet1_didDoc.Id)
	wallet1_kp.VerificationMethodId = wallet1_didDoc.VerificationMethod[0].Id
	
	wallet2_kp := testcrypto.GenerateSecp256k1KeyPair()
	wallet2_didDoc := testssi.GenerateDidDoc(wallet2_kp)
	wallet2_didDoc.VerificationMethod[0].Controller = alice_didDoc.Id
	t.Logf("Wallet 2 DID Id: %s", wallet2_didDoc.Id)
	wallet2_kp.VerificationMethodId = wallet2_didDoc.VerificationMethod[0].Id

	t.Log("Create Org DID")
	org_kp := testcrypto.GenerateEd25519KeyPair()
	org_didDoc := testssi.GenerateDidDoc(org_kp)
	org_didDoc.Controller = []string{alice_didDoc.Id}
	org_didDoc.VerificationMethod = []*types.VerificationMethod{
		wallet1_didDoc.VerificationMethod[0],
		wallet2_didDoc.VerificationMethod[0],
	}
	
	t.Log("3.1 FAIL: Signature is provided by only one of the VMs.")
	t.Logf("Org DID Id: %s", org_didDoc.Id)
	didDocTx2 := testssi.GetRegisterDidDocumentRPC(org_didDoc, []testcrypto.IKeyPair{wallet2_kp})
	_, err2 := msgServer.RegisterDID(goCtx, didDocTx2)
	if err2 == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}

	t.Log("3.2 PASS: Signature is provided by both VMs.")
	t.Logf("Org DID Id: %s", org_didDoc.Id)
	didDocTx2 = testssi.GetRegisterDidDocumentRPC(org_didDoc, []testcrypto.IKeyPair{wallet1_kp, wallet2_kp})
	_, err3 := msgServer.RegisterDID(goCtx, didDocTx2)
	if err3 != nil {
		t.Log(err3)
		t.FailNow()
	}
}

func TestCreateDidTC4(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("4. Alice creates an Org DID where Alice is the controller, and she adds a verification method of her friend Eve.")

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

	t.Log("Create Eve's DID")
	eve_kp := testcrypto.GenerateEd25519KeyPair()
	eve_didDoc := testssi.GenerateDidDoc(eve_kp)
	eve_didDoc.Controller = append(eve_didDoc.Controller, eve_didDoc.Id)
	t.Logf("eve's DID Id: %s", eve_didDoc.Id)
	eve_kp.VerificationMethodId = eve_didDoc.VerificationMethod[0].Id
	didDocTx2 := testssi.GetRegisterDidDocumentRPC(eve_didDoc, []testcrypto.IKeyPair{eve_kp})
	_, err2 := msgServer.RegisterDID(goCtx, didDocTx2)
	if err2 != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Create a random DID Document with Alice being the controller, and Eve's VM being the only Verification Method")
	random_kp := testcrypto.GenerateEd25519KeyPair()
	random_didDoc := testssi.GenerateDidDoc(random_kp)
	random_didDoc.Controller = []string{alice_didDoc.Id}
	random_didDoc.VerificationMethod = []*types.VerificationMethod{
		eve_didDoc.VerificationMethod[0],
	}
	t.Logf("Random DID Id: %s", random_didDoc.Id)
	random_kp.VerificationMethodId = random_didDoc.VerificationMethod[0].Id
	
	t.Log("4.1 FAIL: Only Alice sends their singature")
	didDocTx3 := testssi.GetRegisterDidDocumentRPC(random_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err3 := msgServer.RegisterDID(goCtx, didDocTx3)
	if err3 == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}

	t.Log("4.2 PASS: Both Alice and Eve send their singatures")
	didDocTx3 = testssi.GetRegisterDidDocumentRPC(random_didDoc, []testcrypto.IKeyPair{alice_kp, eve_kp})
	_, err4 := msgServer.RegisterDID(goCtx, didDocTx3)
	if err4 != nil {
		t.Log(err4)
		t.FailNow()
	}
}

func TestCreateDidTC5(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("5. FAIL: Alice tries to register a DID Document with duplicate publicKeyMultibase of type Ed25519VerificationKey2020.")

	t.Log("Create Alice's DID")
	alice_kp := testcrypto.GenerateEd25519KeyPair()
	alice_didDoc := testssi.GenerateDidDoc(alice_kp)
	alice_didDoc.Controller = append(alice_didDoc.Controller, alice_didDoc.Id)
	alice_didDoc.VerificationMethod = append(alice_didDoc.VerificationMethod, alice_didDoc.VerificationMethod[0])
	t.Logf("Alice's DID Id: %s", alice_didDoc.Id)
	alice_kp.VerificationMethodId = alice_didDoc.VerificationMethod[0].Id
	didDocTx := testssi.GetRegisterDidDocumentRPC(alice_didDoc, []testcrypto.IKeyPair{alice_kp})
	_, err := msgServer.RegisterDID(goCtx, didDocTx)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}
}
