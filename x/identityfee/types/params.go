package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(ParamStoreKeyCreateDidFee, sdk.Coin{}, validateFeeParams),
		paramtypes.NewParamSetPair(ParamStoreKeyUpdateDidFee, sdk.Coin{}, validateFeeParams),
		paramtypes.NewParamSetPair(ParamStoreKeyDeactivateDidFee, sdk.Coin{}, validateFeeParams),
		paramtypes.NewParamSetPair(ParamStoreKeyCreateSchemaFee, sdk.Coin{}, validateFeeParams),
		paramtypes.NewParamSetPair(ParamStoreKeyRegisterCredentialStatusFee, sdk.Coin{}, validateFeeParams),
	)
}

func validateFeeParams(i interface{}) error {
	return nil
}