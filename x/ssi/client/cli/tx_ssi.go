package cli

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	ldcontext "github.com/hypersign-protocol/hid-node/x/ssi/ld-context"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/spf13/cobra"
)

const didAliasFlag = "did-alias"

func CmdRegisterDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-did [did-document] ([did-document-proof-1], [did-document-proof-2] .... [did-document-proof-N]) [flags]\n  hid-noded tx ssi register-did --did-alias <name of the DID Alias> [flags]",
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

			var didDoc types.DidDocument
			var didDocumentProofs []*types.DocumentProof
			txAuthorAddr := clientCtx.GetFromAddress()
			txAuthorAddrString := clientCtx.GetFromAddress().String()

			if didAlias == "" {
				// Minimum 2 CLI arguments are expected
				if len(args) < 2 {
					return fmt.Errorf("requires at least 2 arg(s), only received %v", len(args))
				}

				argDidDoc := args[0]

				// Unmarshal DidDocString
				err = clientCtx.Codec.UnmarshalJSON([]byte(argDidDoc), &didDoc)
				if err != nil {
					return err
				}

				// Prepare Signatures
				didDocumentProofs, err = getDocumentProofs(clientCtx, args[1:])
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
				didDocumentProofs = []*types.DocumentProof{
					{
						Type:               types.EcdsaSecp256k1Signature2019,
						VerificationMethod: didDoc.VerificationMethod[0].Id,
						ProofPurpose:       "assertionMethod",
						Created:            time.Now().Format("2006-01-02T15:04:00Z"), // RFC3339 format
					},
					{
						Type:               types.EcdsaSecp256k1Signature2019,
						VerificationMethod: didDoc.VerificationMethod[1].Id,
						ProofPurpose:       "assertionMethod",
						Created:            time.Now().Format("2006-01-02T15:04:00Z"), // RFC3339 format
					},
				}

				// REVERT: get signature for every VM
				for i := 0; i < len(didDocumentProofs); i++ {
					didDocCanonizedHash, err := ldcontext.EcdsaSecp256k1Signature2019Normalize(&didDoc, didDocumentProofs[i])
					if err != nil {
						return err
					}

					keyringBackend, err := cmd.Flags().GetString(flags.FlagKeyringBackend)
					if err != nil {
						return err
					}
					if keyringBackend != "test" {
						return fmt.Errorf("unsupported keyring backend for DID Document Alias Signing: %v", keyringBackend)
					}

					kr, err := keyring.New("hid-node-app", keyringBackend, didAliasConfig.HidNodeConfigDir, nil, clientCtx.Codec)
					if err != nil {
						return err
					}

					signatureBytes, _, err := kr.SignByAddress(txAuthorAddr, didDocCanonizedHash)
					if err != nil {
						return err
					}
					didDocumentProofs[i].ProofValue = base64.StdEncoding.EncodeToString(signatureBytes)

				}
			}

			// Submit RegisterDID Tx
			msg := types.MsgRegisterDID{
				DidDocument:       &didDoc,
				DidDocumentProofs: didDocumentProofs,
				TxAuthor:          txAuthorAddrString,
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
		Use:   "update-did [did-doc] [version-id] ([did-document-proof-1], [did-document-proof-2] .... [did-document-proof-N])",
		Short: "Updates Did Document",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDidDoc := args[0]
			argVersionId := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal DidDocString
			var didDoc types.DidDocument
			err = clientCtx.Codec.UnmarshalJSON([]byte(argDidDoc), &didDoc)
			if err != nil {
				return err
			}

			didDocumentProofs, err := getDocumentProofs(clientCtx, args[2:])
			if err != nil {
				return err
			}

			msg := types.MsgUpdateDID{
				DidDocument:       &didDoc,
				VersionId:         argVersionId,
				DidDocumentProofs: didDocumentProofs,
				TxAuthor:          clientCtx.GetFromAddress().String(),
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
		Short: "Creates Credential Schema",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSchemaDoc := args[0]
			argSchemaProof := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal Schema Document
			var schemaDoc types.CredentialSchemaDocument
			err = clientCtx.Codec.UnmarshalJSON([]byte(argSchemaDoc), &schemaDoc)
			if err != nil {
				return err
			}

			// Unmarshal Schema Proof
			var schemaProof types.DocumentProof
			err = clientCtx.Codec.UnmarshalJSON([]byte(argSchemaProof), &schemaProof)
			if err != nil {
				return err
			}

			msg := types.MsgRegisterCredentialSchema{
				CredentialSchemaDocument: &schemaDoc,
				CredentialSchemaProof:    &schemaProof,
				TxAuthor:                 clientCtx.GetFromAddress().String(),
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

func CmdUpdateSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-schema [schema-doc] [schema-proof]",
		Short: "Updates Credential Schema",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSchemaDoc := args[0]
			argSchemaProof := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal Schema Document
			var schemaDoc types.CredentialSchemaDocument
			err = clientCtx.Codec.UnmarshalJSON([]byte(argSchemaDoc), &schemaDoc)
			if err != nil {
				return err
			}

			// Unmarshal Schema Proof
			var schemaProof types.DocumentProof
			err = clientCtx.Codec.UnmarshalJSON([]byte(argSchemaProof), &schemaProof)
			if err != nil {
				return err
			}

			msg := types.MsgUpdateCredentialSchema{
				CredentialSchemaDocument: &schemaDoc,
				CredentialSchemaProof:    &schemaProof,
				TxAuthor:                 clientCtx.GetFromAddress().String(),
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
		Use:   "deactivate-did [did-id] [version-id] ([did-document-proof-1], [did-document-proof-2] .... [did-document-proof-N])",
		Short: "Deactivates Did Document",
		Args:  cobra.MinimumNArgs(3),
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
			if _, err := queryClient.DidDocumentByID(cmd.Context(), requestParams); err != nil {
				return err
			}

			didDocumentProofs, err := getDocumentProofs(clientCtx, args[2:])
			if err != nil {
				return err
			}

			msg := types.MsgDeactivateDID{
				DidDocumentId:     argDidId,
				VersionId:         argVersionId,
				DidDocumentProofs: didDocumentProofs,
				TxAuthor:          clientCtx.GetFromAddress().String(),
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
				credentialStatus types.CredentialStatusDocument
				proof            types.DocumentProof
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
				CredentialStatusDocument: &credentialStatus,
				CredentialStatusProof:    &proof,
				TxAuthor:                 clientCtx.GetFromAddress().String(),
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

func CmdUpdateCredentialStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-credential-status [credential-status] [proof]",
		Short: "Updates the status of Verifiable Credential",
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
				credentialStatus types.CredentialStatusDocument
				proof            types.DocumentProof
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

			msg := types.MsgUpdateCredentialStatus{
				CredentialStatusDocument: &credentialStatus,
				CredentialStatusProof:    &proof,
				TxAuthor:                 clientCtx.GetFromAddress().String(),
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
