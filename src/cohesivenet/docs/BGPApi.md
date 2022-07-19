# \BGPApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBgpPeer**](BGPApi.md#CreateBgpPeer) | **Post** /ipsec/endpoints/{endpoint_id}/ebgp_peers | Create BGP Peer
[**DeleteBgpPeer**](BGPApi.md#DeleteBgpPeer) | **Delete** /ipsec/endpoints/{endpoint_id}/ebgp_peers/{bgp_peer_id} | Delete BGP Peer
[**GetBgpPeer**](BGPApi.md#GetBgpPeer) | **Get** /ipsec/endpoints/{endpoint_id}/ebgp_peers/{bgp_peer_id} | Get eBGP peer
[**UpdateBgpPeer**](BGPApi.md#UpdateBgpPeer) | **Put** /ipsec/endpoints/{endpoint_id}/ebgp_peers/{bgp_peer_id} | Update BGP Peer



## CreateBgpPeer

> Object CreateBgpPeer(ctx, endpointId).CreateBGPPeerRequest(createBGPPeerRequest).Execute()

Create BGP Peer



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
    endpointId := int32(56) // int32 | ID for IPsec endpoint
    createBGPPeerRequest := *openapiclient.NewCreateBGPPeerRequest("Ipaddress_example", int32(123)) // CreateBGPPeerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BGPApi.CreateBgpPeer(context.Background(), endpointId).CreateBGPPeerRequest(createBGPPeerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.CreateBgpPeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBgpPeer`: Object
    fmt.Fprintf(os.Stdout, "Response from `BGPApi.CreateBgpPeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBgpPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBGPPeerRequest** | [**CreateBGPPeerRequest**](CreateBGPPeerRequest.md) |  | 

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


## DeleteBgpPeer

> Object DeleteBgpPeer(ctx, endpointId, bgpPeerId).Execute()

Delete BGP Peer



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
    endpointId := int32(56) // int32 | ID for IPsec endpoint
    bgpPeerId := int32(56) // int32 | ID for BGP peer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BGPApi.DeleteBgpPeer(context.Background(), endpointId, bgpPeerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.DeleteBgpPeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBgpPeer`: Object
    fmt.Fprintf(os.Stdout, "Response from `BGPApi.DeleteBgpPeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 
**bgpPeerId** | **int32** | ID for BGP peer | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBgpPeerRequest struct via the builder pattern


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


## GetBgpPeer

> BGPPeerResponse GetBgpPeer(ctx, endpointId, bgpPeerId).Verbose(verbose).Execute()

Get eBGP peer



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
    endpointId := int32(56) // int32 | ID for IPsec endpoint
    bgpPeerId := int32(56) // int32 | ID for BGP peer
    verbose := true // bool | True for verbose output (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BGPApi.GetBgpPeer(context.Background(), endpointId, bgpPeerId).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.GetBgpPeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBgpPeer`: BGPPeerResponse
    fmt.Fprintf(os.Stdout, "Response from `BGPApi.GetBgpPeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 
**bgpPeerId** | **int32** | ID for BGP peer | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBgpPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **verbose** | **bool** | True for verbose output | [default to true]

### Return type

[**BGPPeerResponse**](BGPPeerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBgpPeer

> Object UpdateBgpPeer(ctx, endpointId, bgpPeerId).UpdateBGPPeerRequest(updateBGPPeerRequest).Execute()

Update BGP Peer



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
    endpointId := int32(56) // int32 | ID for IPsec endpoint
    bgpPeerId := int32(56) // int32 | ID for BGP peer
    updateBGPPeerRequest := *openapiclient.NewUpdateBGPPeerRequest("Ipaddress_example", int32(123)) // UpdateBGPPeerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BGPApi.UpdateBgpPeer(context.Background(), endpointId, bgpPeerId).UpdateBGPPeerRequest(updateBGPPeerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BGPApi.UpdateBgpPeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBgpPeer`: Object
    fmt.Fprintf(os.Stdout, "Response from `BGPApi.UpdateBgpPeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 
**bgpPeerId** | **int32** | ID for BGP peer | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBgpPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateBGPPeerRequest** | [**UpdateBGPPeerRequest**](UpdateBGPPeerRequest.md) |  | 

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

