package main

import (
	"log"
	"net/http"

	"github.com/scorpio-id/saml/internal/config"
	"github.com/scorpio-id/saml/internal/transport"
)

func main() {
	// parse local config
	cfg := config.NewConfig("internal/config/local.yml")

	// create a new mux router
	router := transport.NewRouter(cfg)
	
	// start the server
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, router))
}