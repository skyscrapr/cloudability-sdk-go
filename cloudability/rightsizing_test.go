package cloudability

import (
	"testing"
	"fmt"
	"net/url"
	"net/http"
	"net/http/httptest"
)

func TestNewRightsizingEndpoint(t *testing.T) {
	testClient := NewClient("testapikey")
	e := testClient.Rightsizing()
	if e.BaseURL.String() != apiV3URL {
		t.Errorf("RightsizingEndpoint BaseURL mismatch. Got %s. Expected %s", e.BaseURL.String(), apiV3URL)
	}
	if e.EndpointPath != rightsizingEndpoint {
		t.Errorf("RightsizingEndpoint EndpointPath mismatch. Got %s. Expected %s", e.EndpointPath, rightsizingEndpoint)
	}
}

func TestGetResource(t *testing.T) {
	vendor := "aws"
	service := "ec2"
	resourceIdentifier := "i-abcde12345"
	
	resourceJSON := []byte(`{
		"result": 
			{
				"service": "ec2",
				"resourceIdentifier": "i-abcde12345",
				"recommendations": [
					{
						"action": "Rightsize",
						"nodeType": "x1e.xlarge",
						"risk": 0
					}
				]
			}
		}`)

//       "name": "My EC2 Instance",
//       "vendorAccountId": "458273444843",
//       "tagsMappings: [
//         {
//           "tagName": "tag_user_Environment",
//           "vendorTagValue": "production"
//         },
//         {
//           "tagName": "tag_user_Name",
//           "vendorTagValue": "My EC2 Instance"
//         }
//       ],
//       "availabilityZone": "ap-southeast-2b",
//       "provider": "NATIVE"
//       "region": "ap-southeast-2",
//       "os": "Linux",
//       "nodeType": "i3.8xlarge",
//       "unitPrice": 1.81,
//       "totalSpend": 434.11,
//       "idle": 0,
//       "localCapacity": 7600,
//       "localDrives": 4,
//       "cpuCapacity": 32,
//       "memoryCapacity": 244,
//       "networkCapacity": 10000,
//       "lastSeen": "2019-10-31T23:00:00Z",
//       "tenancy" : "default",
//       "hoursRunning": 240,
//       "cpuMax": 6,
//       "memoryMax": 6,
//       "recommendations": [
//         {
//           "preferenceOrder": 1,
//           "defaultsOrder": 1,
//           "localCapacity": 120,
//           "localDrives": 1,
//           "cpuCapacity": 4,
//           "memoryCapacity": 122,
//           "previousGenTarget": false,
//           "currentGen": true,
//           "sameMemory": false,
//           "sameFamily": false,
//           "unitPrice": 0.83,
//           "cpuRatio": 0.13,
//           "memoryRatio": 0.5,
//           "diskXPutCapacity": 100,
//           "networkRatio": 1,
//           "cpuRisk": 0,
//           "memoryRisk": 0,
//           "diskRisk": 0,
//           "networkRisk": 0,
//           "risk": 0,
//           "savingsPct": 54,
//           "savings": 233.95,
//           "inDefaults": true,
//           "memoryFit": false,
//           "persistentStorageAdded": false
//         }
//       ],
//       "defaultSameFamily": false,
//       "defaultCurrentGen": true,
//       "defaultMemoryFit": false
//     }
//   ]`)

	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		method := "GET"
		path := fmt.Sprintf("/v3/rightsizing/%s/recommendations/%s", vendor, service)
		filters := fmt.Sprintf("resourceIdentifier==%s", resourceIdentifier)

		rw.WriteHeader(http.StatusOK)

		if req.Method != method {
			t.Errorf("Expected request method ‘%s’, got ‘%s’", method, req.Method)
		}
		if req.URL.EscapedPath() != path {
			t.Errorf("Expected request path ‘%s’, got ‘%s’", path, req.URL.EscapedPath())
		}
		if req.URL.Query().Get("filters") != filters {
			t.Errorf("Expected query parameter filters=‘%s’, got ‘%s’", filters, req.URL.Query().Get("filters"))
		}
		// TODO: Fix this
		// if body != nil {
		// 	jsonReq, err := simplejson.NewFromReader(req.Body)
    	// 	if err != nil {
      	// 		t.Errorf("Error while reading request JSON: %s", err)
    	// 	}
		// 	if !reflect.DeepEqual(jsonReq, req.Body) {
		// 		t.Errorf("Expected body ‘%s’, got ‘%s’", body, req.Body)
		// 	}
		// }
		rw.Write([]byte(resourceJSON))
	}))
	defer testServer.Close()
	testClient := NewClient("testapikey")
	e := testClient.Rightsizing()
	e.BaseURL, _ = url.Parse(testServer.URL)
	resource, err := e.GetResource(vendor, service, resourceIdentifier)
	if err != nil{
		t.Fail()
	}
	if resource.Service != service {
		t.Errorf("Expected resource service ‘%s’, got ‘%s’", service, resource.Service)
	}
	if resource.ResourceIdentifier != resourceIdentifier {
		t.Errorf("Expected resource resourceIdentifier ‘%s’, got ‘%s’", resourceIdentifier, resource.ResourceIdentifier)
	}
	if resource.Recommendations == nil {
		t.Errorf("Expected resource recommendations not nil")
	}
	// if resource.Recommendations[0].Action != action {
	// 	t.Errorf("Expected recommendation action ‘%s’, got ‘%s’", action, resource.Recommendations[0].Action)
	// }
}

