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
	"net/http"
	"regexp"
	"testing"

	. "github.com/bytedance/mockey"
	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFileStorageOnPrem(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.UpdateToken).Return(nil).Build()
					FunctionMocker = Mock(helper.GetNewToken).Return("token", nil).Build()
				},
				Config: ProviderConfig + fileResourceOnPremConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccResourceFileStorageCloud(t *testing.T) {
	var resTerraformName = "apex_navigator_file_storage.cloud_instance"
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetNewToken).Return("token", nil).Build()
				},
				Config: ProviderConfig + fileResourceCloudConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			//Import testing
			{
				ResourceName:  resTerraformName,
				ImportStateId: "{\"system_id\":\"" + systemID + "\",\"target_username\":\"" + powerflexUser + "\",\"target_password\":\"" + powerflexPass + "\", \"target_host\":\"" + "" + "\", \"source_username\":\"" + powerflexUser + "\",\"source_password\":\"" + powerflexPass + "\",\"source_host\":\"" + "" + "\",\"insecure\":true}",
				ImportState:   true,
			},
			// Update file storage besides powerflex activation file should fail
			{
				Config:      ProviderConfig + fileResourceCloudUpdateErrorConfig,
				Check:       resource.ComposeAggregateTestCheckFunc(),
				ExpectError: regexp.MustCompile(`.*Error updating Apex Navigator File Storage*.`),
			},
			// Update powerflex activation file should be successful
			{
				Config: ProviderConfig + fileResourceCloudUpdatePowerScaleActivationConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("apex_navigator_file_storage.cloud_instance", "powerflex.username", "test-username"),
					resource.TestCheckResourceAttr("apex_navigator_file_storage.cloud_instance", "powerflex.password", "test-pass"),
				),
			},
			// Update back to original powerflex username and password
			{
				Config: ProviderConfig + fileResourceCloudConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccResourceFileStorageErrorCases(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.CreateStorageSystem).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + fileResourceCloudConfig,
				ExpectError: regexp.MustCompile(`.*Error creating Apex Navigator File Storage*.`),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GetStorageInstance).Return(nil, &http.Response{StatusCode: 400, Body: http.NoBody}, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + fileResourceCloudConfig,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Storage*.`),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GetStorageInstance).Return(nil, &http.Response{StatusCode: 404, Body: http.NoBody}, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + fileResourceCloudConfig,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Storage*.`),
			},
			// Step1: To test Read error, do a successful create then fail on update
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GetNewToken).Return("token", nil).Build()
				},
				Config: ProviderConfig + fileResourceCloudConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			// Step2: should show error on read
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetStorageInstance).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + fileResourceCloudConfig,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Storage*.`),
			},
		},
	})
}

var fileResourceOnPremConfig = `resource "apex_navigator_file_storage" "onprem_instance" {
	system_id                          = "ONEFS-0afff1f2676bb5f0e6658b17ab7b5139a03e"
	storage_system_type                = "POWERSCALE"
	bandwidth                          = "9999"
	capacity_impact                    = "9999"
	capacity_issue_count               = "9999"
	compression_savings                = "64.64"
	configuration_impact               = "9999"
	configuration_issue_count          = "9999"
	configured_size                    = "9999"
	connectivity_status                = "CONNECTED"
	license_type                       = "AWS-license-mock"         
	contract_coverage_type 			   = "covered"
	contract_expiration_date_timestamp = "0001-01-01 00:00:00 +0000 UTC"
	data_protection_impact             = "9999"
	data_protection_issue_count        = "9999"
	display_identifier                 = "0afff1f2676bb5f0e6658b17ab7b5139a03e"
	free_percent                       = "64.64"
	free_size                          = "9999"
	health_connectivity_status         = "NORMAL"
	health_issue_count                 = "9999"
	health_score                       = "99"
	health_state                       = "GOOD"
	iops                               = "9999"
	ipv4_address                       = "` + OnPremIP + `"
	ipv6_address                       = "` + OnPremIP + `"
	last_contact_timestamp             = "0001-01-01 00:00:00 +0000 UTC"
	latency                            = "9999"
	license_expiration_date_timestamp  = "0001-01-01 00:00:00 +0000 UTC"
	logical_size                       = "9999"
	model                              = "PowerScale software"
	name                               = "Create File OnPrem"
	overall_efficiency                 = "9999.0"
	performance_impact                 = "9999"
	performance_issue_count            = "9999"
	serial_number                      = "serial-9999"
	site_name                          = "Round Rock"
	snaps_savings                      = "64.64"
	system_health_impact               = "9999"
	system_health_issue_count          = "9999"
	thin_savings                       = "9999.00"
	total_size                         = "64"
	used_percent                       = "64.64"
	used_size                          = "64"
	unconfigured_size                  = "1"
	vendor                             = "Dell"
	product_version                    = "1.0"
	deployment_details = {
	  system_on_prem = {
		deployment_type  = "ONPREM"
		site_name        = "SiteName"
		location         = "Location"
		country          = "Country"
		state            = "State"
		city             = "City"
		street_address_1 = "StreetAddress1"
		street_address_2 = "StreetAddress2"
		zip_code         = "ZipCode"
	  }
	}
	powerflex {
		username = "` + powerflexUser + `"
		password = "` + powerflexPass + `"
		scheme   = "` + powerflexScheme + `"
	}
  }

  output "instance_file_storage_onprem" {
	sensitive = true
	value = apex_navigator_file_storage.onprem_instance
  }`

var fileResourceCloudConfig = `resource "apex_navigator_file_storage" "cloud_instance" {
	storage_system_type                = "POWERSCALE"
	name                               = "Create File Cloud"
	product_version                    = "1.0"
	deployment_details = {
	  system_public_cloud = {
		deployment_type            = "PUBLIC_CLOUD"
		cloud_type                 = "AWS"
		cloud_account              = "CloudAccount"
		cloud_region               = "CloudRegion"
		availability_zone_topology = "MULTIPLE_AVAILABILITY_ZONE"
		raw_capacity 			   = "20"
		tier_type                  = "TierType"
		ssh_key_name               = "name"
		iam_instance_profile 	   = "CreateRole"
		vpc = {
		  is_new_vpc               = true
		  vpc_name                 = "my-storage-vpc"
		}
		subnet_options = [
		  {
		  subnet_id   = "id-1"
		  subnet_type = "EXTERNAL"
		},
		{
		  # subnet_id   = "id-2"
		  subnet_type = "INTERNAL"
		},
		{
			# subnet_id   = "id-2"
			subnet_type = "SCG"
		  }
		]
	  }
	}
	powerflex {
		username = "` + powerflexUser + `"
		password = "` + powerflexPass + `"
		scheme   = "` + powerflexScheme + `"
	}
  }

  output "instance_file_storage_cloud" {
	sensitive = true
	value = apex_navigator_file_storage.cloud_instance
  }`

var fileResourceCloudUpdateErrorConfig = `resource "apex_navigator_file_storage" "cloud_instance" {
	storage_system_type                = "POWERSCALE"
	name                               = "Create File Cloud Update"
	product_version                    = "1.0"
	deployment_details = {
	  system_public_cloud = {
		deployment_type            = "PUBLIC_CLOUD"
		cloud_type                 = "AWS"
		cloud_account              = "CloudAccount"
		cloud_region               = "CloudRegion"
		availability_zone_topology = "MULTIPLE_AVAILABILITY_ZONE"
		raw_capacity 			   = "20"
		tier_type                  = "TierType"
		ssh_key_name               = "name"
		iam_instance_profile 	   = "CreateRole"
		vpc = {
		  is_new_vpc               = true
		  vpc_name                 = "my-storage-vpc"
		}
		subnet_options = [
		  {
		  subnet_id   = "id-1"
		  subnet_type = "EXTERNAL"
		},
		{
		  # subnet_id   = "id-2"
		  subnet_type = "INTERNAL"
		},
		{
			# subnet_id   = "id-2"
			subnet_type = "SCG"
		  }
		]
	  }
	}
	powerflex {
		username = "` + powerflexUser + `"
		password = "` + powerflexPass + `"
		scheme   = "` + powerflexScheme + `"
	}
  }

  output "instance_file_storage_cloud" {
	sensitive = true
	value = apex_navigator_file_storage.cloud_instance
  }`

var fileResourceCloudUpdatePowerScaleActivationConfig = `resource "apex_navigator_file_storage" "cloud_instance" {
	storage_system_type                = "POWERSCALE"
	name                               = "Create File Cloud"
	product_version                    = "1.0"
	deployment_details = {
	  system_public_cloud = {
		deployment_type            = "PUBLIC_CLOUD"
		cloud_type                 = "AWS"
		cloud_account              = "CloudAccount"
		cloud_region               = "CloudRegion"
		availability_zone_topology = "MULTIPLE_AVAILABILITY_ZONE"
		raw_capacity 			   = "20"
		tier_type                  = "TierType"
		ssh_key_name               = "name"
		iam_instance_profile 	   = "CreateRole"
		vpc = {
		  is_new_vpc               = true
		  vpc_name                 = "my-storage-vpc"
		}
		subnet_options = [
		  {
		  subnet_id   = "id-1"
		  subnet_type = "EXTERNAL"
		},
		{
		  # subnet_id   = "id-2"
		  subnet_type = "INTERNAL"
		},
		{
			# subnet_id   = "id-2"
			subnet_type = "SCG"
		  }
		]
	  }
	}
	powerflex {
		username = "test-username"
		password = "test-pass"
		scheme   = "` + powerflexScheme + `"
	}
  }

  output "instance_file_storage_cloud" {
	sensitive = true
	value = apex_navigator_file_storage.cloud_instance
  }`
