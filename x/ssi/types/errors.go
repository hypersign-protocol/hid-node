package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ssi module sentinel errors
var (
	ErrBadRequestIsRequired       = sdkerrors.Register(ModuleName, 101, "is required")
	ErrBadRequestIsNotDid         = sdkerrors.Register(ModuleName, 102, "is an invalid Hypersign DID format")
	ErrDidDocExists               = sdkerrors.Register(ModuleName, 103, "DidDoc already exists")
	ErrInvalidDidDoc              = sdkerrors.Register(ModuleName, 104, "DidDoc is invalid")
	ErrVerificationMethodNotFound = sdkerrors.Register(ModuleName, 105, "verification method not found")
	ErrInvalidPublicKey           = sdkerrors.Register(ModuleName, 106, "invalid public key")
	ErrInvalidSignature           = sdkerrors.Register(ModuleName, 107, "invalid signature detected")
	ErrDidDocNotFound             = sdkerrors.Register(ModuleName, 108, "DID Doc not found")
	ErrSchemaExists               = sdkerrors.Register(ModuleName, 109, "Schema already exists")
	ErrInvalidSchemaID            = sdkerrors.Register(ModuleName, 110, "Invalid schema Id")
	ErrBadRequestInvalidVerMethod = sdkerrors.Register(ModuleName, 111, "Invalid verification method")
	ErrInvalidService             = sdkerrors.Register(ModuleName, 112, "Invalid Service")
	ErrUnexpectedDidVersion       = sdkerrors.Register(ModuleName, 113, "Unexpected DID version")
	ErrDidDocDeactivated          = sdkerrors.Register(ModuleName, 114, "DID Document is deactivated")
	ErrInvalidDidElements         = sdkerrors.Register(ModuleName, 115, "Invalid DID elements")
	ErrInvalidDidMethod           = sdkerrors.Register(ModuleName, 116, "Invalid DID method")
)
