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

## PowerFlex Client
 
1. Change LogininResponseYaml (expires_in, refresh_expires_in) from string to int64


## Example openapi-generator-cli command (we are using 7.5.0 version of the generator)
```
openapi-generator-cli generate -g go -i apex-client-openapi-1.1.0.yaml --additional-properties=enumClassPrefix=true
```