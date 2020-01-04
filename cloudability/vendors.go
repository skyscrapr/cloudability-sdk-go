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

type Credential struct {
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

func (e vendorsEndpoint) GetCredentials(vendor string) ([]Credential, error) {
	var credentials []Credential
	err := e.get(fmt.Sprintf("%s/accounts/", vendor), &credentials)
	return credentials, err
}

func (e vendorsEndpoint) GetCredential(vendor string, id int) (*Credential, error) {
	var credential Credential
	err := e.get(fmt.Sprintf("%s/accounts/%s", vendor, strconv.Itoa(id)), &credential)
	return &credential, err
}

func (e vendorsEndpoint) VerifyAccountCredentials(vendor string, accountId string) error {
	err := e.post(fmt.Sprintf("%s/accounts/%s/verification", vendor, accountId), nil)
	return err
}

func (e vendorsEndpoint) NewCredential(vendor string, accountId string, credentialType string) (*Credential, error) {
	var credential Credential
	err := e.post(fmt.Sprintf("%s/accounts", vendor), &credential)
	return &credential, err
}

func (e vendorsEndpoint) DeleteCredential(vendor string, id string) error {
	err := e.delete(fmt.Sprintf("%s/accounts/%s", vendor, id))
	return err
}