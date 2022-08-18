package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/psangwoo/parsdk/client"
	// "github.com/psangwoo/parsdk/client/flags"
	"github.com/psangwoo/parsdk/x/oracle/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdSetAdminAddr())

	cmd.AddCommand(CmdDeleteExchangeRates())

	cmd.AddCommand(CmdDeleteExchangeRate())

	cmd.AddCommand(CmdSetExchangeRate())

	return cmd
}
