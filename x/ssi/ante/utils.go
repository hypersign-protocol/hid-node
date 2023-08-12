package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	ssitypes "github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// isAuthzExecMsg checks if the message is of authz.MsgExec type
func isAuthzExecMsg(msg sdk.Msg) bool {
	switch msg.(type) {
	case *authz.MsgExec:
		return true
	default:
		return false
	}
}

// isSSIMsg checks if the message is of SSI type
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
