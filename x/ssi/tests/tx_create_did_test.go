package tests

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/stretchr/testify/assert"
)

func TestCreateDID(t *testing.T) {
	t.Log("Running test for Valid Create DID Tx")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetDidMethod(&ctx, "hs")
	k.SetDidNamespace(&ctx, "devnet")

	keyPair1 := GeneratePublicPrivateKeyPair()
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

func TestInvalidServiceType(t *testing.T) {
	t.Log("Running test for Invalid Service Type")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	keyPair1 := GeneratePublicPrivateKeyPair()
	rpcElements := GenerateDidDocumentRPCElements(keyPair1)

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

func TestDuplicateServiceId(t *testing.T) {
	t.Log("Running test for Invalid Duplicate Service Id")
	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	keyPair1 := GeneratePublicPrivateKeyPair()
	rpcElements := GenerateDidDocumentRPCElements(keyPair1)

	// Appending a types.Service object with Service Id similar as the first service object
	duplicateServiceId :=  rpcElements.DidDocument.Service[0].Id
	newService := &types.Service{
		Id:              duplicateServiceId,
		Type:            "LinkedDomains",
		ServiceEndpoint: "https://anotherexample.info",
	}

	rpcElements.DidDocument.Service = append(
		rpcElements.DidDocument.Service,
		newService,
	)

	t.Logf("Registering DID with DID Id: %s", rpcElements.DidDocument.GetId())

	updatedDidRpcElements := GetModifiedDidDocumentSignature(
		rpcElements.DidDocument, 
		keyPair1,
		rpcElements.DidDocument.VerificationMethod[0].Id,
	)

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: updatedDidRpcElements.DidDocument,
		Signatures:   updatedDidRpcElements.Signatures,
		Creator:      rpcElements.Creator,
	}


	_, err := msgServer.CreateDID(goCtx, msgCreateDID)
	if err == nil {
		t.Error("DID Document Registeration was expected to fail, as no two or more service endpoints can have similar service Id")
		t.FailNow()
	}
	
	assert.Contains(t, err.Error(), fmt.Sprintf("Service with Id: %s is duplicate: Invalid Service", duplicateServiceId))

	t.Log("Invalid Duplicate Service Id Test Completed")
}
