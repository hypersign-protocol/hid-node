package tests

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
	secp256k1 "github.com/tendermint/tendermint/crypto/secp256k1"
)

func getDidSigningInfo(didDoc *types.Did, signingElements []DidSigningElements) []*types.SignInfo {
	var signingInfos []*types.SignInfo

	for i := 0; i < len(signingElements); i++ {
		signature := SignGeneric(signingElements[i].keyPair, didDoc.GetSignBytes())
		signingInfos = append(signingInfos, &types.SignInfo{
			VerificationMethodId: signingElements[i].vmId,
			Signature:            base64.StdEncoding.EncodeToString(signature),
		})
	}
	return signingInfos
}

func UpdateCredStatus(newStatus string, credRpcElem CredRpcElements, keyPair ed25519KeyPair) CredRpcElements {
	credRpcElem.Status.Claim.CurrentStatus = newStatus
	credRpcElem.Status.Claim.StatusReason = "Status changed for Testing"
	credRpcElem.Proof.Updated = "2022-05-12T00:00:00Z"

	updatedSignature := base64.StdEncoding.EncodeToString(
		ed25519.Sign(keyPair.privateKey, credRpcElem.Status.GetSignBytes()),
	)

	credRpcElem.Proof.ProofValue = updatedSignature

	return credRpcElem
}

func GetModifiedDidDocumentSignature(modifiedDidDocument *types.Did, keyPair ed25519KeyPair, verificationMethodId string) DidRpcElements {
	var signingElements []DidSigningElements

	signingElements = append(signingElements,
		DidSigningElements{
			keyPair: keyPair,
			vmId:    verificationMethodId,
		})

	var signatures []*types.SignInfo = getDidSigningInfo(
		modifiedDidDocument,
		signingElements,
	)

	return DidRpcElements{
		DidDocument: modifiedDidDocument,
		Signatures:  signatures,
		Creator:     Creator,
	}
}

