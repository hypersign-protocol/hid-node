package tests

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func TestRegisterCredentialStatus(t *testing.T) {
	t.Log("Running test for Valid Register Cred Tx")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

	keyPair1 := GenerateSecp256k1KeyPair()
	didRpcElements := GenerateDidDocumentRPCElements(keyPair1, []DidSigningElements{})
	t.Logf("Registering DID with DID Id: %s", didRpcElements.DidDocument.GetId())

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: didRpcElements.DidDocument,
		Signatures:   didRpcElements.Signatures,
		Creator:      didRpcElements.Creator,
	}

	_, err := msgServer.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Log(didRpcElements.DidDocument.Id)
		t.Error(err)
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
	t.Logf("Registering Credential Status with Id: %s", credRpcElements.Status.GetClaim().GetId())

	t.Logf("Block Time: %s", ctx.BlockTime().Format(time.RFC3339))

	_, errCredStatus := msgServer.RegisterCredentialStatus(goCtx, msgRegisterCredentialStatus)
	if errCredStatus != nil {
		t.Error("Credential Status Registeration Failed")
		t.Error(errCredStatus)
		t.FailNow()
	}
	t.Log("Credential Status Registeration Successful")

	t.Log("Valid Register Cred Tx Test Completed")
}

func TestCreateCredentialStatusWithMultiControllerDid(t *testing.T) {
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
			vmId: rpcElements1.DidDocument.VerificationMethod[0].Id,
		},
		DidSigningElements{
			keyPair: keyPair2,
			vmId: rpcElements2.DidDocument.VerificationMethod[0].Id,
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
	credRpcElements := GenerateCredStatusRPCElements(
		keyPair2,
		rpcElementsOrg.DidDocument.Id,
		rpcElements2.DidDocument.VerificationMethod[0],
	)

	msgRegisterCredentialStatus := &types.MsgRegisterCredentialStatus{
		CredentialStatus: credRpcElements.Status,
		Proof:            credRpcElements.Proof,
		Creator:          Creator,
	}
	t.Logf("Registering Credential Status with Id: %s", credRpcElements.Status.GetClaim().GetId())

	t.Logf("Block Time: %s", ctx.BlockTime().Format(time.RFC3339))

	_, errCredStatus := msgServer.RegisterCredentialStatus(goCtx, msgRegisterCredentialStatus)
	if errCredStatus != nil {
		t.Error("Credential Status Registeration Failed")
		t.Error(errCredStatus)
		t.FailNow()
	}
	t.Log("Credential Status Registeration Successful")

	t.Log("Valid Register Cred Tx Test Completed")
}

func TestUpdateCredentialStatus(t *testing.T) {
	t.Log("Running test for updating status of registered credential status")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

	keyPair1 := GenerateEd25519KeyPair()
	didRpcElements := GenerateDidDocumentRPCElements(keyPair1, []DidSigningElements{})
	t.Logf("Registering DID with DID Id: %s", didRpcElements.DidDocument.GetId())

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: didRpcElements.DidDocument,
		Signatures:   didRpcElements.Signatures,
		Creator:      didRpcElements.Creator,
	}

	_, err := msgServer.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Log(didRpcElements.DidDocument.Id)
		t.Error(err)
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
	t.Logf("Registering Credential Status with Id: %s", credRpcElements.Status.GetClaim().GetId())

	t.Logf("Block Time: %s", ctx.BlockTime().Format(time.RFC3339))

	_, errCredStatus := msgServer.RegisterCredentialStatus(goCtx, msgRegisterCredentialStatus)
	if errCredStatus != nil {
		t.Error("Credential Status Registeration Failed")
		t.Error(errCredStatus)
		t.FailNow()
	}
	t.Log("Credential Status Registeration Successful")

	t.Logf("Updating Credential Status (Id: %s) from Live to Revoked", credRpcElements.Status.GetClaim().GetId())

	updatedRpcCredElements := UpdateCredStatus("Revoked", credRpcElements, keyPair1)

	msgUpdateCredentialStatus := &types.MsgRegisterCredentialStatus{
		CredentialStatus: updatedRpcCredElements.Status,
		Proof:            updatedRpcCredElements.Proof,
		Creator:          Creator,
	}

	_, errUpdatedCredStatus := msgServer.RegisterCredentialStatus(goCtx, msgUpdateCredentialStatus)
	if errUpdatedCredStatus != nil {
		t.Error("Credential Status Update Failed")
		t.Error(errUpdatedCredStatus)
		t.FailNow()
	}
	t.Log("Credential Status Update Successful")
}
