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

	ExtendedVerificationMethod struct {
		Id                  string
		Type                string
		Controller          string
		PublicKeyMultibase  string
		BlockchainAccountId string
		Signature           string
	}
)

func CreateExtendedVerificationMethod(vm *VerificationMethod, signature string) *ExtendedVerificationMethod {
	return &ExtendedVerificationMethod{
		Id:                  vm.Id,
		Type:                vm.Type,
		Controller:          vm.Controller,
		PublicKeyMultibase:  vm.PublicKeyMultibase,
		BlockchainAccountId: vm.BlockchainAccountId,
		Signature:           signature,
	}
}

// Struct catering to supported Client Spec's required inputs
type ClientSpecOpts struct {
	ClientSpecType string
	SSIDoc         SsiMsg
	SignerAddress  string
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

// Handle Proof Struct of SSI Docs
type SSIProofInterface interface {
	GetProofValue() string
	GetType() string
	GetVerificationMethod() string
}
