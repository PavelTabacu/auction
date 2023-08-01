package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UsersKeyPrefix is the prefix to retrieve all Users
	UsersKeyPrefix = "Users/value/"
)

// UsersKey returns the store key to retrieve a Users from the index fields
func UsersKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
