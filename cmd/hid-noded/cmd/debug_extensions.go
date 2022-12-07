package cmd

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/hypersign-protocol/vid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"
)

func extendDebug(debugCmd *cobra.Command) *cobra.Command {
	debugCmd.AddCommand(ed25519Cmd())
	return debugCmd
}

// ed25519Cmd returns cobra Command.
func ed25519Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ed25519",
		Short: "ed25519 debug commands",
	}

	cmd.AddCommand(
		ed25519RandomCmd(),
		base64toMultibase58Cmd(),
		multibase58toBase64Cmd(),
		signSSIDocCmd(),
	)

	return cmd
}

// Generate signature for Schema and VC Status Document
func signSSIDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-ssi-doc",
		Short: "Get signature for the signed document",
	}

	cmd.AddCommand(
		signSchemaDocCmd(),
		signCredStatusDocCmd(),
	)
	return cmd
}

func signSchemaDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema-doc [doc] [base64 encoded private-key]",
		Short: "Schema Document signature",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argSchemaDoc := args[0]
			privateKey, err := base64.StdEncoding.DecodeString(args[1])
			if err != nil {
				return err
			}

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

			// Sign Schema Document
			schemaDocBytes := schemaDoc.GetSignBytes()
			signature := base64.StdEncoding.EncodeToString(
				ed25519.Sign(privateKey, schemaDocBytes),
			)

			_, err = fmt.Fprintln(cmd.OutOrStdout(), signature)
			return err
		},
	}
	return cmd
}

func signCredStatusDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cred-status-doc [doc] [base64 encoded private-key]",
		Short: "Credential Status Document signature",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argCredStatusDoc := args[0]
			privateKey, err := base64.StdEncoding.DecodeString(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Unmarshal Credential Status Document
			var credStatusDoc types.CredentialStatus
			err = clientCtx.Codec.UnmarshalJSON([]byte(argCredStatusDoc), &credStatusDoc)
			if err != nil {
				return err
			}

			// Sign Credential Status Document
			credStatusDocBytes := credStatusDoc.GetSignBytes()
			signature := base64.StdEncoding.EncodeToString(
				ed25519.Sign(privateKey, credStatusDocBytes),
			)

			_, err = fmt.Fprintln(cmd.OutOrStdout(), signature)
			return err
		},
	}
	return cmd
}

// ed25519Cmd returns cobra Command.
func ed25519RandomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Generate random ed25519 keypair",
		RunE: func(cmd *cobra.Command, args []string) error {
			pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
			if err != nil {
				return err
			}

			keyInfo := struct {
				PubKeyBase64    string `json:"pub_key_base_64"`
				PubKeyMultibase string `json:"pub_key_multibase"`
				PrivKeyBase64   string `json:"priv_key_base_64"`
			}{
				PubKeyBase64:    base64.StdEncoding.EncodeToString(pubKey),
				PubKeyMultibase: "z" + base58.Encode(pubKey),
				PrivKeyBase64:   base64.StdEncoding.EncodeToString(privKey),
			}

			keyInfoJson, err := json.Marshal(keyInfo)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), string(keyInfoJson))
			return err
		},
	}

	return cmd
}

// Converts base-64 encoded ed25519 public key to multibase58 string
func base64toMultibase58Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "base64-multibase58 [base64-encoded-public-key]",
		Short: "Convert base64 string to multibase58 string",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			base64Str := args[0]
			bytes, err := base64.StdEncoding.DecodeString(base64Str)
			if err != nil {
				return err
			}

			multibase58Str, err := multibase.Encode(multibase.Base58BTC, bytes)
			if err != nil {
				return err
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), multibase58Str)
			return err
		},
	}

	return cmd
}

// Converts multibase58 encoded ed25519 public key to base-64 string
func multibase58toBase64Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "multibase58-base64 [multibase58-public-key]",
		Short: "Convert multibase58 string to base64 string",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			multibase58Str := args[0][1:]
			multibase58Bytes := base58.Decode(multibase58Str)

			base64Str := base64.StdEncoding.EncodeToString(multibase58Bytes)

			_, err := fmt.Fprintln(cmd.OutOrStdout(), base64Str)
			return err
		},
	}

	return cmd
}
