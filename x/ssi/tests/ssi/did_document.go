package ssi

import (
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	testconstants "github.com/hypersign-protocol/hid-node/x/ssi/tests/constants"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetModifiedDidDocumentSignature(
	modifiedDidDocument *types.DidDocument,
	keyPair *testcrypto.KeyPair,
	verificationMethodId string,
) *types.MsgRegisterDID {
	var signingElements []*SsiDocSigningElements

	signingElements = append(signingElements,
		&SsiDocSigningElements{
			KeyPair: keyPair,
			VmId:    verificationMethodId,
		})

	var proofs []*types.DocumentProof = getDocumentProof(
		modifiedDidDocument,
		signingElements,
	)

	return &types.MsgRegisterDID{
		DidDocument:       modifiedDidDocument,
		DidDocumentProofs: proofs,
		TxAuthor:          testconstants.Creator,
	}
}

func GenerateDidDocumentRPCElements(keyPair *testcrypto.KeyPair, signingElements []*SsiDocSigningElements) *types.MsgRegisterDID {
	publicKey, optionalID := testcrypto.GetPublicKeyAndOptionalID(keyPair)
	var didId string
	if optionalID == "" {
		didId = "did:" + testconstants.DidMethod + ":" + testconstants.ChainNamespace + ":" + publicKey
	} else {
		didId = "did:" + testconstants.DidMethod + ":" + testconstants.ChainNamespace + ":" + optionalID
	}

	var verificationMethodId string = didId + "#" + "key-1"

	var vmType string
	switch keyPair.Type {
	case testcrypto.Ed25519KeyPair:
		vmType = types.Ed25519VerificationKey2020
	case testcrypto.Secp256k1Pair:
		vmType = types.EcdsaSecp256k1VerificationKey2019
	}

	var vm = &types.VerificationMethod{
		Id:                 verificationMethodId,
		Type:               vmType,
		Controller:         didId,
		PublicKeyMultibase: publicKey,
	}

	var service *types.Service = &types.Service{
		Id:              didId + "#" + "linkedDomains",
		Type:            "LinkedDomains",
		ServiceEndpoint: "http://www.example.com",
	}

	var controllers []string
	if len(signingElements) > 0 {
		for i := 0; i < len(signingElements); i++ {
			controllers = append(
				controllers,
				stripDidFromVerificationMethod(signingElements[i].VmId))
		}
	} else {
		signingElements = []*SsiDocSigningElements{
			{
				KeyPair: keyPair,
				VmId:    vm.Id,
			},
		}
		controllers = []string{didId}
	}

	var didDocument *types.DidDocument = &types.DidDocument{
		Context: []string{
			"https://www.w3.org/ns/did/v1",
		},
		Id:         didId,
		Controller: controllers,
		VerificationMethod: []*types.VerificationMethod{
			vm,
		},
		Service: []*types.Service{
			service,
		},
		Authentication:  []string{verificationMethodId},
		AssertionMethod: []string{verificationMethodId},
	}

	var signInfo []*types.DocumentProof = getDocumentProof(didDocument, signingElements)

	return &types.MsgRegisterDID{
		DidDocument:       didDocument,
		DidDocumentProofs: signInfo,
		TxAuthor:          testconstants.Creator,
	}
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