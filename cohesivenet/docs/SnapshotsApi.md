# \SnapshotsApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSnapshot**](SnapshotsApi.md#DeleteSnapshot) | **Delete** /snapshots/{snapshot_name} | Delete snapshot
[**GetDownloadSnapshot**](SnapshotsApi.md#GetDownloadSnapshot) | **Get** /snapshots/{snapshot_name} | Download snapshot
[**GetSnapshots**](SnapshotsApi.md#GetSnapshots) | **Get** /snapshots | Get snapshots
[**PostCreateSnapshot**](SnapshotsApi.md#PostCreateSnapshot) | **Post** /snapshots | Create snapshot
[**PutImportSnapshot**](SnapshotsApi.md#PutImportSnapshot) | **Put** /snapshots/running_config | Import snapshot



## DeleteSnapshot

> Object DeleteSnapshot(ctx, snapshotName).Execute()

Delete snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    snapshotName := "snapshotName_example" // string | name of Snapshot

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.DeleteSnapshot(context.Background(), snapshotName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.DeleteSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSnapshot`: Object
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.DeleteSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotName** | **string** | name of Snapshot | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDownloadSnapshot

> *os.File GetDownloadSnapshot(ctx, snapshotName).Execute()

Download snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    snapshotName := "snapshotName_example" // string | name of Snapshot

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.GetDownloadSnapshot(context.Background(), snapshotName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.GetDownloadSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadSnapshot`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.GetDownloadSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotName** | **string** | name of Snapshot | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshots

> SnapshotsListResponse GetSnapshots(ctx).Execute()

Get snapshots



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.GetSnapshots(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.GetSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshots`: SnapshotsListResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.GetSnapshots`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotsRequest struct via the builder pattern


### Return type

[**SnapshotsListResponse**](SnapshotsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateSnapshot

> SnapshotsDetailResponse PostCreateSnapshot(ctx).CreateSnapshotRequest(createSnapshotRequest).Execute()

Create snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createSnapshotRequest := *openapiclient.NewCreateSnapshotRequest() // CreateSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.PostCreateSnapshot(context.Background()).CreateSnapshotRequest(createSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.PostCreateSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateSnapshot`: SnapshotsDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.PostCreateSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSnapshotRequest** | [**CreateSnapshotRequest**](CreateSnapshotRequest.md) |  | 

### Return type

[**SnapshotsDetailResponse**](SnapshotsDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutImportSnapshot

> SnapshotImportStatusResponse PutImportSnapshot(ctx).Body(body).Async(async).Execute()

Import snapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := os.NewFile(1234, "some_file") // *os.File | Snapshot file
    async := true // bool | If true, return a task token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsApi.PutImportSnapshot(context.Background()).Body(body).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsApi.PutImportSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutImportSnapshot`: SnapshotImportStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsApi.PutImportSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutImportSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | Snapshot file | 
 **async** | **bool** | If true, return a task token | 

### Return type

[**SnapshotImportStatusResponse**](SnapshotImportStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

