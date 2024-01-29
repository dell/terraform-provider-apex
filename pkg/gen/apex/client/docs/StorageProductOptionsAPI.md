# \StorageProductOptionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StorageProductOptionsCollection**](StorageProductOptionsAPI.md#StorageProductOptionsCollection) | **Get** /rest/services/storage/v1/storage-product-options | GET offer templates for storage products



## StorageProductOptionsCollection

> StorageProductOptionsCollection200Response StorageProductOptionsCollection(ctx).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).XDELLEMCVISIBILITY(xDELLEMCVISIBILITY).Execute()

GET offer templates for storage products



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
	select_ := "?select=id,name" // string | By default, all properties of resource instances are returned,  but this can cause responses to be large and slow. Get faster  responses by selecting specific properties to be returned. (optional)
	order := "?order=last_name,age.dsc" // string | Sorts the result set by the properties specified. Ascending order is default if not specified. (optional)
	limit := int32(500) // int32 | Optional page size desired for the response. Default value is 100. (optional)
	offset := int32(56) // int32 | Optional page offset for the response. This value determines the starting point within the collection of resource results and needs to be multiple of page size. For instance, for page size, 100, offset should be 0, 100, 200, .. Default value is 0. (optional)
	xDELLEMCVISIBILITY := "xDELLEMCVISIBILITY_example" // string | This header MUST NOT be documented in customer documentation, only in internal documentation. When a request occurs with this header is set to INTERNAL, API elements marked with those visibility levels are valid in the context of that request, otherwise use of any INTERNAL visibility element must result in the same error behavior as if the element was not defined/supported on the system at all. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageProductOptionsAPI.StorageProductOptionsCollection(context.Background()).Filter(filter).Select_(select_).Order(order).Limit(limit).Offset(offset).XDELLEMCVISIBILITY(xDELLEMCVISIBILITY).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageProductOptionsAPI.StorageProductOptionsCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageProductOptionsCollection`: StorageProductOptionsCollection200Response
	fmt.Fprintf(os.Stdout, "Response from `StorageProductOptionsAPI.StorageProductOptionsCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorageProductOptionsCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | This filters rows in a query, by constraining the result to rows matching the property condition(s) specified. Multiple filters  can be combined with AND operators. | 
 **select_** | **string** | By default, all properties of resource instances are returned,  but this can cause responses to be large and slow. Get faster  responses by selecting specific properties to be returned. | 
 **order** | **string** | Sorts the result set by the properties specified. Ascending order is default if not specified. | 
 **limit** | **int32** | Optional page size desired for the response. Default value is 100. | 
 **offset** | **int32** | Optional page offset for the response. This value determines the starting point within the collection of resource results and needs to be multiple of page size. For instance, for page size, 100, offset should be 0, 100, 200, .. Default value is 0. | 
 **xDELLEMCVISIBILITY** | **string** | This header MUST NOT be documented in customer documentation, only in internal documentation. When a request occurs with this header is set to INTERNAL, API elements marked with those visibility levels are valid in the context of that request, otherwise use of any INTERNAL visibility element must result in the same error behavior as if the element was not defined/supported on the system at all. | 

### Return type

[**StorageProductOptionsCollection200Response**](StorageProductOptionsCollection200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

