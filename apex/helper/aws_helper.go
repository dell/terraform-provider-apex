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
	"net/http"

	client "dell/apex-client"

	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

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
		ID:      types.StringPointerValue(permission.Id),
		Version: types.StringPointerValue(permission.Version),
		PermissionsPolicy: models.PermissionsPolicy{
			Version:   types.StringPointerValue(permission.PermissionPolicy.Version),
			Statement: mapStatements(permission.PermissionPolicy.Statement),
		},
	}
}

// mapStatements maps the AWS permssions statements
func mapStatements(statements []client.AwsPermissionPoliciesInstancePermissionPolicyStatementInner) []models.Statement {
	response := make([]models.Statement, 0)
	for _, state := range statements {
		actions := make([]types.String, 0)
		iams := make([]types.String, 0)

		for _, action := range state.Action {
			actions = append(actions, types.StringValue(action))
		}

		if state.Condition != nil && state.Condition.StringLike != nil && state.Condition.StringLike.IamAWSServiceName != nil {
			for _, iam := range state.Condition.StringLike.IamAWSServiceName {
				iams = append(iams, types.StringValue(iam))
			}
		}

		response = append(response, models.Statement{
			Sid:               types.StringPointerValue(state.Sid),
			Effect:            types.StringPointerValue(state.Effect),
			Action:            actions,
			Resource:          types.StringPointerValue(state.Resource),
			IamAwsServiceName: iams,
		})
	}
	return response
}
