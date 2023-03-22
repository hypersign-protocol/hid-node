package types

import (
	fmt "fmt"
	"strings"

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
		ClientSpec          *ClientSpec
	}
)

func CreateExtendedVerificationMethod(vm *VerificationMethod, signInfo *SignInfo) *ExtendedVerificationMethod {
	extendedVm := &ExtendedVerificationMethod{
		Id:                  vm.Id,
		Type:                vm.Type,
		Controller:          vm.Controller,
		PublicKeyMultibase:  vm.PublicKeyMultibase,
		BlockchainAccountId: vm.BlockchainAccountId,
		Signature:           signInfo.Signature,
	}

	if signInfo.ClientSpec != nil {
		extendedVm.ClientSpec = signInfo.ClientSpec
	}

	return extendedVm
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

// CAIP-10 Blockchain Account Id
type BlockchainId struct {
	CAIP10Prefix      string
	ChainId           string
	BlockchainAddress string
}

func NewBlockchainId(blockchainAccountId string) (*BlockchainId, error) {
	segments := strings.Split(blockchainAccountId, ":")

	// Validate blockchainAccountId segments
	if len(segments) != 3 {
		return nil, fmt.Errorf(
			"invalid CAIP-10 format for blockchainAccountId '%v'. Please refer: https://github.com/ChainAgnostic/CAIPs/blob/master/CAIPs/caip-10.md",
			blockchainAccountId,
		)
	}

	return &BlockchainId{
		CAIP10Prefix:      segments[0],
		ChainId:           segments[1],
		BlockchainAddress: segments[2],
	}, nil
}

func (bid *BlockchainId) ValidateSupportedCAIP10Prefix() error {
	if !FindInSlice(SupportedCAIP10Prefixes, bid.CAIP10Prefix) {
		return fmt.Errorf(
			"unsupported CAIP-10 prefix: '%v', supported CAIP-10 prefixes are %v",
			bid.CAIP10Prefix,
			SupportedCAIP10Prefixes,
		)
	}
	return nil
}

func (bid *BlockchainId) ValidateSupportChainId() error {
	supportedChainIds := SupportedCAIP10PrefixChainIdsMap[bid.CAIP10Prefix]

	if !FindInSlice(supportedChainIds, bid.ChainId) {
		return fmt.Errorf(
			"unsupported CAIP-10 chain-id: '%v', supported CAIP-10 chain-ids are %v",
			bid.ChainId,
			supportedChainIds,
		)
	}

	return nil
}

func (bid *BlockchainId) ValidateSupportedBech32Prefix() error {
	extractedBech32Prefix := strings.Split(bid.BlockchainAddress, "1")[0]

	supportedBech32Prefix, supported := CosmosCAIP10ChainIdBech32PrefixMap[bid.ChainId]
	if !supported {
		return fmt.Errorf("chain-id %v is not supported", bid.ChainId)
	}

	if supportedBech32Prefix != extractedBech32Prefix {
		return fmt.Errorf("invalid bech32 prefix for blockchain address: %v", bid.BlockchainAddress)
	}

	return nil
}
