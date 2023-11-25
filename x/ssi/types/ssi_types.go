package types

import (
	fmt "fmt"
	"strings"

	proto "github.com/gogo/protobuf/proto"
	"github.com/hypersign-protocol/hid-node/x/ssi/utils"
)

type (
	SsiMsg interface {
		proto.Message
		GetId() string
		GetSignBytes() []byte
		GetContext() []string
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
		Proof               *DocumentProof
	}
)

func CreateExtendedVerificationMethod(vm *VerificationMethod, documentProof *DocumentProof) *ExtendedVerificationMethod {
	if vm.Id != documentProof.VerificationMethod {
		panic(fmt.Sprintf(
			"unexpected behaviour: while creating ExtendedVerificationMethod the verification method Id of the VM object %v is different from verification method id of Proof %v",
			vm.Id,
			documentProof.VerificationMethod,
		))
	}

	extendedVm := &ExtendedVerificationMethod{
		Id:                  vm.Id,
		Type:                vm.Type,
		Controller:          vm.Controller,
		PublicKeyMultibase:  vm.PublicKeyMultibase,
		BlockchainAccountId: vm.BlockchainAccountId,
		Proof:               documentProof,
	}

	return extendedVm
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
	if !utils.FindInSlice(SupportedCAIP10Prefixes, bid.CAIP10Prefix) {
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

	if !utils.FindInSlice(supportedChainIds, bid.ChainId) {
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
