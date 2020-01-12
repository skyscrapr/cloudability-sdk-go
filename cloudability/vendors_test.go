package cloudability

import (
	"testing"
	"net/url"
)

func TestNewVendorsEndpoint(t *testing.T) {
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	if e.BaseURL.String() != api_v3_url {
		t.Errorf("VendorsEndpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), api_v3_url)
	}
	if e.EndpointPath != vendors_endpoint {
		t.Errorf("VendorssEndpoint EndpointPath mismatch. Got %s. Expected %s", e.EndpointPath, vendors_endpoint)
	}
}

func TestGetVendors(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/vendors", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetVendors()
	if err != nil{
		t.Fail()
	}
}

func TestGetCredentials(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/vendors/aws/accounts", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetAccounts("aws")
	if err != nil{
		t.Fail()
	}
}

func TestGetCredential(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/vendors/aws/accounts/123456789012", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetAccount("aws", "123456789012")
	if err != nil{
		t.Fail()
	}
}

func TestVerifyAccount(t *testing.T) {
	testServer := testRequest(t, "POST", "/v3/vendors/aws/accounts/123456789012/verification", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.VerifyAccount("aws", "123456789012")
	if err != nil{
		t.Fail()
	}
}

func TestNewLinkedAccount(t *testing.T) {
	expectedBody := map[string]string{
		"vendorAccountId": "123456789012",
		"type": "aws_role",
	}
	testServer := testRequest(t, "POST", "/v3/vendors/aws/accounts", expectedBody)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.NewLinkedAccount("aws", &NewLinkedAccountParams{
		VendorAccountId: "123456789012", 
		Type: "aws_role",
	})
	if err != nil{
		t.Fail()
	}
}

func TestNewMasterAccount(t *testing.T) {
	expectedBody := map[string]string{
		"vendorAccountId": "123456789012",
		"type": "aws_role",
	}
	testServer := testRequest(t, "POST", "/v3/vendors/aws/accounts", expectedBody)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	var name, prefix string
	name = "Cloudability"
	prefix = "CostAndUsageReports"
	_, err := e.NewMasterAccount("aws", &NewMasterAccountParams{
		NewLinkedAccountParams: &NewLinkedAccountParams{
			VendorAccountId: "123456789012", 
			Type: "aws_role",
		},
		BucketName: "cloudability-123456789012",
		CostAndUsageReport: &CostAndUsageReport{
			Name: &name,
			Prefix: &prefix,
		},
	})
	if err != nil{
		t.Fail()
	}
}

func TestDeleteCredential(t *testing.T) {
	testServer := testRequest(t, "DELETE", "/v3/vendors/aws/accounts/123456789012", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.DeleteAccount("aws", "123456789012")
	if err != nil{
		t.Fail()
	}
}