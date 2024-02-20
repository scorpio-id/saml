package identity

import (
	"crypto/rand"
	"crypto/rsa"
	"net/http"

	pki "github.com/scorpio-id/pki/pkg/client"
	"github.com/scorpio-id/saml/internal/config"
)

type IDP struct {
	// TODO - X509 needs to be unmarshaled into an x509 cert
	X509    string
	private *rsa.PrivateKey
	client  *pki.X509Client
}

func NewIDP(config *config.Config) (*IDP, error) {
	// create a new RSA key pair
	private, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	// create a HTTP client
	transport := &http.Client{}

	// placeholders for x509 cert (either signed by CA of self-signed)
	var cert string
	var x509client *pki.X509Client

	// get a signed x509 given configured SANs if PKI is enabled
	if config.PKI.Enabled {
		// create a new x509 client 
		x509client := pki.NewX509Client(config.PKI.URL, private, transport)

		// perform client credentials grant to get OAuth JWT
		jwt, err := x509client.AuthenticateCredentials(config.OAuth.Issuer, config.OAuth.ClientID)
		if err != nil {
			return nil, err
		}

		cert, err = x509client.GetCertificate(config.PKI.SANs, jwt)
		if err != nil {
			return nil, err
		}
	} else {
		// FIXME use a self-signed certificate if PKI is not enabled
		// https://go.dev/src/crypto/tls/generate_cert.go
	}
	return &IDP {
		X509: cert,
		private: private,
		client: x509client,
	}, nil
}

func (idp *IDP) CertificateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(idp.X509))
}