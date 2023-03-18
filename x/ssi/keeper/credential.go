package keeper

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k Keeper) RegisterCredentialStatusInState(ctx sdk.Context, cred *types.Credential) uint64 {
	count := k.GetCredentialStatusCount(ctx)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredKey))

	id := cred.GetClaim().GetId()
	credBytes := k.cdc.MustMarshal(cred)

	store.Set([]byte(id), credBytes)
	k.SetCredentialStatusCount(ctx, count+1)
	return count
}

func (k Keeper) GetCredentialStatusFromState(ctx *sdk.Context, id string) (*types.Credential, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredKey))

	var cred types.Credential
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

// Invoked during BeginBlock, it checks the block time against credential
// expiration date. If the expiration date is past current block time, the status
// of the credential(s) are set to `Expired`.
func (k Keeper) SetCredentialStatusToExpired(ctx sdk.Context) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var cred types.Credential
		if err := k.cdc.Unmarshal(iterator.Value(), &cred); err != nil {
			return err
		}

		currentBlockTime := ctx.BlockTime()
		expirationDate, err := time.Parse(time.RFC3339, cred.GetExpirationDate())
		if err != nil {
			return err
		}

		// Set the Credential Status to Expired
		if currentBlockTime.After(expirationDate) {
			cred.Claim.CurrentStatus = "Expired"
			cred.Claim.StatusReason = "Credential Expired"
			cred.Proof.Updated = currentBlockTime.Format(time.RFC3339)

			updatedCredBytes := k.cdc.MustMarshal(&cred)
			store.Set([]byte(cred.Claim.Id), updatedCredBytes)
		}
	}
	return nil
}
