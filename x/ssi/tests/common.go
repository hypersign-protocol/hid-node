package tests

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	//"fmt"

	"github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
	//"github.com/hypersign-protocol/hid-node/x/ssi/utils"
)

type ed25519KeyPair struct {
	publicKey  string
	privateKey ed25519.PrivateKey
}

type DidRpcElements struct {
	DidDocument *types.Did
	Signatures  []*types.SignInfo
	Creator     string
}

type SchemaRpcElements struct {
	SchemaDocument *types.SchemaDocument
	SchemaProof    *types.SchemaProof
	Creator        string
}

type CredRpcElements struct {
	Status  *types.CredentialStatus
	Proof   *types.CredentialProof
	Creator string
}

var Creator string = "hid1kxqk5ejca8nfpw8pg47484rppv359xh7qcasy4"

func getSchemaSigningInfo(schemaDoc *types.SchemaDocument, keyPair ed25519KeyPair, vm *types.VerificationMethod) []*types.SignInfo {
	signature := ed25519.Sign(keyPair.privateKey, schemaDoc.GetSignBytes())
	signInfo := &types.SignInfo{
		VerificationMethodId: vm.GetId(),
		Signature:            base64.StdEncoding.EncodeToString(signature),
	}

	return []*types.SignInfo{
		signInfo,
	}
}

func getDidSigningInfo(didDoc *types.Did, keyPair ed25519KeyPair, vmId string) []*types.SignInfo {
	signature := ed25519.Sign(keyPair.privateKey, didDoc.GetSignBytes())
	signInfo := &types.SignInfo{
		VerificationMethodId: vmId,
		Signature:            base64.StdEncoding.EncodeToString(signature),
	}

	return []*types.SignInfo{
		signInfo,
	}
}

func getMultiSigDidSigningInfo(didDoc *types.Did, keyPairs []ed25519KeyPair, vmIds []string) DidRpcElements {
	if len(keyPairs) != len(vmIds) {
		panic("KeyPairs and vmIds lists should be of equal lengths")
	}

	var signInfoList []*types.SignInfo

	for idx := range keyPairs {
		signature := ed25519.Sign(keyPairs[idx].privateKey, didDoc.GetSignBytes())
		signInfo := &types.SignInfo{
			VerificationMethodId: vmIds[idx],
			Signature:            base64.StdEncoding.EncodeToString(signature),
		}
		signInfoList = append(signInfoList, signInfo)
	}

	return DidRpcElements{
		DidDocument: didDoc,
		Signatures:  signInfoList,
		Creator:     Creator,
	}
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
	var signatures []*types.SignInfo = getDidSigningInfo(
		modifiedDidDocument,
		keyPair,
		verificationMethodId,
	)

	return DidRpcElements{
		DidDocument: modifiedDidDocument,
		Signatures:  signatures,
		Creator:     Creator,
	}
}

func GenerateDidDocumentRPCElements(keyPair ed25519KeyPair) DidRpcElements {
	var didMethod string = "hs"
	var didNamespace string = "devnet"
	var didId = "did:" + didMethod + ":" + didNamespace + ":" + keyPair.publicKey

	var verificationMethodId string = didId + "#" + "key-1"

	var vm = &types.VerificationMethod{
		Id:                 verificationMethodId,
		Type:               "Ed25519VerificationKey2020",
		Controller:         didId,
		PublicKeyMultibase: keyPair.publicKey,
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

	var signInfo []*types.SignInfo = getDidSigningInfo(didDocument, keyPair, vm.Id)

	return DidRpcElements{
		DidDocument: didDocument,
		Signatures:  signInfo,
		Creator:     Creator,
	}

}

func GenerateSchemaDocumentRPCElements(keyPair ed25519KeyPair, Id string, verficationMethodId string) SchemaRpcElements {
	var schemaDocument *types.SchemaDocument = &types.SchemaDocument{
		Type:         "https://w3c-ccg.github.io/vc-json-schemas/schema/1.0/schema.json",
		ModelVersion: "v1.0",
		Name:         "HS Credential",
		Author:       Id,
		Id:           fmt.Sprintf("%s;id=%s;version=1.0", Id, utils.UUID()),
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
		ed25519.Sign(keyPair.privateKey, schemaDocument.GetSignBytes()),
	)

	var schemaProof *types.SchemaProof = &types.SchemaProof{
		Type:               "Ed25519VerificationKey2020",
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

func GenerateCredStatusRPCElements(keyPair ed25519KeyPair, Id string, verficationMethod *types.VerificationMethod) CredRpcElements {
	var credentialStatus *types.CredentialStatus = &types.CredentialStatus{
		Claim: &types.Claim{
			Id:            "did:key:" + utils.UUID(),
			CurrentStatus: "Live",
			StatusReason:  "Valid",
		},
		Issuer:         Id,
		IssuanceDate:   "2022-04-10T04:07:12Z",
		ExpirationDate: "2023-02-22T13:45:55Z",
		CredentialHash: "Hash234",
	}

	var credentialStatusSignature string = base64.StdEncoding.EncodeToString(
		ed25519.Sign(keyPair.privateKey, credentialStatus.GetSignBytes()),
	)

	var credentialProof *types.CredentialProof = &types.CredentialProof{
		Type:               "Ed25519VerificationKey2020",
		Created:            "2022-04-10T04:07:12Z",
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

func GeneratePublicPrivateKeyPair() ed25519KeyPair {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	publicKeyBase58Encoded := "z" + base58.Encode(publicKey)

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

func QueryDid(k *keeper.Keeper, ctx sdk.Context, Id string) *types.DidDocument {
	resolvedDidDocument, errResolve := k.GetDid(&ctx, Id)
	if errResolve != nil {
		panic(errResolve)
	}

	return resolvedDidDocument
}
