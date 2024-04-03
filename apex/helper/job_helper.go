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
	jmsClient "dell/apex-job-client"
	"encoding/json"
	"fmt"
)

// GetErrorMessageFromBody returns the error message from the response body
func GetErrorMessageFromBody(ctx context.Context, body string) (string, error) {
	var bodyMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &bodyMap)
	if err != nil {
		return "", fmt.Errorf("Failed to parse JSON body: %v", err)
	}

	messages, ok := bodyMap["messages"].([]interface{})
	if !ok {
		return "", fmt.Errorf("Invalid format: 'messages' field is missing or not an array")
	}

	if len(messages) == 0 {
		return "", fmt.Errorf("No messages found in the response body")
	}

	messageObj, ok := messages[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("Invalid format: 'message' object is missing or not a map")
	}

	message, ok := messageObj["message"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid format: 'message' field is missing or not a string")
	}
	return message, nil
}

// WaitForJobToComplete returns the job status and waits for the job to complete
func WaitForJobToComplete(ctx context.Context, jobsClient *jmsClient.APIClient, jobID string) (string, error) {
	// Fetching Job Status
	poller := NewPoller(jobsClient)
	return poller.WaitForResource(ctx, jobID)
}

// GetJobStatus returns the job status
func GetJobStatus(ctx context.Context, jobsClient *jmsClient.APIClient, jobID string) (*jmsClient.JobsInstance, error) {
	// Fetching Job Status
	poller := NewPoller(jobsClient)
	return poller.GetJob(ctx, jobID)
}
