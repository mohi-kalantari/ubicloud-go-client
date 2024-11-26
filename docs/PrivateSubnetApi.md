# \PrivateSubnetApi

All URIs are relative to *https://api.ubicloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrivateSubnet**](PrivateSubnetApi.md#CreatePrivateSubnet) | **Post** /project/{project_id}/location/{location}/private-subnet/{private_subnet_name} | Create a new Private Subnet in a specific location of a project
[**DeletePSWithId**](PrivateSubnetApi.md#DeletePSWithId) | **Delete** /project/{project_id}/location/{location}/private-subnet/id/{private_subnet_id} | Delete a specific Private Subnet with ID
[**DeletePrivateSubnet**](PrivateSubnetApi.md#DeletePrivateSubnet) | **Delete** /project/{project_id}/location/{location}/private-subnet/{private_subnet_name} | Delete a specific Private Subnet
[**GetPSDetailsWithId**](PrivateSubnetApi.md#GetPSDetailsWithId) | **Get** /project/{project_id}/location/{location}/private-subnet/id/{private_subnet_id} | Get details of a specific Private Subnet in a location with ID
[**GetPrivateSubnetDetails**](PrivateSubnetApi.md#GetPrivateSubnetDetails) | **Get** /project/{project_id}/location/{location}/private-subnet/{private_subnet_name} | Get details of a specific Private Subnet in a location
[**ListLocationPrivateSubnets**](PrivateSubnetApi.md#ListLocationPrivateSubnets) | **Get** /project/{project_id}/location/{location}/private-subnet | List Private Subnets in a specific location of a project
[**ListPSs**](PrivateSubnetApi.md#ListPSs) | **Get** /project/{project_id}/private-subnet | List visible Private Subnets



## CreatePrivateSubnet

> PrivateSubnet CreatePrivateSubnet(ctx, projectId, location, privateSubnetName).CreatePrivateSubnetRequest(createPrivateSubnetRequest).Execute()

Create a new Private Subnet in a specific location of a project

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
    privateSubnetName := "privateSubnetName_example" // string | Private subnet name
    createPrivateSubnetRequest := *openapiclient.NewCreatePrivateSubnetRequest() // CreatePrivateSubnetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateSubnetApi.CreatePrivateSubnet(context.Background(), projectId, location, privateSubnetName).CreatePrivateSubnetRequest(createPrivateSubnetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubnetApi.CreatePrivateSubnet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrivateSubnet`: PrivateSubnet
    fmt.Fprintf(os.Stdout, "Response from `PrivateSubnetApi.CreatePrivateSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**privateSubnetName** | **string** | Private subnet name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createPrivateSubnetRequest** | [**CreatePrivateSubnetRequest**](CreatePrivateSubnetRequest.md) |  | 

### Return type

[**PrivateSubnet**](PrivateSubnet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePSWithId

> DeletePSWithId(ctx, projectId, location, privateSubnetId).Execute()

Delete a specific Private Subnet with ID

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
    privateSubnetId := "privateSubnetId_example" // string | Private subnet ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PrivateSubnetApi.DeletePSWithId(context.Background(), projectId, location, privateSubnetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubnetApi.DeletePSWithId``: %v\n", err)
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
**privateSubnetId** | **string** | Private subnet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePSWithIdRequest struct via the builder pattern


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


## DeletePrivateSubnet

> DeletePrivateSubnet(ctx, projectId, location, privateSubnetName).Execute()

Delete a specific Private Subnet

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
    privateSubnetName := "privateSubnetName_example" // string | Private subnet name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PrivateSubnetApi.DeletePrivateSubnet(context.Background(), projectId, location, privateSubnetName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubnetApi.DeletePrivateSubnet``: %v\n", err)
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
**privateSubnetName** | **string** | Private subnet name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateSubnetRequest struct via the builder pattern


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


## GetPSDetailsWithId

> PrivateSubnet GetPSDetailsWithId(ctx, projectId, location, privateSubnetId).Execute()

Get details of a specific Private Subnet in a location with ID

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
    privateSubnetId := "privateSubnetId_example" // string | Private subnet ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateSubnetApi.GetPSDetailsWithId(context.Background(), projectId, location, privateSubnetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubnetApi.GetPSDetailsWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPSDetailsWithId`: PrivateSubnet
    fmt.Fprintf(os.Stdout, "Response from `PrivateSubnetApi.GetPSDetailsWithId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**privateSubnetId** | **string** | Private subnet ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPSDetailsWithIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PrivateSubnet**](PrivateSubnet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateSubnetDetails

> PrivateSubnet GetPrivateSubnetDetails(ctx, projectId, location, privateSubnetName).Execute()

Get details of a specific Private Subnet in a location

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
    privateSubnetName := "privateSubnetName_example" // string | Private subnet name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateSubnetApi.GetPrivateSubnetDetails(context.Background(), projectId, location, privateSubnetName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubnetApi.GetPrivateSubnetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateSubnetDetails`: PrivateSubnet
    fmt.Fprintf(os.Stdout, "Response from `PrivateSubnetApi.GetPrivateSubnetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**privateSubnetName** | **string** | Private subnet name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateSubnetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PrivateSubnet**](PrivateSubnet.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocationPrivateSubnets

> ListLocationPrivateSubnets200Response ListLocationPrivateSubnets(ctx, projectId, location).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()

List Private Subnets in a specific location of a project

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
    resp, r, err := apiClient.PrivateSubnetApi.ListLocationPrivateSubnets(context.Background(), projectId, location).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubnetApi.ListLocationPrivateSubnets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocationPrivateSubnets`: ListLocationPrivateSubnets200Response
    fmt.Fprintf(os.Stdout, "Response from `PrivateSubnetApi.ListLocationPrivateSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocationPrivateSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAfter** | **string** | Pagination - Start after | 
 **pageSize** | **int32** | Pagination - Page size | [default to 10]
 **orderColumn** | **string** | Pagination - Order column | [default to &quot;id&quot;]

### Return type

[**ListLocationPrivateSubnets200Response**](ListLocationPrivateSubnets200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPSs

> ListLocationPrivateSubnets200Response ListPSs(ctx, projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()

List visible Private Subnets

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
    resp, r, err := apiClient.PrivateSubnetApi.ListPSs(context.Background(), projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubnetApi.ListPSs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPSs`: ListLocationPrivateSubnets200Response
    fmt.Fprintf(os.Stdout, "Response from `PrivateSubnetApi.ListPSs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPSsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAfter** | **string** | Pagination - Start after | 
 **pageSize** | **int32** | Pagination - Page size | [default to 10]
 **orderColumn** | **string** | Pagination - Order column | [default to &quot;id&quot;]

### Return type

[**ListLocationPrivateSubnets200Response**](ListLocationPrivateSubnets200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

