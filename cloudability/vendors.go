package cloudability

import (
	"fmt"
)

const vendorsEndpoint = "/v3/vendors/"

// VendorsEndpoint - Cloudabiity Vendors Endpoint
type VendorsEndpoint struct {
	*v3Endpoint
}

// Vendors - Vendors Endpoint
func (c *Client) Vendors() *VendorsEndpoint {
	return &VendorsEndpoint{newV3Endpoint(c, vendorsEndpoint)}
}

// Vendor - Cloudability Vendor
type Vendor struct {
	Key string `json:"key"`
	Label string `json:"label"`
	Description string `json:"description"`
}

// Verification - Cloudability Verification
type Verification struct {
	State string `json:"state"`
	LastVerificationAttemptedAt string `json:"lastVerificationAttemptedAt"`
	Message string `json:"message"`
}

// Authorization - Cloudabiity Authorization
type Authorization struct {
	Type string `json:"type"`
	RoleName string `json:"roleName"`
	ExternalID string `json:"externalId"`
	BucketName *string `json:"bucketName,omitempty"`
	CostAndUsageReport *CostAndUsageReport `json:"costAndUsageReport,omitempty"`
}

// Account - Cloudability Account
type Account struct {
	ID string `json:"id"`
	VendorAccountName string `json:"vendorAccountName"`
	VendorAccountID string `json:"vendorAccountId"`
	VendorKey string `json:"vendorKey"`
	Verification *Verification `json:"verification"`
	Authorization *Authorization `json:"authorization"`
	ParentAccountID string `json:"parentAccountId"`
	CreatedAt string `json:"createdAt"`
}

// GetVendors - get all vendors
func (e VendorsEndpoint) GetVendors() ([]Vendor, error) {
	var vendors []Vendor
	err := e.get(e, "", &vendors)
	return vendors, err
}

// GetAccounts - get all accounts for a given vendor
func (e VendorsEndpoint) GetAccounts(vendor string) ([]Account, error) {
	var accounts []Account
	err := e.get(e, fmt.Sprintf("%s/accounts/", vendor), &accounts)
	return accounts, err
}

// GetAccount - get a single account for a given vendor
func (e VendorsEndpoint) GetAccount(vendor string, accountID string) (*Account, error) {
	var account Account
	err := e.get(e, fmt.Sprintf("%s/accounts/%s", vendor, accountID), &account)
	return &account, err
}

// VerifyAccount - verify account
func (e VendorsEndpoint) VerifyAccount(vendor string, accountID string) (*Account, error) {
	var account Account
	err := e.post(e, fmt.Sprintf("%s/accounts/%s/verification", vendor, accountID), nil, &account)
	return &account, err
}

// CostAndUsageReport - cost and usage report
type CostAndUsageReport struct {
	Name string `json:"name,omitempty"`
	Prefix string `json:"prefix,omitempty"`
}

// NewLinkedAccountParams - params required to create a new linked account
type NewLinkedAccountParams struct {
	VendorAccountID string `json:"vendorAccountId"`
	Type string `json:"type"`
}

// NewMasterAccountParams - params required to create a new master account
type NewMasterAccountParams struct {
	*NewLinkedAccountParams 
	BucketName string `json:"bucketName"`
	CostAndUsageReport *CostAndUsageReport `json:"costAndUsageReport"`
}

// NewMasterAccount - Create a new master account
func (e VendorsEndpoint) NewMasterAccount(vendorKey string, newAccountParams *NewMasterAccountParams) (*Account, error) {
	var account Account
	err := e.post(e, fmt.Sprintf("%s/accounts", vendorKey), newAccountParams, &account)
	return &account, err
}

// NewLinkedAccount - Create a new linked account
func (e VendorsEndpoint) NewLinkedAccount(vendorKey string, newAccountParams *NewLinkedAccountParams) (*Account, error) {
	var account Account
	err := e.post(e, fmt.Sprintf("%s/accounts", vendorKey), newAccountParams, &account)
	return &account, err
}

// DeleteAccount - Delete an account
func (e VendorsEndpoint) DeleteAccount(vendor string, accountID string) error {
	err := e.delete(e, fmt.Sprintf("%s/accounts/%s", vendor, accountID))
	return err
}
