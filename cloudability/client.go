package cloudability

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
	"path"
)

const (
	// endpoints
	api_v1_url = "https://app.cloudability.com"
	api_v3_url = "https://api.cloudability.com"
)

// Cloudability http client
type CloudabilityClient struct {
	BusinessMappings *businessMappingsEndpoint
	Users *usersEndpoint
	Vendors *vendorsEndpoint
	Views *viewsEndpoint
}

func NewCloudabilityClient(apikey string) *CloudabilityClient {
	c := &CloudabilityClient{}
	c.BusinessMappings = newBusinessMappingsEndpoint(apikey)
	c.Users = newUsersEndpoint(apikey)
	c.Vendors = newVendorsEndpoint(apikey)
	return c
}

type cloudabilityV3Endpoint struct {
	*http.Client
	BaseURL *url.URL
	EndpointPath string
	UserAgent string
	apikey string
}

type cloudabilityV1Endpoint struct {
	*cloudabilityV3Endpoint
}

func newCloudabilityV3Endpoint(apikey string) *cloudabilityV3Endpoint {
	e := &cloudabilityV3Endpoint{
		Client: &http.Client{Timeout: 10 * time.Second},
		UserAgent: "cloudability-sdk-go",
		apikey: apikey,
	}
	e.BaseURL, _ = url.Parse(api_v3_url)
	return e
}

func newCloudabilityV1Endpoint(apikey string) *cloudabilityV1Endpoint {
	e := &cloudabilityV1Endpoint{newCloudabilityV3Endpoint(apikey)}
	e.BaseURL, _ = url.Parse(api_v1_url)
	return e
}

type resultTemplate struct {
	Result interface{} `json:"result"`
}

func (e *cloudabilityV3Endpoint) get(endpoint string, result interface{}) error {
	endpointPath := path.Join(e.EndpointPath, endpoint)
	req, err := e.newVRequest("GET", endpointPath, nil)
	if err != nil {
		return err
	}
	resultTemplate := &resultTemplate{
		Result: &result,
	}
	_, err = e.execRequest(req, &resultTemplate)
	return err
	
}

func (e *cloudabilityV3Endpoint) execRequest(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := e.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Print(string(bodyBytes))
	}
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		log.Fatal(err)
	}
	return resp, nil
}

func (e *cloudabilityV3Endpoint) newRequest(method string, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := e.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", e.UserAgent)
	return req, nil
}

func (e *cloudabilityV1Endpoint) newVRequest(method string, path string, body interface{}) (*http.Request, error) {
	req, err := e.newRequest(method, path,body)
	q := req.URL.Query()
	q.Add("auth_token", e.apikey)
	req.URL.RawQuery = q.Encode()
	return req, err
}

func (e *cloudabilityV3Endpoint) newVRequest(method string, path string, body interface{}) (*http.Request, error) {
	req, err := e.newRequest(method, path, body)
	req.SetBasicAuth(e.apikey, "")
	return req, err
}
