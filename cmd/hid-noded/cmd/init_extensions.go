package cmd

import (
	"github.com/cosmos/cosmos-sdk/client"
	cosmcfg "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/spf13/cobra"
	tmcfg "github.com/tendermint/tendermint/config"
)

func extendInit(initCmd *cobra.Command) *cobra.Command {
	baseRunE := initCmd.RunE

	initCmd.RunE = func(cmd *cobra.Command, args []string) error {
		err := baseRunE(cmd, args)
		if err != nil {
			return err
		}

		err = applyConfigDefaults(cmd)
		if err != nil {
			return err
		}

		return nil
	}

	return initCmd
}

func applyConfigDefaults(cmd *cobra.Command) error {
	clientCtx := client.GetClientContextFromCmd(cmd)

	err := updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
		config.FastSync.Version = "v0"
		config.P2P.SendRate = 20000000
		config.P2P.RecvRate = 20000000
		config.P2P.MaxPacketMsgPayloadSize = 10240
	})
	if err != nil {
		return err
	}

	err = updateCosmConfig(clientCtx.HomeDir, func(config *cosmcfg.Config) {
		config.BaseConfig.MinGasPrices = "10uhid"
	})
	if err != nil {
		return err
	}

	return nil
}
