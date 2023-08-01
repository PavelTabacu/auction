package simulation

import (
	"math/rand"

	"github.com/PavelTabacu/auction/x/auction/keeper"
	"github.com/PavelTabacu/auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCancelAuction(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCancelAuction{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CancelAuction simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CancelAuction simulation not implemented"), nil, nil
	}
}
