package keeper

import (
	"context"
	"time"

	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ExecuteAuction(ctx sdk.Context, auction types.Auction) (err error) {
	deadline, _ := time.Parse(types.DeadlineLayout, auction.Deadline)
	if deadline.Before(ctx.BlockTime()) {
		if len(auction.BidIds) == 0 {
			return sdkerrors.Wrapf(types.ErrNotFound, "Bidders not found")
		}
		maxBid, _ := k.GetBid(ctx, auction.BidIds[0])
		for i := 0; i < len(auction.BidIds); i++ {
			bid, _ := k.GetBid(ctx, auction.BidIds[i])
			if bid.IsCancel == false && !bid.Price.IsLTE(maxBid.Price) {
				maxBid = bid
			}
		}
		sellerAddr, _ := sdk.AccAddressFromBech32(auction.Seller)
		bidderAddr, _ := sdk.AccAddressFromBech32(maxBid.Buyer)
		price := [](sdk.Coin){auction.CurrentPrice}
		sdkError := k.bankKeeper.SendCoins(ctx, bidderAddr, sellerAddr, price)
		if sdkError != nil {
			return sdkerrors.Wrapf(sdkError, "")
		}

		car, _ := k.GetAsset(ctx, auction.Id)
		bidder, _ := k.GetUsers(ctx, maxBid.Buyer)
		bidder.TokenIds = append(bidder.TokenIds, car.Id)
		k.SetUsers(ctx, bidder)
		seller, _ := k.GetUsers(ctx, auction.Seller)
		carList := seller.TokenIds
		i := 0
		for ; i < len(carList); i++ {
			ownedCar, _ := k.GetAsset(ctx, carList[i])
			if car.Id == ownedCar.Id {
				break
			}
		}
		seller.TokenIds = append(carList[:i], carList[i+1:]...)
		k.SetUsers(ctx, seller)
		car.Creator = maxBid.Buyer
		car.IsAuction = false
		k.SetAsset(ctx, car)

		auction.IsCancel = false
		k.SetAuction(ctx, auction)
		return
	} else {
		return sdkerrors.Wrapf(types.ErrInvalidOperation, "Auction is already ended")
	}
}

func (k Keeper) ExecutionAuctions(goCtx context.Context) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	allAuctions := k.GetAllAuction(ctx)
	for i := 0; i < len(allAuctions); i++ {
		k.ExecuteAuction(ctx, allAuctions[i])
	}

}
