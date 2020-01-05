package cloudability

import "testing"


func TestNewBusinessMappingsEndpoint(t *testing.T) {
	e := newBusinessMappingsEndpoint("testapikey")
	if e.BaseURL.String() != api_v3_url {
		t.Fail()
	}
}