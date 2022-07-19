# \AccessApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessUrl**](AccessApi.md#CreateAccessUrl) | **Post** /access/url | Create access URL
[**CreateApiToken**](AccessApi.md#CreateApiToken) | **Post** /access/token | Create API token
[**DeleteAccessUrl**](AccessApi.md#DeleteAccessUrl) | **Delete** /access/url/{access_url_id} | Delete access URL
[**DeleteAccessUrlBySearch**](AccessApi.md#DeleteAccessUrlBySearch) | **Delete** /access/url | Find and delete access URL
[**DeleteApiToken**](AccessApi.md#DeleteApiToken) | **Delete** /access/token/{token_id} | Delete API token
[**GetAccessUrl**](AccessApi.md#GetAccessUrl) | **Get** /access/url/{access_url_id} | Get access URL
[**GetAccessUrls**](AccessApi.md#GetAccessUrls) | **Get** /access/urls | Get access URLs
[**GetApiToken**](AccessApi.md#GetApiToken) | **Get** /access/token/{token_id} | Get API access token
[**GetApiTokens**](AccessApi.md#GetApiTokens) | **Get** /access/tokens | Get API access tokens
[**GetIdentityControllerSettings**](AccessApi.md#GetIdentityControllerSettings) | **Get** /identity/controller | Get identity Settings for VPN Users
[**GetIdentityVPNSettings**](AccessApi.md#GetIdentityVPNSettings) | **Get** /identity/vpn | Get identity Settings for VPN Users
[**PostTestIdentityControllerSettings**](AccessApi.md#PostTestIdentityControllerSettings) | **Post** /identity/controller/test | Test Controller Identity Settings. Currently only LDAP testing is supported.
[**PostTestIdentityVPNSettings**](AccessApi.md#PostTestIdentityVPNSettings) | **Post** /identity/vpn/test | Test VPN Identity Settings
[**PutExpireAccessUrl**](AccessApi.md#PutExpireAccessUrl) | **Put** /access/url/{access_url_id} | Expire access URL
[**PutExpireApiToken**](AccessApi.md#PutExpireApiToken) | **Put** /access/token/{token_id} | Expire API token
[**PutIdentityControllerSettings**](AccessApi.md#PutIdentityControllerSettings) | **Put** /identity/controller | Update Controller Identity settings
[**PutIdentityVPNSettings**](AccessApi.md#PutIdentityVPNSettings) | **Put** /identity/vpn | Update VPN Identity settings



## CreateAccessUrl

> AccessUrlDetail CreateAccessUrl(ctx).CreateAccessURLRequest(createAccessURLRequest).Execute()

Create access URL



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
    createAccessURLRequest := *openapiclient.NewCreateAccessURLRequest() // CreateAccessURLRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.CreateAccessUrl(context.Background()).CreateAccessURLRequest(createAccessURLRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.CreateAccessUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessUrl`: AccessUrlDetail
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.CreateAccessUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccessURLRequest** | [**CreateAccessURLRequest**](CreateAccessURLRequest.md) |  | 

### Return type

[**AccessUrlDetail**](AccessUrlDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApiToken

> AccessTokenDetail CreateApiToken(ctx).CreateAPITokenRequest(createAPITokenRequest).Execute()

Create API token



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
    createAPITokenRequest := *openapiclient.NewCreateAPITokenRequest() // CreateAPITokenRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.CreateApiToken(context.Background()).CreateAPITokenRequest(createAPITokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.CreateApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiToken`: AccessTokenDetail
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.CreateApiToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAPITokenRequest** | [**CreateAPITokenRequest**](CreateAPITokenRequest.md) |  | 

### Return type

[**AccessTokenDetail**](AccessTokenDetail.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessUrl

> Object DeleteAccessUrl(ctx, accessUrlId).Execute()

Delete access URL



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
    accessUrlId := int32(56) // int32 | Access URL ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.DeleteAccessUrl(context.Background(), accessUrlId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.DeleteAccessUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccessUrl`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.DeleteAccessUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessUrlId** | **int32** | Access URL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessUrlRequest struct via the builder pattern


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


## DeleteAccessUrlBySearch

> Object DeleteAccessUrlBySearch(ctx).DeleteAccessURLRequest(deleteAccessURLRequest).Execute()

Find and delete access URL



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
    deleteAccessURLRequest := openapiclient.DeleteAccessURLRequest{Interface{}: new(interface{})} // DeleteAccessURLRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.DeleteAccessUrlBySearch(context.Background()).DeleteAccessURLRequest(deleteAccessURLRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.DeleteAccessUrlBySearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccessUrlBySearch`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.DeleteAccessUrlBySearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessUrlBySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteAccessURLRequest** | [**DeleteAccessURLRequest**](DeleteAccessURLRequest.md) |  | 

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


## DeleteApiToken

> SimpleStringResponse DeleteApiToken(ctx, tokenId).Execute()

Delete API token



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
    tokenId := int32(56) // int32 | Token ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.DeleteApiToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.DeleteApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApiToken`: SimpleStringResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.DeleteApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **int32** | Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleStringResponse**](SimpleStringResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessUrl

> Object GetAccessUrl(ctx, accessUrlId).Execute()

Get access URL



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
    accessUrlId := int32(56) // int32 | Access URL ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.GetAccessUrl(context.Background(), accessUrlId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetAccessUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessUrl`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetAccessUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessUrlId** | **int32** | Access URL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessUrlRequest struct via the builder pattern


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


## GetAccessUrls

> AccessUrlListResponse GetAccessUrls(ctx).Execute()

Get access URLs



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
    resp, r, err := apiClient.AccessApi.GetAccessUrls(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetAccessUrls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccessUrls`: AccessUrlListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetAccessUrls`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessUrlsRequest struct via the builder pattern


### Return type

[**AccessUrlListResponse**](AccessUrlListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiToken

> Object GetApiToken(ctx, tokenId).Execute()

Get API access token



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
    tokenId := int32(56) // int32 | Token ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.GetApiToken(context.Background(), tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiToken`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **int32** | Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTokenRequest struct via the builder pattern


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


## GetApiTokens

> AccessTokenListResponse GetApiTokens(ctx).Execute()

Get API access tokens



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
    resp, r, err := apiClient.AccessApi.GetApiTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetApiTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiTokens`: AccessTokenListResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetApiTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiTokensRequest struct via the builder pattern


### Return type

[**AccessTokenListResponse**](AccessTokenListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityControllerSettings

> Object GetIdentityControllerSettings(ctx).Execute()

Get identity Settings for VPN Users



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
    resp, r, err := apiClient.AccessApi.GetIdentityControllerSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetIdentityControllerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityControllerSettings`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetIdentityControllerSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityControllerSettingsRequest struct via the builder pattern


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


## GetIdentityVPNSettings

> Object GetIdentityVPNSettings(ctx).Execute()

Get identity Settings for VPN Users



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
    resp, r, err := apiClient.AccessApi.GetIdentityVPNSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.GetIdentityVPNSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityVPNSettings`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.GetIdentityVPNSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityVPNSettingsRequest struct via the builder pattern


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


## PostTestIdentityControllerSettings

> PostTestIdentityVPNSettings200Response PostTestIdentityControllerSettings(ctx).PostTestIdentityControllerSettingsRequest(postTestIdentityControllerSettingsRequest).Execute()

Test Controller Identity Settings. Currently only LDAP testing is supported.



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
    postTestIdentityControllerSettingsRequest := *openapiclient.NewPostTestIdentityControllerSettingsRequest("Provider_example") // PostTestIdentityControllerSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PostTestIdentityControllerSettings(context.Background()).PostTestIdentityControllerSettingsRequest(postTestIdentityControllerSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PostTestIdentityControllerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTestIdentityControllerSettings`: PostTestIdentityVPNSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PostTestIdentityControllerSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTestIdentityControllerSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTestIdentityControllerSettingsRequest** | [**PostTestIdentityControllerSettingsRequest**](PostTestIdentityControllerSettingsRequest.md) |  | 

### Return type

[**PostTestIdentityVPNSettings200Response**](PostTestIdentityVPNSettings200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTestIdentityVPNSettings

> PostTestIdentityVPNSettings200Response PostTestIdentityVPNSettings(ctx).PostTestLdapSettingsRequest(postTestLdapSettingsRequest).Execute()

Test VPN Identity Settings



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
    postTestLdapSettingsRequest := *openapiclient.NewPostTestLdapSettingsRequest("Provider_example") // PostTestLdapSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PostTestIdentityVPNSettings(context.Background()).PostTestLdapSettingsRequest(postTestLdapSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PostTestIdentityVPNSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTestIdentityVPNSettings`: PostTestIdentityVPNSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PostTestIdentityVPNSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTestIdentityVPNSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTestLdapSettingsRequest** | [**PostTestLdapSettingsRequest**](PostTestLdapSettingsRequest.md) |  | 

### Return type

[**PostTestIdentityVPNSettings200Response**](PostTestIdentityVPNSettings200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutExpireAccessUrl

> Object PutExpireAccessUrl(ctx, accessUrlId).ExpireRequest(expireRequest).Execute()

Expire access URL



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
    accessUrlId := int32(56) // int32 | Access URL ID
    expireRequest := *openapiclient.NewExpireRequest() // ExpireRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutExpireAccessUrl(context.Background(), accessUrlId).ExpireRequest(expireRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutExpireAccessUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutExpireAccessUrl`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutExpireAccessUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessUrlId** | **int32** | Access URL ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutExpireAccessUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expireRequest** | [**ExpireRequest**](ExpireRequest.md) |  | 

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


## PutExpireApiToken

> Object PutExpireApiToken(ctx, tokenId).ExpireRequest(expireRequest).Execute()

Expire API token



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
    tokenId := int32(56) // int32 | Token ID
    expireRequest := *openapiclient.NewExpireRequest() // ExpireRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutExpireApiToken(context.Background(), tokenId).ExpireRequest(expireRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutExpireApiToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutExpireApiToken`: Object
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutExpireApiToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **int32** | Token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutExpireApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expireRequest** | [**ExpireRequest**](ExpireRequest.md) |  | 

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


## PutIdentityControllerSettings

> IdentitySettingsResponse1 PutIdentityControllerSettings(ctx).PutIdentityControllerSettingsRequest(putIdentityControllerSettingsRequest).Execute()

Update Controller Identity settings



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
    putIdentityControllerSettingsRequest := openapiclient.putIdentityControllerSettings_request{ERRORUNKNOWN: new(ERRORUNKNOWN)} // PutIdentityControllerSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutIdentityControllerSettings(context.Background()).PutIdentityControllerSettingsRequest(putIdentityControllerSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutIdentityControllerSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutIdentityControllerSettings`: IdentitySettingsResponse1
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutIdentityControllerSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIdentityControllerSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putIdentityControllerSettingsRequest** | [**PutIdentityControllerSettingsRequest**](PutIdentityControllerSettingsRequest.md) |  | 

### Return type

[**IdentitySettingsResponse1**](IdentitySettingsResponse1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIdentityVPNSettings

> IdentitySettingsResponse PutIdentityVPNSettings(ctx).PutIdentityVPNSettingsRequest(putIdentityVPNSettingsRequest).Execute()

Update VPN Identity settings



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
    putIdentityVPNSettingsRequest := openapiclient.putIdentityVPNSettings_request{IdentityLdapSettingsRequest: openapiclient.NewIdentityLdapSettingsRequest("Provider_example", false)} // PutIdentityVPNSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessApi.PutIdentityVPNSettings(context.Background()).PutIdentityVPNSettingsRequest(putIdentityVPNSettingsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessApi.PutIdentityVPNSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutIdentityVPNSettings`: IdentitySettingsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessApi.PutIdentityVPNSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutIdentityVPNSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putIdentityVPNSettingsRequest** | [**PutIdentityVPNSettingsRequest**](PutIdentityVPNSettingsRequest.md) |  | 

### Return type

[**IdentitySettingsResponse**](IdentitySettingsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

