# \SystemAdministrationApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCloudData**](SystemAdministrationApi.md#GetCloudData) | **Get** /cloud_data | Get cloud details
[**GetRemoteSupportDetails**](SystemAdministrationApi.md#GetRemoteSupportDetails) | **Get** /remote_support | Get remote support
[**GetRuntimeStatus**](SystemAdministrationApi.md#GetRuntimeStatus) | **Get** /status | Get runtime status
[**GetSystemStatus**](SystemAdministrationApi.md#GetSystemStatus) | **Get** /status/system | Get system status
[**GetTaskStatus**](SystemAdministrationApi.md#GetTaskStatus) | **Get** /status/task | Get task status
[**PostGenerateSupportKeypair**](SystemAdministrationApi.md#PostGenerateSupportKeypair) | **Post** /remote_support/keypair | Generate support keypair
[**PutServerAction**](SystemAdministrationApi.md#PutServerAction) | **Put** /server | Take server action
[**PutUpdateRemoteSupport**](SystemAdministrationApi.md#PutUpdateRemoteSupport) | **Put** /remote_support | Update remote support config



## GetCloudData

> CloudInfoDetail GetCloudData(ctx).Execute()

Get cloud details



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
    resp, r, err := apiClient.SystemAdministrationApi.GetCloudData(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdministrationApi.GetCloudData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudData`: CloudInfoDetail
    fmt.Fprintf(os.Stdout, "Response from `SystemAdministrationApi.GetCloudData`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudDataRequest struct via the builder pattern


### Return type

[**CloudInfoDetail**](CloudInfoDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteSupportDetails

> RemoteSupportConfigResponse GetRemoteSupportDetails(ctx).Execute()

Get remote support



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
    resp, r, err := apiClient.SystemAdministrationApi.GetRemoteSupportDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdministrationApi.GetRemoteSupportDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteSupportDetails`: RemoteSupportConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAdministrationApi.GetRemoteSupportDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteSupportDetailsRequest struct via the builder pattern


### Return type

[**RemoteSupportConfigResponse**](RemoteSupportConfigResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuntimeStatus

> RuntimeStatusDetail GetRuntimeStatus(ctx).Execute()

Get runtime status



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
    resp, r, err := apiClient.SystemAdministrationApi.GetRuntimeStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdministrationApi.GetRuntimeStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRuntimeStatus`: RuntimeStatusDetail
    fmt.Fprintf(os.Stdout, "Response from `SystemAdministrationApi.GetRuntimeStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuntimeStatusRequest struct via the builder pattern


### Return type

[**RuntimeStatusDetail**](RuntimeStatusDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemStatus

> SystemStatusDetail GetSystemStatus(ctx).Timestamp(timestamp).Execute()

Get system status



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
    timestamp := int32(56) // int32 | Unix timestamp (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAdministrationApi.GetSystemStatus(context.Background()).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdministrationApi.GetSystemStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemStatus`: SystemStatusDetail
    fmt.Fprintf(os.Stdout, "Response from `SystemAdministrationApi.GetSystemStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **int32** | Unix timestamp | 

### Return type

[**SystemStatusDetail**](SystemStatusDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskStatus

> TaskStatusDetail GetTaskStatus(ctx).GetTaskTokenRequest(getTaskTokenRequest).Execute()

Get task status



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
    getTaskTokenRequest := *openapiclient.NewGetTaskTokenRequest() // GetTaskTokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAdministrationApi.GetTaskStatus(context.Background()).GetTaskTokenRequest(getTaskTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdministrationApi.GetTaskStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskStatus`: TaskStatusDetail
    fmt.Fprintf(os.Stdout, "Response from `SystemAdministrationApi.GetTaskStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getTaskTokenRequest** | [**GetTaskTokenRequest**](GetTaskTokenRequest.md) |  | 

### Return type

[**TaskStatusDetail**](TaskStatusDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGenerateSupportKeypair

> *os.File PostGenerateSupportKeypair(ctx).Body(body).Execute()

Generate support keypair



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
    body := os.NewFile(1234, "some_file") // *os.File | Encrypted passphrase file which will be used to generate an X509 key for  accessing the VNS3 Manager in support mode. These are generated and owned by Cohesive.  Contact support@cohesive.net for an encrypted passphrase file. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAdministrationApi.PostGenerateSupportKeypair(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdministrationApi.PostGenerateSupportKeypair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostGenerateSupportKeypair`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SystemAdministrationApi.PostGenerateSupportKeypair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostGenerateSupportKeypairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | Encrypted passphrase file which will be used to generate an X509 key for  accessing the VNS3 Manager in support mode. These are generated and owned by Cohesive.  Contact support@cohesive.net for an encrypted passphrase file.  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/octet-stream, text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutServerAction

> SimpleStatusResponse PutServerAction(ctx).RebootRequest(rebootRequest).Execute()

Take server action



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
    rebootRequest := *openapiclient.NewRebootRequest() // RebootRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAdministrationApi.PutServerAction(context.Background()).RebootRequest(rebootRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdministrationApi.PutServerAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutServerAction`: SimpleStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAdministrationApi.PutServerAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutServerActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rebootRequest** | [**RebootRequest**](RebootRequest.md) |  | 

### Return type

[**SimpleStatusResponse**](SimpleStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateRemoteSupport

> RemoteSupportStatusResponse PutUpdateRemoteSupport(ctx).UpdateRemoteSupportConfigRequest(updateRemoteSupportConfigRequest).Execute()

Update remote support config



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
    updateRemoteSupportConfigRequest := *openapiclient.NewUpdateRemoteSupportConfigRequest() // UpdateRemoteSupportConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemAdministrationApi.PutUpdateRemoteSupport(context.Background()).UpdateRemoteSupportConfigRequest(updateRemoteSupportConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemAdministrationApi.PutUpdateRemoteSupport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateRemoteSupport`: RemoteSupportStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `SystemAdministrationApi.PutUpdateRemoteSupport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateRemoteSupportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateRemoteSupportConfigRequest** | [**UpdateRemoteSupportConfigRequest**](UpdateRemoteSupportConfigRequest.md) |  | 

### Return type

[**RemoteSupportStatusResponse**](RemoteSupportStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

