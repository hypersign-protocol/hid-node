package ldcontext

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/piprate/json-gold/ld"
)

// NormalizeByVerificationMethodType normalizes DID Document based on the input Verification
// Method type
func NormalizeByVerificationMethodType(ssiMsg types.SsiMsg, vmType string, didDocumentProof *types.DocumentProof) ([]byte, error) {
	switch vmType {
	case types.Ed25519VerificationKey2020:
		msgBytes, err := Ed25519Signature2020Normalize(ssiMsg, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return msgBytes, nil
	case types.EcdsaSecp256k1RecoveryMethod2020:
		msgBytes, err := EcdsaSecp256k1RecoverySignature2020Normalize(ssiMsg, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return msgBytes, nil
	case types.Bls12381G2Key2020:
		msgBytes, err := BbsBlsSignature2020Normalize(ssiMsg, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return msgBytes, nil
	case types.EcdsaSecp256k1VerificationKey2019:
		msgBytes, err := EcdsaSecp256k1Signature2019Normalize(ssiMsg, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return msgBytes, nil
	default:
		return ssiMsg.GetSignBytes(), nil
	}
}

// normalizeDocumentWithProof normalizes the SSI document along with Document Proof
// Read more: https://w3c.github.io/vc-di-eddsa/#representation-ed25519signature2020
func normalizeDocumentWithProof(msg types.SsiMsg, docProof *types.DocumentProof) ([]byte, error) {
	// Normalize Document
	var canonizedDocument string
	var context []string

	switch doc := msg.(type) {
	case *types.DidDocument:
		var err error
		jsonLdDidDocument := NewJsonLdDidDocument(doc)
		canonizedDocument, err = normalizeWithURDNA2015(jsonLdDidDocument)
		if err != nil {
			return nil, err
		}
		context = doc.Context
	case *types.CredentialStatusDocument:
		var err error
		jsonLdCredentialStatus := NewJsonLdCredentialStatus(doc)
		canonizedDocument, err = normalizeWithURDNA2015(jsonLdCredentialStatus)
		if err != nil {
			return nil, err
		}
		context = doc.Context
	case *types.CredentialSchemaDocument:
		var err error
		jsonLdCredentialSchema := NewJsonLdCredentialSchema(doc)
		canonizedDocument, err = normalizeWithURDNA2015(jsonLdCredentialSchema)
		if err != nil {
			return nil, err
		}
		context = doc.Context
	}

	canonizedDocumentHash := sha256.Sum256([]byte(canonizedDocument))

	// Normalize Document Proof
	jsonLdDocumentProof := NewJsonLdDocumentProof(docProof, context)
	canonizedDocumentProof, err := normalizeWithURDNA2015(jsonLdDocumentProof)
	if err != nil {
		return nil, err
	}
	canonizedDocumentProofHash := sha256.Sum256([]byte(canonizedDocumentProof))

	var finalNormalizedHash []byte = []byte{}
	// NOTE: The order is: ProofHash + DocumentHash
	finalNormalizedHash = append(finalNormalizedHash, canonizedDocumentProofHash[:]...)
	finalNormalizedHash = append(finalNormalizedHash, canonizedDocumentHash[:]...)

	return finalNormalizedHash, nil
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

// Ed25519Signature2020Normalize normalizes DID Document in accordance with
// EdDSA Cryptosuite v2020 (https://www.w3.org/community/reports/credentials/CG-FINAL-di-eddsa-2020-20220724/)
func Ed25519Signature2020Normalize(ssiMsg types.SsiMsg, didDocProof *types.DocumentProof) ([]byte, error) {
	return normalizeDocumentWithProof(ssiMsg, didDocProof)
}

// EcdsaSecp256k1RecoverySignature2020Normalize normalizes DID Document in accordance with
// the Identity Foundation draft on EcdsaSecp256k1RecoverySignature2020
// Read more: https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/
func EcdsaSecp256k1RecoverySignature2020Normalize(ssiMsg types.SsiMsg, didDocProof *types.DocumentProof) ([]byte, error) {
	return normalizeDocumentWithProof(ssiMsg, didDocProof)
}

// BbsBlsSignature2020Normalize normalizes the DID Document for the
// BbsBlsSignature2020 signature type
// Read more: https://identity.foundation/bbs-signature/draft-irtf-cfrg-bbs-signatures.html
func BbsBlsSignature2020Normalize(ssiMsg types.SsiMsg, didDocProof *types.DocumentProof) ([]byte, error) {
	return normalizeDocumentWithProof(ssiMsg, didDocProof)
}

// EcdsaSecp256k1Signature2019Normalize normalizes the DID Document for the
// EcdsaSecp256k1Signature2019 signature type
// Read more: https://w3c-ccg.github.io/lds-ecdsa-secp256k1-2019/
func EcdsaSecp256k1Signature2019Normalize(ssiMsg types.SsiMsg, didDocProof *types.DocumentProof) ([]byte, error) {
	return normalizeDocumentWithProof(ssiMsg, didDocProof)
}
