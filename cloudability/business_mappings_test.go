package cloudability

import (
	"testing"
	"net/url"
)


func TestNewBusinessMappingsEndpoint(t *testing.T) {
	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	if e.BaseURL.String() != api_v3_url {
		t.Errorf("BusinessMappingsEndpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), api_v3_url)
	}
	if e.EndpointPath != business_mappings_endpoint {
		t.Errorf("BusinessMappingsEndpoint EndpointPath mismatch. Got %s. Expected %s", e.EndpointPath, business_mappings_endpoint)
	}
}

func TestGetBusinessMappings(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/business-mappings", nil)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetBusinessMappings()
	if err != nil{
		t.Fail()
	}
}

func TestGetBusinessMapping(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/business-mappings/1", nil)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetBusinessMapping(1)
	if err != nil{
		t.Fail()
	}
}

func TestNewBusinessMapping(t *testing.T) {
	businessMapping := &BusinessMapping{
		Kind: "test-kind",
		Name: "test-name",
		DefaultValue: "test-default-value",
		// Statements: [], 
	}
	testServer := testRequest(t, "POST", "/v3/business-mappings", businessMapping)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.NewBusinessMapping(businessMapping)
	if err != nil{
		t.Fail()
	}
}

func TestUpdateBusinessMapping(t *testing.T) {
	businessMapping := &BusinessMapping{
		Index: 1,
		Kind: "test-kind",
		Name: "test-name",
		DefaultValue: "test-default-value",
		// Statements: [], 
	}
	testServer := testRequest(t, "PUT", "/v3/business-mappings/1", businessMapping)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.UpdateBusinessMapping(businessMapping)
	if err != nil{
		t.Fail()
	}
}

func TestDeleteBusinessMapping(t *testing.T) {
	testServer := testRequest(t, "DELETE", "/v3/business-mappings/1", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.DeleteBusinessMapping(1)
	if err != nil{
		t.Fail()
	}
}