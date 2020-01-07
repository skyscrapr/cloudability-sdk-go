package cloudability

import (
	"fmt"
	"strconv"
)


type vendorsEndpoint struct {
	*cloudabilityV3Endpoint
}

func newVendorsEndpoint(apikey string) *vendorsEndpoint {
	e := &vendorsEndpoint{newCloudabilityV3Endpoint(apikey)}
	e.EndpointPath = "/v3/vendors/"
	return e
}

type Vendor struct {
	Key string `json:"key"`
	Label string `json:"label"`
	Description string `json:"description"`
}

type Verification struct {
	State string `json:"state"`
	LastVerificationAttemptedAt string `json:"lastVerificationAttemptedAt"`
	Message string `json:"message"`
}

type Authorization struct {
	Type string `json:"type"`
	RoleName string `json:"roleName"`
	ExternalId string `json:"externalId"`
}

type Account struct {
	Id string `json:"id"`
	VendorAccountName string `json:"vendorAccountName"`
	VendorAccountId string `json:"vendorAccountId"`
	VendorKey string `json:"vendorKey"`
	Verification Verification `json:"verification"`
	Authorization Authorization `json:"authorization"`
	ParentAccountId string `json:"parentAccountId"`
	CreatedAt string `json:"createdAt"`
}

func (e vendorsEndpoint) GetVendors() ([]Vendor, error) {
	var vendors []Vendor
	err := e.get("", &vendors)
	return vendors, err
}

func (e vendorsEndpoint) GetAccounts(vendor string) ([]Account, error) {
	var accounts []Account
	err := e.get(fmt.Sprintf("%s/accounts/", vendor), &accounts)
	return accounts, err
}

func (e vendorsEndpoint) GetAccount(vendor string, id int) (*Account, error) {
	var account Account
	err := e.get(fmt.Sprintf("%s/accounts/%s", vendor, strconv.Itoa(id)), &account)
	return &account, err
}

func (e vendorsEndpoint) VerifyAccount(vendor string, accountId string) error {
	err := e.post(fmt.Sprintf("%s/accounts/%s/verification", vendor, accountId), nil, nil)
	return err
}

type newCredentialParams struct {
	VendorAccountId string `json:"vendorAccountId"`
	Type string `json:"type"`
}

func (e vendorsEndpoint) NewAccount(vendorKey string, accountId string, credType string) (*Account, error) {
	var account Account
	body := &newCredentialParams{
		VendorAccountId: accountId,
		Type: credType,
	}
	err := e.post(fmt.Sprintf("%s/accounts", vendorKey), body, &account)
	return &account, err
}

func (e vendorsEndpoint) DeleteAccount(vendor string, id string) error {
	err := e.delete(fmt.Sprintf("%s/accounts/%s", vendor, id))
	return err
}