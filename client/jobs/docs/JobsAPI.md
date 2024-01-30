# \JobsAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobsCancel**](JobsAPI.md#JobsCancel) | **Post** /rest/v1/jobs/{id}/cancel | Cancel
[**JobsCollection**](JobsAPI.md#JobsCollection) | **Get** /rest/v1/jobs | Collection Query
[**JobsInstance**](JobsAPI.md#JobsInstance) | **Get** /rest/v1/jobs/{id} | Instance Query
[**JobsPause**](JobsAPI.md#JobsPause) | **Post** /rest/v1/jobs/{id}/pause | Pause
[**JobsResume**](JobsAPI.md#JobsResume) | **Post** /rest/v1/jobs/{id}/resume | Resume



## JobsCancel

> JobsInstance JobsCancel(ctx, id).Execute()

Cancel



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsAPI.JobsCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.JobsCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JobsCancel`: JobsInstance
    fmt.Fprintf(os.Stdout, "Response from `JobsAPI.JobsCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobsInstance**](JobsInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsCollection

> JobsCollection JobsCollection(ctx).Limit(limit).Offset(offset).Order(order).Filter(filter).Execute()

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
    limit := int32(56) // int32 | The limit on the number of items returned per page (optional) (default to 1000)
    offset := int32(56) // int32 | The zero based index of the first item in the page of results (optional) (default to 0)
    order := "order_example" // string | Sort the jobs based on the specified parameter and order the jobs in ascending or descending order (optional) (default to "timeCreated.dsc")
    filter := "filter_example" // string | Filter the jobs based on provided query. Format: (state in (RUNNING,FAILED)) and (text like \"provision\") and (affectedResourceEntities.actionType eq type)  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsAPI.JobsCollection(context.Background()).Limit(limit).Offset(offset).Order(order).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.JobsCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JobsCollection`: JobsCollection
    fmt.Fprintf(os.Stdout, "Response from `JobsAPI.JobsCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJobsCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The limit on the number of items returned per page | [default to 1000]
 **offset** | **int32** | The zero based index of the first item in the page of results | [default to 0]
 **order** | **string** | Sort the jobs based on the specified parameter and order the jobs in ascending or descending order | [default to &quot;timeCreated.dsc&quot;]
 **filter** | **string** | Filter the jobs based on provided query. Format: (state in (RUNNING,FAILED)) and (text like \&quot;provision\&quot;) and (affectedResourceEntities.actionType eq type)  | 

### Return type

[**JobsCollection**](JobsCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsInstance

> JobsInstance JobsInstance(ctx, id).Execute()

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
    id := "id_example" // string | Unique identifier of the jobs instance.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsAPI.JobsInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.JobsInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JobsInstance`: JobsInstance
    fmt.Fprintf(os.Stdout, "Response from `JobsAPI.JobsInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the jobs instance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobsInstance**](JobsInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsPause

> JobsInstance JobsPause(ctx, id).Execute()

Pause



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsAPI.JobsPause(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.JobsPause``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JobsPause`: JobsInstance
    fmt.Fprintf(os.Stdout, "Response from `JobsAPI.JobsPause`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsPauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobsInstance**](JobsInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobsResume

> JobsInstance JobsResume(ctx, id).Execute()

Resume



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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsAPI.JobsResume(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsAPI.JobsResume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JobsResume`: JobsInstance
    fmt.Fprintf(os.Stdout, "Response from `JobsAPI.JobsResume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsResumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobsInstance**](JobsInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

