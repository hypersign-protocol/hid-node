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
