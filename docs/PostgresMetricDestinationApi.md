# \PostgresMetricDestinationApi

All URIs are relative to *https://api.ubicloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLocationPostgresMetricDestination**](PostgresMetricDestinationApi.md#CreateLocationPostgresMetricDestination) | **Post** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/metric-destination | Create a new Postgres Metric Destination
[**DeleteLocationPostgresMetricDestination**](PostgresMetricDestinationApi.md#DeleteLocationPostgresMetricDestination) | **Delete** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/metric-destination/{metric_destination_id} | Delete a specific Metric Destination



## CreateLocationPostgresMetricDestination

> PostgresDetailed CreateLocationPostgresMetricDestination(ctx, location, postgresDatabaseName, projectId).CreateLocationPostgresMetricDestinationRequest(createLocationPostgresMetricDestinationRequest).Execute()

Create a new Postgres Metric Destination

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mohi-kalantari/ubicloud-go-client"
)

func main() {
    location := "eu-central-h1" // string | The Ubicloud location/region
    postgresDatabaseName := "postgresDatabaseName_example" // string | Postgres database name
    projectId := "projectId_example" // string | ID of the project
    createLocationPostgresMetricDestinationRequest := *openapiclient.NewCreateLocationPostgresMetricDestinationRequest("Password_example", "Url_example", "Username_example") // CreateLocationPostgresMetricDestinationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresMetricDestinationApi.CreateLocationPostgresMetricDestination(context.Background(), location, postgresDatabaseName, projectId).CreateLocationPostgresMetricDestinationRequest(createLocationPostgresMetricDestinationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresMetricDestinationApi.CreateLocationPostgresMetricDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationPostgresMetricDestination`: PostgresDetailed
    fmt.Fprintf(os.Stdout, "Response from `PostgresMetricDestinationApi.CreateLocationPostgresMetricDestination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseName** | **string** | Postgres database name | 
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationPostgresMetricDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createLocationPostgresMetricDestinationRequest** | [**CreateLocationPostgresMetricDestinationRequest**](CreateLocationPostgresMetricDestinationRequest.md) |  | 

### Return type

[**PostgresDetailed**](PostgresDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLocationPostgresMetricDestination

> DeleteLocationPostgresMetricDestination(ctx, location, metricDestinationId, postgresDatabaseName, projectId).Execute()

Delete a specific Metric Destination

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mohi-kalantari/ubicloud-go-client"
)

func main() {
    location := "eu-central-h1" // string | The Ubicloud location/region
    metricDestinationId := "metricDestinationId_example" // string | Postgres Metric Destination ID
    postgresDatabaseName := "postgresDatabaseName_example" // string | Postgres database name
    projectId := "projectId_example" // string | ID of the project

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PostgresMetricDestinationApi.DeleteLocationPostgresMetricDestination(context.Background(), location, metricDestinationId, postgresDatabaseName, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresMetricDestinationApi.DeleteLocationPostgresMetricDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**metricDestinationId** | **string** | Postgres Metric Destination ID | 
**postgresDatabaseName** | **string** | Postgres database name | 
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocationPostgresMetricDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

