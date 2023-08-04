package keeper

import (
	"context"

	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ExecuteAuction(goCtx context.Context, msg *types.MsgExecuteAuction) (*types.MsgExecuteAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	auction, found := k.Keeper.GetAuction(ctx, msg.AuctionId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrNotFound, "Auction not found!")
	}
	if auction.Seller != msg.Creator {
		return nil, sdkerrors.Wrapf(types.ErrInvalidUser, "Must be the poster of the auction, auction poster:%s", auction.Seller)
	}

	err := k.Keeper.ExecuteAuction(ctx, auction)
	if err != nil {
		return nil, err
	}
	return &types.MsgExecuteAuctionResponse{}, nil
}
