package keeper

import (
	"github.com/hypersign-protocol/hid-node/x/hidnode/types"
)

var _ types.QueryServer = Keeper{}
