package cloudability

import (
	"strconv"
	"encoding/json"
)

const businessMappingsEndpoint = "/v3/business-mappings/"

// BusinessMappingsEndpoint - Cloudability BusinessMappingsEndpoint
type BusinessMappingsEndpoint struct {
	*v3Endpoint
}

// BusinessMappings - return a Cloudability BusinessMappingsEndpoint
func (c *Client) BusinessMappings() *BusinessMappingsEndpoint {
	return &BusinessMappingsEndpoint{newV3Endpoint(c, businessMappingsEndpoint)}
}

// BusinessMappingStatement - Cloudability Business Mapping Statement
type BusinessMappingStatement struct {
	MatchExpression string `json:"matchExpression"`
	ValueExpression string `json:"valueExpression"`
}

// BusinessMapping - Cloudability BusinessMapping
type BusinessMapping struct {
	Index int `json:"index"`
	Kind string `json:"kind"`
	Name string `json:"name"`
	DefaultValue string `json:"defaultValue"`
	Statements []*BusinessMappingStatement `json:"statements"`
	UpdatedAt string
}

type businessMappingPayload struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	DefaultValue string `json:"defaultValue"`
	Statements []*BusinessMappingStatement `json:"statements"`
	UpdatedAt string
}

// GetBusinessMappings - Get a list of all existing business mappings.
func (e *BusinessMappingsEndpoint) GetBusinessMappings() ([]BusinessMapping, error) {
	var businessMappings []BusinessMapping
	err := e.get(e, "", &businessMappings)
	return businessMappings, err
}

// GetBusinessMapping - Get an existing business mapping by index.
func (e *BusinessMappingsEndpoint) GetBusinessMapping(index int) (*BusinessMapping, error) {
	var businessMapping BusinessMapping
	err := e.get(e, strconv.Itoa(index), &businessMapping)
	return &businessMapping, err
}

// NewBusinessMapping - Create a new business mapping.
func (e *BusinessMappingsEndpoint) NewBusinessMapping(businessMapping *BusinessMapping) (*BusinessMapping, error) {
	businessMappingPayload := new(businessMappingPayload)
	jsonBusinessMapping, _ := json.Marshal(businessMapping)
	json.Unmarshal(jsonBusinessMapping, businessMappingPayload)
	var newBusinessMapping BusinessMapping
	err := e.post(e, "", businessMappingPayload, &newBusinessMapping)
	return &newBusinessMapping, err
}

// UpdateBusinessMapping - Update an existing business mapping using given index.
func (e *BusinessMappingsEndpoint) UpdateBusinessMapping(businessMapping *BusinessMapping) error {
	businessMappingPayload := new(businessMappingPayload)
	jsonBusinessMapping, _ := json.Marshal(businessMappingPayload)
	json.Unmarshal(jsonBusinessMapping, businessMappingPayload)
	return e.put(e, strconv.Itoa(businessMapping.Index), businessMappingPayload)
}

// DeleteBusinessMapping - Delete an existing business mapping by index.
func (e *BusinessMappingsEndpoint) DeleteBusinessMapping(index int) error {
	err := e.delete(e, strconv.Itoa(index))
	return err
}