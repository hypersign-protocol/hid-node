package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"

	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	testssi "github.com/hypersign-protocol/hid-node/x/ssi/tests/ssi"
)

func TestCredentialStatusTC1(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("FAIL: Alice creates a DID, deactivates it and attempts to create a credential status document")

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

	t.Log("Alice creates a credential status with her deactivated DID Document")
	credentialStatus := testssi.GenerateCredentialStatus(alice_kp, alice_didDoc.Id)
	credentialStatusRPCElements := testssi.GenerateRegisterCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialStatus(goCtx, credentialStatusRPCElements)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}
}

func TestCredentialStatusTC2(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("PASS: Alice creates a DID and attempts to create a credential status document")

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

	t.Log("Alice creates a credential status document")
	credentialStatus := testssi.GenerateCredentialStatus(alice_kp, alice_didDoc.Id)
	credentialStatusRPCElements := testssi.GenerateRegisterCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialStatus(goCtx, credentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestCredentialStatusTC3(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("FAIL: Alice attempts to update credential status without any changes")

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

	t.Log("Alice creates a credential status document")
	credentialStatus := testssi.GenerateCredentialStatus(alice_kp, alice_didDoc.Id)
	credentialStatusRPCElements := testssi.GenerateRegisterCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialStatus(goCtx, credentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Alice attempts to update credential status document without changes")
	updateCredentialStatusRPCElements := testssi.GenerateUpdateCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.UpdateCredentialStatus(goCtx, updateCredentialStatusRPCElements)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}
}

func TestCredentialStatusTC4(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("PASS: Alice suspends the credential status using one of his VMs and then un-suspends it")

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

	t.Log("Alice creates a credential status document")
	credentialStatus := testssi.GenerateCredentialStatus(alice_kp, alice_didDoc.Id)
	credentialStatusRPCElements := testssi.GenerateRegisterCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialStatus(goCtx, credentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Alice attempts to suspend her credential status")
	credentialStatus.Suspended = true
	credentialStatus.Remarks = "Test"
	updateCredentialStatusRPCElements := testssi.GenerateUpdateCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.UpdateCredentialStatus(goCtx, updateCredentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Alice attempts to un-suspend her credential status")
	credentialStatus.Suspended = false
	credentialStatus.Remarks = "Test"
	updateCredentialStatusRPCElements = testssi.GenerateUpdateCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.UpdateCredentialStatus(goCtx, updateCredentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

}

func TestCredentialStatusTC5(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("Alice revokes the credential status using one of his VMs (PASS) and then attempts to un-revoke it (FAIL)")

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

	t.Log("Alice creates a credential status document")
	credentialStatus := testssi.GenerateCredentialStatus(alice_kp, alice_didDoc.Id)
	credentialStatusRPCElements := testssi.GenerateRegisterCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialStatus(goCtx, credentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Alice attempts to revokes her credential status")
	credentialStatus.Revoked = true
	credentialStatus.Remarks = "Test"
	updateCredentialStatusRPCElements := testssi.GenerateUpdateCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.UpdateCredentialStatus(goCtx, updateCredentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Alice attempts to un-revokes her credential status")
	credentialStatus.Revoked = false
	credentialStatus.Remarks = "Test"
	updateCredentialStatusRPCElements = testssi.GenerateUpdateCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.UpdateCredentialStatus(goCtx, updateCredentialStatusRPCElements)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}

}

func TestCredentialStatusTC6(t *testing.T) {
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	t.Log("FAIL: Alice attempts to update her credential status by changing Merkle Root Hash")

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

	t.Log("Alice creates a credential status document")
	credentialStatus := testssi.GenerateCredentialStatus(alice_kp, alice_didDoc.Id)
	credentialStatusRPCElements := testssi.GenerateRegisterCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.RegisterCredentialStatus(goCtx, credentialStatusRPCElements)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("Alice attempts to update her credential status by changing Merkle Root Hash")
	credentialStatus.CredentialMerkleRootHash = "9de17abaffe74f4675c738f5d69c28a329aff8721cb0ed4808d8616e26280ed9"
	updateCredentialStatusRPCElements := testssi.GenerateUpdateCredStatusRPCElements(alice_kp, credentialStatus, alice_didDoc.VerificationMethod[0])
	_, err = msgServer.UpdateCredentialStatus(goCtx, updateCredentialStatusRPCElements)
	if err == nil {
		t.Log(errExpectedToFail)
		t.FailNow()
	}

}


