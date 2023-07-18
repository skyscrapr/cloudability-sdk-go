package cloudability

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"
)

func TestNewUsersEndpoint(t *testing.T) {
	testClient := NewClient("testapikey")
	e := testClient.Users()
	if e.BaseURL.String() != apiV3URL {
		t.Errorf("UsersEndpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), apiV1URL)
	}
	if e.EndpointPath != usersEndpoint {
		t.Errorf("UsersEndpoint EndpointPath mismatch. Got %s. Expected %s", e.EndpointPath, usersEndpoint)
	}
}

// [{"id":91271,"frontdoor_user_id":"53f0db8c-d038-4525-9472-7e3b692ab9ff","frontdoor_login":"hwallez@apptio.com","email":"hwallez@apptio.com","full_name":"Harry Wallez","default_dimension_filter_set_id":0,"shared_dimension_filter_set_ids":[225582,354716,354717,354718,354719,354720,354721,373379,373380,373381]},{"id":91294,"frontdoor_user_id":"9e937cc8-3ea2-4e91-8c10-a07feb9812a3","frontdoor_login":"jplatt@apptio.com","email":"jplatt@apptio.com","full_name":"NULL","default_dimension_filter_set_id":0,"shared_dimension_filter_set_ids":[225582,354716,354717,354718,354719,354720,354721,373379,373380,373381]},{"id":91713,"frontdoor_user_id":"f5a6cb08-3e1d-44aa-b702-4925756e88c8","frontdoor_login":"abhjaw@amazon.com","email":"abhjaw@amazon.com","full_na
func TestGetUsers(t *testing.T) {
	expectedUsers := []User{
		{
			ID:                          1,
			FrontdoorUserId:             "test",
			FrontdoorLogin:              "1@test",
			Email:                       "1@test",
			FullName:                    "1 Test",
			DefaultDimensionFilterSetID: 0,
			SharedDimensionFilterSetIDs: []int{225582},
		},
	}
	testServer := testAPI(t, "GET", "/users", expectedUsers)
	defer testServer.Close()
	testClient := testClient(t, testServer)
	e := testClient.Users()
	users, err := e.GetUsers()
	if err != nil {
		t.Fail()
	}
	if users == nil {
		t.Fail()
	}
	if !reflect.DeepEqual(users, expectedUsers) {
		susers, _ := json.MarshalIndent(users, "", "\t")
		sexpectedUsers, _ := json.MarshalIndent(expectedUsers, "", "\t")
		t.Errorf("Expected user '%s', got '%s'", sexpectedUsers, susers)
	}
}

func TestGetUser(t *testing.T) {
	expectedUser := &User{
		ID:       1,
		Email:    "1@test",
		FullName: "1 Test",
	}
	testServer := testAPI(t, "GET", "/users/1", expectedUser)
	defer testServer.Close()
	testClient := testClient(t, testServer)
	e := testClient.Users()
	user, err := e.GetUser(1)
	if err != nil {
		t.Fail()
	}
	if user == nil {
		t.Fail()
	}
	testCheckStructEqual(t, user, expectedUser)
}

func TestUpdateUser(t *testing.T) {
	testServer := testAPI(t, "PUT", "/users/1", nil)
	defer testServer.Close()
	user := &User{
		ID:                          1,
		Email:                       "test.name@test.com.test",
		FullName:                    "Test Name",
		SharedDimensionFilterSetIDs: []int{0, 1},
		DefaultDimensionFilterSetID: 0,
	}
	testClient := NewClient("testapikey")
	e := testClient.Users()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.UpdateUser(user)
	if err != nil {
		t.Fail()
	}
}
