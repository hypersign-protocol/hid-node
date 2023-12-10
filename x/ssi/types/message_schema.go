package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"cosmossdk.io/errors"
)

const TypeMsgRegisterCredentialSchema = "register_credential_schema"

var _ sdk.Msg = &MsgRegisterCredentialSchema{}

func NewMsgRegisterSchema(
	schemaDoc *CredentialSchemaDocument,
	schemaProof *DocumentProof,
	clientSpecType ClientSpecType,
	txAuthor string,
) *MsgRegisterCredentialSchema {
	return &MsgRegisterCredentialSchema{
		CredentialSchemaDocument: schemaDoc,
		CredentialSchemaProof:    schemaProof,
		TxAuthor:                 txAuthor,
	}
}

func (msg *MsgRegisterCredentialSchema) Route() string {
	return RouterKey
}

func (msg *MsgRegisterCredentialSchema) Type() string {
	return TypeMsgRegisterCredentialSchema
}

func (msg *MsgRegisterCredentialSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.TxAuthor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterCredentialSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *CredentialSchemaDocument) GetSignBytes() []byte {
	return ModuleCdc.MustMarshal(msg)
}

func (msg *MsgRegisterCredentialSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.TxAuthor)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// Update Credential Schema

const TypeMsgUpdateCredentialSchema = "update_credential_schema"

var _ sdk.Msg = &MsgUpdateCredentialSchema{}

func (msg *MsgUpdateCredentialSchema) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCredentialSchema) Type() string {
	return TypeMsgUpdateCredentialSchema
}

func (msg *MsgUpdateCredentialSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.TxAuthor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCredentialSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCredentialSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.TxAuthor)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
