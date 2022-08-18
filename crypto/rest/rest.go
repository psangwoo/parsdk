package rest

import (
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/psangwoo/parsdk/client"
	"github.com/psangwoo/parsdk/client/keys"
)

func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc(
		"/crypto/convert-address",
		convertAddressHandlerFn(cliCtx),
	).Methods("GET")
}

func convertAddressHandlerFn(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var accAddress sdk.AccAddress
		address := r.URL.Query().Get(keys.FlagAddress)
		if len(address) != 0 {
			ad, err := sdk.AccAddressFromBech32(address)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			accAddress = ad
		}

		rest.PostProcessResponse(w, cliCtx, sdk.ValAddress(accAddress))
	}
}
