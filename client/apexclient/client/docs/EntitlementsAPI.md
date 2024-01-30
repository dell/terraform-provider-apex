# \EntitlementsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntitlementsCollection**](EntitlementsAPI.md#EntitlementsCollection) | **Get** /rest/services/storage/v1/entitlements | Collection Query
[**EntitlementsInstance**](EntitlementsAPI.md#EntitlementsInstance) | **Get** /rest/services/storage/v1/entitlements/{id} | Instance Query



## EntitlementsCollection

> EntitlementsCollection200Response EntitlementsCollection(ctx).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()

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
	filter := "(simple)  ?filter=age ge 13 (complex)  ?filter=(age lt 18) or (name like "foo")" // string | This filters rows in a query, by constraining the result to rows matching the property condition(s) specified. Multiple filters can be combined with AND operators. (optional)
	select_ := "?select=id,name" // string | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. (optional)
	order := "?order=last_name,age.dsc" // string | Sorts the result set by the properties specified. Ascending order is default if not specified. (optional)
	limit := int32(500) // int32 | Optional page size desired for the response. Default value may be platform specified, either globally or per resource type. If paging is supported, the max size should be limited also, whether or not the client specifies a limit. A reasonable max page size in on the order of 1000 items but can vary. Smaller could be better for large instance representations. (optional)
	offset := int32(56) // int32 | Optional page size desired for the response. Default value may be platform specified, either globally or per resource type. If paging is supported, the max size should be limited also, whether or not the client specifies a limit. A reasonable max page size in on the order of 1000 items but can vary. Smaller could be better for large instance representations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.EntitlementsCollection(context.Background()).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.EntitlementsCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitlementsCollection`: EntitlementsCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.EntitlementsCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | This filters rows in a query, by constraining the result to rows matching the property condition(s) specified. Multiple filters can be combined with AND operators. | 
 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 
 **order** | **string** | Sorts the result set by the properties specified. Ascending order is default if not specified. | 
 **limit** | **int32** | Optional page size desired for the response. Default value may be platform specified, either globally or per resource type. If paging is supported, the max size should be limited also, whether or not the client specifies a limit. A reasonable max page size in on the order of 1000 items but can vary. Smaller could be better for large instance representations. | 
 **offset** | **int32** | Optional page size desired for the response. Default value may be platform specified, either globally or per resource type. If paging is supported, the max size should be limited also, whether or not the client specifies a limit. A reasonable max page size in on the order of 1000 items but can vary. Smaller could be better for large instance representations. | 

### Return type

[**EntitlementsCollection200Response**](EntitlementsCollection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntitlementsInstance

> EntitlementsInstance EntitlementsInstance(ctx, id).Select_(select_).Execute()

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
	id := "id_example" // string | Unique identifier of the entitlements instance.
	select_ := "?select=id,name" // string | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.EntitlementsInstance(context.Background(), id).Select_(select_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.EntitlementsInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EntitlementsInstance`: EntitlementsInstance
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.EntitlementsInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the entitlements instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEntitlementsInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **select_** | **string** | By default, all properties of resource instances are returned, but this can cause responses to be large and slow. Get faster responses by  selecting specific properties to be returned. | 

### Return type

[**EntitlementsInstance**](EntitlementsInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

