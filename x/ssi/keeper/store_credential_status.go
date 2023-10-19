package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k Keeper) SetCredentialStatusInState(ctx sdk.Context, cred *types.CredentialStatusState) {
	count := k.GetCredentialStatusCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredKey))

	id := cred.CredentialStatusDocument.Id
	credBytes := k.cdc.MustMarshal(cred)

	store.Set([]byte(id), credBytes)
	k.SetCredentialStatusCount(ctx, count+1)
}

func (k Keeper) GetCredentialStatusFromState(ctx *sdk.Context, id string) (*types.CredentialStatusState, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredKey))

	var cred types.CredentialStatusState
	var bytes = store.Get([]byte(id))
	if len(bytes) == 0 {
		return nil, fmt.Errorf("credential status document %s not found", id)
	}

	if err := k.cdc.Unmarshal(bytes, &cred); err != nil {
		return nil, fmt.Errorf("internal: unable to unmarshal credentialStatus id %s from state", id)
	}

	return &cred, nil
}

// Check whether the given Cred is already present in the store
func (k Keeper) HasCredential(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredKey))
	return store.Has([]byte(id))
}

func (k Keeper) GetCredentialStatusCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredCountKey))
	byteKey := []byte(types.CredCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetCredentialStatusCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredCountKey))
	byteKey := []byte(types.CredCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}
