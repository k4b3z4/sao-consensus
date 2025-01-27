package cli

import (
	"strconv"

	"strings"

	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)
var listSeparator = ","

func CmdMigrate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate [data]",
		Short: "Broadcast message migrate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argData := strings.Split(args[0], listSeparator)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMigrate(
				clientCtx.GetFromAddress().String(),
				argData,
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
