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

type JsonLdCredentialStatusBJJ struct {
	Context                  []contextObject     `json:"@context,omitempty"`
	Id                       string              `json:"id,omitempty"`
	Revoked                  bool                `json:"revoked,omitempty"`
	Suspended                bool                `json:"suspended,omitempty"`
	Remarks                  string              `json:"remarks,omitempty"`
	Issuer                   string              `json:"issuer,omitempty"`
	IssuanceDate             string              `json:"issuanceDate,omitempty"`
	CredentialMerkleRootHash string              `json:"credentialMerkleRootHash,omitempty"`
	Proof                    JsonLdDocumentProof `json:"proof,omitempty"`
}

func (doc *JsonLdCredentialStatusBJJ) GetContext() []contextObject {
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

func NewJsonLdCredentialStatusBJJ(credStatusDoc *types.CredentialStatusDocument, docProof *types.DocumentProof) *JsonLdCredentialStatusBJJ {
	if len(credStatusDoc.Context) == 0 {
		panic("atleast one context url must be provided in the Credential Status Document for Canonization")
	}

	var jsonLdCredentialStatus *JsonLdCredentialStatusBJJ = &JsonLdCredentialStatusBJJ{}

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

	jsonLdCredentialStatus.Proof.Type = docProof.Type
	jsonLdCredentialStatus.Proof.Created = docProof.Created
	jsonLdCredentialStatus.Proof.ProofPurpose = docProof.ProofPurpose
	jsonLdCredentialStatus.Proof.VerificationMethod = docProof.VerificationMethod

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

type JsonLdCredentialSchemaBJJ struct {
	Context      []contextObject                 `json:"@context,omitempty"`
	Type         string                          `json:"type,omitempty"`
	ModelVersion string                          `json:"modelVersion,omitempty"`
	Id           string                          `json:"id,omitempty"`
	Name         string                          `json:"name,omitempty"`
	Author       string                          `json:"author,omitempty"`
	Authored     string                          `json:"authored,omitempty"`
	Schema       *types.CredentialSchemaProperty `json:"schema,omitempty"`
	Proof        JsonLdDocumentProof             `json:"proof,omitempty"`
}

func (doc *JsonLdCredentialSchemaBJJ) GetContext() []contextObject {
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

func NewJsonLdCredentialSchemaBJJ(credSchema *types.CredentialSchemaDocument, docProof *types.DocumentProof) *JsonLdCredentialSchemaBJJ {
	if len(credSchema.Context) == 0 {
		panic("atleast one context url must be provided for DID Document for Canonization")
	}

	var jsonLdDoc *JsonLdCredentialSchemaBJJ = &JsonLdCredentialSchemaBJJ{}

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

	jsonLdDoc.Proof.Type = docProof.Type
	jsonLdDoc.Proof.Created = docProof.Created
	jsonLdDoc.Proof.ProofPurpose = docProof.ProofPurpose
	jsonLdDoc.Proof.VerificationMethod = docProof.VerificationMethod

	return jsonLdDoc
}

// It is a similar to `Did` struct, with the exception that the `context` attribute is of type
// `contextObject` instead of `[]string`, which is meant for accomodating Context JSON body
// having arbritrary attributes. It should be used for performing Canonization.
type JsonLdDidDocumentWithoutVM struct {
	Context    []contextObject `json:"@context,omitempty"`
	Id         string          `json:"id,omitempty"`
	Controller []string        `json:"controller,omitempty"`
	// AlsoKnownAs     []string                              `json:"alsoKnownAs,omitempty"`
	Authentication       []verificationMethodWithoutController `json:"authentication,omitempty"`
	AssertionMethod      []verificationMethodWithoutController `json:"assertionMethod,omitempty"`
	CapabilityDelegation []verificationMethodWithoutController `json:"capabilityDelegation,omitempty"`
	CapabilityInvocation []verificationMethodWithoutController `json:"capabilityInvocation,omitempty"`
	KeyAgreement         []verificationMethodWithoutController `json:"keyAgreement,omitempty"`
	Proof                JsonLdDocumentProof                   `json:"proof,omitempty"`
	Service              []*types.Service                      `protobuf:"bytes,11,rep,name=service,proto3" json:"service,omitempty"`
}

func (doc *JsonLdDidDocumentWithoutVM) GetContext() []contextObject {
	return doc.Context
}

// NewJsonLdDidDocument returns a new JsonLdDid struct from input Did
func NewJsonLdDidDocumentWithoutVM(didDoc *types.DidDocument, docProof *types.DocumentProof) *JsonLdDidDocumentWithoutVM {
	if len(didDoc.Context) == 0 {
		panic("atleast one context url must be provided for DID Document for Canonization")
	}

	var jsonLdDoc *JsonLdDidDocumentWithoutVM = &JsonLdDidDocumentWithoutVM{}

	for _, url := range didDoc.Context {
		contextObj, ok := ContextUrlMap[url]
		if !ok {
			panic(fmt.Sprintf("invalid or unsupported context url: %v", url))
		}
		jsonLdDoc.Context = append(jsonLdDoc.Context, contextObj)
	}

	jsonLdDoc.Id = didDoc.Id
	jsonLdDoc.Controller = didDoc.Controller
	// Replace verification method ids with their corresponding Verification Method object
	var vmMap map[string]verificationMethodWithoutController = map[string]verificationMethodWithoutController{}

	for _, vm := range didDoc.VerificationMethod {
		vmMap[vm.Id] = newVerificationMethodWithoutController(vm)
	}

	// If Authentication and AssertionMethod are empty, then populate the
	// verification methods in AssertionMethod
	if len(didDoc.Authentication) == 0 && len(didDoc.AssertionMethod) == 0 {
		for _, vm := range vmMap {
			jsonLdDoc.AssertionMethod = append(jsonLdDoc.AssertionMethod, vm)
			jsonLdDoc.AssertionMethod[len(jsonLdDoc.AssertionMethod)-1].Id = jsonLdDoc.AssertionMethod[len(jsonLdDoc.AssertionMethod)-1].Id + "assertionMethod"
		}
	} else {
		for _, vmId := range didDoc.Authentication {
			vmObj := vmMap[vmId]
			jsonLdDoc.Authentication = append(jsonLdDoc.Authentication, vmObj)
			jsonLdDoc.Authentication[len(jsonLdDoc.Authentication)-1].Id = jsonLdDoc.Authentication[len(jsonLdDoc.Authentication)-1].Id + "authentication"
		}
		for _, vmId := range didDoc.AssertionMethod {
			vmObj := vmMap[vmId]
			jsonLdDoc.AssertionMethod = append(jsonLdDoc.AssertionMethod, vmObj)
			jsonLdDoc.AssertionMethod[len(jsonLdDoc.AssertionMethod)-1].Id = jsonLdDoc.AssertionMethod[len(jsonLdDoc.AssertionMethod)-1].Id + "assertionMethod"
		}

		for _, vmId := range didDoc.CapabilityDelegation {
			vmObj := vmMap[vmId]
			jsonLdDoc.CapabilityDelegation = append(jsonLdDoc.CapabilityDelegation, vmObj)
			jsonLdDoc.CapabilityDelegation[len(jsonLdDoc.CapabilityDelegation)-1].Id = jsonLdDoc.CapabilityDelegation[len(jsonLdDoc.CapabilityDelegation)-1].Id + "capabilityDelegation"
		}

		for _, vmId := range didDoc.CapabilityInvocation {
			vmObj := vmMap[vmId]
			jsonLdDoc.CapabilityInvocation = append(jsonLdDoc.CapabilityInvocation, vmObj)
			jsonLdDoc.CapabilityInvocation[len(jsonLdDoc.CapabilityInvocation)-1].Id = jsonLdDoc.CapabilityInvocation[len(jsonLdDoc.CapabilityInvocation)-1].Id + "capabilityInvocation"
		}

		for _, vmId := range didDoc.KeyAgreement {
			vmObj := vmMap[vmId]
			jsonLdDoc.KeyAgreement = append(jsonLdDoc.KeyAgreement, vmObj)
			jsonLdDoc.KeyAgreement[len(jsonLdDoc.KeyAgreement)-1].Id = jsonLdDoc.KeyAgreement[len(jsonLdDoc.KeyAgreement)-1].Id + "keyAgreement"
		}
	}

	jsonLdDoc.Proof.Type = docProof.Type
	jsonLdDoc.Proof.Created = docProof.Created
	jsonLdDoc.Proof.ProofPurpose = docProof.ProofPurpose
	jsonLdDoc.Proof.VerificationMethod = docProof.VerificationMethod + docProof.ProofPurpose
	jsonLdDoc.Service = didDoc.Service
	return jsonLdDoc
}

type verificationMethodWithoutController struct {
	Id                 string `json:"id,omitempty"`
	Type               string `json:"type,omitempty"`
	PublicKeyMultibase string `json:"publicKeyMultibase,omitempty"`
}

func newVerificationMethodWithoutController(vm *types.VerificationMethod) verificationMethodWithoutController {
	var vmWithoutController verificationMethodWithoutController = verificationMethodWithoutController{
		Id:                 vm.Id,
		Type:               vm.Type,
		PublicKeyMultibase: vm.PublicKeyMultibase,
	}
	return vmWithoutController
}
