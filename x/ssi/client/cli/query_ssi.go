package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-schema [schema-id]",
		Short: "Query Schema for a given schema id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSchemaId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetSchemaRequest{SchemaId: argSchemaId}

			res, err := queryClient.GetSchema(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdSchemas() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema-list",
		Short: "Query Registered Schemas",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySchemasRequest{}

			res, err := queryClient.Schemas(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdSchemaCount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema-count",
		Short: "Get the Schema Count",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QuerySchemaCountRequest{}

			res, err := queryClient.SchemaCount(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}