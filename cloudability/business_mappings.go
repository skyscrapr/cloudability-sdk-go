package cloudability

import (
	"strconv"
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

type businessMappingsResult struct {
	Result []BusinessMapping `json:"result"`
}

func (e businessMappingsEndpoint) BusinessMappings() ([]BusinessMapping, error) {
	var result businessMappingsResult
	err := e.get("", &result)
	return result.Result, err
}

type businessMappingResult struct {
	Result BusinessMapping `json:"result"`
}

func (e businessMappingsEndpoint) BusinessMapping(index int) (*BusinessMapping, error) {
	var result businessMappingResult
	err := e.get(strconv.Itoa(index), &result)
	return &result.Result, err
}