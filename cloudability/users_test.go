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