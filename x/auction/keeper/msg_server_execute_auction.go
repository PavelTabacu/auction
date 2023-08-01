package keeper

import (
	"context"

	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ExecuteAuction(goCtx context.Context, msg *types.MsgExecuteAuction) (*types.MsgExecuteAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgExecuteAuctionResponse{}, nil
}
