package keeper

import (
	"encoding/binary"

	"github.com/PavelTabacu/auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAssetCount get the total number of asset
func (k Keeper) GetAssetCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AssetCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAssetCount set the total number of asset
func (k Keeper) SetAssetCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AssetCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAsset appends a asset in the store with a new id and update the count
func (k Keeper) AppendAsset(
	ctx sdk.Context,
	asset types.Asset,
) uint64 {
	// Create the asset
	count := k.GetAssetCount(ctx)

	// Set the ID of the appended value
	asset.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	appendedValue := k.cdc.MustMarshal(&asset)
	store.Set(GetAssetIDBytes(asset.Id), appendedValue)

	// Update asset count
	k.SetAssetCount(ctx, count+1)

	return count
}

// SetAsset set a specific asset in the store
func (k Keeper) SetAsset(ctx sdk.Context, asset types.Asset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	b := k.cdc.MustMarshal(&asset)
	store.Set(GetAssetIDBytes(asset.Id), b)
}

// GetAsset returns a asset from its id
func (k Keeper) GetAsset(ctx sdk.Context, id uint64) (val types.Asset, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	b := store.Get(GetAssetIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAsset removes a asset from the store
func (k Keeper) RemoveAsset(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	store.Delete(GetAssetIDBytes(id))
}

// GetAllAsset returns all asset
func (k Keeper) GetAllAsset(ctx sdk.Context) (list []types.Asset) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AssetKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Asset
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAssetIDBytes returns the byte representation of the ID
func GetAssetIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAssetIDFromBytes returns ID in uint64 format from a byte array
func GetAssetIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
