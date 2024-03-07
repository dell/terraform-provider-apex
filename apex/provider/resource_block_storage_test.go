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

func TestAccBlockStorageResourceOnPrem(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + blockResourceOnPremConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccBlockStorageResourceCloud(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: ProviderConfig + blockResourceCloudConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			{
				Config:      ProviderConfig + blockResourceCloudUpdateConfig,
				Check:       resource.ComposeAggregateTestCheckFunc(),
				ExpectError: regexp.MustCompile(`.*Unable to update Apex Navigator block storage*.`),
			},
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.GetVolumesCollection).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + volumeCollectionConfig,
				ExpectError: regexp.MustCompile(`.*Unable to Read Apex Navigator Volumes*.`),
			},
		},
	})
}

func TestAccBlockStorageResourceErrorCases(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					FunctionMocker = Mock(helper.CreateBlockStorage).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + blockResourceCloudConfig,
				ExpectError: regexp.MustCompile(`.*Error creating Block Storage*.`),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GetBlockStorageInstance).Return(nil, &http.Response{StatusCode: 400, Body: http.NoBody}, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + blockResourceCloudConfig,
				ExpectError: regexp.MustCompile(`.*Error retrieving created Block storage*.`),
			},
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GetBlockStorageInstance).Return(nil, &http.Response{StatusCode: 404, Body: http.NoBody}, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + blockResourceCloudConfig,
				ExpectError: regexp.MustCompile(`.*Error retrieving created Block storage*.`),
			},
			// Step1: To test Read error, do a successful create then fail on update
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
				},
				Config: ProviderConfig + blockResourceCloudConfig,
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			// Step2: should show error on read
			{
				PreConfig: func() {
					if FunctionMocker != nil {
						FunctionMocker.UnPatch()
					}
					FunctionMocker = Mock(helper.GetBlockStorageInstance).Return(nil, nil, fmt.Errorf("Mock error")).Build()
				},
				Config:      ProviderConfig + blockResourceCloudConfig,
				ExpectError: regexp.MustCompile(`.*Error Reading Apex Navigator block storage*.`),
			},
		},
	})
}

var blockResourceOnPremConfig = `resource "apex_navigator_block_storage" "onprem_instance" {
	system_id                          = "POWERFLEX-ELMSIO0523STQ3-Mock2"
	system_type                        = "POWERFLEX"
	bandwidth                          = "9999"
	capacity_impact                    = "9999"
	capacity_issue_count               = "9999"
	compression_savings                = "64.64"
	configuration_impact               = "9999"
	configuration_issue_count          = "9999"
	configured_size                    = "9999"
	connectivity_status                = "CONNECTED"
	license_type                       = "AWS-license-mock"
	contract_coverage_type             = "covered"
	contract_expiration_date_timestamp = "0001-01-01 00:00:00 +0000 UTC"
	data_protection_impact             = "9999"
	data_protection_issue_count        = "9999"
	display_identifier                 = "ELMSIO0523STQ3"
	free_percent                       = "64.64"
	free_size                          = "9999"
	health_connectivity_status         = "NORMAL"
	health_issue_count                 = "9999"
	health_score                       = "99"
	health_state                       = "GOOD"
	iops                               = "9999"
	ipv4_address                       = "10.164.99.99"
	ipv6_address                       = "10.164.99.99"
	last_contact_timestamp             = "0001-01-01 00:00:00 +0000 UTC"
	latency                            = "9999"
	license_expiration_date_timestamp  = "0001-01-01 00:00:00 +0000 UTC"
	logical_size                       = "9999"
	model                              = "PowerFlex software"
	name                               = "Create OnPrem"
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
  }
  
  output "instance_block_storage_onprem" {
	value = apex_navigator_block_storage.onprem_instance
  }`

var blockResourceCloudConfig = `
resource "apex_navigator_block_storage" "cloud_instance" {
	system_type                        = "POWERFLEX"
	name                               = "Create Cloud"
	product_version                      = "1.0"
	deployment_details = {
	  system_public_cloud = {
		deployment_type            = "PUBLIC_CLOUD"
		cloud_type                 = "AWS"
		cloud_account              = "CloudAccount"
		cloud_region               = "CloudRegion"
		availability_zone_topology = "MULTIPLE_AVAILABILITY_ZONE"
		minimum_iops               = "1"
		minimum_capacity           = "2"
		tier_type                  = "TierType"
		ssh_key_name               = "name"
		vpc = {
		  is_new_vpc               = true
		  vpc_name                 = "my-storage-vpc"
		}
		subnet_options = [
		  {
		  subnet_id   = "id-1"
		  # cidr_block  = "10.10.10/21"
		  subnet_type = "EXTERNAL"
		},
		{
		  # subnet_id   = "id-2"
		  cidr_block  = "10.10.10/22"
		  subnet_type = "INTERNAL"
		}
		]
	  }
	}
  }
  
  output "instance_block_storage_cloud" {
	value = apex_navigator_block_storage.cloud_instance
  }`

var blockResourceCloudUpdateConfig = `
  resource "apex_navigator_block_storage" "cloud_instance" {
	  system_type                        = "POWERFLEX"
	  name                               = "Create Cloud Update"
	  product_version                      = "1.0"
	  deployment_details = {
		system_public_cloud = {
		  deployment_type            = "PUBLIC_CLOUD"
		  cloud_type                 = "AWS"
		  cloud_account              = "CloudAccount"
		  cloud_region               = "CloudRegion"
		  availability_zone_topology = "MULTIPLE_AVAILABILITY_ZONE"
		  minimum_iops               = "1"
		  minimum_capacity           = "2"
		  tier_type                  = "TierType"
		  ssh_key_name               = "name"
		  vpc = {
			is_new_vpc               = true
			vpc_name                 = "my-storage-vpc"
		  }
		  subnet_options = [
			{
			subnet_id   = "id-1"
			# cidr_block  = "10.10.10/21"
			subnet_type = "EXTERNAL"
		  },
		  {
			# subnet_id   = "id-2"
			cidr_block  = "10.10.10/22"
			subnet_type = "INTERNAL"
		  }
		  ]
		}
	  }
	}
	
	output "instance_block_storage_cloud" {
	  value = apex_navigator_block_storage.cloud_instance
	}`
