package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"gitlab.com/gitlab-org/gitlab-workhorse/internal/api"
)

var servicesProxyPath = fmt.Sprintf("%s/-/jobs/1/proxy", testProject)

func TestServicesDeniedHTTPProxy(t *testing.T) {
	ts := testAuthServer(nil, 403, "Access denied")
	defer ts.Close()
	ws := startWorkhorseServer(ts.URL)
	defer ws.Close()

	resp, err := http.Get(ws.URL + "/" + servicesProxyPath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 403, resp.StatusCode)
}

func TestServicesHTTPProxy(t *testing.T) {
	message := []byte("ACK")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ensure we copy the Authorization header in the request
		assert.Equal(t, "foo", r.Header.Get("Authorization"))
		w.Header().Set("ReturnHeader", "foo")
		w.Write(message)
	}))
	defer srv.Close()

	out := &api.Response{
		Service: &api.ServiceProxySettings{
			Url:    "http://" + srv.Listener.Addr().String(),
			Header: http.Header{"Authorization": []string{"foo"}},
		},
	}

	out.Service.Validate()

	ts := testAuthServer(nil, 200, out)
	defer ts.Close()
	ws := startWorkhorseServer(ts.URL)
	defer ws.Close()

	resp, err := http.Get(ws.URL + "/" + servicesProxyPath)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	body := resp.Body
	data, err := ioutil.ReadAll(body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, message, data)
	// Ensure we copy the headers from the proxied server
	assert.Equal(t, "foo", resp.Header.Get("ReturnHeader"))
}

func TestServicesHTTPProxyPOST(t *testing.T) {
	payload := []byte("example body")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		data, err := ioutil.ReadAll(r.Body)
		require.NoError(t, err)
		assert.Equal(t, payload, data)
	}))
	defer srv.Close()

	out := &api.Response{
		Service: &api.ServiceProxySettings{
			Url:    "http://" + srv.Listener.Addr().String(),
			Header: http.Header{"Authorization": []string{"foo"}},
		},
	}

	out.Service.Validate()

	ts := testAuthServer(nil, 200, out)
	defer ts.Close()
	ws := startWorkhorseServer(ts.URL)
	defer ws.Close()

	resp, err := http.Post(ws.URL+"/"+servicesProxyPath, "text/plain", bytes.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}
