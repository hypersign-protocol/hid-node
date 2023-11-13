package ldcontext

import (
	"encoding/json"
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/piprate/json-gold/ld"
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

// It is a similar to `Did` struct, with the exception that the `context` attribute is of type
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

// normalizeWithURDNA2015 performs RDF Canonization upon JsonLdDid using URDNA2015
// algorithm and returns the canonized document in string
func normalizeWithURDNA2015(jsonLdDocument JsonLdDocument) (string, error) {
	return normalize(ld.AlgorithmURDNA2015, jsonLdDocument)
}

func normalize(algorithm string, jsonLdDocument JsonLdDocument) (string, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Algorithm = algorithm // ld.AlgorithmURDNA2015
	options.Format = "application/n-quads"

	normalisedJsonLd, err := proc.Normalize(jsonLdDocToInterface(jsonLdDocument), options)
	if err != nil {
		return "", fmt.Errorf("unable to Normalize DID Document: %v", err.Error())
	}

	canonizedDocString := normalisedJsonLd.(string)
	if canonizedDocString == "" {
		return "", fmt.Errorf("normalization of JSON-LD document yielded empty RDF string")
	}

	return canonizedDocString, nil
}

// Convert JsonLdDid to interface
func jsonLdDocToInterface(jsonLd any) interface{} {
	var intf interface{}

	jsonLdBytes, err := json.Marshal(jsonLd)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonLdBytes, &intf)
	if err != nil {
		panic(err)
	}

	return intf
}
