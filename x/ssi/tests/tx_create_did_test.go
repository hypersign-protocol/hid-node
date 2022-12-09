package tests

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateDIDUsingEd25519KeyPair(t *testing.T) {
	t.Log("Running test for Valid Create DID Tx")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

	keyPair1 := GenerateEd25519KeyPair()
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
	t.Log("Did Registeration Successful")

	t.Log("Create DID Tx Test Completed")
}

func TestCreateDIDUsingSecp256k1KeyPair(t *testing.T) {
	t.Log("Create DID with Secp256k1")

	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

	keyPair1 := GenerateSecp256k1KeyPair()
	rpcElements := GenerateDidDocumentRPCElements(keyPair1)

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: rpcElements.DidDocument,
		Signatures:   rpcElements.Signatures,
		Creator:      rpcElements.Creator,
	}
	t.Logf("Signature in base64: %v", msgCreateDID.Signatures[0].Signature)
	_, err := msgServer.CreateDID(goCtx, msgCreateDID)
	if err != nil {
		t.Error("DID Registeration Failed")
		t.Error(err)
		t.FailNow()
	}
	t.Log("Did Registeration Successful")

	t.Log("Create DID Tx Test Completed")
}

func TestInvalidServiceType(t *testing.T) {
	t.Log("Running test for Invalid Service Type")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	keyPair1 := GenerateEd25519KeyPair()
	rpcElements := GenerateDidDocumentRPCElements(keyPair1)
	// Set Namespace
	k.SetChainNamespace(&ctx, "devnet")

	// Changing Service Type from "LinkedDomains" to "DIDComm" which is not supported type
	invalidServiceType := "DIDComm"
	rpcElements.DidDocument.Service[0].Type = invalidServiceType

	updatedRpcElements := GetModifiedDidDocumentSignature(
		rpcElements.DidDocument,
		keyPair1,
		rpcElements.DidDocument.VerificationMethod[0].Id,
	)
	t.Logf("Registering DID with DID Id: %s", rpcElements.DidDocument.GetId())

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: updatedRpcElements.DidDocument,
		Signatures:   updatedRpcElements.Signatures,
		Creator:      rpcElements.Creator,
	}

	_, err := msgServer.CreateDID(goCtx, msgCreateDID)
	if err == nil {
		t.Error("DID Document Registeration was expected to fail, as the service type provided was invalid")
		t.FailNow()
	}

	assert.Contains(t, err.Error(), fmt.Sprintf("Service Type %s is Invalid: Invalid Service", invalidServiceType))

	t.Log("Test Completed")
}
