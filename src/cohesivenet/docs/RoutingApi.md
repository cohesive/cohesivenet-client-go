# \RoutingApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRoute**](RoutingApi.md#DeleteRoute) | **Delete** /routes/{route_id} | Delete route
[**DisableRoute**](RoutingApi.md#DisableRoute) | **Put** /routes/{route_id}/disable | Disable route
[**EnableRoute**](RoutingApi.md#EnableRoute) | **Put** /routes/{route_id}/enable | Enable route
[**GetRoutes**](RoutingApi.md#GetRoutes) | **Get** /routes | Get routes
[**PostCreateRoute**](RoutingApi.md#PostCreateRoute) | **Post** /routes | Create route



## DeleteRoute

> Object DeleteRoute(ctx, routeId).Execute()

Delete route



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
    routeId := int32(56) // int32 | ID for Route

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.DeleteRoute(context.Background(), routeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.DeleteRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRoute`: Object
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.DeleteRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **int32** | ID for Route | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


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


## DisableRoute

> RouteDetailResponse DisableRoute(ctx, routeId).Execute()

Disable route



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
    routeId := int32(56) // int32 | ID for Route

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.DisableRoute(context.Background(), routeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.DisableRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableRoute`: RouteDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.DisableRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **int32** | ID for Route | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouteDetailResponse**](RouteDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableRoute

> Object EnableRoute(ctx, routeId).Execute()

Enable route



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
    routeId := int32(56) // int32 | ID for Route

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.EnableRoute(context.Background(), routeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.EnableRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableRoute`: Object
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.EnableRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeId** | **int32** | ID for Route | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableRouteRequest struct via the builder pattern


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


## GetRoutes

> RoutesListResponse GetRoutes(ctx).Table(table).Execute()

Get routes



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
    table := "table_example" // string | name of table to get routes for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.GetRoutes(context.Background()).Table(table).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.GetRoutes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoutes`: RoutesListResponse
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.GetRoutes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **table** | **string** | name of table to get routes for | 

### Return type

[**RoutesListResponse**](RoutesListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateRoute

> Object PostCreateRoute(ctx).CreateRouteRequest(createRouteRequest).Execute()

Create route



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
    createRouteRequest := *openapiclient.NewCreateRouteRequest("Cidr_example") // CreateRouteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoutingApi.PostCreateRoute(context.Background()).CreateRouteRequest(createRouteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingApi.PostCreateRoute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateRoute`: Object
    fmt.Fprintf(os.Stdout, "Response from `RoutingApi.PostCreateRoute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRouteRequest** | [**CreateRouteRequest**](CreateRouteRequest.md) |  | 

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

