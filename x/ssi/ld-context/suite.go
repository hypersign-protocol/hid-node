package ldcontext

import (
	"crypto/sha256"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// EdDSACryptoSuite2020Canonize canonizes DID Document in accordance with
// EdDSA Cryptosuite v2020 (https://www.w3.org/community/reports/credentials/CG-FINAL-di-eddsa-2020-20220724/)
func EdDSACryptoSuite2020Canonize(didDoc *types.Did) ([]byte, error) {
	jsonLdDid := NewJsonLdDid(didDoc)
	canonizedDidDocument, err := jsonLdDid.NormalizeWithURDNA2015()
	if err != nil {
		return nil, err
	}
	canonizedDidDocumentHash := sha256.Sum256([]byte(canonizedDidDocument))
	return canonizedDidDocumentHash[:], nil
}

// EcdsaSecp256k1RecoverySignature2020Canonize canonizes DID Document in accordance with
// the Identity Foundation draft on EcdsaSecp256k1RecoverySignature2020 
// Read more: https://identity.foundation/EcdsaSecp256k1RecoverySignature2020/
// LD Context: https://ns.did.ai/suites/secp256k1-2020/v1
func EcdsaSecp256k1RecoverySignature2020Canonize(didDoc *types.Did) ([]byte, error) {
	jsonLdDid := NewJsonLdDid(didDoc)
	canonizedDidDocument, err := jsonLdDid.NormalizeWithURDNA2015()
	if err != nil {
		return nil, err
	}

	canonizedDidDocumentHash := sha256.Sum256([]byte(canonizedDidDocument))
	return canonizedDidDocumentHash[:], nil
}

// BbsBlsSignature2020Canonize canonizes the DID Document for the
// BbsBlsSignature2020 signature type
func BbsBlsSignature2020Canonize(didDoc *types.Did) ([]byte, error) {
	jsonLdDid := NewJsonLdDid(didDoc)
	canonizedDidDocument, err := jsonLdDid.NormalizeWithURDNA2015()
	if err != nil {
		return nil, err
	}

	canonizedDidDocumentHash := sha256.Sum256([]byte(canonizedDidDocument))
	return canonizedDidDocumentHash[:], nil
}


// EcdsaSecp256k1Signature2019Canonize canonizes the DID Document for the
// EcdsaSecp256k1Signature2019 signature type
func EcdsaSecp256k1Signature2019Canonize(didDoc *types.Did) ([]byte, error) {
	jsonLdDid := NewJsonLdDid(didDoc)
	canonizedDidDocument, err := jsonLdDid.NormalizeWithURDNA2015()
	if err != nil {
		return nil, err
	}

	canonizedDidDocumentHash := sha256.Sum256([]byte(canonizedDidDocument))
	return canonizedDidDocumentHash[:], nil
}
