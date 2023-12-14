package cloudability

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func testAPI(t *testing.T, method string, path string, body interface{}) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			t.Errorf("Expected request method ‘%s’, got ‘%s’", method, req.Method)
		}
		if req.URL.Path != path {
			t.Errorf("Expected request to ‘%s’, got ‘%s’", path, req.URL.Path)
		}

		switch b := body.(type) {
		case string:
			rw.Header().Set("Content-Type", "text/plain")
			rw.Write([]byte(b))
		default:
			rw.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(rw).Encode(body); err != nil {
				t.Errorf("Error encoding response: %s", err)
			}
		}
	}))
	return server
}

func testClient(t *testing.T, s *httptest.Server) *Client {
	client := NewClient("testapikey")
	u, err := url.Parse(s.URL)
	if err != nil {
		t.Fail()
	}
	client.V1BaseURL = u
	client.V3BaseURL = u
	return client
}

func testCheckStructEqual(t *testing.T, got interface{}, expected interface{}) {
	if !reflect.DeepEqual(got, expected) {
		sgot, _ := json.MarshalIndent(got, "", "\t")
		sexpected, _ := json.MarshalIndent(expected, "", "\t")
		t.Errorf("Expected '%s', got '%s'", sexpected, sgot)
	}
}
