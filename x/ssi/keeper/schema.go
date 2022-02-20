package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/utils"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k Keeper) GetSchemaCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey and SchemaCountKey (which is "Schema-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaCountKey))
	// Convert the SchemaCountKey to bytes
	byteKey := []byte(types.SchemaCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

// Check whether the given Schema already exists in the store
func (k Keeper) HasSchema(ctx sdk.Context, id string) bool {
	extractedID := utils.ExtractIDFromSchema(id)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SchemaKey))
	return store.Has(utils.UnsafeStrToBytes(extractedID))
}

func (k Keeper) SetSchemaCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey and SchemaCountKey (which is "Schema-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaCountKey))
	// Convert the SchemaCountKey to bytes
	byteKey := []byte(types.SchemaCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of Schema-count- to count
	store.Set(byteKey, bz)
}

func (k Keeper) AppendSchema(ctx sdk.Context, schema types.Schema) uint64 {
	// Get the current number of Schemas in the store
	count := k.GetSchemaCount(ctx)
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaKey))
	// Marshal the Schema into bytes
	schemaBytes := k.cdc.MustMarshal(&schema)
	// Strip the id part from schema.ID
	schemaID := utils.ExtractIDFromSchema(schema.Id)
	store.Set(utils.UnsafeStrToBytes(schemaID), schemaBytes)
	// Update the Schema count
	k.SetSchemaCount(ctx, count+1)
	return count
}
