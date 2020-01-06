package cloudability

import (
	"testing"
	"net/url"
)


func TestNewBusinessMappingsEndpoint(t *testing.T) {
	e := newBusinessMappingsEndpoint("testapikey")
	if e.BaseURL.String() != api_v3_url {
		t.Fail()
	}
}

func TestGetBusinessMappings(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/business-mappings", nil)
	defer testServer.Close()

	e := newBusinessMappingsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	_, err := e.GetBusinessMappings()
	if err != nil{
		t.Fail()
	}
}

func TestGetBusinessMapping(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/business-mappings/1", nil)
	defer testServer.Close()

	e := newBusinessMappingsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	_, err := e.GetBusinessMapping(1)
	if err != nil{
		t.Fail()
	}
}

func TestNewBusinessMapping(t *testing.T) {
	businessMapping := &BusinessMapping{
		Name: "TestName",
		DefaultValue: "TestDefaultValue",
	}
	testServer := testRequest(t, "POST", "/v3/business-mappings", businessMapping)
	defer testServer.Close()

	e := newBusinessMappingsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	_, err := e.NewBusinessMapping(businessMapping)
	if err != nil{
		t.Fail()
	}
}