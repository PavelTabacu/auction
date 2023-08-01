package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInitiateAuction = "initiate_auction"

var _ sdk.Msg = &MsgInitiateAuction{}

func NewMsgInitiateAuction(creator string, tokenId uint64, startPrice sdk.Coin, deadline string) *MsgInitiateAuction {
	return &MsgInitiateAuction{
		Creator:    creator,
		TokenId:    tokenId,
		StartPrice: startPrice,
		Deadline:   deadline,
	}
}

func (msg *MsgInitiateAuction) Route() string {
	return RouterKey
}

func (msg *MsgInitiateAuction) Type() string {
	return TypeMsgInitiateAuction
}

func (msg *MsgInitiateAuction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInitiateAuction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInitiateAuction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	return nil
}
