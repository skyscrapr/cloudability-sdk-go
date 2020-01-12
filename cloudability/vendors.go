package cloudability

import (
	"fmt"
)

const vendors_endpoint = "/v3/vendors/"

type vendorsEndpoint struct {
	*v3Endpoint
}

func (c *Client) Vendors() *vendorsEndpoint {
	return &vendorsEndpoint{newV3Endpoint(c, vendors_endpoint)}
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
	BucketName *string `json:"bucketName,omitempty"`
	CostAndUsageReport *CostAndUsageReport `json:"costAndUsageReport,omitempty"`
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
	err := e.get(e, "", &vendors)
	return vendors, err
}

func (e vendorsEndpoint) GetAccounts(vendor string) ([]Account, error) {
	var accounts []Account
	err := e.get(e, fmt.Sprintf("%s/accounts/", vendor), &accounts)
	return accounts, err
}

func (e vendorsEndpoint) GetAccount(vendor string, accountId string) (*Account, error) {
	var account Account
	err := e.get(e, fmt.Sprintf("%s/accounts/%s", vendor, accountId), &account)
	return &account, err
}

func (e vendorsEndpoint) VerifyAccount(vendor string, accountId string) (*Account, error) {
	var account Account
	err := e.post(e, fmt.Sprintf("%s/accounts/%s/verification", vendor, accountId), nil, &account)
	return &account, err
}

type CostAndUsageReport struct {
	Name *string `json:"name,omitempty"`
	Prefix *string `json:"prefix,omitempty"`
}

type NewLinkedAccountParams struct {
	VendorAccountId string `json:"vendorAccountId"`
	Type string `json:"type"` 
}

type NewMasterAccountParams struct {
	*NewLinkedAccountParams 
	BucketName string `json:"bucketName"`
	CostAndUsageReport *CostAndUsageReport `json:"costAndUsageReport"`
}

func (e vendorsEndpoint) NewMasterAccount(vendorKey string, newAccountParams *NewMasterAccountParams) (*Account, error) {
	var account Account
	err := e.post(e, fmt.Sprintf("%s/accounts", vendorKey), newAccountParams, &account)
	return &account, err
}

func (e vendorsEndpoint) NewLinkedAccount(vendorKey string, newAccountParams *NewLinkedAccountParams) (*Account, error) {
	var account Account
	err := e.post(e, fmt.Sprintf("%s/accounts", vendorKey), newAccountParams, &account)
	return &account, err
}

func (e vendorsEndpoint) DeleteAccount(vendor string, accountId string) error {
	err := e.delete(e, fmt.Sprintf("%s/accounts/%s", vendor, accountId))
	return err
}