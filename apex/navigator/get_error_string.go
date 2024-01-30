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

package navigator

import (
	"io"
	"net/http"
)

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

// ResourceRetrieveError is the Error message for retrieveing resource ID failure
const ResourceRetrieveError = "Could not retrieve resourceID, unexpected error: "

// JobRetrieveError is the Error message for retrieveing job ID failure
const JobRetrieveError = "Could not retrieve job, unexpected error: "
