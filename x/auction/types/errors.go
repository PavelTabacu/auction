package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/auction module sentinel errors
var (
	ErrSample           = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrNotFound         = sdkerrors.Register(ModuleName, 1101, "Not found error")
	ErrInvalidAsset     = sdkerrors.Register(ModuleName, 1102, "Invalid asset")
	ErrInvalidUser      = sdkerrors.Register(ModuleName, 1103, "Invalid user")
	ErrInvalidPrice     = sdkerrors.Register(ModuleName, 1104, "Invalid Price")
	ErrInvalidAuction   = sdkerrors.Register(ModuleName, 1105, "Invalid Auction")
	ErrInvalidOperation = sdkerrors.Register(ModuleName, 1106, "Invalid operation")
)
