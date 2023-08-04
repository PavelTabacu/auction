package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateAuction = "update_auction"

var _ sdk.Msg = &MsgUpdateAuction{}

func NewMsgUpdateAuction(creator string, auctionId uint64, startPrice sdk.Coin, deadline string) *MsgUpdateAuction {
	return &MsgUpdateAuction{
		Creator:    creator,
		AuctionId:  auctionId,
		StartPrice: startPrice,
		Deadline:   deadline,
	}
}

func (msg *MsgUpdateAuction) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAuction) Type() string {
	return TypeMsgUpdateAuction
}

func (msg *MsgUpdateAuction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAuction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
