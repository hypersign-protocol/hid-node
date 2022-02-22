package cli

import (
	"crypto/ed25519"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/spf13/cobra"
)

const VerKeyFlag = "ver-key"

func CmdCreateDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-did [did-doc-string] [verification-method-id]",
		Short: "Registers the DidDocString",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDidDocString := args[0]
			verificationMethodId := args[1]

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

			verKeyPriv, err := getVerKey(cmd, clientCtx)
			if err != nil {
				return err
			}

			// // Build identity message
			signBytes := didDoc.GetSignBytes()
			signatureBytes := ed25519.Sign(verKeyPriv, signBytes)

			signInfo := types.SignInfo{
				VerificationMethodId: verificationMethodId,
				Signature:            base64.StdEncoding.EncodeToString(signatureBytes),
			}

			msg := types.MsgCreateDID{
				DidDocString: &didDoc,
				Signatures:   []*types.SignInfo{&signInfo},
				Creator:      clientCtx.GetFromAddress().String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(VerKeyFlag, "", "Base64 encoded ed25519 private key to sign identity message with. ")
	return cmd
}

func CmdUpdateDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-did [did-doc-string] [version-id] [verification-method-id]",
		Short: "Updates the DID",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDidDocString := args[0]
			argVersionId := args[1]
			argVerificationMethodId := args[2]

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

			verKeyPriv, err := getVerKey(cmd, clientCtx)
			if err != nil {
				return err
			}

			// // Build identity message
			signBytes := didDoc.GetSignBytes()
			signatureBytes := ed25519.Sign(verKeyPriv, signBytes)

			signInfo := types.SignInfo{
				VerificationMethodId: argVerificationMethodId,
				Signature:            base64.StdEncoding.EncodeToString(signatureBytes),
			}

			msg := types.MsgUpdateDID{
				Creator:      clientCtx.GetFromAddress().String(),
				DidDocString: &didDoc,
				VersionId:    argVersionId,
				Signatures:   []*types.SignInfo{&signInfo},
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(VerKeyFlag, "", "Base64 encoded ed25519 private key to sign identity message with. ")
	return cmd
}

func CmdCreateSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-schema [schema] [verification-method-id]",
		Short: "Broadcast message createSchema",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSchema := args[0]
			argVerificationMethodId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal Schema
			var schema types.Schema
			err = clientCtx.Codec.UnmarshalJSON([]byte(argSchema), &schema)
			if err != nil {
				return err
			}

			verKeyPriv, err := getVerKey(cmd, clientCtx)
			if err != nil {
				return err
			}

			signBytes := schema.GetSignBytes()
			signatureBytes := ed25519.Sign(verKeyPriv, signBytes)

			signInfo := types.SignInfo{
				VerificationMethodId: argVerificationMethodId,
				Signature:            base64.StdEncoding.EncodeToString(signatureBytes),
			}

			msg := types.MsgCreateSchema{
				Schema:     &schema,
				Signatures: []*types.SignInfo{&signInfo},
				Creator:    clientCtx.GetFromAddress().String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(VerKeyFlag, "", "Base64 encoded ed25519 private key to sign identity message with. ")
	return cmd
}

func CmdDeactivateDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deactivate-did [did-doc-string] [version-id] [verification-method-id]",
		Short: "Deactivates the DID",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDidDocString := args[0]
			argVersionId := args[1]
			argVerificationMethodId := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var didDoc types.Did
			err = clientCtx.Codec.UnmarshalJSON([]byte(argDidDocString), &didDoc)
			if err != nil {
				return err
			}

			verKeyPriv, err := getVerKey(cmd, clientCtx)
			if err != nil {
				return err
			}

			// // Build identity message
			signBytes := didDoc.GetSignBytes()
			signatureBytes := ed25519.Sign(verKeyPriv, signBytes)

			signInfo := types.SignInfo{
				VerificationMethodId: argVerificationMethodId,
				Signature:            base64.StdEncoding.EncodeToString(signatureBytes),
			}


			msg := types.MsgDeactivateDID{
				Creator:      clientCtx.GetFromAddress().String(),
				DidDocString: &didDoc,
				VersionId:    argVersionId,
				Signatures:   []*types.SignInfo{&signInfo},
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String(VerKeyFlag, "", "Base64 encoded ed25519 private key to sign identity message with. ")
	return cmd
}
