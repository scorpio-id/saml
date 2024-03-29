package pkg

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"net/url"

	xrv "github.com/mattermost/xml-roundtrip-validator"

	"github.com/scorpio-id/saml/internal/logger"
	"github.com/scorpio-id/saml/internal/provider"
)

// ParseMetadata parses arbitrary SAML IDP metadata.
//
// Note: this is needed because IDP metadata is sometimes wrapped in
// an <EntitiesDescriptor>, and sometimes the top level element is an
// <EntityDescriptor>.
func ParseMetadata(data []byte) (*provider.EntityDescriptor, error) {
	entity := &provider.EntityDescriptor{}

	if err := xrv.Validate(bytes.NewBuffer(data)); err != nil {
		return nil, err
	}

	err := xml.Unmarshal(data, entity)

	// this comparison is ugly, but it is how the error is generated in encoding/xml
	if err != nil && err.Error() == "expected element type <EntityDescriptor> but have <EntitiesDescriptor>" {
		entities := &provider.EntitiesDescriptor{}
		if err := xml.Unmarshal(data, entities); err != nil {
			return nil, err
		}

		for i, e := range entities.EntityDescriptors {
			if len(e.IDPSSODescriptors) > 0 {
				return &entities.EntityDescriptors[i], nil
			}
		}
		return nil, errors.New("no entity found with IDPSSODescriptor")
	}
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// FetchMetadata returns metadata from an IDP metadata URL.
func FetchMetadata(ctx context.Context, httpClient *http.Client, metadataURL url.URL) (*provider.EntityDescriptor, error) {
	req, err := http.NewRequest("GET", metadataURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.DefaultLogger.Printf("Error while closing response body during fetch metadata: %v", err)
		}
	}()
	if resp.StatusCode >= 400 {
		return nil, errors.New("HTTP 400")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return ParseMetadata(data)
}