package ldcontext

import (
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/piprate/json-gold/ld"
)

type contextObject map[string]interface{}

// It is a similar to `Did` struct, with the exception that the `context` attribute is of type
// `contextObject` instead of `[]string`, which is meant for accomodating Context JSON body
// having arbritrary attributes. It should be used for performing Canonization.
type JsonLdDid struct {
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

// NewJsonLdDid returns a new JsonLdDid struct from input Did
func NewJsonLdDid(didDoc *types.Did) *JsonLdDid {
	if len(didDoc.Context) == 0 {
		panic("atleast one context url must be provided for DID Document for Canonization")
	}

	var jsonLdDoc *JsonLdDid = &JsonLdDid{}

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

// NormalizeWithURDNA2015 performs RDF Canonization upon JsonLdDid using URDNA2015 
// algorithm and returns the canonized document in string
func (doc *JsonLdDid) NormalizeWithURDNA2015() (string, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Algorithm = ld.AlgorithmURDNA2015
	options.Format = "application/n-quads"

	normalisedJsonLdDid, err := proc.Normalize(doc, options)
	if err != nil {
		return "", fmt.Errorf("unable to Normalize DID Document: %v", err.Error())
	}

	canonizedDocString := normalisedJsonLdDid.(string)
	return canonizedDocString, nil
}
