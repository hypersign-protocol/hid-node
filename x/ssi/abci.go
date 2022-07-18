package ssi

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/keeper"
)

// BeginBlocker is called at the beginning of every block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	// Set all the credential status that have passed their expiration date
	// to Expired
	if err := k.SetCredentialStatusToExpired(ctx); err != nil {
		panic(err)
	}
}
