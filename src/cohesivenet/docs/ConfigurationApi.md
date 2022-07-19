# \ConfigurationApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomVariable**](ConfigurationApi.md#DeleteCustomVariable) | **Delete** /system/variables/{name} | Delete custom variable from system
[**GetConfig**](ConfigurationApi.md#GetConfig) | **Get** /config | Get configuration details
[**GetKeyset**](ConfigurationApi.md#GetKeyset) | **Get** /keyset | Get topology keyset
[**GetLicense**](ConfigurationApi.md#GetLicense) | **Get** /license | Get license details
[**GetMsConfig**](ConfigurationApi.md#GetMsConfig) | **Get** /ms | Get MS configuration
[**GetSSLCerts**](ConfigurationApi.md#GetSSLCerts) | **Get** /system/ssl | Get SSL Certificates
[**GetSslInstallStatus**](ConfigurationApi.md#GetSslInstallStatus) | **Get** /system/ssl/install/{uuid} | Get SSL installation status
[**GetVariableCollections**](ConfigurationApi.md#GetVariableCollections) | **Get** /system/variable-collections | Get system variable collection details
[**GetVariables**](ConfigurationApi.md#GetVariables) | **Get** /system/variables | Get system variables grouped by collections
[**PostCreateCustomVariable**](ConfigurationApi.md#PostCreateCustomVariable) | **Post** /system/variables | create custom variable for system
[**PostSendTestMsAlert**](ConfigurationApi.md#PostSendTestMsAlert) | **Post** /ms/alert/test | Send test VNS3:ms alert
[**PostSetMsConfig**](ConfigurationApi.md#PostSetMsConfig) | **Post** /ms | Set MS for controller
[**PutConfig**](ConfigurationApi.md#PutConfig) | **Put** /config | Update configuration
[**PutInstallSslKeypair**](ConfigurationApi.md#PutInstallSslKeypair) | **Put** /system/ssl/install | Install SSL cert and key pair
[**PutKeyset**](ConfigurationApi.md#PutKeyset) | **Put** /keyset | Generate keyset
[**PutLicenseUpgrade**](ConfigurationApi.md#PutLicenseUpgrade) | **Put** /license/upgrade | Upgrade license
[**PutSetLicenseParameters**](ConfigurationApi.md#PutSetLicenseParameters) | **Put** /license/parameters | Set license parameters
[**PutUpdateAdminUi**](ConfigurationApi.md#PutUpdateAdminUi) | **Put** /admin_ui | Update admin UI settings
[**PutUpdateApiPassword**](ConfigurationApi.md#PutUpdateApiPassword) | **Put** /api_password | Update API password
[**PutUpdateCustomVariable**](ConfigurationApi.md#PutUpdateCustomVariable) | **Put** /system/variables/{name} | Update custom variable value
[**PutUploadSslKeypair**](ConfigurationApi.md#PutUploadSslKeypair) | **Put** /system/ssl/keypair | Upload new SSL cert and key pair
[**UpdateMsConfig**](ConfigurationApi.md#UpdateMsConfig) | **Put** /ms | Update MS config for controller
[**UploadLicense**](ConfigurationApi.md#UploadLicense) | **Put** /license | Upload license



## DeleteCustomVariable

> Object DeleteCustomVariable(ctx, name).Execute()

Delete custom variable from system



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
    name := "name_example" // string | name of variable

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.DeleteCustomVariable(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.DeleteCustomVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCustomVariable`: Object
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.DeleteCustomVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomVariableRequest struct via the builder pattern


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


## GetConfig

> ConfigDetail GetConfig(ctx).Execute()

Get configuration details



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
    resp, r, err := apiClient.ConfigurationApi.GetConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfig`: ConfigDetail
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


### Return type

[**ConfigDetail**](ConfigDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyset

> KeysetDetail GetKeyset(ctx).Execute()

Get topology keyset



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
    resp, r, err := apiClient.ConfigurationApi.GetKeyset(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetKeyset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyset`: KeysetDetail
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetKeyset`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeysetRequest struct via the builder pattern


### Return type

[**KeysetDetail**](KeysetDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicense

> LicenseDetail GetLicense(ctx).Execute()

Get license details



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
    resp, r, err := apiClient.ConfigurationApi.GetLicense(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicense`: LicenseDetail
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetLicense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseRequest struct via the builder pattern


### Return type

[**LicenseDetail**](LicenseDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMsConfig

> AlertDetailResponse GetMsConfig(ctx).Ip(ip).Execute()

Get MS configuration



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
    ip := "ip_example" // string | name of resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.GetMsConfig(context.Background()).Ip(ip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetMsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMsConfig`: AlertDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetMsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ip** | **string** | name of resource | 

### Return type

[**AlertDetailResponse**](AlertDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSLCerts

> SSLCertsResponse GetSSLCerts(ctx).Execute()

Get SSL Certificates



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
    resp, r, err := apiClient.ConfigurationApi.GetSSLCerts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetSSLCerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSLCerts`: SSLCertsResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetSSLCerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSLCertsRequest struct via the builder pattern


### Return type

[**SSLCertsResponse**](SSLCertsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSslInstallStatus

> Object GetSslInstallStatus(ctx, uuid).Execute()

Get SSL installation status



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
    uuid := "uuid_example" // string | uuid of resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.GetSslInstallStatus(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetSslInstallStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSslInstallStatus`: Object
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetSslInstallStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSslInstallStatusRequest struct via the builder pattern


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


## GetVariableCollections

> VariableCollectionsListResponse GetVariableCollections(ctx).Collections(collections).Execute()

Get system variable collection details



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
    collections := "collections_example" // string | filter variables by collections (accepts csv A,B,C) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.GetVariableCollections(context.Background()).Collections(collections).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetVariableCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariableCollections`: VariableCollectionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetVariableCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collections** | **string** | filter variables by collections (accepts csv A,B,C) | 

### Return type

[**VariableCollectionsListResponse**](VariableCollectionsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariables

> VariablesListResponse GetVariables(ctx).Execute()

Get system variables grouped by collections



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
    resp, r, err := apiClient.ConfigurationApi.GetVariables(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.GetVariables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVariables`: VariablesListResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.GetVariables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariablesRequest struct via the builder pattern


### Return type

[**VariablesListResponse**](VariablesListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateCustomVariable

> Object PostCreateCustomVariable(ctx).CreateCustomVariableRequest(createCustomVariableRequest).Execute()

create custom variable for system



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
    createCustomVariableRequest := *openapiclient.NewCreateCustomVariableRequest("Name_example", "Value_example") // CreateCustomVariableRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.PostCreateCustomVariable(context.Background()).CreateCustomVariableRequest(createCustomVariableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PostCreateCustomVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateCustomVariable`: Object
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PostCreateCustomVariable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateCustomVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomVariableRequest** | [**CreateCustomVariableRequest**](CreateCustomVariableRequest.md) |  | 

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


## PostSendTestMsAlert

> SimpleBooleanResponse PostSendTestMsAlert(ctx).Execute()

Send test VNS3:ms alert



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
    resp, r, err := apiClient.ConfigurationApi.PostSendTestMsAlert(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PostSendTestMsAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSendTestMsAlert`: SimpleBooleanResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PostSendTestMsAlert`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSendTestMsAlertRequest struct via the builder pattern


### Return type

[**SimpleBooleanResponse**](SimpleBooleanResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSetMsConfig

> MSConfig PostSetMsConfig(ctx).SetMSRequest1(setMSRequest1).Execute()

Set MS for controller



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
    setMSRequest1 := *openapiclient.NewSetMSRequest1("Ip_example") // SetMSRequest1 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.PostSetMsConfig(context.Background()).SetMSRequest1(setMSRequest1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PostSetMsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSetMsConfig`: MSConfig
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PostSetMsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSetMsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setMSRequest1** | [**SetMSRequest1**](SetMSRequest1.md) |  | 

### Return type

[**MSConfig**](MSConfig.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConfig

> Object PutConfig(ctx).UpdateConfigRequest(updateConfigRequest).Execute()

Update configuration



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
    updateConfigRequest := *openapiclient.NewUpdateConfigRequest() // UpdateConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.PutConfig(context.Background()).UpdateConfigRequest(updateConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PutConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutConfig`: Object
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PutConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigRequest** | [**UpdateConfigRequest**](UpdateConfigRequest.md) |  | 

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


## PutInstallSslKeypair

> ServerSSLDetailResponse PutInstallSslKeypair(ctx).Execute()

Install SSL cert and key pair



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
    resp, r, err := apiClient.ConfigurationApi.PutInstallSslKeypair(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PutInstallSslKeypair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutInstallSslKeypair`: ServerSSLDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PutInstallSslKeypair`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPutInstallSslKeypairRequest struct via the builder pattern


### Return type

[**ServerSSLDetailResponse**](ServerSSLDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutKeyset

> Object PutKeyset(ctx).UpdateKeysetRequest(updateKeysetRequest).Execute()

Generate keyset



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
    updateKeysetRequest := *openapiclient.NewUpdateKeysetRequest("Token_example") // UpdateKeysetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.PutKeyset(context.Background()).UpdateKeysetRequest(updateKeysetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PutKeyset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutKeyset`: Object
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PutKeyset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutKeysetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateKeysetRequest** | [**UpdateKeysetRequest**](UpdateKeysetRequest.md) |  | 

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


## PutLicenseUpgrade

> UpgradeLicenseResponse PutLicenseUpgrade(ctx).Body(body).Execute()

Upgrade license



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
    body := os.NewFile(1234, "some_file") // *os.File | License file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.PutLicenseUpgrade(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PutLicenseUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutLicenseUpgrade`: UpgradeLicenseResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PutLicenseUpgrade`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLicenseUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | License file | 

### Return type

[**UpgradeLicenseResponse**](UpgradeLicenseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSetLicenseParameters

> LicenseParametersDetail PutSetLicenseParameters(ctx).PutLicenseParametersRequest(putLicenseParametersRequest).Execute()

Set license parameters



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
    putLicenseParametersRequest := *openapiclient.NewPutLicenseParametersRequest() // PutLicenseParametersRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.PutSetLicenseParameters(context.Background()).PutLicenseParametersRequest(putLicenseParametersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PutSetLicenseParameters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSetLicenseParameters`: LicenseParametersDetail
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PutSetLicenseParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutSetLicenseParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putLicenseParametersRequest** | [**PutLicenseParametersRequest**](PutLicenseParametersRequest.md) |  | 

### Return type

[**LicenseParametersDetail**](LicenseParametersDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateAdminUi

> AdminUISettingsDetail PutUpdateAdminUi(ctx).UpdateAdminUISettingsRequest(updateAdminUISettingsRequest).Execute()

Update admin UI settings



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
    updateAdminUISettingsRequest := *openapiclient.NewUpdateAdminUISettingsRequest() // UpdateAdminUISettingsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.PutUpdateAdminUi(context.Background()).UpdateAdminUISettingsRequest(updateAdminUISettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PutUpdateAdminUi``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateAdminUi`: AdminUISettingsDetail
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PutUpdateAdminUi`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateAdminUiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateAdminUISettingsRequest** | [**UpdateAdminUISettingsRequest**](UpdateAdminUISettingsRequest.md) |  | 

### Return type

[**AdminUISettingsDetail**](AdminUISettingsDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateApiPassword

> PasswordResetResponse PutUpdateApiPassword(ctx).UpdatePasswordRequest(updatePasswordRequest).Execute()

Update API password



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
    updatePasswordRequest := *openapiclient.NewUpdatePasswordRequest() // UpdatePasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.PutUpdateApiPassword(context.Background()).UpdatePasswordRequest(updatePasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PutUpdateApiPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateApiPassword`: PasswordResetResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PutUpdateApiPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateApiPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatePasswordRequest** | [**UpdatePasswordRequest**](UpdatePasswordRequest.md) |  | 

### Return type

[**PasswordResetResponse**](PasswordResetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateCustomVariable

> VariableDetailResponse PutUpdateCustomVariable(ctx, name).UpdateCustomVariableRequest(updateCustomVariableRequest).Execute()

Update custom variable value



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
    name := "name_example" // string | name of variable
    updateCustomVariableRequest := *openapiclient.NewUpdateCustomVariableRequest() // UpdateCustomVariableRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.PutUpdateCustomVariable(context.Background(), name).UpdateCustomVariableRequest(updateCustomVariableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PutUpdateCustomVariable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateCustomVariable`: VariableDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PutUpdateCustomVariable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name of variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateCustomVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomVariableRequest** | [**UpdateCustomVariableRequest**](UpdateCustomVariableRequest.md) |  | 

### Return type

[**VariableDetailResponse**](VariableDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUploadSslKeypair

> Object PutUploadSslKeypair(ctx).UpdateServerSSLRequest(updateServerSSLRequest).Execute()

Upload new SSL cert and key pair



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
    updateServerSSLRequest := *openapiclient.NewUpdateServerSSLRequest("Cert_example", "Key_example") // UpdateServerSSLRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.PutUploadSslKeypair(context.Background()).UpdateServerSSLRequest(updateServerSSLRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.PutUploadSslKeypair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUploadSslKeypair`: Object
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.PutUploadSslKeypair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutUploadSslKeypairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateServerSSLRequest** | [**UpdateServerSSLRequest**](UpdateServerSSLRequest.md) |  | 

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


## UpdateMsConfig

> Object UpdateMsConfig(ctx).SetMSRequest(setMSRequest).Execute()

Update MS config for controller



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
    setMSRequest := *openapiclient.NewSetMSRequest(false) // SetMSRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.UpdateMsConfig(context.Background()).SetMSRequest(setMSRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.UpdateMsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMsConfig`: Object
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.UpdateMsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setMSRequest** | [**SetMSRequest**](SetMSRequest.md) |  | 

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


## UploadLicense

> InitLicenseDetail UploadLicense(ctx).Body(body).Execute()

Upload license



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
    body := os.NewFile(1234, "some_file") // *os.File | License file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigurationApi.UploadLicense(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationApi.UploadLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadLicense`: InitLicenseDetail
    fmt.Fprintf(os.Stdout, "Response from `ConfigurationApi.UploadLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | License file | 

### Return type

[**InitLicenseDetail**](InitLicenseDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

