package serviceproxy

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"net/http"

	"gitlab.com/gitlab-org/gitlab-workhorse/internal/api"
	"gitlab.com/gitlab-org/gitlab-workhorse/internal/helper"
	"gitlab.com/gitlab-org/gitlab-workhorse/internal/transporthelper"
)

func Handler(myAPI *api.API) http.Handler {
	return myAPI.PreAuthorizeHandler(func(w http.ResponseWriter, r *http.Request, a *api.Response) {
		if err := a.Service.Validate(); err != nil {
			helper.Fail500(w, r, err)
			return
		}

		r.Header.Add("Authorization", a.Service.Header.Get("Authorization"))

		proxyRequest(w, r, a.Service)
	}, "authorize")
}

var transportWithTimeouts = transporthelper.TransportWithTimeouts()

var httpTransport = transporthelper.TracingRoundTripper(transportWithTimeouts)

func proxyRequest(w http.ResponseWriter, r *http.Request, s *api.ServiceProxySettings) {
	var err error

	r.URL, err = s.URL()
	if err != nil {
		helper.Fail500(w, r, err)
		return
	}

	if len(s.CAPem) > 0 {
		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM([]byte(s.CAPem))
		transportWithTimeouts.TLSClientConfig = &tls.Config{RootCAs: pool}
	}

	resp, err := httpTransport.RoundTrip(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
