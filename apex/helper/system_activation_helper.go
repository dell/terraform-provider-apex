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
	apexClient "dell/apex-client"
	powerflexClient "dell/powerflex-client"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/antihax/optional"
	"github.com/dell/terraform-provider-apex/apex/client"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ActivateSystemPowerflexSystem checks to see if the current powerflex is activated, if not it will activate it.
func ActivateSystemPowerflexSystem(ctx context.Context, clientAPI *apexClient.APIClient, systemID string, powerflexClientModel models.ActivationClientModel, storageType apexClient.StorageProductEnum) error {

	// Check if the ip was set otherwise grab it from apex
	if powerflexClientModel.Host.ValueString() == "" {
		system, _, err := clientAPI.StorageSystemsAPI.StorageSystemsInstance(ctx, systemID).Execute()
		if err != nil {
			return err
		}
		if system.Ipv4Address == nil || *system.Ipv4Address == "" {
			return fmt.Errorf("no ip address found for system %s, please set manually in the powerflex config", systemID)
		}
		powerflexClientModel.Host = types.StringPointerValue(system.Ipv4Address)
	}

	systemActivationToken := getApexActivationToken(ctx, clientAPI, systemID, apexClient.StorageProductEnum(storageType))
	tflog.Debug(ctx, "Previous System Activation Token: "+systemActivationToken)
	// There is no token related to the systemId so we need to create one and relate it
	if systemActivationToken == "" {
		tokenNew, activateTokenErr := getPowerFlexActivationToken(ctx, powerflexClientModel)
		if activateTokenErr != nil {
			tflog.Debug(ctx, "Error getting new activation token "+tokenNew)
			return activateTokenErr
		}
		post := clientAPI.StorageSystemTokensAPI.StorageSystemTokensCreate(ctx)
		post = post.StorageSystemTokensCreateRequest(apexClient.StorageSystemTokensCreateRequest{
			AccessToken: tokenNew,
			SystemType:  apexClient.StorageProductEnum(storageType),
			SystemId:    systemID,
		})
		_, _, err := post.Execute()
		if err != nil {
			tflog.Debug(ctx, "Error activating new token "+tokenNew)
			return err
		}

		// There is already a token related to the systemId
		// Check if it is still valid, if so return else activate a new token and update the relationship on apex
	} else {
		active, _, err := clientAPI.StorageSystemTokensAPI.StorageSystemTokensInstance(ctx, systemActivationToken).Execute()
		if err != nil {
			tflog.Debug(ctx, "Error checking token instance "+err.Error())
			return err
		}

		// If token is no longer valid, attempt to activate a new one
		if !active.IsTokenValid {
			token := ""
			// If token is from a powerflex system create a client and get a new token.
			if active.SystemType == "POWERFLEX" {
				tokenNew, activateTokenErr := getPowerFlexActivationToken(ctx, powerflexClientModel)
				if activateTokenErr != nil {
					tflog.Debug(ctx, "Error getting new patch token instance "+activateTokenErr.Error())
					return activateTokenErr
				}
				token = tokenNew
			} else {
				return fmt.Errorf("system type %s not supported", active.SystemType)
			}

			// Update apex with the newly created activation token
			patch := clientAPI.StorageSystemTokensAPI.StorageSystemTokensModify(ctx, systemActivationToken)
			patch = patch.StorageSystemTokensModifyRequest(apexClient.StorageSystemTokensModifyRequest{
				AccessToken: token,
			})
			_, err = patch.Execute()
			if err != nil {
				tflog.Info(ctx, "Error getting new patching")
				return err
			}
		}
	}

	return nil
}

// getApexActivationToken does a post with a fake token to the Apex activate token API
// if the API returns a 409 there is an existing token, we can extract the token from the error response and use it to validate a new token
// if the API returns anything else the system does not have an activation token currently.
// @return string current activation token. If it is an empty string we were not able to extract and user should try to post a new one
func getApexActivationToken(ctx context.Context, clientAPI *apexClient.APIClient, systemID string, storageType apexClient.StorageProductEnum) string {
	post := clientAPI.StorageSystemTokensAPI.StorageSystemTokensCreate(ctx)
	post = post.StorageSystemTokensCreateRequest(apexClient.StorageSystemTokensCreateRequest{
		AccessToken: "test-dummy-token",
		SystemType:  storageType,
		SystemId:    systemID,
	})
	_, status, err := post.Execute()

	// 409 attempt to extract the old token from the error response
	if status.StatusCode == http.StatusConflict {
		if err == nil {
			return ""
		}
		errByte, ok := err.(*apexClient.GenericOpenAPIError)
		if ok {
			return extractTokenFromError(ctx, errByte.Body())
		}
	}
	return ""
}

// extractTokenFromError Extracts the token from the status body
func extractTokenFromError(ctx context.Context, body []byte) string {
	var parsedData apexClient.ErrorResponse
	err := json.Unmarshal(body, &parsedData)
	tflog.Debug(ctx, "Extracting from body Error Response: "+string(body))
	if err != nil {
		return ""
	}
	// If the token is in the arguments it will be a stringyfied "array" look like below
	// "[6848e5a6-47c6-4ec1-a548-14fbbe41c4a3, system-id]"
	// We want to extract the token from that string
	if len(parsedData.Messages) > 0 && len(parsedData.Messages[0].Arguments) > 0 && parsedData.Messages[0].Arguments[0] != "[]" {
		arg := parsedData.Messages[0].Arguments[0]
		argSplit := strings.Split(arg, ",")
		// Split by `,` and get the first element
		if len(argSplit) < 1 {
			return ""
		}
		arg = argSplit[0]
		// Remove the leading `[`
		arg = arg[1:]
		return arg
	}

	return ""
}

// getPowerFlexActivationToken gets a new powerflex activation token
func getPowerFlexActivationToken(ctx context.Context, powerflexClientModel models.ActivationClientModel) (string, error) {
	// Create a client
	createdClient, pClientErr := client.CreatePowerFlexClient(ctx, powerflexClientModel.Host.ValueString(), powerflexClientModel.Scheme.ValueString(), powerflexClientModel.Insecure.ValueBool())
	if pClientErr != nil {
		return "", pClientErr
	}

	// Attempt a login
	resp, _, loginErr := createdClient.DefaultApi.PostRestAuthLogin(ctx, &powerflexClient.PostRestAuthLoginOpts{
		LoginCredentialsYaml: optional.NewInterface(powerflexClient.LoginCredentialsYaml{
			Username: powerflexClientModel.Username.ValueString(),
			Password: powerflexClientModel.Password.ValueString(),
		}),
	})
	if loginErr != nil {
		return "", loginErr
	}

	// Return the access token
	return resp.AccessToken, nil
}

// SetPowerflexClientState sets the powerflex client state
func SetPowerflexClientState(source models.ActivationClientModel) *models.ActivationClientModel {
	target := &models.ActivationClientModel{
		Host:     handleNullState(source.Host),
		Insecure: source.Insecure,
		Username: handleNullState(source.Username),
		Password: handleNullState(source.Password),
		Scheme:   handleNullState(source.Scheme),
	}
	return target
}

// handleNullState handles the null or unknown state
func handleNullState(value basetypes.StringValue) basetypes.StringValue {
	if value.IsNull() || value.IsUnknown() || value.ValueString() == "" {
		return types.StringValue("")
	}
	return value
}
