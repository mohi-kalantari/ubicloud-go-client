# \KubernetesClusterApi

All URIs are relative to *https://api.ubicloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKubernetesCluster**](KubernetesClusterApi.md#CreateKubernetesCluster) | **Post** /project/{project_id}/kubernetes-cluster | Create a new kubernetes cluster
[**CreateLocationKubernetesCluster**](KubernetesClusterApi.md#CreateLocationKubernetesCluster) | **Post** /project/{project_id}/location/{location}/kubernetes-cluster/{kubernetes_cluster_name} | Create Kubernetes Cluster in a specific location of a project
[**GetKubernetesCluster**](KubernetesClusterApi.md#GetKubernetesCluster) | **Get** /project/{project_id}/kubernetes-cluster | Return the list of kubernetes clusters in the project
[**GetKubernetesClusterDetails**](KubernetesClusterApi.md#GetKubernetesClusterDetails) | **Get** /project/{project_id}/location/{location}/kubernetes-cluster/{kubernetes_cluster_name} | Get details of a specific Kubernetes Cluster in a location
[**ListLocationKubernetesClusters**](KubernetesClusterApi.md#ListLocationKubernetesClusters) | **Get** /project/{project_id}/location/{location}/kubernetes-cluster | List kubernetes clusters in a specific location of a project



## CreateKubernetesCluster

> KubernetesCluster CreateKubernetesCluster(ctx, projectId).CreateKubernetesClusterRequest(createKubernetesClusterRequest).Execute()

Create a new kubernetes cluster

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
    createKubernetesClusterRequest := *openapiclient.NewCreateKubernetesClusterRequest("Subnet_example", "Location_example", "KubernetesVersion_example") // CreateKubernetesClusterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesClusterApi.CreateKubernetesCluster(context.Background(), projectId).CreateKubernetesClusterRequest(createKubernetesClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClusterApi.CreateKubernetesCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKubernetesCluster`: KubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `KubernetesClusterApi.CreateKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createKubernetesClusterRequest** | [**CreateKubernetesClusterRequest**](CreateKubernetesClusterRequest.md) |  | 

### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLocationKubernetesCluster

> KubernetesCluster CreateLocationKubernetesCluster(ctx, projectId, location, kubernetesClusterName).CreateLocationKubernetesClusterRequest(createLocationKubernetesClusterRequest).Execute()

Create Kubernetes Cluster in a specific location of a project

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
    kubernetesClusterName := "kubernetesClusterName_example" // string | Kubernetes cluster name
    createLocationKubernetesClusterRequest := *openapiclient.NewCreateLocationKubernetesClusterRequest("Subnet_example", "KubernetesVersion_example") // CreateLocationKubernetesClusterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesClusterApi.CreateLocationKubernetesCluster(context.Background(), projectId, location, kubernetesClusterName).CreateLocationKubernetesClusterRequest(createLocationKubernetesClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClusterApi.CreateLocationKubernetesCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocationKubernetesCluster`: KubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `KubernetesClusterApi.CreateLocationKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**kubernetesClusterName** | **string** | Kubernetes cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocationKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createLocationKubernetesClusterRequest** | [**CreateLocationKubernetesClusterRequest**](CreateLocationKubernetesClusterRequest.md) |  | 

### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesCluster

> GetKubernetesCluster200Response GetKubernetesCluster(ctx, projectId).Execute()

Return the list of kubernetes clusters in the project

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
    resp, r, err := apiClient.KubernetesClusterApi.GetKubernetesCluster(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClusterApi.GetKubernetesCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesCluster`: GetKubernetesCluster200Response
    fmt.Fprintf(os.Stdout, "Response from `KubernetesClusterApi.GetKubernetesCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetKubernetesCluster200Response**](GetKubernetesCluster200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesClusterDetails

> KubernetesCluster GetKubernetesClusterDetails(ctx, projectId, location, kubernetesClusterName).Execute()

Get details of a specific Kubernetes Cluster in a location

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
    kubernetesClusterName := "kubernetesClusterName_example" // string | Kubernetes cluster name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesClusterApi.GetKubernetesClusterDetails(context.Background(), projectId, location, kubernetesClusterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClusterApi.GetKubernetesClusterDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesClusterDetails`: KubernetesCluster
    fmt.Fprintf(os.Stdout, "Response from `KubernetesClusterApi.GetKubernetesClusterDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**kubernetesClusterName** | **string** | Kubernetes cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesClusterDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**KubernetesCluster**](KubernetesCluster.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocationKubernetesClusters

> GetKubernetesCluster200Response ListLocationKubernetesClusters(ctx, projectId, location).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()

List kubernetes clusters in a specific location of a project

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
    resp, r, err := apiClient.KubernetesClusterApi.ListLocationKubernetesClusters(context.Background(), projectId, location).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesClusterApi.ListLocationKubernetesClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocationKubernetesClusters`: GetKubernetesCluster200Response
    fmt.Fprintf(os.Stdout, "Response from `KubernetesClusterApi.ListLocationKubernetesClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocationKubernetesClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAfter** | **string** | Pagination - Start after | 
 **pageSize** | **int32** | Pagination - Page size | [default to 10]
 **orderColumn** | **string** | Pagination - Order column | [default to &quot;id&quot;]

### Return type

[**GetKubernetesCluster200Response**](GetKubernetesCluster200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

