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
	"fmt"
	"time"

	"net/http"

	jobsClient "dell/apex-job-client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Poller interface
type Poller interface {
	WaitForResource(ctx context.Context, id string) (string, error)
	CancelJob(ctx context.Context, id string) (*http.Response, error)
	GetJob(ctx context.Context, id string) (*jobsClient.JobsInstance, error)
}

// NewPoller returns a new poller object
func NewPoller(jmsClient *jobsClient.APIClient) Poller {
	return &poller{
		jmsClient: jmsClient,
	}
}

type poller struct {
	jmsClient *jobsClient.APIClient
}

// WaitForResource simulates a long operation of a job.
func (p poller) WaitForResource(ctx context.Context, id string) (string, error) {
	// try 3 times for getting updated tokens
	var retryUnauthorizedCounter = 0
	for {
		// TODO: resolve unmarshal error that causes infinite loop on terminating job status (failed, succeeded)
		job, status, err := p.jmsClient.JobsAPI.JobsInstance(ctx, id).Execute()
		//if token gets expired before job is completed
		if status != nil && status.StatusCode == http.StatusUnauthorized {
			err = UpdateToken(ctx, nil, p.jmsClient)
			if err != nil {
				retryUnauthorizedCounter = retryUnauthorizedCounter + 1
				if retryUnauthorizedCounter == 3 {
					tflog.Error(ctx, "Error updating new token", map[string]interface{}{"Error": err})
					return "", err
				}
				time.Sleep(5 * time.Second)
			}
			continue
		}
		if err != nil && (status != nil && status.StatusCode != http.StatusOK) || job == nil {
			//  If Error, don't check for status/resource, just poll again
			tflog.Error(ctx, "Error getting job", map[string]interface{}{"Error": err})
			time.Sleep(5 * time.Second)
			continue
		}

		if err := status.Body.Close(); err != nil {
			tflog.Error(ctx, "Error Closing response body:", map[string]interface{}{"Error": err})
		}

		switch *job.State {
		case jobsClient.JOBSTATEENUM_QUEUED:
		case jobsClient.JOBSTATEENUM_RUNNING:
		case jobsClient.JOBSTATEENUM_SUCCEEDED:
			if job.Resource.Id == "" {
				return "", fmt.Errorf("resource id not returned when job completed")
			}
			return job.Resource.Id, nil
		case jobsClient.JOBSTATEENUM_FAILED:
			resp, _ := job.GetResponseOk()
			var message = ""
			if resp != nil {
				body, ok := resp.GetBodyOk()
				if ok {
					message, err = GetErrorMessageFromBody(ctx, body)
					if err != nil {
						return "", fmt.Errorf("job failed: " + err.Error())
					}
				}
			}
			return "", fmt.Errorf("job failed: " + message)
		case jobsClient.JOBSTATEENUM_CANCELLING:
		case jobsClient.JOBSTATEENUM_CANCELLED:
			return "", fmt.Errorf("job cancelled")
		case jobsClient.JOBSTATEENUM_PAUSING:
		case jobsClient.JOBSTATEENUM_PAUSED:
			resp, cancelErr := p.CancelJob(ctx, id)
			if cancelErr != nil {
				return "", fmt.Errorf("job failed to cancel")
			}
			if respErr := resp.Body.Close(); respErr != nil {
				fmt.Print("Error Closing response body")
			}
		case jobsClient.JOBSTATEENUM_RESUMING:
		case jobsClient.JOBSTATEENUM_SCHEDULED:
		case jobsClient.JOBSTATEENUM_COMPLETED_WITH_MESSAGES:
			return job.Resource.Id, nil
		}
		time.Sleep(5 * time.Second)
	}
}

// GetJob returns a job
func (p poller) GetJob(ctx context.Context, id string) (*jobsClient.JobsInstance, error) {
	job, status, err := p.jmsClient.JobsAPI.JobsInstance(ctx, id).Execute()
	if err != nil && status.StatusCode != http.StatusOK || job == nil {
		return nil, err
	}

	if err := status.Body.Close(); err != nil {
		fmt.Print("Error Closing response body:", err)
	}
	return job, nil
}

// CancelJob cancels an existing job.
func (p poller) CancelJob(ctx context.Context, id string) (*http.Response, error) {
	_, status, err := p.jmsClient.JobsAPI.JobsCancel(ctx, id).Execute()
	return status, err
}
