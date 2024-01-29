# \AwsPermissionPoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AwsPermissionPoliciesCollection**](AwsPermissionPoliciesAPI.md#AwsPermissionPoliciesCollection) | **Get** /rest/services/storage/v1/aws-permission-policies | Get AWS permission policy
[**AwsPermissionPoliciesInstance**](AwsPermissionPoliciesAPI.md#AwsPermissionPoliciesInstance) | **Get** /rest/services/storage/v1/aws-permission-policies/{permission_policy_id} | Get AWS permission policy instance



## AwsPermissionPoliciesCollection

> AwsPermissionPoliciesCollection200Response AwsPermissionPoliciesCollection(ctx).Execute()

Get AWS permission policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AwsPermissionPoliciesAPI.AwsPermissionPoliciesCollection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AwsPermissionPoliciesAPI.AwsPermissionPoliciesCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsPermissionPoliciesCollection`: AwsPermissionPoliciesCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `AwsPermissionPoliciesAPI.AwsPermissionPoliciesCollection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAwsPermissionPoliciesCollectionRequest struct via the builder pattern


### Return type

[**AwsPermissionPoliciesCollection200Response**](AwsPermissionPoliciesCollection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsPermissionPoliciesInstance

> AwsPermissionPoliciesInstance AwsPermissionPoliciesInstance(ctx, permissionPolicyId).Execute()

Get AWS permission policy instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	permissionPolicyId := "permissionPolicyId_example" // string | AWS Permission Policy Instance ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AwsPermissionPoliciesAPI.AwsPermissionPoliciesInstance(context.Background(), permissionPolicyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AwsPermissionPoliciesAPI.AwsPermissionPoliciesInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsPermissionPoliciesInstance`: AwsPermissionPoliciesInstance
	fmt.Fprintf(os.Stdout, "Response from `AwsPermissionPoliciesAPI.AwsPermissionPoliciesInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**permissionPolicyId** | **string** | AWS Permission Policy Instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAwsPermissionPoliciesInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsPermissionPoliciesInstance**](AwsPermissionPoliciesInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

