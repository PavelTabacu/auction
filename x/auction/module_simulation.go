package auction

import (
	"math/rand"

	"github.com/PavelTabacu/auction/testutil/sample"
	auctionsimulation "github.com/PavelTabacu/auction/x/auction/simulation"
	"github.com/PavelTabacu/auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = auctionsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateAsset = "op_weight_msg_asset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAsset int = 100

	opWeightMsgUpdateAsset = "op_weight_msg_asset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAsset int = 100

	opWeightMsgDeleteAsset = "op_weight_msg_asset"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAsset int = 100

	opWeightMsgInitiateAuction = "op_weight_msg_initiate_auction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInitiateAuction int = 100

	opWeightMsgCancelAuction = "op_weight_msg_cancel_auction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelAuction int = 100

	opWeightMsgExecuteAuction = "op_weight_msg_execute_auction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgExecuteAuction int = 100

	opWeightMsgSendBid = "op_weight_msg_send_bid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendBid int = 100

	opWeightMsgUpdateBid = "op_weight_msg_update_bid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBid int = 100

	opWeightMsgCancelBid = "op_weight_msg_cancel_bid"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelBid int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	auctionGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AssetList: []types.Asset{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		AssetCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&auctionGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAsset int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAsset, &weightMsgCreateAsset, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAsset = defaultWeightMsgCreateAsset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAsset,
		auctionsimulation.SimulateMsgCreateAsset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAsset int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAsset, &weightMsgUpdateAsset, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAsset = defaultWeightMsgUpdateAsset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAsset,
		auctionsimulation.SimulateMsgUpdateAsset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAsset int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAsset, &weightMsgDeleteAsset, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAsset = defaultWeightMsgDeleteAsset
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAsset,
		auctionsimulation.SimulateMsgDeleteAsset(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInitiateAuction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInitiateAuction, &weightMsgInitiateAuction, nil,
		func(_ *rand.Rand) {
			weightMsgInitiateAuction = defaultWeightMsgInitiateAuction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInitiateAuction,
		auctionsimulation.SimulateMsgInitiateAuction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelAuction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelAuction, &weightMsgCancelAuction, nil,
		func(_ *rand.Rand) {
			weightMsgCancelAuction = defaultWeightMsgCancelAuction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelAuction,
		auctionsimulation.SimulateMsgCancelAuction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgExecuteAuction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgExecuteAuction, &weightMsgExecuteAuction, nil,
		func(_ *rand.Rand) {
			weightMsgExecuteAuction = defaultWeightMsgExecuteAuction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgExecuteAuction,
		auctionsimulation.SimulateMsgExecuteAuction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendBid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSendBid, &weightMsgSendBid, nil,
		func(_ *rand.Rand) {
			weightMsgSendBid = defaultWeightMsgSendBid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendBid,
		auctionsimulation.SimulateMsgSendBid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateBid, &weightMsgUpdateBid, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBid = defaultWeightMsgUpdateBid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBid,
		auctionsimulation.SimulateMsgUpdateBid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelBid int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelBid, &weightMsgCancelBid, nil,
		func(_ *rand.Rand) {
			weightMsgCancelBid = defaultWeightMsgCancelBid
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelBid,
		auctionsimulation.SimulateMsgCancelBid(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateAsset,
			defaultWeightMsgCreateAsset,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				auctionsimulation.SimulateMsgCreateAsset(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateAsset,
			defaultWeightMsgUpdateAsset,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				auctionsimulation.SimulateMsgUpdateAsset(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteAsset,
			defaultWeightMsgDeleteAsset,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				auctionsimulation.SimulateMsgDeleteAsset(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgInitiateAuction,
			defaultWeightMsgInitiateAuction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				auctionsimulation.SimulateMsgInitiateAuction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCancelAuction,
			defaultWeightMsgCancelAuction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				auctionsimulation.SimulateMsgCancelAuction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgExecuteAuction,
			defaultWeightMsgExecuteAuction,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				auctionsimulation.SimulateMsgExecuteAuction(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSendBid,
			defaultWeightMsgSendBid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				auctionsimulation.SimulateMsgSendBid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateBid,
			defaultWeightMsgUpdateBid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				auctionsimulation.SimulateMsgUpdateBid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCancelBid,
			defaultWeightMsgCancelBid,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				auctionsimulation.SimulateMsgCancelBid(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
