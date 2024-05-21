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
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const expectedBufferSize = 1024 * 16

var config *Config
var once sync.Once

// GetNewToken retrieves a new token by running a Python script, reading from a named pipe,
// and sending an API request.
//
// ctx: The context.Context object for the function.
// Returns:
// - string: The token from the response from the API request.
// - error: An error if any occurred during the process.
func GetNewToken(ctx context.Context) (string, error) {

	// values taken from environmental variables
	once.Do(func() {
		config = &Config{
			ScriptPath:    filepath.Clean(os.Getenv("APEX_SAML_TOKEN_SCRIPT_NAME")),
			DirPath:       filepath.Clean(os.Getenv("APEX_SAML_TOKEN_SCRIPT_DIR")),
			TokenFileName: filepath.Clean(os.Getenv("APEX_SAML_TOKEN_OUTPUT_FILE_NAME")),
			URLLink:       os.Getenv("APEX_DI_TOKEN_URL"),
			AccessKey:     os.Getenv("APEX_TOKEN_ACCESS_KEY"),
			Secret:        os.Getenv("APEX_TOKEN_SECRET"),
		}
	})

	err := validateConfig(config)
	if err != nil {
		return "", err
	}
	// Set the command to run the Python script
	// Run the command with a timeout of 3 minutes
	err1 := runCommandWithTimeout(ctx, 3*time.Minute, "python", config.DirPath, config.ScriptPath)
	if err1 != nil {
		tflog.Error(ctx, "Error running SAML get token:", map[string]interface{}{
			"Error":   err,
			"DirPath": config.DirPath,
			"Script:": config.ScriptPath,
		})
		return "", fmt.Errorf("failed to run script for SAML get token: %w", err1)
	}

	// Read data from the pipe
	buffer, n, shouldReturn1, err := readFromPipe(ctx)
	if shouldReturn1 {
		return "", err
	}
	//get header and data
	header, data, shouldReturn, err := getHeaderData(ctx, n, buffer)
	if shouldReturn {
		return "", err
	}
	method := "POST"
	response, err := sendGetTokenRequest(config.URLLink, method, header, data.Encode())

	if err != nil {
		tflog.Error(ctx, "Error sending API request:", map[string]interface{}{
			"Error": err,
		})
		return "", err
	}
	if response == "" {
		tflog.Error(ctx, "Error in getting token")
		return "", errors.New("error in getting token")
	}
	return response, nil
}

// readFromPipe reads data from a named pipe and returns the read data as a byte slice,
// the number of bytes read, a boolean indicating whether an error occurred, and an error object.
//
// The function takes a context.Context object as a parameter.
// It opens the named pipe specified by config.TokenFileName for reading using os.OpenFile.
// If an error occurs during opening, it logs the error and returns nil, 0, true, and an error object.
// Otherwise, it reads data from the pipe into a buffer of size EXPECTED_BUFFER_SIZE using file.Read.
// If an error occurs during reading, it logs the error and returns nil, 0, true, and an error object.
// Otherwise, it returns the buffer, the number of bytes read, false, and nil.
func readFromPipe(ctx context.Context) ([]byte, int, bool, error) {
	//Open the named pipe for reading
	file, err := os.OpenFile(config.TokenFileName, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		tflog.Error(ctx, "Error reading from file:", map[string]interface{}{
			"Error": err,
			"file":  config.TokenFileName,
		})
		return nil, 0, true, fmt.Errorf("failed to open file: %w", err)
	}
	// Close the file when done
	defer file.Close()

	// Read data from the pipe
	buffer := make([]byte, expectedBufferSize)
	n, err := file.Read(buffer)
	if err != nil {
		tflog.Error(ctx, "Error reading from file:", map[string]interface{}{
			"Error": err,
		})
		return nil, 0, true, fmt.Errorf("failed to read from file: %w", err)
	}

	return buffer[:n], n, false, nil
}

// getHeaderData generates the header and data for making a request.
//
// It takes the following parameters:
// - ctx: the context.Context object for the request
// - n: the length of the buffer containing the SAML token
// - buffer: the byte buffer containing the SAML token
//
// It returns the following:
// - header: a map[string]string containing the headers for the request
// - data: a url.Values object containing the data for the request
// - error: a boolean indicating if an error occurred
// - err: an error object if an error occurred
func getHeaderData(ctx context.Context, n int, buffer []byte) (map[string]string, url.Values, bool, error) {
	authKey := config.AccessKey + ":" + config.Secret
	encodedValue := base64.StdEncoding.EncodeToString([]byte(authKey))

	header := map[string]string{
		"Authorization": "Basic" + " " + encodedValue,
		"Content-Type":  "application/x-www-form-urlencoded",
		"Cookie":        "lwp=c=us&r=us&l=en&cs=19&s=dhs",
	}
	// Validate the token length
	if n < 1 || n > expectedBufferSize {
		tflog.Error(ctx, "Error in reading SAML token. Invalid token length")
		return nil, nil, true, fmt.Errorf("invalid token length: %d", n)
	}
	data := url.Values{}
	data.Set("grant_type", "urn:ietf:params:oauth:grant-type:token-exchange")
	data.Set("subject_token", string(buffer))
	data.Set("subject_token_type", "urn:ietf:params:oauth:token-type:saml2")
	return header, data, false, nil
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
	if len(header) < 10 {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	} else {
		return "", errors.New("too many headers")
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

// Config represents the configuration settings for the token helper.
// The struct contains the following fields:
// - ScriptPath: the path to the script file that retrieves the token.
// - DirPath: the directory where the script file is located.
// - TokenFileName: the name of the file where the token is stored.
// - URLLink: the URL link for the API request.
// - AccessKey: the access key for the API request.
// - Secret: the secret key for the API request.
type Config struct {
	ScriptPath    string
	DirPath       string
	TokenFileName string
	URLLink       string
	AccessKey     string
	Secret        string
}

// NewConfig creates a new Config struct with values obtained from environment variables.
//
// It returns a pointer to the Config struct and an error if any of the required environment variables are not set.
func validateConfig(localConfig *Config) error {
	if localConfig.ScriptPath == "" {
		return errors.New("APEX_SAML_TOKEN_SCRIPT_NAME not set")
	}

	if config.DirPath == "" {
		return errors.New("APEX_SAML_TOKEN_SCRIPT_DIR not set")
	}

	if config.TokenFileName == "" {
		return errors.New("APEX_SAML_TOKEN_OUTPUT_FILE_NAME not set")
	}

	if config.URLLink == "" {
		return errors.New("APEX_DI_TOKEN_URL not set")
	}

	if config.AccessKey == "" {
		return errors.New("APEX_TOKEN_ACCESS_KEY not set")
	}
	if config.Secret == "" {
		return errors.New("APEX_TOKEN_SECRET not set")
	}

	return nil
}

func runCommandWithTimeout(ctx context.Context, duration time.Duration, command string, dirPath string, args ...string) error {
	ctx1, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	cmd := exec.CommandContext(ctx1, command, args...)
	cmd.Dir = dirPath
	_, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("command timed out after %s", duration)
	}

	if err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	return nil
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
		if client != nil {
			client.GetConfig().AddDefaultHeader("Authorization", "Bearer "+token)
		}
		if jmsClient != nil {
			jmsClient.GetConfig().AddDefaultHeader("Authorization", "Bearer "+token)
		}
	}
	return nil
}
