package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	ssitypes "github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// isSSIMsg checks if the message is of SSI type or not
func isSSIMsg(msg sdk.Msg) bool {
	switch msg.(type) {
	case *ssitypes.MsgCreateDID:
		return true
	case *ssitypes.MsgUpdateDID:
		return true
	case *ssitypes.MsgDeactivateDID:
		return true
	case *ssitypes.MsgCreateSchema:
		return true
	case *ssitypes.MsgRegisterCredentialStatus:
		return true
	default:
		return false
	}
}

// isSSIMsgPresentInTx checks if there is any SSI message present in the transaction
func isSSIMsgPresentInTx(tx sdk.Tx) bool {
	for _, msg := range tx.GetMsgs() {
		if isSSIMsg(msg) {
			return true
		}
	}

	return false
}

// isNonSSIMsgPresentInTx checks if there is any non-SSI message present in the transaction
func isNonSSIMsgPresentInTx(tx sdk.Tx) bool {
	for _, msg := range tx.GetMsgs() {
		if !isSSIMsg(msg) {
			return true
		}
	}

	return false
}
