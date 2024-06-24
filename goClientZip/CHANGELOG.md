# Copyright (c) 2023 Dell Inc., or its subsidiaries. All Rights Reserved.
#
# Licensed under the Mozilla Public License Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://mozilla.org/MPL/2.0/
#
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

## Apex Jobs Client 
Changes in apex-jobs-client-openapi-1.0.0.yaml : 
Line 328 changed type of Body(ResponseContent -> Body) to string
Updated example for Body(ResponseContent -> Body) : Line 318,402

apex-jobs-client-openapi-1.0.0 : .\model_response_content.go line 26 Changed type of Body from map[string]interface{} to *string

## Apex Client 1.0

Remove - refresh_timestamp clonesModel as a required paramter from the clones model (since APEX does not return in every object and this can make the datasource fail). 
Add - POWERSCALE to the StorageProductEnum (This is temperary until we have the new client which includes the PowerScale APIs)

## Apex Client 1.1

Add - ISILON and POWERSCALE to the StorageSystemTypeEnum
Remove -oneOf: from TargetSystemOptions and update indentation
Remove -oneOf: from UpdateMobilityGroupInput and update indentation
Remove - ID as a Required field from StorageProductInstance since it is not always returned in the REST Request
Remove - ID as a Required field from SupportMap since it is not always returned in the REST Request

## Apex Client 1.1.1

Add - ISILON to the StorageSystemTypeEnum
Remove `x-internal: true` from CreateTargetBaseInput
Remove - ID as a Required field from StorageProductsInstance since it is not always returned in the REST Request
Remove - ID as a Required field from SupportMap since it is not always returned in the REST Request
Remove - refresh_timestamp clonesModel as a required paramter from the clones model (since APEX does not return in every object and this can make the datasource fail). 

Changes in generated code:

model_power_flex_tier_enum.go 
	line 23 BALANCED -> POWERFLEXTIERENUM_BALANCED
	line 24 PERFORMANCE_OPTIMIZED-> POWERFLEXTIERENUM_PERFORMANCE_OPTIMIZED

model_storage_product_enum.go
	line 23 POWERFLEX -> STORAGEPRODUCTENUM_POWERFLEX  
	line 24 POWERSCALE ->STORAGEPRODUCTENUM_POWERSCALE

model_power_scale_tier_enum.go 
 	line 23 BALANCED -> POWERSCALETIERENUM_BALANCED
	line 24 PERFORMANCE_OPTIMIZED-> POWERSCALETIERENUM_PERFORMANCE_OPTIMIZED

model_storage_type_enum.go
	line 23 BLOCK -> STORAGETYPEENUM_BLOCK
	line 24 FILE -> STORAGETYPEENUM_FILE

model_storage_system_type_enum.go
	line 23 POWERFLEX-> STORAGESYSTEMTYPEENUM_POWERFLEX
	line 24 POWERSCALE -> STORAGESYSTEMTYPEENUM_POWERSCALE
	line 25 ISILON -> STORAGESYSTEMTYPEENUM_ISILON

model_tier_enum.go
	line 23 BALANCED -> TIERENUM_BALANCED
	line 24 PERFORMANCE_OPTIMIZED-> TIERENUM_PERFORMANCE_OPTIMIZED
	line 25 COST_OPTIMIZED -> TIERENUM_COST_OPTIMIZED

model_subnet_type_enum.go
	line 23 UNDEFINED -> SUBNETTYPEENUM_UNDEFINED
	line 24 EXTERNAL -> SUBNETTYPEENUM_EXTERNAL
	line 25 INTERNAL -> SUBNETTYPEENUM_INTERNAL
	line 26 SCG -> SUBNETTYPEENUM_SCG

model_subnet_options.go
	line 35, line 45 UNDEFINED -> SUBNETTYPEENUM_UNDEFINED

model_availability_zone_topology_enum.go
	line 23 SINGLE_AVAILABILITY_ZONE -> AVAILABILITYZONETOPOLOGYENUM_SINGLE_AVAILABILITY_ZONE
	line 24 MULTIPLE_AVAILABILITY_ZONE -> AVAILABILITYZONETOPOLOGYENUM_MULTIPLE_AVAILABILITY_ZONE

## PowerFlex Client
 
1. Change LogininResponseYaml (expires_in, refresh_expires_in) from string to int64


## Example openapi-generator-cli command (we are using 7.5.0 version of the generator)
```
openapi-generator-cli generate -g go -i apex-client-openapi-1.1.0.yaml --additional-properties=enumClassPrefix=true
```