package ssi

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	
	testconstants "github.com/hypersign-protocol/hid-node/x/ssi/tests/constants"
	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func GetRegisterDidDocumentRPC(
	didDocument *types.DidDocument,
	keyPairs []testcrypto.IKeyPair,
) *types.MsgRegisterDID {
	var proofs []*types.DocumentProof = getDocumentProof(
		didDocument,
		keyPairs,
	)

	return &types.MsgRegisterDID{
		DidDocument:       didDocument,
		DidDocumentProofs: proofs,
		TxAuthor:          testconstants.Creator,
	}
}

func GetUpdateDidDocumentRPC(
	k *keeper.Keeper,
	ctx sdk.Context,
	didDocument *types.DidDocument,
	keyPairs []testcrypto.IKeyPair,
) *types.MsgUpdateDID {
	// Get Version ID
	didDocFromState := QueryDid(k, ctx, didDocument.Id)
	versionId := didDocFromState.DidDocumentMetadata.VersionId

	var proofs []*types.DocumentProof = getDocumentProof(
		didDocument,
		keyPairs,
	)

	return &types.MsgUpdateDID{
		DidDocument:       didDocument,
		DidDocumentProofs: proofs,
		TxAuthor:          testconstants.Creator,
		VersionId:         versionId,
	}
}

func GetDeactivateDidDocumentRPC(
	k *keeper.Keeper,
	ctx sdk.Context,
	didDocument *types.DidDocument,
	keyPairs []testcrypto.IKeyPair,
) *types.MsgDeactivateDID {
	// Get Version ID
	didId := didDocument.Id
	didDocFromState := QueryDid(k, ctx, didId)
	versionId := didDocFromState.DidDocumentMetadata.VersionId

	var proofs []*types.DocumentProof = getDocumentProof(
		didDocument,
		keyPairs,
	)

	return &types.MsgDeactivateDID{
		DidDocumentId:     didId,
		DidDocumentProofs: proofs,
		TxAuthor:          testconstants.Creator,
		VersionId:         versionId,
	}
}

func GenerateDidDoc(keyPair testcrypto.IKeyPair) *types.DidDocument {
	publicKey, optionalID := testcrypto.GetPublicKeyAndOptionalID(keyPair)
	var didId string
	if optionalID == "" {
		didId = "did:" + testconstants.DidMethod + ":" + testconstants.ChainNamespace + ":" + publicKey
	} else {
		didId = "did:" + testconstants.DidMethod + ":" + testconstants.ChainNamespace + ":" + optionalID
	}

	var verificationMethodId string = didId + "#" + "key-1"

	var vmType string = keyPair.GetType() 

	var vm = &types.VerificationMethod{
		Id:                 verificationMethodId,
		Type:               vmType,
		Controller:         didId,
		PublicKeyMultibase: publicKey,
	}

	if optionalID != "" {
		if vm.Type == types.EcdsaSecp256k1RecoveryMethod2020 {
			vm.PublicKeyMultibase = ""
			vm.BlockchainAccountId = "eip155:1:" + optionalID
		}
		if vm.Type == types.EcdsaSecp256k1VerificationKey2019 {
			vm.BlockchainAccountId = "cosmos:prajna:" + optionalID
		}
	}
 
	vmContextUrl := GetContextFromKeyPair(keyPair)
	var didDocument *types.DidDocument = &types.DidDocument{
		Context: []string{
			"https://www.w3.org/ns/did/v1",
			vmContextUrl,
		},
		Id:         didId,
		Controller: []string{},
		VerificationMethod: []*types.VerificationMethod{
			vm,
		},
	}

	return didDocument
}

func QueryDid(k *keeper.Keeper, ctx sdk.Context, Id string) *types.DidDocumentState {
	resolvedDidDocument, errResolve := k.DidDocumentByID(&ctx, &types.QueryDidDocumentRequest{
		DidId: Id,
	})
	if errResolve != nil {
		panic(errResolve)
	}

	return &types.DidDocumentState{
		DidDocument:         resolvedDidDocument.DidDocument,
		DidDocumentMetadata: resolvedDidDocument.DidDocumentMetadata,
	}
}
