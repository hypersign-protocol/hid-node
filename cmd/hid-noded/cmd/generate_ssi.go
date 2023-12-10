package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/hypersign-protocol/hid-node/app"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const fromFlag = "from"
const didAliasFlag = "did-alias"
const keyringBackendFlag = "keyring-backend"
const didNamespaceFlag = "did-namespace"

func generateSSICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ssi-tools",
		Short: "commands to experiment around Self Sovereign Identity (SSI) documents",
	}

	cmd.AddCommand(generateDidCmd())
	cmd.AddCommand(showDidByAliasCmd())
	cmd.AddCommand(listAllDidAliasesCmd())

	return cmd
}

func listAllDidAliasesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-did-aliases",
		Short: "List all DID Document alias names",
		RunE: func(cmd *cobra.Command, _ []string) error {
			didAliasConfig, err := types.GetDidAliasConfig(cmd)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			result := []map[string]string{}

			if _, err := os.Stat(didAliasConfig.DidAliasDir); err != nil {
				if os.IsNotExist(err) {
					fmt.Fprintf(cmd.ErrOrStderr(), "%v\n\n", []string{})
					return nil
				}
			}
			didJsonFiles, err := os.ReadDir(didAliasConfig.DidAliasDir)
			if err != nil {
				return err
			}

			// Consider only those files whose extensions are '.json'
			for _, didJsonFile := range didJsonFiles {
				isDidJsonFile := !didJsonFile.IsDir() && (strings.Split(didJsonFile.Name(), ".")[1] == "json")
				if isDidJsonFile {
					unit := map[string]string{}
					didDocBytes, err := os.ReadFile(filepath.Join(didAliasConfig.DidAliasDir, didJsonFile.Name()))
					if err != nil {
						return err
					}

					var didDoc types.DidDocument
					err = clientCtx.Codec.UnmarshalJSON(didDocBytes, &didDoc)
					if err != nil {
						// Ignore any files which are not able to parse into type.Did
						continue
					}

					unit["did"] = didDoc.Id
					unit["alias"] = strings.Split(didJsonFile.Name(), ".")[0]
					result = append(result, unit)
				} else {
					continue
				}
			}

			// Indent Map
			resultBytes, err := json.MarshalIndent(result, "", " ")
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(cmd.ErrOrStderr(), "%v\n", string(resultBytes))
			return err
		},
	}
	return cmd
}

func showDidByAliasCmd() *cobra.Command {
	exampleString := "hid-noded ssi-tools show-did-by-alias didsample3"

	cmd := &cobra.Command{
		Use:     "show-did-by-alias [alias-name]",
		Args:    cobra.ExactArgs(1),
		Example: exampleString,
		Short:   "Retrieve the Did Document by an alias name",
		RunE: func(cmd *cobra.Command, args []string) error {
			didAliasConfig, err := types.GetDidAliasConfig(cmd)
			if err != nil {
				return err
			}

			aliasName := args[0]
			aliasFile := aliasName + ".json"

			if _, err := os.Stat(didAliasConfig.DidAliasDir); err != nil {
				if os.IsNotExist(err) {
					fmt.Fprintf(cmd.ErrOrStderr(), "DID Document alias '%v' does not exist\n", aliasName)
					return nil
				}
			}

			didDocBytes, err := os.ReadFile(filepath.Join(didAliasConfig.DidAliasDir, aliasFile))
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "DID Document alias '%v' does not exist\n", aliasName)
				return nil
			}

			_, err = fmt.Fprintf(cmd.ErrOrStderr(), "%v\n", string(didDocBytes))
			return err
		},
	}

	return cmd
}

