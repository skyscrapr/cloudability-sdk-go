package cloudability

import (
	"net/url"
	"testing"
)

func TestNewVendorsEndpoint(t *testing.T) {
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	if e.BaseURL.String() != apiV3URL {
		t.Errorf("VendorsEndpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), apiV3URL)
	}
	if e.EndpointPath != vendorsEndpoint {
		t.Errorf("VendorssEndpoint EndpointPath mismatch. Got %s. Expected %s", e.EndpointPath, vendorsEndpoint)
	}
}

func TestGetVendors(t *testing.T) {
	expectedVendors := []Vendor{
		{
			Key:         "aws",
			Label:       "aws",
			Description: "aws",
		},
		{
			Key:         "azure",
			Label:       "azure",
			Description: "azure",
		},
	}
	result := v3Result[[]Vendor]{expectedVendors}
	testServer := testAPI(t, "GET", "/vendors", result)
	defer testServer.Close()
	testClient := testClient(t, testServer)
	e := testClient.Vendors()
	vendors, err := e.GetVendors()
	if err != nil {
		t.Errorf("Unexpected Errro: %s", err)
	}
	if vendors == nil {
		t.Fail()
	}
	testCheckStructEqual(t, vendors, expectedVendors)
}

func TestGetAccounts(t *testing.T) {
	expectedAccounts := []Account{
		{
			ID:                "1",
			VendorAccountName: "Account1",
			VendorAccountID:   "1",
			VendorKey:         "aws",
		},
	}
	result := v3Result[[]Account]{expectedAccounts}
	testServer := testAPI(t, "GET", "/vendors/aws/accounts", result)
	defer testServer.Close()
	testClient := testClient(t, testServer)
	e := testClient.Vendors()
	accounts, err := e.GetAccounts("aws")
	if err != nil {
		t.Fail()
	}
	if accounts == nil {
		t.Fail()
	}
	testCheckStructEqual(t, accounts, expectedAccounts)
}

func TestGetAccount(t *testing.T) {
	expectedAccount := &Account{
		ID:                "1",
		VendorAccountName: "Account1",
		VendorAccountID:   "1",
		VendorKey:         "aws",
	}
	result := v3Result[Account]{*expectedAccount}
	testServer := testAPI(t, "GET", "/vendors/aws/accounts/123456789012", result)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	account, err := e.GetAccount("aws", "123456789012")
	if err != nil {
		t.Fail()
	}
	testCheckStructEqual(t, account, expectedAccount)
}

func TestVerifyAccount(t *testing.T) {
	expectedAccount := &Account{
		ID:                "1",
		VendorAccountName: "Account1",
		VendorAccountID:   "1",
		VendorKey:         "aws",
	}
	result := v3Result[Account]{*expectedAccount}
	testServer := testAPI(t, "POST", "/vendors/aws/accounts/123456789012/verification", result)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	account, err := e.VerifyAccount("aws", "123456789012")
	if err != nil {
		t.Fail()
	}
	testCheckStructEqual(t, account, expectedAccount)
}

func TestNewLinkedAccount(t *testing.T) {
	expectedBody := map[string]string{
		"vendorAccountId": "123456789012",
		"type":            "aws_role",
	}
	testServer := testAPI(t, "POST", "/vendors/aws/accounts", expectedBody)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.NewLinkedAccount("aws", &NewLinkedAccountParams{
		VendorAccountID: "123456789012",
		Type:            "aws_role",
	})
	if err != nil {
		t.Fail()
	}
}

func TestNewMasterAccount(t *testing.T) {
	expectedBody := map[string]string{
		"vendorAccountId": "123456789012",
		"type":            "aws_role",
	}
	testServer := testAPI(t, "POST", "/vendors/aws/accounts", expectedBody)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	var name, prefix string
	name = "Cloudability"
	prefix = "CostAndUsageReports"
	_, err := e.NewMasterAccount("aws", &NewMasterAccountParams{
		NewLinkedAccountParams: &NewLinkedAccountParams{
			VendorAccountID: "123456789012",
			Type:            "aws_role",
		},
		BucketName: "cloudability-123456789012",
		CostAndUsageReport: &CostAndUsageReport{
			Name:   name,
			Prefix: prefix,
		},
	})
	if err != nil {
		t.Fail()
	}
}

func TestDeleteAccount(t *testing.T) {
	testServer := testAPI(t, "DELETE", "/vendors/aws/accounts/123456789012", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Vendors()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.DeleteAccount("aws", "123456789012")
	if err != nil {
		t.Fail()
	}
}
