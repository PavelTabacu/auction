package types_test

import (
	"testing"

	"github.com/PavelTabacu/auction/x/auction/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated asset",
			genState: &types.GenesisState{
				AssetList: []types.Asset{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid asset count",
			genState: &types.GenesisState{
				AssetList: []types.Asset{
					{
						Id: 1,
					},
				},
				AssetCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated auction",
			genState: &types.GenesisState{
				AuctionList: []types.Auction{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid auction count",
			genState: &types.GenesisState{
				AuctionList: []types.Auction{
					{
						Id: 1,
					},
				},
				AuctionCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated bid",
			genState: &types.GenesisState{
				BidList: []types.Bid{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid bid count",
			genState: &types.GenesisState{
				BidList: []types.Bid{
					{
						Id: 1,
					},
				},
				BidCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated users",
			genState: &types.GenesisState{
				UsersList: []types.Users{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
