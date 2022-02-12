package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateDID = "update_did"

var _ sdk.Msg = &MsgUpdateDID{}

func NewMsgUpdateDID(creator string, didDocString *DidDocStructUpdateDID, signatures []*SignInfo) *MsgUpdateDID {
	return &MsgUpdateDID{
		Creator:      creator,
		DidDocString: didDocString,
		Signatures:   signatures,
	}
}

func (msg *MsgUpdateDID) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDID) Type() string {
	return TypeMsgUpdateDID
}

func (msg *MsgUpdateDID) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDID) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *DidDocStructUpdateDID) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDID) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *DidDocStructUpdateDID) GetSigners() []Signer {
	// TODO: Implemt this when working on the updated DidDoc standard
	// if len(msg.Controller) > 0 {
	// 	result := make([]Signer, len(msg.Controller))

	// 	for i, controller := range msg.Controller {
	// 		if controller == msg.Id {
	// 			result[i] = Signer{
	// 				Signer:             controller,
	// 				Authentication:     msg.Authentication,
	// 				VerificationMethod: msg.VerificationMethod,
	// 			}
	// 		} else {
	// 			result[i] = Signer{
	// 				Signer: controller,
	// 			}
	// 		}
	// 	}

	// 	return result
	// }

	if len(msg.Authentication) > 0 {
		return []Signer{
			{
				Signer:          msg.Id,
				Authentication:  msg.Authentication,
				PublicKeyStruct: msg.PublicKey,
			},
		}
	}

	return []Signer{}
}
