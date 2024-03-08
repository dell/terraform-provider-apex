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
	"io"
	"net/http"
	"time"

	jmsClient "github.com/dell/terraform-provider-apex/client/jobsclient/client"
)

// ConvertTimeToString converts time to string
func ConvertTimeToString(pTime *time.Time) string {
	if pTime != nil {
		return pTime.String()
	}

	return ""
}

// GetErrorString returns the error string from the response
func GetErrorString(err error, status *http.Response) string {
	newErr := ""
	if err != nil {
		newErr = err.Error()
	}
	if status != nil {
		bodyBytes, _ := io.ReadAll(status.Body)
		newErr = newErr + " " + string(bodyBytes)
	}
	return newErr
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
