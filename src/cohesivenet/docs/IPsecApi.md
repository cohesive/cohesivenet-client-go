# \IPsecApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIpsecEndpoint**](IPsecApi.md#DeleteIpsecEndpoint) | **Delete** /ipsec/endpoints/{endpoint_id} | Delete IPsec endpoint
[**DeleteIpsecEndpointTunnel**](IPsecApi.md#DeleteIpsecEndpointTunnel) | **Delete** /ipsec/endpoints/{endpoint_id}/tunnels/{tunnel_id} | Delete IPsec tunnel
[**DeleteIpsecTrafficPair**](IPsecApi.md#DeleteIpsecTrafficPair) | **Delete** /ipsec/endpoints/{endpoint_id}/traffic_pairs/{pair_id} | Delete IPsec traffic pair
[**DisableIpsecTrafficPair**](IPsecApi.md#DisableIpsecTrafficPair) | **Put** /ipsec/endpoints/{endpoint_id}/traffic_pairs/{pair_id}/disable | Disable IPsec traffic pair
[**EnableIpsecTrafficPair**](IPsecApi.md#EnableIpsecTrafficPair) | **Put** /ipsec/endpoints/{endpoint_id}/traffic_pairs/{pair_id}/enable | Enable IPsec traffic pair
[**GetConnectedSubnets**](IPsecApi.md#GetConnectedSubnets) | **Get** /status/connected_subnets | Get connected subnets
[**GetIpsecDetails**](IPsecApi.md#GetIpsecDetails) | **Get** /ipsec | Get IPsec details
[**GetIpsecEndpoint**](IPsecApi.md#GetIpsecEndpoint) | **Get** /ipsec/endpoints/{endpoint_id} | Get IPsec endpoint
[**GetIpsecLinkHistory**](IPsecApi.md#GetIpsecLinkHistory) | **Get** /status/link_history | Get IPsec link history
[**GetIpsecStatus**](IPsecApi.md#GetIpsecStatus) | **Get** /status/ipsec | Get IPsec status
[**PostCreateIpsecEndpoint**](IPsecApi.md#PostCreateIpsecEndpoint) | **Post** /ipsec/endpoints | Create IPsec endpoint
[**PostCreateIpsecEndpointTunnel**](IPsecApi.md#PostCreateIpsecEndpointTunnel) | **Post** /ipsec/endpoints/{endpoint_id}/tunnels | Create IPsec endpoint tunnel
[**PostCreateIpsecTrafficPair**](IPsecApi.md#PostCreateIpsecTrafficPair) | **Post** /ipsec/endpoints/{endpoint_id}/traffic_pairs | Create IPsec traffic pair
[**PostRestartIpsecAction**](IPsecApi.md#PostRestartIpsecAction) | **Post** /ipsec | Restart ipsec subystem
[**PutUpdateIpsecConfig**](IPsecApi.md#PutUpdateIpsecConfig) | **Put** /ipsec | Update IPsec config
[**PutUpdateIpsecEndpoint**](IPsecApi.md#PutUpdateIpsecEndpoint) | **Put** /ipsec/endpoints/{endpoint_id} | Update IPsec endpoint
[**PutUpdateIpsecEndpointTunnel**](IPsecApi.md#PutUpdateIpsecEndpointTunnel) | **Put** /ipsec/endpoints/{endpoint_id}/tunnels/{tunnel_id} | Update IPsec endpoint tunnel
[**PutUpdateIpsecTrafficPair**](IPsecApi.md#PutUpdateIpsecTrafficPair) | **Put** /ipsec/endpoints/{endpoint_id}/traffic_pairs/{pair_id} | Update IPsec traffic pair



## DeleteIpsecEndpoint

> Object DeleteIpsecEndpoint(ctx, endpointId).Execute()

Delete IPsec endpoint



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.DeleteIpsecEndpoint(context.Background(), endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.DeleteIpsecEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIpsecEndpoint`: Object
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.DeleteIpsecEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpsecEndpointRequest struct via the builder pattern


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


## DeleteIpsecEndpointTunnel

> Object DeleteIpsecEndpointTunnel(ctx, endpointId, tunnelId).Execute()

Delete IPsec tunnel



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
    tunnelId := int32(56) // int32 | ID for tunnel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.DeleteIpsecEndpointTunnel(context.Background(), endpointId, tunnelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.DeleteIpsecEndpointTunnel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIpsecEndpointTunnel`: Object
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.DeleteIpsecEndpointTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 
**tunnelId** | **int32** | ID for tunnel | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpsecEndpointTunnelRequest struct via the builder pattern


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


## DeleteIpsecTrafficPair

> Object DeleteIpsecTrafficPair(ctx, endpointId, pairId).Execute()

Delete IPsec traffic pair



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
    pairId := int32(56) // int32 | ID for traffic pair

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.DeleteIpsecTrafficPair(context.Background(), endpointId, pairId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.DeleteIpsecTrafficPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIpsecTrafficPair`: Object
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.DeleteIpsecTrafficPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 
**pairId** | **int32** | ID for traffic pair | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpsecTrafficPairRequest struct via the builder pattern


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


## DisableIpsecTrafficPair

> Object DisableIpsecTrafficPair(ctx, endpointId, pairId).Execute()

Disable IPsec traffic pair



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
    pairId := int32(56) // int32 | ID for traffic pair

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.DisableIpsecTrafficPair(context.Background(), endpointId, pairId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.DisableIpsecTrafficPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableIpsecTrafficPair`: Object
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.DisableIpsecTrafficPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 
**pairId** | **int32** | ID for traffic pair | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableIpsecTrafficPairRequest struct via the builder pattern


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


## EnableIpsecTrafficPair

> Object EnableIpsecTrafficPair(ctx, endpointId, pairId).Execute()

Enable IPsec traffic pair



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
    pairId := int32(56) // int32 | ID for traffic pair

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.EnableIpsecTrafficPair(context.Background(), endpointId, pairId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.EnableIpsecTrafficPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableIpsecTrafficPair`: Object
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.EnableIpsecTrafficPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 
**pairId** | **int32** | ID for traffic pair | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableIpsecTrafficPairRequest struct via the builder pattern


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


## GetConnectedSubnets

> ConnectedSubnetsDetailResponse GetConnectedSubnets(ctx).ExtendedOutput(extendedOutput).Execute()

Get connected subnets



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
    extendedOutput := true // bool | Receive verbose information about resources (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.GetConnectedSubnets(context.Background()).ExtendedOutput(extendedOutput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.GetConnectedSubnets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectedSubnets`: ConnectedSubnetsDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.GetConnectedSubnets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectedSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extendedOutput** | **bool** | Receive verbose information about resources | [default to false]

### Return type

[**ConnectedSubnetsDetailResponse**](ConnectedSubnetsDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpsecDetails

> IpsecSystemDetail GetIpsecDetails(ctx).Execute()

Get IPsec details



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
    resp, r, err := apiClient.IPsecApi.GetIpsecDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.GetIpsecDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIpsecDetails`: IpsecSystemDetail
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.GetIpsecDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpsecDetailsRequest struct via the builder pattern


### Return type

[**IpsecSystemDetail**](IpsecSystemDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpsecEndpoint

> Object GetIpsecEndpoint(ctx, endpointId).Execute()

Get IPsec endpoint



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.GetIpsecEndpoint(context.Background(), endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.GetIpsecEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIpsecEndpoint`: Object
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.GetIpsecEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpsecEndpointRequest struct via the builder pattern


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


## GetIpsecLinkHistory

> LinkHistoryDetail GetIpsecLinkHistory(ctx).Remote(remote).Local(local).Tunnelid(tunnelid).Execute()

Get IPsec link history



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
    remote := "remote_example" // string | Address string in CIDR format to display link history to a remote endpoint. (optional)
    local := "local_example" // string | Address string in CIDR format which will display status of the local route (optional)
    tunnelid := int32(56) // int32 | Will display link history of just the tunnel specified, which may be only one tunnel to a remote endpoint. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.GetIpsecLinkHistory(context.Background()).Remote(remote).Local(local).Tunnelid(tunnelid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.GetIpsecLinkHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIpsecLinkHistory`: LinkHistoryDetail
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.GetIpsecLinkHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpsecLinkHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remote** | **string** | Address string in CIDR format to display link history to a remote endpoint. | 
 **local** | **string** | Address string in CIDR format which will display status of the local route | 
 **tunnelid** | **int32** | Will display link history of just the tunnel specified, which may be only one tunnel to a remote endpoint. | 

### Return type

[**LinkHistoryDetail**](LinkHistoryDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpsecStatus

> IpsecTunnelListResponseKeyValue GetIpsecStatus(ctx).UpDownStatusOnly(upDownStatusOnly).Execute()

Get IPsec status



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
    upDownStatusOnly := true // bool | Only retrieve tunnel status. True is more performant but has less info. Defaults to false. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.GetIpsecStatus(context.Background()).UpDownStatusOnly(upDownStatusOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.GetIpsecStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIpsecStatus`: IpsecTunnelListResponseKeyValue
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.GetIpsecStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIpsecStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upDownStatusOnly** | **bool** | Only retrieve tunnel status. True is more performant but has less info. Defaults to false. | [default to false]

### Return type

[**IpsecTunnelListResponseKeyValue**](IpsecTunnelListResponseKeyValue.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateIpsecEndpoint

> IpsecRemoteEndpointDetail PostCreateIpsecEndpoint(ctx).CreateIpsecEndpointRequest(createIpsecEndpointRequest).Execute()

Create IPsec endpoint



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
    createIpsecEndpointRequest := *openapiclient.NewCreateIpsecEndpointRequest("Name_example", "Ipaddress_example", "Secret_example") // CreateIpsecEndpointRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.PostCreateIpsecEndpoint(context.Background()).CreateIpsecEndpointRequest(createIpsecEndpointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.PostCreateIpsecEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateIpsecEndpoint`: IpsecRemoteEndpointDetail
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.PostCreateIpsecEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateIpsecEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIpsecEndpointRequest** | [**CreateIpsecEndpointRequest**](CreateIpsecEndpointRequest.md) |  | 

### Return type

[**IpsecRemoteEndpointDetail**](IpsecRemoteEndpointDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateIpsecEndpointTunnel

> Object PostCreateIpsecEndpointTunnel(ctx, endpointId).CreateIpsecTunnelRequest(createIpsecTunnelRequest).Execute()

Create IPsec endpoint tunnel



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
    createIpsecTunnelRequest := *openapiclient.NewCreateIpsecTunnelRequest("RemoteSubnet_example") // CreateIpsecTunnelRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.PostCreateIpsecEndpointTunnel(context.Background(), endpointId).CreateIpsecTunnelRequest(createIpsecTunnelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.PostCreateIpsecEndpointTunnel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateIpsecEndpointTunnel`: Object
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.PostCreateIpsecEndpointTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateIpsecEndpointTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIpsecTunnelRequest** | [**CreateIpsecTunnelRequest**](CreateIpsecTunnelRequest.md) |  | 

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


## PostCreateIpsecTrafficPair

> IpsecTrafficPairResponse PostCreateIpsecTrafficPair(ctx, endpointId).CreateIpsecTrafficPairRequest(createIpsecTrafficPairRequest).Execute()

Create IPsec traffic pair



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
    createIpsecTrafficPairRequest := *openapiclient.NewCreateIpsecTrafficPairRequest("RemoteSubnet_example", "LocalSubnet_example") // CreateIpsecTrafficPairRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.PostCreateIpsecTrafficPair(context.Background(), endpointId).CreateIpsecTrafficPairRequest(createIpsecTrafficPairRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.PostCreateIpsecTrafficPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateIpsecTrafficPair`: IpsecTrafficPairResponse
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.PostCreateIpsecTrafficPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateIpsecTrafficPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIpsecTrafficPairRequest** | [**CreateIpsecTrafficPairRequest**](CreateIpsecTrafficPairRequest.md) |  | 

### Return type

[**IpsecTrafficPairResponse**](IpsecTrafficPairResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRestartIpsecAction

> RestartStatus PostRestartIpsecAction(ctx).RestartRequest(restartRequest).Execute()

Restart ipsec subystem



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
    restartRequest := *openapiclient.NewRestartRequest(false) // RestartRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.PostRestartIpsecAction(context.Background()).RestartRequest(restartRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.PostRestartIpsecAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostRestartIpsecAction`: RestartStatus
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.PostRestartIpsecAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostRestartIpsecActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restartRequest** | [**RestartRequest**](RestartRequest.md) |  | 

### Return type

[**RestartStatus**](RestartStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateIpsecConfig

> Object PutUpdateIpsecConfig(ctx).UpdateIpsecAddressRequest(updateIpsecAddressRequest).Execute()

Update IPsec config



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
    updateIpsecAddressRequest := *openapiclient.NewUpdateIpsecAddressRequest("IpsecLocalIpaddress_example") // UpdateIpsecAddressRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.PutUpdateIpsecConfig(context.Background()).UpdateIpsecAddressRequest(updateIpsecAddressRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.PutUpdateIpsecConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateIpsecConfig`: Object
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.PutUpdateIpsecConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateIpsecConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateIpsecAddressRequest** | [**UpdateIpsecAddressRequest**](UpdateIpsecAddressRequest.md) |  | 

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


## PutUpdateIpsecEndpoint

> Object PutUpdateIpsecEndpoint(ctx, endpointId).UpdateIpsecEndpointRequest(updateIpsecEndpointRequest).Execute()

Update IPsec endpoint



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
    updateIpsecEndpointRequest := *openapiclient.NewUpdateIpsecEndpointRequest() // UpdateIpsecEndpointRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.PutUpdateIpsecEndpoint(context.Background(), endpointId).UpdateIpsecEndpointRequest(updateIpsecEndpointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.PutUpdateIpsecEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateIpsecEndpoint`: Object
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.PutUpdateIpsecEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateIpsecEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIpsecEndpointRequest** | [**UpdateIpsecEndpointRequest**](UpdateIpsecEndpointRequest.md) |  | 

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


## PutUpdateIpsecEndpointTunnel

> IpsecTunnelDetail PutUpdateIpsecEndpointTunnel(ctx, endpointId, tunnelId).UpdateIpsecTunnelRequest(updateIpsecTunnelRequest).Execute()

Update IPsec endpoint tunnel



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
    tunnelId := int32(56) // int32 | ID for tunnel
    updateIpsecTunnelRequest := *openapiclient.NewUpdateIpsecTunnelRequest() // UpdateIpsecTunnelRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.PutUpdateIpsecEndpointTunnel(context.Background(), endpointId, tunnelId).UpdateIpsecTunnelRequest(updateIpsecTunnelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.PutUpdateIpsecEndpointTunnel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateIpsecEndpointTunnel`: IpsecTunnelDetail
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.PutUpdateIpsecEndpointTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 
**tunnelId** | **int32** | ID for tunnel | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateIpsecEndpointTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIpsecTunnelRequest** | [**UpdateIpsecTunnelRequest**](UpdateIpsecTunnelRequest.md) |  | 

### Return type

[**IpsecTunnelDetail**](IpsecTunnelDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateIpsecTrafficPair

> Object PutUpdateIpsecTrafficPair(ctx, endpointId, pairId).UpdateIpsecTrafficPairRequest(updateIpsecTrafficPairRequest).Execute()

Update IPsec traffic pair



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
    pairId := int32(56) // int32 | ID for traffic pair
    updateIpsecTrafficPairRequest := *openapiclient.NewUpdateIpsecTrafficPairRequest() // UpdateIpsecTrafficPairRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IPsecApi.PutUpdateIpsecTrafficPair(context.Background(), endpointId, pairId).UpdateIpsecTrafficPairRequest(updateIpsecTrafficPairRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPsecApi.PutUpdateIpsecTrafficPair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateIpsecTrafficPair`: Object
    fmt.Fprintf(os.Stdout, "Response from `IPsecApi.PutUpdateIpsecTrafficPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**endpointId** | **int32** | ID for IPsec endpoint | 
**pairId** | **int32** | ID for traffic pair | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateIpsecTrafficPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIpsecTrafficPairRequest** | [**UpdateIpsecTrafficPairRequest**](UpdateIpsecTrafficPairRequest.md) |  | 

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

