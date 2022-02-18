package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const TypeMsgCreateDID = "create_did"

var _ sdk.Msg = &MsgCreateDID{}

func NewMsgCreateDID(did string, didDocString *DidDocStructCreateDID) *MsgCreateDID {
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

// TODO: Implement the following without the field Creator
// func accAddrByKeyRef(keyring keyring.Keyring, keyRef string) (sdk.AccAddress, error) {
// 	// Firstly check if the keyref is a key name of a key registered in a keyring
// 	info, err := keyring.Key(keyRef)

// 	if err == nil {
// 		return info.GetAddress(), nil
// 	}

// 	if !sdkerr.IsOf(err, sdkerr.ErrIO, sdkerr.ErrKeyNotFound) {
// 		return nil, err
// 	}

// 	// Fallback: convert keyref to address
// 	return sdk.AccAddressFromBech32(keyRef)
// }

// func (msg *MsgCreateDID) GetSigners() []sdk.AccAddress {
// 	ctx := client
// 	signerAccAddr, err := accAddrByKeyRef(ctx.Keyring, ctx.From)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return []sdk.AccAddress{signerAccAddr}
// }

func (msg *MsgCreateDID) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *DidDocStructCreateDID) GetSigners() []Signer {
	if len(msg.Controller) > 0 {
		result := make([]Signer, len(msg.Controller))

		for i, controller := range msg.Controller {
			if controller == msg.Id {
				result[i] = Signer{
					Signer:             controller,
					Authentication:     msg.Authentication,
					VerificationMethod: msg.VerificationMethod,
				}
			} else {
				result[i] = Signer{
					Signer: controller,
				}
			}
		}

		return result
	}

	if len(msg.Authentication) > 0 {
		return []Signer{
			{
				Signer:             msg.Id,
				Authentication:     msg.Authentication,
				VerificationMethod: msg.VerificationMethod,
			},
		}
	}

	return []Signer{}
}

func (msg *MsgCreateDID) GetSignBytes() []byte {
	return ModuleCdc.MustMarshal(msg)
}

func (msg *DidDocStructCreateDID) GetSignBytes() []byte {
	return ModuleCdc.MustMarshal(msg)
}

func (msg *MsgCreateDID) ValidateBasic() error {
	did := msg.GetDidDocString().GetId()
	if did == "" {
		return ErrBadRequestIsRequired.Wrap("DID")
	}
	return nil
}
