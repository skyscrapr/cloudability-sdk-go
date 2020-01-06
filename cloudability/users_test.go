package cloudability

import (
	"testing"
	"net/url"
)


func TestNewUsersEndpoint(t *testing.T) {
	e := newUsersEndpoint("testapikey")
	if e.BaseURL.String() != api_v1_url {
		t.Fail()
	}
}


func TestGetUsers(t *testing.T) {
	testServer := testRequest(t, "GET", "/api/1/users", nil)
	defer testServer.Close()

	e := newUsersEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	_, err := e.GetUsers()
	if err != nil{
		t.Fail()
	}
}

func TestGetUser(t *testing.T) {
	testServer := testRequest(t, "GET", "/api/1/users/1", nil)
	defer testServer.Close()

	e := newUsersEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
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
		DefaultDimensionFilterId: 1,
	}
	e := newUsersEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	err := e.NewUser(user)
	if err != nil{
		t.Fail()
	}
}

func TestUpdateUser(t *testing.T) {
	testServer := testRequest(t, "PUT", "/api/1/users/1", nil)
	defer testServer.Close()
	user := &User{
		Id: 1,
		Email: "test.name@test.com.test",
		FullName: "Test Name",
		Role: "test_role",
		Restricted: false,
		// TODO: Fix this
		// SharedDimensionFilterSetIds: [0,1],
		DefaultDimensionFilterId: 1,
	}
	e := newUsersEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	err := e.UpdateUser(user)
	if err != nil{
		t.Fail()
	}
}

func TestDeleteUser(t *testing.T) {
	testServer := testRequest(t, "DELETE", "/api/1/users/1", nil)
	defer testServer.Close()
	e := newUsersEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	err := e.DeleteUser(1)
	if err != nil{
		t.Fail()
	}
}