package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateDID = "create_did"

var _ sdk.Msg = &MsgCreateDID{}

func NewMsgCreateDID(creator string, did string, didDocString string, createdAt string) *MsgCreateDID {
	return &MsgCreateDID{
		Creator:      creator,
		Did:          did,
		DidDocString: didDocString,
		CreatedAt:    createdAt,
	}
}

func (msg *MsgCreateDID) Route() string {
	return RouterKey
}

func (msg *MsgCreateDID) Type() string {
	return TypeMsgCreateDID
}

func (msg *MsgCreateDID) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDID) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDID) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
