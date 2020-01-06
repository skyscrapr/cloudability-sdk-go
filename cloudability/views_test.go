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

func TestNewView(t *testing.T) {
	view := &View{
		Id: 1,

	}
	testServer := testRequest(t, "POST", "/v3/views", view)
	defer testServer.Close()

	e := newViewsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	_, err := e.NewView(view)
	if err != nil{
		t.Fail()
	}
}

func TestUpdateView(t *testing.T) {
	testServer := testRequest(t, "PUT", "/v3/views/1", nil)
	defer testServer.Close()
	view := &View{
		Id: 1,
		Title: "Test View",

	}
	e := newViewsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	err := e.UpdateView(view)
	if err != nil{
		t.Fail()
	}
}

func TestDeleteView(t *testing.T) {
	testServer := testRequest(t, "DELETE", "/v3/views/1", nil)
	defer testServer.Close()
	e := newViewsEndpoint("testapikey")
	e.BaseURL, _= url.Parse(testServer.URL)
	err := e.DeleteView(1)
	if err != nil{
		t.Fail()
	}
}