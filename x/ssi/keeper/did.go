package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/utils"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k Keeper) GetDidCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey and DidCountKey (which is "Did-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DidCountKey))
	// Convert the DidCountKey to bytes
	byteKey := []byte(types.DidCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

// Check whether the given DID is already present in the store
func (k Keeper) HasDid(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	return store.Has(utils.UnsafeStrToBytes(id))
}

// Retrieves the DID from the store
func (k Keeper) GetDid(ctx *sdk.Context, id string) (*types.Did, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))

	if !k.HasDid(*ctx, id) {
		return nil, sdkerrors.ErrNotFound
	}

	var value types.Did
	var bytes = store.Get([]byte(id))
	if err := k.cdc.Unmarshal(bytes, &value); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidType, err.Error())
	}

	return &value, nil
}

func (k Keeper) SetDidCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey and SchemaCountKey (which is "Did-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DidCountKey))
	// Convert the DidCountKey to bytes
	byteKey := []byte(types.DidCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of Did-count- to count
	store.Set(byteKey, bz)
}

// SetDid set a specific did in the store
func (k Keeper) SetDid(ctx sdk.Context, did types.Did) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	b := k.cdc.MustMarshal(&did)
	store.Set([]byte(did.Id), b)
	return nil
}

func (k Keeper) AppendDID(ctx sdk.Context, didSpec types.Did) uint64 {
	// Get the current number of DIDs in the store
	count := k.GetDidCount(ctx)
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DidKey))
	// Marshal the DID into bytes
	didDocString := k.cdc.MustMarshal(&didSpec)
	store.Set(utils.UnsafeStrToBytes(didSpec.Id), didDocString)
	// Update the Did count
	k.SetDidCount(ctx, count+1)
	return count
}
