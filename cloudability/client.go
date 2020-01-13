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
type Client struct {
	*http.Client
	BaseURL      *url.URL
	UserAgent    string
	apikey       string
}

// This constructor creates a Cloudability client.
func NewClient(apikey string) *Client {
	c := &Client{
		Client:    &http.Client{Timeout: 10 * time.Second},
		UserAgent: "cloudability-sdk-go",
		apikey:    apikey,
	}
	return c
}

type APIError struct {
	Error errorDetail `json:"error"`
}

type errorDetail struct {
	Code float64 `json:"code"`
	Messages []string `json:"messages"`
}

type endpointI interface {
	buildURL(endpoint string) *url.URL
	newRequest(method string, u *url.URL, body interface{}) (*http.Request, error)
}

type endpoint struct {
	*Client
	BaseURL *url.URL
	EndpointPath string
}

type v1Endpoint struct {
	*endpoint
}

type v3Endpoint struct {
	*endpoint
}

func newEndpoint(c *Client, baseURL string, endpointPath string) *endpoint {
	e := &endpoint{Client: c}	
	e.BaseURL, _ = url.Parse(baseURL)
	e.EndpointPath = endpointPath
	return e
}

func newV1Endpoint(c *Client, endpointPath string) *v1Endpoint {
	return &v1Endpoint{newEndpoint(c, api_v1_url, endpointPath)}
}

func newV3Endpoint(c *Client, endpointPath string) *v3Endpoint {
	return &v3Endpoint{newEndpoint(c, api_v3_url, endpointPath)}
}

func (e* endpoint) buildURL(endpointPath string) *url.URL{
	rel := &url.URL{Path: path.Join(e.EndpointPath, endpointPath)}
	return e.BaseURL.ResolveReference(rel)
}

type resultTemplate struct {
	Result interface{} `json:"result"`
}

func (c *Client) get(e endpointI, endpoint string, result interface{}) error {
	return c.do(e, "GET", endpoint, nil, result)
}

func (c *Client) post(e endpointI, endpoint string, body interface{}, result interface{}) error {
	return c.do(e, "POST", endpoint, body, result)
}

func (c *Client) put(e endpointI, endpoint string, body interface{}) error {
	return c.do(e, "PUT", endpoint, body, nil)
}

func (c *Client) delete(e endpointI, endpoint string) error {
	return c.do(e, "DELETE", endpoint, nil, nil)
}

func (c *Client) do(e endpointI, method string, path string, body interface{}, result interface{}) error {
	u := e.buildURL(path)
	req, err := e.newRequest(method, u, body)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, result)
	return err
}

func (c *Client) doRequest(req *http.Request, result interface{}) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
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

func (c *Client) newRequest(method string, u *url.URL, body interface{}) (*http.Request, error) {
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
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (e *v1Endpoint) newRequest(method string, u *url.URL, body interface{}) (*http.Request, error) {
	req, err := e.Client.newRequest(method, u, body)
	q := req.URL.Query()
	q.Add("auth_token", e.apikey)
	req.URL.RawQuery = q.Encode()
	return req, err
}

func (e *v3Endpoint) newRequest(method string, u *url.URL, body interface{}) (*http.Request, error) {
	req, err := e.Client.newRequest(method, u, body)
	req.SetBasicAuth(e.apikey, "")
	return req, err
}
