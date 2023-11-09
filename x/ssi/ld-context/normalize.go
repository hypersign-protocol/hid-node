package ldcontext

import (
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// NormalizeByVerificationMethodType canonizes DID Document based on the input Verification
// Method type
func NormalizeByVerificationMethodType(didDoc *types.Did, vmType string) ([]byte, error) {
	switch vmType {
	case types.Ed25519VerificationKey2020:
		didDocBytes, err := EdDSACryptoSuite2020Canonize(didDoc)
		if err != nil {
			return nil, err
		}
		return didDocBytes, nil
	case types.EcdsaSecp256k1RecoveryMethod2020:
		didDocBytes, err := EcdsaSecp256k1RecoverySignature2020Canonize(didDoc)
		if err != nil {
			return nil, err
		}
		return didDocBytes, nil
	case types.Bls12381G2Key2020:
		didDocBytes, err := BbsBlsSignature2020Canonize(didDoc)
		if err != nil {
			return nil, err
		}
		return didDocBytes, nil
	case types.EcdsaSecp256k1VerificationKey2019:
		didDocBytes, err := EcdsaSecp256k1Signature2019Canonize(didDoc)
		if err != nil {
			return nil, err
		}
		return didDocBytes, nil
	default:
		return didDoc.GetSignBytes(), nil
	}
}
