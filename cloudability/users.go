package cloudability

import (
	"strconv"
	"encoding/json"
)

const usersEndpoint = "users/"

// UsersEndpoint - Cloudability Users Endpoint
type UsersEndpoint struct {
	*v1Endpoint
}

// Users - Cloudability Users Endpoint
func (c *Client) Users() *UsersEndpoint {
	return &UsersEndpoint{newV1Endpoint(c, usersEndpoint)}
}

// User - Cloudability User
type User struct {
	ID int `json:"id,omitempty"`
	Email string `json:"email"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
	Restricted bool `json:"restricted"`
	SharedDimensionFilterSetIDs []int `json:"shared_dimension_filter_set_ids"`
	DefaultDimensionFilterID int `json:"default_dimension_filer_set_id"`
}

// GetUsers - Get all users
func (e *UsersEndpoint) GetUsers() ([]User, error) {
	var users []User
	err := e.get(e, "", &users)
	return users, err
}

// GetUser - Get user
func (e *UsersEndpoint) GetUser(id int) (*User, error) {
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
	SharedDimensionFilterSetIDs []int `json:"shared_dimension_filter_set_ids"`
	DefaultDimensionFilterID int `json:"default_dimension_filer_set_id"`
}

// NewUser - Create a user
func (e *UsersEndpoint) NewUser(user *User) error {
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
	SharedDimensionFilterSetIDs []int `json:"shared_dimension_filter_set_ids"`
	DefaultDimensionFilterID int `json:"default_dimension_filer_set_id"`
}

// UpdateUser - Update a user
func (e *UsersEndpoint) UpdateUser(user *User) error {
	userPayload := new(userUpdatePayload)
	jsonUser, _ := json.Marshal(user)
    json.Unmarshal(jsonUser, userPayload)
	
	userPayloadWrapper := &userUpdatePayloadWrapper{
		User: userPayload,
	}
	return e.put(e, strconv.Itoa(user.ID), userPayloadWrapper)
}

// DeleteUser - Delete a user
func (e *UsersEndpoint) DeleteUser(id int) error {
	return e.delete(e, strconv.Itoa(id))
}