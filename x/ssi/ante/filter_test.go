package ante

import (
	"testing"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ssitypes "github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/stretchr/testify/require"
)

func TestFilterMsgsIntoSSIAndNonSSI(t *testing.T) {
	t.Run("two SSI messages", func(t *testing.T) {
		msgList := []sdk.Msg{
			&ssitypes.MsgCreateDID{},
			&ssitypes.MsgCreateSchema{},
		}

		ssiMsgs, nonSSImsgs := filterMsgsIntoSSIAndNonSSI(msgList)

		require.Equal(t, 2, len(ssiMsgs))
		require.Equal(t, 0, len(nonSSImsgs))
	})

	t.Run("one SSI message and one non SSI message", func(t *testing.T) {
		msgList := []sdk.Msg{
			&ssitypes.MsgCreateDID{},
			&authz.MsgGrant{},
		}

		ssiMsgs, nonSSImsgs := filterMsgsIntoSSIAndNonSSI(msgList)

		require.Equal(t, 1, len(ssiMsgs))
		require.Equal(t, 1, len(nonSSImsgs))
	})

	t.Run("one SSI message and one AuthzExec message containing two SSI message", func(t *testing.T) {
		msgList := []sdk.Msg{
			&ssitypes.MsgCreateDID{},
			&authz.MsgExec{
				Msgs: makeMsgsForAny(
					&ssitypes.MsgCreateDID{},
					&ssitypes.MsgCreateSchema{},
				),
			},
		}

		ssiMsgs, nonSSImsgs := filterMsgsIntoSSIAndNonSSI(msgList)

		require.Equal(t, 3, len(ssiMsgs))
		require.Equal(t, 0, len(nonSSImsgs))
	})

	t.Run("one SSI message and one AuthzExec message containing two non SSI message", func(t *testing.T) {
		msgList := []sdk.Msg{
			&ssitypes.MsgCreateDID{},
			&authz.MsgExec{
				Msgs: makeMsgsForAny(
					&banktypes.MsgSend{},
					&banktypes.MsgSend{},
				),
			},
		}

		ssiMsgs, nonSSImsgs := filterMsgsIntoSSIAndNonSSI(msgList)

		require.Equal(t, 1, len(ssiMsgs))
		require.Equal(t, 2, len(nonSSImsgs))
	})

	t.Run("two non SSI message and one AuthzExec message containing two SSI messages", func(t *testing.T) {
		msgList := []sdk.Msg{
			&banktypes.MsgSend{},
			&banktypes.MsgSend{},
			&authz.MsgExec{
				Msgs: makeMsgsForAny(
					&ssitypes.MsgCreateDID{},
					&ssitypes.MsgCreateSchema{},
				),
			},
		}

		ssiMsgs, nonSSIMsgs := filterMsgsIntoSSIAndNonSSI(msgList)

		require.Equal(t, 2, len(ssiMsgs))
		require.Equal(t, 2, len(nonSSIMsgs))
	})
}

func makeMsgsForAny(msgs ...sdk.Msg) []*cdctypes.Any {
	anyMsgs := []*cdctypes.Any{}

	for _, msg := range msgs {
		anyMsg, err := cdctypes.NewAnyWithValue(msg)
		if err != nil {
			panic(err)
		}
		anyMsgs = append(anyMsgs, anyMsg)
	}

	return anyMsgs
}
