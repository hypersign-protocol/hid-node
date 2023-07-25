package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	DefaultCreateDIDFee                = sdk.NewInt64Coin("uhid", 4000)
	DefaultUpdateDIDFee                = sdk.NewInt64Coin("uhid", 1000)
	DefaultDeactivateDIDFee            = sdk.NewInt64Coin("uhid", 1000)
	DefaultCreateSchemaFee             = sdk.NewInt64Coin("uhid", 2000)
	DefaultRegisterCredentialStatusFee = sdk.NewInt64Coin("uhid", 2000)
)

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CreateDidFee:                &DefaultCreateDIDFee,
		UpdateDidFee:                &DefaultUpdateDIDFee,
		DeactivateDidFee:            &DefaultDeactivateDIDFee,
		CreateSchemaFee:             &DefaultCreateSchemaFee,
		RegisterCredentialStatusFee: &DefaultRegisterCredentialStatusFee,
	}
}

func (gs GenesisState) Validate() error {
	return nil
}
