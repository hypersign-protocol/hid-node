package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ssi module sentinel errors
var (
	ErrDidDocExists                    = sdkerrors.Register(ModuleName, 101, "didDoc already exists")
	ErrInvalidDidDoc                   = sdkerrors.Register(ModuleName, 102, "didDoc is invalid")
	ErrVerificationMethodNotFound      = sdkerrors.Register(ModuleName, 103, "verification method not found")
	ErrInvalidSignature                = sdkerrors.Register(ModuleName, 104, "invalid signature detected")
	ErrDidDocNotFound                  = sdkerrors.Register(ModuleName, 105, "didDoc not found")
	ErrSchemaExists                    = sdkerrors.Register(ModuleName, 106, "schema already exists")
	ErrInvalidSchemaID                 = sdkerrors.Register(ModuleName, 107, "invalid schema Id")
	ErrUnexpectedDidVersion            = sdkerrors.Register(ModuleName, 108, "unexpected DID version")
	ErrDidDocDeactivated               = sdkerrors.Register(ModuleName, 109, "didDoc is deactivated")
	ErrInvalidCredentialStatus         = sdkerrors.Register(ModuleName, 110, "invalid Credential Status")
	ErrInvalidDate                     = sdkerrors.Register(ModuleName, 111, "invalid Date")
	ErrInvalidCredentialField          = sdkerrors.Register(ModuleName, 112, "invalid Credential Field")
	ErrInvalidCredentialMerkleRootHash = sdkerrors.Register(ModuleName, 113, "invalid Credential merkle root hash")
	ErrInvalidClientSpecType           = sdkerrors.Register(ModuleName, 114, "invalid Client Spec Type")
	ErrCredentialStatusNotFound        = sdkerrors.Register(ModuleName, 115, "credential status document not found")
	ErrCredentialStatusExists          = sdkerrors.Register(ModuleName, 116, "credential status document already exists")
	ErrInvalidCredentialStatusID       = sdkerrors.Register(ModuleName, 117, "invalid credential status Id")
	ErrInvalidProof                    = sdkerrors.Register(ModuleName, 118, "invalid document proof")
	ErrInvalidCredentialSchema         = sdkerrors.Register(ModuleName, 119, "invalid credential schema")
)
