package auction

import (
	"github.com/PavelTabacu/auction/x/auction/keeper"
	"github.com/PavelTabacu/auction/x/auction/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the asset
	for _, elem := range genState.AssetList {
		k.SetAsset(ctx, elem)
	}

	// Set asset count
	k.SetAssetCount(ctx, genState.AssetCount)
	// Set all the auction
	for _, elem := range genState.AuctionList {
		k.SetAuction(ctx, elem)
	}

	// Set auction count
	k.SetAuctionCount(ctx, genState.AuctionCount)
	// Set all the bid
	for _, elem := range genState.BidList {
		k.SetBid(ctx, elem)
	}

	// Set bid count
	k.SetBidCount(ctx, genState.BidCount)
	// Set all the users
	for _, elem := range genState.UsersList {
		k.SetUsers(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AssetList = k.GetAllAsset(ctx)
	genesis.AssetCount = k.GetAssetCount(ctx)
	genesis.AuctionList = k.GetAllAuction(ctx)
	genesis.AuctionCount = k.GetAuctionCount(ctx)
	genesis.BidList = k.GetAllBid(ctx)
	genesis.BidCount = k.GetBidCount(ctx)
	genesis.UsersList = k.GetAllUsers(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
