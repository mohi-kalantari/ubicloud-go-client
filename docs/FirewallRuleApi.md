# \FirewallRuleApi

All URIs are relative to *https://api.ubicloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFirewallRule**](FirewallRuleApi.md#CreateFirewallRule) | **Post** /project/{project_id}/firewall/{firewall_name}/firewall-rule | Create a new firewall rule
[**CreateLocationFirewallFirewallRule**](FirewallRuleApi.md#CreateLocationFirewallFirewallRule) | **Post** /project/{project_id}/location/{location}/firewall/{firewall_name}/firewall-rule/{firewall_rule_id} | Create a new firewall rule
[**CreateLocationFirewallFirewallRuleWithId**](FirewallRuleApi.md#CreateLocationFirewallFirewallRuleWithId) | **Post** /project/{project_id}/location/{location}/firewall/id/{firewall_name}/firewall-rule/{firewall_rule_id} | Create a new firewall rule
[**CreateLocationFirewallRule**](FirewallRuleApi.md#CreateLocationFirewallRule) | **Post** /project/{project_id}/location/{location}/firewall/{firewall_name}/firewall-rule | Create a new firewall rule
[**CreateLocationFirewallRuleWithId**](FirewallRuleApi.md#CreateLocationFirewallRuleWithId) | **Post** /project/{project_id}/location/{location}/firewall/id/{firewall_name}/firewall-rule | Create a new firewall rule
[**DeleteFirewallRule**](FirewallRuleApi.md#DeleteFirewallRule) | **Delete** /project/{project_id}/firewall/{firewall_name}/firewall-rule/{firewall_rule_id} | Delete a specific firewall rule
[**DeleteLocationFirewallFirewallRule**](FirewallRuleApi.md#DeleteLocationFirewallFirewallRule) | **Delete** /project/{project_id}/location/{location}/firewall/{firewall_name}/firewall-rule/{firewall_rule_id} | Delete a specific firewall rule
[**DeleteLocationFirewallFirewallRuleWithId**](FirewallRuleApi.md#DeleteLocationFirewallFirewallRuleWithId) | **Delete** /project/{project_id}/location/{location}/firewall/id/{firewall_name}/firewall-rule/{firewall_rule_id} | Delete a specific firewall rule
[**DeleteLocationPostgresFirewallRule**](FirewallRuleApi.md#DeleteLocationPostgresFirewallRule) | **Delete** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/firewall-rule/{firewall_rule_id} | Delete a specific firewall rule
[**DeleteLocationPostgresFirewallRuleWithId**](FirewallRuleApi.md#DeleteLocationPostgresFirewallRuleWithId) | **Delete** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/firewall-rule/{firewall_rule_id} | Delete a specific Postgres firewall rule
[**GetFirewallRuleDetails**](FirewallRuleApi.md#GetFirewallRuleDetails) | **Get** /project/{project_id}/firewall/{firewall_name}/firewall-rule/{firewall_rule_id} | Get details of a firewall rule
[**GetLocationFirewallFirewallRuleDetails**](FirewallRuleApi.md#GetLocationFirewallFirewallRuleDetails) | **Get** /project/{project_id}/location/{location}/firewall/{firewall_name}/firewall-rule/{firewall_rule_id} | Get details of a firewall rule
[**GetLocationFirewallFirewallRuleDetailsWithId**](FirewallRuleApi.md#GetLocationFirewallFirewallRuleDetailsWithId) | **Get** /project/{project_id}/location/{location}/firewall/id/{firewall_name}/firewall-rule/{firewall_rule_id} | Get details of a firewall rule



## CreateFirewallRule

> FirewallRule CreateFirewallRule(ctx, projectId, firewallName).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()

Create a new firewall rule

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
    projectId := "projectId_example" // string | ID of the project
    firewallName := "firewallName_example" // string | Name of the firewall
    createFirewallRuleRequest := *openapiclient.NewCreateFirewallRuleRequest("Cidr_example") // CreateFirewallRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRuleApi.CreateFirewallRule(context.Background(), projectId, firewallName).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.CreateFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFirewallRule`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRuleApi.CreateFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createFirewallRuleRequest** | [**CreateFirewallRuleRequest**](CreateFirewallRuleRequest.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocationFirewallFirewallRule

> FirewallRule CreateLocationFirewallFirewallRule(ctx, location, projectId, firewallName, firewallRuleId).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()

Create a new firewall rule

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
    projectId := "projectId_example" // string | ID of the project
    firewallName := "firewallName_example" // string | Name of the firewall
    firewallRuleId := "firewallRuleId_example" // string | ID of the firewall rule
    createFirewallRuleRequest := *openapiclient.NewCreateFirewallRuleRequest("Cidr_example") // CreateFirewallRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRuleApi.CreateLocationFirewallFirewallRule(context.Background(), location, projectId, firewallName, firewallRuleId).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.CreateLocationFirewallFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationFirewallFirewallRule`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRuleApi.CreateLocationFirewallFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 
**firewallRuleId** | **string** | ID of the firewall rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationFirewallFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **createFirewallRuleRequest** | [**CreateFirewallRuleRequest**](CreateFirewallRuleRequest.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocationFirewallFirewallRuleWithId

> FirewallRule CreateLocationFirewallFirewallRuleWithId(ctx, location, projectId, firewallName, firewallRuleId).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()

Create a new firewall rule

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
    projectId := "projectId_example" // string | ID of the project
    firewallName := "firewallName_example" // string | Name of the firewall
    firewallRuleId := "firewallRuleId_example" // string | ID of the firewall rule
    createFirewallRuleRequest := *openapiclient.NewCreateFirewallRuleRequest("Cidr_example") // CreateFirewallRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRuleApi.CreateLocationFirewallFirewallRuleWithId(context.Background(), location, projectId, firewallName, firewallRuleId).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.CreateLocationFirewallFirewallRuleWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationFirewallFirewallRuleWithId`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRuleApi.CreateLocationFirewallFirewallRuleWithId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 
**firewallRuleId** | **string** | ID of the firewall rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationFirewallFirewallRuleWithIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **createFirewallRuleRequest** | [**CreateFirewallRuleRequest**](CreateFirewallRuleRequest.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocationFirewallRule

> FirewallRule CreateLocationFirewallRule(ctx, firewallName, location, projectId).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()

Create a new firewall rule

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
    firewallName := "firewallName_example" // string | Name of the firewall
    location := "eu-central-h1" // string | The Ubicloud location/region
    projectId := "projectId_example" // string | ID of the project
    createFirewallRuleRequest := *openapiclient.NewCreateFirewallRuleRequest("Cidr_example") // CreateFirewallRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRuleApi.CreateLocationFirewallRule(context.Background(), firewallName, location, projectId).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.CreateLocationFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationFirewallRule`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRuleApi.CreateLocationFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallName** | **string** | Name of the firewall | 
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createFirewallRuleRequest** | [**CreateFirewallRuleRequest**](CreateFirewallRuleRequest.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocationFirewallRuleWithId

> FirewallRule CreateLocationFirewallRuleWithId(ctx, firewallName, location, projectId).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()

Create a new firewall rule

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
    firewallName := "firewallName_example" // string | Name of the firewall
    location := "eu-central-h1" // string | The Ubicloud location/region
    projectId := "projectId_example" // string | ID of the project
    createFirewallRuleRequest := *openapiclient.NewCreateFirewallRuleRequest("Cidr_example") // CreateFirewallRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRuleApi.CreateLocationFirewallRuleWithId(context.Background(), firewallName, location, projectId).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.CreateLocationFirewallRuleWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationFirewallRuleWithId`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRuleApi.CreateLocationFirewallRuleWithId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firewallName** | **string** | Name of the firewall | 
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationFirewallRuleWithIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createFirewallRuleRequest** | [**CreateFirewallRuleRequest**](CreateFirewallRuleRequest.md) |  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewallRule

> DeleteFirewallRule(ctx, projectId, firewallName, firewallRuleId).Execute()

Delete a specific firewall rule

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
    projectId := "projectId_example" // string | ID of the project
    firewallName := "firewallName_example" // string | Name of the firewall
    firewallRuleId := "firewallRuleId_example" // string | ID of the firewall rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallRuleApi.DeleteFirewallRule(context.Background(), projectId, firewallName, firewallRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.DeleteFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 
**firewallRuleId** | **string** | ID of the firewall rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleRequest struct via the builder pattern


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


## DeleteLocationFirewallFirewallRule

> DeleteLocationFirewallFirewallRule(ctx, location, projectId, firewallName, firewallRuleId).Execute()

Delete a specific firewall rule

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
    projectId := "projectId_example" // string | ID of the project
    firewallName := "firewallName_example" // string | Name of the firewall
    firewallRuleId := "firewallRuleId_example" // string | ID of the firewall rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallRuleApi.DeleteLocationFirewallFirewallRule(context.Background(), location, projectId, firewallName, firewallRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.DeleteLocationFirewallFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 
**firewallRuleId** | **string** | ID of the firewall rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocationFirewallFirewallRuleRequest struct via the builder pattern


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


## DeleteLocationFirewallFirewallRuleWithId

> DeleteLocationFirewallFirewallRuleWithId(ctx, location, projectId, firewallName, firewallRuleId).Execute()

Delete a specific firewall rule

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
    projectId := "projectId_example" // string | ID of the project
    firewallName := "firewallName_example" // string | Name of the firewall
    firewallRuleId := "firewallRuleId_example" // string | ID of the firewall rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallRuleApi.DeleteLocationFirewallFirewallRuleWithId(context.Background(), location, projectId, firewallName, firewallRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.DeleteLocationFirewallFirewallRuleWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 
**firewallRuleId** | **string** | ID of the firewall rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocationFirewallFirewallRuleWithIdRequest struct via the builder pattern


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


## DeleteLocationPostgresFirewallRule

> DeleteLocationPostgresFirewallRule(ctx, projectId, location, postgresDatabaseName, firewallRuleId).Execute()

Delete a specific firewall rule

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
    projectId := "projectId_example" // string | ID of the project
    location := "eu-central-h1" // string | The Ubicloud location/region
    postgresDatabaseName := "postgresDatabaseName_example" // string | Postgres database name
    firewallRuleId := "firewallRuleId_example" // string | ID of the firewall rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallRuleApi.DeleteLocationPostgresFirewallRule(context.Background(), projectId, location, postgresDatabaseName, firewallRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.DeleteLocationPostgresFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseName** | **string** | Postgres database name | 
**firewallRuleId** | **string** | ID of the firewall rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocationPostgresFirewallRuleRequest struct via the builder pattern


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


## DeleteLocationPostgresFirewallRuleWithId

> DeleteLocationPostgresFirewallRuleWithId(ctx, location, postgresDatabaseId, projectId, firewallRuleId).Execute()

Delete a specific Postgres firewall rule

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
    r, err := apiClient.FirewallRuleApi.DeleteLocationPostgresFirewallRuleWithId(context.Background(), location, postgresDatabaseId, projectId, firewallRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.DeleteLocationPostgresFirewallRuleWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
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

Other parameters are passed through a pointer to a apiDeleteLocationPostgresFirewallRuleWithIdRequest struct via the builder pattern


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


## GetFirewallRuleDetails

> FirewallRule GetFirewallRuleDetails(ctx, projectId, firewallName, firewallRuleId).Execute()

Get details of a firewall rule

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
    projectId := "projectId_example" // string | ID of the project
    firewallName := "firewallName_example" // string | Name of the firewall
    firewallRuleId := "firewallRuleId_example" // string | ID of the firewall rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRuleApi.GetFirewallRuleDetails(context.Background(), projectId, firewallName, firewallRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.GetFirewallRuleDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallRuleDetails`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRuleApi.GetFirewallRuleDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 
**firewallRuleId** | **string** | ID of the firewall rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRuleDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationFirewallFirewallRuleDetails

> FirewallRule GetLocationFirewallFirewallRuleDetails(ctx, location, projectId, firewallName, firewallRuleId).Execute()

Get details of a firewall rule

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
    projectId := "projectId_example" // string | ID of the project
    firewallName := "firewallName_example" // string | Name of the firewall
    firewallRuleId := "firewallRuleId_example" // string | ID of the firewall rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRuleApi.GetLocationFirewallFirewallRuleDetails(context.Background(), location, projectId, firewallName, firewallRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.GetLocationFirewallFirewallRuleDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationFirewallFirewallRuleDetails`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRuleApi.GetLocationFirewallFirewallRuleDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 
**firewallRuleId** | **string** | ID of the firewall rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationFirewallFirewallRuleDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationFirewallFirewallRuleDetailsWithId

> FirewallRule GetLocationFirewallFirewallRuleDetailsWithId(ctx, location, projectId, firewallName, firewallRuleId).Execute()

Get details of a firewall rule

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
    projectId := "projectId_example" // string | ID of the project
    firewallName := "firewallName_example" // string | Name of the firewall
    firewallRuleId := "firewallRuleId_example" // string | ID of the firewall rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallRuleApi.GetLocationFirewallFirewallRuleDetailsWithId(context.Background(), location, projectId, firewallName, firewallRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallRuleApi.GetLocationFirewallFirewallRuleDetailsWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationFirewallFirewallRuleDetailsWithId`: FirewallRule
    fmt.Fprintf(os.Stdout, "Response from `FirewallRuleApi.GetLocationFirewallFirewallRuleDetailsWithId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 
**firewallRuleId** | **string** | ID of the firewall rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationFirewallFirewallRuleDetailsWithIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

