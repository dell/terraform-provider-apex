# \ClonesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClonesCollection**](ClonesAPI.md#ClonesCollection) | **Get** /rest/services/storage/v1/clones | Get all clones
[**ClonesCreate**](ClonesAPI.md#ClonesCreate) | **Post** /rest/services/storage/v1/clones | Create Clone
[**ClonesDelete**](ClonesAPI.md#ClonesDelete) | **Delete** /rest/services/storage/v1/clones/{clone_id} | Delete a clone
[**ClonesInstance**](ClonesAPI.md#ClonesInstance) | **Get** /rest/services/storage/v1/clones/{clone_id} | Get a clone
[**ClonesMap**](ClonesAPI.md#ClonesMap) | **Post** /rest/services/storage/v1/clones/{clone_id}/map | Map hosts to clone
[**ClonesModify**](ClonesAPI.md#ClonesModify) | **Patch** /rest/services/storage/v1/clones/{clone_id} | Update clone
[**ClonesRefresh**](ClonesAPI.md#ClonesRefresh) | **Post** /rest/services/storage/v1/clones/{clone_id}/refresh | Refresh Clone
[**ClonesUnmap**](ClonesAPI.md#ClonesUnmap) | **Post** /rest/services/storage/v1/clones/{clone_id}/unmap | Unmap hosts from clone



## ClonesCollection

> ClonesCollection200Response ClonesCollection(ctx).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()

Get all clones



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
	resp, r, err := apiClient.ClonesAPI.ClonesCollection(context.Background()).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClonesAPI.ClonesCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClonesCollection`: ClonesCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `ClonesAPI.ClonesCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClonesCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | This filters rows in a query, by constraining the result to rows matching the property condition(s) specified. Multiple filters  can be combined with AND operators. | 
 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 
 **order** | **string** | Sorts the result set by the properties specified. Ascending order is default if not specified. | 
 **limit** | **int32** | Optional page size desired for the response. Default value is 500. | 
 **offset** | **int32** | Set the starting point within the collection of returned results. An  offset can only be set to a multiple of the page size. For example, for  a page size of 100, an offset can be to 0, 100, 200, 300, and so on.  The offset’s default value is 0. | 

### Return type

[**ClonesCollection200Response**](ClonesCollection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClonesCreate

> Job ClonesCreate(ctx).Async(async).CloneCreateInput(cloneCreateInput).Execute()

Create Clone



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
	cloneCreateInput := *openapiclient.NewCloneCreateInput("test clone", "POWERFLEX-AWSSIO08200000__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5") // CloneCreateInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClonesAPI.ClonesCreate(context.Background()).Async(async).CloneCreateInput(cloneCreateInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClonesAPI.ClonesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClonesCreate`: Job
	fmt.Fprintf(os.Stdout, "Response from `ClonesAPI.ClonesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClonesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async** | **bool** | Asynchronous operation support | [default to false]
 **cloneCreateInput** | [**CloneCreateInput**](CloneCreateInput.md) |  | 

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


## ClonesDelete

> Job ClonesDelete(ctx, cloneId).Async(async).DeleteCloneInput(deleteCloneInput).Execute()

Delete a clone



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
	cloneId := "POWERFLEX-AWSSIO08200000__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5" // string | ID of the clone
	async := true // bool | Asynchronous operation support (optional) (default to false)
	deleteCloneInput := *openapiclient.NewDeleteCloneInput(false) // DeleteCloneInput | Based on the parameter in the request, either fully clean up a clone from the system or remove the clone from APEX management but do not remove the system objects. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClonesAPI.ClonesDelete(context.Background(), cloneId).Async(async).DeleteCloneInput(deleteCloneInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClonesAPI.ClonesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClonesDelete`: Job
	fmt.Fprintf(os.Stdout, "Response from `ClonesAPI.ClonesDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloneId** | **string** | ID of the clone | 

### Other Parameters

Other parameters are passed through a pointer to a apiClonesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Asynchronous operation support | [default to false]
 **deleteCloneInput** | [**DeleteCloneInput**](DeleteCloneInput.md) | Based on the parameter in the request, either fully clean up a clone from the system or remove the clone from APEX management but do not remove the system objects. | 

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


## ClonesInstance

> Clone ClonesInstance(ctx, cloneId).Select_(select_).Execute()

Get a clone



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
	cloneId := "POWERFLEX-AWSSIO08200000__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5" // string | ID of the clone
	select_ := "?select=id,name" // string | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClonesAPI.ClonesInstance(context.Background(), cloneId).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClonesAPI.ClonesInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClonesInstance`: Clone
	fmt.Fprintf(os.Stdout, "Response from `ClonesAPI.ClonesInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloneId** | **string** | ID of the clone | 

### Other Parameters

Other parameters are passed through a pointer to a apiClonesInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 

### Return type

[**Clone**](Clone.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClonesMap

> Job ClonesMap(ctx, cloneId).Async(async).MapInput(mapInput).Execute()

Map hosts to clone



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
	cloneId := "POWERFLEX-AWSSIO08200000__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5" // string | ID of the clone
	async := true // bool | Asynchronous operation support (optional) (default to false)
	mapInput := *openapiclient.NewMapInput([]string{"POWERFLEX-ELMSIOENG10004__HOST__a1810e4300000007"}) // MapInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClonesAPI.ClonesMap(context.Background(), cloneId).Async(async).MapInput(mapInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClonesAPI.ClonesMap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClonesMap`: Job
	fmt.Fprintf(os.Stdout, "Response from `ClonesAPI.ClonesMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloneId** | **string** | ID of the clone | 

### Other Parameters

Other parameters are passed through a pointer to a apiClonesMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Asynchronous operation support | [default to false]
 **mapInput** | [**MapInput**](MapInput.md) |  | 

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


## ClonesModify

> Job ClonesModify(ctx, cloneId).Async(async).UpdateCloneInput(updateCloneInput).Execute()

Update clone



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
	cloneId := "POWERFLEX-AWSSIO08200000__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5" // string | ID of the clone
	async := true // bool | Asynchronous operation support (optional) (default to false)
	updateCloneInput := *openapiclient.NewUpdateCloneInput() // UpdateCloneInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClonesAPI.ClonesModify(context.Background(), cloneId).Async(async).UpdateCloneInput(updateCloneInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClonesAPI.ClonesModify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClonesModify`: Job
	fmt.Fprintf(os.Stdout, "Response from `ClonesAPI.ClonesModify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloneId** | **string** | ID of the clone | 

### Other Parameters

Other parameters are passed through a pointer to a apiClonesModifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Asynchronous operation support | [default to false]
 **updateCloneInput** | [**UpdateCloneInput**](UpdateCloneInput.md) |  | 

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


## ClonesRefresh

> Job ClonesRefresh(ctx, cloneId).Async(async).Execute()

Refresh Clone



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
	cloneId := "POWERFLEX-AWSSIO08200000__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5" // string | ID of the clone
	async := true // bool | Asynchronous operation support (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClonesAPI.ClonesRefresh(context.Background(), cloneId).Async(async).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClonesAPI.ClonesRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClonesRefresh`: Job
	fmt.Fprintf(os.Stdout, "Response from `ClonesAPI.ClonesRefresh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloneId** | **string** | ID of the clone | 

### Other Parameters

Other parameters are passed through a pointer to a apiClonesRefreshRequest struct via the builder pattern


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


## ClonesUnmap

> Job ClonesUnmap(ctx, cloneId).Async(async).UnmapInput(unmapInput).Execute()

Unmap hosts from clone



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
	cloneId := "POWERFLEX-AWSSIO08200000__DATAMOBILITYGROUP__a90fcfaf-c61e-4b4d-8f89-65c6ef00dfd5" // string | ID of the clone
	async := true // bool | Asynchronous operation support (optional) (default to false)
	unmapInput := *openapiclient.NewUnmapInput([]string{"POWERFLEX-ELMSIOENG10004__DATAMOBILITYHOST__31182e76-6828-44f1-b7d8-5038b440f001"}) // UnmapInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClonesAPI.ClonesUnmap(context.Background(), cloneId).Async(async).UnmapInput(unmapInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClonesAPI.ClonesUnmap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClonesUnmap`: Job
	fmt.Fprintf(os.Stdout, "Response from `ClonesAPI.ClonesUnmap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloneId** | **string** | ID of the clone | 

### Other Parameters

Other parameters are passed through a pointer to a apiClonesUnmapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Asynchronous operation support | [default to false]
 **unmapInput** | [**UnmapInput**](UnmapInput.md) |  | 

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

