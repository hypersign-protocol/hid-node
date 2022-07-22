package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateSchema = "create_schema"

var _ sdk.Msg = &MsgCreateSchema{}

func NewMsgCreateSchema(creator string, schemaDoc *SchemaDocument, schemaProof *SchemaProof) *MsgCreateSchema {
	return &MsgCreateSchema{
		Creator:     creator,
		SchemaDoc:   schemaDoc,
		SchemaProof: schemaProof,
	}
}

func (msg *MsgCreateSchema) Route() string {
	return RouterKey
}

func (msg *MsgCreateSchema) Type() string {
	return TypeMsgCreateSchema
}

func (msg *MsgCreateSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *SchemaDocument) GetSignBytes() []byte {
	return ModuleCdc.MustMarshal(msg)
}

func (msg *MsgCreateSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
