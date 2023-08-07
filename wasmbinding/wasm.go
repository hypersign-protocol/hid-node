package wasmbinding

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	ssiKeeper "github.com/hypersign-protocol/hid-node/x/ssi/keeper"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)
// RegisterCustomPlugins returns wasmkeeper.Option that we can use to connect handlers for implemented custom queries to the App
func RegisterCustomPlugins(ssiKeeper *ssiKeeper.Keeper) []wasm.Option {
	// For Query
	queryPlugin := NewQueryPlugin(ssiKeeper)

	queryPluginWasmOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(queryPlugin),
	})

	// For Execute
	messagePluginOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(ssiKeeper),
	)

	return []wasm.Option{
		queryPluginWasmOpt,
		messagePluginOpt,
	}
}
