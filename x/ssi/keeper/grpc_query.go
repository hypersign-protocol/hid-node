package keeper

import (
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

var _ types.QueryServer = Keeper{}