func GenerateDidDocumentRPCElements(keyPair GenericKeyPair) DidRpcElements {
	var publicKey string = GetPublicKeyGeneric(keyPair)
	var didId = "did:" + DidMethod + ":" + ChainNamespace + ":" + publicKey

	var verificationMethodId string = didId + "#" + "key-1"

	var vmType string
	switch keyPair.(type) {
	case ed25519KeyPair:
		vmType = keeper.Ed25519VerificationKey2020
	case secp256k1KeyPair:
		vmType = keeper.EcdsaSecp256k1VerificationKey2019
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

	var didDocument *types.Did = &types.Did{
		Context: []string{
			"https://www.w3.org/ns/did/v1",
		},
		Id:         didId,
		Controller: []string{didId},
		VerificationMethod: []*types.VerificationMethod{
			vm,
		},
		Service: []*types.Service{
			service,
		},
		Authentication:  []string{verificationMethodId},
		AssertionMethod: []string{verificationMethodId},
	}

	signingElements := []DidSigningElements{
		DidSigningElements{
			keyPair: keyPair,
			vmId:    vm.Id,
		},
	}

	var signInfo []*types.SignInfo = getDidSigningInfo(didDocument, signingElements)

	return DidRpcElements{
		DidDocument: didDocument,
		Signatures:  signInfo,
		Creator:     Creator,
	}
}

func GenerateSchemaDocumentRPCElements(keyPair GenericKeyPair, Id string, verficationMethodId string) SchemaRpcElements {
	var schemaId string = "sch:" + DidMethod + ":" + "devnet" + ":" + strings.Split(Id, ":")[3] + ":1.0"
	var schemaDocument *types.SchemaDocument = &types.SchemaDocument{
		Type:         "https://w3c-ccg.github.io/vc-json-schemas/schema/1.0/schema.json",
		ModelVersion: "v1.0",
		Name:         "HS Credential",
		Author:       Id,
		Id:           schemaId,
		Authored:     "2022-04-10T04:07:12Z",
		Schema: &types.SchemaProperty{
			Schema:               "https://json-schema.org/draft-07/schema#",
			Description:          "test",
			Type:                 "Object",
			Properties:           "{myString:{type:string}}",
			Required:             []string{"myString"},
			AdditionalProperties: false,
		},
	}

	var schemaDocumentSignature string = base64.StdEncoding.EncodeToString(
		SignGeneric(keyPair, schemaDocument.GetSignBytes()))

	var schemaProof *types.SchemaProof = &types.SchemaProof{
		Type:               "Ed25519Signature2020",
		Created:            "2022-04-10T04:07:12Z",
		VerificationMethod: verficationMethodId,
		ProofValue:         schemaDocumentSignature,
		ProofPurpose:       "assertionMethod",
	}

	return SchemaRpcElements{
		SchemaDocument: schemaDocument,
		SchemaProof:    schemaProof,
		Creator:        Creator,
	}
}

func GenerateCredStatusRPCElements(keyPair GenericKeyPair, Id string, verficationMethod *types.VerificationMethod) CredRpcElements {
	var credentialId = "vc:" + DidMethod + ":" + "devnet:" + strings.Split(Id, ":")[3]
	var credHash = sha256.Sum256([]byte("Hash1234"))
	var credentialStatus *types.CredentialStatus = &types.CredentialStatus{
		Claim: &types.Claim{
			Id:            credentialId,
			CurrentStatus: "Live",
			StatusReason:  "Valid",
		},
		Issuer:         Id,
		IssuanceDate:   "2022-04-10T04:07:12Z",
		ExpirationDate: "2023-02-22T13:45:55Z",
		CredentialHash: hex.EncodeToString(credHash[:]),
	}

	var credentialStatusSignature string = base64.StdEncoding.EncodeToString(
		SignGeneric(keyPair, credentialStatus.GetSignBytes()))

	var credentialProof *types.CredentialProof = &types.CredentialProof{
		Type:               "Ed25519Signature2020",
		Created:            "2022-04-10T04:07:12Z",
		Updated:            "2022-04-10T04:07:12Z",
		VerificationMethod: verficationMethod.Id,
		ProofValue:         credentialStatusSignature,
		ProofPurpose:       "assertionMethod",
	}

	return CredRpcElements{
		Status:  credentialStatus,
		Proof:   credentialProof,
		Creator: Creator,
	}
}

func GenerateSecp256k1KeyPair() secp256k1KeyPair {
	privateKey := secp256k1.GenPrivKey()

	publicKey := privateKey.PubKey().Bytes()

	publicKeyMultibase, err := multibase.Encode(multibase.Base58BTC, publicKey)
	if err != nil {
		panic("Error while encoding multibase string")
	}
	return secp256k1KeyPair{
		publicKey:  publicKeyMultibase,
		privateKey: &privateKey,
	}
}

func GenerateEd25519KeyPair() ed25519KeyPair {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	publicKeyBase58Encoded, err := multibase.Encode(multibase.Base58BTC, publicKey)
	if err != nil {
		panic("Error while encoding multibase string")
	}

	return ed25519KeyPair{
		publicKey:  publicKeyBase58Encoded,
		privateKey: privateKey,
	}
}

func CreateDidTx(msgServer types.MsgServer, ctx context.Context, keyPair ed25519KeyPair) (string, error) {
	rpcElements := GenerateDidDocumentRPCElements(keyPair)

	msgCreateDID := &types.MsgCreateDID{
		DidDocString: rpcElements.DidDocument,
		Signatures:   rpcElements.Signatures,
		Creator:      rpcElements.Creator,
	}

	_, err := msgServer.CreateDID(ctx, msgCreateDID)
	if err != nil {
		return "", err
	}

	return rpcElements.DidDocument.Id, nil
}

func UpdateDidTx(msgServer types.MsgServer, ctx context.Context, rpcElements DidRpcElements, versionId string) error {
	msgUpdateDID := &types.MsgUpdateDID{
		DidDocString: rpcElements.DidDocument,
		Signatures:   rpcElements.Signatures,
		Creator:      rpcElements.Creator,
		VersionId:    versionId,
	}

	_, err := msgServer.UpdateDID(ctx, msgUpdateDID)
	if err != nil {
		return err
	}
	return nil
}

func QueryDid(k *keeper.Keeper, ctx sdk.Context, Id string) *types.DidDocumentState {
	resolvedDidDocument, errResolve := k.GetDid(&ctx, Id)
	if errResolve != nil {
		panic(errResolve)
	}

	return resolvedDidDocument
}
