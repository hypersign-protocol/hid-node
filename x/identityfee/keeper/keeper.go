package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/hypersign-protocol/hid-node/x/identityfee/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		paramSpace paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
) *Keeper {
	return &Keeper{
		cdc:        cdc,
		paramSpace: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
