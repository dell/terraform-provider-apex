/*
Copyright (c) 2024 Dell Inc., or its subsidiaries. All Rights Reserved.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://mozilla.org/MPL/2.0/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package provider

import (
	"fmt"
	"log"
	"os"
	"testing"

	. "github.com/bytedance/mockey"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/joho/godotenv"
)

var ProviderConfig = ""

// Used for Mocking responses from functions
var FunctionMocker *Mocker
var cloneHost = setDefault(os.Getenv("CLONE_HOST"), "test_clone_host")
var cloneID = setDefault(os.Getenv("CLONE_ID"), "test_clone_id")
var cloneRefreshID = setDefault(os.Getenv("CLONE_REFRESH_ID"), "test_clone_refresh_id")
var cloneUnmapID = setDefault(os.Getenv("CLONE_UNMAP_ID"), "test_clone_unmap_id")
var cloneUnmapHost = setDefault(os.Getenv("CLONE_UNMAP_HOST"), "test_clone_unmap_host")
var mobilityName = setDefault(os.Getenv("MOBILITY_NAME"), "test_mobility_name")
var systemID = setDefault(os.Getenv("SYSTEM_ID"), "test_system_id")
var volumeID = setDefault(os.Getenv("VOLUME_ID"), "test_volume_id")
var sourceMobilityTargetGroupID = setDefault(os.Getenv("SOURCE_MOBILITY_TARGET_GROUP_ID"), "test_source_group_id")
var mobilityTargetOptions = setDefault(os.Getenv("MOBILITY_TARGET_OPTIONS"), "test_target_options")
var mobilityTargetSystemID = setDefault(os.Getenv("MOBILITY_TARGET_SYSTEM_ID"), "test_system_id")
var powerflexUser = setDefault(os.Getenv("POWERFLEX_USER"), "test_user")
var powerflexPass = setDefault(os.Getenv("POWERFLEX_PASS"), "test_pass")
var powerflexTargetUser = setDefault(os.Getenv("POWERFLEX_TARGET_USER"), "test_target_user")
var powerflexTargetPass = setDefault(os.Getenv("POWERFLEX_TARGET_PASS"), "test_target_pass")
var powerflexScheme = setDefault(os.Getenv("POWERFLEX_TARGET_PASS"), "http")
var OnPremIP = setDefault(os.Getenv("ON_PREM_IP"), "localhost:3005")
var mobilitySourceID = setDefault(os.Getenv("MOBILITY_SOURCE_ID"), "test_mobility_source_id")
var mobilityTargetID = setDefault(os.Getenv("MOBILITY_TARGET_ID"), "test_mobility_target_id")

// Datasource variables
var blockStorageID1 = setDefault(os.Getenv("BLOCK_STORAGE_ID_1"), "POWERFLEX-ELMSIO0523STQ3-Mock")
var blockStorageID2 = setDefault(os.Getenv("BLOCK_STORAGE_ID_2"), "POWERFLEX-ELMSIO0523STQ3-Mock2")
var cloneID1 = setDefault(os.Getenv("CLONE_ID_1"), "clone_read_Id")
var cloneID2 = setDefault(os.Getenv("CLONE_ID_2"), "clone_read_Id_2")
var hostID1 = setDefault(os.Getenv("HOST_ID1"), "test_host_id1")
var hostID2 = setDefault(os.Getenv("HOST_ID2"), "test_host_id2")
var mobilityID1 = setDefault(os.Getenv("MOBILITY_ID_1"), "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__fcdecfaf-c61e-4b4d-8f89-65c6ef00d0000")
var mobilityID2 = setDefault(os.Getenv("MOBILITY_ID_2"), "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__fcdecfaf-c61e-4b4d-8f89-65c6ef00d0001")
var mobilityTargetID1 = setDefault(os.Getenv("MOBILITY_TARGET_ID_1"), "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5")
var mobilityTargetID2 = setDefault(os.Getenv("MOBILITY_TARGET_ID_2"), "POWERFLEX-ELMSIO0823QVTV__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd1")
var sourcePoolsID1 = setDefault(os.Getenv("SOURCE_POOL_ID1"), "Pool_ID1")
var sourcePoolsID2 = setDefault(os.Getenv("SOURCE_POOL_ID2"), "Pool_ID2")
var volumeID1 = setDefault(os.Getenv("VOLUME_ID1"), "volume_id1")
var volumeID2 = setDefault(os.Getenv("VOLUME_ID2"), "volume_id2")
var storageProduct1 = setDefault(os.Getenv("STORAGE_PRODUCT_1"), "POWERFLEX")

var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"apex": providerserver.NewProtocol6WithError(New("test")()),
	}
)

func init() {
	err := godotenv.Load("apex.env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
		return
	}

	host := os.Getenv("APEX_HOST")
	token := setDefault(os.Getenv("APEX_TOKEN"), "test_token")

	ProviderConfig = fmt.Sprintf(` 
		provider "apex" {
			host      = "%s"
			token     = "%s"
		}
	`, host, token)
}

func testAccPreCheck(t *testing.T) {
	// Check that the required environment variables are set.
	if os.Getenv("APEX_HOST") == "" {
		t.Fatal("APEX_HOST environment variable not set")
	}

	// Make sure to unpatch before each new test is run
	if FunctionMocker != nil {
		FunctionMocker.UnPatch()
	}
}

func setDefault(osInput string, defaultStr string) string {
	if osInput == "" {
		return defaultStr
	}
	return osInput
}
