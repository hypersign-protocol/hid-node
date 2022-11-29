package verification

import (
	"fmt"
	"regexp"
	"time"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func VerifyCredentialHash(credHash string) bool {
	var matchFound bool = false
	var supportedCredentialHash map[string]string = map[string]string{
		"sha256": "[a-f0-9]{64}",
	}

	for _, regexPattern := range supportedCredentialHash {
		matchFound, _ = regexp.MatchString(regexPattern, credHash)
		if matchFound {
			return true
		}
	}

	return matchFound
}

func VerifyCredentialStatusDates(issuanceDate time.Time, expirationDate time.Time) error {
	var dateDiff int64 = int64(expirationDate.Sub(issuanceDate)) / 1e9 // converting nanoseconds to seconds
	if dateDiff < 0 {
		return sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("expiration date %s cannot be less than issuance date %s", expirationDate, issuanceDate))
	}

	return nil
}

func VerifyCredentialProofDates(credProof *types.CredentialProof, credRegistration bool) error {
	var dateDiff int64

	proofCreatedDate := credProof.GetCreated()
	proofCreatedDateParsed, err := time.Parse(time.RFC3339, proofCreatedDate)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid created date format: %s", proofCreatedDate))
	}

	proofUpdatedDate := credProof.GetUpdated()
	proofUpdatedDateParsed, err := time.Parse(time.RFC3339, proofUpdatedDate)
	if err != nil {
		return sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("invalid created date format: %s", proofUpdatedDate))
	}

	// If credRegistration is True, check for equity of updated and created dates will proceeed
	// Else, check for updated date being greater than created date will proceeed
	if credRegistration {
		if !proofUpdatedDateParsed.Equal(proofCreatedDateParsed) {
			return sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("updated date %s should be similar to created date %s", proofUpdatedDate, proofCreatedDate))
		}
	} else {
		dateDiff = int64(proofUpdatedDateParsed.Sub(proofCreatedDateParsed)) / 1e9 // converting nanoseconds to seconds
		if dateDiff <= 0 {
			return sdkerrors.Wrapf(types.ErrInvalidDate, fmt.Sprintf("update date %s cannot be less than or equal to created date %s in case of credential status update", proofUpdatedDate, proofCreatedDate))
		}
	}

	return nil
}
