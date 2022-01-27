package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/utils"
	"github.com/hypersign-protocol/hid-node/x/did/types"
)

func (k Keeper) GetSchemaCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaCountKey))
	// Convert the PostCountKey to bytes
	byteKey := []byte(types.SchemaCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found (for example, it's the first post)
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetSchemaCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey (which is "blog") and PostCountKey (which is "Post-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaCountKey))
	// Convert the PostCountKey to bytes
	byteKey := []byte(types.SchemaCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of Post-count- to count
	store.Set(byteKey, bz)
}

func (k Keeper) AppendSchema(ctx sdk.Context, schema types.Schema) uint64 {
	// Get the current number of posts in the store
	count := k.GetSchemaCount(ctx)
	// Assign an ID to the post based on the number of posts in the store
	schema.Id = count
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaKey))

	// TODO: Follow x/did/keeper/didSpec.go
	store.Set(utils.UnsafeStrToBytes(schema.SchemaID), utils.UnsafeStrToBytes(schema.SchemaStr))
	// Update the post count
	k.SetSchemaCount(ctx, count+1)
	return count
}
