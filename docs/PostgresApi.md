# \PostgresApi

All URIs are relative to *https://api.ubicloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePostgresDatabase**](PostgresApi.md#CreatePostgresDatabase) | **Post** /project/{project_id}/location/{location}/postgres/{postgres_database_name} | Create a new Postgres database in a specific location of a project
[**DeletePostgresDatabase**](PostgresApi.md#DeletePostgresDatabase) | **Delete** /project/{project_id}/location/{location}/postgres/{postgres_database_name} | Delete a specific Postgres database
[**DeletePostgresDatabaseWithID**](PostgresApi.md#DeletePostgresDatabaseWithID) | **Delete** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id} | Delete a specific Postgres database with ID
[**FailoverPostgresDatabaseWithID**](PostgresApi.md#FailoverPostgresDatabaseWithID) | **Post** /project/{project_id}/location/{location}/postgres/_{postgres_database_id}/failover | Failover a specific Postgres database with ID
[**GetPostgresDatabaseDetails**](PostgresApi.md#GetPostgresDatabaseDetails) | **Get** /project/{project_id}/location/{location}/postgres/{postgres_database_name} | Get details of a specific Postgres database in a location
[**GetPostgresDetailsWithId**](PostgresApi.md#GetPostgresDetailsWithId) | **Get** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id} | Get details of a specific Postgres database in a location with ID
[**ListLocationPostgresDatabases**](PostgresApi.md#ListLocationPostgresDatabases) | **Get** /project/{project_id}/location/{location}/postgres | List Postgres databases in a specific location of a project
[**ListPostgresDatabases**](PostgresApi.md#ListPostgresDatabases) | **Get** /project/{project_id}/postgres | List visible Postgres databases
[**ResetSuperuserPassword**](PostgresApi.md#ResetSuperuserPassword) | **Post** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/reset-superuser-password | Reset superuser password of the Postgres database
[**ResetSuperuserPasswordWithID**](PostgresApi.md#ResetSuperuserPasswordWithID) | **Post** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/reset-superuser-password | Reset super-user password of the Postgres database
[**RestorePostgresDatabase**](PostgresApi.md#RestorePostgresDatabase) | **Post** /project/{project_id}/location/{location}/postgres/{postgres_database_name}/restore | Restore a new Postgres database in a specific location of a project
[**RestorePostgresDatabaseWithID**](PostgresApi.md#RestorePostgresDatabaseWithID) | **Post** /project/{project_id}/location/{location}/postgres/id/{postgres_database_id}/restore | Restore a new Postgres database in a specific location of a project with ID



## CreatePostgresDatabase

> PostgresDetailed CreatePostgresDatabase(ctx, projectId, location, postgresDatabaseName).CreatePostgresDatabaseRequest(createPostgresDatabaseRequest).Execute()

Create a new Postgres database in a specific location of a project

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
    createPostgresDatabaseRequest := *openapiclient.NewCreatePostgresDatabaseRequest("Size_example") // CreatePostgresDatabaseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.CreatePostgresDatabase(context.Background(), projectId, location, postgresDatabaseName).CreatePostgresDatabaseRequest(createPostgresDatabaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.CreatePostgresDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePostgresDatabase`: PostgresDetailed
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.CreatePostgresDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseName** | **string** | Postgres database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostgresDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **createPostgresDatabaseRequest** | [**CreatePostgresDatabaseRequest**](CreatePostgresDatabaseRequest.md) |  | 

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


## DeletePostgresDatabase

> DeletePostgresDatabase(ctx, projectId, location, postgresDatabaseName).Execute()

Delete a specific Postgres database

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PostgresApi.DeletePostgresDatabase(context.Background(), projectId, location, postgresDatabaseName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.DeletePostgresDatabase``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostgresDatabaseRequest struct via the builder pattern


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


## DeletePostgresDatabaseWithID

> DeletePostgresDatabaseWithID(ctx, projectId, location, postgresDatabaseId).Execute()

Delete a specific Postgres database with ID

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
    postgresDatabaseId := "postgresDatabaseId_example" // string | Postgres database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PostgresApi.DeletePostgresDatabaseWithID(context.Background(), projectId, location, postgresDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.DeletePostgresDatabaseWithID``: %v\n", err)
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
**postgresDatabaseId** | **string** | Postgres database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostgresDatabaseWithIDRequest struct via the builder pattern


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


## FailoverPostgresDatabaseWithID

> PostgresDetailed FailoverPostgresDatabaseWithID(ctx, projectId, location, postgresDatabaseId).Execute()

Failover a specific Postgres database with ID

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
    postgresDatabaseId := "postgresDatabaseId_example" // string | Postgres database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.FailoverPostgresDatabaseWithID(context.Background(), projectId, location, postgresDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.FailoverPostgresDatabaseWithID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FailoverPostgresDatabaseWithID`: PostgresDetailed
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.FailoverPostgresDatabaseWithID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseId** | **string** | Postgres database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiFailoverPostgresDatabaseWithIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PostgresDetailed**](PostgresDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostgresDatabaseDetails

> PostgresDetailed GetPostgresDatabaseDetails(ctx, projectId, location, postgresDatabaseName).Execute()

Get details of a specific Postgres database in a location

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.GetPostgresDatabaseDetails(context.Background(), projectId, location, postgresDatabaseName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.GetPostgresDatabaseDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostgresDatabaseDetails`: PostgresDetailed
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.GetPostgresDatabaseDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseName** | **string** | Postgres database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostgresDatabaseDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PostgresDetailed**](PostgresDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostgresDetailsWithId

> PostgresDetailed GetPostgresDetailsWithId(ctx, projectId, location, postgresDatabaseId).Execute()

Get details of a specific Postgres database in a location with ID

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
    postgresDatabaseId := "postgresDatabaseId_example" // string | Postgres database ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.GetPostgresDetailsWithId(context.Background(), projectId, location, postgresDatabaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.GetPostgresDetailsWithId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostgresDetailsWithId`: PostgresDetailed
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.GetPostgresDetailsWithId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseId** | **string** | Postgres database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostgresDetailsWithIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PostgresDetailed**](PostgresDetailed.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocationPostgresDatabases

> ListLocationPostgresDatabases200Response ListLocationPostgresDatabases(ctx, projectId, location).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()

List Postgres databases in a specific location of a project

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
    resp, r, err := apiClient.PostgresApi.ListLocationPostgresDatabases(context.Background(), projectId, location).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.ListLocationPostgresDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocationPostgresDatabases`: ListLocationPostgresDatabases200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.ListLocationPostgresDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocationPostgresDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAfter** | **string** | Pagination - Start after | 
 **pageSize** | **int32** | Pagination - Page size | [default to 10]
 **orderColumn** | **string** | Pagination - Order column | [default to &quot;id&quot;]

### Return type

[**ListLocationPostgresDatabases200Response**](ListLocationPostgresDatabases200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostgresDatabases

> ListLocationPostgresDatabases200Response ListPostgresDatabases(ctx, projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()

List visible Postgres databases

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
    resp, r, err := apiClient.PostgresApi.ListPostgresDatabases(context.Background(), projectId).StartAfter(startAfter).PageSize(pageSize).OrderColumn(orderColumn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.ListPostgresDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPostgresDatabases`: ListLocationPostgresDatabases200Response
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.ListPostgresDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPostgresDatabasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAfter** | **string** | Pagination - Start after | 
 **pageSize** | **int32** | Pagination - Page size | [default to 10]
 **orderColumn** | **string** | Pagination - Order column | [default to &quot;id&quot;]

### Return type

[**ListLocationPostgresDatabases200Response**](ListLocationPostgresDatabases200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetSuperuserPassword

> PostgresDetailed ResetSuperuserPassword(ctx, projectId, location, postgresDatabaseName).ResetSuperuserPasswordWithIDRequest(resetSuperuserPasswordWithIDRequest).Execute()

Reset superuser password of the Postgres database

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
    resetSuperuserPasswordWithIDRequest := *openapiclient.NewResetSuperuserPasswordWithIDRequest("Password_example") // ResetSuperuserPasswordWithIDRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.ResetSuperuserPassword(context.Background(), projectId, location, postgresDatabaseName).ResetSuperuserPasswordWithIDRequest(resetSuperuserPasswordWithIDRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.ResetSuperuserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetSuperuserPassword`: PostgresDetailed
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.ResetSuperuserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseName** | **string** | Postgres database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetSuperuserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resetSuperuserPasswordWithIDRequest** | [**ResetSuperuserPasswordWithIDRequest**](ResetSuperuserPasswordWithIDRequest.md) |  | 

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


## ResetSuperuserPasswordWithID

> PostgresDetailed ResetSuperuserPasswordWithID(ctx, projectId, location, postgresDatabaseId).ResetSuperuserPasswordWithIDRequest(resetSuperuserPasswordWithIDRequest).Execute()

Reset super-user password of the Postgres database

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
    postgresDatabaseId := "postgresDatabaseId_example" // string | Postgres database ID
    resetSuperuserPasswordWithIDRequest := *openapiclient.NewResetSuperuserPasswordWithIDRequest("Password_example") // ResetSuperuserPasswordWithIDRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.ResetSuperuserPasswordWithID(context.Background(), projectId, location, postgresDatabaseId).ResetSuperuserPasswordWithIDRequest(resetSuperuserPasswordWithIDRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.ResetSuperuserPasswordWithID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetSuperuserPasswordWithID`: PostgresDetailed
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.ResetSuperuserPasswordWithID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseId** | **string** | Postgres database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetSuperuserPasswordWithIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resetSuperuserPasswordWithIDRequest** | [**ResetSuperuserPasswordWithIDRequest**](ResetSuperuserPasswordWithIDRequest.md) |  | 

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


## RestorePostgresDatabase

> PostgresDetailed RestorePostgresDatabase(ctx, projectId, location, postgresDatabaseName).RestorePostgresDatabaseWithIDRequest(restorePostgresDatabaseWithIDRequest).Execute()

Restore a new Postgres database in a specific location of a project

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
    restorePostgresDatabaseWithIDRequest := *openapiclient.NewRestorePostgresDatabaseWithIDRequest("Name_example", "RestoreTarget_example") // RestorePostgresDatabaseWithIDRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.RestorePostgresDatabase(context.Background(), projectId, location, postgresDatabaseName).RestorePostgresDatabaseWithIDRequest(restorePostgresDatabaseWithIDRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.RestorePostgresDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestorePostgresDatabase`: PostgresDetailed
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.RestorePostgresDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseName** | **string** | Postgres database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestorePostgresDatabaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restorePostgresDatabaseWithIDRequest** | [**RestorePostgresDatabaseWithIDRequest**](RestorePostgresDatabaseWithIDRequest.md) |  | 

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


## RestorePostgresDatabaseWithID

> Postgres RestorePostgresDatabaseWithID(ctx, projectId, location, postgresDatabaseId).RestorePostgresDatabaseWithIDRequest(restorePostgresDatabaseWithIDRequest).Execute()

Restore a new Postgres database in a specific location of a project with ID

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
    postgresDatabaseId := "postgresDatabaseId_example" // string | Postgres database ID
    restorePostgresDatabaseWithIDRequest := *openapiclient.NewRestorePostgresDatabaseWithIDRequest("Name_example", "RestoreTarget_example") // RestorePostgresDatabaseWithIDRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresApi.RestorePostgresDatabaseWithID(context.Background(), projectId, location, postgresDatabaseId).RestorePostgresDatabaseWithIDRequest(restorePostgresDatabaseWithIDRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresApi.RestorePostgresDatabaseWithID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestorePostgresDatabaseWithID`: Postgres
    fmt.Fprintf(os.Stdout, "Response from `PostgresApi.RestorePostgresDatabaseWithID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project | 
**location** | **string** | The Ubicloud location/region | 
**postgresDatabaseId** | **string** | Postgres database ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestorePostgresDatabaseWithIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **restorePostgresDatabaseWithIDRequest** | [**RestorePostgresDatabaseWithIDRequest**](RestorePostgresDatabaseWithIDRequest.md) |  | 

### Return type

[**Postgres**](Postgres.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

