package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

type SSIMsg sdk.Msg
type NonSSIMsg sdk.Msg

// filterMsgsIntoSSIAndNonSSI filters transaction messages into SSI and non-SSI messages
func filterMsgsIntoSSIAndNonSSI(msgs []sdk.Msg) ([]SSIMsg, []NonSSIMsg) {
	msgList := []sdk.Msg{}
	for _, msg := range msgs {
		// Append every sub-message of Authz's MsgExec to msgList
		if isAuthzExecMsg(msg) {
			subMsgs, err := msg.(*authz.MsgExec).GetMessages()
			if err != nil {
				panic(err)
			}
			msgList = append(msgList, subMsgs...)
		} else {
			msgList = append(msgList, msg)
		}
	}

	// Categorize messages into SSI and non-SSI
	ssiMsgs := []SSIMsg{}
	nonSsiMsgs := []NonSSIMsg{}

	for _, msg := range msgList {
		if isSSIMsg(msg) {
			ssiMsgs = append(ssiMsgs, msg)
		} else {
			nonSsiMsgs = append(nonSsiMsgs, msg)
		}
	}

	return ssiMsgs, nonSsiMsgs
}
