package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

// SetChainNamespace sets the Chain namespace in store
func (k Keeper) SetChainNamespace(ctx *sdk.Context, namespace string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ChainNamespaceKey))
	byteKey := []byte(types.ChainNamespaceKey)
	store.Set(byteKey, []byte(namespace))
}

// GetChainNamespace gets the Chain namespace from store
func (k Keeper) GetChainNamespace(ctx *sdk.Context) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.ChainNamespaceKey))
	byteKey := []byte(types.ChainNamespaceKey)
	bz := store.Get(byteKey)
	return string(bz)
}

// getDidDocumentCount gets the did document count from store
func (k Keeper) getDidDocumentCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DidCountKey))
	byteKey := []byte(types.DidCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// hasDidDocument checks whether did document exists in store
func (k Keeper) hasDidDocument(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	return store.Has([]byte(id))
}

// getDidDocumentState gets the did document from store
func (k Keeper) getDidDocumentState(ctx *sdk.Context, id string) (*types.DidDocumentState, error) {
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

// setDidCount sets the did document count in store
func (k Keeper) setDidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.DidCountKey))
	byteKey := []byte(types.DidCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// updateDidDocumentInStore updates existing did document in the store
func (k Keeper) updateDidDocumentInStore(ctx sdk.Context, didDoc types.DidDocumentState) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))
	b := k.cdc.MustMarshal(&didDoc)
	store.Set([]byte(didDoc.DidDocument.Id), b)
	return nil
}

// Creates record for a new DID Document
func (k Keeper) registerDidDocumentInStore(ctx sdk.Context, didDoc *types.DidDocumentState) {
	didCount := k.getDidDocumentCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DidKey))

	idBytes := []byte(didDoc.GetDidDocument().GetId())
	didDocBytes := k.cdc.MustMarshal(didDoc)

	store.Set(idBytes, didDocBytes)
	k.setDidCount(ctx, didCount+1)
}

// Set the BlockchainAccountId in Store
func (k Keeper) setBlockchainAddressInStore(ctx *sdk.Context, blockchainAccountId string, didId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BlockchainAccountIdStoreKey))
	store.Set([]byte(blockchainAccountId), []byte(didId))
}

// Get the BlockchainAccountId from Store
func (k Keeper) getBlockchainAddressFromStore(ctx *sdk.Context, blockchainAccountId string) []byte {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BlockchainAccountIdStoreKey))
	return store.Get([]byte(blockchainAccountId))
}

// Remove the BlockchainAccountId from Store
func (k Keeper) removeBlockchainAddressInStore(ctx *sdk.Context, blockchainAccountId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BlockchainAccountIdStoreKey))
	store.Delete([]byte(blockchainAccountId))
}
