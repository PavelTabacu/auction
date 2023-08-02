package keeper

import (
	"context"

	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SendBid(goCtx context.Context, msg *types.MsgSendBid) (*types.MsgSendBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	auction, found := k.Keeper.GetAuction(ctx, msg.AuctionId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNotFound, "Auction not found!")
	}
	// if uint64(ctx.BlockHeight()) > auction.PostDate+auction.Duration {
	// 	return nil, sdkerrors.Wrapf(types.ErrInvalidOperation, "Auction is already ended!")
	// } else if auction.IsCancel == true {
	// 	return nil, sdkerrors.Wrapf(types.ErrInvalidOperation, "Auction is already cancelled!")
	// }
	if auction.CurrentPrice.IsGTE(msg.Price) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidPrice, "Lower than or equal to current price. Current Price: %d", auction.CurrentPrice)
	}
	if auction.Seller == msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrInvalidOperation, "Sellers cannot bid to their own auctions.")
	}
	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)
	addAmount := sdk.NewInt(1000)
	requiredPrice := auction.CurrentPrice.AddAmount(addAmount)
	if k.bankKeeper.GetBalance(ctx, buyer, "stake").IsLT(requiredPrice) {
		return nil, sdkerrors.Wrapf(types.ErrInvalidPrice, "Balance lower than required price %v", requiredPrice)
	}

	bid := types.Bid{
		Buyer:     msg.Creator,
		AuctionId: msg.AuctionId,
		Price:     msg.Price,
		IsCancel:  false,
	}

	bidId := k.Keeper.AppendBid(ctx, bid)

	auction.BidIds = append(auction.BidIds, bidId)
	auction.CurrentPrice = msg.Price
	k.Keeper.SetAuction(ctx, auction)

	user, found := k.Keeper.GetUsers(ctx, msg.Creator)
	if !found {
		user.Index = msg.Creator
	}
	user.BidIds = append(user.BidIds, bidId)
	k.Keeper.SetUsers(ctx, user)

	return &types.MsgSendBidResponse{}, nil
}
