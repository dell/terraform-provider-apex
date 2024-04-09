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
	powerscaleClient "dell/powerscale-client"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/dell/terraform-provider-apex/apex/client"
	"github.com/dell/terraform-provider-apex/apex/models"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AsyncCheckActivation checks to see if the current platform (ie.powerflex or powerscale) is activated, if not it will activate it in a go routine.
func AsyncCheckActivation(ctx context.Context, stop chan bool, err chan error, clientAPI *apexClient.APIClient, systemID string, activateSystemClientSystem models.ActivationAsyncClientModel, storageType apexClient.StorageProductEnum) {
	tflog.Info(ctx, fmt.Sprintf("Aysnc check activation: %d", stop))
	for {
		select {
		// If stop is sent end loop
		case <-stop:
			err <- nil
			return
		// Keep checking activation ever 30 seconds
		default:
			clientSystem := models.ActivationClientModel{
				Host:     types.StringValue(activateSystemClientSystem.Host.ValueString()),
				Username: types.StringValue(activateSystemClientSystem.Username.ValueString()),
				Password: types.StringValue(activateSystemClientSystem.Password.ValueString()),
				Scheme:   types.StringValue(activateSystemClientSystem.Scheme.ValueString()),
				Insecure: types.BoolValue(activateSystemClientSystem.Insecure.ValueBool()),
			}
			activeErr := ActivateSystemClientSystem(ctx, clientAPI, systemID, clientSystem, storageType)
			time.Sleep(time.Duration(activateSystemClientSystem.PollInterval.ValueInt64()) * time.Second)
			if activeErr != nil {
				tflog.Debug(ctx, fmt.Sprintf("Aysnc check activation error!!!: %d", err))
				// return error stop routine
				err <- activeErr
				return
			}

		}
	}
}

