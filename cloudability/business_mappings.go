package cloudability

import (
	"strconv"
	"encoding/json"
)

type businessMappingsEndpoint struct {
	*cloudabilityV3Endpoint
}

func newBusinessMappingsEndpoint(apikey string) *businessMappingsEndpoint {
	e := &businessMappingsEndpoint{newCloudabilityV3Endpoint(apikey)}
	e.EndpointPath = "/v3/business-mappings/"
	return e
}

type BusinessMappingStatement struct {
	MatchExpression string `json:"matchExpression"`
	ValueExpression string `json:"valueExpression"`
}

type BusinessMapping struct {
	Index int `json:"index"`
	Kind string `json:"kind"`
	Name string `json:"name"`
	DefaultValue string `json:"defaultValue"`
	Statements []BusinessMappingStatement `json:"statements"`
	UpdatedAt string
}

type businessMappingPayload struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	DefaultValue string `json:"defaultValue"`
	Statements []BusinessMappingStatement `json:"statements"`
	UpdatedAt string
}

// Get a list of all existing business mappings.
func (e *businessMappingsEndpoint) GetBusinessMappings() ([]BusinessMapping, error) {
	var businessMappings []BusinessMapping
	err := e.get(e, "", &businessMappings)
	return businessMappings, err
}

// Get an existing business mapping by index.
func (e *businessMappingsEndpoint) GetBusinessMapping(index int) (*BusinessMapping, error) {
	var businessMapping BusinessMapping
	err := e.get(e, strconv.Itoa(index), &businessMapping)
	return &businessMapping, err
}

// Create a new business mapping.
func (e *businessMappingsEndpoint) NewBusinessMapping(businessMapping *BusinessMapping) (*BusinessMapping, error) {
	businessMappingPayload := new(businessMappingPayload)
	jsonBusinessMapping, _ := json.Marshal(businessMappingPayload)
	json.Unmarshal(jsonBusinessMapping, businessMappingPayload)
	var newBusinessMapping BusinessMapping
	err := e.post(e, "", businessMappingPayload, &newBusinessMapping)
	return &newBusinessMapping, err
}

// Update an existing business mapping using given index.
func (e *businessMappingsEndpoint) UpdateBusinessMapping(businessMapping *BusinessMapping) error {
	businessMappingPayload := new(businessMappingPayload)
	jsonBusinessMapping, _ := json.Marshal(businessMappingPayload)
	json.Unmarshal(jsonBusinessMapping, businessMappingPayload)
	return e.put(e, strconv.Itoa(businessMapping.Index), businessMappingPayload)
}

// Delete an existing business mapping by index.
func (e *businessMappingsEndpoint) DeleteBusinessMapping(index int) error {
	err := e.delete(e, strconv.Itoa(index))
	return err
}