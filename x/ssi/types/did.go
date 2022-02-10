package types

import "github.com/multiformats/go-multibase"

func (p PublicKeyStruct) GetPublicKey() ([]byte, error) {
	if len(p.PublicKeyBase58) > 0 {
		_, key, err := multibase.Decode(p.PublicKeyBase58)
		if err != nil {
			return nil, ErrInvalidPublicKey.Wrapf("Cannot decode verification method '%s' public key", p.Id)
		}
		return key, nil
	}

	return nil, ErrInvalidPublicKey.Wrapf("verification method '%s' public key not found", p.Id)
}
