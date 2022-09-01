package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateDID())
	cmd.AddCommand(CmdUpdateDID())
	cmd.AddCommand(CmdCreateSchema())
	cmd.AddCommand(CmdDeactivateDID())
	cmd.AddCommand(CmdRegisterCredentialStatus())
	// this line is used by starport scaffolding # 1

	return cmd
}
