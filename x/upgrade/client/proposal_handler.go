package client

import (
	govclient "github.com/psangwoo/parsdk/x/gov/client"
	"github.com/psangwoo/parsdk/x/upgrade/client/cli"
	"github.com/psangwoo/parsdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal, rest.ProposalCancelRESTHandler)
