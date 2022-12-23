package verification

var ServiceTypes = []string{
	"LinkedDomains",
}

const DidMethod string = "hid"

// Acceptable Credential Status.
// Ref: https://github.com/hypersign-protocol/hid-node/discussions/141#discussioncomment-2825349
const ClaimStatus_live      = "Live"
const ClaimStatus_suspended = "Suspended"
const ClaimStatus_revoked   = "Revoked"
const ClaimStatus_expired   = "Expired"

func GetAcceptedCredentialClaimStatuses() []string {
	return []string{
		ClaimStatus_live,
		ClaimStatus_suspended,
		ClaimStatus_revoked,
		ClaimStatus_expired,
	}	
}
