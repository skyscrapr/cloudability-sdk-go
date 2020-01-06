package cloudability

import (
	"testing"
	"net/url"
)


func TestNewViewsEndpoint(t *testing.T) {
	e := newViewsEndpoint("testapikey")
	if e.BaseURL.String() != api_v3_url {
		t.Fail()
	}
}

func TestGetViews(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/views", nil)
	defer testServer.Close()

	e := newViewsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	_, err := e.GetViews()
	if err != nil{
		t.Fail()
	}
}

func TestGetView(t *testing.T) {
	testServer := testRequest(t, "GET", "/v3/views/1", nil)
	defer testServer.Close()

	e := newViewsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	_, err := e.GetView(1)
	if err != nil{
		t.Fail()
	}
}