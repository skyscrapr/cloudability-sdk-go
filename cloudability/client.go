// Cloudability package provides a client for the cloudability api.
package cloudability

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"
	"errors"
)

const (
	api_v1_url = "https://app.cloudability.com"
	api_v3_url = "https://api.cloudability.com"
)

// Cloudability client.
type CloudabilityClient struct {
	BusinessMappings *businessMappingsEndpoint
	Users            *usersEndpoint
	Vendors          *vendorsEndpoint
	Views            *viewsEndpoint
}

// This constructor creates a client that provides access to all available cloudability endpoints.
func NewCloudabilityClient(apikey string) *CloudabilityClient {
	c := &CloudabilityClient{}
	c.BusinessMappings = newBusinessMappingsEndpoint(apikey)
	c.Users = newUsersEndpoint(apikey)
	c.Vendors = newVendorsEndpoint(apikey)
	c.Views = newViewsEndpoint(apikey)
	return c
}

type cloudabilityEndpointI interface {
	newRequest(method string, path string, body interface{}) (*http.Request, error)
}

type cloudabilityEndpoint struct {
	*http.Client
	BaseURL      *url.URL
	EndpointPath string
	UserAgent    string
	apikey       string
}

type cloudabilityV3Endpoint struct {
	*cloudabilityEndpoint
}

type cloudabilityV1Endpoint struct {
	*cloudabilityEndpoint
}

type resultTemplate struct {
	Result interface{} `json:"result"`
}

func newCloudabilityEndpoint(apikey string) *cloudabilityEndpoint {
	e := &cloudabilityEndpoint{
		Client:    &http.Client{Timeout: 10 * time.Second},
		UserAgent: "cloudability-sdk-go",
		apikey:    apikey,
	}
	return e
}

func newCloudabilityV3Endpoint(apikey string) *cloudabilityV3Endpoint {
	e := &cloudabilityV3Endpoint{newCloudabilityEndpoint(apikey)}
	e.BaseURL, _ = url.Parse(api_v3_url)
	return e
}

func newCloudabilityV1Endpoint(apikey string) *cloudabilityV1Endpoint {
	e := &cloudabilityV1Endpoint{newCloudabilityEndpoint(apikey)}
	e.BaseURL, _ = url.Parse(api_v1_url)
	return e
}

func (e *cloudabilityEndpoint) get(ce cloudabilityEndpointI, endpoint string, result interface{}) error {
	return e.exec(ce, "GET", endpoint, nil, result)
}

func (e *cloudabilityEndpoint) post(ce cloudabilityEndpointI, endpoint string, body interface{}, result interface{}) error {
	return e.exec(ce, "POST", endpoint, body, result)
}

func (e *cloudabilityEndpoint) put(ce cloudabilityEndpointI, endpoint string, body interface{}) error {
	return e.exec(ce, "PUT", endpoint, body, nil)
}

func (e *cloudabilityEndpoint) delete(ce cloudabilityEndpointI, endpoint string) error {
	return e.exec(ce, "DELETE", endpoint, nil, nil)
}

func (e *cloudabilityEndpoint) exec(ce cloudabilityEndpointI, method string, endpoint string, body interface{}, result interface{}) error {
	endpointPath := path.Join(e.EndpointPath, endpoint)
	req, err := ce.newRequest(method, endpointPath, body)
	if err != nil {
		return err
	}
	_, err = e.execRequest(req, result)
	return err
}

func (e *cloudabilityEndpoint) execRequest(req *http.Request, result interface{}) (*http.Response, error) {
	resp, err := e.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(bodyBytes))
	}
	if result != nil {
		resultTemplate := &resultTemplate{
			Result: &result,
		}
		err = json.NewDecoder(resp.Body).Decode(resultTemplate)
		if err != nil {
			log.Fatal(err)
		}
	}
	return resp, nil
}

func (e *cloudabilityEndpoint) newRequest(method string, path string, body interface{}) (*http.Request, error) {
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

func (e *cloudabilityV3Endpoint) newRequest(method string, path string, body interface{}) (*http.Request, error) {
	req, err := e.cloudabilityEndpoint.newRequest(method, path, body)
	req.SetBasicAuth(e.apikey, "")
	return req, err
}

func (e *cloudabilityV1Endpoint) newRequest(method string, path string, body interface{}) (*http.Request, error) {
	req, err := e.cloudabilityEndpoint.newRequest(method, path, body)
	q := req.URL.Query()
	q.Add("auth_token", e.apikey)
	req.URL.RawQuery = q.Encode()
	return req, err
}
