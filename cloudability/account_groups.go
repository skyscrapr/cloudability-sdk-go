package cloudability

import (
	"encoding/json"
	"strconv"
)

const accountGroupsEndpoint = "/account_groups/"

// AccountGroupsEndpoint - Cloudability Account Groups Endpoint
type AccountGroupsEndpoint struct {
	*v1Endpoint
}

// AccountGroups - Cloudability AccountGroups Endpoint
func (c *Client) AccountGroups() *AccountGroupsEndpoint {
	return &AccountGroupsEndpoint{newV1Endpoint(c, accountGroupsEndpoint)}
}

// AccountGroup - Cloudability AccountGroup
type AccountGroup struct {
	ID                      int      `json:"id,omitempty"`
	Name                    string   `json:"name"`
	Position                int      `json:"position"`
	AccountGroupEntryValues []string `json:"account_group_entry_values,omitempty"`
}

// GetAccountGroups - Get all account groups
func (e *AccountGroupsEndpoint) GetAccountGroups() ([]AccountGroup, error) {
	var accountGroups []AccountGroup
	err := e.get(e, "", &accountGroups)
	return accountGroups, err
}

// GetAccountGroup - Get account group
func (e *AccountGroupsEndpoint) GetAccountGroup(id int) (*AccountGroup, error) {
	var accountGroup AccountGroup
	err := e.get(e, strconv.Itoa(id), &accountGroup)
	return &accountGroup, err
}

type accountGroupNewPayload struct {
	Name     string `json:"name"`
	Position int    `json:"position"`
}

// NewAccountGroup - Create an account group
func (e *AccountGroupsEndpoint) NewAccountGroup(accountGroup *AccountGroup) error {
	accountGroupPayload := new(accountGroupNewPayload)
	jsonAccountGroup, _ := json.Marshal(accountGroup)
	json.Unmarshal(jsonAccountGroup, accountGroupPayload)
	return e.post(e, "", accountGroupPayload, nil)
}

type accountGroupUpdatePayload struct {
	Name     string `json:"name"`
	Position int    `json:"position"`
}

// UpdateAccountGroup - Update an account group
func (e *AccountGroupsEndpoint) UpdateAccountGroup(accountGroup *AccountGroup) error {
	accountGroupPayload := new(accountGroupUpdatePayload)
	jsonAccountGroup, _ := json.Marshal(accountGroup)
	json.Unmarshal(jsonAccountGroup, accountGroupPayload)
	return e.put(e, strconv.Itoa(accountGroup.ID), accountGroupPayload)
}

// DeleteAccountGroup - Delete an account group
func (e *AccountGroupsEndpoint) DeleteAccountGroup(id int) error {
	return e.delete(e, strconv.Itoa(id))
}
