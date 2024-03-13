package ssi

import (
	"strings"

	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"

	testconstants "github.com/hypersign-protocol/hid-node/x/ssi/tests/constants"
	testcrypto "github.com/hypersign-protocol/hid-node/x/ssi/tests/crypto"
)

func GenerateSchema(keyPair testcrypto.IKeyPair, authorId string) *types.CredentialSchemaDocument {
	var schemaId = "sch:" + testconstants.DidMethod + ":" + testconstants.ChainNamespace + ":" + strings.Split(authorId, ":")[3] + ":" + "1.0"
	var vmContextUrl = GetContextFromKeyPair(keyPair)

	var credentialSchema *types.CredentialSchemaDocument = &types.CredentialSchemaDocument{
		Context:      []string{
			ldcontext.CredentialSchemaContext,
			vmContextUrl,
		},
		Type:         "https://w3c-ccg.github.io/vc-json-schemas/v1/schema/1.0/schema.json",
		ModelVersion: "1.0",
		Id:           schemaId,
		Name:         "SomeSchema",
		Author:       authorId,
		Authored:     "2022-04-10T02:07:12Z",
		Schema: &types.CredentialSchemaProperty{
			Schema:               "http://json-schema.org/draft-07/schema",
			Description:          "Student ID Credential Schema",
			Type:                 "https://schema.org/object",
			Properties:           "{\"jayeshL\":{\"type\":\"string\"}}",
			Required:             []string{"jayeshL"},
			AdditionalProperties: false,
		},
	}

	return credentialSchema
}

func GenerateSchemaRPCElements(keyPair testcrypto.IKeyPair, credentialSchema *types.CredentialSchemaDocument, verficationMethod *types.VerificationMethod) *types.MsgRegisterCredentialSchema {

	var credentialProof *types.DocumentProof = &types.DocumentProof{
		Created:            "2022-04-10T04:07:12Z",
		VerificationMethod: verficationMethod.Id,
		ProofPurpose:       "assertionMethod",
	}

	var credentialStatusSignature string = testcrypto.SignGeneric(keyPair, credentialSchema, credentialProof)
	credentialProof.ProofValue = credentialStatusSignature

	return &types.MsgRegisterCredentialSchema{
		CredentialSchemaDocument: credentialSchema,
		CredentialSchemaProof:    credentialProof,
		TxAuthor:                 testconstants.Creator,
	}
}
