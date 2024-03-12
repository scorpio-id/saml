package transport

import (
	"flag"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/scorpio-id/saml/internal/config"
	"github.com/scorpio-id/saml/internal/data"
	"github.com/scorpio-id/saml/internal/identity"
	"github.com/scorpio-id/saml/internal/logger"
	"github.com/scorpio-id/saml/internal/provider"
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

	logr := logger.DefaultLogger
	baseURLstr := flag.String("idp", "", "The URL to the IDP")
	flag.Parse()

	baseURL, err := url.Parse(*baseURLstr)
	if err != nil {
		logr.Fatalf("cannot parse base URL: %v", err)
	}

	// TODO - Tuesday, begin swap to idp in identity package
	// FIXME - relationship between server below and IDP
	s, err := provider.New(provider.Options{
		URL:         *baseURL,
		Key:         idp.Private,
		Logger:      logr,
		Certificate: idp.Leaf,
		Store:       &data.MemoryStore{},
	})
	if err != nil {
		logr.Fatalf("%s", err)
	}

	// TODO - leave out for now, example purposes
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

	router.HandleFunc("/metadata", func(w http.ResponseWriter, r *http.Request) {
		s.IdpConfigMu.RLock()
		defer s.IdpConfigMu.RUnlock()
		s.IDP.ServeMetadata(w, r)
	})
	router.HandleFunc("/sso", func(w http.ResponseWriter, r *http.Request) {
		s.IdpConfigMu.RLock()
		defer s.IdpConfigMu.RUnlock()
		s.IDP.ServeSSO(w, r)
	})

	// mux.Handle("/login", s.HandleLogin)
	// mux.Handle("/login/:shortcut", s.HandleIDPInitiated)
	// mux.Handle("/login/:shortcut/*", s.HandleIDPInitiated)

	router.HandleFunc("/login", s.HandleLogin)
	router.HandleFunc("/login/{shortcut}", s.HandleIDPInitiated)

	// TODO - check if you can use *
	router.HandleFunc("/login/{shortcut}/{*}", s.HandleIDPInitiated)

	// mux.Get("/services/", s.HandleListServices)
	// mux.Get("/services/:id", s.HandleGetService)
	// mux.Put("/services/:id", s.HandlePutService)
	// mux.Post("/services/:id", s.HandlePutService)
	// mux.Delete("/services/:id", s.HandleDeleteService)

	// ignored trailing slash in "/services/" from above
	router.HandleFunc("/services", s.HandleListServices).Methods("GET")
	router.HandleFunc("/services/{id}", s.HandleGetService).Methods("GET")
	router.HandleFunc("/services/{id}", s.HandlePutService).Methods("POST", "PUT")
	router.HandleFunc("/services/{id}", s.HandleDeleteService).Methods("DELETE")

	// mux.Get("/users/", s.HandleListUsers)
	// mux.Get("/users/:id", s.HandleGetUser)
	// mux.Put("/users/:id", s.HandlePutUser)
	// mux.Delete("/users/:id", s.HandleDeleteUser)

	router.HandleFunc("/users", s.HandleListUsers).Methods("GET")
	router.HandleFunc("/users/{id}", s.HandleGetUser).Methods("GET")
	router.HandleFunc("/users/{id}", s.HandlePutUser).Methods("PUT")
	router.HandleFunc("/users/{id}", s.HandleDeleteUser).Methods("DELETE")

	// sessionPath := regexp.MustCompile("/sessions/(?P<id>.*)")
	// mux.Get("/sessions/", s.HandleListSessions)
	// mux.Get(sessionPath, s.HandleGetSession)
	// mux.Delete(sessionPath, s.HandleDeleteSession)

	router.HandleFunc("/sessions", s.HandleListSessions).Methods("GET")
	router.HandleFunc("/sessions/(?P<id>.*)", s.HandleGetSession).Methods("GET")
	router.HandleFunc("/sessions/(?P<id>.*)", s.HandleDeleteSession).Methods("DELETE")

	// mux.Get("/shortcuts/", s.HandleListShortcuts)
	// mux.Get("/shortcuts/:id", s.HandleGetShortcut)
	// mux.Put("/shortcuts/:id", s.HandlePutShortcut)
	// mux.Delete("/shortcuts/:id", s.HandleDeleteShortcut)

	router.HandleFunc("/shortcuts", s.HandleListShortcuts).Methods("GET")
	router.HandleFunc("/shortcuts/{id}", s.HandleGetShortcut).Methods("GET")
	router.HandleFunc("/shortcuts/{id}", s.HandlePutShortcut).Methods("PUT")
	router.HandleFunc("/shortcuts/{id}", s.HandleDeleteShortcut).Methods("DELETE")

	return router
}
