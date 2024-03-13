package crypto

import (
	"encoding/hex"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/multiformats/go-multibase"
)

func GenerateBabyJubJubKeyPair() *BabyJubJubKeyPair {
	// Generate Key Pair
	babyJubjubPrivKey := babyjub.NewRandPrivKey()
	babyJubjubPubKey := babyJubjubPrivKey.Public()

	// Prepare Priv key
	var privKeyBytes [32]byte = babyJubjubPrivKey
	var privKeyHex string = hex.EncodeToString(privKeyBytes[:])

	// Prepare Public Key
	var pubKeyHex string = babyJubjubPubKey.Compress().String()

	// Prepare Multibase Public Key
	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		panic(err)
	}
	pubKeyMultibase, err := multibase.Encode(multibase.Base58BTC, pubKeyBytes)
	if err != nil {
		panic(err)
	}

	return &BabyJubJubKeyPair{
		Type:       types.BabyJubJubKey2021,
		PublicKey:  pubKeyMultibase,
		PrivateKey: privKeyHex,
	}
}
