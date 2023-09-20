package ldcontext

import (
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func NormalizeByVerificationMethodType(didDoc *types.Did, vmType string) ([]byte, error) {
	switch vmType {
	case types.Ed25519VerificationKey2020:
		didDocBytes, err := EdDSACryptoSuite2020Canonize(didDoc)
		if err != nil {
			return nil, err
		}
		return didDocBytes, nil
	default:
		return didDoc.GetSignBytes(), nil
	}
}
