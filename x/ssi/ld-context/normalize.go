package ldcontext

import (
	"encoding/json"
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/piprate/json-gold/ld"
)

// NormalizeByProofType normalizes DID Document based on the input Proof type
func NormalizeByProofType(ssiMsg types.SsiMsg, didDocumentProof *types.DocumentProof) ([]byte, error) {
	switch didDocumentProof.Type {
	case types.Ed25519Signature2020:
		msgBytes, err := Ed25519Signature2020Normalize(ssiMsg, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return msgBytes, nil
	case types.EcdsaSecp256k1RecoverySignature2020:
		msgBytes, err := EcdsaSecp256k1RecoverySignature2020Normalize(ssiMsg, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return msgBytes, nil
	case types.BbsBlsSignature2020:
		msgBytes, err := BbsBlsSignature2020Normalize(ssiMsg, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return msgBytes, nil
	case types.EcdsaSecp256k1Signature2019:
		msgBytes, err := EcdsaSecp256k1Signature2019Normalize(ssiMsg, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return msgBytes, nil
	case types.BJJSignature2021:
		msgBytes, err := BJJSignature2021Normalize(ssiMsg)
		if err != nil {
			return nil, err
		}
		return msgBytes, nil
	default:
		return nil, fmt.Errorf("unsupported proof type: %v", didDocumentProof.Type)
	}
}

func normalizeDocument(msg types.SsiMsg, algorithm string) (string, error) {
	var canonizedDocument string

	switch doc := msg.(type) {
	case *types.DidDocument:
		var err error
		var jsonLdDocument JsonLdDocument = NewJsonLdDidDocument(doc)
		canonizedDocument, err = normalize(jsonLdDocument, algorithm)
		if err != nil {
			return "", err
		}
	case *types.CredentialStatusDocument:
		var err error
		jsonLdCredentialStatus := NewJsonLdCredentialStatus(doc)
		canonizedDocument, err = normalize(jsonLdCredentialStatus, algorithm)
		if err != nil {
			return "", err
		}
	case *types.CredentialSchemaDocument:
		var err error
		jsonLdCredentialSchema := NewJsonLdCredentialSchema(doc)
		canonizedDocument, err = normalize(jsonLdCredentialSchema, algorithm)
		if err != nil {
			return "", err
		}
	}

	return canonizedDocument, nil
}

func normalizeDocumentProof(docProof *types.DocumentProof, algorithm string, docContext []string) (string, error) {
	jsonLdDocumentProof := NewJsonLdDocumentProof(docProof, docContext)
	canonizedDocumentProof, err := normalize(jsonLdDocumentProof, algorithm)
	if err != nil {
		return "", err
	}

	return canonizedDocumentProof, nil
}

func normalize(jsonLdDocument JsonLdDocument, algorithm string) (string, error) {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Algorithm = algorithm
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
