package keeper

import (
	"context"

	"github.com/PavelTabacu/auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AssetAll(goCtx context.Context, req *types.QueryAllAssetRequest) (*types.QueryAllAssetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var assets []types.Asset
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	assetStore := prefix.NewStore(store, types.KeyPrefix(types.AssetKey))

	pageRes, err := query.Paginate(assetStore, req.Pagination, func(key []byte, value []byte) error {
		var asset types.Asset
		if err := k.cdc.Unmarshal(value, &asset); err != nil {
			return err
		}

		assets = append(assets, asset)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAssetResponse{Asset: assets, Pagination: pageRes}, nil
}

func (k Keeper) Asset(goCtx context.Context, req *types.QueryGetAssetRequest) (*types.QueryGetAssetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	asset, found := k.GetAsset(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAssetResponse{Asset: asset}, nil
}
