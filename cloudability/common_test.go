package cloudability

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func testV1API(t *testing.T, method string, path string, body interface{}) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			t.Errorf("Expected request method ‘%s’, got ‘%s’", method, req.Method)
		}
		if req.URL.Path != path {
			t.Errorf("Expected request to ‘%s’, got ‘%s’", path, req.URL.Path)
		}
		if body != nil {
			buf := new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(body)
			if err != nil {
				t.Errorf("Error converting body into JSON: %s", err)
			}
			rw.Write(buf.Bytes())
		} else {
			// TODO: Fix this. I don't think it's right
			rw.Write([]byte(`{}`))
		}
	}))
	return server
}

func testV3API(t *testing.T, method string, path string, body interface{}) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			t.Errorf("Expected request method ‘%s’, got ‘%s’", method, req.Method)
		}
		if req.URL.Path != path {
			t.Errorf("Expected request to ‘%s’, got ‘%s’", path, req.URL.Path)
		}
		if body != nil {
			result := v3ResultTemplate{
				Result: body,
			}
			buf := new(bytes.Buffer)
			err := json.NewEncoder(buf).Encode(result)
			if err != nil {
				t.Errorf("Error converting body into JSON: %s", err)
			}
			rw.Write(buf.Bytes())
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
