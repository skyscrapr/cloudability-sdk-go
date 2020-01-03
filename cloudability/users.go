package cloudability

import (
	"strconv"
)

type usersEndpoint struct {
	*cloudabilityV1Endpoint
}

func newUsersEndpoint(apikey string) *usersEndpoint {
	e := &usersEndpoint{newCloudabilityV1Endpoint(apikey)}
	e.EndpointPath = "/api/1/users/"
	return e
}

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
	Restricted bool `json:"restricted"`
	SharedDimensionFilterSetIds []int `json:"shared_dimension_filter_set_ids"`
	DefaultDimensionFilterId int `json:"default_dimension_filer_set_id"`
}

func (e usersEndpoint) Users() ([]User, error) {
	var users []User
	err := e.get("", &users)
	return users, err
}

func (e usersEndpoint) User(index int) (*User, error) {
	var user User
	err := e.get(strconv.Itoa(index), &user)
	return &user, err
}