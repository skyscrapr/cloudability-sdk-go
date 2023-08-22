package cloudability

import (
	"encoding/json"
	"strconv"
)

const containersEndpoint = "/containers/"

// ContainersEndpoint - Cloudability Containers Provisioning Endpoint
type ContainersEndpoint struct {
	*v3Endpoint
}

// Containers - Cloudability ContainersProvisioning Endpoint
func (c *Client) Containers() *ContainersEndpoint {
	return &ContainersEndpoint{newV3Endpoint(c, containersEndpoint)}
}

type Cluster struct {
	ID                int    `json:"id,omitempty"`
	ClusterName       string `json:"clusterName"`
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
	ClusterVersion    string `json:"clusterVersion,omitempty"`
	CreatedAt         string `json:"createdAt"`
}

type clustersResponse struct {
	Result []Cluster `json:"result"`
}

type clusterPayload struct {
	ClusterName       string `json:"clusterName"`
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
	ClusterVersion    string `json:"clusterVersion,omitempty"`
}

// GetCluster - Get an existing cluster by ID.
func (e ContainersEndpoint) GetCluster(id string) (*Cluster, error) {
	var result clustersResponse
	err := e.get(e, "provisioning/", &result)
	if err != nil {
		return nil, err
	}
	for _, cluster := range result.Result {
		if strconv.Itoa(cluster.ID) == id {
			return &cluster, nil
		}
	}
	return nil, err
}

// NewCluster - Create a new cluster.
func (e *ContainersEndpoint) NewCluster(clusterProvisioning *Cluster) (*Cluster, error) {
	clusterProvisioningPayload := new(clusterPayload)
	jsonCluster, err := json.Marshal(clusterProvisioning)
	err = json.Unmarshal(jsonCluster, clusterProvisioningPayload)
	var result v3Result[*Cluster]
	err = e.post(e, "provisioning/", clusterProvisioningPayload, &result)
	return result.Result, err
}

// UpdateCluster - Update an existing container by ID.
func (e *ContainersEndpoint) UpdateCluster(clusterProvisioning *Cluster) error {
	payload := new(clusterPayload)
	jsonCluster, _ := json.Marshal(clusterProvisioning)
	json.Unmarshal(jsonCluster, payload)
	return e.put(e, "provisioning/"+strconv.Itoa(clusterProvisioning.ID), payload)
}
