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

package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	powerflexClient "dell/powerflex-client"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CreatePowerFlexClient creates a powerflex client
func CreatePowerFlexClient(ctx context.Context, powerflexEndpoint string, powerflexScheme string, insecure bool) (*powerflexClient.APIClient, error) {

	jar, err := cookiejar.New(nil)
	if err != nil {
		tflog.Error(ctx, "Got error while creating cookie jar")
	}

	httpclient := &http.Client{
		Timeout: (time.Duration(60) * time.Second),
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

	cfg := &powerflexClient.Configuration{
		Host:          powerflexEndpoint,
		Scheme:        powerflexScheme,
		HTTPClient:    httpclient,
		DefaultHeader: make(map[string]string),
		Debug:         false,
		Servers: []powerflexClient.ServerConfiguration{
			{
				Url:         powerflexScheme + "://" + powerflexEndpoint,
				Description: "No description provided",
			},
		},
	}

	apiClient := powerflexClient.NewAPIClient(cfg)
	return apiClient, nil
}
