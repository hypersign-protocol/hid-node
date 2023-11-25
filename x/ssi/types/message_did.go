package types

import (
	"encoding/hex"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

func (msg *DidDocument) GetSignBytes() []byte {
	return ModuleCdc.MustMarshal(msg)
}

// MsgCreateDID Type Methods
const TypeMsgCreateDID = "register_did"

var _ sdk.Msg = &MsgRegisterDID{}

func NewMsgCreateDID(didDoc *DidDocument, documentProofs []*DocumentProof, txAuthor string) *MsgRegisterDID {
	return &MsgRegisterDID{
		DidDocument:       didDoc,
		DidDocumentProofs: documentProofs,
		TxAuthor:          txAuthor,
	}
}

func (msg *MsgRegisterDID) Route() string {
	return RouterKey
}

func (msg *MsgRegisterDID) Type() string {
	return TypeMsgCreateDID
}

func (msg *MsgRegisterDID) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.TxAuthor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterDID) GetSignBytes() []byte {
	return ModuleCdc.MustMarshal(msg)
}

func (msg *MsgRegisterDID) ValidateBasic() error {
	didDoc := msg.DidDocument
	if err := didDoc.ValidateDidDocument(); err != nil {
		return err
	}
	return nil
}

// MsgUpdateDID Type Methods

const TypeMsgUpdateDID = "update_did"

var _ sdk.Msg = &MsgUpdateDID{}

func NewMsgUpdateDID(
	didDoc *DidDocument,
	documentProofs []*DocumentProof,
	versionId string,
	txAuthor string,
) *MsgUpdateDID {
	return &MsgUpdateDID{
		DidDocument:       didDoc,
		DidDocumentProofs: documentProofs,
		VersionId:         versionId,
		TxAuthor:          txAuthor,
	}
}

func (msg *MsgUpdateDID) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDID) Type() string {
	return TypeMsgUpdateDID
}

func (msg *MsgUpdateDID) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.TxAuthor)
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
	didDoc := msg.DidDocument
	if err := didDoc.ValidateDidDocument(); err != nil {
		return err
	}
	return nil
}

// MsgDeactivateDID Type Methods

const TypeMsgDeactivateDID = "deactivate_did"

var _ sdk.Msg = &MsgDeactivateDID{}

func NewMsgDeactivateDID(didId string, versionId string, documentProofs []*DocumentProof, txAuthor string) *MsgDeactivateDID {
	return &MsgDeactivateDID{
		DidDocumentId:     didId,
		VersionId:         versionId,
		DidDocumentProofs: documentProofs,
		TxAuthor:          txAuthor,
	}
}

func (msg *MsgDeactivateDID) Route() string {
	return RouterKey
}

func (msg *MsgDeactivateDID) Type() string {
	return TypeMsgDeactivateDID
}

func (msg *MsgDeactivateDID) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.TxAuthor)
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
	_, err := sdk.AccAddressFromBech32(msg.TxAuthor)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func CreateNewMetadata(ctx sdk.Context) DidDocumentMetadata {
	return DidDocumentMetadata{
		VersionId:   strings.ToUpper(hex.EncodeToString(tmhash.Sum([]byte(ctx.TxBytes())))),
		Deactivated: false,
		Created:     ctx.BlockTime().Format(time.RFC3339),
		Updated:     ctx.BlockTime().Format(time.RFC3339),
	}
}
