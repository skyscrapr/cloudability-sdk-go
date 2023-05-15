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
	if e.BaseURL.String() != apiV1URL {
		t.Errorf("UsersEndpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), apiV1URL)
	}
	if e.EndpointPath != usersEndpoint {
		t.Errorf("UsersEndpoint EndpointPath mismatch. Got %s. Expected %s", e.EndpointPath, usersEndpoint)
	}
}

func TestGetUsers(t *testing.T) {
	expectedUsers := []User{
		{
			ID:       1,
			Email:    "1@test",
			FullName: "1 Test",
		},
	}
	testServer := testV1API(t, "GET", "/users", &expectedUsers)
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
	testServer := testV1API(t, "GET", "/users/1", &expectedUser)
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
	if !reflect.DeepEqual(user, expectedUser) {
		suser, _ := json.MarshalIndent(user, "", "\t")
		sexpectedUser, _ := json.MarshalIndent(expectedUser, "", "\t")
		t.Errorf("Expected user '%s', got '%s'", sexpectedUser, suser)
	}
}

func TestNewUser(t *testing.T) {
	testServer := testV1API(t, "POST", "/users", nil)
	defer testServer.Close()
	user := &User{
		Email:      "test.name@test.com.test",
		FullName:   "Test Name",
		Role:       "test_role",
		Restricted: false,
		// TODO: Fix this
		// SharedDimensionFilterSetIds: [0,1],
		DefaultDimensionFilterID: 1,
	}
	testClient := NewClient("testapikey")
	e := testClient.Users()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.NewUser(user)
	if err != nil {
		t.Fail()
	}
}

func TestUpdateUser(t *testing.T) {
	testServer := testV1API(t, "PUT", "/users/1", nil)
	defer testServer.Close()
	user := &User{
		ID:         1,
		Email:      "test.name@test.com.test",
		FullName:   "Test Name",
		Role:       "test_role",
		Restricted: false,
		// TODO: Fix this
		// SharedDimensionFilterSetIds: [0,1],
		DefaultDimensionFilterID: 1,
	}
	testClient := NewClient("testapikey")
	e := testClient.Users()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.UpdateUser(user)
	if err != nil {
		t.Fail()
	}
}

func TestDeleteUser(t *testing.T) {
	testServer := testV1API(t, "DELETE", "/users/1", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Users()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.DeleteUser(1)
	if err != nil {
		t.Fail()
	}
}
