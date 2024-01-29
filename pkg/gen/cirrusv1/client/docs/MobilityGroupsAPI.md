# \MobilityGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MobilityGroupsCollection**](MobilityGroupsAPI.md#MobilityGroupsCollection) | **Get** /rest/services/storage/v1/mobility-groups | Get Source Mobility Groups
[**MobilityGroupsCopy**](MobilityGroupsAPI.md#MobilityGroupsCopy) | **Post** /rest/services/storage/v1/mobility-groups/{mobility_group_id}/copy | Start copy operation on DMG
[**MobilityGroupsCreate**](MobilityGroupsAPI.md#MobilityGroupsCreate) | **Post** /rest/services/storage/v1/mobility-groups | Create Source Mobility Group
[**MobilityGroupsDelete**](MobilityGroupsAPI.md#MobilityGroupsDelete) | **Delete** /rest/services/storage/v1/mobility-groups/{mobility_group_id} | Delete source mobility group
[**MobilityGroupsInstance**](MobilityGroupsAPI.md#MobilityGroupsInstance) | **Get** /rest/services/storage/v1/mobility-groups/{mobility_group_id} | Get Source Mobility Group
[**MobilityGroupsModify**](MobilityGroupsAPI.md#MobilityGroupsModify) | **Patch** /rest/services/storage/v1/mobility-groups/{mobility_group_id} | Update source mobility group



## MobilityGroupsCollection

> MobilityGroupsCollection200Response MobilityGroupsCollection(ctx).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()

Get Source Mobility Groups



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
	resp, r, err := apiClient.MobilityGroupsAPI.MobilityGroupsCollection(context.Background()).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobilityGroupsAPI.MobilityGroupsCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilityGroupsCollection`: MobilityGroupsCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `MobilityGroupsAPI.MobilityGroupsCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMobilityGroupsCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | This filters rows in a query, by constraining the result to rows matching the property condition(s) specified. Multiple filters  can be combined with AND operators. | 
 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 
 **order** | **string** | Sorts the result set by the properties specified. Ascending order is default if not specified. | 
 **limit** | **int32** | Optional page size desired for the response. Default value is 500. | 
 **offset** | **int32** | Set the starting point within the collection of returned results. An  offset can only be set to a multiple of the page size. For example, for  a page size of 100, an offset can be to 0, 100, 200, 300, and so on.  The offset’s default value is 0. | 

### Return type

[**MobilityGroupsCollection200Response**](MobilityGroupsCollection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MobilityGroupsCopy

> Job MobilityGroupsCopy(ctx, mobilityGroupId).Async(async).StartCopyInput(startCopyInput).Execute()

Start copy operation on DMG



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
	mobilityGroupId := "mobilityGroupId_example" // string | Mobility group identifier
	async := true // bool | Asynchronous operation support (optional) (default to false)
	startCopyInput := *openapiclient.NewStartCopyInput() // StartCopyInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobilityGroupsAPI.MobilityGroupsCopy(context.Background(), mobilityGroupId).Async(async).StartCopyInput(startCopyInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobilityGroupsAPI.MobilityGroupsCopy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilityGroupsCopy`: Job
	fmt.Fprintf(os.Stdout, "Response from `MobilityGroupsAPI.MobilityGroupsCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobilityGroupId** | **string** | Mobility group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiMobilityGroupsCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Asynchronous operation support | [default to false]
 **startCopyInput** | [**StartCopyInput**](StartCopyInput.md) |  | 

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


## MobilityGroupsCreate

> Job MobilityGroupsCreate(ctx).Async(async).SourceMobilityGroupInput(sourceMobilityGroupInput).Execute()

Create Source Mobility Group



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
	sourceMobilityGroupInput := *openapiclient.NewSourceMobilityGroupInput("Source group", "POWERFLEX-ELMSIOPRODTST001", openapiclient.StorageSystemTypeEnum("POWERFLEX"), []string{"POWERFLEX-ELMSIO82000000__LUN__1000"}) // SourceMobilityGroupInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobilityGroupsAPI.MobilityGroupsCreate(context.Background()).Async(async).SourceMobilityGroupInput(sourceMobilityGroupInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobilityGroupsAPI.MobilityGroupsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilityGroupsCreate`: Job
	fmt.Fprintf(os.Stdout, "Response from `MobilityGroupsAPI.MobilityGroupsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMobilityGroupsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async** | **bool** | Asynchronous operation support | [default to false]
 **sourceMobilityGroupInput** | [**SourceMobilityGroupInput**](SourceMobilityGroupInput.md) |  | 

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


## MobilityGroupsDelete

> Job MobilityGroupsDelete(ctx, mobilityGroupId).Async(async).Execute()

Delete source mobility group



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
	mobilityGroupId := "mobilityGroupId_example" // string | Mobility group identifier
	async := true // bool | Asynchronous operation support (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobilityGroupsAPI.MobilityGroupsDelete(context.Background(), mobilityGroupId).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobilityGroupsAPI.MobilityGroupsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilityGroupsDelete`: Job
	fmt.Fprintf(os.Stdout, "Response from `MobilityGroupsAPI.MobilityGroupsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobilityGroupId** | **string** | Mobility group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiMobilityGroupsDeleteRequest struct via the builder pattern


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


## MobilityGroupsInstance

> MobilityGroup MobilityGroupsInstance(ctx, mobilityGroupId).Select_(select_).Execute()

Get Source Mobility Group



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
	mobilityGroupId := "mobilityGroupId_example" // string | Mobility group identifier
	select_ := "?select=id,name" // string | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobilityGroupsAPI.MobilityGroupsInstance(context.Background(), mobilityGroupId).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobilityGroupsAPI.MobilityGroupsInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilityGroupsInstance`: MobilityGroup
	fmt.Fprintf(os.Stdout, "Response from `MobilityGroupsAPI.MobilityGroupsInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobilityGroupId** | **string** | Mobility group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiMobilityGroupsInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 

### Return type

[**MobilityGroup**](MobilityGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MobilityGroupsModify

> Job MobilityGroupsModify(ctx, mobilityGroupId).Async(async).UpdateMobilityGroupInput(updateMobilityGroupInput).Execute()

Update source mobility group



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
	mobilityGroupId := "mobilityGroupId_example" // string | Mobility group identifier
	async := true // bool | Asynchronous operation support (optional) (default to false)
	updateMobilityGroupInput := *openapiclient.NewUpdateMobilityGroupInput() // UpdateMobilityGroupInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MobilityGroupsAPI.MobilityGroupsModify(context.Background(), mobilityGroupId).Async(async).UpdateMobilityGroupInput(updateMobilityGroupInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MobilityGroupsAPI.MobilityGroupsModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MobilityGroupsModify`: Job
	fmt.Fprintf(os.Stdout, "Response from `MobilityGroupsAPI.MobilityGroupsModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mobilityGroupId** | **string** | Mobility group identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiMobilityGroupsModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Asynchronous operation support | [default to false]
 **updateMobilityGroupInput** | [**UpdateMobilityGroupInput**](UpdateMobilityGroupInput.md) |  | 

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

