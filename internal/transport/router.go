package transport

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scorpio-id/saml/internal/config"
	"github.com/scorpio-id/saml/internal/idp"
)

// NewRouter creates a new mux router with applied SAML IDP configurations
func NewRouter(cfg config.Config) (*mux.Router) {

	// create gorilla mux router
	router := mux.NewRouter()

	router.HandleFunc("/idp", idp.HelloFromIDP).Methods(http.MethodGet, http.MethodOptions)

	return router
}