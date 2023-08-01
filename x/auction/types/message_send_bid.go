package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendBid = "send_bid"

var _ sdk.Msg = &MsgSendBid{}

func NewMsgSendBid(creator string, auctionId uint64, price sdk.Coin) *MsgSendBid {
	return &MsgSendBid{
		Creator:   creator,
		AuctionId: auctionId,
		Price:     price,
	}
}

func (msg *MsgSendBid) Route() string {
	return RouterKey
}

func (msg *MsgSendBid) Type() string {
	return TypeMsgSendBid
}

func (msg *MsgSendBid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendBid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendBid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
