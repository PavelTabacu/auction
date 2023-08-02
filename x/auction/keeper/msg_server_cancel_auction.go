package keeper

import (
	"context"

	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CancelAuction(goCtx context.Context, msg *types.MsgCancelAuction) (*types.MsgCancelAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	auction, found := k.Keeper.GetAuction(ctx, msg.AuctionId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNotFound, "Auction not found!")
	}
	if auction.Seller != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrInvalidUser, "Must be the Seller of the auction, auction Seller:%s", auction.Seller)
	}
	if auction.IsCancel == true {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOperation, "Auction is already cancelled!")
	}
	auction.IsCancel = true
	k.Keeper.SetAuction(ctx, auction)

	return &types.MsgCancelAuctionResponse{}, nil
}
