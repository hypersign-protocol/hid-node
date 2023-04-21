package cli

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/spf13/cobra"
)

const didAliasFlag = "did-alias"

func CmdCreateDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-did [did-doc-string] ([vm-id-1] [sign-key-1] [sign-key-algo-1] ... [vm-id-N] [sign-key-N] [sign-key-algo-N]) [flags]\n  hid-noded tx ssi create-did --did-alias <name of the DID Alias> [flags]",
		Short: "Registers a DID Document",
		Args:  cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			didAlias, err := cmd.Flags().GetString(didAliasFlag)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var didDoc types.Did
			var signInfos []*types.SignInfo
			txAuthorAddr := clientCtx.GetFromAddress()
			txAuthorAddrString := clientCtx.GetFromAddress().String()

			if didAlias == "" {
				// Minimum 4 CLI arguments are expected
				if len(args) < 4 {
					return fmt.Errorf("requires at least 4 arg(s), only received %v", len(args))
				}

				argDidDocString := args[0]

				// Unmarshal DidDocString
				err = clientCtx.Codec.UnmarshalJSON([]byte(argDidDocString), &didDoc)
				if err != nil {
					return err
				}

				// Prepare Signatures
				signInfos, err = getSignatures(cmd, didDoc.GetSignBytes(), args[1:])
				if err != nil {
					return err
				}
			} else {
				// Get the DID Document from local
				didAliasConfig, err := types.GetDidAliasConfig(cmd)
				if err != nil {
					return fmt.Errorf("failed to read DID Alias config: %v", err.Error())
				}

				aliasFile := didAlias + ".json"
				didDocBytes, err := os.ReadFile(filepath.Join(didAliasConfig.DidAliasDir, aliasFile))
				if err != nil {
					fmt.Fprintf(cmd.ErrOrStderr(), "DID Document alias '%v' does not exist\n", didAlias)
					return nil
				}

				err = clientCtx.Codec.UnmarshalJSON(didDocBytes, &didDoc)
				if err != nil {
					return err
				}

				// Ensure the --from flag value matches with publicKey multibase

				// Since DID Alias will always have one verification method object, it is safe to
				// choose the 0th index
				publicKeyMultibase := didDoc.VerificationMethod[0].PublicKeyMultibase

				if err := validateDidAliasSignerAddress(txAuthorAddrString, publicKeyMultibase); err != nil {
					return fmt.Errorf("%v: %v", err.Error(), didAlias)
				}

				// Sign the DID Document using Keyring to get theSignInfo. Currently, "test" keyring-backend is only supported
				keyringBackend, err := cmd.Flags().GetString(flags.FlagKeyringBackend)
				if err != nil {
					return err
				}
				if keyringBackend != "test" {
					return fmt.Errorf("unsupporeted keyring backend for DID Document Alias Signing: %v", keyringBackend)
				}

				kr, err := keyring.New("hid-node-app", keyringBackend, didAliasConfig.HidNodeConfigDir, nil)
				if err != nil {
					return err
				}

				signatureBytes, _, err := kr.SignByAddress(txAuthorAddr, didDoc.GetSignBytes())
				if err != nil {
					return err
				}
				signatureStr := base64.StdEncoding.EncodeToString(signatureBytes)

				signInfos = []*types.SignInfo{
					{
						VerificationMethodId: didDoc.VerificationMethod[0].Id,
						Signature:            signatureStr,
					},
				}
			}

			// Submit CreateDID Tx
			msg := types.MsgCreateDID{
				DidDocString: &didDoc,
				Signatures:   signInfos,
				Creator:      txAuthorAddrString,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String(didAliasFlag, "", "alias of the generated DID Document which can be referred to while registering on-chain")
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
