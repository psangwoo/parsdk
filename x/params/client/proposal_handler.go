package client

import (
	govclient "github.com/psangwoo/parsdk/x/gov/client"
	"github.com/psangwoo/parsdk/x/params/client/cli"
	"github.com/psangwoo/parsdk/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd, rest.ProposalRESTHandler)
