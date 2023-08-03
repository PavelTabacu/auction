package types

import "time"

const (
	// ModuleName defines the module name
	ModuleName = "auction"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_auction"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	AssetKey      = "Asset/value/"
	AssetCountKey = "Asset/count/"
)

const (
	AuctionKey      = "Auction/value/"
	AuctionCountKey = "Auction/count/"
)

const (
	BidKey      = "Bid/value/"
	BidCountKey = "Bid/count/"
)

const (
	MaxTurnDuration = time.Duration(24 * 3_600 * 1000_000_000) // 1 day
	DeadlineLayout  = "2006-01-02 15:04:05.999999999 +0000 UTC"
)
