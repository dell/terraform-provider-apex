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
	"crypto/tls"
	apexClient "dell/apex-client"
	jmsClient "dell/apex-job-client"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GetNewToken retrieves a new token by running a Python script, reading from a named pipe,
// and sending an API request.
//
// ctx: The context.Context object for the function.
// Returns:
// - string: The token from the response from the API request.
// - error: An error if any occurred during the process.
func GetNewToken(ctx context.Context) (string, error) {

	// values taken from environmental variables
	// Set the directory path
	dirPath := os.Getenv("APEX_SAML_TOKEN_SCRIPT_DIR")
	// Set the Python script path
	scriptPath := os.Getenv("APEX_SAML_TOKEN_SCRIPT_NAME")

	// Set the command to run the Python script
	cmd := exec.Command("python", scriptPath)
	cmd.Dir = dirPath

	// Run the command and capture the output
	_, err := cmd.CombinedOutput()
	if err != nil {
		tflog.Error(ctx, "Error running SAML get token:", map[string]interface{}{
			"Error":   err,
			"DirPath": dirPath,
			"Script:": scriptPath,
		})
	}

	tokenFileName := os.Getenv("APEX_SAML_TOKEN_OUTPUT_FILE_NAME")
	//Open the named pipe for reading
	file, err := os.OpenFile(tokenFileName, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		tflog.Error(ctx, "Error reading from file:", map[string]interface{}{
			"Error": err,
			"file":  tokenFileName,
		})
		return "", err
	}
	defer file.Close() // Close the file when done

	// Read data from the pipe
	buffer := make([]byte, 1024*16)
	n, err := file.Read(buffer)
	if err != nil {
		tflog.Error(ctx, "Error reading from file:", map[string]interface{}{
			"Error": err,
		})
		return "", err
	}

	urlLink := os.Getenv("APEX_DI_TOKEN_URL")
	accesskey := os.Getenv("APEX_TOKEN_ACCESS_KEY")
	secret := os.Getenv("APEX_TOKEN_SECRET")
	authKey := accesskey + ":" + secret

	encodedValue := base64.StdEncoding.EncodeToString([]byte(authKey))
	method := "POST"
	header := map[string]string{
		"Authorization": "Basic" + " " + encodedValue,
		"Content-Type":  "application/x-www-form-urlencoded",
		"Cookie":        "lwp=c=us&r=us&l=en&cs=19&s=dhs",
	}

	data := url.Values{}
	data.Set("grant_type", "urn:ietf:params:oauth:grant-type:token-exchange")
	data.Set("subject_token", string(buffer[:n]))
	data.Set("subject_token_type", "urn:ietf:params:oauth:token-type:saml2")

	response, err := sendGetTokenRequest(urlLink, method, header, data.Encode())
	if err != nil {
		tflog.Error(ctx, "Error sending API request:", map[string]interface{}{
			"Error": err,
		})
		return "", err
	}
	return response, nil
}

// sendGetTokenRequest sends a GET request to the specified URL with the given method, header, data, and request body.
//
// Parameters:
// - urlstr: the URL to send the request to (string)
// - method: the HTTP method to use (string)
// - header: a map of headers to include in the request (map[string]string)
// - data: the data to include in the request body (interface{})
// - reqBody: the request body as a string (string)
//
// Returns:
// - responseBody: the response body as a string (string)
// - error: an error if the request fails (error)
func sendGetTokenRequest(urlstr string, method string, header map[string]string, reqBody string) (string, error) {
	req, err := http.NewRequest(method, urlstr, strings.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: true,
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return extractVariable(responseBody)
}

// AccessTokenResponse represents the structure of the JSON response
type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	IssuedTo     string `json:"issued_to"`
	TokenTypeURL string `json:"token_type_url"`
}

// extractVariable extracts the access token from the given response body.
//
// It takes a byte slice `responseBody` as input, which represents the response body containing the access token.
// It returns a string representing the access token and an error if any occurred during the unmarshaling process.
func extractVariable(responseBody []byte) (string, error) {
	var response AccessTokenResponse
	err := json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		return "", err
	}
	return response.AccessToken, nil
}

// UpdateToken updates the token for the given API clients.
//
// It takes a context.Context, a pointer to an apexClient.APIClient, and a pointer to a jmsClient.APIClient as parameters.
// It returns an error.
func UpdateToken(ctx context.Context, client *apexClient.APIClient, jmsClient *jmsClient.APIClient) error {
	token, err := GetNewToken(ctx)
	if err != nil {
		tflog.Error(ctx, "Error getting new token", map[string]interface{}{
			"Error": err,
		})
		return err
	}
	if token != "" {
		if client == nil {
			client.GetConfig().AddDefaultHeader("Authorization", "Bearer "+token)
		}
		if jmsClient == nil {
			jmsClient.GetConfig().AddDefaultHeader("Authorization", "Bearer "+token)
		}
	}
	return nil
}