// ActivateSystemClientSystem checks to see if the current platform (ie.powerflex or powerscale) is activated, if not it will activate it.
func ActivateSystemClientSystem(ctx context.Context, clientAPI *apexClient.APIClient, systemID string, activateSystemClientSystem models.ActivationClientModel, storageType apexClient.StorageProductEnum) error {

	// Check if the ip was set otherwise grab it from apex
	if activateSystemClientSystem.Host.ValueString() == "" {
		system, _, err := clientAPI.StorageSystemsAPI.StorageSystemsInstance(ctx, systemID).Execute()
		if err != nil {
			return err
		}
		if system.Ipv4Address == nil || *system.Ipv4Address == "" {
			return fmt.Errorf("no ip or hostname found for system %s, please set the `host` field manually in the client config", systemID)
		}
		activateSystemClientSystem.Host = types.StringPointerValue(system.Ipv4Address)
	}

	systemActivationToken := getApexActivationToken(ctx, clientAPI, systemID, apexClient.StorageProductEnum(storageType))
	tflog.Debug(ctx, "Previous System Activation Token: "+systemActivationToken)
	// There is no token related to the systemId so we need to create one and relate it
	if systemActivationToken == "" {
		tokenNew, clientTokenError := getClientActivationToken(ctx, activateSystemClientSystem, storageType)
		if clientTokenError != nil {
			tflog.Debug(ctx, "Error getting new activation token "+clientTokenError.Error())
			return clientTokenError
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

		// If token is no longer valid, or will expire in less the 5 minutes attempt to activate a new one
		fiveMinutesAgo := time.Now().Add(-5 * time.Minute)
		if !active.IsTokenValid || active.ExpirationTimestamp.Compare(fiveMinutesAgo) < 0 {
			token, clientTokenError := getClientActivationToken(ctx, activateSystemClientSystem, apexClient.StorageProductEnum(active.SystemType))
			if clientTokenError != nil {
				tflog.Debug(ctx, "Error getting new activation token "+clientTokenError.Error())
				return clientTokenError
			}

			// Update apex with the newly created activation token
			patch := clientAPI.StorageSystemTokensAPI.StorageSystemTokensModify(ctx, systemActivationToken)
			patch = patch.StorageSystemTokensModifyRequest(apexClient.StorageSystemTokensModifyRequest{
				AccessToken: token,
			})
			_, err = patch.Execute()
			if err != nil {
				tflog.Debug(ctx, "Error patching activation token")
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
	if status != nil && status.StatusCode == http.StatusConflict {
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

// getClientActivationToken decifers which client we need to do an activation against and gets the activation token
func getClientActivationToken(ctx context.Context, activateSystemClientSystem models.ActivationClientModel, storageType apexClient.StorageProductEnum) (string, error) {
	// If token is from a powerflex system create a client and get a new token.
	switch storageType {
	case apexClient.STORAGEPRODUCTENUM_POWERFLEX:
		tokenNew, activateTokenErr := getPowerFlexActivationToken(ctx, activateSystemClientSystem)
		if activateTokenErr != nil {
			tflog.Debug(ctx, "Error getting new patch token instance "+activateTokenErr.Error())
			return "", activateTokenErr
		}
		return tokenNew, nil
	case apexClient.STORAGEPRODUCTENUM_POWERSCALE:
		tokenNew, activateTokenErr := getPowerScaleActivationToken(ctx, activateSystemClientSystem)
		if activateTokenErr != nil {
			tflog.Debug(ctx, "Error getting new patch token instance "+activateTokenErr.Error())
			return "", activateTokenErr
		}
		return tokenNew, nil
	default:
		return "", fmt.Errorf("system type %s not supported", storageType)
	}
}

func getPowerScaleActivationToken(ctx context.Context, activateSystemClientSystem models.ActivationClientModel) (string, error) {
	// Create a client
	createdClient, pClientErr := client.CreatePowerScaleClient(ctx, activateSystemClientSystem.Host.ValueString(), activateSystemClientSystem.Scheme.ValueString(), activateSystemClientSystem.Insecure.ValueBool())
	if pClientErr != nil {
		return "", pClientErr
	}

	resp, loginErr := createdClient.DefaultApi.PostRestAuthLogin(ctx, &powerscaleClient.PostRestAuthLoginOpts{
		LoginCredentialsYaml: optional.NewInterface(powerscaleClient.LoginCredentialsYaml{
			Username: activateSystemClientSystem.Username.ValueString(),
			Password: activateSystemClientSystem.Password.ValueString(),
			// The services which we require access too
			Services: []string{"platform", "namespace"},
			// 15 minute timeout
			TimeoutAbsolute: 900,
			TimeoutInactive: 900,
		}),
	})
	if loginErr != nil {
		return "", loginErr
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "isisessid" {
			return cookie.Value, nil
		}
	}
	return "", fmt.Errorf("unable to authenticate with powerscale")
}

// getPowerFlexActivationToken gets a new powerflex activation token
func getPowerFlexActivationToken(ctx context.Context, activateSystemClientSystem models.ActivationClientModel) (string, error) {
	// Create a client
	createdClient, pClientErr := client.CreatePowerFlexClient(ctx, activateSystemClientSystem.Host.ValueString(), activateSystemClientSystem.Scheme.ValueString(), activateSystemClientSystem.Insecure.ValueBool())
	if pClientErr != nil {
		return "", pClientErr
	}

	// Attempt a login
	resp, _, loginErr := createdClient.DefaultApi.PostRestAuthLogin(ctx, &powerflexClient.PostRestAuthLoginOpts{
		LoginCredentialsYaml: optional.NewInterface(powerflexClient.LoginCredentialsYaml{
			Username: activateSystemClientSystem.Username.ValueString(),
			Password: activateSystemClientSystem.Password.ValueString(),
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

// SetPowerflexAysncClientState sets the powerflex client state
func SetPowerflexAysncClientState(source models.ActivationAsyncClientModel) *models.ActivationAsyncClientModel {
	target := &models.ActivationAsyncClientModel{
		Host:         handleNullState(source.Host),
		Insecure:     source.Insecure,
		Username:     handleNullState(source.Username),
		Password:     handleNullState(source.Password),
		Scheme:       handleNullState(source.Scheme),
		PollInterval: types.Int64Value(source.PollInterval.ValueInt64()),
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
