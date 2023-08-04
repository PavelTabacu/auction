package keeper

import (
	"context"
	"time"

	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ExecuteAuction(ctx sdk.Context, auction types.Auction) (err error) {
	if auction.IsCancel {
		return sdkerrors.Wrapf(types.ErrInvalidAuction, "The auction is not valid")
	}
	if len(auction.BidIds) == 0 {
		return sdkerrors.Wrapf(types.ErrNotFound, "Bidders not found")
	}
	maxBid, _ := k.GetBid(ctx, auction.BidIds[0])
	for i := 0; i < len(auction.BidIds); i++ {
		bid, _ := k.GetBid(ctx, auction.BidIds[i])
		if bid.IsCancel == false && maxBid.Price.IsLTE(bid.Price) {
			maxBid = bid
		}
	}
	sellerAddr, _ := sdk.AccAddressFromBech32(auction.Seller)
	buyerAddr, _ := sdk.AccAddressFromBech32(maxBid.Buyer)
	price := [](sdk.Coin){auction.CurrentPrice}
	sdkError := k.bankKeeper.SendCoins(ctx, buyerAddr, sellerAddr, price)
	if sdkError != nil {
		return sdkerrors.Wrapf(sdkError, "coin send error")
	}

	asset, _ := k.GetAsset(ctx, auction.AssetId)
	buyer, _ := k.GetUsers(ctx, maxBid.Buyer)
	buyer.AssetIds = append(buyer.AssetIds, asset.Id)
	k.SetUsers(ctx, buyer)
	seller, _ := k.GetUsers(ctx, auction.Seller)
	assetIds := seller.AssetIds
	if len(assetIds) == 0 {
		panic("The number of assetIds for Seller is not zero")
	}
	if len(assetIds) == 1 {
		seller.AssetIds = [](uint64){}
	} else {
		i := 0
		for ; i < len(assetIds); i++ {
			ownedAsset, _ := k.GetAsset(ctx, assetIds[i])
			if asset.Id == ownedAsset.Id {
				break
			}
		}
		seller.AssetIds = append(assetIds[:i], assetIds[i+1:]...)
	}
	k.SetUsers(ctx, seller)
	asset.Owner = maxBid.Buyer
	asset.IsAuction = false
	k.SetAsset(ctx, asset)

	auction.IsCancel = false
	k.SetAuction(ctx, auction)

	return nil
}

func (k Keeper) ExecutionAuctions(goCtx context.Context) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	allAuctions := k.GetAllAuction(ctx)
	for i := 0; i < len(allAuctions); i++ {
		deadline, _ := time.Parse(types.DeadlineLayout, allAuctions[i].Deadline)
		if deadline.Before(ctx.BlockTime()) {
			k.ExecuteAuction(ctx, allAuctions[i])
		}
	}
}
