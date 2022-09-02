package types

type (
	IdentityMsg interface {
		GetSigners() []Signer
		GetSignBytes() []byte
	}

	Signer struct {
		Signer               string
		Authentication       []string
		AssertionMethod      []string
		VerificationMethod   []*VerificationMethod
		KeyAgreement         []string
		CapabilityInvocation []string
		CapabilityDelegation []string
	}

	ValidDid struct {
		DidId   string
		IsValid bool
	}
)
