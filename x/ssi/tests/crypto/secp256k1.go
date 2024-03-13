package crypto

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"

	ethercrypto "github.com/ethereum/go-ethereum/crypto"
)

func GenerateSecp256k1KeyPair() *Secp256k1Pair {
	privateKey := secp256k1.GenPrivKey()

	publicKey := privateKey.PubKey().Bytes()

	publicKeyMultibase, err := multibase.Encode(multibase.Base58BTC, publicKey)
	if err != nil {
		panic("Error while encoding multibase string")
	}
	return &Secp256k1Pair{
		Type: types.EcdsaSecp256k1VerificationKey2019,
		PublicKey:  publicKeyMultibase,
		PrivateKey: base64.StdEncoding.EncodeToString(privateKey),
	}
}

func GenerateSecp256k1RecoveryKeyPair() *Secp256k1RecoveryPair {
	privateKeyObj := secp256k1.GenPrivKey()
	privateKey := privateKeyObj.Bytes()

	publicKeyCompressed := privateKeyObj.PubKey().Bytes()

	publicKeyUncompressed, err := ethercrypto.DecompressPubkey(publicKeyCompressed)
	if err != nil {
		panic(err)
	}
	ethereumAddress := ethercrypto.PubkeyToAddress(*publicKeyUncompressed).Hex()

	return &Secp256k1RecoveryPair{
		Type: types.EcdsaSecp256k1RecoveryMethod2020,
		PublicKey:  hex.EncodeToString(publicKeyCompressed),
		PrivateKey: hex.EncodeToString(privateKey),
		OptionalID: ethereumAddress,
	}
}
