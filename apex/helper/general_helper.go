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
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types"
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
	var newErr string
	if err != nil {
		newErr = err.Error()
	}
	if status != nil {
		if bodyBytes, err := io.ReadAll(status.Body); err == nil {
			newErr += " " + string(bodyBytes)
		}
	}
	return newErr
}

// CreateFilter creates a filter query string
func CreateFilter(filter []string, key string) string {
	filterString := ""
	for k, value := range filter {
		if k == len(filter)-1 {
			filterString += `(` + key + ` eq "` + value + `")`
		} else {
			filterString += `(` + key + ` eq "` + value + `") or `
		}
	}
	return filterString
}

// ConvertToStringSlice converts []types.String to []string
func ConvertToStringSlice(typesSlice []types.String) []string {
	result := make([]string, len(typesSlice))
	for i, v := range typesSlice {
		result[i] = v.ValueString()
	}
	return result
}
