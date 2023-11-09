package ldcontext

import (
	"crypto/sha256"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// NormalizeByVerificationMethodType normalizes DID Document based on the input Verification
// Method type
func NormalizeByVerificationMethodType(didDoc *types.DidDocument, vmType string, didDocumentProof *types.DocumentProof) ([]byte, error) {
	switch vmType {
	case types.Ed25519VerificationKey2020:
		didDocBytes, err := Ed25519Signature2020Normalize(didDoc, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return didDocBytes, nil
	case types.EcdsaSecp256k1RecoveryMethod2020:
		didDocBytes, err := EcdsaSecp256k1RecoverySignature2020Normalize(didDoc, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return didDocBytes, nil
	case types.Bls12381G2Key2020:
		didDocBytes, err := BbsBlsSignature2020Normalize(didDoc, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return didDocBytes, nil
	case types.EcdsaSecp256k1VerificationKey2019:
		didDocBytes, err := EcdsaSecp256k1Signature2019Normalize(didDoc, didDocumentProof)
		if err != nil {
			return nil, err
		}
		return didDocBytes, nil
	default:
		return didDoc.GetSignBytes(), nil
	}
}

// normalizeDocumentWithProof normalizes the DidDocument along with Document Proof
// Read more: https://w3c.github.io/vc-di-eddsa/#representation-ed25519signature2020
func normalizeDocumentWithProof(didDoc *types.DidDocument, didDocProof *types.DocumentProof) ([]byte, error) {
	jsonLdDid := NewJsonLdDid(didDoc)
	canonizedDidDocument, err := jsonLdDid.NormalizeWithURDNA2015()
	if err != nil {
		return nil, err
	}
	canonizedDidDocumentHash := sha256.Sum256([]byte(canonizedDidDocument))

	jsonLdDocumentProof := NewJsonLdDocumentProof(didDocProof, didDoc.Context)
	canonizedDocumentProof, err := jsonLdDocumentProof.NormalizeWithURDNA2015()
	if err != nil {
		return nil, err
	}
	canonizedDocumentProofHash := sha256.Sum256([]byte(canonizedDocumentProof))

	var finalNormalizedHash []byte = []byte{}
	// NOTE: The order is: ProofHash + DocumentHash
	finalNormalizedHash = append(finalNormalizedHash, canonizedDocumentProofHash[:]...)
	finalNormalizedHash = append(finalNormalizedHash, canonizedDidDocumentHash[:]...)

	return finalNormalizedHash, nil
}

// Ed25519Signature2020Normalize normalizes DID Document in accordance with
// EdDSA Cryptosuite v2020 (https://www.w3.org/community/reports/credentials/CG-FINAL-di-eddsa-2020-20220724/)
func Ed25519Signature2020Normalize(didDoc *types.DidDocument, didDocProof *types.DocumentProof) ([]byte, error) {
	return normalizeDocumentWithProof(didDoc, didDocProof)
}

// EcdsaSecp256k1RecoverySignature2020Normalize normalizes DID Document in accordance with
// the Identity Foundation draft on EcdsaSecp256k1RecoverySignature2020
// Read more: https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/
func EcdsaSecp256k1RecoverySignature2020Normalize(didDoc *types.DidDocument, didDocProof *types.DocumentProof) ([]byte, error) {
	return normalizeDocumentWithProof(didDoc, didDocProof)
}

// BbsBlsSignature2020Normalize normalizes the DID Document for the
// BbsBlsSignature2020 signature type
// Read more: https://identity.foundation/bbs-signature/draft-irtf-cfrg-bbs-signatures.html
func BbsBlsSignature2020Normalize(didDoc *types.DidDocument, didDocProof *types.DocumentProof) ([]byte, error) {
	return normalizeDocumentWithProof(didDoc, didDocProof)
}

// EcdsaSecp256k1Signature2019Normalize normalizes the DID Document for the
// EcdsaSecp256k1Signature2019 signature type
// Read more: https://w3c-ccg.github.io/lds-ecdsa-secp256k1-2019/
func EcdsaSecp256k1Signature2019Normalize(didDoc *types.DidDocument, didDocProof *types.DocumentProof) ([]byte, error) {
	return normalizeDocumentWithProof(didDoc, didDocProof)
}
