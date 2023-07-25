package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/hypersign-protocol/hid-node/x/identityfee/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for identityfee module
func GetQueryCmd() *cobra.Command {
	// Group identityfee queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(cmdListFees())
	return cmd
}


func cmdListFees() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-fees",
		Short: "List fee for all SSI based transactions",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.QuerySSIFee(cmd.Context(), &types.QuerySSIFeeRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}