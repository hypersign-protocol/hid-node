package types

import (
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
)

// Valid Proof Types
const (
	authentication       = "authentication"
	assertionMethod      = "assertionMethod"
	keyAgreement         = "keyAgreement"
	capabilityInvocation = "capabilityInvocation"
	capabilityDelegation = "capabilityDelegation"
)

var validProofPurposes []string = []string{
	authentication,
	assertionMethod,
	keyAgreement,
	capabilityInvocation,
	capabilityDelegation,
}

// Validate Document Proof
func (proof *DocumentProof) Validate() error {
	// Validate Proof Type
	if !utils.FindInSlice(validProofPurposes, proof.ProofPurpose) {
		return invalidProofErrorMsg("invalid proof purpose %v", proof.ProofPurpose)
	}

	// Validate Proof Created
	_, err := time.Parse(time.RFC3339, proof.Created)
	if err != nil {
		return invalidProofErrorMsg("invalid created date format: %v", proof.Created)
	}

	// Validate VerificationMethod
	if len(proof.VerificationMethod) == 0 {
		return invalidProofErrorMsg("'verificationMethod' attribute in document proof cannot be empty")
	}

	if len(proof.ProofValue) == 0 {
		return invalidProofErrorMsg("'proofValue' attribute in document proof cannot be empty")
	}

	return nil
}

func invalidProofErrorMsg(errMsg string, errMsgArgs ...interface{}) error {
	return sdkerrors.Wrapf(ErrInvalidProof, errMsg, errMsgArgs...)
}
