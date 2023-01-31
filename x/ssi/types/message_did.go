package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Returns a list of controllers present in the Did document along with
// their verification methods.
func (msg *Did) GetSigners() []Signer {
	nControllers := len(msg.Controller)

	if nControllers > 0 {
		signers := make([]Signer, nControllers)
		for i, controller := range msg.Controller {
			if controller == msg.Id {
				signers[i] = Signer{
					Signer:               controller,
					Authentication:       msg.Authentication,
					AssertionMethod:      msg.AssertionMethod,
					KeyAgreement:         msg.KeyAgreement,
					CapabilityInvocation: msg.CapabilityInvocation,
					CapabilityDelegation: msg.CapabilityDelegation,
					VerificationMethod:   msg.VerificationMethod,
				}
			} else {
				signers[i] = Signer{
					Signer: controller,
				}
			}
		}
		return signers
	}

	return []Signer{}
}

func (msg *Did) GetSignBytes() []byte {
	return ModuleCdc.MustMarshal(msg)
}

// MsgCreateDID Type Methods

const TypeMsgCreateDID = "create_did"

var _ sdk.Msg = &MsgCreateDID{}

func NewMsgCreateDID(did string, didDocString *Did) *MsgCreateDID {
	return &MsgCreateDID{
		DidDocString: didDocString,
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
	return ModuleCdc.MustMarshal(msg)
}

func (msg *MsgCreateDID) ValidateBasic() error {
	msgDidDocument := msg.GetDidDocString()
	err := didDocumentStatelessVerification(msgDidDocument)
	return err
}

// MsgUpdateDID Type Methods

const TypeMsgUpdateDID = "update_did"

var _ sdk.Msg = &MsgUpdateDID{}

func NewMsgUpdateDID(creator string, didDocString *Did, signatures []*SignInfo) *MsgUpdateDID {
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

func (msg *MsgUpdateDID) ValidateBasic() error {
	msgDidDocument := msg.GetDidDocString()
	err := didDocumentStatelessVerification(msgDidDocument)
	return err
}

// MsgDeactivateDID Type Methods

const TypeMsgDeactivateDID = "deactivate_did"

var _ sdk.Msg = &MsgDeactivateDID{}

func NewMsgDeactivateDID(creator string, didId string, versionId string, signatures []*SignInfo) *MsgDeactivateDID {
	return &MsgDeactivateDID{
		Creator:    creator,
		DidId:      didId,
		VersionId:  versionId,
		Signatures: signatures,
	}
}

func (msg *MsgDeactivateDID) Route() string {
	return RouterKey
}

func (msg *MsgDeactivateDID) Type() string {
	return TypeMsgDeactivateDID
}

func (msg *MsgDeactivateDID) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeactivateDID) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeactivateDID) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
