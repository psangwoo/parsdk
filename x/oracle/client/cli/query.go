package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/psangwoo/parsdk/client"
	// "github.com/psangwoo/parsdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/psangwoo/parsdk/x/oracle/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group oracle queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	//	cmd.AddCommand(CmdListParams())
	//	cmd.AddCommand(CmdShowParams())

	cmd.AddCommand(CmdListAdminAddr())
	//	cmd.AddCommand(CmdShowAdminAddr())

	cmd.AddCommand(CmdListExchangeRate())
	cmd.AddCommand(CmdShowExchangeRate())

	return cmd
}
