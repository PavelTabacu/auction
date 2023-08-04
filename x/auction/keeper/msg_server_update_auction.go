package keeper

import (
	"context"
	"time"

	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateAuction(goCtx context.Context, msg *types.MsgUpdateAuction) (*types.MsgUpdateAuctionResponse, error) {
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
	deadline, _ := time.Parse(types.DeadlineLayout, msg.Deadline)
	if deadline.Before(ctx.BlockTime()) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOperation, "invalid Deadline")
	}

	auction.IsCancel = false
	auction.CurrentPrice = msg.StartPrice
	auction.Deadline = msg.Deadline
	k.Keeper.SetAuction(ctx, auction)

	return &types.MsgUpdateAuctionResponse{}, nil
}
