package main

import (
	"log"
	"net/http"

	"github.com/scorpio-id/saml/internal/config"
	"github.com/scorpio-id/saml/internal/idp"
	"github.com/scorpio-id/saml/internal/transport"
)

func main() {
		// parse local config
		cfg := config.NewConfig("internal/config/local.yml")

		// create an IDP
		idp, err := idp.NewIDP(&cfg)
		if err != nil {
			log.Fatal(err)
		}

		// create a new mux router
		router := transport.NewRouter(cfg, idp)
	
		// start the server
		log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))
}