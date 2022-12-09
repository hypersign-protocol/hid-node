package tests

import (
	"crypto/ed25519"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	secp256k1 "github.com/tendermint/tendermint/crypto/secp256k1"
)

// Structs

type ed25519KeyPair struct {
	publicKey  string
	privateKey ed25519.PrivateKey
}

type secp256k1KeyPair struct {
	publicKey  string
	privateKey *secp256k1.PrivKey
}

type DidRpcElements struct {
	DidDocument *types.Did
	Signatures  []*types.SignInfo
	Creator     string
}

type SchemaRpcElements struct {
	SchemaDocument *types.SchemaDocument
	SchemaProof    *types.SchemaProof
	Creator        string
}

type CredRpcElements struct {
	Status  *types.CredentialStatus
	Proof   *types.CredentialProof
	Creator string
}

type DidSigningElements struct {
	keyPair interface{}
	vmId    string
}

// Interfaces

// An interface to support multiple Key Pair Structs
type GenericKeyPair = interface{}
