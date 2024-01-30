# \VolumesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumesCollection**](VolumesAPI.md#VolumesCollection) | **Get** /rest/services/storage/v1/volumes | Collection Query
[**VolumesInstance**](VolumesAPI.md#VolumesInstance) | **Get** /rest/services/storage/v1/volumes/{id} | Instance Query



## VolumesCollection

> VolumesCollection200Response VolumesCollection(ctx).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()

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
	resp, r, err := apiClient.VolumesAPI.VolumesCollection(context.Background()).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesCollection`: VolumesCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVolumesCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | This filters rows in a query, by constraining the result to rows matching the property condition(s) specified. Multiple filters  can be combined with AND operators. | 
 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 
 **order** | **string** | Sorts the result set by the properties specified. Ascending order is default if not specified. | 
 **limit** | **int32** | Optional page size desired for the response. Default value is 500. | 
 **offset** | **int32** | Set the starting point within the collection of returned results. An  offset can only be set to a multiple of the page size. For example, for  a page size of 100, an offset can be to 0, 100, 200, 300, and so on.  The offset’s default value is 0. | 

### Return type

[**VolumesCollection200Response**](VolumesCollection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumesInstance

> VolumesInstance VolumesInstance(ctx, id).Select_(select_).Execute()

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
	id := "id_example" // string | The volume ID to retrieve
	select_ := "?select=id,name" // string | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.VolumesInstance(context.Background(), id).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.VolumesInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VolumesInstance`: VolumesInstance
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.VolumesInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The volume ID to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumesInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 

### Return type

[**VolumesInstance**](VolumesInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

