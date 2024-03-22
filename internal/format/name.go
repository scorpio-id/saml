package format

import "github.com/beevik/etree"

// NameID is the format of the id
type NameID string

// Element returns an XML element representation of n.
func (n NameID) Element() *etree.Element {
	el := etree.NewElement("")
	el.SetText(string(n))
	return el
}

// Name ID formats
const (
	UnspecifiedNameIDFormat  NameID = "urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"
	TransientNameIDFormat    NameID = "urn:oasis:names:tc:SAML:2.0:nameid-format:transient"
	EmailAddressNameIDFormat NameID = "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress"
	PersistentNameIDFormat   NameID = "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent"
)
