package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelBid = "cancel_bid"

var _ sdk.Msg = &MsgCancelBid{}

func NewMsgCancelBid(creator string, auctionId uint64) *MsgCancelBid {
	return &MsgCancelBid{
		Creator:   creator,
		AuctionId: auctionId,
	}
}

func (msg *MsgCancelBid) Route() string {
	return RouterKey
}

func (msg *MsgCancelBid) Type() string {
	return TypeMsgCancelBid
}

func (msg *MsgCancelBid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelBid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelBid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.AuctionId < 0 {
		return sdkerrors.Wrapf(ErrInvalidAuction, "invalid AuctionId (%s)", err)
	}

	return nil
}
