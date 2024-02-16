package cmd

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/ripemd160" //nolint: staticcheck

	bech32 "github.com/cosmos/cosmos-sdk/types/bech32"
)

// publicKeyToBech32Address converts publicKey byteArray to Bech32 encoded blockchain address
func publicKeyToBech32Address(addressPrefix string, pubKeyBytes []byte) string {
	// Throw error if the length of secp256k1 publicKey is not 33
	if len(pubKeyBytes) != 33 {
		panic(fmt.Sprintf("invalid secp256k1 public key length %v", len(pubKeyBytes)))
	}

	// Hash pubKeyBytes as: RIPEMD160(SHA256(public_key_bytes))
	pubKeySha256Hash := sha256.Sum256(pubKeyBytes)
	ripemd160hash := ripemd160.New()
	ripemd160hash.Write(pubKeySha256Hash[:])
	addressBytes := ripemd160hash.Sum(nil)

	// Convert addressBytes to bech32 encoded address
	address, err := bech32.ConvertAndEncode(addressPrefix, addressBytes)
	if err != nil {
		panic(err)
	}
	return address
}
