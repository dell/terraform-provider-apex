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

package provider

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	client "dell/apex-client"
	jmsClient "dell/apex-job-client"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/dell/terraform-provider-apex/apex/helper"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Clients is the structure for the API client and JMS client.
type Clients struct {
	APIClient *client.APIClient
	JMSClient *jmsClient.APIClient
}

// NewApexJmsClient creates a new instance of Clients, which contains two API clients: apiClient and jobsClient.
//
// The function takes the following parameters:
// - ctx: the context.Context object for cancellation and timeouts.
// - hostURL: the URL of the host API.
// - jmsURL: the URL of the JMS API.
// - insecure: a boolean indicating whether to skip SSL certificate verification or not.
//
// It returns a pointer to Clients and an error.
func NewApexJmsClient(ctx context.Context, hostURL url.URL, jmsURL url.URL, token string, insecure bool) (*Clients, error) {
	// Setup a User-Agent for your API client :
	userAgent := "terraform-provider-apex/1.0.0"
	jar, err := cookiejar.New(nil)
	if err != nil {
		tflog.Error(ctx, "Got error while creating cookie jar")
	}

	httpclient := &http.Client{
		Timeout: time.Second * time.Duration(10000000),
		Jar:     jar,
	}
	if insecure {
		/* #nosec */
		httpclient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				InsecureSkipVerify: true,
			},
		}
	} else {
		// Loading system certs by default if insecure is set to false
		pool, err := x509.SystemCertPool()
		if err != nil {
			errSysCerts := errors.New("unable to initialize cert pool from system")
			return nil, errSysCerts
		}
		httpclient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				RootCAs:            pool,
				InsecureSkipVerify: false,
			},
		}
	}
	cfg := &client.Configuration{
		Host:       hostURL.Host,
		Scheme:     hostURL.Scheme,
		UserAgent:  userAgent,
		HTTPClient: httpclient,
		Servers: []client.ServerConfiguration{
			{
				URL:         hostURL.Path,
				Description: "API Client for Terraform Provider",
			},
		},
	}
	if token == "" || len(token) < 30 {
		var err1 error
		token, err1 = helper.GetNewToken(ctx)
		if err1 != nil {
			tflog.Error(ctx, "Error in getting new token", map[string]interface{}{
				"Error": err1,
			})
			return nil, err1
		}
	}

	cfg.DefaultHeader = getHeaders()
	cfg.AddDefaultHeader("Authorization", "Bearer "+token)

	apiClient := client.NewAPIClient(cfg)

	jobCfg := &jmsClient.Configuration{
		Host:       jmsURL.Host,
		Scheme:     jmsURL.Scheme,
		HTTPClient: httpclient,
		Servers: jmsClient.ServerConfigurations{
			{
				URL:         jmsURL.Path,
				Description: "API Client for Terraform Provider",
			},
		},
	}
	jobCfg.DefaultHeader = getHeaders()
	jobCfg.AddDefaultHeader("Authorization", "Bearer "+token)

	jobsClient := jmsClient.NewAPIClient(jobCfg)

	return &Clients{apiClient, jobsClient}, nil
}

func getHeaders() map[string]string {
	header := make(map[string]string)
	header["Accept"] = "*/*"
	return header
}
