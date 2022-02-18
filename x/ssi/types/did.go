package types

import "github.com/multiformats/go-multibase"

func (v VerificationMethod) GetPublicKey() ([]byte, error) {
	if len(v.PublicKeyMultibase) > 0 {
		_, key, err := multibase.Decode(v.PublicKeyMultibase)
		if err != nil {
			return nil, ErrInvalidPublicKey.Wrapf("Cannot decode verification method '%s' public key", v.Id)
		}
		return key, nil
	}

	return nil, ErrInvalidPublicKey.Wrapf("verification method '%s' public key not found", v.Id)
}
