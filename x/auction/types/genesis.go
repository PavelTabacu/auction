package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AssetList:   []Asset{},
		AuctionList: []Auction{},
		BidList:     []Bid{},
		UsersList:   []Users{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in asset
	assetIdMap := make(map[uint64]bool)
	assetCount := gs.GetAssetCount()
	for _, elem := range gs.AssetList {
		if _, ok := assetIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for asset")
		}
		if elem.Id >= assetCount {
			return fmt.Errorf("asset id should be lower or equal than the last id")
		}
		assetIdMap[elem.Id] = true
	}
	// Check for duplicated ID in auction
	auctionIdMap := make(map[uint64]bool)
	auctionCount := gs.GetAuctionCount()
	for _, elem := range gs.AuctionList {
		if _, ok := auctionIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for auction")
		}
		if elem.Id >= auctionCount {
			return fmt.Errorf("auction id should be lower or equal than the last id")
		}
		auctionIdMap[elem.Id] = true
	}
	// Check for duplicated ID in bid
	bidIdMap := make(map[uint64]bool)
	bidCount := gs.GetBidCount()
	for _, elem := range gs.BidList {
		if _, ok := bidIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for bid")
		}
		if elem.Id >= bidCount {
			return fmt.Errorf("bid id should be lower or equal than the last id")
		}
		bidIdMap[elem.Id] = true
	}
	// Check for duplicated index in users
	usersIndexMap := make(map[string]struct{})

	for _, elem := range gs.UsersList {
		index := string(UsersKey(elem.Index))
		if _, ok := usersIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for users")
		}
		usersIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
