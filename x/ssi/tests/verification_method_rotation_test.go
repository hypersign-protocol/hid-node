package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/stretchr/testify/assert"
)


func TestVerificationMethodRotation(t *testing.T) {
	t.Logf("Verification Rotation Test Started")

	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	// Create a DID with pubKey1
	keyPair1 := GeneratePublicPrivateKeyPair()
	didId1, err := CreateDidTx(msgServer, goCtx, keyPair1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("DID registerd with ID: %s", didId1)

	// Update the DID by adding pubKey2 
	keyPair2 := GeneratePublicPrivateKeyPair()

	resolvedDidDocument := QueryDid(k, ctx, didId1)
	versionId := resolvedDidDocument.GetMetadata().GetVersionId()

	newVerificationMethod := &types.VerificationMethod{
		Id:                 didId1 + "#" + keyPair2.publicKey,
		Type:               "Ed25519VerificationKey",
		Controller:         didId1,
		PublicKeyMultibase: keyPair2.publicKey,
	}
	resolvedDidDocument.Did.VerificationMethod = append(
		resolvedDidDocument.Did.VerificationMethod, 
		newVerificationMethod)
	resolvedDidDocument.Did.Authentication = append(
		resolvedDidDocument.Did.Authentication,
		newVerificationMethod.Id)

	updatedDidRpcElements := GetModifiedDidDocumentSignature(
		resolvedDidDocument.Did, 
		keyPair1,
		resolvedDidDocument.Did.VerificationMethod[0].Id,
	)

	UpdateDidTx(msgServer, goCtx, updatedDidRpcElements, versionId)
	
	// Remove the first public key using the second public key
	resolvedDidDocument = QueryDid(k, ctx, didId1)
	versionId = resolvedDidDocument.GetMetadata().GetVersionId()
	
	resolvedDidDocument.Did.VerificationMethod = resolvedDidDocument.Did.VerificationMethod[1:]
	resolvedDidDocument.Did.Authentication = resolvedDidDocument.Did.Authentication[1:]
	
	updatedDidRpcElements = GetModifiedDidDocumentSignature(
		resolvedDidDocument.Did, 
		keyPair2,
		resolvedDidDocument.Did.VerificationMethod[0].Id,
	)

	UpdateDidTx(msgServer, goCtx, updatedDidRpcElements, versionId)
	
	// Assert if the new VM is on the only VM present
	assert.Equal(t, 1, len(resolvedDidDocument.Did.VerificationMethod), "Updated DID Document should have only one Verfiication Method")
	assert.EqualValues(t, newVerificationMethod, resolvedDidDocument.Did.VerificationMethod[0], "Unexpected Verification method")
}
