package keeper

import (
	"context"

	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) InitiateAuction(goCtx context.Context, msg *types.MsgInitiateAuction) (*types.MsgInitiateAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	token, found := k.Keeper.GetAsset(ctx, msg.AssetId)
	if !found {
		// panic("The token not found!")
		return nil, sdkerrors.Wrapf(types.ErrNotFound, "The token not found!")
	}
	if token.IsAuction == true {
		return nil, sdkerrors.Wrapf(types.ErrInvalidAuction, "The token is already in an auction!")
	}
	if token.Creator != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrInvalidUser, "Not the owner of the token, token owner:%s", token.Creator)
	}

	auction := types.Auction{
		Seller:       msg.Creator,
		AssetId:      msg.AssetId,
		StartPrice:   msg.StartPrice,
		CurrentPrice: msg.StartPrice,
		AuctionDate:  "2023/08/02",
		Deadline:     msg.Deadline,
		BidIds:       nil,
		IsCancel:     false,
	}
	auctionId := k.Keeper.AppendAuction(ctx, auction)

	token.IsAuction = true
	k.Keeper.SetAsset(ctx, token)

	user, found := k.Keeper.GetUsers(ctx, msg.Creator)
	if !found {
		user.Index = msg.Creator
	}
	user.AuctionIds = append(user.AuctionIds, auctionId)
	k.Keeper.SetUsers(ctx, user)

	return &types.MsgInitiateAuctionResponse{}, nil
}
