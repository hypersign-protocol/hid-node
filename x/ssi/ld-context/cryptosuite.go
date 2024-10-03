package ldcontext

import (
	"context"
	"crypto/sha256"
	"strings"

	"encoding/json"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/iden3/go-schema-processor/merklize"
	"github.com/piprate/json-gold/ld"
)

// Ed25519Signature2020Normalize normalizes DID Document in accordance with
// EdDSA Cryptosuite v2020 (https://www.w3.org/community/reports/credentials/CG-FINAL-di-eddsa-2020-20220724/)
func Ed25519Signature2020Normalize(ssiMsg types.SsiMsg, docProof *types.DocumentProof) ([]byte, error) {
	normalizationAlgorithm := ld.AlgorithmURDNA2015

	// Normalize Document
	normalizedDocumentString, err := normalizeDocument(ssiMsg, normalizationAlgorithm)
	if err != nil {
		return nil, err
	}

	// Normalize Document Proof
	normalizedDocumentProofString, err := normalizeDocumentProof(docProof, normalizationAlgorithm, ssiMsg.GetContext())
	if err != nil {
		return nil, err
	}

	// Hash Normalized Document and Document Proof strings with SHA-256 hash
	// Combine them in the order: DocumentProofHash + DocumentHash
	normalizedDocumentProofHash := sha256.Sum256([]byte(normalizedDocumentProofString))
	normalizedDocumentHash := sha256.Sum256([]byte(normalizedDocumentString))

	var combinedHash []byte
	combinedHash = append(combinedHash, normalizedDocumentProofHash[:]...)
	combinedHash = append(combinedHash, normalizedDocumentHash[:]...)
	return combinedHash, nil
}

// EcdsaSecp256k1RecoverySignature2020Normalize normalizes DID Document in accordance with
// the Identity Foundation draft on EcdsaSecp256k1RecoverySignature2020
// Read more: https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/
func EcdsaSecp256k1RecoverySignature2020Normalize(ssiMsg types.SsiMsg, docProof *types.DocumentProof) ([]byte, error) {
	normalizationAlgorithm := ld.AlgorithmURDNA2015

	// Normalize Document
	normalizedDocumentString, err := normalizeDocument(ssiMsg, normalizationAlgorithm)
	if err != nil {
		return nil, err
	}

	// Normalize Document Proof
	normalizedDocumentProofString, err := normalizeDocumentProof(docProof, normalizationAlgorithm, ssiMsg.GetContext())
	if err != nil {
		return nil, err
	}

	// Hash Normalized Document and Document Proof strings with SHA-256 hash
	// Combine them in the order: DocumentProofHash + DocumentHash
	normalizedDocumentProofHash := sha256.Sum256([]byte(normalizedDocumentProofString))
	normalizedDocumentHash := sha256.Sum256([]byte(normalizedDocumentString))

	var combinedHash []byte
	combinedHash = append(combinedHash, normalizedDocumentProofHash[:]...)
	combinedHash = append(combinedHash, normalizedDocumentHash[:]...)
	return combinedHash, nil
}

// BbsBlsSignature2020Normalize normalizes the DID Document for the
// BbsBlsSignature2020 signature type
// Read more: https://identity.foundation/bbs-signature/draft-irtf-cfrg-bbs-signatures.html
func BbsBlsSignature2020Normalize(ssiMsg types.SsiMsg, docProof *types.DocumentProof) ([]byte, error) {
	normalizationAlgorithm := ld.AlgorithmURDNA2015

	// Normalize Document
	normalizedDocumentString, err := normalizeDocument(ssiMsg, normalizationAlgorithm)
	if err != nil {
		return nil, err
	}

	// Normalize Document Proof
	normalizedDocumentProofString, err := normalizeDocumentProof(docProof, normalizationAlgorithm, ssiMsg.GetContext())
	if err != nil {
		return nil, err
	}

	// Hash Normalized Document and Document Proof strings with SHA-256 hash
	// Combine them in the order: DocumentProofHash + DocumentHash
	normalizedDocumentProofHash := sha256.Sum256([]byte(normalizedDocumentProofString))
	normalizedDocumentHash := sha256.Sum256([]byte(normalizedDocumentString))

	var combinedHash []byte
	combinedHash = append(combinedHash, normalizedDocumentProofHash[:]...)
	combinedHash = append(combinedHash, normalizedDocumentHash[:]...)
	return combinedHash, nil
}

// EcdsaSecp256k1Signature2019Normalize normalizes the DID Document for the
// EcdsaSecp256k1Signature2019 signature type
// Read more: https://w3c-ccg.github.io/lds-ecdsa-secp256k1-2019/
func EcdsaSecp256k1Signature2019Normalize(ssiMsg types.SsiMsg, docProof *types.DocumentProof) ([]byte, error) {
	normalizationAlgorithm := ld.AlgorithmURDNA2015

	// Normalize Document
	normalizedDocumentString, err := normalizeDocument(ssiMsg, normalizationAlgorithm)
	if err != nil {
		return nil, err
	}

	// Normalize Document Proof
	normalizedDocumentProofString, err := normalizeDocumentProof(docProof, normalizationAlgorithm, ssiMsg.GetContext())
	if err != nil {
		return nil, err
	}

	// Hash Normalized Document and Document Proof strings with SHA-256 hash
	// Combine them in the order: DocumentProofHash + DocumentHash
	normalizedDocumentProofHash := sha256.Sum256([]byte(normalizedDocumentProofString))
	normalizedDocumentHash := sha256.Sum256([]byte(normalizedDocumentString))

	var combinedHash []byte
	combinedHash = append(combinedHash, normalizedDocumentProofHash[:]...)
	combinedHash = append(combinedHash, normalizedDocumentHash[:]...)
	return combinedHash, nil
}

// BJJSignature2021Normalize performs canonization of SSI documents
// based on the spec: https://iden3-communication.io/BJJSignature2021/
func BJJSignature2021Normalize(ssiMsg types.SsiMsg, docProof *types.DocumentProof) ([]byte, error) {
	var jsonLDString string
	switch doc := ssiMsg.(type) {
	case *types.DidDocument:
		jsonLDBytes, err := json.Marshal(NewJsonLdDidDocumentWithoutVM(doc, docProof))
		if err != nil {
			return nil, err
		}
		jsonLDString = string(jsonLDBytes)
	case *types.CredentialSchemaDocument:
		credentialSchemaDocument := NewJsonLdCredentialSchemaBJJ(doc, docProof)
		jsonLDBytes, err := json.Marshal(credentialSchemaDocument)
		if err != nil {
			return nil, err
		}
		jsonLDString = string(jsonLDBytes)
	case *types.CredentialStatusDocument:
		credentialStatusDocument := NewJsonLdCredentialStatusBJJ(doc, docProof)
		jsonLDBytes, err := json.Marshal(credentialStatusDocument)
		if err != nil {
			return nil, err
		}
		jsonLDString = string(jsonLDBytes)
	}

	// The following canonization is done in order to check whether the canonized string
	// is empty or not
	_, err := normalizeDocument(ssiMsg, ld.AlgorithmURDNA2015)
	if err != nil {
		return nil, err
	}

	mz, err := merklize.MerklizeJSONLD(context.Background(), strings.NewReader(jsonLDString))
	if err != nil {
		return nil, err
	}
	jsonLdDocumentMerkleRoot := mz.Root().BigInt().Bytes()

	return jsonLdDocumentMerkleRoot, nil
}
