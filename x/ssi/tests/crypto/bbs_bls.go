package crypto

import (
	"crypto/sha256"
	"encoding/base64"

	bbs "github.com/hyperledger/aries-framework-go/component/kmscrypto/crypto/primitive/bbs12381g2pub"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
)

func GenerateBbsBlsKeyPair() *BbsBlsKeyPair {
	pubKey, privKey, err := bbs.GenerateKeyPair(sha256.New, nil)
	if err != nil {
		panic(err)
	}

	// Convert Public Key Object to Multibase
	pubKeyBytes, err := pubKey.Marshal()
	if err != nil {
		panic(err)
	}

	publicKeyMultibase, err := multibase.Encode(multibase.Base58BTC, pubKeyBytes)
	if err != nil {
		panic(err)
	}

	// Convert Private Object to Bytes
	privKeyBytes, err := privKey.Marshal()
	if err != nil {
		panic(err)
	}

	return &BbsBlsKeyPair{
		Type:       types.Bls12381G2Key2020,
		PublicKey:  publicKeyMultibase,
		PrivateKey: base64.StdEncoding.EncodeToString(privKeyBytes),
	}
}
