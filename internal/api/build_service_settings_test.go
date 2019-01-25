package api

import (
	"net/http"
	"testing"
)

func buildService(url string) *BuildServiceSettings {
	return &BuildServiceSettings{
		Url: url,
	}
}

func buildServiceCa(buildService *BuildServiceSettings) *BuildServiceSettings {
	buildService = buildService.Clone()
	buildService.CAPem = "Valid CA data"

	return buildService
}

func buildServiceHeader(buildService *BuildServiceSettings, values ...string) *BuildServiceSettings {
	if len(values) == 0 {
		values = []string{"Dummy Value"}
	}

	buildService = buildService.Clone()
	buildService.Header = http.Header{
		"Header": values,
	}

	return buildService
}

func TestBuildServiceClone(t *testing.T) {
	a := buildServiceCa(buildServiceHeader(buildService("http:")))
	b := a.Clone()

	if a == b {
		t.Fatalf("Address of cloned build service didn't change")
	}

	if &a.Header == &b.Header {
		t.Fatalf("Address of cloned header didn't change")
	}
}

func TestBuildServiceValidate(t *testing.T) {
	for i, tc := range []struct {
		buildService *BuildServiceSettings
		valid        bool
		msg          string
	}{
		{nil, false, "nil build service"},
		{buildService(""), false, "empty URL"},
		{buildService("ws:"), false, "websocket URL"},
		{buildService("wss:"), false, "secure websocket URL"},
		{buildService("http:"), true, "HTTP URL"},
		{buildService("https:"), true, "HTTPS URL"},
		{buildServiceCa(buildService("http:")), true, "any CA pem"},
		{buildServiceHeader(buildService("http:")), true, "any headers"},
		{buildServiceCa(buildServiceHeader(buildService("http:"))), true, "PEM and headers"},
	} {
		if err := tc.buildService.Validate(); (err != nil) == tc.valid {
			t.Fatalf("test case %d: "+tc.msg+": valid=%v: %s: %+v", i, tc.valid, err, tc.buildService)
		}
	}
}

func TestBuildServiceIsEqual(t *testing.T) {
	serv := buildService("http:")

	servHeader2 := buildServiceHeader(serv, "extra")
	servHeader3 := buildServiceHeader(serv)
	servHeader3.Header.Add("Extra", "extra")

	servCa2 := buildServiceCa(serv)
	servCa2.CAPem = "other value"

	for i, tc := range []struct {
		serviceA *BuildServiceSettings
		serviceB *BuildServiceSettings
		expected bool
	}{
		{nil, nil, true},
		{serv, nil, false},
		{nil, serv, false},
		{serv, serv, true},
		{serv.Clone(), serv.Clone(), true},
		{serv, buildService("foo:"), false},
		{serv, buildService(serv.Url), true},
		{buildServiceHeader(serv), buildServiceHeader(serv), true},
		{servHeader2, servHeader2, true},
		{servHeader3, servHeader3, true},
		{buildServiceHeader(serv), servHeader2, false},
		{buildServiceHeader(serv), servHeader3, false},
		{buildServiceHeader(serv), serv, false},
		{serv, buildServiceHeader(serv), false},
		{buildServiceCa(serv), buildServiceCa(serv), true},
		{buildServiceCa(serv), serv, false},
		{serv, buildServiceCa(serv), false},
		{buildServiceCa(buildServiceHeader(serv)), buildServiceCa(buildServiceHeader(serv)), true},
		{servCa2, buildServiceCa(serv), false},
	} {
		if actual := tc.serviceA.IsEqual(tc.serviceB); tc.expected != actual {
			t.Fatalf(
				"test case %d: Comparison:\n-%+v\n+%+v\nexpected=%v: actual=%v",
				i, tc.serviceA, tc.serviceB, tc.expected, actual,
			)
		}
	}
}
