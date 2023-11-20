package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterDID{}, "ssi/RegisterDID", nil)
	cdc.RegisterConcrete(&MsgUpdateDID{}, "ssi/UpdateDID", nil)
	cdc.RegisterConcrete(&MsgRegisterCredentialSchema{}, "ssi/RegisterCredentialSchema", nil)
	cdc.RegisterConcrete(&MsgDeactivateDID{}, "ssi/DeactivateDID", nil)
	cdc.RegisterConcrete(&MsgRegisterCredentialStatus{}, "ssi/RegisterCredentialStatus", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterDID{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgUpdateDID{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterCredentialSchema{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgDeactivateDID{},
	)
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgRegisterCredentialStatus{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
