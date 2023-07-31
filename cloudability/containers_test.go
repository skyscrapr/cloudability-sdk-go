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
	testServer := testAPI(t, "GET", "/containers/provisioning/1", nil)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.Containers()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.GetContainer(1)
	if err != nil {
		t.Fail()
	}
}

func TestNewContainers(t *testing.T) {
	cluster := &Cluster{
		ClusterName:       "test-cluster-name",
		ClusterVersion:    "test-cluster-version",
		KubernetesVersion: "test-kubernetes-version",
	}
	testServer := testAPI(t, "POST", "/containers/provisioning", cluster)
	defer testServer.Close()

	testClient := NewClient("testapikey")
	e := testClient.Containers()
	e.BaseURL, _ = url.Parse(testServer.URL)
	_, err := e.NewContainers(cluster)
	if err != nil {
		t.Fail()
	}
}

func TestUpdateContainers(t *testing.T) {
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
	err := e.UpdateContainers(cluster)
	if err != nil {
		t.Fail()
	}
}
