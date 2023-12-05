package cloudability

import (
	"net/url"
	"testing"
)

func TestContainersProvisioningEndpoint(t *testing.T) {
	testClient := NewClient("testapikey")
	e := testClient.Containers()
	if e.BaseURL.String() != apiV3URL {
		t.Errorf("ContaintersProvisioningEndpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), apiV3URL)
	}
	if e.EndpointPath != containersEndpoint {
		t.Errorf("ContaintersProvisioningEndpoint EndpointPath mismatch. Got %s. Expected %s", e.EndpointPath, containersEndpoint)
	}
}

func TestGetCluster(t *testing.T) {
	testServer := testAPI(t, "GET", "/containers/provisioning", nil)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.Containers()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetCluster("1")
	if err != nil {
		t.Fail()
	}
}

func TestGetClusterConfig(t *testing.T) {
	mockYAML := `
apiVersion: v1
kind: Namespace
metadata:
  name: cloudability
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: cloudability
  namespace: cloudability
`

	testServer := testAPI(t, "GET", "/containers/provisioning/1/config", mockYAML)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.Containers()
	e.BaseURL, _ = url.Parse(testServer.URL)
	config, err := e.GetClusterConfig("1")
	if err != nil {
		t.Fail()
	}
	if config != mockYAML {
		t.Fail()
	}
}

func TestNewCluster(t *testing.T) {
	cluster := &Cluster{
		ClusterName:       "test-cluster-name",
		ClusterVersion:    "test-cluster-version",
		KubernetesVersion: "1.11",
	}
	testServer := testAPI(t, "POST", "/containers/provisioning", cluster)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.Containers()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.NewCluster(cluster)
	if err != nil {
		t.Fail()
	}
}

func TestUpdateCluster(t *testing.T) {
	cluster := &Cluster{
		ClusterName:       "test-cluster-name",
		ClusterVersion:    "test-cluster-version",
		KubernetesVersion: "test-kubernetes-version",
	}
	testServer := testAPI(t, "PUT", "/containers/provisioning/0", cluster)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.Containers()
	e.BaseURL, _ = url.Parse(testServer.URL)
	err := e.UpdateCluster(cluster)
	if err != nil {
		t.Fail()
	}
}
