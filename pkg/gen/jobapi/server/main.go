/*
 * APEX Orchestration Platform - Job Management System (JMS) API
 *
 * Provides management and visibility for APEX Orchestration Platform Jobs
 *
 * API version: IGNORED - see resource tag's x-api-version for the specific version of this resource.
 * Contact: apex.mars@dell.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"net/http"

	server "eos2git.cec.lab.emc.com/CIRRUS/cirrus-terraform-providers/pkg/gen/jobapi/server/go"
)

func main() {
	log.Printf("Server started")

	JobsAPIService := server.NewJobsAPIService()
	JobsAPIController := server.NewJobsAPIController(JobsAPIService)

	router := server.NewRouter(JobsAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
