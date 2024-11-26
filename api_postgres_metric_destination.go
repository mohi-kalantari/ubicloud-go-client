/*
Clover API

API for managing resources on Ubicloud

API version: 0.1.0
Contact: support@ubicloud.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// PostgresMetricDestinationApiService PostgresMetricDestinationApi service
type PostgresMetricDestinationApiService service

type ApiCreateLocationPostgresMetricDestinationRequest struct {
	ctx context.Context
	ApiService *PostgresMetricDestinationApiService
	location string
	postgresDatabaseName string
	projectId string
	createLocationPostgresMetricDestinationRequest *CreateLocationPostgresMetricDestinationRequest
}

func (r ApiCreateLocationPostgresMetricDestinationRequest) CreateLocationPostgresMetricDestinationRequest(createLocationPostgresMetricDestinationRequest CreateLocationPostgresMetricDestinationRequest) ApiCreateLocationPostgresMetricDestinationRequest {
	r.createLocationPostgresMetricDestinationRequest = &createLocationPostgresMetricDestinationRequest
	return r
}

func (r ApiCreateLocationPostgresMetricDestinationRequest) Execute() (*PostgresDetailed, *http.Response, error) {
	return r.ApiService.CreateLocationPostgresMetricDestinationExecute(r)
}

/*
CreateLocationPostgresMetricDestination Create a new Postgres Metric Destination

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param location The Ubicloud location/region
 @param postgresDatabaseName Postgres database name
 @param projectId ID of the project
 @return ApiCreateLocationPostgresMetricDestinationRequest
*/
func (a *PostgresMetricDestinationApiService) CreateLocationPostgresMetricDestination(ctx context.Context, location string, postgresDatabaseName string, projectId string) ApiCreateLocationPostgresMetricDestinationRequest {
	return ApiCreateLocationPostgresMetricDestinationRequest{
		ApiService: a,
		ctx: ctx,
		location: location,
		postgresDatabaseName: postgresDatabaseName,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return PostgresDetailed
func (a *PostgresMetricDestinationApiService) CreateLocationPostgresMetricDestinationExecute(r ApiCreateLocationPostgresMetricDestinationRequest) (*PostgresDetailed, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresDetailed
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresMetricDestinationApiService.CreateLocationPostgresMetricDestination")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{project_id}/location/{location}/postgres/{postgres_database_name}/metric-destination"
	localVarPath = strings.Replace(localVarPath, "{"+"location"+"}", url.PathEscape(parameterValueToString(r.location, "location")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"postgres_database_name"+"}", url.PathEscape(parameterValueToString(r.postgresDatabaseName, "postgresDatabaseName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createLocationPostgresMetricDestinationRequest == nil {
		return localVarReturnValue, nil, reportError("createLocationPostgresMetricDestinationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createLocationPostgresMetricDestinationRequest
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiDeleteLocationPostgresMetricDestinationRequest struct {
	ctx context.Context
	ApiService *PostgresMetricDestinationApiService
	location string
	metricDestinationId string
	postgresDatabaseName string
	projectId string
}

func (r ApiDeleteLocationPostgresMetricDestinationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLocationPostgresMetricDestinationExecute(r)
}

/*
DeleteLocationPostgresMetricDestination Delete a specific Metric Destination

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param location The Ubicloud location/region
 @param metricDestinationId Postgres Metric Destination ID
 @param postgresDatabaseName Postgres database name
 @param projectId ID of the project
 @return ApiDeleteLocationPostgresMetricDestinationRequest
*/
func (a *PostgresMetricDestinationApiService) DeleteLocationPostgresMetricDestination(ctx context.Context, location string, metricDestinationId string, postgresDatabaseName string, projectId string) ApiDeleteLocationPostgresMetricDestinationRequest {
	return ApiDeleteLocationPostgresMetricDestinationRequest{
		ApiService: a,
		ctx: ctx,
		location: location,
		metricDestinationId: metricDestinationId,
		postgresDatabaseName: postgresDatabaseName,
		projectId: projectId,
	}
}

// Execute executes the request
func (a *PostgresMetricDestinationApiService) DeleteLocationPostgresMetricDestinationExecute(r ApiDeleteLocationPostgresMetricDestinationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresMetricDestinationApiService.DeleteLocationPostgresMetricDestination")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{project_id}/location/{location}/postgres/{postgres_database_name}/metric-destination/{metric_destination_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"location"+"}", url.PathEscape(parameterValueToString(r.location, "location")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"metric_destination_id"+"}", url.PathEscape(parameterValueToString(r.metricDestinationId, "metricDestinationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"postgres_database_name"+"}", url.PathEscape(parameterValueToString(r.postgresDatabaseName, "postgresDatabaseName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}