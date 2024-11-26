# \VirtualMachineApi

All URIs are relative to *https://api.ubicloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVM**](VirtualMachineApi.md#CreateVM) | **Post** /project/{project_id}/location/{location}/vm/{vm_name} | Create a new VM in a specific location of a project
[**DeleteVM**](VirtualMachineApi.md#DeleteVM) | **Delete** /project/{project_id}/location/{location}/vm/{vm_name} | Delete a specific VM
[**DeleteVMWithId**](VirtualMachineApi.md#DeleteVMWithId) | **Delete** /project/{project_id}/location/{location}/vm/id/{vm_id} | Delete a specific VM with ID
[**GetVMDetails**](VirtualMachineApi.md#GetVMDetails) | **Get** /project/{project_id}/location/{location}/vm/{vm_name} | Get details of a specific VM in a location
[**GetVMDetailsWithId**](VirtualMachineApi.md#GetVMDetailsWithId) | **Get** /project/{project_id}/location/{location}/vm/id/{vm_id} | Get details of a specific VM in a location with ID
[**ListLocationVMs**](VirtualMachineApi.md#ListLocationVMs) | **Get** /project/{project_id}/location/{location}/vm | List VMs in a specific location of a project
[**ListProjectVMs**](VirtualMachineApi.md#ListProjectVMs) | **Get** /project/{project_id}/vm | List all VMs created under the given project ID and visible to logged in user



## CreateVM

> VmDetailed CreateVM(ctx, projectId, location, vmName).CreateVMRequest(createVMRequest).Execute()

Create a new VM in a specific location of a project

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
    vmName := "vmName_example" // string | Virtual machine name
    createVMRequest := *openapiclient.NewCreateVMRequest("PublicKey_example") // CreateVMRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualMachineApi.CreateVM(context.Background(), projectId, location, vmName).CreateVMRequest(createVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineApi.CreateVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVM`: VmDetailed
    fmt.Fprintf(os.Stdout, "Response from `VirtualMachineApi.CreateVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**vmName** | **string** | Virtual machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createVMRequest** | [**CreateVMRequest**](CreateVMRequest.md) |  | 

### Return type

[**VmDetailed**](VmDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVM

> DeleteVM(ctx, projectId, location, vmName).Execute()

Delete a specific VM

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
    vmName := "vmName_example" // string | Virtual machine name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VirtualMachineApi.DeleteVM(context.Background(), projectId, location, vmName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineApi.DeleteVM``: %v\n", err)
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
**vmName** | **string** | Virtual machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMRequest struct via the builder pattern


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


## DeleteVMWithId

> DeleteVMWithId(ctx, projectId, location, vmId).Execute()

Delete a specific VM with ID

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
    vmId := "vmId_example" // string | Virtual machine ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VirtualMachineApi.DeleteVMWithId(context.Background(), projectId, location, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineApi.DeleteVMWithId``: %v\n", err)
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
**vmId** | **string** | Virtual machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMWithIdRequest struct via the builder pattern


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


## GetVMDetails

> VmDetailed GetVMDetails(ctx, projectId, location, vmName).Execute()

Get details of a specific VM in a location

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
    vmName := "vmName_example" // string | Virtual machine name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualMachineApi.GetVMDetails(context.Background(), projectId, location, vmName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineApi.GetVMDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMDetails`: VmDetailed
    fmt.Fprintf(os.Stdout, "Response from `VirtualMachineApi.GetVMDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**vmName** | **string** | Virtual machine name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**VmDetailed**](VmDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMDetailsWithId

> VmDetailed GetVMDetailsWithId(ctx, projectId, location, vmId).Execute()

Get details of a specific VM in a location with ID

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
    vmId := "vmId_example" // string | Virtual machine ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualMachineApi.GetVMDetailsWithId(context.Background(), projectId, location, vmId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineApi.GetVMDetailsWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMDetailsWithId`: VmDetailed
    fmt.Fprintf(os.Stdout, "Response from `VirtualMachineApi.GetVMDetailsWithId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**vmId** | **string** | Virtual machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMDetailsWithIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**VmDetailed**](VmDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocationVMs

> ListLocationVMs200Response ListLocationVMs(ctx, location, projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()

List VMs in a specific location of a project

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
    startAfter := "startAfter_example" // string | Pagination - Start after (optional)
    pageSize := int32(56) // int32 | Pagination - Page size (optional) (default to 10)
    orderColumn := "orderColumn_example" // string | Pagination - Order column (optional) (default to "id")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualMachineApi.ListLocationVMs(context.Background(), location, projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineApi.ListLocationVMs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocationVMs`: ListLocationVMs200Response
    fmt.Fprintf(os.Stdout, "Response from `VirtualMachineApi.ListLocationVMs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocationVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAfter** | **string** | Pagination - Start after | 
 **pageSize** | **int32** | Pagination - Page size | [default to 10]
 **orderColumn** | **string** | Pagination - Order column | [default to &quot;id&quot;]

### Return type

[**ListLocationVMs200Response**](ListLocationVMs200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjectVMs

> ListLocationVMs200Response ListProjectVMs(ctx, projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()

List all VMs created under the given project ID and visible to logged in user

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
    resp, r, err := apiClient.VirtualMachineApi.ListProjectVMs(context.Background(), projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualMachineApi.ListProjectVMs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProjectVMs`: ListLocationVMs200Response
    fmt.Fprintf(os.Stdout, "Response from `VirtualMachineApi.ListProjectVMs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAfter** | **string** | Pagination - Start after | 
 **pageSize** | **int32** | Pagination - Page size | [default to 10]
 **orderColumn** | **string** | Pagination - Order column | [default to &quot;id&quot;]

### Return type

[**ListLocationVMs200Response**](ListLocationVMs200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

