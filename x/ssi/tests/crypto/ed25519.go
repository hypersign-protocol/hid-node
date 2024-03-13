package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
)

func GenerateEd25519KeyPair() *Ed25519KeyPair {
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

	return &Ed25519KeyPair{
		Type:       types.Ed25519VerificationKey2020,
		PublicKey:  publicKeyMultibase,
		PrivateKey: privKeyBase64String,
	}
}
