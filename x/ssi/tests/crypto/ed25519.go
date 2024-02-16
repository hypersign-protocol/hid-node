package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"

	"github.com/multiformats/go-multibase"
)

func GenerateEd25519KeyPair() *KeyPair {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	var publicKeyWithHeader []byte
	publicKeyWithHeader = append(publicKeyWithHeader, append([]byte{0xed, 0x01}, publicKey...)...)

	publicKeyMultibase, err := multibase.Encode(multibase.Base58BTC, publicKeyWithHeader)
	if err != nil {
		panic("Error while encoding multibase string")
	}

	privKeyBase64String := base64.StdEncoding.EncodeToString(privateKey)

	return &KeyPair{
		Type:       Ed25519KeyPair,
		PublicKey:  publicKeyMultibase,
		PrivateKey: privKeyBase64String,
	}
}
