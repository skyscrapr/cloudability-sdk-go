package cloudability

import (
	"encoding/json"
	"strconv"
)

const businessMappingsEndpoint = "/"

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
	Index                  int                         `json:"index"`
	Kind                   string                      `json:"kind"`
	Name                   string                      `json:"name"`
	DefaultValue           string                      `json:"defaultValue,omitempty"`
	DefaultValueExpression string                      `json:"defaultValueExpression,omitempty"`
	Statements             []*BusinessMappingStatement `json:"statements"`
	UpdatedAt              string
	NumberFormat           string `json:"numberFormat"`
	PreMatchExpression     string `json:"preMatchExpression,omitempty"`
}

type businessMappingPayload struct {
	Kind                   string                      `json:"kind"`
	Name                   string                      `json:"name"`
	DefaultValue           string                      `json:"defaultValue,omitempty"`
	DefaultValueExpression string                      `json:"defaultValueExpression,omitempty"`
	Statements             []*BusinessMappingStatement `json:"statements"`
	UpdatedAt              string
	NumberFormat           string `json:"numberFormat"`
	PreMatchExpression     string `json:"preMatchExpression,omitempty"`
}

// GetBusinessDimensions - Get a list of all existing business dimensions.
func (e *BusinessMappingsEndpoint) GetBusinessDimensions() ([]BusinessMapping, error) {
	var result v3Result[[]BusinessMapping]
	err := e.get(e, "business-mappings/dimensions", &result)
	return result.Result, err
}

// GetBusinessDimension - Get an existing business dimension by index.
func (e *BusinessMappingsEndpoint) GetBusinessDimension(index int) (*BusinessMapping, error) {
	var result v3Result[*BusinessMapping]
	err := e.get(e, "business-mappings/"+strconv.Itoa(index), &result)
	return result.Result, err
}

// NewBusinessDimension - Create a new business dimension.
func (e *BusinessMappingsEndpoint) NewBusinessDimension(dimension *BusinessMapping) (*BusinessMapping, error) {
	payload := new(businessMappingPayload)
	jsonBusinessMapping, _ := json.Marshal(dimension)
	json.Unmarshal(jsonBusinessMapping, payload)
	var result v3Result[*BusinessMapping]
	err := e.post(e, "business-mappings/", payload, &result)
	return result.Result, err
}

// UpdateBusinessDimension - Update an existing business dimension using given index.
func (e *BusinessMappingsEndpoint) UpdateBusinessDimension(dimension *BusinessMapping) error {
	payload := new(businessMappingPayload)
	jsonBusinessMapping, _ := json.Marshal(dimension)
	json.Unmarshal(jsonBusinessMapping, payload)
	return e.put(e, "business-mappings/"+strconv.Itoa(dimension.Index), payload)
}

// DeleteBusinessDimension - Delete an existing business dimension by index.
func (e *BusinessMappingsEndpoint) DeleteBusinessDimension(index int) error {
	return e.delete(e, "business-mappings/"+strconv.Itoa(index))
}

// GetBusinessMetrics - Get a list of all existing business metrics.
func (e *BusinessMappingsEndpoint) GetBusinessMetrics() ([]BusinessMapping, error) {
	var result v3Result[[]BusinessMapping]
	err := e.get(e, "internal/business-mappings/metrics", &result)
	return result.Result, err
}

// GetBusinessDimension - Get an existing business dimension by index.
func (e *BusinessMappingsEndpoint) GetBusinessMetric(index int) (*BusinessMapping, error) {
	var result v3Result[*BusinessMapping]
	err := e.get(e, "internal/business-mappings/"+strconv.Itoa(index)+"/metrics", &result)
	return result.Result, err
}

// NewBusinessMetric - Create a new business metric.
func (e *BusinessMappingsEndpoint) NewBusinessMetric(metric *BusinessMapping) (*BusinessMapping, error) {
	payload := new(businessMappingPayload)
	jsonBusinessMapping, _ := json.Marshal(metric)
	json.Unmarshal(jsonBusinessMapping, payload)
	var result v3Result[*BusinessMapping]
	err := e.post(e, "internal/business-mappings/metrics", payload, &result)
	return result.Result, err
}

// UpdateBusinessMetric - Update an existing business metric using given index.
func (e *BusinessMappingsEndpoint) UpdateBusinessMetric(metric *BusinessMapping) error {
	payload := new(businessMappingPayload)
	jsonBusinessMapping, _ := json.Marshal(metric)
	json.Unmarshal(jsonBusinessMapping, payload)
	return e.put(e, "internal/business-mappings/"+strconv.Itoa(metric.Index)+"/metrics", payload)
}

// DeleteBusinessMetric - Delete an existing business metric by index.
func (e *BusinessMappingsEndpoint) DeleteBusinessMetric(index int) error {
	return e.delete(e, "internal/business-mappings/"+strconv.Itoa(index)+"/metrics")
}
