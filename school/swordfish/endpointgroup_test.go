// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var endpointGroupBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#EndpointGroup.EndpointGroup",
		"@odata.type": "#EndpointGroup.v1_1_2.EndpointGroup",
		"@odata.id": "/redfish/v1/EndpointGroup",
		"Id": "EndpointGroup-1",
		"Name": "EndpointGroupOne",
		"Description": "EndpointGroup One",
		"AccessState": "Optimized",
		"Endpoints": {
			"@odata.id": "/redfish/v1/Endpoints"
		},
		"GroupType": "Server",
		"Preferred": true,
		"TargetEndpointGroupIdentifier": 5
	}`)

// TestEndpointGroup tests the parsing of EndpointGroup objects.
func TestEndpointGroup(t *testing.T) {
	var result EndpointGroup
	err := json.NewDecoder(endpointGroupBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "EndpointGroup-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "EndpointGroupOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AccessState != OptimizedAccessState {
		t.Errorf("Access state is %s", result.AccessState)
	}

	if result.endpoints != "/redfish/v1/Endpoints" {
		t.Errorf("Invalid endpoints: %s", result.endpoints)
	}

	if result.GroupType != ServerGroupType {
		t.Errorf("Invalid group type: %s", result.GroupType)
	}

	if !result.Preferred {
		t.Error("Preferred should be true")
	}

	if result.TargetEndpointGroupIdentifier != 5 {
		t.Errorf("Invalid target endpoint group id: %d", result.TargetEndpointGroupIdentifier)
	}
}
