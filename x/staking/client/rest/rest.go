package rest

import (
	"github.com/gorilla/mux"

	"github.com/psangwoo/parsdk/client"
	"github.com/psangwoo/parsdk/client/rest"
)

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}
