# \PeeringApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePeer**](PeeringApi.md#DeletePeer) | **Delete** /peering/peers/{peer_id} | Delete peer
[**GetPeeringStatus**](PeeringApi.md#GetPeeringStatus) | **Get** /peering | Get peering status
[**PostCreatePeer**](PeeringApi.md#PostCreatePeer) | **Post** /peering/peers | Create peer
[**PutSelfPeeringId**](PeeringApi.md#PutSelfPeeringId) | **Put** /peering/self | Set peering ID
[**PutUpdatePeer**](PeeringApi.md#PutUpdatePeer) | **Put** /peering/peers/{peer_id} | Update peer



## DeletePeer

> Object DeletePeer(ctx, peerId).Execute()

Delete peer



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
    peerId := int32(56) // int32 | Peer ID for controller peer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeeringApi.DeletePeer(context.Background(), peerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeeringApi.DeletePeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePeer`: Object
    fmt.Fprintf(os.Stdout, "Response from `PeeringApi.DeletePeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**peerId** | **int32** | Peer ID for controller peer | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePeerRequest struct via the builder pattern


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


## GetPeeringStatus

> PeersDetailResponse GetPeeringStatus(ctx).Execute()

Get peering status



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
    resp, r, err := apiClient.PeeringApi.GetPeeringStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeeringApi.GetPeeringStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPeeringStatus`: PeersDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `PeeringApi.GetPeeringStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPeeringStatusRequest struct via the builder pattern


### Return type

[**PeersDetailResponse**](PeersDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreatePeer

> Object PostCreatePeer(ctx).CreatePeerRequest(createPeerRequest).Execute()

Create peer



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
    createPeerRequest := *openapiclient.NewCreatePeerRequest(int32(123), "Name_example") // CreatePeerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeeringApi.PostCreatePeer(context.Background()).CreatePeerRequest(createPeerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeeringApi.PostCreatePeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreatePeer`: Object
    fmt.Fprintf(os.Stdout, "Response from `PeeringApi.PostCreatePeer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreatePeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPeerRequest** | [**CreatePeerRequest**](CreatePeerRequest.md) |  | 

### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSelfPeeringId

> Object PutSelfPeeringId(ctx).PeerSelfRequest(peerSelfRequest).Execute()

Set peering ID



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
    peerSelfRequest := *openapiclient.NewPeerSelfRequest(int32(123)) // PeerSelfRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeeringApi.PutSelfPeeringId(context.Background()).PeerSelfRequest(peerSelfRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeeringApi.PutSelfPeeringId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSelfPeeringId`: Object
    fmt.Fprintf(os.Stdout, "Response from `PeeringApi.PutSelfPeeringId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutSelfPeeringIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **peerSelfRequest** | [**PeerSelfRequest**](PeerSelfRequest.md) |  | 

### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdatePeer

> Object PutUpdatePeer(ctx, peerId).UpdatePeerRequest(updatePeerRequest).Execute()

Update peer



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
    peerId := int32(56) // int32 | Peer ID for controller peer
    updatePeerRequest := *openapiclient.NewUpdatePeerRequest() // UpdatePeerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PeeringApi.PutUpdatePeer(context.Background(), peerId).UpdatePeerRequest(updatePeerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PeeringApi.PutUpdatePeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdatePeer`: Object
    fmt.Fprintf(os.Stdout, "Response from `PeeringApi.PutUpdatePeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**peerId** | **int32** | Peer ID for controller peer | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdatePeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePeerRequest** | [**UpdatePeerRequest**](UpdatePeerRequest.md) |  | 

### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

