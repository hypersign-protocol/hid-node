package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/ssi module sentinel errors
var (
	ErrDidDocExists                    = errors.Register(ModuleName, 101, "didDoc already exists")
	ErrInvalidDidDoc                   = errors.Register(ModuleName, 102, "didDoc is invalid")
	ErrVerificationMethodNotFound      = errors.Register(ModuleName, 103, "verification method not found")
	ErrInvalidSignature                = errors.Register(ModuleName, 104, "invalid signature detected")
	ErrDidDocNotFound                  = errors.Register(ModuleName, 105, "didDoc not found")
	ErrSchemaExists                    = errors.Register(ModuleName, 106, "schema already exists")
	ErrInvalidSchemaID                 = errors.Register(ModuleName, 107, "invalid schema Id")
	ErrUnexpectedDidVersion            = errors.Register(ModuleName, 108, "unexpected DID version")
	ErrDidDocDeactivated               = errors.Register(ModuleName, 109, "didDoc is deactivated")
	ErrInvalidCredentialStatus         = errors.Register(ModuleName, 110, "invalid Credential Status")
	ErrInvalidDate                     = errors.Register(ModuleName, 111, "invalid Date")
	ErrInvalidCredentialField          = errors.Register(ModuleName, 112, "invalid Credential Field")
	ErrInvalidCredentialMerkleRootHash = errors.Register(ModuleName, 113, "invalid Credential merkle root hash")
	ErrInvalidClientSpecType           = errors.Register(ModuleName, 114, "invalid Client Spec Type")
	ErrCredentialStatusNotFound        = errors.Register(ModuleName, 115, "credential status document not found")
	ErrCredentialStatusExists          = errors.Register(ModuleName, 116, "credential status document already exists")
	ErrInvalidCredentialStatusID       = errors.Register(ModuleName, 117, "invalid credential status Id")
	ErrInvalidProof                    = errors.Register(ModuleName, 118, "invalid document proof")
	ErrInvalidCredentialSchema         = errors.Register(ModuleName, 119, "invalid credential schema")
)
