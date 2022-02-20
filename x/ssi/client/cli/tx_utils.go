package cli

import (
	"crypto/ed25519"
	"encoding/base64"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

func getVerKey(cmd *cobra.Command, clientCtx client.Context) (ed25519.PrivateKey, error) {
	// Get the verification key from --ver-key flag
	verKeyPrivBase64, err := cmd.Flags().GetString(VerKeyFlag)
	if err != nil {
		return nil, err
	}

	// Decode key into bytes
	verKeyPrivBytes, err := base64.StdEncoding.DecodeString(verKeyPrivBase64)
	if err != nil {
		return nil, err
	}

	return verKeyPrivBytes, nil
}
