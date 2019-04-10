package api

import (
	"fmt"
	"net/http"
	"net/url"

	"gitlab.com/gitlab-org/gitlab-workhorse/internal/helper"
)

type ServiceProxySettings struct {
	// The URL to connect to.
	Url string

	// Any headers (e.g., Authorization) to send with the request
	Header http.Header

	// The CA roots to validate the remote endpoint with, for wss:// URLs. The
	// system-provided CA pool will be used if this is blank. PEM-encoded data.
	CAPem string
}

func (t *ServiceProxySettings) URL() (*url.URL, error) {
	return url.Parse(t.Url)
}

func (t *ServiceProxySettings) Clone() *ServiceProxySettings {
	// Doesn't clone the strings, but that's OK as strings are immutable in go
	cloned := *t
	cloned.Header = helper.HeaderClone(t.Header)
	return &cloned
}

func (t *ServiceProxySettings) Validate() error {
	if t == nil {
		return fmt.Errorf("service details not specified")
	}

	parsedURL, err := t.URL()
	if err != nil {
		return fmt.Errorf("invalid URL")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("invalid scheme: %q", parsedURL.Scheme)
	}

	return nil
}

func (t *ServiceProxySettings) IsEqual(other *ServiceProxySettings) bool {
	if t == nil && other == nil {
		return true
	}

	if t == nil || other == nil {
		return false
	}

	if len(t.Header) != len(other.Header) {
		return false
	}

	for header, values := range t.Header {
		if len(values) != len(other.Header[header]) {
			return false
		}
		for i, value := range values {
			if other.Header[header][i] != value {
				return false
			}
		}
	}

	return t.Url == other.Url &&
		t.CAPem == other.CAPem
}
