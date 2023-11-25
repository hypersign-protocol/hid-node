package ldcontext

import (
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

type contextObject map[string]interface{}

type JsonLdDocument interface {
	GetContext() []contextObject
}

// It is a similar to `Did` struct, with the exception that the `context` attribute is of type
// `contextObject` instead of `[]string`, which is meant for accomodating Context JSON body
// having arbritrary attributes. It should be used for performing Canonization.
type JsonLdDidDocument struct {
	Context              []contextObject             `json:"@context,omitempty"`
	Id                   string                      `json:"id,omitempty"`
	Controller           []string                    `json:"controller,omitempty"`
	AlsoKnownAs          []string                    `json:"alsoKnownAs,omitempty"`
	VerificationMethod   []*types.VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication       []string                    `json:"authentication,omitempty"`
	AssertionMethod      []string                    `json:"assertionMethod,omitempty"`
	KeyAgreement         []string                    `json:"keyAgreement,omitempty"`
	CapabilityInvocation []string                    `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation []string                    `json:"capabilityDelegation,omitempty"`
	Service              []*types.Service            `json:"service,omitempty"`
}

func (doc *JsonLdDidDocument) GetContext() []contextObject {
	return doc.Context
}

// NewJsonLdDidDocument returns a new JsonLdDid struct from input Did
func NewJsonLdDidDocument(didDoc *types.DidDocument) *JsonLdDidDocument {
	if len(didDoc.Context) == 0 {
		panic("atleast one context url must be provided for DID Document for Canonization")
	}

	var jsonLdDoc *JsonLdDidDocument = &JsonLdDidDocument{}

	for _, url := range didDoc.Context {
		contextObj, ok := ContextUrlMap[url]
		if !ok {
			panic(fmt.Sprintf("invalid or unsupported context url: %v", url))
		}
		jsonLdDoc.Context = append(jsonLdDoc.Context, contextObj)
	}

	jsonLdDoc.Id = didDoc.Id
	jsonLdDoc.AlsoKnownAs = didDoc.AlsoKnownAs
	jsonLdDoc.AssertionMethod = didDoc.AssertionMethod
	jsonLdDoc.Authentication = didDoc.Authentication
	jsonLdDoc.CapabilityDelegation = didDoc.CapabilityDelegation
	jsonLdDoc.CapabilityInvocation = didDoc.CapabilityInvocation
	jsonLdDoc.Service = didDoc.Service
	jsonLdDoc.VerificationMethod = didDoc.VerificationMethod
	jsonLdDoc.Controller = didDoc.Controller
	jsonLdDoc.KeyAgreement = didDoc.KeyAgreement

	return jsonLdDoc
}

// It is a similar to `CredentialStatusDocument` struct, with the exception that the `context` attribute is of type
// `contextObject` instead of `[]string`, which is meant for accomodating Context JSON body
// having arbritrary attributes. It should be used for performing Canonization.
type JsonLdCredentialStatus struct {
	Context                  []contextObject `json:"@context,omitempty"`
	Id                       string          `json:"id,omitempty"`
	Revoked                  bool            `json:"revoked,omitempty"`
	Suspended                bool            `json:"suspended,omitempty"`
	Remarks                  string          `json:"remarks,omitempty"`
	Issuer                   string          `json:"issuer,omitempty"`
	IssuanceDate             string          `json:"issuanceDate,omitempty"`
	CredentialMerkleRootHash string          `json:"credentialMerkleRootHash,omitempty"`
}

func (doc *JsonLdCredentialStatus) GetContext() []contextObject {
	return doc.Context
}

// NewJsonLdCredentialStatus returns a new JsonLdCredentialStatus struct from input Credential Status
func NewJsonLdCredentialStatus(credStatusDoc *types.CredentialStatusDocument) *JsonLdCredentialStatus {
	if len(credStatusDoc.Context) == 0 {
		panic("atleast one context url must be provided in the Credential Status Document for Canonization")
	}

	var jsonLdCredentialStatus *JsonLdCredentialStatus = &JsonLdCredentialStatus{}

	for _, url := range credStatusDoc.Context {
		contextObj, ok := ContextUrlMap[url]
		if !ok {
			panic(fmt.Sprintf("invalid or unsupported context url: %v", url))
		}
		jsonLdCredentialStatus.Context = append(jsonLdCredentialStatus.Context, contextObj)
	}

	jsonLdCredentialStatus.Id = credStatusDoc.Id
	jsonLdCredentialStatus.Revoked = credStatusDoc.Revoked
	jsonLdCredentialStatus.Remarks = credStatusDoc.Remarks
	jsonLdCredentialStatus.Suspended = credStatusDoc.Suspended
	jsonLdCredentialStatus.Issuer = credStatusDoc.Issuer
	jsonLdCredentialStatus.IssuanceDate = credStatusDoc.IssuanceDate
	jsonLdCredentialStatus.CredentialMerkleRootHash = credStatusDoc.CredentialMerkleRootHash

	return jsonLdCredentialStatus
}

// Document Proof

type JsonLdDocumentProof struct {
	Context            []contextObject `json:"@context,omitempty"`
	Type               string          `json:"type,omitempty"`
	Created            string          `json:"created,omitempty"`
	VerificationMethod string          `json:"verificationMethod,omitempty"`
	ProofPurpose       string          `json:"proofPurpose,omitempty"`
}

func (doc *JsonLdDocumentProof) GetContext() []contextObject {
	return doc.Context
}

func NewJsonLdDocumentProof(didDocProof *types.DocumentProof, didContexts []string) *JsonLdDocumentProof {
	if len(didContexts) == 0 {
		panic("atleast one context url must be provided for DID Document for Canonization")
	}

	var jsonLdDoc *JsonLdDocumentProof = &JsonLdDocumentProof{}

	for _, url := range didContexts {
		contextObj, ok := ContextUrlMap[url]
		if !ok {
			panic(fmt.Sprintf("invalid or unsupported context url: %v", url))
		}
		jsonLdDoc.Context = append(jsonLdDoc.Context, contextObj)
	}

	jsonLdDoc.Created = didDocProof.Created
	jsonLdDoc.ProofPurpose = didDocProof.ProofPurpose
	jsonLdDoc.Type = didDocProof.Type
	jsonLdDoc.VerificationMethod = didDocProof.VerificationMethod

	return jsonLdDoc
}

// It is a similar to `CredentialSchemaDocument` struct, with the exception that the `context` attribute is of type
// `contextObject` instead of `[]string`, which is meant for accomodating Context JSON body
// having arbritrary attributes. It should be used for performing Canonization.
type JsonLdCredentialSchema struct {
	Context      []contextObject                 `json:"@context,omitempty"`
	Type         string                          `json:"type,omitempty"`
	ModelVersion string                          `json:"modelVersion,omitempty"`
	Id           string                          `json:"id,omitempty"`
	Name         string                          `json:"name,omitempty"`
	Author       string                          `json:"author,omitempty"`
	Authored     string                          `json:"authored,omitempty"`
	Schema       *types.CredentialSchemaProperty `json:"schema,omitempty"`
}

func (doc *JsonLdCredentialSchema) GetContext() []contextObject {
	return doc.Context
}

func NewJsonLdCredentialSchema(credSchema *types.CredentialSchemaDocument) *JsonLdCredentialSchema {
	if len(credSchema.Context) == 0 {
		panic("atleast one context url must be provided for DID Document for Canonization")
	}

	var jsonLdDoc *JsonLdCredentialSchema = &JsonLdCredentialSchema{}

	for _, url := range credSchema.Context {
		contextObj, ok := ContextUrlMap[url]
		if !ok {
			panic(fmt.Sprintf("invalid or unsupported context url: %v", url))
		}
		jsonLdDoc.Context = append(jsonLdDoc.Context, contextObj)
	}

	jsonLdDoc.Type = credSchema.Type
	jsonLdDoc.ModelVersion = credSchema.ModelVersion
	jsonLdDoc.Id = credSchema.Id
	jsonLdDoc.Name = credSchema.Name
	jsonLdDoc.Author = credSchema.Author
	jsonLdDoc.Authored = credSchema.Authored
	jsonLdDoc.Schema = credSchema.Schema

	return jsonLdDoc
}
