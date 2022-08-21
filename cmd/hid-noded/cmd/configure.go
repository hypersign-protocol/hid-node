package cmd

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	cosmcfg "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
	tmcfg "github.com/tendermint/tendermint/config"
)

// configureCmd returns configure cobra Command.
func configureCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Adjust node parameters",
	}

	cmd.AddCommand(
		minGasPricesCmd(defaultNodeHome),
		p2pCmd(defaultNodeHome),
		rpcLaddrCmd(defaultNodeHome),
		createEmptyBlocksCmd(defaultNodeHome),
		fastsyncVersionCmd(defaultNodeHome))

	return cmd
}

// p2pCmd returns configure cobra Command.
func p2pCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "p2p",
		Short: "Adjust peer-to-peer (p2p) parameters",
	}

	cmd.AddCommand(
		seedModeCmd(defaultNodeHome),
		seedsCmd(defaultNodeHome),
		externalAddressCmd(defaultNodeHome),
		persistentPeersCmd(defaultNodeHome),
		sendRateCmd(defaultNodeHome),
		recvRateCmd(defaultNodeHome),
		maxPacketMsgPayloadSizeCmd(defaultNodeHome),
		p2pLaddrCmd(defaultNodeHome))

	return cmd
}

// minGasPricesCmd returns configuration cobra Command.
func minGasPricesCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "min-gas-prices [value]",
		Short: "The minimum gas prices a validator is willing to accept (default  \"10uhid\")",
		Example: "min-gas-prices 10uhid",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			return updateCosmConfig(clientCtx.HomeDir, func(config *cosmcfg.Config) {
				config.MinGasPrices = args[0]
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

// seedModeCmd returns configuration cobra Command.
func seedModeCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "seed-mode (true|false)",
		Short: "Seed mode, in which node constantly crawls the network and looks for peers.",
		Long: "Seed mode, in which node constantly crawls the network and looks for peers. If another node asks " +
			"it for addresses, it responds and disconnects.",
		Example: "seed-mode true",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			value, err := strconv.ParseBool(args[0])
			if err != nil {
				return errors.Wrap(err, "can't parse seed mode")
			}

			return updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
				config.P2P.SeedMode = value
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

// seedsCmd returns configuration cobra Command.
func seedsCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "seeds [value]",
		Short: "Comma separated list of seed nodes to connect to in <node-id@ip-address-or-dns-name:port> format",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			return updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
				config.P2P.Seeds = args[0]
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

// externalAddressCmd returns configuration cobra Command.
func externalAddressCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "external-address [value]",
		Short: "Address to advertise to peers for them to connect to the node",
		Long: "Address to advertise to peers for them to connect to the node. If empty, the node will use the same " +
			"port as the laddr, and will attach it to the listener or use UPnP to figure out the address. IP " +
			"address/DNS name P2P port are required.",
		Example: "external-address 159.89.10.97:26656",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			return updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
				config.P2P.ExternalAddress = args[0]
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

// persistentPeersCmd returns configuration cobra Command.
func persistentPeersCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "persistent-peers [value]",
		Short: "Comma separated list of nodes to keep persistent connections to in " +
			"<node-id@ip-address-or-dns-name:p2p-port> format",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			return updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
				config.P2P.PersistentPeers = args[0]
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

// sendRateCmd returns configuration cobra Command.
func sendRateCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "send-rate [value]",
		Short:   "Rate at which packets can be sent, in bytes/second",
		Example: "send-rate 20000000",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			value, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return errors.Wrap(err, "can't parse send rate")
			}

			return updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
				config.P2P.SendRate = value
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

// recvRateCmd returns configuration cobra Command.
func recvRateCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "recv-rate [value]",
		Short:   "Rate at which packets can be received, in bytes/second",
		Example: "recv-rate 20000000",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			value, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return errors.Wrap(err, "can't parse recv rate")
			}

			return updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
				config.P2P.RecvRate = value
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

// maxPacketMsgPayloadSizeCmd returns configuration cobra Command.
func maxPacketMsgPayloadSizeCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "max-packet-msg-payload-size [value]",
		Short:   "Maximum size of a message packet payload, in bytes",
		Example: "max-packet-msg-payload-size 10240",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			value, err := strconv.ParseInt(args[0], 10, 32)
			if err != nil {
				return errors.Wrap(err, "can't parse max-packet-msg-payload-size")
			}

			return updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
				config.P2P.MaxPacketMsgPayloadSize = int(value)
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

// p2pLaddrCmd returns configuration cobra Command.
func p2pLaddrCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "laddr [value]",
		Short:   "Address to listen for incoming connections",
		Example: "laddr \"tcp://0.0.0.0:26656\"",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			return updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
				config.P2P.ListenAddress = args[0]
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

// createEmptyBlocksCmd returns configuration cobra Command.
func createEmptyBlocksCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-empty-blocks [value]",
		Short:   "EmptyBlocks mode (true|false)",
		Example: "create-empty-blocks false",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			value, err := strconv.ParseBool(args[0])
			if err != nil {
				return errors.Wrap(err, "can't parse create-empty-blocks")
			}

			return updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
				config.Consensus.CreateEmptyBlocks = value
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

// createEmptyBlocksCmd returns configuration cobra Command.
func rpcLaddrCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rpc-laddr [value]",
		Short:   "TCP or UNIX socket address for the RPC server to listen on",
		Example: "rpc-laddr \"tcp://127.0.0.1:26657\"",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			return updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
				config.RPC.ListenAddress = args[0]
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}

// fastsyncVersionCmd returns configuration cobra Command.
func fastsyncVersionCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "fastsync-version [value]",
		Short:   "Fast Sync version to use (v0|v1|v2)",
		Example: "fastsync-version v2",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			return updateTmConfig(clientCtx.HomeDir, func(config *tmcfg.Config) {
				config.FastSync.Version = args[0]
			})
		},
	}

	cmd.Flags().String(flags.FlagHome, defaultNodeHome, "The application home directory")

	return cmd
}
