package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgExecuteAuction = "execute_auction"

var _ sdk.Msg = &MsgExecuteAuction{}

func NewMsgExecuteAuction(creator string, auctionId uint64) *MsgExecuteAuction {
	return &MsgExecuteAuction{
		Creator:   creator,
		AuctionId: auctionId,
	}
}

func (msg *MsgExecuteAuction) Route() string {
	return RouterKey
}

func (msg *MsgExecuteAuction) Type() string {
	return TypeMsgExecuteAuction
}

func (msg *MsgExecuteAuction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgExecuteAuction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgExecuteAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
