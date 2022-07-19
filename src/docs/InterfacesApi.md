# \InterfacesApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInterfaceAddress**](InterfacesApi.md#CreateInterfaceAddress) | **Post** /interfaces/system/{interface_id}/addresses | Create system interface address
[**DeleteGreInterface**](InterfacesApi.md#DeleteGreInterface) | **Delete** /interfaces/edge_gre/{interface_id} | Delete GRE interface
[**DeleteInterfaceAddress**](InterfacesApi.md#DeleteInterfaceAddress) | **Delete** /interfaces/system/{interface_id}/addresses/{address_id} | Delete system interface address
[**DeleteSystemInterface**](InterfacesApi.md#DeleteSystemInterface) | **Delete** /interfaces/system/{interface_id} | Delete system interface
[**GetGreInterfaceDetails**](InterfacesApi.md#GetGreInterfaceDetails) | **Get** /interfaces/edge_gre/{interface_id} | Get GRE interface details
[**GetGreInterfaces**](InterfacesApi.md#GetGreInterfaces) | **Get** /interfaces/edge_gre | Get Edge GRE interfaces
[**GetInterfaceAddress**](InterfacesApi.md#GetInterfaceAddress) | **Get** /interfaces/system/{interface_id}/addresses/{address_id} | Get system interface address
[**GetInterfaceAddresses**](InterfacesApi.md#GetInterfaceAddresses) | **Get** /interfaces/system/{interface_id}/addresses | Get all system interface addresses
[**GetInterfaces**](InterfacesApi.md#GetInterfaces) | **Get** /interfaces | Get all interfaces
[**GetSystemInterfaceDetails**](InterfacesApi.md#GetSystemInterfaceDetails) | **Get** /interfaces/system/{interface_id} | Get system interface
[**GetSystemInterfaces**](InterfacesApi.md#GetSystemInterfaces) | **Get** /interfaces/system | Get all system interfaces
[**PostCreateGreInterface**](InterfacesApi.md#PostCreateGreInterface) | **Post** /interfaces/edge_gre | Create edge GRE interface
[**PostCreateSystemInterface**](InterfacesApi.md#PostCreateSystemInterface) | **Post** /interfaces/system | Create system interface
[**PostInterfacesAction**](InterfacesApi.md#PostInterfacesAction) | **Post** /interfaces/action | Take action on all interfaces
[**PutUpdateGreInterface**](InterfacesApi.md#PutUpdateGreInterface) | **Put** /interfaces/edge_gre/{interface_id} | Update GRE interface
[**PutUpdateInterfaceAddress**](InterfacesApi.md#PutUpdateInterfaceAddress) | **Put** /interfaces/system/{interface_id}/addresses/{address_id} | Update system interface address
[**PutUpdateSystemInterface**](InterfacesApi.md#PutUpdateSystemInterface) | **Put** /interfaces/system/{interface_id} | Update system interface



## CreateInterfaceAddress

> interface{} CreateInterfaceAddress(ctx, interfaceId).CreateInterfaceAddressRequest(createInterfaceAddressRequest).Execute()

Create system interface address



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
    interfaceId := "interfaceId_example" // string | ID for system interface
    createInterfaceAddressRequest := *openapiclient.NewCreateInterfaceAddressRequest("IpInternal_example") // CreateInterfaceAddressRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.CreateInterfaceAddress(context.Background(), interfaceId).CreateInterfaceAddressRequest(createInterfaceAddressRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.CreateInterfaceAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInterfaceAddress`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.CreateInterfaceAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | ID for system interface | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInterfaceAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInterfaceAddressRequest** | [**CreateInterfaceAddressRequest**](CreateInterfaceAddressRequest.md) |  | 

### Return type

**interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGreInterface

> Object DeleteGreInterface(ctx, interfaceId).Execute()

Delete GRE interface



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
    interfaceId := "interfaceId_example" // string | ID for system interface

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.DeleteGreInterface(context.Background(), interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.DeleteGreInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGreInterface`: Object
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.DeleteGreInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | ID for system interface | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGreInterfaceRequest struct via the builder pattern


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


## DeleteInterfaceAddress

> Object DeleteInterfaceAddress(ctx, interfaceId, addressId).Execute()

Delete system interface address



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
    interfaceId := "interfaceId_example" // string | ID for system interface
    addressId := int32(56) // int32 | ID for interface address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.DeleteInterfaceAddress(context.Background(), interfaceId, addressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.DeleteInterfaceAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteInterfaceAddress`: Object
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.DeleteInterfaceAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | ID for system interface | 
**addressId** | **int32** | ID for interface address | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInterfaceAddressRequest struct via the builder pattern


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


## DeleteSystemInterface

> Object DeleteSystemInterface(ctx, interfaceId).Execute()

Delete system interface



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
    interfaceId := "interfaceId_example" // string | ID for system interface

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.DeleteSystemInterface(context.Background(), interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.DeleteSystemInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSystemInterface`: Object
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.DeleteSystemInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | ID for system interface | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemInterfaceRequest struct via the builder pattern


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


## GetGreInterfaceDetails

> Object GetGreInterfaceDetails(ctx, interfaceId).Execute()

Get GRE interface details



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
    interfaceId := "interfaceId_example" // string | ID for system interface

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.GetGreInterfaceDetails(context.Background(), interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.GetGreInterfaceDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGreInterfaceDetails`: Object
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.GetGreInterfaceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | ID for system interface | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGreInterfaceDetailsRequest struct via the builder pattern


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


## GetGreInterfaces

> GREEndpointListResponse GetGreInterfaces(ctx).Execute()

Get Edge GRE interfaces



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
    resp, r, err := apiClient.InterfacesApi.GetGreInterfaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.GetGreInterfaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGreInterfaces`: GREEndpointListResponse
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.GetGreInterfaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGreInterfacesRequest struct via the builder pattern


### Return type

[**GREEndpointListResponse**](GREEndpointListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterfaceAddress

> Object GetInterfaceAddress(ctx, interfaceId, addressId).Execute()

Get system interface address



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
    interfaceId := "interfaceId_example" // string | ID for system interface
    addressId := int32(56) // int32 | ID for interface address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.GetInterfaceAddress(context.Background(), interfaceId, addressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.GetInterfaceAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterfaceAddress`: Object
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.GetInterfaceAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | ID for system interface | 
**addressId** | **int32** | ID for interface address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceAddressRequest struct via the builder pattern


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


## GetInterfaceAddresses

> InterfaceAddressListResponse GetInterfaceAddresses(ctx, interfaceId).Execute()

Get all system interface addresses



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
    interfaceId := "interfaceId_example" // string | ID for system interface

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.GetInterfaceAddresses(context.Background(), interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.GetInterfaceAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterfaceAddresses`: InterfaceAddressListResponse
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.GetInterfaceAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | ID for system interface | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InterfaceAddressListResponse**](InterfaceAddressListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInterfaces

> SystemInterfaceListResponse GetInterfaces(ctx).Execute()

Get all interfaces



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
    resp, r, err := apiClient.InterfacesApi.GetInterfaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.GetInterfaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterfaces`: SystemInterfaceListResponse
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.GetInterfaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfacesRequest struct via the builder pattern


### Return type

[**SystemInterfaceListResponse**](SystemInterfaceListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemInterfaceDetails

> Object GetSystemInterfaceDetails(ctx, interfaceId).Execute()

Get system interface



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
    interfaceId := "interfaceId_example" // string | ID for system interface

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.GetSystemInterfaceDetails(context.Background(), interfaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.GetSystemInterfaceDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemInterfaceDetails`: Object
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.GetSystemInterfaceDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | ID for system interface | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemInterfaceDetailsRequest struct via the builder pattern


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


## GetSystemInterfaces

> Object GetSystemInterfaces(ctx).Execute()

Get all system interfaces



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
    resp, r, err := apiClient.InterfacesApi.GetSystemInterfaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.GetSystemInterfaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemInterfaces`: Object
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.GetSystemInterfaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemInterfacesRequest struct via the builder pattern


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


## PostCreateGreInterface

> GREEndpointDetail PostCreateGreInterface(ctx).PostCreateGreInterfaceRequest(postCreateGreInterfaceRequest).Execute()

Create edge GRE interface



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
    postCreateGreInterfaceRequest := *openapiclient.NewPostCreateGreInterfaceRequest() // PostCreateGreInterfaceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.PostCreateGreInterface(context.Background()).PostCreateGreInterfaceRequest(postCreateGreInterfaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.PostCreateGreInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateGreInterface`: GREEndpointDetail
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.PostCreateGreInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateGreInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postCreateGreInterfaceRequest** | [**PostCreateGreInterfaceRequest**](PostCreateGreInterfaceRequest.md) |  | 

### Return type

[**GREEndpointDetail**](GREEndpointDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateSystemInterface

> SystemInterfaceDetail PostCreateSystemInterface(ctx).PostCreateSystemInterfaceRequest(postCreateSystemInterfaceRequest).Execute()

Create system interface



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
    postCreateSystemInterfaceRequest := *openapiclient.NewPostCreateSystemInterfaceRequest("Name_example") // PostCreateSystemInterfaceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.PostCreateSystemInterface(context.Background()).PostCreateSystemInterfaceRequest(postCreateSystemInterfaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.PostCreateSystemInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateSystemInterface`: SystemInterfaceDetail
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.PostCreateSystemInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateSystemInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postCreateSystemInterfaceRequest** | [**PostCreateSystemInterfaceRequest**](PostCreateSystemInterfaceRequest.md) |  | 

### Return type

[**SystemInterfaceDetail**](SystemInterfaceDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInterfacesAction

> Object PostInterfacesAction(ctx).InterfaceActionRequest(interfaceActionRequest).Execute()

Take action on all interfaces



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
    interfaceActionRequest := *openapiclient.NewInterfaceActionRequest() // InterfaceActionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.PostInterfacesAction(context.Background()).InterfaceActionRequest(interfaceActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.PostInterfacesAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostInterfacesAction`: Object
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.PostInterfacesAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostInterfacesActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interfaceActionRequest** | [**InterfaceActionRequest**](InterfaceActionRequest.md) |  | 

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


## PutUpdateGreInterface

> Object PutUpdateGreInterface(ctx, interfaceId).PostCreateGreInterfaceRequest(postCreateGreInterfaceRequest).Execute()

Update GRE interface



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
    interfaceId := "interfaceId_example" // string | ID for system interface
    postCreateGreInterfaceRequest := *openapiclient.NewPostCreateGreInterfaceRequest() // PostCreateGreInterfaceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.PutUpdateGreInterface(context.Background(), interfaceId).PostCreateGreInterfaceRequest(postCreateGreInterfaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.PutUpdateGreInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateGreInterface`: Object
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.PutUpdateGreInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | ID for system interface | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateGreInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postCreateGreInterfaceRequest** | [**PostCreateGreInterfaceRequest**](PostCreateGreInterfaceRequest.md) |  | 

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


## PutUpdateInterfaceAddress

> Object PutUpdateInterfaceAddress(ctx, interfaceId, addressId).PutUpdateInterfaceAddressRequest(putUpdateInterfaceAddressRequest).Execute()

Update system interface address



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
    interfaceId := "interfaceId_example" // string | ID for system interface
    addressId := int32(56) // int32 | ID for interface address
    putUpdateInterfaceAddressRequest := *openapiclient.NewPutUpdateInterfaceAddressRequest() // PutUpdateInterfaceAddressRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.PutUpdateInterfaceAddress(context.Background(), interfaceId, addressId).PutUpdateInterfaceAddressRequest(putUpdateInterfaceAddressRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.PutUpdateInterfaceAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateInterfaceAddress`: Object
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.PutUpdateInterfaceAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | ID for system interface | 
**addressId** | **int32** | ID for interface address | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateInterfaceAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putUpdateInterfaceAddressRequest** | [**PutUpdateInterfaceAddressRequest**](PutUpdateInterfaceAddressRequest.md) |  | 

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


## PutUpdateSystemInterface

> Object PutUpdateSystemInterface(ctx, interfaceId).PutUpdateSystemInterfaceRequest(putUpdateSystemInterfaceRequest).Execute()

Update system interface



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
    interfaceId := "interfaceId_example" // string | ID for system interface
    putUpdateSystemInterfaceRequest := *openapiclient.NewPutUpdateSystemInterfaceRequest() // PutUpdateSystemInterfaceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InterfacesApi.PutUpdateSystemInterface(context.Background(), interfaceId).PutUpdateSystemInterfaceRequest(putUpdateSystemInterfaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InterfacesApi.PutUpdateSystemInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateSystemInterface`: Object
    fmt.Fprintf(os.Stdout, "Response from `InterfacesApi.PutUpdateSystemInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**interfaceId** | **string** | ID for system interface | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateSystemInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putUpdateSystemInterfaceRequest** | [**PutUpdateSystemInterfaceRequest**](PutUpdateSystemInterfaceRequest.md) |  | 

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

