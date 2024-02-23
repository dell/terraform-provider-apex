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

	client "github.com/dell/terraform-provider-apex/client/jobsclient/client"
)

// Poller interface
type Poller interface {
	WaitForResource(ctx context.Context, id string) (string, error)
	CancelJob(ctx context.Context, id string) (*http.Response, error)
	GetJob(ctx context.Context, id string) (*client.JobsInstance, error)
}

// NewPoller returns a new poller object
func NewPoller(jmsClient *client.APIClient) Poller {
	return &poller{
		client: jmsClient,
	}
}

type poller struct {
	client *client.APIClient
}

// WaitForResource simulates a long operation of a job.
func (p poller) WaitForResource(ctx context.Context, id string) (string, error) {
	for {
		// TODO: resolve unmarshal error that causes infinite loop on terminating job status (failed, succeeded)
		job, status, err := p.client.JobsAPI.JobsInstance(ctx, id).Execute()
		if err != nil && status.StatusCode != http.StatusOK || job == nil {
			//  If Error, don't check for status/resource, just poll again
			fmt.Print("Error getting job")
			time.Sleep(5 * time.Second)
			continue
		}

		if err := status.Body.Close(); err != nil {
			fmt.Print("Error Closing response body:", err)
		}

		switch *job.State {
		case client.JOBSTATEENUM_QUEUED:
		case client.JOBSTATEENUM_RUNNING:
		case client.JOBSTATEENUM_SUCCEEDED:
			if job.Resource.Id == "" {
				return "", fmt.Errorf("resource id not returned when job completed")
			}
			return job.Resource.Id, nil
		case client.JOBSTATEENUM_FAILED:
			return "", fmt.Errorf("job failed")
		case client.JOBSTATEENUM_CANCELLING:
		case client.JOBSTATEENUM_CANCELLED:
			return "", fmt.Errorf("job cancelled")
		case client.JOBSTATEENUM_PAUSING:
		case client.JOBSTATEENUM_PAUSED:
			resp, cancelErr := p.CancelJob(ctx, id)
			if cancelErr != nil {
				return "", fmt.Errorf("job failed to cancel")
			}
			if respErr := resp.Body.Close(); respErr != nil {
				fmt.Print("Error Closing response body")
			}
		case client.JOBSTATEENUM_RESUMING:
		case client.JOBSTATEENUM_SCHEDULED:
		case client.JOBSTATEENUM_COMPLETED_WITH_MESSAGES:
			return job.Resource.Id, nil
		}
		time.Sleep(5 * time.Second)
	}
}

// GetJob returns a job
func (p poller) GetJob(ctx context.Context, id string) (*client.JobsInstance, error) {
	job, status, err := p.client.JobsAPI.JobsInstance(ctx, id).Execute()
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
	_, status, err := p.client.JobsAPI.JobsCancel(ctx, id).Execute()
	return status, err
}
