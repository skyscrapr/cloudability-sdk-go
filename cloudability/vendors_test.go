package cloudability

import (
	"testing"
	"net/url"
)


func TestNewVendorsEndpoint(t *testing.T) {
	e := newVendorsEndpoint("testapikey")
	if e.BaseURL.String() != api_v3_url {
		t.Fail()
	}
}

func TestGetVendors(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/vendors", nil)
	defer testServer.Close()

	e := newVendorsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	_, err := e.GetVendors()
	if err != nil{
		t.Fail()
	}
}

func TestGetCredentials(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/vendors/aws/accounts", nil)
	defer testServer.Close()

	e := newVendorsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	_, err := e.GetAccounts("aws")
	if err != nil{
		t.Fail()
	}
}

func TestGetCredential(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/vendors/aws/accounts/123456789012", nil)
	defer testServer.Close()

	e := newVendorsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	_, err := e.GetAccount("aws", 123456789012)
	if err != nil{
		t.Fail()
	}
}

func TestVerifyCredential(t *testing.T) {
	testServer := testRequest(t, "POST", "/v3/vendors/aws/accounts/123456789012/verification", nil)
	defer testServer.Close()

	e := newVendorsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	err := e.VerifyAccount("aws", "123456789012")
	if err != nil{
		t.Fail()
	}
}

func TestNewCredential(t *testing.T) {
	expectedBody := map[string]string{
		"vendorAccountId": "123456789012",
		"type": "aws_role",
	}
	testServer := testRequest(t, "POST", "/v3/vendors/aws/accounts", expectedBody)
	defer testServer.Close()

	e := newVendorsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	_, err := e.NewAccount("aws", "123456789012", "aws_role")
	if err != nil{
		t.Fail()
	}
}

func TestDeleteCredential(t *testing.T) {
	testServer := testRequest(t, "DELETE", "/v3/vendors/aws/accounts/123456789012", nil)
	defer testServer.Close()

	e := newVendorsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	err := e.DeleteAccount("aws", "123456789012")
	if err != nil{
		t.Fail()
	}
}