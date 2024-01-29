/*
 * APEX Orchestration Platform - Job Management System (JMS) API
 *
 * Provides management and visibility for APEX Orchestration Platform Jobs
 *
 * API version: IGNORED - see resource tag's x-api-version for the specific version of this resource.
 * Contact: apex.mars@dell.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// JobsAPIController binds http requests to an api service and writes the service results to the http response
type JobsAPIController struct {
	service      JobsAPIServicer
	errorHandler ErrorHandler
}

// JobsAPIOption for how the controller is set up.
type JobsAPIOption func(*JobsAPIController)

// WithJobsAPIErrorHandler inject ErrorHandler into controller
func WithJobsAPIErrorHandler(h ErrorHandler) JobsAPIOption {
	return func(c *JobsAPIController) {
		c.errorHandler = h
	}
}

// NewJobsAPIController creates a default api controller
func NewJobsAPIController(s JobsAPIServicer, opts ...JobsAPIOption) Router {
	controller := &JobsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the JobsAPIController
func (c *JobsAPIController) Routes() Routes {
	return Routes{
		"JobsCancel": Route{
			strings.ToUpper("Post"),
			"/rest/v1/jobs/{id}/cancel",
			c.JobsCancel,
		},
		"JobsCollection": Route{
			strings.ToUpper("Get"),
			"/rest/v1/jobs",
			c.JobsCollection,
		},
		"JobsInstance": Route{
			strings.ToUpper("Get"),
			"/rest/v1/jobs/{id}",
			c.JobsInstance,
		},
		"JobsPause": Route{
			strings.ToUpper("Post"),
			"/rest/v1/jobs/{id}/pause",
			c.JobsPause,
		},
		"JobsResume": Route{
			strings.ToUpper("Post"),
			"/rest/v1/jobs/{id}/resume",
			c.JobsResume,
		},
	}
}

// JobsCancel - Cancel
func (c *JobsAPIController) JobsCancel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	result, err := c.service.JobsCancel(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// JobsCollection - Collection Query
func (c *JobsAPIController) JobsCollection(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limitParam, err := parseNumericParameter[int32](
		query.Get("limit"),
		WithDefaultOrParse[int32](1000, parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	offsetParam, err := parseNumericParameter[int32](
		query.Get("offset"),
		WithDefaultOrParse[int32](0, parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	orderParam := "timeCreated.dsc"
	if query.Has("order") {
		orderParam = query.Get("order")
	}
	var filterParam string
	if query.Has("filter") {
		filterParam = query.Get("filter")
	}
	result, err := c.service.JobsCollection(r.Context(), limitParam, offsetParam, orderParam, filterParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// JobsInstance - Instance Query
func (c *JobsAPIController) JobsInstance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	result, err := c.service.JobsInstance(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// JobsPause - Pause
func (c *JobsAPIController) JobsPause(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	result, err := c.service.JobsPause(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// JobsResume - Resume
func (c *JobsAPIController) JobsResume(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &RequiredError{"id"}, nil)
		return
	}
	result, err := c.service.JobsResume(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}
