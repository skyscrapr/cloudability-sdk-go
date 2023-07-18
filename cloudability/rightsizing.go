package cloudability

import (
	"fmt"
)

const rightsizingEndpoint = "/v3/rightsizing/"

// RightsizingEndpoint Cloudability Rightsizing Endpoint
type RightsizingEndpoint struct {
	*v3Endpoint
}

// Rightsizing Endpoint
func (c *Client) Rightsizing() *RightsizingEndpoint {
	return &RightsizingEndpoint{newV3Endpoint(c, rightsizingEndpoint)}
}

// Resource Cloudability Rightsizing Resource
type Resource struct {
	Vendor             string           `json:"vendor"`
	Service            string           `json:"service"`
	ResourceIdentifier string           `json:"resourceIdentifier"`
	Recommendations    []Recommendation `json:"recommendations"`
}

// Recommendation Cloudability Rightsizing Resource Recommendation
type Recommendation struct {
	Action string `json:"action"`
}

// GetResource return a Rightsizing Resource
func (e RightsizingEndpoint) GetResource(vendor string, service string, resourceID string) (*Resource, error) {
	var result v3Result[*Resource]
	err := e.get(e, fmt.Sprintf("%s/recommendations/%s?filters=resourceIdentifier==%s", vendor, service, resourceID), &result)
	return result.Result, err
}
