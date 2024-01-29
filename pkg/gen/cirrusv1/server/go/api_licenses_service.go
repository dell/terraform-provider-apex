/*
 * APEX Navigator for Multicloud Storage REST APIs
 *
 * The public API definitions for APEX Navigator for Multicloud Storage
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"context"
	"net/http"
	"errors"
)

// LicensesAPIService is a service that implements the logic for the LicensesAPIServicer
// This service should implement the business logic for every endpoint for the LicensesAPI API.
// Include any external packages or services that will be required by this service.
type LicensesAPIService struct {
}

// NewLicensesAPIService creates a default api service
func NewLicensesAPIService() LicensesAPIServicer {
	return &LicensesAPIService{}
}

// LicensesCollection - Collection Query
func (s *LicensesAPIService) LicensesCollection(ctx context.Context, filter string, select_ string, order string, limit int32, offset int32) (ImplResponse, error) {
	// TODO - update LicensesCollection with the required logic for this service method.
	// Add api_licenses_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, LicensesCollection200Response{}) or use other options such as http.Ok ...
	// return Response(200, LicensesCollection200Response{}), nil

	// TODO: Uncomment the next line to return response Response(206, LicensesCollection206Response{}) or use other options such as http.Ok ...
	// return Response(206, LicensesCollection206Response{}), nil

	// TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(400, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(401, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(403, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(403, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(429, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(429, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(500, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(503, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(503, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("LicensesCollection method not implemented")
}

// LicensesInstance - Instance Query
func (s *LicensesAPIService) LicensesInstance(ctx context.Context, id string, select_ string) (ImplResponse, error) {
	// TODO - update LicensesInstance with the required logic for this service method.
	// Add api_licenses_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, LicensesInstance{}) or use other options such as http.Ok ...
	// return Response(200, LicensesInstance{}), nil

	// TODO: Uncomment the next line to return response Response(400, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(400, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(401, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(401, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(403, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(403, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(404, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(404, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(429, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(429, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(500, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(500, ErrorResponse{}), nil

	// TODO: Uncomment the next line to return response Response(503, ErrorResponse{}) or use other options such as http.Ok ...
	// return Response(503, ErrorResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("LicensesInstance method not implemented")
}
