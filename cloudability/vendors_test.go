package cloudability

import "testing"


func TestNewVendorsEndpoint(t *testing.T) {
	e := newVendorsEndpoint("testapikey")
	if e.BaseURL.String() != api_v3_url {
		t.Fail()
	}
}