package transport

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scorpio-id/saml/internal/config"
	"github.com/scorpio-id/saml/internal/idp"
)

// NewRouter creates a new mux router with applied SAML IDP configurations
// TODO - fix the awful naming on idp
func NewRouter(cfg config.Config, idp *idp.IDP) (*mux.Router) {

	// create gorilla mux router
	router := mux.NewRouter()

	router.HandleFunc("/cert", idp.GetCert).Methods(http.MethodGet, http.MethodOptions)

	return router
}