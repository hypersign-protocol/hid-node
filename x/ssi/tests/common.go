package tests

import (
	"context"
	"strings"

	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
	testconstants "github.com/hypersign-protocol/hid-node/x/ssi/tests/constants"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)


func GenerateSchemaDocumentRPCElements(keyPair testcrypto.IKeyPair, authorId string, verficationMethodId string) *types.MsgRegisterCredentialSchema {
	var schemaId string = "sch:" + testconstants.DidMethod + ":" + testconstants.ChainNamespace + ":" + strings.Split(authorId, ":")[3] + ":1.0"
	var schemaDocument *types.CredentialSchemaDocument = &types.CredentialSchemaDocument{
		Type:         "https://w3c-ccg.github.io/vc-json-schemas/schema/1.0/schema.json",
		ModelVersion: "v1.0",
		Name:         "HS Credential",
		Author:       authorId,
		Id:           schemaId,
		Authored:     "2022-04-10T04:07:12Z",
		Schema: &types.CredentialSchemaProperty{
			Schema:               "https://json-schema.org/draft-07/schema#",
			Description:          "test",
			Type:                 "Object",
			Properties:           "{myString:{type:string}}",
			Required:             []string{"myString"},
			AdditionalProperties: false,
		},
	}

	var schemaProof *types.DocumentProof = &types.DocumentProof{
		Created:            "2022-04-10T04:07:12Z",
		VerificationMethod: verficationMethodId,
		ProofPurpose:       "assertionMethod",
	}

	var schemaDocumentSignature string = testcrypto.SignGeneric(keyPair, schemaDocument, schemaProof)

	schemaProof.ProofValue = schemaDocumentSignature

	return &types.MsgRegisterCredentialSchema{
		CredentialSchemaDocument: schemaDocument,
		CredentialSchemaProof:    schemaProof,
		TxAuthor:                 testconstants.Creator,
	}
}


func UpdateDidTx(msgServer types.MsgServer, ctx context.Context, rpcElements *types.MsgUpdateDID, versionId string) error {
	_, err := msgServer.UpdateDID(ctx, rpcElements)
	if err != nil {
		return err
	}
	return nil
}

