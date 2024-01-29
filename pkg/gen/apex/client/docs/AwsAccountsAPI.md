# \AwsAccountsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AwsAccountsCollection**](AwsAccountsAPI.md#AwsAccountsCollection) | **Get** /rest/services/storage/v1/aws-accounts | List connected AWS accounts
[**AwsAccountsCreate**](AwsAccountsAPI.md#AwsAccountsCreate) | **Post** /rest/services/storage/v1/aws-accounts | Connect AWS account
[**AwsAccountsDelete**](AwsAccountsAPI.md#AwsAccountsDelete) | **Delete** /rest/services/storage/v1/aws-accounts/{account_id} | Disconnect AWS account
[**AwsAccountsGenerateTrustPolicy**](AwsAccountsAPI.md#AwsAccountsGenerateTrustPolicy) | **Post** /rest/services/storage/v1/aws-accounts/generate-trust-policy | Retrieve AWS Trust Policy
[**AwsAccountsInstance**](AwsAccountsAPI.md#AwsAccountsInstance) | **Get** /rest/services/storage/v1/aws-accounts/{account_id} | Get connected AWS account
[**AwsAccountsModify**](AwsAccountsAPI.md#AwsAccountsModify) | **Patch** /rest/services/storage/v1/aws-accounts/{account_id} | Update AWS role selection



## AwsAccountsCollection

> AwsAccountsCollection200Response AwsAccountsCollection(ctx).Execute()

List connected AWS accounts



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
	resp, r, err := apiClient.AwsAccountsAPI.AwsAccountsCollection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AwsAccountsAPI.AwsAccountsCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsAccountsCollection`: AwsAccountsCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `AwsAccountsAPI.AwsAccountsCollection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAwsAccountsCollectionRequest struct via the builder pattern


### Return type

[**AwsAccountsCollection200Response**](AwsAccountsCollection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsAccountsCreate

> RedactedAwsAccountInstance AwsAccountsCreate(ctx).Async(async).CorrelationId(correlationId).AwsAccountsCreateInput(awsAccountsCreateInput).Execute()

Connect AWS account



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
	async := true // bool | Asynchronous operation support (optional) (default to false)
	correlationId := "correlationId_example" // string | Value is persisted in the resulting job, and any child jobs; allows queries to find disconnected synchronous jobs lost to a network disconnect or client or server failure. (optional)
	awsAccountsCreateInput := *openapiclient.NewAwsAccountsCreateInput("arn:aws:iam::123456789123:role/exampleRole") // AwsAccountsCreateInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AwsAccountsAPI.AwsAccountsCreate(context.Background()).Async(async).CorrelationId(correlationId).AwsAccountsCreateInput(awsAccountsCreateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AwsAccountsAPI.AwsAccountsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsAccountsCreate`: RedactedAwsAccountInstance
	fmt.Fprintf(os.Stdout, "Response from `AwsAccountsAPI.AwsAccountsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsAccountsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async** | **bool** | Asynchronous operation support | [default to false]
 **correlationId** | **string** | Value is persisted in the resulting job, and any child jobs; allows queries to find disconnected synchronous jobs lost to a network disconnect or client or server failure. | 
 **awsAccountsCreateInput** | [**AwsAccountsCreateInput**](AwsAccountsCreateInput.md) |  | 

### Return type

[**RedactedAwsAccountInstance**](RedactedAwsAccountInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsAccountsDelete

> Job AwsAccountsDelete(ctx, accountId).Async(async).CorrelationId(correlationId).Execute()

Disconnect AWS account



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
	accountId := "123456789123" // string | AWS account id.
	async := true // bool | Asynchronous operation support (optional) (default to false)
	correlationId := "correlationId_example" // string | Value is persisted in the resulting job, and any child jobs; allows queries to find disconnected synchronous jobs lost to a network disconnect or client or server failure. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AwsAccountsAPI.AwsAccountsDelete(context.Background(), accountId).Async(async).CorrelationId(correlationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AwsAccountsAPI.AwsAccountsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsAccountsDelete`: Job
	fmt.Fprintf(os.Stdout, "Response from `AwsAccountsAPI.AwsAccountsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | AWS account id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAwsAccountsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Asynchronous operation support | [default to false]
 **correlationId** | **string** | Value is persisted in the resulting job, and any child jobs; allows queries to find disconnected synchronous jobs lost to a network disconnect or client or server failure. | 

### Return type

[**Job**](Job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsAccountsGenerateTrustPolicy

> AwsTrustPolicy AwsAccountsGenerateTrustPolicy(ctx).Async(async).CorrelationId(correlationId).AwsTrustPolicyInput(awsTrustPolicyInput).Execute()

Retrieve AWS Trust Policy



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
	async := true // bool | Asynchronous operation support (optional) (default to false)
	correlationId := "correlationId_example" // string | Value is persisted in the resulting job, and any child jobs; allows queries to find disconnected synchronous jobs lost to a network disconnect or client or server failure. (optional)
	awsTrustPolicyInput := *openapiclient.NewAwsTrustPolicyInput("123456789123") // AwsTrustPolicyInput | The request body when attempting to generate an AWS trust policy. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AwsAccountsAPI.AwsAccountsGenerateTrustPolicy(context.Background()).Async(async).CorrelationId(correlationId).AwsTrustPolicyInput(awsTrustPolicyInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AwsAccountsAPI.AwsAccountsGenerateTrustPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsAccountsGenerateTrustPolicy`: AwsTrustPolicy
	fmt.Fprintf(os.Stdout, "Response from `AwsAccountsAPI.AwsAccountsGenerateTrustPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAwsAccountsGenerateTrustPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async** | **bool** | Asynchronous operation support | [default to false]
 **correlationId** | **string** | Value is persisted in the resulting job, and any child jobs; allows queries to find disconnected synchronous jobs lost to a network disconnect or client or server failure. | 
 **awsTrustPolicyInput** | [**AwsTrustPolicyInput**](AwsTrustPolicyInput.md) | The request body when attempting to generate an AWS trust policy. | 

### Return type

[**AwsTrustPolicy**](AwsTrustPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsAccountsInstance

> AwsAccountsInstance AwsAccountsInstance(ctx, accountId).Execute()

Get connected AWS account



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
	accountId := "123456789123" // string | AWS account id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AwsAccountsAPI.AwsAccountsInstance(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AwsAccountsAPI.AwsAccountsInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsAccountsInstance`: AwsAccountsInstance
	fmt.Fprintf(os.Stdout, "Response from `AwsAccountsAPI.AwsAccountsInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | AWS account id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAwsAccountsInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsAccountsInstance**](AwsAccountsInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AwsAccountsModify

> Job AwsAccountsModify(ctx, accountId).Async(async).CorrelationId(correlationId).AwsAccountModifyInput(awsAccountModifyInput).Execute()

Update AWS role selection



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
	accountId := "123456789123" // string | AWS account id.
	async := true // bool | Asynchronous operation support (optional) (default to false)
	correlationId := "correlationId_example" // string | Value is persisted in the resulting job, and any child jobs; allows queries to find disconnected synchronous jobs lost to a network disconnect or client or server failure. (optional)
	awsAccountModifyInput := *openapiclient.NewAwsAccountModifyInput("arn:aws:iam::123456789123:role/exampleRole") // AwsAccountModifyInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AwsAccountsAPI.AwsAccountsModify(context.Background(), accountId).Async(async).CorrelationId(correlationId).AwsAccountModifyInput(awsAccountModifyInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AwsAccountsAPI.AwsAccountsModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AwsAccountsModify`: Job
	fmt.Fprintf(os.Stdout, "Response from `AwsAccountsAPI.AwsAccountsModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | AWS account id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAwsAccountsModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Asynchronous operation support | [default to false]
 **correlationId** | **string** | Value is persisted in the resulting job, and any child jobs; allows queries to find disconnected synchronous jobs lost to a network disconnect or client or server failure. | 
 **awsAccountModifyInput** | [**AwsAccountModifyInput**](AwsAccountModifyInput.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

