package cloudability

import (
	"strconv"
	"encoding/json"
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
	Id int `json:"id,omitempty"`
	Email string `json:"email"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
	Restricted bool `json:"restricted"`
	SharedDimensionFilterSetIds []int `json:"shared_dimension_filter_set_ids"`
	DefaultDimensionFilterId int `json:"default_dimension_filer_set_id"`
}

func (e *usersEndpoint) GetUsers() ([]User, error) {
	var users []User
	err := e.get(e, "", &users)
	return users, err
}

func (e *usersEndpoint) GetUser(id int) (*User, error) {
	var user User
	err := e.get(e, strconv.Itoa(id), &user)
	return &user, err
}

type userNewPayloadWrapper struct {
	User *userNewPayload `json:"user"`
}

type userNewPayload struct {
	Email string `json:"email"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
	Restricted bool `json:"restricted"`
	SharedDimensionFilterSetIds []int `json:"shared_dimension_filter_set_ids"`
	DefaultDimensionFilterId int `json:"default_dimension_filer_set_id"`
}

func (e *usersEndpoint) NewUser(user *User) error {
	userPayload := new(userNewPayload)
	jsonUser, _ := json.Marshal(user)
    json.Unmarshal(jsonUser, userPayload)
	
	userPayloadWrapper := &userNewPayloadWrapper{
		User: userPayload,
	}
	return e.post(e, "", userPayloadWrapper, nil)
}

type userUpdatePayloadWrapper struct {
	User *userUpdatePayload `json:"user"`
}

type userUpdatePayload struct {
	FullName string `json:"full_name"`
	Role string `json:"role"`
	Restricted bool `json:"restricted"`
	SharedDimensionFilterSetIds []int `json:"shared_dimension_filter_set_ids"`
	DefaultDimensionFilterId int `json:"default_dimension_filer_set_id"`
}

func (e *usersEndpoint) UpdateUser(user *User) error {
	userPayload := new(userUpdatePayload)
	jsonUser, _ := json.Marshal(user)
    json.Unmarshal(jsonUser, userPayload)
	
	userPayloadWrapper := &userUpdatePayloadWrapper{
		User: userPayload,
	}
	return e.put(e, strconv.Itoa(user.Id), userPayloadWrapper)
}

func (e *usersEndpoint) DeleteUser(id int) error {
	return e.delete(e, strconv.Itoa(id))
}