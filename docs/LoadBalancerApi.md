# \LoadBalancerApi

All URIs are relative to *https://api.ubicloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachVmLocationLoadBalancer**](LoadBalancerApi.md#AttachVmLocationLoadBalancer) | **Post** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name}/attach-vm | Attach a VM to a Load Balancer in a specific location of a project
[**CreateLoadBalancer**](LoadBalancerApi.md#CreateLoadBalancer) | **Post** /project/{project_id}/load-balancer/{load_balancer_name} | Create a new Load Balancer in a project
[**CreateLocationLoadBalancer**](LoadBalancerApi.md#CreateLocationLoadBalancer) | **Post** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name} | Create a new Load Balancer in a specific location of a project
[**DeleteLoadBalancer**](LoadBalancerApi.md#DeleteLoadBalancer) | **Delete** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name} | Delete a specific Load Balancer
[**DeleteLoadBalancerWithID**](LoadBalancerApi.md#DeleteLoadBalancerWithID) | **Delete** /project/{project_id}/location/{location}/load-balancer/id/{load_balancer_id} | Delete a specific Load Balancer with ID
[**DetachVmLocationLoadBalancer**](LoadBalancerApi.md#DetachVmLocationLoadBalancer) | **Post** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name}/detach-vm | Detach a VM from a Load Balancer in a specific location of a project
[**GetLoadBalancer**](LoadBalancerApi.md#GetLoadBalancer) | **Get** /project/{project_id}/load-balancer/{load_balancer_name} | Get details of a specific Load Balancer
[**GetLoadBalancerDetails**](LoadBalancerApi.md#GetLoadBalancerDetails) | **Get** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name} | Get details of a specific Load Balancer in a location
[**GetLoadBalancerDetailsWithId**](LoadBalancerApi.md#GetLoadBalancerDetailsWithId) | **Get** /project/{project_id}/location/{location}/load-balancer/id/{load_balancer_id} | Get details of a specific Load Balancer in a location with ID
[**ListLoadBalancers**](LoadBalancerApi.md#ListLoadBalancers) | **Get** /project/{project_id}/load-balancer | List Load Balancers in a specific project
[**ListLocationLoadBalancers**](LoadBalancerApi.md#ListLocationLoadBalancers) | **Get** /project/{project_id}/location/{location}/load-balancer | List Load Balancers in a specific location of a project
[**PatchLocationLoadBalancer**](LoadBalancerApi.md#PatchLocationLoadBalancer) | **Patch** /project/{project_id}/location/{location}/load-balancer/{load_balancer_name} | Update a Load Balancer in a specific location of a project



## AttachVmLocationLoadBalancer

> LoadBalancerDetailed AttachVmLocationLoadBalancer(ctx, projectId, location, loadBalancerName).AttachVmLocationLoadBalancerRequest(attachVmLocationLoadBalancerRequest).Execute()

Attach a VM to a Load Balancer in a specific location of a project

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
    loadBalancerName := "loadBalancerName_example" // string | Name of the load balancer
    attachVmLocationLoadBalancerRequest := *openapiclient.NewAttachVmLocationLoadBalancerRequest("VmId_example") // AttachVmLocationLoadBalancerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancerApi.AttachVmLocationLoadBalancer(context.Background(), projectId, location, loadBalancerName).AttachVmLocationLoadBalancerRequest(attachVmLocationLoadBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.AttachVmLocationLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachVmLocationLoadBalancer`: LoadBalancerDetailed
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.AttachVmLocationLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**loadBalancerName** | **string** | Name of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachVmLocationLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **attachVmLocationLoadBalancerRequest** | [**AttachVmLocationLoadBalancerRequest**](AttachVmLocationLoadBalancerRequest.md) |  | 

### Return type

[**LoadBalancerDetailed**](LoadBalancerDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadBalancer

> LoadBalancerDetailed CreateLoadBalancer(ctx, projectId, loadBalancerName).CreateLoadBalancerRequest(createLoadBalancerRequest).Execute()

Create a new Load Balancer in a project

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
    loadBalancerName := "loadBalancerName_example" // string | Name of the load balancer
    createLoadBalancerRequest := *openapiclient.NewCreateLoadBalancerRequest("Algorithm_example", int32(123), "HealthCheckProtocol_example", "PrivateSubnetId_example", int32(123)) // CreateLoadBalancerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancerApi.CreateLoadBalancer(context.Background(), projectId, loadBalancerName).CreateLoadBalancerRequest(createLoadBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.CreateLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoadBalancer`: LoadBalancerDetailed
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.CreateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**loadBalancerName** | **string** | Name of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createLoadBalancerRequest** | [**CreateLoadBalancerRequest**](CreateLoadBalancerRequest.md) |  | 

### Return type

[**LoadBalancerDetailed**](LoadBalancerDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocationLoadBalancer

> LoadBalancerDetailed CreateLocationLoadBalancer(ctx, projectId, location, loadBalancerName).CreateLoadBalancerRequest(createLoadBalancerRequest).Execute()

Create a new Load Balancer in a specific location of a project

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
    loadBalancerName := "loadBalancerName_example" // string | Name of the load balancer
    createLoadBalancerRequest := *openapiclient.NewCreateLoadBalancerRequest("Algorithm_example", int32(123), "HealthCheckProtocol_example", "PrivateSubnetId_example", int32(123)) // CreateLoadBalancerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancerApi.CreateLocationLoadBalancer(context.Background(), projectId, location, loadBalancerName).CreateLoadBalancerRequest(createLoadBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.CreateLocationLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationLoadBalancer`: LoadBalancerDetailed
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.CreateLocationLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**loadBalancerName** | **string** | Name of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createLoadBalancerRequest** | [**CreateLoadBalancerRequest**](CreateLoadBalancerRequest.md) |  | 

### Return type

[**LoadBalancerDetailed**](LoadBalancerDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancer

> DeleteLoadBalancer(ctx, projectId, location, loadBalancerName).Execute()

Delete a specific Load Balancer

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
    loadBalancerName := "loadBalancerName_example" // string | Name of the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LoadBalancerApi.DeleteLoadBalancer(context.Background(), projectId, location, loadBalancerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.DeleteLoadBalancer``: %v\n", err)
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
**loadBalancerName** | **string** | Name of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerRequest struct via the builder pattern


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


## DeleteLoadBalancerWithID

> DeleteLoadBalancerWithID(ctx, projectId, location, loadBalancerId).Execute()

Delete a specific Load Balancer with ID

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
    loadBalancerId := "loadBalancerId_example" // string | ID of the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LoadBalancerApi.DeleteLoadBalancerWithID(context.Background(), projectId, location, loadBalancerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.DeleteLoadBalancerWithID``: %v\n", err)
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
**loadBalancerId** | **string** | ID of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerWithIDRequest struct via the builder pattern


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


## DetachVmLocationLoadBalancer

> LoadBalancerDetailed DetachVmLocationLoadBalancer(ctx, projectId, location, loadBalancerName).DetachVmLocationLoadBalancerRequest(detachVmLocationLoadBalancerRequest).Execute()

Detach a VM from a Load Balancer in a specific location of a project

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
    loadBalancerName := "loadBalancerName_example" // string | Name of the load balancer
    detachVmLocationLoadBalancerRequest := *openapiclient.NewDetachVmLocationLoadBalancerRequest("VmId_example") // DetachVmLocationLoadBalancerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancerApi.DetachVmLocationLoadBalancer(context.Background(), projectId, location, loadBalancerName).DetachVmLocationLoadBalancerRequest(detachVmLocationLoadBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.DetachVmLocationLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachVmLocationLoadBalancer`: LoadBalancerDetailed
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.DetachVmLocationLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**loadBalancerName** | **string** | Name of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachVmLocationLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **detachVmLocationLoadBalancerRequest** | [**DetachVmLocationLoadBalancerRequest**](DetachVmLocationLoadBalancerRequest.md) |  | 

### Return type

[**LoadBalancerDetailed**](LoadBalancerDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancer

> LoadBalancerDetailed GetLoadBalancer(ctx, projectId, loadBalancerName).Execute()

Get details of a specific Load Balancer

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
    loadBalancerName := "loadBalancerName_example" // string | Name of the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancerApi.GetLoadBalancer(context.Background(), projectId, loadBalancerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.GetLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancer`: LoadBalancerDetailed
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.GetLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**loadBalancerName** | **string** | Name of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoadBalancerDetailed**](LoadBalancerDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerDetails

> LoadBalancerDetailed GetLoadBalancerDetails(ctx, projectId, location, loadBalancerName).Execute()

Get details of a specific Load Balancer in a location

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
    loadBalancerName := "loadBalancerName_example" // string | Name of the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancerApi.GetLoadBalancerDetails(context.Background(), projectId, location, loadBalancerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.GetLoadBalancerDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancerDetails`: LoadBalancerDetailed
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.GetLoadBalancerDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**loadBalancerName** | **string** | Name of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**LoadBalancerDetailed**](LoadBalancerDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoadBalancerDetailsWithId

> LoadBalancerDetailed GetLoadBalancerDetailsWithId(ctx, projectId, location, loadBalancerId).Execute()

Get details of a specific Load Balancer in a location with ID

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
    loadBalancerId := "loadBalancerId_example" // string | ID of the load balancer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancerApi.GetLoadBalancerDetailsWithId(context.Background(), projectId, location, loadBalancerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.GetLoadBalancerDetailsWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoadBalancerDetailsWithId`: LoadBalancerDetailed
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.GetLoadBalancerDetailsWithId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**loadBalancerId** | **string** | ID of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerDetailsWithIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**LoadBalancerDetailed**](LoadBalancerDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoadBalancers

> ListLoadBalancers200Response ListLoadBalancers(ctx, projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()

List Load Balancers in a specific project

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
    startAfter := "startAfter_example" // string | Pagination - Start after (optional)
    pageSize := int32(56) // int32 | Pagination - Page size (optional) (default to 10)
    orderColumn := "orderColumn_example" // string | Pagination - Order column (optional) (default to "id")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancerApi.ListLoadBalancers(context.Background(), projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.ListLoadBalancers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoadBalancers`: ListLoadBalancers200Response
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.ListLoadBalancers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAfter** | **string** | Pagination - Start after | 
 **pageSize** | **int32** | Pagination - Page size | [default to 10]
 **orderColumn** | **string** | Pagination - Order column | [default to &quot;id&quot;]

### Return type

[**ListLoadBalancers200Response**](ListLoadBalancers200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocationLoadBalancers

> ListLoadBalancers200Response ListLocationLoadBalancers(ctx, projectId, location).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()

List Load Balancers in a specific location of a project

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
    startAfter := "startAfter_example" // string | Pagination - Start after (optional)
    pageSize := int32(56) // int32 | Pagination - Page size (optional) (default to 10)
    orderColumn := "orderColumn_example" // string | Pagination - Order column (optional) (default to "id")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancerApi.ListLocationLoadBalancers(context.Background(), projectId, location).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.ListLocationLoadBalancers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocationLoadBalancers`: ListLoadBalancers200Response
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.ListLocationLoadBalancers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocationLoadBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAfter** | **string** | Pagination - Start after | 
 **pageSize** | **int32** | Pagination - Page size | [default to 10]
 **orderColumn** | **string** | Pagination - Order column | [default to &quot;id&quot;]

### Return type

[**ListLoadBalancers200Response**](ListLoadBalancers200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchLocationLoadBalancer

> LoadBalancerDetailed PatchLocationLoadBalancer(ctx, projectId, location, loadBalancerName).PatchLocationLoadBalancerRequest(patchLocationLoadBalancerRequest).Execute()

Update a Load Balancer in a specific location of a project

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
    loadBalancerName := "loadBalancerName_example" // string | Name of the load balancer
    patchLocationLoadBalancerRequest := *openapiclient.NewPatchLocationLoadBalancerRequest() // PatchLocationLoadBalancerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoadBalancerApi.PatchLocationLoadBalancer(context.Background(), projectId, location, loadBalancerName).PatchLocationLoadBalancerRequest(patchLocationLoadBalancerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerApi.PatchLocationLoadBalancer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchLocationLoadBalancer`: LoadBalancerDetailed
    fmt.Fprintf(os.Stdout, "Response from `LoadBalancerApi.PatchLocationLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**loadBalancerName** | **string** | Name of the load balancer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchLocationLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchLocationLoadBalancerRequest** | [**PatchLocationLoadBalancerRequest**](PatchLocationLoadBalancerRequest.md) |  | 

### Return type

[**LoadBalancerDetailed**](LoadBalancerDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

