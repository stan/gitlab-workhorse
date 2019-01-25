package transporthelper

import (
	"net"
	"net/http"
	"time"

	"gitlab.com/gitlab-org/labkit/correlation"
	"gitlab.com/gitlab-org/labkit/tracing"
)

// TransportWithTimeouts returns more restrictive http.Transport
// than for http.DefaultTransport,
// They define shorter TLS Handshake, and more agressive connection closing
// to prevent the connection hanging and reduce FD usage
func TransportWithTimeouts() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		MaxIdleConns:          2,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
		ResponseHeaderTimeout: 30 * time.Second,
	}
}

func TracingRoundTripper(transport *http.Transport) http.RoundTripper {
	return tracing.NewRoundTripper(correlation.NewInstrumentedRoundTripper(transport))
}
