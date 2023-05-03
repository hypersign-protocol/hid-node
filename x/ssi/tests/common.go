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

func GenerateDidDocumentRPCElements(keyPair GenericKeyPair, signingElements []DidSigningElements) DidRpcElements {
	publicKey, optionalID := GetPublicKeyAndOptionalID(keyPair)
	var didId string
	if optionalID == "" {
		didId = "did:" + DidMethod + ":" + ChainNamespace + ":" + publicKey
	} else {
		didId = "did:" + DidMethod + ":" + ChainNamespace + ":" + optionalID
	}

	var verificationMethodId string = didId + "#" + "key-1"

	var vmType string
	switch keyPair.(type) {
	case ed25519KeyPair:
		vmType = types.Ed25519VerificationKey2020
	case secp256k1KeyPair:
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
				stripDidFromVerificationMethod(signingElements[i].vmId))
		}
	} else {
		signingElements = []DidSigningElements{
			{
				keyPair: keyPair,
				vmId:    vm.Id,
			},
		}
		controllers = []string{didId}
	}

	var didDocument *types.Did = &types.Did{
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

	var signInfo []*types.SignInfo = getDidSigningInfo(didDocument, signingElements)

	return DidRpcElements{
		DidDocument: didDocument,
		Signatures:  signInfo,
		Creator:     Creator,
	}
}

func GenerateSchemaDocumentRPCElements(keyPair GenericKeyPair, authorId string, verficationMethodId string) SchemaRpcElements {
	var schemaId string = "sch:" + DidMethod + ":" + "devnet" + ":" + strings.Split(authorId, ":")[3] + ":1.0"
	var schemaDocument *types.SchemaDocument = &types.SchemaDocument{
		Type:         "https://w3c-ccg.github.io/vc-json-schemas/schema/1.0/schema.json",
		ModelVersion: "v1.0",
		Name:         "HS Credential",
		Author:       authorId,
		Id:           schemaId,
		Authored:     "2022-04-10T04:07:12Z",
		Schema: &types.SchemaProperty{
			Description:          "test",
			Type:                 "Object",
			Properties:           "{myString:{type:string}}",
			Required:             []string{"myString"},
			AdditionalProperties: false,
		},
	}

	var schemaDocumentSignature string = base64.StdEncoding.EncodeToString(
		SignGeneric(keyPair, schemaDocument.GetSignBytes()),
	)

	var proofType string
	switch keyPair.(type) {
	case ed25519KeyPair:
		proofType = types.VerificationKeySignatureMap["Ed25519VerificationKey2020"]
	case secp256k1KeyPair:
		proofType = types.VerificationKeySignatureMap["EcdsaSecp256k1VerificationKey2019"]
	}

	var schemaProof *types.SchemaProof = &types.SchemaProof{
		Type:               proofType,
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

func GenerateCredStatusRPCElements(keyPair GenericKeyPair, issuerId string, verficationMethod *types.VerificationMethod) CredRpcElements {
	var credentialId = "vc:" + DidMethod + ":" + "devnet:" + strings.Split(issuerId, ":")[3]
	var credHash = sha256.Sum256([]byte("Hash1234"))
	var credentialStatus *types.CredentialStatus = &types.CredentialStatus{
		Claim: &types.Claim{
			Id:            credentialId,
			CurrentStatus: "Live",
			StatusReason:  "Valid",
		},
		Issuer:         issuerId,
		IssuanceDate:   "2022-04-10T04:07:12Z",
		ExpirationDate: "2023-02-22T13:45:55Z",
		CredentialHash: hex.EncodeToString(credHash[:]),
	}

	var credentialStatusSignature string = base64.StdEncoding.EncodeToString(
		SignGeneric(keyPair, credentialStatus.GetSignBytes()),
	)

	var proofType string
	switch keyPair.(type) {
	case ed25519KeyPair:
		proofType = types.VerificationKeySignatureMap["Ed25519VerificationKey2020"]
	case secp256k1KeyPair:
		proofType = types.VerificationKeySignatureMap["EcdsaSecp256k1VerificationKey2019"]
	}

	var credentialProof *types.CredentialProof = &types.CredentialProof{
		Type:               proofType,
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

func CreateDidTx(msgServer types.MsgServer, ctx context.Context, keyPair ed25519KeyPair) (string, error) {
	rpcElements := GenerateDidDocumentRPCElements(keyPair, []DidSigningElements{})

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
	resolvedDidDocument, errResolve := k.GetDidDocumentState(&ctx, Id)
	if errResolve != nil {
		panic(errResolve)
	}

	return resolvedDidDocument
}
