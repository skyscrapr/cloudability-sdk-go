package cloudability

import "testing"


func TestNewViewsEndpoint(t *testing.T) {
	e := newViewsEndpoint("testapikey")
	if e.BaseURL.String() != api_v3_url {
		t.Fail()
	}
}