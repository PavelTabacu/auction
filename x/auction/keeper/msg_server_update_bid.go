package keeper

import (
	"context"
	"time"

	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateBid(goCtx context.Context, msg *types.MsgUpdateBid) (*types.MsgUpdateBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	auction, found := k.Keeper.GetAuction(ctx, msg.AuctionId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNotFound, "Auction not found!")
	}
	if auction.IsCancel {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOperation, "The auction is already canceld")
	}
	deadline, _ := time.Parse(types.DeadlineLayout, auction.Deadline)
	if deadline.Before(ctx.BlockTime()) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOperation, "The auction is already done")
	}

	for i := len(auction.BidIds) - 1; i >= 0; i-- {
		bid, _ := k.Keeper.GetBid(ctx, auction.BidIds[i])
		if bid.Buyer == msg.Creator {
			if !bid.Price.IsGTE(msg.Price) {
				bid.Price = msg.Price
				k.Keeper.SetBid(ctx, bid)

				break
			}
		}
	}

	return &types.MsgUpdateBidResponse{}, nil
}
