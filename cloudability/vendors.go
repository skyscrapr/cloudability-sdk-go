package cloudability


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

func (e vendorsEndpoint) Vendors() ([]Vendor, error) {
	var vendors []Vendor
	err := e.get("", &vendors)
	return vendors, err
}

func (e vendorsEndpoint) Credentials() ([]Credential, error) {
	var credentials []Credential
	err := e.get("", &credentials)
	return credentials, err
}

func (e vendorsEndpoint) Credential() (Credential, error) {
	var credential Credential
	err := e.get("", &credential)
	return credential, err
}