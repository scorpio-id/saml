package service

import (
	"crypto/rsa"
	"crypto/x509"
	"net/http"
	"net/url"
	"time"

	"github.com/beevik/etree"
	dsig "github.com/russellhaering/goxmldsig"

	"github.com/scorpio-id/saml/internal/data"
	"github.com/scorpio-id/saml/internal/format"
)

// SignatureVerifier verifies a signature
//
// Can be implemented in order to override ServiceProvider's default
// way of verifying signatures.
type SignatureVerifier interface {
	VerifySignature(validationContext *dsig.ValidationContext, el *etree.Element) error
}

// ServiceProvider implements SAML Service provider.
//
// In SAML, service providers delegate responsibility for identifying
// clients to an identity provider. If you are writing an application
// that uses passwords (or whatever) stored somewhere else, then you
// are service provider.
//
// See the example directory for an example of a web application using
// the service provider interface.
type ServiceProvider struct {
	// Entity ID is optional - if not specified then MetadataURL will be used
	EntityID string

	// Key is the RSA private key we use to sign requests.
	Key *rsa.PrivateKey

	// Certificate is the RSA public part of Key.
	Certificate   *x509.Certificate
	Intermediates []*x509.Certificate

	// HTTPClient to use during SAML artifact resolution
	HTTPClient *http.Client

	// MetadataURL is the full URL to the metadata endpoint on this host,
	// i.e. https://example.com/saml/metadata
	MetadataURL url.URL

	// AcsURL is the full URL to the SAML Assertion Customer Service endpoint
	// on this host, i.e. https://example.com/saml/acs
	AcsURL url.URL

	// SloURL is the full URL to the SAML Single Logout endpoint on this host.
	// i.e. https://example.com/saml/slo
	SloURL url.URL

	// IDPMetadata is the metadata from the identity provider.
	IDPMetadata *data.EntityDescriptor

	// AuthnNameIDFormat is the format used in the NameIDPolicy for
	// authentication requests
	AuthnNameIDFormat format.NameID

	// MetadataValidDuration is a duration used to calculate validUntil
	// attribute in the metadata endpoint
	MetadataValidDuration time.Duration

	// ForceAuthn allows you to force re-authentication of users even if the user
	// has a SSO session at the IdP.
	ForceAuthn *bool

	// RequestedAuthnContext allow you to specify the requested authentication
	// context in authentication requests
	RequestedAuthnContext *data.RequestedAuthnContext

	// AllowIdpInitiated
	AllowIDPInitiated bool

	// DefaultRedirectURI where untracked requests (as of IDPInitiated) are redirected to
	DefaultRedirectURI string

	// SignatureVerifier, if non-nil, allows you to implement an alternative way
	// to verify signatures.
	SignatureVerifier SignatureVerifier

	// SignatureMethod, if non-empty, authentication requests will be signed
	SignatureMethod string

	// LogoutBindings specify the bindings available for SLO endpoint. If empty,
	// HTTP-POST binding is used.
	LogoutBindings []string
}
