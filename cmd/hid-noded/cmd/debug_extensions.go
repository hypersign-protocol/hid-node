package cmd

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"
	bbs "github.com/hyperledger/aries-framework-go/component/kmscrypto/crypto/primitive/bbs12381g2pub"
	hidnodecli "github.com/hypersign-protocol/hid-node/x/ssi/client/cli"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
	"github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"
	secp256k1 "github.com/tendermint/tendermint/crypto/secp256k1"
)

func extendDebug(debugCmd *cobra.Command) *cobra.Command {
	debugCmd.AddCommand(
		ed25519Cmd(),
		secp256k1Cmd(),
		bbsCmd(),
		signSSIDocCmd(),
	)
	return debugCmd
}

func bbsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bbs",
		Short: "bbs debug commands",
	}

	cmd.AddCommand(
		blsRandomCmd(),
	)

	return cmd
}

func secp256k1Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secp256k1",
		Short: "secp256k1 debug commands",
	}

	cmd.AddCommand(
		secp256k1RandomCmd(),
		secp256k1Bech32AddressCmd(),
		secp256k1EthRandomCmd(),
	)

	return cmd
}

func blsRandomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Generate random blsBbs keypair",
		RunE: func(cmd *cobra.Command, args []string) error {
			pubKey, privKey, err := bbs.GenerateKeyPair(sha256.New, nil)
			if err != nil {
				return err
			}

			// Convert Public Key Object to Multibase
			pubKeyBytes, err := pubKey.Marshal()
			if err != nil {
				return err
			}

			publicKeyMultibase, err := multibase.Encode(multibase.Base58BTC, pubKeyBytes)
			if err != nil {
				return err
			}

			// Convert Private Object to Bytes
			privKeyBytes, err := privKey.Marshal()
			if err != nil {
				return err
			}

			keyInfo := struct {
				PubKeyBase64    string `json:"pub_key_base_64"`
				PubKeyMultibase string `json:"pub_key_multibase"`
				PrivKeyBase64   string `json:"priv_key_base_64"`
			}{
				PubKeyBase64:    base64.StdEncoding.EncodeToString(pubKeyBytes),
				PubKeyMultibase: publicKeyMultibase,
				PrivKeyBase64:   base64.StdEncoding.EncodeToString(privKeyBytes),
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

func secp256k1RandomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "random",
		Short: "Generate random secp256k1 keypair",
		RunE: func(cmd *cobra.Command, args []string) error {
			privateKeyObj := secp256k1.GenPrivKey()

			privateKey := privateKeyObj.Bytes()
			publicKeyCompressed := privateKeyObj.PubKey().Bytes()

			publicKeyMultibase, err := multibase.Encode(multibase.Base58BTC, publicKeyCompressed)
			if err != nil {
				panic(err)
			}

			keyInfo := struct {
				PubKeyBase64    string `json:"pub_key_base_64"`
				PubKeyMultibase string `json:"pub_key_multibase"`
				PrivKeyBase64   string `json:"priv_key_base_64"`
			}{
				PubKeyBase64:    base64.StdEncoding.EncodeToString(publicKeyCompressed),
				PubKeyMultibase: publicKeyMultibase,
				PrivKeyBase64:   base64.StdEncoding.EncodeToString(privateKey),
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

func secp256k1Bech32AddressCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bech32-addr [base64-encoded-public-key] [prefix]",
		Short: "Converts a compressed base64 encoded secp256k1 public key to bech32 address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			base64PubKey := args[0]
			addrPrefix := args[1]

			publicKeyBytes, err := base64.StdEncoding.DecodeString(base64PubKey)
			if err != nil {
				panic(err)
			}

			bech32address := publicKeyToBech32Address(addrPrefix, publicKeyBytes)

			_, err = fmt.Fprintln(cmd.OutOrStdout(), bech32address)
			return err
		},
	}

	return cmd
}

func secp256k1EthRandomCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "eth-hex-random",
		Short: "Generate random Ethereum hex-encoded secp256k1 keypair",
		RunE: func(cmd *cobra.Command, args []string) error {
			privateKeyObj := secp256k1.GenPrivKey()
			privateKey := privateKeyObj.Bytes()

			publicKeyCompressed := privateKeyObj.PubKey().Bytes()

			publicKeyUncompressed, err := ethercrypto.DecompressPubkey(publicKeyCompressed)
			if err != nil {
				return err
			}
			ethereumAddress := ethercrypto.PubkeyToAddress(*publicKeyUncompressed).Hex()

			keyInfo := struct {
				PubKeyBase64    string `json:"pub_key_hex"`
				PrivKeyBase64   string `json:"priv_key_hex"`
				EthereumAddress string `json:"ethereum_address"`
			}{
				PubKeyBase64:    hex.EncodeToString(publicKeyCompressed),
				PrivKeyBase64:   hex.EncodeToString(privateKey),
				EthereumAddress: ethereumAddress,
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
		Use:   "schema-doc [doc] [private-key] [signing-algo]",
		Short: "Schema Document signature",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argSchemaDoc := args[0]
			argPrivateKey := args[1]
			argSigningAlgo := args[2]

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
			schemaDocBytes := schemaDoc.GetSignBytes()

			// Sign Schema Document
			var signature string
			switch argSigningAlgo {
			case "ed25519":
				signature, err = hidnodecli.GetEd25519Signature(argPrivateKey, schemaDocBytes)
				if err != nil {
					return err
				}
			case "secp256k1":
				signature, err = hidnodecli.GetSecp256k1Signature(argPrivateKey, schemaDocBytes)
				if err != nil {
					return err
				}
			case "recover-eth":
				signature, err = hidnodecli.GetEthRecoverySignature(argPrivateKey, schemaDocBytes)
				if err != nil {
					return err
				}
			case "bbs":
				signature, err = hidnodecli.GetBBSSignature(argPrivateKey, schemaDocBytes)
				if err != nil {
					return err
				}
			default:
				panic("recieved unsupported signing-algo. Supported algorithms are: ['ed25519', 'secp256k1', 'recover-eth', 'bbs']")
			}

			_, err = fmt.Fprintln(cmd.OutOrStdout(), signature)
			return err
		},
	}
	return cmd
}

func signCredStatusDocCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cred-status-doc [doc] [private-key] [signing-algo]",
		Short: "Credential Status Document signature",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argCredStatusDoc := args[0]
			argPrivateKey := args[1]
			argSigningAlgo := args[2]

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
			credStatusDocBytes := credStatusDoc.GetSignBytes()

			// Sign Credential Status Document
			var signature string
			switch argSigningAlgo {
			case "ed25519":
				signature, err = hidnodecli.GetEd25519Signature(argPrivateKey, credStatusDocBytes)
				if err != nil {
					return err
				}
			case "secp256k1":
				signature, err = hidnodecli.GetSecp256k1Signature(argPrivateKey, credStatusDocBytes)
				if err != nil {
					return err
				}
			case "recover-eth":
				signature, err = hidnodecli.GetEthRecoverySignature(argPrivateKey, credStatusDocBytes)
				if err != nil {
					return err
				}
			case "bbs":
				signature, err = hidnodecli.GetBBSSignature(argPrivateKey, credStatusDocBytes)
				if err != nil {
					return err
				}
			default:
				panic("recieved unsupported signing-algo. Supported algorithms are: ['ed25519', 'secp256k1', 'recover-eth', 'bbs']")
			}

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
