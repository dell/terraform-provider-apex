/*
APEX Navigator for Multicloud Storage REST APIs

The public API definitions for APEX Navigator for Multicloud Storage

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// StorageProductOptionsAPIService StorageProductOptionsAPI service
type StorageProductOptionsAPIService service

type ApiStorageProductOptionsCollectionRequest struct {
	ctx context.Context
	ApiService *StorageProductOptionsAPIService
	filter *string
	select_ *string
	order *string
	limit *int32
	offset *int32
	xDELLEMCVISIBILITY *string
}

// This filters rows in a query, by constraining the result to rows matching the property condition(s) specified. Multiple filters  can be combined with AND operators.
func (r ApiStorageProductOptionsCollectionRequest) Filter(filter string) ApiStorageProductOptionsCollectionRequest {
	r.filter = &filter
	return r
}

// By default, all properties of resource instances are returned,  but this can cause responses to be large and slow. Get faster  responses by selecting specific properties to be returned.
func (r ApiStorageProductOptionsCollectionRequest) Select_(select_ string) ApiStorageProductOptionsCollectionRequest {
	r.select_ = &select_
	return r
}

// Sorts the result set by the properties specified. Ascending order is default if not specified.
func (r ApiStorageProductOptionsCollectionRequest) Order(order string) ApiStorageProductOptionsCollectionRequest {
	r.order = &order
	return r
}

// Optional page size desired for the response. Default value is 100.
func (r ApiStorageProductOptionsCollectionRequest) Limit(limit int32) ApiStorageProductOptionsCollectionRequest {
	r.limit = &limit
	return r
}

// Optional page offset for the response. This value determines the starting point within the collection of resource results and needs to be multiple of page size. For instance, for page size, 100, offset should be 0, 100, 200, .. Default value is 0.
func (r ApiStorageProductOptionsCollectionRequest) Offset(offset int32) ApiStorageProductOptionsCollectionRequest {
	r.offset = &offset
	return r
}

// This header MUST NOT be documented in customer documentation, only in internal documentation. When a request occurs with this header is set to INTERNAL, API elements marked with those visibility levels are valid in the context of that request, otherwise use of any INTERNAL visibility element must result in the same error behavior as if the element was not defined/supported on the system at all.
func (r ApiStorageProductOptionsCollectionRequest) XDELLEMCVISIBILITY(xDELLEMCVISIBILITY string) ApiStorageProductOptionsCollectionRequest {
	r.xDELLEMCVISIBILITY = &xDELLEMCVISIBILITY
	return r
}

func (r ApiStorageProductOptionsCollectionRequest) Execute() (*StorageProductOptionsCollection200Response, *http.Response, error) {
	return r.ApiService.StorageProductOptionsCollectionExecute(r)
}

/*
StorageProductOptionsCollection GET offer templates for storage products

Gets available offer template(s) based on input

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStorageProductOptionsCollectionRequest
*/
func (a *StorageProductOptionsAPIService) StorageProductOptionsCollection(ctx context.Context) ApiStorageProductOptionsCollectionRequest {
	return ApiStorageProductOptionsCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StorageProductOptionsCollection200Response
func (a *StorageProductOptionsAPIService) StorageProductOptionsCollectionExecute(r ApiStorageProductOptionsCollectionRequest) (*StorageProductOptionsCollection200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StorageProductOptionsCollection200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StorageProductOptionsAPIService.StorageProductOptionsCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rest/services/storage/v1/storage-product-options"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.select_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "select", r.select_, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xDELLEMCVISIBILITY != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-DELL-EMC-VISIBILITY", r.xDELLEMCVISIBILITY, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
