package types

import (
	"github.com/multiformats/go-multibase"
)

func (v VerificationMethod) GetPublicKeyAndVerificationMethodType() (string, string, error) {
	var publicKey string = v.PublicKeyMultibase
	var verificationMethodType string = v.Type

	if len(publicKey) == 0 {
		return "", "", ErrInvalidPublicKey.Wrapf("verification method '%s' public key not found", v.Id)
	}

	if publicKey[0] != 'z' {
		return "", "", ErrInvalidPublicKey.Wrapf(
			"public key is expected to be in multibase base58btc encoding, recieved public key: %s",
			publicKey,
		)
	}

	// Check if the public key is decoded successfully
	_, _, err := multibase.Decode(publicKey)
	if err != nil {
		panic(err)
	}

	return publicKey, verificationMethodType, nil
}
