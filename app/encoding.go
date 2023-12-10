package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry codecTypes.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// MakeEncodingConfig creates an EncodingConfig for an amino based test configuration.
// NewEncodingConfig creates an EncodingConfig instance.
func NewEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := codecTypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)

	encodingConfig := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}

	return encodingConfig
}

func MakeEncodingConfig() EncodingConfig {
	encodingConfig := NewEncodingConfig()

	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}