func generateDidCmd() *cobra.Command {
	exampleString1 := "hid-noded ssi-tools generate-did --from hid1kspgn6f5hmurulx4645ch6rf0kt90jpv5ydykp --keyring-backend test --did-alias example1"
	exampleString2 := "hid-noded ssi-tools generate-did --from node1 --keyring-backend test --did-alias example2"
	exampleString3 := "hid-noded ssi-tools generate-did --from node1 --keyring-backend test --did-alias example3 --did-namespace devnet"

	cmd := &cobra.Command{
		Use:     "generate-did",
		Short:   "Generates a DID Document",
		Example: exampleString1 + "\n" + exampleString2 + "\n" + exampleString3,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Get the flags
			account, err := cmd.Flags().GetString(fromFlag)
			if err != nil {
				return err
			}
			if account == "" {
				return fmt.Errorf("no value provided for --from flag")
			}

			didAlias, err := cmd.Flags().GetString(didAliasFlag)
			if err != nil {
				return err
			}
			if didAlias == "" {
				return fmt.Errorf("no value provided for --did-alias flag")
			}

			keyringBackend, err := cmd.Flags().GetString(keyringBackendFlag)
			if err != nil {
				return err
			}
			if keyringBackend == "" {
				return fmt.Errorf("no value provided for --keyring-backend flag")
			}

			didNamespace, err := cmd.Flags().GetString(didNamespaceFlag)
			if err != nil {
				return err
			}

			// Get Public Key from keyring account
			var kr keyring.Keyring
			appName := "hid-noded-keyring"
			didAliasConfig, err := types.GetDidAliasConfig(cmd)
			if err != nil {
				return err
			}

			switch keyringBackend {
			case "test":
				kr, err = keyring.New(appName, "test", didAliasConfig.HidNodeConfigDir, nil, nil)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unsupported keyring-backend : %v", keyringBackend)
			}

			// Handle both key name as well as key address
			userKeyInfo, errAccountFetch := kr.Key(account)
			if errAccountFetch != nil {
				if accountAddr, err := sdk.AccAddressFromBech32(account); err != nil {
					return err
				} else {
					userKeyInfo, errAccountFetch = kr.KeyByAddress(accountAddr)
					if errAccountFetch != nil {
						return errAccountFetch
					}
				}
			}

			pubKey, err := userKeyInfo.GetPubKey()
			if err != nil {
				return err
			}
			pubKeyMultibase, err := multibase.Encode(multibase.Base58BTC, pubKey.Bytes())
			if err != nil {
				return err
			}
			userKeyInfoAddress, err := userKeyInfo.GetAddress()
			if err != nil {
				return err
			}
			userBlockchainAddress := sdk.MustBech32ifyAddressBytes(
				app.Bech32Prefix,
				userKeyInfoAddress.Bytes(),
			)

			// Generate a DID document with both publicKeyMultibase and blockchainAccountId
			didDoc := generateDidDoc(didNamespace, pubKeyMultibase, userBlockchainAddress)

			// Construct the JSON and store it in $HOME/.hid-node/generated-ssi-docs
			if _, err := os.Stat(didAliasConfig.DidAliasDir); err != nil {
				if os.IsNotExist(err) {
					if err := os.Mkdir(didAliasConfig.DidAliasDir, os.ModePerm); err != nil {
						return err
					}
				} else {
					return err
				}
			}

			didJsonBytes, err := json.MarshalIndent(didDoc, "", " ")
			if err != nil {
				return err
			}
			didJsonFilename := didAlias + ".json"
			didJsonPath := filepath.Join(didAliasConfig.DidAliasDir, didJsonFilename)
			if err := os.WriteFile(didJsonPath, didJsonBytes, 0644); err != nil {
				return err
			}

			_, err = fmt.Fprintf(cmd.ErrOrStderr(), "DID Document alias '%v' (didId: %v) has been successfully generated at %v\n", didAlias, didDoc.Id, didJsonPath)

			return err
		},
	}

	cmd.Flags().String(fromFlag, "", "name of account while will sign the DID Document")
	cmd.Flags().String(didAliasFlag, "", "alias of the generated DID Document which can be referred to while registering on-chain")
	cmd.Flags().String(keyringBackendFlag, "", "supported keyring backend: (test)")
	cmd.Flags().String(didNamespaceFlag, "", "namespace of DID Document Id")
	return cmd
}
