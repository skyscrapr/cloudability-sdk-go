package cloudability

import (
	"net/url"
	"testing"
)

func TestNewViewsEndpoint(t *testing.T) {
	testClient := NewClient("testapikey")
	e := testClient.Views()
	if e.BaseURL.String() != apiV3URL {
		t.Errorf("ViewsEndpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), apiV3URL)
	}
	if e.EndpointPath != viewsEndpoint {
		t.Errorf("ViewsEndpoint EndpointPath mismatch. Got %s. Expected %s", e.EndpointPath, viewsEndpoint)
	}
}

func TestGetViews(t *testing.T) {
	testServer := testAPI(t, "GET", "/views", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Views()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetViews()
	if err != nil {
		t.Fail()
	}
}

func TestGetView(t *testing.T) {
	testServer := testAPI(t, "GET", "/views/1", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Views()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetView("1")
	if err != nil {
		t.Fail()
	}
}

func TestNewView(t *testing.T) {
	view := &View{
		ID: "1",
	}
	testServer := testAPI(t, "POST", "/views", view)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Views()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.NewView(view)
	if err != nil {
		t.Fail()
	}
}

func TestUpdateView(t *testing.T) {
	testServer := testAPI(t, "PUT", "/views/1", nil)
	defer testServer.Close()
	view := &View{
		ID:    "1",
		Title: "Test View",
	}
	testClient := NewClient("testapikey")
	e := testClient.Views()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.UpdateView(view)
	if err != nil {
		t.Fail()
	}
}

func TestDeleteView(t *testing.T) {
	testServer := testAPI(t, "DELETE", "/views/1", nil)
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Views()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.DeleteView("1")
	if err != nil {
		t.Fail()
	}
}
