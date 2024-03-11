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

const ()

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

var mobilitySourceID = setDefault(os.Getenv("MOBILITY_SOURCE_ID"), "test_mobility_source_id")
var mobilityTargetID = setDefault(os.Getenv("MOBILITY_TARGET_ID"), "test_mobility_target_id")
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

	ProviderConfig = fmt.Sprintf(` 
		provider "apex" {
			host      = "%s"
		}
	`, host)
}

func testAccPreCheck(t *testing.T) {
	// Check that the required environment variables are set.
	if os.Getenv("APEX_HOST") == "" {
		t.Fatal("APEX_HOST environment variable not set")
	}

	t.Log(ProviderConfig)

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
