# \DefaultApi

All URIs are relative to *https://api.ubicloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLocationKubernetesVM**](DefaultApi.md#CreateLocationKubernetesVM) | **Post** /project/{project_id}/location/{location}/kubernetes-vm/{kubernetes_vm_name} | Create Kubernetes VM in a specific location of a project



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
    resp, r, err := apiClient.DefaultApi.CreateLocationKubernetesVM(context.Background(), projectId, location, kubernetesVmName).CreateLocationKubernetesVMRequest(createLocationKubernetesVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateLocationKubernetesVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationKubernetesVM`: VmDetailed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateLocationKubernetesVM`: %v\n", resp)
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

