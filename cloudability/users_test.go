package cloudability

import "testing"


func TestNewUsersEndpoint(t *testing.T) {
	e := newUsersEndpoint("testapikey")
	if e.BaseURL.String() != api_v1_url {
		t.Fail()
	}
}