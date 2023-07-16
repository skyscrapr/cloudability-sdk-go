package cloudability

import (
	"encoding/json"
	"strconv"
)

const usersEndpoint = "/users/"

// UsersEndpoint - Cloudability Users Endpoint
type UsersEndpoint struct {
	*v3Endpoint
}

// Users - Cloudability Users Endpoint
func (c *Client) Users() *UsersEndpoint {
	return &UsersEndpoint{newV3Endpoint(c, usersEndpoint)}
}

// User - Cloudability User
type User struct {
	ID                          int    `json:"id"`
	FrontdoorUserId             string `json:"frontdoor_user_id"`
	FrontdoorLogin              string `json:"frontdoor_login"`
	Email                       string `json:"email"`
	FullName                    string `json:"full_name"`
	DefaultDimensionFilterSetID int    `json:"default_dimension_filter_set_id"`
	SharedDimensionFilterSetIDs []int  `json:"shared_dimension_filter_set_ids"`
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

type userUpdatePayloadWrapper struct {
	User *userUpdatePayload `json:"user"`
}

type userUpdatePayload struct {
	FullName                           string `json:"full_name"`
	NewSharedDimensionFilterSetIDs     []int  `json:"new_shared_dimension_filter_set_ids"`
	UnshareExistingDimensionFilterSets bool   `json:"unshare_existing_dimension_filter_sets"`
	DefaultDimensionFilterSetID        int    `json:"default_dimension_filter_set_id"`
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
