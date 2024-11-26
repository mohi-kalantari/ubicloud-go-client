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


// PostgresFirewallRuleApiService PostgresFirewallRuleApi service
type PostgresFirewallRuleApiService service

type ApiCreateLocationPostgresFirewallRuleRequest struct {
	ctx context.Context
	ApiService *PostgresFirewallRuleApiService
	location string
	postgresDatabaseName string
	projectId string
	createLocationPostgresFirewallRuleWithIdRequest *CreateLocationPostgresFirewallRuleWithIdRequest
}

func (r ApiCreateLocationPostgresFirewallRuleRequest) CreateLocationPostgresFirewallRuleWithIdRequest(createLocationPostgresFirewallRuleWithIdRequest CreateLocationPostgresFirewallRuleWithIdRequest) ApiCreateLocationPostgresFirewallRuleRequest {
	r.createLocationPostgresFirewallRuleWithIdRequest = &createLocationPostgresFirewallRuleWithIdRequest
	return r
}

func (r ApiCreateLocationPostgresFirewallRuleRequest) Execute() (*PostgresFirewallRule, *http.Response, error) {
	return r.ApiService.CreateLocationPostgresFirewallRuleExecute(r)
}

/*
CreateLocationPostgresFirewallRule Create a new postgres firewall rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param location The Ubicloud location/region
 @param postgresDatabaseName Postgres database name
 @param projectId ID of the project
 @return ApiCreateLocationPostgresFirewallRuleRequest
*/
func (a *PostgresFirewallRuleApiService) CreateLocationPostgresFirewallRule(ctx context.Context, location string, postgresDatabaseName string, projectId string) ApiCreateLocationPostgresFirewallRuleRequest {
	return ApiCreateLocationPostgresFirewallRuleRequest{
		ApiService: a,
		ctx: ctx,
		location: location,
		postgresDatabaseName: postgresDatabaseName,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return PostgresFirewallRule
func (a *PostgresFirewallRuleApiService) CreateLocationPostgresFirewallRuleExecute(r ApiCreateLocationPostgresFirewallRuleRequest) (*PostgresFirewallRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresFirewallRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresFirewallRuleApiService.CreateLocationPostgresFirewallRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{project_id}/location/{location}/postgres/{postgres_database_name}/firewall-rule"
	localVarPath = strings.Replace(localVarPath, "{"+"location"+"}", url.PathEscape(parameterValueToString(r.location, "location")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"postgres_database_name"+"}", url.PathEscape(parameterValueToString(r.postgresDatabaseName, "postgresDatabaseName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createLocationPostgresFirewallRuleWithIdRequest == nil {
		return localVarReturnValue, nil, reportError("createLocationPostgresFirewallRuleWithIdRequest is required and must be specified")
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
	localVarPostBody = r.createLocationPostgresFirewallRuleWithIdRequest
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

type ApiCreateLocationPostgresFirewallRuleWithIdRequest struct {
	ctx context.Context
	ApiService *PostgresFirewallRuleApiService
	location string
	postgresDatabaseId string
	projectId string
	createLocationPostgresFirewallRuleWithIdRequest *CreateLocationPostgresFirewallRuleWithIdRequest
}

func (r ApiCreateLocationPostgresFirewallRuleWithIdRequest) CreateLocationPostgresFirewallRuleWithIdRequest(createLocationPostgresFirewallRuleWithIdRequest CreateLocationPostgresFirewallRuleWithIdRequest) ApiCreateLocationPostgresFirewallRuleWithIdRequest {
	r.createLocationPostgresFirewallRuleWithIdRequest = &createLocationPostgresFirewallRuleWithIdRequest
	return r
}

func (r ApiCreateLocationPostgresFirewallRuleWithIdRequest) Execute() (*PostgresFirewallRule, *http.Response, error) {
	return r.ApiService.CreateLocationPostgresFirewallRuleWithIdExecute(r)
}

/*
CreateLocationPostgresFirewallRuleWithId Create a new Postgres firewall rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param location The Ubicloud location/region
 @param postgresDatabaseId Postgres database ID
 @param projectId ID of the project
 @return ApiCreateLocationPostgresFirewallRuleWithIdRequest
*/
func (a *PostgresFirewallRuleApiService) CreateLocationPostgresFirewallRuleWithId(ctx context.Context, location string, postgresDatabaseId string, projectId string) ApiCreateLocationPostgresFirewallRuleWithIdRequest {
	return ApiCreateLocationPostgresFirewallRuleWithIdRequest{
		ApiService: a,
		ctx: ctx,
		location: location,
		postgresDatabaseId: postgresDatabaseId,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return PostgresFirewallRule
func (a *PostgresFirewallRuleApiService) CreateLocationPostgresFirewallRuleWithIdExecute(r ApiCreateLocationPostgresFirewallRuleWithIdRequest) (*PostgresFirewallRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresFirewallRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresFirewallRuleApiService.CreateLocationPostgresFirewallRuleWithId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{project_id}/location/{location}/postgres/_{postgres_database_id}/firewall-rule"
	localVarPath = strings.Replace(localVarPath, "{"+"location"+"}", url.PathEscape(parameterValueToString(r.location, "location")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"postgres_database_id"+"}", url.PathEscape(parameterValueToString(r.postgresDatabaseId, "postgresDatabaseId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createLocationPostgresFirewallRuleWithIdRequest == nil {
		return localVarReturnValue, nil, reportError("createLocationPostgresFirewallRuleWithIdRequest is required and must be specified")
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
	localVarPostBody = r.createLocationPostgresFirewallRuleWithIdRequest
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

type ApiCreateLocationPostgresFirewallRuleWithIdWithIdRequest struct {
	ctx context.Context
	ApiService *PostgresFirewallRuleApiService
	location string
	postgresDatabaseId string
	projectId string
	firewallRuleId string
	createLocationPostgresFirewallRuleWithIdWithIdRequest *CreateLocationPostgresFirewallRuleWithIdWithIdRequest
}

func (r ApiCreateLocationPostgresFirewallRuleWithIdWithIdRequest) CreateLocationPostgresFirewallRuleWithIdWithIdRequest(createLocationPostgresFirewallRuleWithIdWithIdRequest CreateLocationPostgresFirewallRuleWithIdWithIdRequest) ApiCreateLocationPostgresFirewallRuleWithIdWithIdRequest {
	r.createLocationPostgresFirewallRuleWithIdWithIdRequest = &createLocationPostgresFirewallRuleWithIdWithIdRequest
	return r
}

func (r ApiCreateLocationPostgresFirewallRuleWithIdWithIdRequest) Execute() (*PostgresFirewallRule, *http.Response, error) {
	return r.ApiService.CreateLocationPostgresFirewallRuleWithIdWithIdExecute(r)
}

/*
CreateLocationPostgresFirewallRuleWithIdWithId Create a new Postgres firewall rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param location The Ubicloud location/region
 @param postgresDatabaseId Postgres database ID
 @param projectId ID of the project
 @param firewallRuleId ID of the firewall rule
 @return ApiCreateLocationPostgresFirewallRuleWithIdWithIdRequest
*/
func (a *PostgresFirewallRuleApiService) CreateLocationPostgresFirewallRuleWithIdWithId(ctx context.Context, location string, postgresDatabaseId string, projectId string, firewallRuleId string) ApiCreateLocationPostgresFirewallRuleWithIdWithIdRequest {
	return ApiCreateLocationPostgresFirewallRuleWithIdWithIdRequest{
		ApiService: a,
		ctx: ctx,
		location: location,
		postgresDatabaseId: postgresDatabaseId,
		projectId: projectId,
		firewallRuleId: firewallRuleId,
	}
}

// Execute executes the request
//  @return PostgresFirewallRule
func (a *PostgresFirewallRuleApiService) CreateLocationPostgresFirewallRuleWithIdWithIdExecute(r ApiCreateLocationPostgresFirewallRuleWithIdWithIdRequest) (*PostgresFirewallRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresFirewallRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresFirewallRuleApiService.CreateLocationPostgresFirewallRuleWithIdWithId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/firewall-rule/{firewall_rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"location"+"}", url.PathEscape(parameterValueToString(r.location, "location")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"postgres_database_id"+"}", url.PathEscape(parameterValueToString(r.postgresDatabaseId, "postgresDatabaseId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"firewall_rule_id"+"}", url.PathEscape(parameterValueToString(r.firewallRuleId, "firewallRuleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createLocationPostgresFirewallRuleWithIdWithIdRequest == nil {
		return localVarReturnValue, nil, reportError("createLocationPostgresFirewallRuleWithIdWithIdRequest is required and must be specified")
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
	localVarPostBody = r.createLocationPostgresFirewallRuleWithIdWithIdRequest
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

type ApiGetLocationPostgresFirewallRuleDetailsWithIdRequest struct {
	ctx context.Context
	ApiService *PostgresFirewallRuleApiService
	location string
	postgresDatabaseId string
	projectId string
	firewallRuleId string
}

func (r ApiGetLocationPostgresFirewallRuleDetailsWithIdRequest) Execute() (*PostgresFirewallRule, *http.Response, error) {
	return r.ApiService.GetLocationPostgresFirewallRuleDetailsWithIdExecute(r)
}

/*
GetLocationPostgresFirewallRuleDetailsWithId Get details of a Postgres firewall rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param location The Ubicloud location/region
 @param postgresDatabaseId Postgres database ID
 @param projectId ID of the project
 @param firewallRuleId ID of the firewall rule
 @return ApiGetLocationPostgresFirewallRuleDetailsWithIdRequest
*/
func (a *PostgresFirewallRuleApiService) GetLocationPostgresFirewallRuleDetailsWithId(ctx context.Context, location string, postgresDatabaseId string, projectId string, firewallRuleId string) ApiGetLocationPostgresFirewallRuleDetailsWithIdRequest {
	return ApiGetLocationPostgresFirewallRuleDetailsWithIdRequest{
		ApiService: a,
		ctx: ctx,
		location: location,
		postgresDatabaseId: postgresDatabaseId,
		projectId: projectId,
		firewallRuleId: firewallRuleId,
	}
}

// Execute executes the request
//  @return PostgresFirewallRule
func (a *PostgresFirewallRuleApiService) GetLocationPostgresFirewallRuleDetailsWithIdExecute(r ApiGetLocationPostgresFirewallRuleDetailsWithIdRequest) (*PostgresFirewallRule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresFirewallRule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresFirewallRuleApiService.GetLocationPostgresFirewallRuleDetailsWithId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/firewall-rule/{firewall_rule_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"location"+"}", url.PathEscape(parameterValueToString(r.location, "location")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"postgres_database_id"+"}", url.PathEscape(parameterValueToString(r.postgresDatabaseId, "postgresDatabaseId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"firewall_rule_id"+"}", url.PathEscape(parameterValueToString(r.firewallRuleId, "firewallRuleId")), -1)

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

type ApiListLocationPostgresFirewallRulesRequest struct {
	ctx context.Context
	ApiService *PostgresFirewallRuleApiService
	location string
	postgresDatabaseId string
	projectId string
}

func (r ApiListLocationPostgresFirewallRulesRequest) Execute() (*ListLocationPostgresFirewallRules200Response, *http.Response, error) {
	return r.ApiService.ListLocationPostgresFirewallRulesExecute(r)
}

/*
ListLocationPostgresFirewallRules List location Postgres firewall rules

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param location The Ubicloud location/region
 @param postgresDatabaseId Postgres database ID
 @param projectId ID of the project
 @return ApiListLocationPostgresFirewallRulesRequest
*/
func (a *PostgresFirewallRuleApiService) ListLocationPostgresFirewallRules(ctx context.Context, location string, postgresDatabaseId string, projectId string) ApiListLocationPostgresFirewallRulesRequest {
	return ApiListLocationPostgresFirewallRulesRequest{
		ApiService: a,
		ctx: ctx,
		location: location,
		postgresDatabaseId: postgresDatabaseId,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return ListLocationPostgresFirewallRules200Response
func (a *PostgresFirewallRuleApiService) ListLocationPostgresFirewallRulesExecute(r ApiListLocationPostgresFirewallRulesRequest) (*ListLocationPostgresFirewallRules200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListLocationPostgresFirewallRules200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresFirewallRuleApiService.ListLocationPostgresFirewallRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/project/{project_id}/location/{location}/postgres/_{postgres_database_id}/firewall-rule"
	localVarPath = strings.Replace(localVarPath, "{"+"location"+"}", url.PathEscape(parameterValueToString(r.location, "location")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"postgres_database_id"+"}", url.PathEscape(parameterValueToString(r.postgresDatabaseId, "postgresDatabaseId")), -1)
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