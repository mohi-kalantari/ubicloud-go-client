# \KubernetesVMApi

All URIs are relative to *https://api.ubicloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLocationKubernetesVM**](KubernetesVMApi.md#CreateLocationKubernetesVM) | **Post** /project/{project_id}/location/{location}/kubernetes-vm/{kubernetes_vm_name} | Create Kubernetes VM in a specific location of a project
[**DeleteKubernetesVMWithName**](KubernetesVMApi.md#DeleteKubernetesVMWithName) | **Delete** /project/{project_id}/location/{location}/kubernetes-vm/{kubernetes_vm_name} | Delete a specific Kubernetes VM with name
[**GetKubernetesVMDetails**](KubernetesVMApi.md#GetKubernetesVMDetails) | **Get** /project/{project_id}/location/{location}/kubernetes-vm/{kubernetes_vm_name} | Get details of a specific Kubernetes VM in a location
[**ListLocationKubernetesVMs**](KubernetesVMApi.md#ListLocationKubernetesVMs) | **Get** /project/{project_id}/location/{location}/kubernetes-vm | List Kubernetes VMs in a specific location of a project



## CreateLocationKubernetesVM

> VmDetailed CreateLocationKubernetesVM(ctx, projectId, location, kubernetesVmName).CreateLocationKubernetesVMRequest(createLocationKubernetesVMRequest).Execute()

Create Kubernetes VM in a specific location of a project

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
    kubernetesVmName := "kubernetesVmName_example" // string | Kubernetes vm name
    createLocationKubernetesVMRequest := *openapiclient.NewCreateLocationKubernetesVMRequest("BootImage_example", false, "PrivateSubnetId_example", "Size_example", int32(123), "UnixUser_example", []string{"Commands_example"}) // CreateLocationKubernetesVMRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesVMApi.CreateLocationKubernetesVM(context.Background(), projectId, location, kubernetesVmName).CreateLocationKubernetesVMRequest(createLocationKubernetesVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesVMApi.CreateLocationKubernetesVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationKubernetesVM`: VmDetailed
    fmt.Fprintf(os.Stdout, "Response from `KubernetesVMApi.CreateLocationKubernetesVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**kubernetesVmName** | **string** | Kubernetes vm name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationKubernetesVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createLocationKubernetesVMRequest** | [**CreateLocationKubernetesVMRequest**](CreateLocationKubernetesVMRequest.md) |  | 

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


## DeleteKubernetesVMWithName

> DeleteKubernetesVMWithName(ctx, projectId, location, kubernetesVmName).Execute()

Delete a specific Kubernetes VM with name

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
    kubernetesVmName := "kubernetesVmName_example" // string | Kubernetes vm name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesVMApi.DeleteKubernetesVMWithName(context.Background(), projectId, location, kubernetesVmName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesVMApi.DeleteKubernetesVMWithName``: %v\n", err)
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
**kubernetesVmName** | **string** | Kubernetes vm name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKubernetesVMWithNameRequest struct via the builder pattern


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


## GetKubernetesVMDetails

> VmDetailed GetKubernetesVMDetails(ctx, projectId, location, kubernetesVmName).Execute()

Get details of a specific Kubernetes VM in a location

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
    kubernetesVmName := "kubernetesVmName_example" // string | Kubernetes vm name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesVMApi.GetKubernetesVMDetails(context.Background(), projectId, location, kubernetesVmName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesVMApi.GetKubernetesVMDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesVMDetails`: VmDetailed
    fmt.Fprintf(os.Stdout, "Response from `KubernetesVMApi.GetKubernetesVMDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**kubernetesVmName** | **string** | Kubernetes vm name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesVMDetailsRequest struct via the builder pattern


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


## ListLocationKubernetesVMs

> ListLocationKubernetesVMs200Response ListLocationKubernetesVMs(ctx, location, projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()

List Kubernetes VMs in a specific location of a project

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
    resp, r, err := apiClient.KubernetesVMApi.ListLocationKubernetesVMs(context.Background(), location, projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesVMApi.ListLocationKubernetesVMs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocationKubernetesVMs`: ListLocationKubernetesVMs200Response
    fmt.Fprintf(os.Stdout, "Response from `KubernetesVMApi.ListLocationKubernetesVMs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**location** | **string** | The Ubicloud location/region | 
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocationKubernetesVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAfter** | **string** | Pagination - Start after | 
 **pageSize** | **int32** | Pagination - Page size | [default to 10]
 **orderColumn** | **string** | Pagination - Order column | [default to &quot;id&quot;]

### Return type

[**ListLocationKubernetesVMs200Response**](ListLocationKubernetesVMs200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

