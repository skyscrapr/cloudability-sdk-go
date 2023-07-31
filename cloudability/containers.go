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
}

type clusterPayload struct {
	ClusterName       string `json:"clusterName"`
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
	ClusterVersion    string `json:"clusterVersion,omitempty"`
}

// GetContainer - Get an existing cluster by ID.
func (e ContainersEndpoint) GetContainer(id int) (*Cluster, error) {
	var result v3Result[*Cluster]
	err := e.get(e, "provisioning/"+strconv.Itoa(id), &result)
	return result.Result, err
}

// NewContainers - Create a new Container Provisioning.
func (e *ContainersEndpoint) NewContainers(clusterProvisioning *Cluster) (*Cluster, error) {
	clusterProvisioningPayload := new(clusterPayload)
	jsonClusterProvisioning, _ := json.Marshal(clusterProvisioning)
	json.Unmarshal(jsonClusterProvisioning, clusterProvisioningPayload)
	var result v3Result[*Cluster]
	err := e.post(e, "provisioning/", clusterProvisioningPayload, nil)
	return result.Result, err
}

// UpdateContainers - Update an existing container by ID.
func (e *ContainersEndpoint) UpdateContainers(clusterProvisioning *Cluster) error {
	payload := new(clusterPayload)
	jsonContainersProvisioning, _ := json.Marshal(clusterProvisioning)
	json.Unmarshal(jsonContainersProvisioning, payload)
	return e.put(e, "provisioning/"+strconv.Itoa(clusterProvisioning.ID), payload)
}
