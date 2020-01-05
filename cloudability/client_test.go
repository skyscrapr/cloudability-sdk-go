package cloudability

import "testing"


func TestNewCloudabilityClient(t *testing.T) {
	apikey := "testapikey"
	testClient := NewCloudabilityClient(apikey)
	if testClient.BusinessMappings.apikey != apikey {
		t.Errorf("BusinessMappings endpoint apikey mismatch")
	}
	if testClient.Views.apikey != apikey {
		t.Errorf("Views endpoint apikey mismatch")
	}
	if testClient.Users.apikey != apikey {
		t.Errorf("Users endpoint apikey mismatch")
	}
	if testClient.Vendors.apikey != apikey {
		t.Errorf("Vendors endpoint apikey mismatch")
	}
}

func TestNewCloudabilityV3Endpoint(t *testing.T) {
	e := newCloudabilityV3Endpoint("testapikey")
	if e.BaseURL.String() != api_v3_url {
		t.Fail()
	} 
}

func TestNewCloudabilityV1Endpoint(t *testing.T) {
	e := newCloudabilityV1Endpoint("testapikey")
	if e.BaseURL.String() != api_v1_url {
		t.Fail()
	} 
}

func TestV3NewRequest(t *testing.T) {
	testapikey := "testapikey"
	v3e := newCloudabilityV3Endpoint(testapikey)
	v3req, _ := v3e.newRequest("GET", "test-endpoint-path", nil)
	if u, p, ok := v3req.BasicAuth(); ok {
		if u != testapikey {
			t.Fail()
		}
		if p != "" {
			t.Fail()
		}
	}
}

func TestV1NewRequest(t *testing.T) {
	testapikey := "testapikey"
	v1e := newCloudabilityV1Endpoint(testapikey)	
	v1req, _ := v1e.newRequest("GET", "test-endpoint-path", nil)
	q := v1req.URL.Query()
	if q.Get("auth_token") != testapikey {
		t.Fail()
	}
}