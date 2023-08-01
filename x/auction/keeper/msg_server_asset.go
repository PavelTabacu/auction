package keeper

import (
	"context"
	"fmt"

	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAsset(goCtx context.Context, msg *types.MsgCreateAsset) (*types.MsgCreateAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var asset = types.Asset{
		Creator:     msg.Creator,
		TokenId:     msg.TokenId,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Owner:       msg.Owner,
		Description: msg.Description,
		Url:         msg.Url,
		IsAuction:   msg.IsAuction,
	}

	id := k.AppendAsset(
		ctx,
		asset,
	)

	return &types.MsgCreateAssetResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateAsset(goCtx context.Context, msg *types.MsgUpdateAsset) (*types.MsgUpdateAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var asset = types.Asset{
		Creator:     msg.Creator,
		Id:          msg.Id,
		TokenId:     msg.TokenId,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Owner:       msg.Owner,
		Description: msg.Description,
		Url:         msg.Url,
		IsAuction:   msg.IsAuction,
	}

	// Checks that the element exists
	val, found := k.GetAsset(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAsset(ctx, asset)

	return &types.MsgUpdateAssetResponse{}, nil
}

func (k msgServer) DeleteAsset(goCtx context.Context, msg *types.MsgDeleteAsset) (*types.MsgDeleteAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetAsset(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAsset(ctx, msg.Id)

	return &types.MsgDeleteAssetResponse{}, nil
}
