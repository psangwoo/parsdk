package cli

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/psangwoo/parsdk/client"
	"github.com/psangwoo/parsdk/client/flags"
	"github.com/psangwoo/parsdk/client/tx"
	"github.com/psangwoo/parsdk/x/oracle/types"
)

func CmdDeleteExchangeRates() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-exchange-rates [pairs]",
		Short: "Broadcast message DeleteExchangeRates",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pairs := strings.Split(args[0], ",")

			msg := types.NewMsgDeleteExchangeRates(clientCtx.GetFromAddress().String(), pairs)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
