package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hypersign-protocol/hid-node/x/ssi/types"
)

func (k Keeper) RegisterCred(ctx sdk.Context, cred *types.Credential) uint64 {
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CredKey))
	// Get ID from DID Doc
	id := cred.GetClaim().GetId()
	// Marshal the DID into bytes
	credBytes := k.cdc.MustMarshal(cred)
	store.Set([]byte(id), credBytes)
	return 1
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

// func (k Keeper) GetCredentialsFromStore(ctx sdk.Context, querySchemaStr string) []*types.Schema {
// 	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaKey))
// 	var schemas []*types.Schema
// 	iterator := sdk.KVStorePrefixIterator(store, []byte{})

// 	for ; iterator.Valid(); iterator.Next() {
// 		var schema types.Schema
// 		k.cdc.MustUnmarshal(iterator.Value(), &schema)

// 		if querySchemaStr == schema.Id[0:len(schema.Id)-12] || querySchemaStr == schema.Id {
// 			schemas = append(schemas, &schema)
// 		}
// 	}

// 	return schemas
// }
