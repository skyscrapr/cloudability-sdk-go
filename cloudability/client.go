// Package cloudability provides a client for the cloudability api.
package cloudability

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	apiV1URL = "https://app.cloudability.com/api/1"
	apiV3URL = "https://api.cloudability.com/v3"
)

// Client - Cloudability client.
type Client struct {
	*http.Client
	V1BaseURL *url.URL
	V3BaseURL *url.URL
	UserAgent string
	apikey    string
}

// NewClient - This constructor creates a Cloudability client.
func NewClient(apikey string) *Client {
	c := &Client{
		Client:    &http.Client{Timeout: 30 * time.Second},
		UserAgent: "cloudability-sdk-go",
		apikey:    apikey,
	}
	c.V1BaseURL, _ = url.Parse(apiV1URL)
	c.V3BaseURL, _ = url.Parse(apiV3URL)
	return c
}

// APIError - Cloudability Error
type APIError struct {
	Error errorDetail `json:"error"`
}

type errorDetail struct {
	Status   float64  `json:"status"`
	Code     []string `json:"code"`
	Messages []string `json:"messages"`
}

type endpointI interface {
	buildURL(endpoint string) *url.URL
	newRequest(method string, u *url.URL, body interface{}) (*http.Request, error)
	doRequest(req *http.Request, result interface{}) (*http.Response, error)
}

type endpoint struct {
	*Client
	BaseURL      *url.URL
	EndpointPath string
}

type v1Endpoint struct {
	*endpoint
}

type v3Endpoint struct {
	*endpoint
}

func SetTimeout(c *Client, t time.Duration) {
	c.Timeout = t
}

func newEndpoint(c *Client, baseURL *url.URL, endpointPath string) *endpoint {
	e := &endpoint{
		Client:       c,
		BaseURL:      baseURL,
		EndpointPath: endpointPath,
	}
	return e
}

func newV1Endpoint(c *Client, endpointPath string) *v1Endpoint {
	return &v1Endpoint{newEndpoint(c, c.V1BaseURL, endpointPath)}
}

func newV3Endpoint(c *Client, endpointPath string) *v3Endpoint {
	return &v3Endpoint{newEndpoint(c, c.V3BaseURL, endpointPath)}
}

func (e *endpoint) buildURL(endpointPath string) *url.URL {
	u, err := url.Parse(endpointPath)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(e.EndpointPath, u.Path)
	u.Path = path.Join(e.BaseURL.Path, u.Path)
	return e.BaseURL.ResolveReference(u)
}

type v3ResultTemplate struct {
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
	resp, err := e.doRequest(req, result)
	if err == nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) doRequest(req *http.Request, result interface{}) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(bodyBytes))
	}
	return resp, nil
}

func (e *v1Endpoint) doRequest(req *http.Request, result interface{}) (*http.Response, error) {
	resp, err := e.Client.doRequest(req, result)
	if result != nil {
		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			log.Fatal(err)
		}
	}
	return resp, err
}

func (e *v3Endpoint) doRequest(req *http.Request, result interface{}) (*http.Response, error) {
	resp, err := e.Client.doRequest(req, result)
	if err != nil {
		return resp, err
	}
	if result != nil {
		resultTemplate := &v3ResultTemplate{
			Result: &result,
		}
		err = json.NewDecoder(resp.Body).Decode(resultTemplate)
		if err != nil {
			log.Fatal(err)
		}
	}
	return resp, err
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
