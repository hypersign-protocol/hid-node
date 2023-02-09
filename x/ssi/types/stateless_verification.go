package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/common"
)

func didDocumentStatelessVerification(didDoc *Did) error {
	didId := didDoc.GetId()
	if didId == "" {
		return ErrBadRequestIsRequired.Wrap("didDoc Id is required")
	}

	// Verification Method of Type EcdsaSecp256k1RecoveryMethod2020 should only have
	// `blockchainAccountId` field populated
	verificationMethods := didDoc.GetVerificationMethod()
	for i := 0; i < len(verificationMethods); i++ {
		switch verificationMethods[i].Type {
		case common.EcdsaSecp256k1RecoveryMethod2020:
			if verificationMethods[i].GetBlockchainAccountId() == "" {
				return sdkerrors.Wrapf(
					ErrBadRequestInvalidVerMethod,
					"blockchainAccountId cannot be empty for verification method %s as it is of type %s",
					verificationMethods[i].Id,
					verificationMethods[i].Type,
				)
			}
			if verificationMethods[i].GetPublicKeyMultibase() != "" {
				return sdkerrors.Wrapf(
					ErrBadRequestInvalidVerMethod,
					"publicKeyMultibase should be empty for verification method %s as it is type %s",
					verificationMethods[i].Id,
					verificationMethods[i].Type,
				)
			}

		default:
			if verificationMethods[i].GetBlockchainAccountId() != "" {
				return sdkerrors.Wrapf(
					ErrBadRequestInvalidVerMethod,
					"blockchainAccountId should be empty for verification method %s as it is of type %s",
					verificationMethods[i].Id,
					verificationMethods[i].Type,
				)
			}
			if verificationMethods[i].GetPublicKeyMultibase() == "" {
				return sdkerrors.Wrapf(
					ErrBadRequestInvalidVerMethod,
					"publicKeyMultibase cannot be empty for verification method %s as it is type %s",
					verificationMethods[i].Id,
					verificationMethods[i].Type,
				)
			}
		}
	}

	return nil
}
