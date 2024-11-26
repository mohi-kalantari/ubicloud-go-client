# \FirewallApi

All URIs are relative to *https://api.ubicloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionLocationFirewallAttachSubnet**](FirewallApi.md#ActionLocationFirewallAttachSubnet) | **Post** /project/{project_id}/location/{location}/firewall/_{firewall_id}/attach-subnet | Attach a subnet to firewall
[**ActionLocationFirewallDetachSubnet**](FirewallApi.md#ActionLocationFirewallDetachSubnet) | **Post** /project/{project_id}/location/{location}/firewall/_{firewall_id}/detach-subnet | Detach a subnet from firewall
[**CreateFirewall**](FirewallApi.md#CreateFirewall) | **Post** /project/{project_id}/firewall | Create a new firewall
[**CreateLocationFirewall**](FirewallApi.md#CreateLocationFirewall) | **Post** /project/{project_id}/location/{location}/firewall/{firewall_name} | Create a new firewall
[**DeleteFirewall**](FirewallApi.md#DeleteFirewall) | **Delete** /project/{project_id}/firewall/{firewall_name} | Delete a specific firewall
[**DeleteLocationFirewall**](FirewallApi.md#DeleteLocationFirewall) | **Delete** /project/{project_id}/location/{location}/firewall/{firewall_name} | Delete a specific firewall
[**DeleteLocationFirewallWithId**](FirewallApi.md#DeleteLocationFirewallWithId) | **Delete** /project/{project_id}/location/{location}/firewall/id/{firewall_id} | Delete a specific firewall
[**GetFirewall**](FirewallApi.md#GetFirewall) | **Get** /project/{project_id}/firewall | Return the list of firewalls in the project
[**GetFirewallDetails**](FirewallApi.md#GetFirewallDetails) | **Get** /project/{project_id}/firewall/{firewall_name} | Get details of a specific firewall
[**GetLocationFirewall**](FirewallApi.md#GetLocationFirewall) | **Get** /project/{project_id}/location/{location}/firewall | Return the list of firewalls in the project and location
[**GetLocationFirewallDetails**](FirewallApi.md#GetLocationFirewallDetails) | **Get** /project/{project_id}/location/{location}/firewall/{firewall_name} | Get details of a specific firewall
[**GetLocationFirewallDetailsWithId**](FirewallApi.md#GetLocationFirewallDetailsWithId) | **Get** /project/{project_id}/location/{location}/firewall/id/{firewall_id} | Get details of a specific firewall



## ActionLocationFirewallAttachSubnet

> FirewallDetailed ActionLocationFirewallAttachSubnet(ctx, location, projectId, firewallId).ActionLocationFirewallAttachSubnetRequest(actionLocationFirewallAttachSubnetRequest).Execute()

Attach a subnet to firewall

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
    firewallId := "firewallId_example" // string | ID of the firewall
    actionLocationFirewallAttachSubnetRequest := *openapiclient.NewActionLocationFirewallAttachSubnetRequest("PrivateSubnetId_example") // ActionLocationFirewallAttachSubnetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.ActionLocationFirewallAttachSubnet(context.Background(), location, projectId, firewallId).ActionLocationFirewallAttachSubnetRequest(actionLocationFirewallAttachSubnetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.ActionLocationFirewallAttachSubnet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionLocationFirewallAttachSubnet`: FirewallDetailed
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.ActionLocationFirewallAttachSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 
**firewallId** | **string** | ID of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionLocationFirewallAttachSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **actionLocationFirewallAttachSubnetRequest** | [**ActionLocationFirewallAttachSubnetRequest**](ActionLocationFirewallAttachSubnetRequest.md) |  | 

### Return type

[**FirewallDetailed**](FirewallDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActionLocationFirewallDetachSubnet

> FirewallDetailed ActionLocationFirewallDetachSubnet(ctx, location, projectId, firewallId).ActionLocationFirewallDetachSubnetRequest(actionLocationFirewallDetachSubnetRequest).Execute()

Detach a subnet from firewall

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
    firewallId := "firewallId_example" // string | ID of the firewall
    actionLocationFirewallDetachSubnetRequest := *openapiclient.NewActionLocationFirewallDetachSubnetRequest("PrivateSubnetId_example") // ActionLocationFirewallDetachSubnetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.ActionLocationFirewallDetachSubnet(context.Background(), location, projectId, firewallId).ActionLocationFirewallDetachSubnetRequest(actionLocationFirewallDetachSubnetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.ActionLocationFirewallDetachSubnet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActionLocationFirewallDetachSubnet`: FirewallDetailed
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.ActionLocationFirewallDetachSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 
**firewallId** | **string** | ID of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionLocationFirewallDetachSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **actionLocationFirewallDetachSubnetRequest** | [**ActionLocationFirewallDetachSubnetRequest**](ActionLocationFirewallDetachSubnetRequest.md) |  | 

### Return type

[**FirewallDetailed**](FirewallDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFirewall

> Firewall CreateFirewall(ctx, projectId).CreateFirewallRequest(createFirewallRequest).Execute()

Create a new firewall

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
    createFirewallRequest := *openapiclient.NewCreateFirewallRequest("Name_example") // CreateFirewallRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.CreateFirewall(context.Background(), projectId).CreateFirewallRequest(createFirewallRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.CreateFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFirewall`: Firewall
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.CreateFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFirewallRequest** | [**CreateFirewallRequest**](CreateFirewallRequest.md) |  | 

### Return type

[**Firewall**](Firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocationFirewall

> Firewall CreateLocationFirewall(ctx, location, projectId, firewallName).CreateLocationFirewallRequest(createLocationFirewallRequest).Execute()

Create a new firewall

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
    createLocationFirewallRequest := *openapiclient.NewCreateLocationFirewallRequest() // CreateLocationFirewallRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.CreateLocationFirewall(context.Background(), location, projectId, firewallName).CreateLocationFirewallRequest(createLocationFirewallRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.CreateLocationFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationFirewall`: Firewall
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.CreateLocationFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createLocationFirewallRequest** | [**CreateLocationFirewallRequest**](CreateLocationFirewallRequest.md) |  | 

### Return type

[**Firewall**](Firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFirewall

> DeleteFirewall(ctx, projectId, firewallName).Execute()

Delete a specific firewall

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallApi.DeleteFirewall(context.Background(), projectId, firewallName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewall``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRequest struct via the builder pattern


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


## DeleteLocationFirewall

> DeleteLocationFirewall(ctx, location, projectId, firewallName).Execute()

Delete a specific firewall

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallApi.DeleteLocationFirewall(context.Background(), location, projectId, firewallName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteLocationFirewall``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocationFirewallRequest struct via the builder pattern


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


## DeleteLocationFirewallWithId

> DeleteLocationFirewallWithId(ctx, location, projectId, firewallId).Execute()

Delete a specific firewall

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
    firewallId := "firewallId_example" // string | ID of the firewall

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallApi.DeleteLocationFirewallWithId(context.Background(), location, projectId, firewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteLocationFirewallWithId``: %v\n", err)
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
**firewallId** | **string** | ID of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocationFirewallWithIdRequest struct via the builder pattern


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


## GetFirewall

> GetFirewall200Response GetFirewall(ctx, projectId).Execute()

Return the list of firewalls in the project

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetFirewall(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewall`: GetFirewall200Response
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFirewall200Response**](GetFirewall200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallDetails

> Firewall GetFirewallDetails(ctx, projectId, firewallName).Execute()

Get details of a specific firewall

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetFirewallDetails(context.Background(), projectId, firewallName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallDetails`: Firewall
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Firewall**](Firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationFirewall

> GetFirewall200Response GetLocationFirewall(ctx, location, projectId).Execute()

Return the list of firewalls in the project and location

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetLocationFirewall(context.Background(), location, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetLocationFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationFirewall`: GetFirewall200Response
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetLocationFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetFirewall200Response**](GetFirewall200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationFirewallDetails

> FirewallDetailed GetLocationFirewallDetails(ctx, location, projectId, firewallName).Execute()

Get details of a specific firewall

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetLocationFirewallDetails(context.Background(), location, projectId, firewallName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetLocationFirewallDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationFirewallDetails`: FirewallDetailed
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetLocationFirewallDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 
**firewallName** | **string** | Name of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationFirewallDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FirewallDetailed**](FirewallDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationFirewallDetailsWithId

> Firewall GetLocationFirewallDetailsWithId(ctx, location, projectId, firewallId).Execute()

Get details of a specific firewall

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
    firewallId := "firewallId_example" // string | ID of the firewall

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetLocationFirewallDetailsWithId(context.Background(), location, projectId, firewallId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetLocationFirewallDetailsWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationFirewallDetailsWithId`: Firewall
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetLocationFirewallDetailsWithId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 
**firewallId** | **string** | ID of the firewall | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationFirewallDetailsWithIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Firewall**](Firewall.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

