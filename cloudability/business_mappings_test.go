package cloudability

import (
	"testing"
	"net/url"
)


func TestNewBusinessMappingsEndpoint(t *testing.T) {
	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	if e.BaseURL.String() != apiV3URL {
		t.Errorf("BusinessMappingsEndpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), apiV3URL)
	}
	if e.EndpointPath != businessMappingsEndpoint {
		t.Errorf("BusinessMappingsEndpoint EndpointPath mismatch. Got %s. Expected %s", e.EndpointPath, businessMappingsEndpoint)
	}
}

func TestGetBusinessMappings(t *testing.T) {
	testServer := testV1API(t, "GET", "/business-mappings", nil)
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
	testServer := testV1API(t, "GET", "/business-mappings/1", nil)
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
	testServer := testV1API(t, "POST", "/business-mappings", businessMapping)
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
	testServer := testV1API(t, "PUT", "/business-mappings/1", businessMapping)
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
	testServer := testV1API(t, "DELETE", "/business-mappings/1", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.DeleteBusinessMapping(1)
	if err != nil{
		t.Fail()
	}
}