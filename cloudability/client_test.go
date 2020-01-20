package cloudability

import (
	"testing"
	"time"
	"net/url"
)

func TestNewClient(t *testing.T) {
	duration := 10 * time.Second
	userAgent := "cloudability-sdk-go"
	apikey := "testapikey"
	testClient := NewClient(apikey)
	if testClient.Client.Timeout != duration {
		t.Errorf("HTTP client timeout mismatch. Got %s. Expected %s", testClient.Client.Timeout, duration)
	}
	if testClient.UserAgent != userAgent {
		t.Errorf("Cloudability client useragent mismatch. Got %s. Expected %s", testClient.UserAgent, userAgent)
	}
	if testClient.apikey != apikey {
		t.Errorf("Cloudability client apikey mismatch. Got %s. Expected %s", testClient.apikey, apikey)
	}
}

func TestNewEndpoint(t *testing.T) {
	baseURL := "api.test.com"
	endpointPath := "test-endpoint"
	apikey := "testapikey"
	testClient := NewClient(apikey)
	e := newEndpoint(testClient, baseURL, endpointPath)
	if e.BaseURL.String() != baseURL {
		t.Errorf("Endpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), baseURL)
	}
	if e.EndpointPath != endpointPath {
		t.Errorf("Endpoint BaseURL mismatch. Got %s. Expected %s", e.EndpointPath, endpointPath)
	} 
}

func TestNewV1Endpoint(t *testing.T) {
	endpointPath := "test-endpoint"
	apikey := "testapikey"
	testClient := NewClient(apikey)
	e := newV1Endpoint(testClient, endpointPath)
	if e.BaseURL.String() != api_v1_url {
		t.Errorf("V3Endpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), api_v1_url)
	}
	if e.EndpointPath != endpointPath {
		t.Errorf("V3Endpoint BaseURL mismatch. Got %s. Expected %s", e.EndpointPath, endpointPath)
	} 
}

func TestNewV3Endpoint(t *testing.T) {
	endpointPath := "test-endpoint"
	apikey := "testapikey"
	testClient := NewClient(apikey)
	e := newV3Endpoint(testClient, endpointPath)
	if e.BaseURL.String() != api_v3_url {
		t.Errorf("V3Endpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), api_v3_url)
	}
	if e.EndpointPath != endpointPath {
		t.Errorf("V3Endpoint BaseURL mismatch. Got %s. Expected %s", e.EndpointPath, endpointPath)
	} 
}

func TestV3NewRequest(t *testing.T) {
	endpointPath := "test-endpoint"
	u := &url.URL{Path: "api.test.com/test-endpoint"}
	apikey := "testapikey"
	testClient := NewClient(apikey)
	v3e := newV3Endpoint(testClient, endpointPath)
	v3req, _ := v3e.newRequest("GET", u, nil)
	if u, p, ok := v3req.BasicAuth(); ok {
		if u != apikey {
			t.Errorf("v3Endpoint.NewRequest BasicAuth username mismatch. Got %s. Expected %s", u, apikey)
		}
		if p != "" {
			t.Errorf("v3Endpoint.NewRequest BasicAuth password mismatch. Got %s. Expected %s", p, "")
		}
	}
}

func TestV1NewRequest(t *testing.T) {
	endpointPath := "test-endpoint"
	u := &url.URL{Path: "api.test.com/test-endpoint"}
	apikey := "testapikey"
	testClient := NewClient(apikey)
	v1e := newV1Endpoint(testClient, endpointPath)
	v1req, _ := v1e.newRequest("GET", u, nil)
	q := v1req.URL.Query()
	if q.Get("auth_token") != apikey {
		t.Errorf("v1Endpoint.NewRequest Query parameter 'auth_token' mismatch. Got %s. Expected %s", q.Get("auth_token"), apikey)
	}
}