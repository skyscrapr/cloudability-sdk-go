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

func (e *businessMappingsEndpoint) GetBusinessMappings() ([]BusinessMapping, error) {
	var businessMappings []BusinessMapping
	err := e.get("", &businessMappings)
	return businessMappings, err
}

func (e *businessMappingsEndpoint) GetBusinessMapping(index int) (*BusinessMapping, error) {
	var businessMapping BusinessMapping
	err := e.get(strconv.Itoa(index), &businessMapping)
	return &businessMapping, err
}

func (e *businessMappingsEndpoint) NewBusinessMapping(businessMapping *BusinessMapping) (*BusinessMapping, error) {
	var newBusinessMapping BusinessMapping
	err := e.post(strconv.Itoa(businessMapping.Index), businessMapping, &newBusinessMapping)
	return &newBusinessMapping, err
}
