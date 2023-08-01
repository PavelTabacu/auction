package keeper

import (
	"github.com/PavelTabacu/auction/x/auction/types"
)

var _ types.QueryServer = Keeper{}
