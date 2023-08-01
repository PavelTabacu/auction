package cli

import (
	"strconv"

	"github.com/PavelTabacu/auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdInitiateAuction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "initiate-auction [token-id] [start-price] [deadline]",
		Short: "Broadcast message initiate-auction",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argStartPrice, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			argDeadline := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInitiateAuction(
				clientCtx.GetFromAddress().String(),
				argTokenId,
				argStartPrice,
				argDeadline,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
