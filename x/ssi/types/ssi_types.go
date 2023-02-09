package types

import (
	proto "github.com/gogo/protobuf/proto"
)

type (
	SsiMsg interface {
		proto.Message
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

// Struct catering to supported Client Spec's required inputs
type ClientSpecOpts struct {
	SSIDoc   SsiMsg
	SignerAddress string
}

// Cosmos ADR SignDoc Struct Definitions
type (
	Val struct {
		Data   string `json:"data"`
		Signer string `json:"signer"`
	}

	Msg struct {
		Type  string `json:"type"`
		Value Val    `json:"value"`
	}

	Fees struct {
		Amount []string `json:"amount"`
		Gas    string   `json:"gas"`
	}

	SignDoc struct {
		AccountNumber string `json:"account_number"`
		ChainId       string `json:"chain_id"`
		Fee           Fees   `json:"fee"`
		Memo          string `json:"memo"`
		Msgs          []Msg  `json:"msgs"`
		Sequence      string `json:"sequence"`
	}
)
