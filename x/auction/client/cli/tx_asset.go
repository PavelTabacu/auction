package cli

import (
	"strconv"

	"github.com/PavelTabacu/auction/x/auction/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-asset [token-id] [name] [symbol] [owner] [description] [url] [is-auction]",
		Short: "Create a new asset",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTokenId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argName := args[1]
			argSymbol := args[2]
			argOwner := args[3]
			argDescription := args[4]
			argUrl := args[5]
			argIsAuction, err := cast.ToBoolE(args[6])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAsset(clientCtx.GetFromAddress().String(), argTokenId, argName, argSymbol, argOwner, argDescription, argUrl, argIsAuction)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-asset [id] [token-id] [name] [symbol] [owner] [description] [url] [is-auction]",
		Short: "Update a asset",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argTokenId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argName := args[2]

			argSymbol := args[3]

			argOwner := args[4]

			argDescription := args[5]

			argUrl := args[6]

			argIsAuction, err := cast.ToBoolE(args[7])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAsset(clientCtx.GetFromAddress().String(), id, argTokenId, argName, argSymbol, argOwner, argDescription, argUrl, argIsAuction)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-asset [id]",
		Short: "Delete a asset by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteAsset(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
