package keeper

import (
	"github.com/hypersign-protocol/hid-node/x/did/types"
)

var _ types.QueryServer = Keeper{}
