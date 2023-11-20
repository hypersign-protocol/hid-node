package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterCredentialStatus = "register_credential_status"

var _ sdk.Msg = &MsgRegisterCredentialStatus{}

func (msg *MsgUpdateCredentialStatus) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCredentialStatus) Type() string {
	return TypeMsgRegisterCredentialSchema
}

func (msg *MsgUpdateCredentialStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.TxAuthor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCredentialStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCredentialStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.TxAuthor)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid transaction author's address (%s)", err)
	}
	return nil
}

// Update Credential Status

const TypeMsgUpdateCredentialStatus = "update_credential_status"

var _ sdk.Msg = &MsgRegisterCredentialStatus{}

func NewMsgRegisterCredentialStatus(
	credentialStatusDocument *CredentialStatusDocument,
	credentialStatusProof *DocumentProof,
	txAuthor string,
) *MsgRegisterCredentialStatus {
	return &MsgRegisterCredentialStatus{
		CredentialStatusDocument: credentialStatusDocument,
		CredentialStatusProof:    credentialStatusProof,
		TxAuthor:                 txAuthor,
	}
}

func (msg *MsgRegisterCredentialStatus) Route() string {
	return RouterKey
}

func (msg *MsgRegisterCredentialStatus) Type() string {
	return TypeMsgRegisterCredentialSchema
}

func (msg *MsgRegisterCredentialStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.TxAuthor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterCredentialStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *CredentialStatusDocument) GetSignBytes() []byte {
	return ModuleCdc.MustMarshal(msg)
}

func (msg *MsgRegisterCredentialStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.TxAuthor)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid transaction author's address (%s)", err)
	}
	return nil
}
