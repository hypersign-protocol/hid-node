package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterCredentialStatus = "register_credential_status"

var _ sdk.Msg = &MsgRegisterCredentialStatus{}

func NewMsgRegisterCredentialStatus(creator string, credentialStatus *CredentialStatus) *MsgRegisterCredentialStatus {
	return &MsgRegisterCredentialStatus{
		Creator:          creator,
		CredentialStatus: credentialStatus,
	}
}

func (msg *MsgRegisterCredentialStatus) Route() string {
	return RouterKey
}

func (msg *MsgRegisterCredentialStatus) Type() string {
	return TypeMsgCreateSchema
}

func (msg *MsgRegisterCredentialStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterCredentialStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *CredentialStatus) GetSignBytes() []byte {
	return ModuleCdc.MustMarshal(msg)
}

func (msg *MsgRegisterCredentialStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
