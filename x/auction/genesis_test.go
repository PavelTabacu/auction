package auction_test

import (
	"testing"

	keepertest "github.com/PavelTabacu/auction/testutil/keeper"
	"github.com/PavelTabacu/auction/testutil/nullify"
	"github.com/PavelTabacu/auction/x/auction"
	"github.com/PavelTabacu/auction/x/auction/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AssetList: []types.Asset{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AssetCount: 2,
		AuctionList: []types.Auction{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AuctionCount: 2,
		BidList: []types.Bid{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		BidCount: 2,
		UsersList: []types.Users{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AuctionKeeper(t)
	auction.InitGenesis(ctx, *k, genesisState)
	got := auction.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AssetList, got.AssetList)
	require.Equal(t, genesisState.AssetCount, got.AssetCount)
	require.ElementsMatch(t, genesisState.AuctionList, got.AuctionList)
	require.Equal(t, genesisState.AuctionCount, got.AuctionCount)
	require.ElementsMatch(t, genesisState.BidList, got.BidList)
	require.Equal(t, genesisState.BidCount, got.BidCount)
	require.ElementsMatch(t, genesisState.UsersList, got.UsersList)
	// this line is used by starport scaffolding # genesis/test/assert
}
