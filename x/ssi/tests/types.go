package tests

import (
	"crypto/ed25519"

	secp256k1 "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// Structs

type secp256k1KeyPair struct {
	publicKey  string
	privateKey *secp256k1.PrivateKey
}

type ed25519KeyPair struct {
	publicKey  string
	privateKey ed25519.PrivateKey
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
