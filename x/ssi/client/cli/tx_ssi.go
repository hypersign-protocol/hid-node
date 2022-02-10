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
			var didDoc types.DidDocStruct
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
