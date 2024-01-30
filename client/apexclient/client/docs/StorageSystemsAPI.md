# \StorageSystemsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageSystemsCollection**](StorageSystemsAPI.md#StorageSystemsCollection) | **Get** /rest/services/storage/v1/storage-systems | Collection Query
[**StorageSystemsCreate**](StorageSystemsAPI.md#StorageSystemsCreate) | **Post** /rest/services/storage/v1/storage-systems | Deploy a new storage system in cloud
[**StorageSystemsDelete**](StorageSystemsAPI.md#StorageSystemsDelete) | **Delete** /rest/services/storage/v1/storage-systems/{id} | Decommission storage system
[**StorageSystemsFinalizeTrust**](StorageSystemsAPI.md#StorageSystemsFinalizeTrust) | **Post** /rest/services/storage/v1/storage-systems/{id}/finalize-trust | Finalize system trust
[**StorageSystemsInitializeTrust**](StorageSystemsAPI.md#StorageSystemsInitializeTrust) | **Post** /rest/services/storage/v1/storage-systems/{id}/initialize-trust | Initialize system trust
[**StorageSystemsInstance**](StorageSystemsAPI.md#StorageSystemsInstance) | **Get** /rest/services/storage/v1/storage-systems/{id} | Instance Query



## StorageSystemsCollection

> StorageSystemsCollection200Response StorageSystemsCollection(ctx).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()

Collection Query



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
	filter := "(simple)  ?filter=age ge 13 (complex)  ?filter=(age lt 18) or (name like "foo")" // string | This filters rows in a query, by constraining the result to rows matching the property condition(s) specified. Multiple filters  can be combined with AND operators. (optional)
	select_ := "?select=id,name" // string | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. (optional)
	order := "?order=last_name,age.dsc" // string | Sorts the result set by the properties specified. Ascending order is default if not specified. (optional)
	limit := int32(500) // int32 | Optional page size desired for the response. Default value is 500. (optional)
	offset := int32(56) // int32 | Set the starting point within the collection of returned results. An  offset can only be set to a multiple of the page size. For example, for  a page size of 100, an offset can be to 0, 100, 200, 300, and so on.  The offset’s default value is 0. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageSystemsAPI.StorageSystemsCollection(context.Background()).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageSystemsAPI.StorageSystemsCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageSystemsCollection`: StorageSystemsCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `StorageSystemsAPI.StorageSystemsCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageSystemsCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | This filters rows in a query, by constraining the result to rows matching the property condition(s) specified. Multiple filters  can be combined with AND operators. | 
 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 
 **order** | **string** | Sorts the result set by the properties specified. Ascending order is default if not specified. | 
 **limit** | **int32** | Optional page size desired for the response. Default value is 500. | 
 **offset** | **int32** | Set the starting point within the collection of returned results. An  offset can only be set to a multiple of the page size. For example, for  a page size of 100, an offset can be to 0, 100, 200, 300, and so on.  The offset’s default value is 0. | 

### Return type

[**StorageSystemsCollection200Response**](StorageSystemsCollection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageSystemsCreate

> Job StorageSystemsCreate(ctx).Async(async).XDELLEMCVISIBILITY(xDELLEMCVISIBILITY).StorageSystemDeploymentRequest(storageSystemDeploymentRequest).Execute()

Deploy a new storage system in cloud



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
	xDELLEMCVISIBILITY := "xDELLEMCVISIBILITY_example" // string | This header MUST NOT be documented in customer documentation, only in internal documentation. When a request occurs with this header is set to INTERNAL, API elements marked with those visibility levels are valid in the context of that request, otherwise use of any INTERNAL visibility element must result in the same error behavior as if the element was not defined/supported on the system at all. (optional)
	storageSystemDeploymentRequest := *openapiclient.NewStorageSystemDeploymentRequest("Name_example", openapiclient.StorageSystemDeploymentRequest_cloud_options{AWSCloudOptions: openapiclient.NewAWSCloudOptions("CloudType_example", "AccountId_example", "us-east-1", "SshKeyName_example", *openapiclient.NewVpc(), []openapiclient.SubnetOptions{*openapiclient.NewSubnetOptions()})}, openapiclient.StorageSystemDeploymentRequest_storage_options{PowerFlexStorageOptions: openapiclient.NewPowerFlexStorageOptions("POWERFLEX")}, false) // StorageSystemDeploymentRequest | Deployment request input (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageSystemsAPI.StorageSystemsCreate(context.Background()).Async(async).XDELLEMCVISIBILITY(xDELLEMCVISIBILITY).StorageSystemDeploymentRequest(storageSystemDeploymentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageSystemsAPI.StorageSystemsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageSystemsCreate`: Job
	fmt.Fprintf(os.Stdout, "Response from `StorageSystemsAPI.StorageSystemsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageSystemsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async** | **bool** | Asynchronous operation support | [default to false]
 **xDELLEMCVISIBILITY** | **string** | This header MUST NOT be documented in customer documentation, only in internal documentation. When a request occurs with this header is set to INTERNAL, API elements marked with those visibility levels are valid in the context of that request, otherwise use of any INTERNAL visibility element must result in the same error behavior as if the element was not defined/supported on the system at all. | 
 **storageSystemDeploymentRequest** | [**StorageSystemDeploymentRequest**](StorageSystemDeploymentRequest.md) | Deployment request input | 

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


## StorageSystemsDelete

> Job StorageSystemsDelete(ctx, id).Async(async).Execute()

Decommission storage system



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
	id := "id_example" // string | Storage system unique identifier.
	async := true // bool | Asynchronous operation support (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageSystemsAPI.StorageSystemsDelete(context.Background(), id).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageSystemsAPI.StorageSystemsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageSystemsDelete`: Job
	fmt.Fprintf(os.Stdout, "Response from `StorageSystemsAPI.StorageSystemsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Storage system unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageSystemsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Asynchronous operation support | [default to false]

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


## StorageSystemsFinalizeTrust

> StorageSystemsFinalizeTrust(ctx, id).StorageSystemsFinalizeTrustPostRequest(storageSystemsFinalizeTrustPostRequest).Execute()

Finalize system trust



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
	id := "id_example" // string | Storage system unique identifier.
	storageSystemsFinalizeTrustPostRequest := *openapiclient.NewStorageSystemsFinalizeTrustPostRequest("ClientId_example", "ClientSecret_example") // StorageSystemsFinalizeTrustPostRequest | Request object to finalize the trust between on-prem storage system and APEX Navigator for Multicloud Storage. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StorageSystemsAPI.StorageSystemsFinalizeTrust(context.Background(), id).StorageSystemsFinalizeTrustPostRequest(storageSystemsFinalizeTrustPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageSystemsAPI.StorageSystemsFinalizeTrust``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Storage system unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageSystemsFinalizeTrustRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageSystemsFinalizeTrustPostRequest** | [**StorageSystemsFinalizeTrustPostRequest**](StorageSystemsFinalizeTrustPostRequest.md) | Request object to finalize the trust between on-prem storage system and APEX Navigator for Multicloud Storage. | 

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


## StorageSystemsInitializeTrust

> StorageSystemsInitializeTrustResponse StorageSystemsInitializeTrust(ctx, id).Execute()

Initialize system trust



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
	id := "id_example" // string | Storage system unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageSystemsAPI.StorageSystemsInitializeTrust(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageSystemsAPI.StorageSystemsInitializeTrust``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageSystemsInitializeTrust`: StorageSystemsInitializeTrustResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageSystemsAPI.StorageSystemsInitializeTrust`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Storage system unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageSystemsInitializeTrustRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StorageSystemsInitializeTrustResponse**](StorageSystemsInitializeTrustResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorageSystemsInstance

> StorageSystemsInstance StorageSystemsInstance(ctx, id).Select_(select_).Execute()

Instance Query



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
	id := "id_example" // string | The storage system ID to retrieve
	select_ := "?select=id,name" // string | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageSystemsAPI.StorageSystemsInstance(context.Background(), id).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageSystemsAPI.StorageSystemsInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageSystemsInstance`: StorageSystemsInstance
	fmt.Fprintf(os.Stdout, "Response from `StorageSystemsAPI.StorageSystemsInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The storage system ID to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiStorageSystemsInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 

### Return type

[**StorageSystemsInstance**](StorageSystemsInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

