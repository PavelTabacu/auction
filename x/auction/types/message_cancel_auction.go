package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelAuction = "cancel_auction"

var _ sdk.Msg = &MsgCancelAuction{}

func NewMsgCancelAuction(creator string, auctionId uint64) *MsgCancelAuction {
	return &MsgCancelAuction{
		Creator:   creator,
		AuctionId: auctionId,
	}
}

func (msg *MsgCancelAuction) Route() string {
	return RouterKey
}

func (msg *MsgCancelAuction) Type() string {
	return TypeMsgCancelAuction
}

func (msg *MsgCancelAuction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelAuction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.AuctionId < 0 {
		return sdkerrors.Wrapf(ErrInvalidAuction, "invalid auctionId (%s)", err)
	}

	return nil
}
