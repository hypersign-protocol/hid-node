package crypto

import (
	"encoding/base64"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/multiformats/go-multibase"
)

func GenerateSecp256k1KeyPair() *KeyPair {
	privateKey := secp256k1.GenPrivKey()

	publicKey := privateKey.PubKey().Bytes()

	publicKeyMultibase, err := multibase.Encode(multibase.Base58BTC, publicKey)
	if err != nil {
		panic("Error while encoding multibase string")
	}
	return &KeyPair{
		PublicKey:  publicKeyMultibase,
		PrivateKey: base64.StdEncoding.EncodeToString(privateKey),
	}
}
