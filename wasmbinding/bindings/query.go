package bindings

import (
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// SsiContractQuery contains custom queries for x/ssi chain
type SsiContractQuery struct {
	DidDocument *QueryDidDocument `json:"did_document,omitempty"`
	DidDocumentFromAddress *QueryDidDocumentFromAddress `json:"did_document_from_address,omitempty"`
}

type QueryDidDocument struct {
	DidId string `json:"did_id,omitempty"`
}

type QueryDidDocumentResponse struct {
	Context              []string                    `json:"context"`
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

type QueryDidDocumentFromAddress struct {
	Address string `json:"address,omitempty"`
}

type QueryDidDocumentFromAddressResponse struct {
	DidId string `json:"did_id,omitempty"`
}

