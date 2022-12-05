package tests

import (
	"crypto/ed25519"

	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// Structs

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
