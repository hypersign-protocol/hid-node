package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/hypersign-protocol/hid-node/x/did/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-did [did] [did-doc-string] [created-at]",
		Short: "Broadcast message createDID",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDid := args[0]
			argDidDocString := args[1]
			argCreatedAt := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDID(
				clientCtx.GetFromAddress().String(),
				argDid,
				argDidDocString,
				argCreatedAt,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
