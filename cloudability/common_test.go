package cloudability

import (
	"testing"
	"net/http"
	"net/http/httptest"
)


func testRequest(t *testing.T, method string, path string, body interface{}) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			t.Errorf("Expected request method ‘%s’, got ‘%s’", method, req.Method)
		}
		if req.URL.EscapedPath() != path {
			t.Errorf("Expected request to ‘%s’, got ‘%s’", path, req.URL.EscapedPath())
		}
		// TODO: Fix this
		// if body != nil {
		// 	jsonReq, err := simplejson.NewFromReader(req.Body)
    	// 	if err != nil {
      	// 		t.Errorf("Error while reading request JSON: %s", err)
    	// 	}
		// 	if !reflect.DeepEqual(jsonReq, req.Body) {
		// 		t.Errorf("Expected body ‘%s’, got ‘%s’", body, req.Body)
		// 	}
		// }
		rw.Write([]byte(`{}`))
	}))
	return server
}