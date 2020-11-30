package cloudability

import (
	"testing"
	"net/url"
)


func TestNewAccountGroupsEndpoint(t *testing.T) {
	testClient := NewClient("testapikey")
	e := testClient.AccountGroups()
	if e.BaseURL.String() != apiV1URL {
		t.Errorf("AccountGroupsEndpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), apiV1URL)
	}
	if e.EndpointPath != accountGroupsEndpoint {
		t.Errorf("AccountGroupsEndpoint EndpointPath mismatch. Got %s. Expected %s", e.EndpointPath, accountGroupsEndpoint)
	}
}

func TestGetAccountGroups(t *testing.T) {
	expectedAccountGroups := []AccountGroup{
		{
			ID: 1,
			Name: "red",
			Position: 1,
			AccountGroupEntryValues: []string {
				"red1",
				"red2",
				"red3",
			},
		},
		{
			ID: 2,
			Name: "orange",
			Position: 2,
			AccountGroupEntryValues: []string {
				"orange1",
				"orange2",
			},
		},
		{
			ID: 3,
			Name: "mauve",
			Position: 3,
			AccountGroupEntryValues: []string {},
		},
	}
	testServer := testV1API(t, "GET", "/account_groups", expectedAccountGroups)
	defer testServer.Close()
	testClient := testClient(t, testServer)
	e := testClient.AccountGroups()
	accountGroups, err := e.GetAccountGroups()
	if err != nil{
		t.Fail()
	}
	if accountGroups == nil{
		t.Fail()
	}
	// testCheckStructEqual(t, accountGroups, expectedAccountGroups)
}

func TestGetAccountGroup(t *testing.T) {
	expectedAccountGroup := &AccountGroup{
		ID: 2,
		Name: "orange",
		Position: 2,
		AccountGroupEntryValues: []string {
			"orange1",
			"orange2",
		},
	}
	testServer := testV1API(t, "GET", "/account_groups/2", &expectedAccountGroup)
	defer testServer.Close()
	testClient := testClient(t, testServer)
	e := testClient.AccountGroups()
	accountGroup, err := e.GetAccountGroup(2)
	if err != nil{
		t.Fail()
	}
	if accountGroup == nil{
		t.Fail()
	}
	testCheckStructEqual(t, accountGroup, expectedAccountGroup)
}


func TestNewAccountGroup(t *testing.T) {
	testServer := testV1API(t, "POST", "/account_groups", nil)
	defer testServer.Close()
	accountGroup := &AccountGroup{
		Name: "purple",
		Position: 5,
	}
	testClient := NewClient("testapikey")
	e := testClient.AccountGroups()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.NewAccountGroup(accountGroup)
	if err != nil{
		t.Fail()
	}
}

func TestUpdateAccountGroup(t *testing.T) {
	testServer := testV1API(t, "PUT", "/account_groups/1", nil)
	defer testServer.Close()
	accountGroup := &AccountGroup{
		ID: 1,         
		Name: "more purple",
	}
	testClient := NewClient("testapikey")
	e := testClient.AccountGroups()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.UpdateAccountGroup(accountGroup)
	if err != nil{
		t.Fail()
	}
}

func TestDeleteAccountGroup(t *testing.T) {
	testServer := testV1API(t, "DELETE", "/account_groups/1", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.AccountGroups()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.DeleteAccountGroup(1)
	if err != nil{
		t.Fail()
	}
}