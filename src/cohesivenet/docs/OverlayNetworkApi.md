# \OverlayNetworkApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClientpackTag**](OverlayNetworkApi.md#DeleteClientpackTag) | **Delete** /clientpack/{clientpack_name} | Delete clientpack tag
[**ExportFirewallRules**](OverlayNetworkApi.md#ExportFirewallRules) | **Get** /v2/firewall/rules/export | V2 Export firewall rules
[**GetClientpack**](OverlayNetworkApi.md#GetClientpack) | **Get** /clientpacks/{clientpack_name} | Get clientpack details
[**GetClientpacks**](OverlayNetworkApi.md#GetClientpacks) | **Get** /clientpacks | Get clientpacks
[**GetClientsStatus**](OverlayNetworkApi.md#GetClientsStatus) | **Get** /status/clients | Get clients status
[**GetDownloadClientpack**](OverlayNetworkApi.md#GetDownloadClientpack) | **Get** /clientpack | Download clientpack
[**GetDownloadNamedClientpack**](OverlayNetworkApi.md#GetDownloadNamedClientpack) | **Get** /clientpack/{clientpack_name} | Download clientpack by name
[**PostAddClientpacks**](OverlayNetworkApi.md#PostAddClientpacks) | **Post** /clientpacks/add_clientpacks | Create new clientpack
[**PostCheckoutNextClientpack**](OverlayNetworkApi.md#PostCheckoutNextClientpack) | **Post** /clientpacks/next_available | Checkout next clientpack
[**PostCreateClientpackTag**](OverlayNetworkApi.md#PostCreateClientpackTag) | **Post** /clientpack/{clientpack_name} | Create clientpack tag
[**PostResetAllClients**](OverlayNetworkApi.md#PostResetAllClients) | **Post** /clients/reset_all | Reset all clients
[**PostResetClient**](OverlayNetworkApi.md#PostResetClient) | **Post** /client/reset | Reset client
[**PutDisconnectClientpack**](OverlayNetworkApi.md#PutDisconnectClientpack) | **Put** /clientpack/{clientpack_name} | Disconnect clientpack
[**PutUpdateAllClientpacks**](OverlayNetworkApi.md#PutUpdateAllClientpacks) | **Put** /clientpacks | Update all clientpacks
[**PutUpdateClientpack**](OverlayNetworkApi.md#PutUpdateClientpack) | **Put** /clientpack | Update clientpack



## DeleteClientpackTag

> Object DeleteClientpackTag(ctx, clientpackName).ClientpackTagKeyRequest(clientpackTagKeyRequest).Execute()

Delete clientpack tag



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
    clientpackName := "clientpackName_example" // string | name of clientpack
    clientpackTagKeyRequest := *openapiclient.NewClientpackTagKeyRequest("Key_example") // ClientpackTagKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.DeleteClientpackTag(context.Background(), clientpackName).ClientpackTagKeyRequest(clientpackTagKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.DeleteClientpackTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClientpackTag`: Object
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.DeleteClientpackTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientpackName** | **string** | name of clientpack | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientpackTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientpackTagKeyRequest** | [**ClientpackTagKeyRequest**](ClientpackTagKeyRequest.md) |  | 

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


## ExportFirewallRules

> *os.File ExportFirewallRules(ctx).Group(group).Tables(tables).Type_(type_).Execute()

V2 Export firewall rules



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
    group := "group_example" // string | Export group only (optional)
    tables := "tables_example" // string | Tables to export. Accepts csv. Defaults to all. (optional)
    type_ := "type__example" // string | One of json or plaintext (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.ExportFirewallRules(context.Background()).Group(group).Tables(tables).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.ExportFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFirewallRules`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.ExportFirewallRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | **string** | Export group only | 
 **tables** | **string** | Tables to export. Accepts csv. Defaults to all. | 
 **type_** | **string** | One of json or plaintext | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientpack

> ClientpackDetailResponse GetClientpack(ctx, clientpackName).Execute()

Get clientpack details



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
    clientpackName := "clientpackName_example" // string | name of clientpack

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.GetClientpack(context.Background(), clientpackName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.GetClientpack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientpack`: ClientpackDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.GetClientpack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientpackName** | **string** | name of clientpack | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientpackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientpackDetailResponse**](ClientpackDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientpacks

> ClientpackListResponse GetClientpacks(ctx).Sorted(sorted).Execute()

Get clientpacks



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
    sorted := true // bool | Sort resources (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.GetClientpacks(context.Background()).Sorted(sorted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.GetClientpacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientpacks`: ClientpackListResponse
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.GetClientpacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientpacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sorted** | **bool** | Sort resources | [default to false]

### Return type

[**ClientpackListResponse**](ClientpackListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsStatus

> OverlayClientsListResponse GetClientsStatus(ctx).Execute()

Get clients status



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
    resp, r, err := apiClient.OverlayNetworkApi.GetClientsStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.GetClientsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientsStatus`: OverlayClientsListResponse
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.GetClientsStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsStatusRequest struct via the builder pattern


### Return type

[**OverlayClientsListResponse**](OverlayClientsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDownloadClientpack

> *os.File GetDownloadClientpack(ctx).Name(name).Fileformat(fileformat).Execute()

Download clientpack



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
    name := "name_example" // string | name of clientpack. Typical IP address with underscores. e.g. 100_127_255_200.
    fileformat := "fileformat_example" // string | One of tarball, tar.gz, zip, conf, ovpn

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.GetDownloadClientpack(context.Background()).Name(name).Fileformat(fileformat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.GetDownloadClientpack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadClientpack`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.GetDownloadClientpack`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadClientpackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | name of clientpack. Typical IP address with underscores. e.g. 100_127_255_200. | 
 **fileformat** | **string** | One of tarball, tar.gz, zip, conf, ovpn | 

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


## GetDownloadNamedClientpack

> *os.File GetDownloadNamedClientpack(ctx, clientpackName).Execute()

Download clientpack by name



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
    clientpackName := "clientpackName_example" // string | name of clientpack

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.GetDownloadNamedClientpack(context.Background(), clientpackName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.GetDownloadNamedClientpack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadNamedClientpack`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.GetDownloadNamedClientpack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientpackName** | **string** | name of clientpack | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadNamedClientpackRequest struct via the builder pattern


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


## PostAddClientpacks

> Object PostAddClientpacks(ctx).AddClientpackRequest(addClientpackRequest).Execute()

Create new clientpack



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
    addClientpackRequest := *openapiclient.NewAddClientpackRequest("RequestedIps_example") // AddClientpackRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.PostAddClientpacks(context.Background()).AddClientpackRequest(addClientpackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.PostAddClientpacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAddClientpacks`: Object
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.PostAddClientpacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAddClientpacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addClientpackRequest** | [**AddClientpackRequest**](AddClientpackRequest.md) |  | 

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


## PostCheckoutNextClientpack

> Object PostCheckoutNextClientpack(ctx).CalculateNextClientpackRequest(calculateNextClientpackRequest).Execute()

Checkout next clientpack



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
    calculateNextClientpackRequest := *openapiclient.NewCalculateNextClientpackRequest() // CalculateNextClientpackRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.PostCheckoutNextClientpack(context.Background()).CalculateNextClientpackRequest(calculateNextClientpackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.PostCheckoutNextClientpack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCheckoutNextClientpack`: Object
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.PostCheckoutNextClientpack`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCheckoutNextClientpackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculateNextClientpackRequest** | [**CalculateNextClientpackRequest**](CalculateNextClientpackRequest.md) |  | 

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


## PostCreateClientpackTag

> ClientpackTagsResponse PostCreateClientpackTag(ctx, clientpackName).CreateClientpackTagRequest(createClientpackTagRequest).Execute()

Create clientpack tag



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
    clientpackName := "clientpackName_example" // string | name of clientpack
    createClientpackTagRequest := *openapiclient.NewCreateClientpackTagRequest("Key_example", "Value_example") // CreateClientpackTagRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.PostCreateClientpackTag(context.Background(), clientpackName).CreateClientpackTagRequest(createClientpackTagRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.PostCreateClientpackTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateClientpackTag`: ClientpackTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.PostCreateClientpackTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientpackName** | **string** | name of clientpack | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateClientpackTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createClientpackTagRequest** | [**CreateClientpackTagRequest**](CreateClientpackTagRequest.md) |  | 

### Return type

[**ClientpackTagsResponse**](ClientpackTagsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostResetAllClients

> BulkClientResetStatusResponse PostResetAllClients(ctx).Execute()

Reset all clients



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
    resp, r, err := apiClient.OverlayNetworkApi.PostResetAllClients(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.PostResetAllClients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostResetAllClients`: BulkClientResetStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.PostResetAllClients`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostResetAllClientsRequest struct via the builder pattern


### Return type

[**BulkClientResetStatusResponse**](BulkClientResetStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostResetClient

> ClientpackStatusResponse PostResetClient(ctx).ResetOverlayClientRequest(resetOverlayClientRequest).Execute()

Reset client



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
    resetOverlayClientRequest := *openapiclient.NewResetOverlayClientRequest("Name_example") // ResetOverlayClientRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.PostResetClient(context.Background()).ResetOverlayClientRequest(resetOverlayClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.PostResetClient``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostResetClient`: ClientpackStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.PostResetClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostResetClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetOverlayClientRequest** | [**ResetOverlayClientRequest**](ResetOverlayClientRequest.md) |  | 

### Return type

[**ClientpackStatusResponse**](ClientpackStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDisconnectClientpack

> Object PutDisconnectClientpack(ctx, clientpackName).DisconnetClientRequest(disconnetClientRequest).Execute()

Disconnect clientpack



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
    clientpackName := "clientpackName_example" // string | name of clientpack
    disconnetClientRequest := *openapiclient.NewDisconnetClientRequest(false) // DisconnetClientRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.PutDisconnectClientpack(context.Background(), clientpackName).DisconnetClientRequest(disconnetClientRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.PutDisconnectClientpack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutDisconnectClientpack`: Object
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.PutDisconnectClientpack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientpackName** | **string** | name of clientpack | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDisconnectClientpackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disconnetClientRequest** | [**DisconnetClientRequest**](DisconnetClientRequest.md) |  | 

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


## PutUpdateAllClientpacks

> UpdateClientpacksStatusResponse PutUpdateAllClientpacks(ctx).UpdateConfigClientpackRequest(updateConfigClientpackRequest).Execute()

Update all clientpacks



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
    updateConfigClientpackRequest := *openapiclient.NewUpdateConfigClientpackRequest() // UpdateConfigClientpackRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.PutUpdateAllClientpacks(context.Background()).UpdateConfigClientpackRequest(updateConfigClientpackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.PutUpdateAllClientpacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateAllClientpacks`: UpdateClientpacksStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.PutUpdateAllClientpacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateAllClientpacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigClientpackRequest** | [**UpdateConfigClientpackRequest**](UpdateConfigClientpackRequest.md) |  | 

### Return type

[**UpdateClientpacksStatusResponse**](UpdateClientpacksStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateClientpack

> UpdateClientpack PutUpdateClientpack(ctx).UpdateClientpackRequest(updateClientpackRequest).Execute()

Update clientpack



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
    updateClientpackRequest := *openapiclient.NewUpdateClientpackRequest("Name_example") // UpdateClientpackRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OverlayNetworkApi.PutUpdateClientpack(context.Background()).UpdateClientpackRequest(updateClientpackRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OverlayNetworkApi.PutUpdateClientpack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateClientpack`: UpdateClientpack
    fmt.Fprintf(os.Stdout, "Response from `OverlayNetworkApi.PutUpdateClientpack`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateClientpackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateClientpackRequest** | [**UpdateClientpackRequest**](UpdateClientpackRequest.md) |  | 

### Return type

[**UpdateClientpack**](UpdateClientpack.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

