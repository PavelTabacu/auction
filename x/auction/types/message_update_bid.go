package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateBid = "update_bid"

var _ sdk.Msg = &MsgUpdateBid{}

func NewMsgUpdateBid(creator string, bidId uint64, price sdk.Coin) *MsgUpdateBid {
	return &MsgUpdateBid{
		Creator: creator,
		BidId:   bidId,
		Price:   price,
	}
}

func (msg *MsgUpdateBid) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBid) Type() string {
	return TypeMsgUpdateBid
}

func (msg *MsgUpdateBid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBid) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBid) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.BidId < 0 {
		return sdkerrors.Wrapf(ErrInvalidAuction, "invalid bidId (%s)", err)
	}

	return nil
}
