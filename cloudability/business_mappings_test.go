package cloudability

import (
	"net/url"
	"os"
	"testing"
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
	testServer := testAPI(t, "GET", "/business-mappings", nil)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetBusinessMappings()
	if err != nil {
		t.Fail()
	}
}

func TestGetBusinessMapping(t *testing.T) {
	testServer := testAPI(t, "GET", "/business-mappings/1", nil)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetBusinessMapping(1)
	if err != nil {
		t.Fail()
	}
}

func TestNewBusinessMapping(t *testing.T) {
	businessMapping := &BusinessMapping{
		Kind:         "test-kind",
		Name:         "test-name",
		DefaultValue: "test-default-value",
		// Statements: [],
	}
	testServer := testAPI(t, "POST", "/business-mappings", businessMapping)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.NewBusinessMapping(businessMapping)
	if err != nil {
		t.Fail()
	}
}

func TestUpdateBusinessMapping(t *testing.T) {
	businessMapping := &BusinessMapping{
		Index:        1,
		Kind:         "test-kind",
		Name:         "test-name",
		DefaultValue: "test-default-value",
		// Statements: [],
	}
	testServer := testAPI(t, "PUT", "/business-mappings/1", businessMapping)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.UpdateBusinessMapping(businessMapping)
	if err != nil {
		t.Fail()
	}
}

func TestDeleteBusinessMapping(t *testing.T) {
	testServer := testAPI(t, "DELETE", "/business-mappings/1", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.BusinessMappings()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.DeleteBusinessMapping(1)
	if err != nil {
		t.Fail()
	}
}

func TestMultipleBusinessMapping(t *testing.T) {
	bm1 := BusinessMapping{
		Name:         "bm1",
		DefaultValue: "unallocated",
		Statements: []*BusinessMappingStatement{
			{
				MatchExpression: "EXISTS TAG['Business_Unit']",
				ValueExpression: "TAG['bm1']",
			},
			{
				MatchExpression: "DIMENSION['vendor_account_identifier'] == '999999999999'",
				ValueExpression: "'Mergers and Acquisitions'",
			},
		},
	}

	bm2 := BusinessMapping{
		Name:         "bm2",
		DefaultValue: "unallocated",
		Statements: []*BusinessMappingStatement{
			{
				MatchExpression: "EXISTS TAG['Business_Unit']",
				ValueExpression: "TAG['bm2']",
			},
			{
				MatchExpression: "DIMENSION['vendor_account_identifier'] == '999999999999'",
				ValueExpression: "'Mergers and Acquisitions'",
			},
		},
	}

	bm3 := BusinessMapping{
		Name:         "bm3",
		DefaultValue: "unallocated",
		Statements: []*BusinessMappingStatement{
			{
				MatchExpression: "EXISTS TAG['Business_Unit']",
				ValueExpression: "TAG['bm3']",
			},
			{
				MatchExpression: "DIMENSION['vendor_account_identifier'] == '999999999999'",
				ValueExpression: "'Mergers and Acquisitions'",
			},
		},
	}

	apikey := os.Getenv("CLOUDABILITY_APIKEY")
	client := NewClient(apikey)
	newbm1, err := client.BusinessMappings().NewBusinessMapping(&bm1)
	if err != nil {
		t.Fatalf("Error creating bm1: %s", err)
	}
	newbm2, err := client.BusinessMappings().NewBusinessMapping(&bm2)
	if err != nil {
		t.Fatalf("Error creating bm2: %s", err)
	}
	newbm3, err := client.BusinessMappings().NewBusinessMapping(&bm3)
	if err != nil {
		t.Fatalf("Error creating bm3: %s", err)
	}
	if newbm1.Name != "bm1" {
		t.Fatalf(`New Business Mapping. Got: %q, Want: %q`, newbm1.Name, "bm1")
	}
	if newbm2.Name != "bm2" {
		t.Fatalf(`New Business Mapping. Got: %q, Want: %q`, newbm2.Name, "bm2")
	}
	if newbm3.Name != "bm3" {
		t.Fatalf(`New Business Mapping. Got: %q, Want: %q`, newbm3.Name, "bm3")
	}
	err = client.BusinessMappings().DeleteBusinessMapping(newbm1.Index)
	if err != nil {
		t.Fatalf("Error deleting bm1: %s", err)
	}
	err = client.BusinessMappings().DeleteBusinessMapping(newbm2.Index)
	if err != nil {
		t.Fatalf("Error deleting bm1: %s", err)
	}
	err = client.BusinessMappings().DeleteBusinessMapping(newbm3.Index)
	if err != nil {
		t.Fatalf("Error deleting bm1: %s", err)
	}
}
