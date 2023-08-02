package keeper

import (
	"context"

	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CancelBid(goCtx context.Context, msg *types.MsgCancelBid) (*types.MsgCancelBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	bid, found := k.Keeper.GetBid(ctx, msg.BidId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNotFound, "Auction not found!")
	}
	if bid.Buyer == msg.Creator {
		bid.IsCancel = true
		k.Keeper.SetBid(ctx, bid)
		return &types.MsgCancelBidResponse{}, nil
	}

	return &types.MsgCancelBidResponse{}, nil
}
