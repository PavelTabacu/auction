package keeper

import (
	"github.com/PavelTabacu/auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetUsers set a specific users in the store from its index
func (k Keeper) SetUsers(ctx sdk.Context, users types.Users) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersKeyPrefix))
	b := k.cdc.MustMarshal(&users)
	store.Set(types.UsersKey(
		users.Index,
	), b)
}

// GetUsers returns a users from its index
func (k Keeper) GetUsers(
	ctx sdk.Context,
	index string,

) (val types.Users, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersKeyPrefix))

	b := store.Get(types.UsersKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUsers removes a users from the store
func (k Keeper) RemoveUsers(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersKeyPrefix))
	store.Delete(types.UsersKey(
		index,
	))
}

// GetAllUsers returns all users
func (k Keeper) GetAllUsers(ctx sdk.Context) (list []types.Users) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UsersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Users
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
