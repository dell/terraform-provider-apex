# \MobilityTargetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MobilityTargetsCollection**](MobilityTargetsAPI.md#MobilityTargetsCollection) | **Get** /rest/services/storage/v1/mobility-targets | Get Mobility Targets
[**MobilityTargetsCreate**](MobilityTargetsAPI.md#MobilityTargetsCreate) | **Post** /rest/services/storage/v1/mobility-targets | Create a mobility target
[**MobilityTargetsDelete**](MobilityTargetsAPI.md#MobilityTargetsDelete) | **Delete** /rest/services/storage/v1/mobility-targets/{mobility_target_id} | Delete mobility target
[**MobilityTargetsInstance**](MobilityTargetsAPI.md#MobilityTargetsInstance) | **Get** /rest/services/storage/v1/mobility-targets/{mobility_target_id} | Get Mobility Target
[**MobilityTargetsModify**](MobilityTargetsAPI.md#MobilityTargetsModify) | **Patch** /rest/services/storage/v1/mobility-targets/{mobility_target_id} | Update mobility target



## MobilityTargetsCollection

> MobilityTargetsCollection200Response MobilityTargetsCollection(ctx).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()

Get Mobility Targets



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
	resp, r, err := apiClient.MobilityTargetsAPI.MobilityTargetsCollection(context.Background()).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobilityTargetsAPI.MobilityTargetsCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilityTargetsCollection`: MobilityTargetsCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `MobilityTargetsAPI.MobilityTargetsCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMobilityTargetsCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | This filters rows in a query, by constraining the result to rows matching the property condition(s) specified. Multiple filters  can be combined with AND operators. | 
 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 
 **order** | **string** | Sorts the result set by the properties specified. Ascending order is default if not specified. | 
 **limit** | **int32** | Optional page size desired for the response. Default value is 500. | 
 **offset** | **int32** | Set the starting point within the collection of returned results. An  offset can only be set to a multiple of the page size. For example, for  a page size of 100, an offset can be to 0, 100, 200, 300, and so on.  The offset’s default value is 0. | 

### Return type

[**MobilityTargetsCollection200Response**](MobilityTargetsCollection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MobilityTargetsCreate

> Job MobilityTargetsCreate(ctx).Async(async).CreateTargetInput(createTargetInput).Execute()

Create a mobility target



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
	createTargetInput := *openapiclient.NewCreateTargetInput("POWERFLEX-ELMSIO08200000__DATAMOBILITYGROUP__fcdecfaf-c61e-4b4d-8f89-65c6ef00d1234", "Analytics target", "POWERFLEX-AWSSIO08200000", openapiclient.StorageSystemTypeEnum("POWERFLEX"), *openapiclient.NewTargetSystemOptions()) // CreateTargetInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobilityTargetsAPI.MobilityTargetsCreate(context.Background()).Async(async).CreateTargetInput(createTargetInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobilityTargetsAPI.MobilityTargetsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilityTargetsCreate`: Job
	fmt.Fprintf(os.Stdout, "Response from `MobilityTargetsAPI.MobilityTargetsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMobilityTargetsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async** | **bool** | Asynchronous operation support | [default to false]
 **createTargetInput** | [**CreateTargetInput**](CreateTargetInput.md) |  | 

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


## MobilityTargetsDelete

> Job MobilityTargetsDelete(ctx, mobilityTargetId).Async(async).Execute()

Delete mobility target



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
	mobilityTargetId := "mobilityTargetId_example" // string | Mobility target identifier
	async := true // bool | Asynchronous operation support (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobilityTargetsAPI.MobilityTargetsDelete(context.Background(), mobilityTargetId).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobilityTargetsAPI.MobilityTargetsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilityTargetsDelete`: Job
	fmt.Fprintf(os.Stdout, "Response from `MobilityTargetsAPI.MobilityTargetsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobilityTargetId** | **string** | Mobility target identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiMobilityTargetsDeleteRequest struct via the builder pattern


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


## MobilityTargetsInstance

> MobilityTarget MobilityTargetsInstance(ctx, mobilityTargetId).Select_(select_).Execute()

Get Mobility Target



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
	mobilityTargetId := "mobilityTargetId_example" // string | Mobility target identifier
	select_ := "?select=id,name" // string | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobilityTargetsAPI.MobilityTargetsInstance(context.Background(), mobilityTargetId).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobilityTargetsAPI.MobilityTargetsInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilityTargetsInstance`: MobilityTarget
	fmt.Fprintf(os.Stdout, "Response from `MobilityTargetsAPI.MobilityTargetsInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobilityTargetId** | **string** | Mobility target identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiMobilityTargetsInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 

### Return type

[**MobilityTarget**](MobilityTarget.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MobilityTargetsModify

> Job MobilityTargetsModify(ctx, mobilityTargetId).Async(async).UpdateMobilityTargetInput(updateMobilityTargetInput).Execute()

Update mobility target



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
	mobilityTargetId := "mobilityTargetId_example" // string | Mobility target identifier
	async := true // bool | Asynchronous operation support (optional) (default to false)
	updateMobilityTargetInput := *openapiclient.NewUpdateMobilityTargetInput() // UpdateMobilityTargetInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobilityTargetsAPI.MobilityTargetsModify(context.Background(), mobilityTargetId).Async(async).UpdateMobilityTargetInput(updateMobilityTargetInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobilityTargetsAPI.MobilityTargetsModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilityTargetsModify`: Job
	fmt.Fprintf(os.Stdout, "Response from `MobilityTargetsAPI.MobilityTargetsModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobilityTargetId** | **string** | Mobility target identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiMobilityTargetsModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Asynchronous operation support | [default to false]
 **updateMobilityTargetInput** | [**UpdateMobilityTargetInput**](UpdateMobilityTargetInput.md) |  | 

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

