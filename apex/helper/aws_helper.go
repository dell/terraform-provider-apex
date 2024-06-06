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

package helper

import (
	"context"
	"encoding/json"
	"net/http"

	client "dell/apex-client"

	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// GenerateAwsTrustPolicy aws trust policy
func GenerateAwsTrustPolicy(clientAPI *client.APIClient, accountID string) (*client.AwsTrustPolicy, *http.Response, error) {
	return clientAPI.AwsAccountsAPI.AwsAccountsGenerateTrustPolicy(context.Background()).AwsTrustPolicyInput(client.AwsTrustPolicyInput{AwsAccountId: accountID}).Execute()
}

// ConnectAccount connects an aws account
func ConnectAccount(clientAPI *client.APIClient, arn string) (*client.RedactedAwsAccountInstance, *http.Response, error) {
	return clientAPI.AwsAccountsAPI.AwsAccountsCreate(context.Background()).AwsAccountsCreateInput(client.AwsAccountsCreateInput{RoleArn: arn}).Execute()
}

// GetAccountInstance gets an aws account
func GetAccountInstance(clientAPI *client.APIClient, accountID string) (*client.AwsAccountsInstance, *http.Response, error) {
	return clientAPI.AwsAccountsAPI.AwsAccountsInstance(context.Background(), accountID).Execute()
}

// UpdateAccount updates an aws account
func UpdateAccount(clientAPI *client.APIClient, accountID string, roleArn string) (*client.Job, *http.Response, error) {
	return clientAPI.AwsAccountsAPI.AwsAccountsModify(context.Background(), accountID).AwsAccountModifyInput(client.AwsAccountModifyInput{RoleArn: roleArn}).Execute()
}

// DisconnectAccount disconnects an aws account
func DisconnectAccount(clientAPI *client.APIClient, accountID string) (*client.Job, *http.Response, error) {
	return clientAPI.AwsAccountsAPI.AwsAccountsDelete(context.Background(), accountID).Execute()
}

// GetAwsAccountCollection gets list of block storage instances
func GetAwsAccountCollection(clientAPI *client.APIClient, filter []basetypes.StringValue) (*client.AwsAccountsCollection200Response, *http.Response, error) {
	accounts, http, err := clientAPI.AwsAccountsAPI.AwsAccountsCollection(context.Background()).Execute()
	// If there is an error or the filter is empty just return the full collection and response error
	if err != nil || len(filter) == 0 {
		return accounts, http, err
	}

	// Otherwise filter the collection (this API does not have the filter query param so we need to do it manually)
	response := client.AwsAccountsCollection200Response{
		Results: make([]client.RedactedAwsAccountInstance, 0),
	}
	for _, value := range filter {
		for _, account := range accounts.Results {
			if *account.Id == *value.ValueStringPointer() {
				response.Results = append(response.Results, account)
			}
		}
	}
	return &response, http, nil
}

// MapAwsAccount maps the AWS account
func MapAwsAccount(account client.RedactedAwsAccountInstance) models.AwsAccountModel {

	return models.AwsAccountModel{
		ID:     types.StringPointerValue(account.Id),
		Alias:  types.StringPointerValue(account.AwsAccountAlias),
		Status: types.StringPointerValue(account.Status),
	}
}

// GetAwsPermssionCollection gets list of block storage instances
func GetAwsPermssionCollection(clientAPI *client.APIClient, filter []basetypes.StringValue) (*client.AwsPermissionPoliciesCollection200Response, *http.Response, error) {
	permissions, http, err := clientAPI.AwsPermissionPoliciesAPI.AwsPermissionPoliciesCollection(context.Background()).Execute()
	// If there is an error or the filter is empty just return the full collection and response error
	if err != nil || len(filter) == 0 {
		return permissions, http, err
	}

	// Otherwise filter the collection (this API does not have the filter query param so we need to do it manually)
	response := client.AwsPermissionPoliciesCollection200Response{
		Results: make([]client.AwsPermissionPoliciesInstance, 0),
	}
	for _, value := range filter {
		for _, permission := range permissions.Results {
			if *permission.Id == *value.ValueStringPointer() {
				response.Results = append(response.Results, permission)
			}
		}
	}
	return &response, http, nil
}

// MapAwsPermission maps the AWS permission policy
func MapAwsPermission(permission client.AwsPermissionPoliciesInstance) models.AwsPermissionsModel {
	return models.AwsPermissionsModel{
		ID:                types.StringPointerValue(permission.Id),
		Version:           types.StringPointerValue(permission.Version),
		PermissionsPolicy: types.StringValue(marshalPermissionPolicy(permission.PermissionPolicy)),
	}
}

// marshalPermissionPolicy creates the marshal string of the permission policy
func marshalPermissionPolicy(permissionPolicy map[string]interface{}) string {
	b, _ := json.Marshal(permissionPolicy)
	return string(b)
}

// MapGeneratedStatements maps the AWS permssions statements to terraform state
func MapGeneratedStatements(statements []client.AwsTrustPolicyStatementInner) basetypes.ListValue {

	var mapStatements []attr.Value

	principal := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"aws": types.StringType,
		},
	}
	condtion := types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"string_equals": types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"sts_external_id": types.StringType,
				},
			},
			"bool": types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"aws_multi_factor_auth_present": types.StringType,
				},
			},
		},
	}
	attrTypes := map[string]attr.Type{
		"effect":    types.StringType,
		"action":    types.StringType,
		"principal": principal,
		"condition": condtion,
	}
	for _, state := range statements {
		aws := ""
		stsExternalID := ""
		awsMultiFactor := ""
		if state.Principal != nil && state.Principal.AWS != nil {
			aws = *state.Principal.AWS
		}
		if state.Condition != nil {
			if state.Condition.StringEquals != nil && state.Condition.StringEquals.StsExternalId != nil {
				stsExternalID = *state.Condition.StringEquals.StsExternalId
			}
			if state.Condition.Bool != nil && state.Condition.Bool.AwsMultiFactorAuthPresent != nil {
				awsMultiFactor = *state.Condition.Bool.AwsMultiFactorAuthPresent
			}
		}
		object, _ := types.ObjectValue(attrTypes, map[string]attr.Value{
			"effect": types.StringPointerValue(state.Effect),
			"action": types.StringPointerValue(state.Action),
			"principal": types.ObjectValueMust(map[string]attr.Type{
				"aws": types.StringType,
			}, map[string]attr.Value{
				"aws": types.StringValue(aws),
			}),
			"condition": types.ObjectValueMust(map[string]attr.Type{
				"string_equals": types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"sts_external_id": types.StringType,
					},
				},
				"bool": types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"aws_multi_factor_auth_present": types.StringType,
					},
				},
			}, map[string]attr.Value{
				"string_equals": types.ObjectValueMust(map[string]attr.Type{
					"sts_external_id": types.StringType,
				}, map[string]attr.Value{
					"sts_external_id": types.StringValue(stsExternalID),
				}),
				"bool": types.ObjectValueMust(map[string]attr.Type{
					"aws_multi_factor_auth_present": types.StringType,
				}, map[string]attr.Value{
					"aws_multi_factor_auth_present": types.StringValue(awsMultiFactor),
				}),
			}),
		})
		mapStatements = append(mapStatements, object)
	}
	typeObject := types.ObjectType{
		AttrTypes: attrTypes,
	}
	return basetypes.NewListValueMust(typeObject, mapStatements)
}
