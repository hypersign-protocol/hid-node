package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/hypersign-protocol/vid-node/x/ssi/types"
	"github.com/spf13/cobra"
)

func CmdCreateDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-did [did-doc-string] [vm-id-1] [sign-key-1] [sign-key-algo-1] ... [vm-id-N] [sign-key-N] [sign-key-algo-N]",
		Short: "Registers a DID Document",
		Args:  cobra.MinimumNArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDidDocString := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal DidDocString
			var didDoc types.Did
			err = clientCtx.Codec.UnmarshalJSON([]byte(argDidDocString), &didDoc)
			if err != nil {
				return err
			}

			// Prepare Signatures
			signInfos, err := getSignatures(cmd, didDoc.GetSignBytes(), args[1:])
			if err != nil {
				return err
			}

			msg := types.MsgCreateDID{
				DidDocString: &didDoc,
				Signatures:   signInfos,
				Creator:      clientCtx.GetFromAddress().String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdUpdateDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-did [did-doc-string] [version-id] [vm-id-1] [sign-key-1] [sign-key-algo-1] ... [vm-id-N] [sign-key-N] [sign-key-algo-N]",
		Short: "Updates Did Document",
		Args:  cobra.MinimumNArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDidDocString := args[0]
			argVersionId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal DidDocString
			var didDoc types.Did
			err = clientCtx.Codec.UnmarshalJSON([]byte(argDidDocString), &didDoc)
			if err != nil {
				return err
			}

			signInfos, err := getSignatures(cmd, didDoc.GetSignBytes(), args[2:])
			if err != nil {
				return err
			}

			msg := types.MsgUpdateDID{
				Creator:      clientCtx.GetFromAddress().String(),
				DidDocString: &didDoc,
				VersionId:    argVersionId,
				Signatures:   signInfos,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdCreateSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-schema [schema-doc] [schema-proof]",
		Short: "Creates Schema",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSchemaDoc := args[0]
			argSchemaProof := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal Schema Document
			var schemaDoc types.SchemaDocument
			err = clientCtx.Codec.UnmarshalJSON([]byte(argSchemaDoc), &schemaDoc)
			if err != nil {
				return err
			}

			// Unmarshal Schema Proof
			var schemaProof types.SchemaProof
			err = clientCtx.Codec.UnmarshalJSON([]byte(argSchemaProof), &schemaProof)
			if err != nil {
				return err
			}

			msg := types.MsgCreateSchema{
				SchemaDoc:   &schemaDoc,
				SchemaProof: &schemaProof,
				Creator:     clientCtx.GetFromAddress().String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdDeactivateDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deactivate-did [did-id] [version-id] [vm-id-1] [sign-key-1] [sign-key-algo-1] ... [vm-id-N] [sign-key-N] [sign-key-algo-N]",
		Short: "Deactivates Did Document",
		Args:  cobra.MinimumNArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDidId := args[0]
			argVersionId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Query Did Document from store using Did Id
			queryClient := types.NewQueryClient(clientCtx)
			requestParams := &types.QueryDidDocumentRequest{DidId: argDidId}
			resolvedDidDocument, err := queryClient.QueryDidDocument(cmd.Context(), requestParams)
			if err != nil {
				return err
			}
			didDoc := resolvedDidDocument.GetDidDocument()

			signInfos, err := getSignatures(cmd, didDoc.GetSignBytes(), args[2:])
			if err != nil {
				return err
			}

			msg := types.MsgDeactivateDID{
				Creator:    clientCtx.GetFromAddress().String(),
				DidId:      argDidId,
				VersionId:  argVersionId,
				Signatures: signInfos,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdRegisterCredentialStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-credential-status [credential-status] [proof]",
		Short: "Registers the status of Verifiable Credential",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCredStatus := args[0]
			argProof := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal Credential Status
			var (
				credentialStatus types.CredentialStatus
				proof            types.CredentialProof
			)

			err = clientCtx.Codec.UnmarshalJSON([]byte(argCredStatus), &credentialStatus)
			if err != nil {
				return err
			}

			// Unmarshal Proof
			err = clientCtx.Codec.UnmarshalJSON([]byte(argProof), &proof)
			if err != nil {
				return err
			}

			msg := types.MsgRegisterCredentialStatus{
				CredentialStatus: &credentialStatus,
				Proof:            &proof,
				Creator:          clientCtx.GetFromAddress().String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
