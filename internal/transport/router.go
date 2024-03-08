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

	// logr := logger.DefaultLogger
	// baseURLstr := flag.String("idp", "", "The URL to the IDP")
	// flag.Parse()

	// baseURL, err := url.Parse(*baseURLstr)
	// if err != nil {
	// 	logr.Fatalf("cannot parse base URL: %v", err)
	// }

	// // TODO - Tuesday, begin swap to idp in identity package
	// idpServer, err := provider.New(provider.Options{
	// 	URL:         *baseURL,
	// 	Key:         key,
	// 	Logger:      logr,
	// 	Certificate: cert,
	// 	Store:       &samlidp.MemoryStore{},
	// })
	// if err != nil {
	// 	logr.Fatalf("%s", err)
	// }

	// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("hunter2"), bcrypt.DefaultCost)
	// err = idpServer.Store.Put("/users/alice", samlidp.User{Name: "alice",
	// 	HashedPassword: hashedPassword,
	// 	Groups:         []string{"Administrators", "Users"},
	// 	Email:          "alice@example.com",
	// 	CommonName:     "Alice Smith",
	// 	Surname:        "Smith",
	// 	GivenName:      "Alice",
	// })
	// if err != nil {
	// 	logr.Fatalf("%s", err)
	// }

	// err = idpServer.Store.Put("/users/bob", samlidp.User{
	// 	Name:           "bob",
	// 	HashedPassword: hashedPassword,
	// 	Groups:         []string{"Users"},
	// 	Email:          "bob@example.com",
	// 	CommonName:     "Bob Smith",
	// 	Surname:        "Smith",
	// 	GivenName:      "Bob",
	// })
	// if err != nil {
	// 	logr.Fatalf("%s", err)
	// }

	router.HandleFunc("/certificate", idp.CertificateHandler).Methods(http.MethodGet, http.MethodOptions)

	// mux.Get("/metadata", func(w http.ResponseWriter, r *http.Request) {
	// 	s.idpConfigMu.RLock()
	// 	defer s.idpConfigMu.RUnlock()
	// 	s.IDP.ServeMetadata(w, r)
	// })
	// mux.Handle("/sso", func(w http.ResponseWriter, r *http.Request) {
	// 	s.idpConfigMu.RLock()
	// 	defer s.idpConfigMu.RUnlock()
	// 	s.IDP.ServeSSO(w, r)
	// })

	// mux.Handle("/login", s.HandleLogin)
	// mux.Handle("/login/:shortcut", s.HandleIDPInitiated)
	// mux.Handle("/login/:shortcut/*", s.HandleIDPInitiated)

	// mux.Get("/services/", s.HandleListServices)
	// mux.Get("/services/:id", s.HandleGetService)
	// mux.Put("/services/:id", s.HandlePutService)
	// mux.Post("/services/:id", s.HandlePutService)
	// mux.Delete("/services/:id", s.HandleDeleteService)

	// mux.Get("/users/", s.HandleListUsers)
	// mux.Get("/users/:id", s.HandleGetUser)
	// mux.Put("/users/:id", s.HandlePutUser)
	// mux.Delete("/users/:id", s.HandleDeleteUser)

	// sessionPath := regexp.MustCompile("/sessions/(?P<id>.*)")
	// mux.Get("/sessions/", s.HandleListSessions)
	// mux.Get(sessionPath, s.HandleGetSession)
	// mux.Delete(sessionPath, s.HandleDeleteSession)

	// mux.Get("/shortcuts/", s.HandleListShortcuts)
	// mux.Get("/shortcuts/:id", s.HandleGetShortcut)
	// mux.Put("/shortcuts/:id", s.HandlePutShortcut)
	// mux.Delete("/shortcuts/:id", s.HandleDeleteShortcut)

	return router
}
