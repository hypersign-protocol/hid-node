package tests

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/stretchr/testify/assert"
)

func TestVerificationMethodRotation(t *testing.T) {
	t.Logf("Verification Rotation Test Started")

	k, ctx := TestKeeper(t)
	msgServer := keeper.NewMsgServerImpl(*k)
	goCtx := sdk.WrapSDKContext(ctx)

	k.SetChainNamespace(&ctx, "devnet")

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

	// Replace the old public key with new one
	resolvedDidDocument.Did.VerificationMethod[0].PublicKeyMultibase = keyPair2.publicKey

	updatedDidRpcElements := GetModifiedDidDocumentSignature(
		resolvedDidDocument.Did,
		keyPair1,
		resolvedDidDocument.Did.VerificationMethod[0].Id,
	)

	err = UpdateDidTx(msgServer, goCtx, updatedDidRpcElements, versionId)
	if err != nil {
		t.Error("Unable to rotate key")
		t.Error(err)
		t.FailNow()
	}

	resolvedDidDocument = QueryDid(k, ctx, didId1)
	// Assert if the update was successful
	assert.Equal(t, keyPair2.publicKey, resolvedDidDocument.Did.VerificationMethod[0].PublicKeyMultibase)
}
