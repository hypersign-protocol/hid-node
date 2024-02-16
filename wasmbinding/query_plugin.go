package wasmbinding

import (
	ssiKeeper "github.com/hypersign-protocol/hid-node/x/ssi/keeper"
)

type QueryPlugin struct {
	ssiKeeper *ssiKeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(ssiKeeper *ssiKeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		ssiKeeper: ssiKeeper,
	}
}
