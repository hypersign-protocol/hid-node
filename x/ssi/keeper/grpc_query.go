package keeper

import (
	"github.com/hypersign-protocol/vid-node/x/ssi/types"
)

var _ types.QueryServer = Keeper{}
