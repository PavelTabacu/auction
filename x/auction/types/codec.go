package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateAsset{}, "auction/CreateAsset", nil)
	cdc.RegisterConcrete(&MsgUpdateAsset{}, "auction/UpdateAsset", nil)
	cdc.RegisterConcrete(&MsgDeleteAsset{}, "auction/DeleteAsset", nil)
	cdc.RegisterConcrete(&MsgInitiateAuction{}, "auction/InitiateAuction", nil)
	cdc.RegisterConcrete(&MsgCancelAuction{}, "auction/CancelAuction", nil)
	cdc.RegisterConcrete(&MsgExecuteAuction{}, "auction/ExecuteAuction", nil)
	cdc.RegisterConcrete(&MsgSendBid{}, "auction/SendBid", nil)
	cdc.RegisterConcrete(&MsgUpdateBid{}, "auction/UpdateBid", nil)
	cdc.RegisterConcrete(&MsgCancelBid{}, "auction/CancelBid", nil)
	cdc.RegisterConcrete(&MsgUpdateAuction{}, "auction/UpdateAuction", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAsset{},
		&MsgUpdateAsset{},
		&MsgDeleteAsset{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitiateAuction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelAuction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExecuteAuction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendBid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateBid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelBid{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateAuction{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
