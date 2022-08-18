package client

import (
	"github.com/psangwoo/parsdk/x/distribution/client/cli"
	"github.com/psangwoo/parsdk/x/distribution/client/rest"
	govclient "github.com/psangwoo/parsdk/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
