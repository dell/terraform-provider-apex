# \StorageSystemTokensAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageSystemTokensCreate**](StorageSystemTokensAPI.md#StorageSystemTokensCreate) | **Post** /rest/services/storage/v1/storage-system-tokens | Activate the storage system token.
[**StorageSystemTokensInstance**](StorageSystemTokensAPI.md#StorageSystemTokensInstance) | **Get** /rest/services/storage/v1/storage-system-tokens/{id} | Instance by id.
[**StorageSystemTokensModify**](StorageSystemTokensAPI.md#StorageSystemTokensModify) | **Patch** /rest/services/storage/v1/storage-system-tokens/{id} | Update the storage system token.



## StorageSystemTokensCreate

> ResourceId StorageSystemTokensCreate(ctx).StorageSystemTokensCreateRequest(storageSystemTokensCreateRequest).Execute()

Activate the storage system token.



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
	storageSystemTokensCreateRequest := *openapiclient.NewStorageSystemTokensCreateRequest(openapiclient.StorageProductEnum("POWERFLEX"), "SystemId_example", "AccessToken_example") // StorageSystemTokensCreateRequest | Storage system tokens create request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageSystemTokensAPI.StorageSystemTokensCreate(context.Background()).StorageSystemTokensCreateRequest(storageSystemTokensCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageSystemTokensAPI.StorageSystemTokensCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageSystemTokensCreate`: ResourceId
	fmt.Fprintf(os.Stdout, "Response from `StorageSystemTokensAPI.StorageSystemTokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageSystemTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storageSystemTokensCreateRequest** | [**StorageSystemTokensCreateRequest**](StorageSystemTokensCreateRequest.md) | Storage system tokens create request | 

### Return type

[**ResourceId**](ResourceId.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageSystemTokensInstance

> StorageSystemTokensInstance StorageSystemTokensInstance(ctx, id).Execute()

Instance by id.



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
	id := "id_example" // string | The storage system token ID to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageSystemTokensAPI.StorageSystemTokensInstance(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageSystemTokensAPI.StorageSystemTokensInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageSystemTokensInstance`: StorageSystemTokensInstance
	fmt.Fprintf(os.Stdout, "Response from `StorageSystemTokensAPI.StorageSystemTokensInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The storage system token ID to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageSystemTokensInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorageSystemTokensInstance**](StorageSystemTokensInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageSystemTokensModify

> StorageSystemTokensModify(ctx, id).StorageSystemTokensModifyRequest(storageSystemTokensModifyRequest).Execute()

Update the storage system token.



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
	id := "id_example" // string | The storage system token ID to retrieve.
	storageSystemTokensModifyRequest := *openapiclient.NewStorageSystemTokensModifyRequest("AccessToken_example") // StorageSystemTokensModifyRequest | Storage system tokens modify request (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageSystemTokensAPI.StorageSystemTokensModify(context.Background(), id).StorageSystemTokensModifyRequest(storageSystemTokensModifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageSystemTokensAPI.StorageSystemTokensModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The storage system token ID to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageSystemTokensModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageSystemTokensModifyRequest** | [**StorageSystemTokensModifyRequest**](StorageSystemTokensModifyRequest.md) | Storage system tokens modify request | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

