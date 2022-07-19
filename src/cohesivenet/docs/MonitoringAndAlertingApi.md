# \MonitoringAndAlertingApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePacketMonitor**](MonitoringAndAlertingApi.md#CreatePacketMonitor) | **Post** /packet_monitor | Create packet monitor
[**DeleteAlert**](MonitoringAndAlertingApi.md#DeleteAlert) | **Delete** /alert/{alert_id} | Delete alert
[**DeletePacketMonitor**](MonitoringAndAlertingApi.md#DeletePacketMonitor) | **Delete** /packet_monitor/{name} | Delete packet monitor
[**DeleteWebhook**](MonitoringAndAlertingApi.md#DeleteWebhook) | **Delete** /webhook/{webhook_id} | Delete webhook
[**DownloadPacketMonitorData**](MonitoringAndAlertingApi.md#DownloadPacketMonitorData) | **Get** /packet_monitor/{name}/download | Download packet data
[**GetAlert**](MonitoringAndAlertingApi.md#GetAlert) | **Get** /alert/{alert_id} | Get alert definition details
[**GetAlertEvents**](MonitoringAndAlertingApi.md#GetAlertEvents) | **Get** /alert_events | Get all alert events
[**GetAlerts**](MonitoringAndAlertingApi.md#GetAlerts) | **Get** /alerts | Get all alerts
[**GetPacketMonitor**](MonitoringAndAlertingApi.md#GetPacketMonitor) | **Get** /packet_monitor/{name} | Get packet monitor
[**GetPacketMonitors**](MonitoringAndAlertingApi.md#GetPacketMonitors) | **Get** /packet_monitors | Get all packet monitors
[**GetWebhook**](MonitoringAndAlertingApi.md#GetWebhook) | **Get** /webhook/{webhook_id} | Get webhook details
[**GetWebhooks**](MonitoringAndAlertingApi.md#GetWebhooks) | **Get** /webhooks | Get all webhooks
[**PostCreateAlert**](MonitoringAndAlertingApi.md#PostCreateAlert) | **Post** /alert | Define new alert
[**PostCreateWebhook**](MonitoringAndAlertingApi.md#PostCreateWebhook) | **Post** /webhook | Create new webhook integration
[**PostTestAlert**](MonitoringAndAlertingApi.md#PostTestAlert) | **Post** /alert/{alert_id}/test | Send test alert
[**PostToggleEnableAlert**](MonitoringAndAlertingApi.md#PostToggleEnableAlert) | **Post** /alert/{alert_id}/toggle_enabled | Enable alert
[**PutStartPacketMonitor**](MonitoringAndAlertingApi.md#PutStartPacketMonitor) | **Put** /packet_monitor/{name}/start | Start packet monitor
[**PutStopPacketMonitor**](MonitoringAndAlertingApi.md#PutStopPacketMonitor) | **Put** /packet_monitor/{name}/stop | Stop packet monitor
[**PutUpdateAlert**](MonitoringAndAlertingApi.md#PutUpdateAlert) | **Put** /alert/{alert_id} | Edit alert definition
[**PutUpdateWebhook**](MonitoringAndAlertingApi.md#PutUpdateWebhook) | **Put** /webhook/{webhook_id} | Update webhook configuration



## CreatePacketMonitor

> PacketMonitorDetailResponse CreatePacketMonitor(ctx).CreatePacketMonitorRequest(createPacketMonitorRequest).Execute()

Create packet monitor



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
    createPacketMonitorRequest := *openapiclient.NewCreatePacketMonitorRequest("Name_example", "Type_example", "Interface_example", "Filter_example", "Duration_example", "Destination_example") // CreatePacketMonitorRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.CreatePacketMonitor(context.Background()).CreatePacketMonitorRequest(createPacketMonitorRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.CreatePacketMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePacketMonitor`: PacketMonitorDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.CreatePacketMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePacketMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPacketMonitorRequest** | [**CreatePacketMonitorRequest**](CreatePacketMonitorRequest.md) |  | 

### Return type

[**PacketMonitorDetailResponse**](PacketMonitorDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlert

> Object DeleteAlert(ctx, alertId).Execute()

Delete alert



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
    alertId := int32(56) // int32 | ID for Alert definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.DeleteAlert(context.Background(), alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.DeleteAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAlert`: Object
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.DeleteAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int32** | ID for Alert definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertRequest struct via the builder pattern


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


## DeletePacketMonitor

> SimpleOutputResponse DeletePacketMonitor(ctx, name).Execute()

Delete packet monitor



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
    name := "name_example" // string | Unique name for packet monitor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.DeletePacketMonitor(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.DeletePacketMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePacketMonitor`: SimpleOutputResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.DeletePacketMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Unique name for packet monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePacketMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleOutputResponse**](SimpleOutputResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> Object DeleteWebhook(ctx, webhookId).Execute()

Delete webhook



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
    webhookId := int32(56) // int32 | ID for webhook integration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.DeleteWebhook(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.DeleteWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWebhook`: Object
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32** | ID for webhook integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


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


## DownloadPacketMonitorData

> *os.File DownloadPacketMonitorData(ctx, name).Execute()

Download packet data



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
    name := "name_example" // string | Unique name for packet monitor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.DownloadPacketMonitorData(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.DownloadPacketMonitorData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadPacketMonitorData`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.DownloadPacketMonitorData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Unique name for packet monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadPacketMonitorDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlert

> Object GetAlert(ctx, alertId).Execute()

Get alert definition details



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
    alertId := int32(56) // int32 | ID for Alert definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.GetAlert(context.Background(), alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.GetAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlert`: Object
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.GetAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int32** | ID for Alert definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertRequest struct via the builder pattern


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


## GetAlertEvents

> AlertEventTypesListResponse GetAlertEvents(ctx).Execute()

Get all alert events



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
    resp, r, err := apiClient.MonitoringAndAlertingApi.GetAlertEvents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.GetAlertEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertEvents`: AlertEventTypesListResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.GetAlertEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertEventsRequest struct via the builder pattern


### Return type

[**AlertEventTypesListResponse**](AlertEventTypesListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlerts

> AlertsListResponse GetAlerts(ctx).Execute()

Get all alerts



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
    resp, r, err := apiClient.MonitoringAndAlertingApi.GetAlerts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.GetAlerts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlerts`: AlertsListResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.GetAlerts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsRequest struct via the builder pattern


### Return type

[**AlertsListResponse**](AlertsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPacketMonitor

> Object GetPacketMonitor(ctx, name).Execute()

Get packet monitor



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
    name := "name_example" // string | Unique name for packet monitor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.GetPacketMonitor(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.GetPacketMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPacketMonitor`: Object
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.GetPacketMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Unique name for packet monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPacketMonitorRequest struct via the builder pattern


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


## GetPacketMonitors

> PacketMonitorsListResponse GetPacketMonitors(ctx).Execute()

Get all packet monitors



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
    resp, r, err := apiClient.MonitoringAndAlertingApi.GetPacketMonitors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.GetPacketMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPacketMonitors`: PacketMonitorsListResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.GetPacketMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPacketMonitorsRequest struct via the builder pattern


### Return type

[**PacketMonitorsListResponse**](PacketMonitorsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> Object GetWebhook(ctx, webhookId).Execute()

Get webhook details



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
    webhookId := int32(56) // int32 | ID for webhook integration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.GetWebhook(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.GetWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhook`: Object
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32** | ID for webhook integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


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


## GetWebhooks

> WebhooksListResponse GetWebhooks(ctx).Execute()

Get all webhooks



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
    resp, r, err := apiClient.MonitoringAndAlertingApi.GetWebhooks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.GetWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooks`: WebhooksListResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.GetWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


### Return type

[**WebhooksListResponse**](WebhooksListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateAlert

> Object PostCreateAlert(ctx).CreateAlertRequest(createAlertRequest).Execute()

Define new alert



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
    createAlertRequest := *openapiclient.NewCreateAlertRequest("Name_example", "Url_example", int32(123)) // CreateAlertRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.PostCreateAlert(context.Background()).CreateAlertRequest(createAlertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.PostCreateAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateAlert`: Object
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.PostCreateAlert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAlertRequest** | [**CreateAlertRequest**](CreateAlertRequest.md) |  | 

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


## PostCreateWebhook

> WebhookDetailResponse PostCreateWebhook(ctx).CreateWebookRequest(createWebookRequest).Execute()

Create new webhook integration



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
    createWebookRequest := *openapiclient.NewCreateWebookRequest("Name_example") // CreateWebookRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.PostCreateWebhook(context.Background()).CreateWebookRequest(createWebookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.PostCreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateWebhook`: WebhookDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.PostCreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWebookRequest** | [**CreateWebookRequest**](CreateWebookRequest.md) |  | 

### Return type

[**WebhookDetailResponse**](WebhookDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTestAlert

> PostTestAlert200Response PostTestAlert(ctx, alertId).Execute()

Send test alert



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
    alertId := int32(56) // int32 | ID for Alert definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.PostTestAlert(context.Background(), alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.PostTestAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostTestAlert`: PostTestAlert200Response
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.PostTestAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int32** | ID for Alert definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTestAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostTestAlert200Response**](PostTestAlert200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostToggleEnableAlert

> Object PostToggleEnableAlert(ctx, alertId).Execute()

Enable alert



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
    alertId := int32(56) // int32 | ID for Alert definition

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.PostToggleEnableAlert(context.Background(), alertId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.PostToggleEnableAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostToggleEnableAlert`: Object
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.PostToggleEnableAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int32** | ID for Alert definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostToggleEnableAlertRequest struct via the builder pattern


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


## PutStartPacketMonitor

> Object PutStartPacketMonitor(ctx, name).Execute()

Start packet monitor



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
    name := "name_example" // string | Unique name for packet monitor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.PutStartPacketMonitor(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.PutStartPacketMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutStartPacketMonitor`: Object
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.PutStartPacketMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Unique name for packet monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutStartPacketMonitorRequest struct via the builder pattern


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


## PutStopPacketMonitor

> Object PutStopPacketMonitor(ctx, name).Execute()

Stop packet monitor



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
    name := "name_example" // string | Unique name for packet monitor

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.PutStopPacketMonitor(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.PutStopPacketMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutStopPacketMonitor`: Object
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.PutStopPacketMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Unique name for packet monitor | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutStopPacketMonitorRequest struct via the builder pattern


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


## PutUpdateAlert

> Object PutUpdateAlert(ctx, alertId).UpdateAlertRequest(updateAlertRequest).Execute()

Edit alert definition



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
    alertId := int32(56) // int32 | ID for Alert definition
    updateAlertRequest := *openapiclient.NewUpdateAlertRequest() // UpdateAlertRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.PutUpdateAlert(context.Background(), alertId).UpdateAlertRequest(updateAlertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.PutUpdateAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateAlert`: Object
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.PutUpdateAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**alertId** | **int32** | ID for Alert definition | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAlertRequest** | [**UpdateAlertRequest**](UpdateAlertRequest.md) |  | 

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


## PutUpdateWebhook

> Object PutUpdateWebhook(ctx, webhookId).UpdateWebhookRequest(updateWebhookRequest).Execute()

Update webhook configuration



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
    webhookId := int32(56) // int32 | ID for webhook integration
    updateWebhookRequest := *openapiclient.NewUpdateWebhookRequest() // UpdateWebhookRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoringAndAlertingApi.PutUpdateWebhook(context.Background(), webhookId).UpdateWebhookRequest(updateWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoringAndAlertingApi.PutUpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateWebhook`: Object
    fmt.Fprintf(os.Stdout, "Response from `MonitoringAndAlertingApi.PutUpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int32** | ID for webhook integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWebhookRequest** | [**UpdateWebhookRequest**](UpdateWebhookRequest.md) |  | 

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

