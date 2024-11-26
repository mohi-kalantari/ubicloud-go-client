# \PostgresFirewallRuleApi

All URIs are relative to *https://api.ubicloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLocationPostgresFirewallRule**](PostgresFirewallRuleApi.md#CreateLocationPostgresFirewallRule) | **Post** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/firewall-rule | Create a new postgres firewall rule
[**CreateLocationPostgresFirewallRuleWithId**](PostgresFirewallRuleApi.md#CreateLocationPostgresFirewallRuleWithId) | **Post** /project/{project_id}/location/{location}/postgres/_{postgres_database_id}/firewall-rule | Create a new Postgres firewall rule
[**CreateLocationPostgresFirewallRuleWithIdWithId**](PostgresFirewallRuleApi.md#CreateLocationPostgresFirewallRuleWithIdWithId) | **Post** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/firewall-rule/{firewall_rule_id} | Create a new Postgres firewall rule
[**GetLocationPostgresFirewallRuleDetailsWithId**](PostgresFirewallRuleApi.md#GetLocationPostgresFirewallRuleDetailsWithId) | **Get** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/firewall-rule/{firewall_rule_id} | Get details of a Postgres firewall rule
[**ListLocationPostgresFirewallRules**](PostgresFirewallRuleApi.md#ListLocationPostgresFirewallRules) | **Get** /project/{project_id}/location/{location}/postgres/_{postgres_database_id}/firewall-rule | List location Postgres firewall rules



## CreateLocationPostgresFirewallRule

> PostgresFirewallRule CreateLocationPostgresFirewallRule(ctx, location, postgresDatabaseName, projectId).CreateLocationPostgresFirewallRuleWithIdRequest(createLocationPostgresFirewallRuleWithIdRequest).Execute()

Create a new postgres firewall rule

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
    createLocationPostgresFirewallRuleWithIdRequest := *openapiclient.NewCreateLocationPostgresFirewallRuleWithIdRequest("Cidr_example") // CreateLocationPostgresFirewallRuleWithIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresFirewallRuleApi.CreateLocationPostgresFirewallRule(context.Background(), location, postgresDatabaseName, projectId).CreateLocationPostgresFirewallRuleWithIdRequest(createLocationPostgresFirewallRuleWithIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresFirewallRuleApi.CreateLocationPostgresFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationPostgresFirewallRule`: PostgresFirewallRule
    fmt.Fprintf(os.Stdout, "Response from `PostgresFirewallRuleApi.CreateLocationPostgresFirewallRule`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateLocationPostgresFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createLocationPostgresFirewallRuleWithIdRequest** | [**CreateLocationPostgresFirewallRuleWithIdRequest**](CreateLocationPostgresFirewallRuleWithIdRequest.md) |  | 

### Return type

[**PostgresFirewallRule**](PostgresFirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocationPostgresFirewallRuleWithId

> PostgresFirewallRule CreateLocationPostgresFirewallRuleWithId(ctx, location, postgresDatabaseId, projectId).CreateLocationPostgresFirewallRuleWithIdRequest(createLocationPostgresFirewallRuleWithIdRequest).Execute()

Create a new Postgres firewall rule

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
    postgresDatabaseId := "postgresDatabaseId_example" // string | Postgres database ID
    projectId := "projectId_example" // string | ID of the project
    createLocationPostgresFirewallRuleWithIdRequest := *openapiclient.NewCreateLocationPostgresFirewallRuleWithIdRequest("Cidr_example") // CreateLocationPostgresFirewallRuleWithIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresFirewallRuleApi.CreateLocationPostgresFirewallRuleWithId(context.Background(), location, postgresDatabaseId, projectId).CreateLocationPostgresFirewallRuleWithIdRequest(createLocationPostgresFirewallRuleWithIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresFirewallRuleApi.CreateLocationPostgresFirewallRuleWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationPostgresFirewallRuleWithId`: PostgresFirewallRule
    fmt.Fprintf(os.Stdout, "Response from `PostgresFirewallRuleApi.CreateLocationPostgresFirewallRuleWithId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseId** | **string** | Postgres database ID | 
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationPostgresFirewallRuleWithIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createLocationPostgresFirewallRuleWithIdRequest** | [**CreateLocationPostgresFirewallRuleWithIdRequest**](CreateLocationPostgresFirewallRuleWithIdRequest.md) |  | 

### Return type

[**PostgresFirewallRule**](PostgresFirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocationPostgresFirewallRuleWithIdWithId

> PostgresFirewallRule CreateLocationPostgresFirewallRuleWithIdWithId(ctx, location, postgresDatabaseId, projectId, firewallRuleId).CreateLocationPostgresFirewallRuleWithIdWithIdRequest(createLocationPostgresFirewallRuleWithIdWithIdRequest).Execute()

Create a new Postgres firewall rule

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
    postgresDatabaseId := "postgresDatabaseId_example" // string | Postgres database ID
    projectId := "projectId_example" // string | ID of the project
    firewallRuleId := "firewallRuleId_example" // string | ID of the firewall rule
    createLocationPostgresFirewallRuleWithIdWithIdRequest := *openapiclient.NewCreateLocationPostgresFirewallRuleWithIdWithIdRequest("Cidr_example") // CreateLocationPostgresFirewallRuleWithIdWithIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresFirewallRuleApi.CreateLocationPostgresFirewallRuleWithIdWithId(context.Background(), location, postgresDatabaseId, projectId, firewallRuleId).CreateLocationPostgresFirewallRuleWithIdWithIdRequest(createLocationPostgresFirewallRuleWithIdWithIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresFirewallRuleApi.CreateLocationPostgresFirewallRuleWithIdWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationPostgresFirewallRuleWithIdWithId`: PostgresFirewallRule
    fmt.Fprintf(os.Stdout, "Response from `PostgresFirewallRuleApi.CreateLocationPostgresFirewallRuleWithIdWithId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseId** | **string** | Postgres database ID | 
**projectId** | **string** | ID of the project | 
**firewallRuleId** | **string** | ID of the firewall rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationPostgresFirewallRuleWithIdWithIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **createLocationPostgresFirewallRuleWithIdWithIdRequest** | [**CreateLocationPostgresFirewallRuleWithIdWithIdRequest**](CreateLocationPostgresFirewallRuleWithIdWithIdRequest.md) |  | 

### Return type

[**PostgresFirewallRule**](PostgresFirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationPostgresFirewallRuleDetailsWithId

> PostgresFirewallRule GetLocationPostgresFirewallRuleDetailsWithId(ctx, location, postgresDatabaseId, projectId, firewallRuleId).Execute()

Get details of a Postgres firewall rule

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
    postgresDatabaseId := "postgresDatabaseId_example" // string | Postgres database ID
    projectId := "projectId_example" // string | ID of the project
    firewallRuleId := "firewallRuleId_example" // string | ID of the firewall rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresFirewallRuleApi.GetLocationPostgresFirewallRuleDetailsWithId(context.Background(), location, postgresDatabaseId, projectId, firewallRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresFirewallRuleApi.GetLocationPostgresFirewallRuleDetailsWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationPostgresFirewallRuleDetailsWithId`: PostgresFirewallRule
    fmt.Fprintf(os.Stdout, "Response from `PostgresFirewallRuleApi.GetLocationPostgresFirewallRuleDetailsWithId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseId** | **string** | Postgres database ID | 
**projectId** | **string** | ID of the project | 
**firewallRuleId** | **string** | ID of the firewall rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationPostgresFirewallRuleDetailsWithIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PostgresFirewallRule**](PostgresFirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocationPostgresFirewallRules

> ListLocationPostgresFirewallRules200Response ListLocationPostgresFirewallRules(ctx, location, postgresDatabaseId, projectId).Execute()

List location Postgres firewall rules

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
    postgresDatabaseId := "postgresDatabaseId_example" // string | Postgres database ID
    projectId := "projectId_example" // string | ID of the project

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresFirewallRuleApi.ListLocationPostgresFirewallRules(context.Background(), location, postgresDatabaseId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresFirewallRuleApi.ListLocationPostgresFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocationPostgresFirewallRules`: ListLocationPostgresFirewallRules200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresFirewallRuleApi.ListLocationPostgresFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseId** | **string** | Postgres database ID | 
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocationPostgresFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ListLocationPostgresFirewallRules200Response**](ListLocationPostgresFirewallRules200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

