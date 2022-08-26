package verification

var ServiceTypes = []string{
	"LinkedDomains",
}

const DidMethod string = "hid"

// Acceptable Credential Status.
// Ref: https://github.com/hypersign-protocol/hid-node/discussions/141#discussioncomment-2825349
var AcceptedCredStatuses = []string{
	"Live",
	"Suspended",
	"Revoked",
	"Expired",
}
