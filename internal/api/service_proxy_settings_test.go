package api

import (
	"net/http"
	"testing"
)

func service(url string) *ServiceProxySettings {
	return &ServiceProxySettings{
		Url: url,
	}
}

func serviceCa(service *ServiceProxySettings) *ServiceProxySettings {
	service = service.Clone()
	service.CAPem = "Valid CA data"

	return service
}

func serviceHeader(service *ServiceProxySettings, values ...string) *ServiceProxySettings {
	if len(values) == 0 {
		values = []string{"Dummy Value"}
	}

	service = service.Clone()
	service.Header = http.Header{
		"Header": values,
	}

	return service
}

func TestServiceClone(t *testing.T) {
	a := serviceCa(serviceHeader(service("http:")))
	b := a.Clone()

	if a == b {
		t.Fatalf("Address of cloned build service didn't change")
	}

	if &a.Header == &b.Header {
		t.Fatalf("Address of cloned header didn't change")
	}
}

func TestServiceValidate(t *testing.T) {
	for i, tc := range []struct {
		service *ServiceProxySettings
		valid   bool
		msg     string
	}{
		{nil, false, "nil build service"},
		{service(""), false, "empty URL"},
		{service("ws:"), false, "websocket URL"},
		{service("wss:"), false, "secure websocket URL"},
		{service("http:"), true, "HTTP URL"},
		{service("https:"), true, "HTTPS URL"},
		{serviceCa(service("http:")), true, "any CA pem"},
		{serviceHeader(service("http:")), true, "any headers"},
		{serviceCa(serviceHeader(service("http:"))), true, "PEM and headers"},
	} {
		if err := tc.service.Validate(); (err != nil) == tc.valid {
			t.Fatalf("test case %d: "+tc.msg+": valid=%v: %s: %+v", i, tc.valid, err, tc.service)
		}
	}
}

func TestServiceIsEqual(t *testing.T) {
	serv := service("http:")

	servHeader2 := serviceHeader(serv, "extra")
	servHeader3 := serviceHeader(serv)
	servHeader3.Header.Add("Extra", "extra")

	servCa2 := serviceCa(serv)
	servCa2.CAPem = "other value"

	for i, tc := range []struct {
		serviceA *ServiceProxySettings
		serviceB *ServiceProxySettings
		expected bool
	}{
		{nil, nil, true},
		{serv, nil, false},
		{nil, serv, false},
		{serv, serv, true},
		{serv.Clone(), serv.Clone(), true},
		{serv, service("foo:"), false},
		{serv, service(serv.Url), true},
		{serviceHeader(serv), serviceHeader(serv), true},
		{servHeader2, servHeader2, true},
		{servHeader3, servHeader3, true},
		{serviceHeader(serv), servHeader2, false},
		{serviceHeader(serv), servHeader3, false},
		{serviceHeader(serv), serv, false},
		{serv, serviceHeader(serv), false},
		{serviceCa(serv), serviceCa(serv), true},
		{serviceCa(serv), serv, false},
		{serv, serviceCa(serv), false},
		{serviceCa(serviceHeader(serv)), serviceCa(serviceHeader(serv)), true},
		{servCa2, serviceCa(serv), false},
	} {
		if actual := tc.serviceA.IsEqual(tc.serviceB); tc.expected != actual {
			t.Fatalf(
				"test case %d: Comparison:\n-%+v\n+%+v\nexpected=%v: actual=%v",
				i, tc.serviceA, tc.serviceB, tc.expected, actual,
			)
		}
	}
}
