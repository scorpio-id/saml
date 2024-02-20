package transport

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scorpio-id/saml/internal/config"
	"github.com/scorpio-id/saml/internal/identity"
)

// NewRouter creates a new mux router with applied SAML IDP configurations
// TODO - fix the awful naming on idp
func NewRouter(cfg config.Config) *mux.Router {

	// create gorilla mux router
	router := mux.NewRouter()

	// create an IDP
	idp, err := identity.NewIDP(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/certificate", idp.CertificateHandler).Methods(http.MethodGet, http.MethodOptions)

	return router
}
