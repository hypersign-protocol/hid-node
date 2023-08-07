package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/utils"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// Set the Chain namespace
func (k Keeper) SetChainNamespace(ctx *sdk.Context, namespace string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ChainNamespaceKey))
	byteKey := []byte(types.ChainNamespaceKey)
	store.Set(byteKey, []byte(namespace))
}

// Get the Chain namespace
func (k Keeper) GetChainNamespace(ctx *sdk.Context) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ChainNamespaceKey))
	byteKey := []byte(types.ChainNamespaceKey)
	bz := store.Get(byteKey)
	return string(bz)
}

// Get the count of registered Did Documents
func (k Keeper) GetDidCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DidCountKey))
	byteKey := []byte(types.DidCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// Check whether the Did document exist in the store
func (k Keeper) HasDid(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	return store.Has(utils.UnsafeStrToBytes(id))
}

// Retrieves the DID from the store
func (k Keeper) GetDidDocumentState(ctx *sdk.Context, id string) (*types.DidDocumentState, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))

	var didDocState types.DidDocumentState
	var bytes = store.Get([]byte(id))

	if len(bytes) == 0 {
		return nil, fmt.Errorf("DID Document %s not found", id)
	}

	if err := k.cdc.Unmarshal(bytes, &didDocState); err != nil {
		return nil, fmt.Errorf("internal: unable to unmarshal didDocBytes of id %s", id)
	}

	return &didDocState, nil
}

// Sets the Did Document Count
func (k Keeper) SetDidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DidCountKey))
	byteKey := []byte(types.DidCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// Updates an existing Did document present in the store
func (k Keeper) UpdateDidDocumentInStore(ctx sdk.Context, didDoc types.DidDocumentState) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	b := k.cdc.MustMarshal(&didDoc)
	store.Set([]byte(didDoc.DidDocument.Id), b)
	return nil
}

// Creates record for a new DID Document
func (k Keeper) RegisterDidDocumentInStore(ctx sdk.Context, didDoc *types.DidDocumentState) uint64 {
	didCount := k.GetDidCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))

	idBytes := []byte(didDoc.GetDidDocument().GetId())
	didDocBytes := k.cdc.MustMarshal(didDoc)

	store.Set(idBytes, didDocBytes)
	k.SetDidCount(ctx, didCount+1)
	return didCount
}

// Set the BlockchainAccountId in Store
func (k Keeper) SetBlockchainAddressInStore(ctx *sdk.Context, blockchainAccountId string, didId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BlockchainAccountIdStoreKey))
	store.Set([]byte(blockchainAccountId), []byte(didId))
}

// Get the BlockchainAccountId from Store
func (k Keeper) GetBlockchainAddressFromStore(ctx *sdk.Context, blockchainAccountId string) []byte {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BlockchainAccountIdStoreKey))
	return store.Get([]byte(blockchainAccountId))
}

// Remove the BlockchainAccountId from Store
func (k Keeper) RemoveBlockchainAddressInStore(ctx *sdk.Context, blockchainAccountId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BlockchainAccountIdStoreKey))
	store.Delete([]byte(blockchainAccountId))
}
