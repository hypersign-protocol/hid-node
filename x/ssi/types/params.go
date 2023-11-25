package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	DefaultRegisterDIDFee              = sdk.NewInt64Coin("uhid", 4000)
	DefaultUpdateDIDFee                = sdk.NewInt64Coin("uhid", 1000)
	DefaultDeactivateDIDFee            = sdk.NewInt64Coin("uhid", 1000)
	DefaultRegisterCredentialSchemaFee = sdk.NewInt64Coin("uhid", 2000)
	DefaultUpdateCredentialSchemaFee   = sdk.NewInt64Coin("uhid", 2000)
	DefaultRegisterCredentialStatusFee = sdk.NewInt64Coin("uhid", 2000)
	DefaultUpdateCredentialStatusFee   = sdk.NewInt64Coin("uhid", 2000)
)

func DefaultParams() *Params {
	return &Params{
		RegisterDidFee:              &DefaultRegisterDIDFee,
		UpdateDidFee:                &DefaultUpdateDIDFee,
		DeactivateDidFee:            &DefaultDeactivateDIDFee,
		RegisterCredentialSchemaFee: &DefaultRegisterCredentialSchemaFee,
		UpdateCredentialSchemaFee:   &DefaultUpdateCredentialSchemaFee,
		RegisterCredentialStatusFee: &DefaultRegisterCredentialStatusFee,
		UpdateCredentialStatusFee:   &DefaultUpdateCredentialStatusFee,
	}
}

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(ParamStoreKeyRegisterDidFee, sdk.Coin{}, validateFeeParams),
		paramtypes.NewParamSetPair(ParamStoreKeyUpdateDidFee, sdk.Coin{}, validateFeeParams),
		paramtypes.NewParamSetPair(ParamStoreKeyDeactivateDidFee, sdk.Coin{}, validateFeeParams),
		paramtypes.NewParamSetPair(ParamStoreKeyRegisterCredentialSchemaFee, sdk.Coin{}, validateFeeParams),
		paramtypes.NewParamSetPair(ParamStoreKeyUpdateCredentialSchemaFee, sdk.Coin{}, validateFeeParams),
		paramtypes.NewParamSetPair(ParamStoreKeyRegisterCredentialStatusFee, sdk.Coin{}, validateFeeParams),
		paramtypes.NewParamSetPair(ParamStoreKeyUpdateCredentialStatusFee, sdk.Coin{}, validateFeeParams),
	)
}

func validateFeeParams(i interface{}) error {
	v, ok := i.(sdk.Coin)

	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.Denom != "uhid" {
		return fmt.Errorf("fee param denom must be 'uhid', got %v", v.Denom)
	}

	return nil
}
