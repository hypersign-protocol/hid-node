package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k Keeper) RegisterCred(ctx sdk.Context, cred *types.Credential) uint64 {
	// Get Cred count
	count := k.GetCredentialCount(ctx)
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredKey))
	// Get ID from DID Doc
	id := cred.GetClaim().GetId()
	// Marshal the DID into bytes
	credBytes := k.cdc.MustMarshal(cred)
	store.Set([]byte(id), credBytes)
	k.SetCredentialCount(ctx, count+1)
	return count
}

func (k Keeper) GetCredential(ctx *sdk.Context, id string) (*types.Credential, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredKey))

	var cred types.Credential
	var bytes = store.Get([]byte(id))
	if err := k.cdc.Unmarshal(bytes, &cred); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, err.Error())
	}

	return &cred, nil
}

// Check whether the given Cred is already present in the store
func (k Keeper) HasCredential(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CredKey))
	return store.Has([]byte(id))
}

func (k Keeper) GetCredentialCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredCountKey))
	byteKey := []byte(types.CredCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetCredentialCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredCountKey))
	byteKey := []byte(types.CredCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}
