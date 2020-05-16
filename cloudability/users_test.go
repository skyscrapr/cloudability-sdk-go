package cloudability

import (
	"testing"
	"net/url"
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
	testServer := testRequest(t, "GET", "/api/1/users", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Users()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetUsers()
	if err != nil{
		t.Fail()
	}
}

func TestGetUser(t *testing.T) {
	testServer := testRequest(t, "GET", "/api/1/users/1", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Users()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetUser(1)
	if err != nil{
		t.Fail()
	}
}

func TestNewUser(t *testing.T) {
	testServer := testRequest(t, "POST", "/api/1/users", nil)
	defer testServer.Close()
	user := &User{
		Email: "test.name@test.com.test",
		FullName: "Test Name",
		Role: "test_role",
		Restricted: false,
		// TODO: Fix this
		// SharedDimensionFilterSetIds: [0,1],
		DefaultDimensionFilterID: 1,
	}
	testClient := NewClient("testapikey")
	e := testClient.Users()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.NewUser(user)
	if err != nil{
		t.Fail()
	}
}

func TestUpdateUser(t *testing.T) {
	testServer := testRequest(t, "PUT", "/api/1/users/1", nil)
	defer testServer.Close()
	user := &User{
		ID: 1,
		Email: "test.name@test.com.test",
		FullName: "Test Name",
		Role: "test_role",
		Restricted: false,
		// TODO: Fix this
		// SharedDimensionFilterSetIds: [0,1],
		DefaultDimensionFilterID: 1,
	}
	testClient := NewClient("testapikey")
	e := testClient.Users()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.UpdateUser(user)
	if err != nil{
		t.Fail()
	}
}

func TestDeleteUser(t *testing.T) {
	testServer := testRequest(t, "DELETE", "/api/1/users/1", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Users()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.DeleteUser(1)
	if err != nil{
		t.Fail()
	}
}