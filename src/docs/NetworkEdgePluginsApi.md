# \NetworkEdgePluginsApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearPluginInstanceConfFileHistory**](NetworkEdgePluginsApi.md#ClearPluginInstanceConfFileHistory) | **Post** /plugin-instances/{id}/configurations/{slug}/clear | Clear plugin instance configuration file versions except for the current
[**CreatePluginExport**](NetworkEdgePluginsApi.md#CreatePluginExport) | **Post** /plugins/{id}/exports | Create new plugin export file
[**CreatePluginInstanceExport**](NetworkEdgePluginsApi.md#CreatePluginInstanceExport) | **Post** /plugin-instances/{id}/exports | Create new plugin instance export file
[**DeleteContainer**](NetworkEdgePluginsApi.md#DeleteContainer) | **Delete** /container_system/containers/{uuid} | Delete container
[**DeleteContainerImage**](NetworkEdgePluginsApi.md#DeleteContainerImage) | **Delete** /container_system/images/{uuid} | Delete container image
[**DeletePlugin**](NetworkEdgePluginsApi.md#DeletePlugin) | **Delete** /plugins/{id} | Delete plugin
[**DeletePluginExport**](NetworkEdgePluginsApi.md#DeletePluginExport) | **Delete** /plugins/{id}/exports/{export_name} | Delete plugin export file
[**DeletePluginInstance**](NetworkEdgePluginsApi.md#DeletePluginInstance) | **Delete** /plugin-instances/{id} | Delete plugin instance
[**DeletePluginInstanceConfigVersion**](NetworkEdgePluginsApi.md#DeletePluginInstanceConfigVersion) | **Delete** /plugin-instances/{id}/configurations/{slug}/versions/{version} | Delete plugin instance config file version
[**DeletePluginInstanceExport**](NetworkEdgePluginsApi.md#DeletePluginInstanceExport) | **Delete** /plugin-instances/{id}/exports/{export_name} | Delete plugin instance export file
[**DeletePluginInstanceUser**](NetworkEdgePluginsApi.md#DeletePluginInstanceUser) | **Delete** /plugin-instances/{id}/users/{username} | Delete plugin instance SSH user
[**GetContainerImages**](NetworkEdgePluginsApi.md#GetContainerImages) | **Get** /container_system/images | Get container images
[**GetContainerLogs**](NetworkEdgePluginsApi.md#GetContainerLogs) | **Get** /container_system/containers/{uuid}/logs | Get container logs
[**GetContainerSystemIps**](NetworkEdgePluginsApi.md#GetContainerSystemIps) | **Get** /container_system/ip_addresses | Get container system IP addresses
[**GetContainerSystemStatus**](NetworkEdgePluginsApi.md#GetContainerSystemStatus) | **Get** /container_system | Get container system status
[**GetDownloadPluginExport**](NetworkEdgePluginsApi.md#GetDownloadPluginExport) | **Get** /plugins/{id}/exports/{export_name} | Download plugin export file
[**GetDownloadPluginInstanceExport**](NetworkEdgePluginsApi.md#GetDownloadPluginInstanceExport) | **Get** /plugin-instances/{id}/exports/{export_name} | Download plugin instance export file
[**GetPlugin**](NetworkEdgePluginsApi.md#GetPlugin) | **Get** /plugins/{id} | Get plugin
[**GetPluginCatalog**](NetworkEdgePluginsApi.md#GetPluginCatalog) | **Get** /plugin-catalog | Get the list of plugins available in the catalog
[**GetPluginExports**](NetworkEdgePluginsApi.md#GetPluginExports) | **Get** /plugins/{id}/exports | Get available plugin export files
[**GetPluginInstance**](NetworkEdgePluginsApi.md#GetPluginInstance) | **Get** /plugin-instances/{id} | Get plugin instance
[**GetPluginInstanceConfigContent**](NetworkEdgePluginsApi.md#GetPluginInstanceConfigContent) | **Get** /plugin-instances/{id}/configurations/{slug}/content | Read plugin instance config file
[**GetPluginInstanceConfigFiles**](NetworkEdgePluginsApi.md#GetPluginInstanceConfigFiles) | **Get** /plugin-instances/{id}/configurations | Get plugin instance configs
[**GetPluginInstanceExports**](NetworkEdgePluginsApi.md#GetPluginInstanceExports) | **Get** /plugin-instances/{id}/exports | Get available plugin instance export files
[**GetPluginInstanceFirewall**](NetworkEdgePluginsApi.md#GetPluginInstanceFirewall) | **Get** /plugin-instances/{id}/firewall | Get plugin instance firewall
[**GetPluginInstanceLogContent**](NetworkEdgePluginsApi.md#GetPluginInstanceLogContent) | **Get** /plugin-instances/{id}/logs/{slug}/content | Read plugin instance log
[**GetPluginInstanceLogFiles**](NetworkEdgePluginsApi.md#GetPluginInstanceLogFiles) | **Get** /plugin-instances/{id}/logs | Get plugin instance logs
[**GetPluginInstanceProcesses**](NetworkEdgePluginsApi.md#GetPluginInstanceProcesses) | **Get** /plugin-instances/{id}/processes | Get plugin instance processes
[**GetPluginInstances**](NetworkEdgePluginsApi.md#GetPluginInstances) | **Get** /plugin-instances | Get plugin instances
[**GetPluginInstanceusers**](NetworkEdgePluginsApi.md#GetPluginInstanceusers) | **Get** /plugin-instances/{id}/users | Get list of plugin instance users
[**GetPluginManagerConfig**](NetworkEdgePluginsApi.md#GetPluginManagerConfig) | **Get** /plugins/{id}/manager | Get plugin manager config json
[**GetPluginSystemExports**](NetworkEdgePluginsApi.md#GetPluginSystemExports) | **Get** /plugin-exports | Get plugin exports
[**GetPlugins**](NetworkEdgePluginsApi.md#GetPlugins) | **Get** /plugins | Get plugins
[**GetRunningContainers**](NetworkEdgePluginsApi.md#GetRunningContainers) | **Get** /container_system/containers | Get running containers
[**InstallPlugin**](NetworkEdgePluginsApi.md#InstallPlugin) | **Post** /plugins | Install new plugin from the catalog
[**PostActionContainerSystem**](NetworkEdgePluginsApi.md#PostActionContainerSystem) | **Post** /container_system | Take action on container system
[**PostCommitContainer**](NetworkEdgePluginsApi.md#PostCommitContainer) | **Post** /container_system/containers/{uuid}/commit | Commit container
[**PostCommitPluginInstance**](NetworkEdgePluginsApi.md#PostCommitPluginInstance) | **Post** /plugin-instances/{id}/commit | Commit plugin instance
[**PostCreateContainerImage**](NetworkEdgePluginsApi.md#PostCreateContainerImage) | **Post** /container_system/images | Create container image
[**PostCreateManagerConfig**](NetworkEdgePluginsApi.md#PostCreateManagerConfig) | **Post** /plugin-instances/{id}/manager | Create Manager Config
[**PostExportImage**](NetworkEdgePluginsApi.md#PostExportImage) | **Post** /container_system/images/{uuid}/exports | Create container image export
[**PostStartContainer**](NetworkEdgePluginsApi.md#PostStartContainer) | **Post** /container_system/containers | Start container
[**PutConfigureContainerSystem**](NetworkEdgePluginsApi.md#PutConfigureContainerSystem) | **Put** /container_system | Update container system
[**PutPluginInstanceFirewallRule**](NetworkEdgePluginsApi.md#PutPluginInstanceFirewallRule) | **Put** /plugin-instances/{id}/firewall | Create plugin instance firewall rule
[**PutPluginInstanceUser**](NetworkEdgePluginsApi.md#PutPluginInstanceUser) | **Put** /plugin-instances/{id}/users | Update or create a plugin instance user
[**PutStopContainer**](NetworkEdgePluginsApi.md#PutStopContainer) | **Put** /container_system/containers/{uuid} | Stop container
[**PutUpdateContainerImage**](NetworkEdgePluginsApi.md#PutUpdateContainerImage) | **Put** /container_system/images/{uuid} | Update container image
[**PutUpdateManagerConfig**](NetworkEdgePluginsApi.md#PutUpdateManagerConfig) | **Put** /plugin-instances/{id}/manager | Update Manager Config
[**PutUpdatePlugin**](NetworkEdgePluginsApi.md#PutUpdatePlugin) | **Put** /plugins/{id} | Update plugin data
[**PutUpdatePluginInstanceConfigContent**](NetworkEdgePluginsApi.md#PutUpdatePluginInstanceConfigContent) | **Put** /plugin-instances/{id}/configurations/{slug}/content | Update plugin instance config file
[**RevertPluginInstanceConfigFile**](NetworkEdgePluginsApi.md#RevertPluginInstanceConfigFile) | **Post** /plugin-instances/{id}/configurations/{slug}/revert | Revert plugin instance config file
[**RunPluginInstanceExecutableCommand**](NetworkEdgePluginsApi.md#RunPluginInstanceExecutableCommand) | **Post** /plugin-instances/{id}/commands/execute | Run plugin instance executable command
[**RunPluginInstanceProcessAction**](NetworkEdgePluginsApi.md#RunPluginInstanceProcessAction) | **Post** /plugin-instances/{id}/processes/action | Run plugin instance process action
[**StartPluginInstance**](NetworkEdgePluginsApi.md#StartPluginInstance) | **Post** /plugin-instances | Start new plugin instance from a plugin



## ClearPluginInstanceConfFileHistory

> Object ClearPluginInstanceConfFileHistory(ctx, id, slug).Execute()

Clear plugin instance configuration file versions except for the current



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
    id := int32(56) // int32 | ID for plugin instance (container)
    slug := openapiclient.FirewallRuleTuple_inner{Int32: new(int32)} // FirewallRuleTupleInner | Either the index of the config file in the manager configuration  or the name of the confiig file. e.g. 0 or \"name\" 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.ClearPluginInstanceConfFileHistory(context.Background(), id, slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.ClearPluginInstanceConfFileHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClearPluginInstanceConfFileHistory`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.ClearPluginInstanceConfFileHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 
**slug** | [**FirewallRuleTupleInner**](.md) | Either the index of the config file in the manager configuration  or the name of the confiig file. e.g. 0 or \&quot;name\&quot;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearPluginInstanceConfFileHistoryRequest struct via the builder pattern


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


## CreatePluginExport

> PluginExportDetailResponse CreatePluginExport(ctx, id).CreatePluginExportRequest(createPluginExportRequest).Execute()

Create new plugin export file



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
    id := int32(56) // int32 | ID for plugin (container image)
    createPluginExportRequest := *openapiclient.NewCreatePluginExportRequest("Name_example") // CreatePluginExportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.CreatePluginExport(context.Background(), id).CreatePluginExportRequest(createPluginExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.CreatePluginExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePluginExport`: PluginExportDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.CreatePluginExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin (container image) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePluginExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPluginExportRequest** | [**CreatePluginExportRequest**](CreatePluginExportRequest.md) |  | 

### Return type

[**PluginExportDetailResponse**](PluginExportDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePluginInstanceExport

> Object CreatePluginInstanceExport(ctx, id).CreatePluginInstanceExportRequest(createPluginInstanceExportRequest).Execute()

Create new plugin instance export file



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
    id := int32(56) // int32 | ID for plugin instance
    createPluginInstanceExportRequest := *openapiclient.NewCreatePluginInstanceExportRequest("Name_example") // CreatePluginInstanceExportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.CreatePluginInstanceExport(context.Background(), id).CreatePluginInstanceExportRequest(createPluginInstanceExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.CreatePluginInstanceExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePluginInstanceExport`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.CreatePluginInstanceExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePluginInstanceExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPluginInstanceExportRequest** | [**CreatePluginInstanceExportRequest**](CreatePluginInstanceExportRequest.md) |  | 

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


## DeleteContainer

> DeleteContainerDetailResponse DeleteContainer(ctx, uuid).Execute()

Delete container



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
    resp, r, err := apiClient.NetworkEdgePluginsApi.DeleteContainer(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.DeleteContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContainer`: DeleteContainerDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.DeleteContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteContainerDetailResponse**](DeleteContainerDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerImage

> DeleteContainerImageDetailResponse DeleteContainerImage(ctx, uuid).Force(force).Execute()

Delete container image



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
    force := true // bool | force operation with cleanup (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.DeleteContainerImage(context.Background(), uuid).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.DeleteContainerImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContainerImage`: DeleteContainerImageDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.DeleteContainerImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force operation with cleanup | [default to false]

### Return type

[**DeleteContainerImageDetailResponse**](DeleteContainerImageDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlugin

> Object DeletePlugin(ctx, id).Execute()

Delete plugin



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
    id := int32(56) // int32 | ID for plugin (container image)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.DeletePlugin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.DeletePlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePlugin`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.DeletePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin (container image) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePluginRequest struct via the builder pattern


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


## DeletePluginExport

> Object DeletePluginExport(ctx, id, exportName).Execute()

Delete plugin export file



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
    id := int32(56) // int32 | ID for plugin (container image)
    exportName := "exportName_example" // string | Name of export file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.DeletePluginExport(context.Background(), id, exportName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.DeletePluginExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePluginExport`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.DeletePluginExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin (container image) | 
**exportName** | **string** | Name of export file | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePluginExportRequest struct via the builder pattern


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


## DeletePluginInstance

> Object DeletePluginInstance(ctx, id).DeletePluginInstanceRequest(deletePluginInstanceRequest).Execute()

Delete plugin instance



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
    id := int32(56) // int32 | ID for plugin instance
    deletePluginInstanceRequest := *openapiclient.NewDeletePluginInstanceRequest() // DeletePluginInstanceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.DeletePluginInstance(context.Background(), id).DeletePluginInstanceRequest(deletePluginInstanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.DeletePluginInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePluginInstance`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.DeletePluginInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePluginInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deletePluginInstanceRequest** | [**DeletePluginInstanceRequest**](DeletePluginInstanceRequest.md) |  | 

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


## DeletePluginInstanceConfigVersion

> Object DeletePluginInstanceConfigVersion(ctx, id, slug, version).Execute()

Delete plugin instance config file version



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
    id := int32(56) // int32 | ID for plugin instance (container)
    slug := openapiclient.FirewallRuleTuple_inner{Int32: new(int32)} // FirewallRuleTupleInner | Either the index of the config file in the configuration  or the name of the config file. e.g. 0 or \"name\" 
    version := int32(56) // int32 | Version to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.DeletePluginInstanceConfigVersion(context.Background(), id, slug, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.DeletePluginInstanceConfigVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePluginInstanceConfigVersion`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.DeletePluginInstanceConfigVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 
**slug** | [**FirewallRuleTupleInner**](.md) | Either the index of the config file in the configuration  or the name of the config file. e.g. 0 or \&quot;name\&quot;  | 
**version** | **int32** | Version to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePluginInstanceConfigVersionRequest struct via the builder pattern


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


## DeletePluginInstanceExport

> Object DeletePluginInstanceExport(ctx, id, exportName).Execute()

Delete plugin instance export file



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
    id := int32(56) // int32 | ID for plugin instance
    exportName := "exportName_example" // string | Name of export file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.DeletePluginInstanceExport(context.Background(), id, exportName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.DeletePluginInstanceExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePluginInstanceExport`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.DeletePluginInstanceExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance | 
**exportName** | **string** | Name of export file | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePluginInstanceExportRequest struct via the builder pattern


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


## DeletePluginInstanceUser

> Object DeletePluginInstanceUser(ctx, id, username).Execute()

Delete plugin instance SSH user



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
    id := int32(56) // int32 | ID for plugin instance
    username := "username_example" // string | Name of user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.DeletePluginInstanceUser(context.Background(), id, username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.DeletePluginInstanceUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePluginInstanceUser`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.DeletePluginInstanceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance | 
**username** | **string** | Name of user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePluginInstanceUserRequest struct via the builder pattern


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


## GetContainerImages

> ContainerImageListResponse GetContainerImages(ctx).Uuid(uuid).Execute()

Get container images



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
    uuid := "uuid_example" // string | UUID for Container (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetContainerImages(context.Background()).Uuid(uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetContainerImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerImages`: ContainerImageListResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetContainerImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uuid** | **string** | UUID for Container | 

### Return type

[**ContainerImageListResponse**](ContainerImageListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerLogs

> ContainerLogsResponse GetContainerLogs(ctx, uuid).Lines(lines).Execute()

Get container logs



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
    lines := int32(56) // int32 | Number of log lines to fetch

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetContainerLogs(context.Background(), uuid).Lines(lines).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetContainerLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerLogs`: ContainerLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetContainerLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lines** | **int32** | Number of log lines to fetch | 

### Return type

[**ContainerLogsResponse**](ContainerLogsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerSystemIps

> ContainerSystemIPListResponse GetContainerSystemIps(ctx).Execute()

Get container system IP addresses



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
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetContainerSystemIps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetContainerSystemIps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerSystemIps`: ContainerSystemIPListResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetContainerSystemIps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerSystemIpsRequest struct via the builder pattern


### Return type

[**ContainerSystemIPListResponse**](ContainerSystemIPListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerSystemStatus

> ContainerSystemStatusDetailResponse GetContainerSystemStatus(ctx).Execute()

Get container system status



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
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetContainerSystemStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetContainerSystemStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerSystemStatus`: ContainerSystemStatusDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetContainerSystemStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerSystemStatusRequest struct via the builder pattern


### Return type

[**ContainerSystemStatusDetailResponse**](ContainerSystemStatusDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDownloadPluginExport

> *os.File GetDownloadPluginExport(ctx, id, exportName).Execute()

Download plugin export file



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
    id := int32(56) // int32 | ID for plugin (container image)
    exportName := "exportName_example" // string | Name of export file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetDownloadPluginExport(context.Background(), id, exportName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetDownloadPluginExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadPluginExport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetDownloadPluginExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin (container image) | 
**exportName** | **string** | Name of export file | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadPluginExportRequest struct via the builder pattern


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


## GetDownloadPluginInstanceExport

> *os.File GetDownloadPluginInstanceExport(ctx, id, exportName).Execute()

Download plugin instance export file



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
    id := int32(56) // int32 | ID for plugin instance
    exportName := "exportName_example" // string | Name of export file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetDownloadPluginInstanceExport(context.Background(), id, exportName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetDownloadPluginInstanceExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDownloadPluginInstanceExport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetDownloadPluginInstanceExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance | 
**exportName** | **string** | Name of export file | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadPluginInstanceExportRequest struct via the builder pattern


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


## GetPlugin

> Object GetPlugin(ctx, id).Execute()

Get plugin



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
    id := int32(56) // int32 | ID for plugin (container image)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPlugin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlugin`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin (container image) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginRequest struct via the builder pattern


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


## GetPluginCatalog

> PluginCatalogResponse GetPluginCatalog(ctx).Execute()

Get the list of plugins available in the catalog



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
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginCatalog(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginCatalog`: PluginCatalogResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginCatalogRequest struct via the builder pattern


### Return type

[**PluginCatalogResponse**](PluginCatalogResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginExports

> Object GetPluginExports(ctx, id).Execute()

Get available plugin export files



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
    id := int32(56) // int32 | ID for plugin (container image)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginExports(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginExports`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginExports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin (container image) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginExportsRequest struct via the builder pattern


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


## GetPluginInstance

> Object GetPluginInstance(ctx, id).SystemData(systemData).Execute()

Get plugin instance



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
    id := int32(56) // int32 | ID for plugin instance (container)
    systemData := true // bool | Fetch instance system data (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginInstance(context.Background(), id).SystemData(systemData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginInstance`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **systemData** | **bool** | Fetch instance system data | [default to true]

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


## GetPluginInstanceConfigContent

> Object GetPluginInstanceConfigContent(ctx, id, slug).Version(version).Execute()

Read plugin instance config file



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
    id := int32(56) // int32 | ID for plugin instance (container)
    slug := openapiclient.FirewallRuleTuple_inner{Int32: new(int32)} // FirewallRuleTupleInner | Either the index of the config file in the configuration  or the name of the config file. e.g. 0 or \"name\" 
    version := int32(56) // int32 | Optional version to read content from (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginInstanceConfigContent(context.Background(), id, slug).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginInstanceConfigContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginInstanceConfigContent`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginInstanceConfigContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 
**slug** | [**FirewallRuleTupleInner**](.md) | Either the index of the config file in the configuration  or the name of the config file. e.g. 0 or \&quot;name\&quot;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginInstanceConfigContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **int32** | Optional version to read content from | 

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


## GetPluginInstanceConfigFiles

> PluginManagerConfigConfigFilesResponse GetPluginInstanceConfigFiles(ctx, id).Execute()

Get plugin instance configs



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
    id := int32(56) // int32 | ID for plugin instance (container)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginInstanceConfigFiles(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginInstanceConfigFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginInstanceConfigFiles`: PluginManagerConfigConfigFilesResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginInstanceConfigFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginInstanceConfigFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PluginManagerConfigConfigFilesResponse**](PluginManagerConfigConfigFilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginInstanceExports

> Object GetPluginInstanceExports(ctx, id).Execute()

Get available plugin instance export files



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
    id := int32(56) // int32 | ID for plugin instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginInstanceExports(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginInstanceExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginInstanceExports`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginInstanceExports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginInstanceExportsRequest struct via the builder pattern


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


## GetPluginInstanceFirewall

> PluginInstanceFirewallResponse GetPluginInstanceFirewall(ctx, id).Execute()

Get plugin instance firewall



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
    id := int32(56) // int32 | ID for plugin instance (container)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginInstanceFirewall(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginInstanceFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginInstanceFirewall`: PluginInstanceFirewallResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginInstanceFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginInstanceFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PluginInstanceFirewallResponse**](PluginInstanceFirewallResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginInstanceLogContent

> SimpleStringListResponse GetPluginInstanceLogContent(ctx, id, slug).Lines(lines).Execute()

Read plugin instance log



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
    id := int32(56) // int32 | ID for plugin instance (container)
    slug := openapiclient.FirewallRuleTuple_inner{Int32: new(int32)} // FirewallRuleTupleInner | Either the index of the log file in the configuration  or the name of the log file. e.g. 0 or \"name\" 
    lines := int32(56) // int32 | Number of log lines to return (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginInstanceLogContent(context.Background(), id, slug).Lines(lines).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginInstanceLogContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginInstanceLogContent`: SimpleStringListResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginInstanceLogContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 
**slug** | [**FirewallRuleTupleInner**](.md) | Either the index of the log file in the configuration  or the name of the log file. e.g. 0 or \&quot;name\&quot;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginInstanceLogContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lines** | **int32** | Number of log lines to return | [default to 25]

### Return type

[**SimpleStringListResponse**](SimpleStringListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginInstanceLogFiles

> PluginManagerConfigListLogFilesResponse GetPluginInstanceLogFiles(ctx, id).Execute()

Get plugin instance logs



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
    id := int32(56) // int32 | ID for plugin instance (container)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginInstanceLogFiles(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginInstanceLogFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginInstanceLogFiles`: PluginManagerConfigListLogFilesResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginInstanceLogFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginInstanceLogFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PluginManagerConfigListLogFilesResponse**](PluginManagerConfigListLogFilesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginInstanceProcesses

> PluginManagerConfigProcessManagerResponse GetPluginInstanceProcesses(ctx, id).Execute()

Get plugin instance processes



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
    id := int32(56) // int32 | ID for plugin instance (container)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginInstanceProcesses(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginInstanceProcesses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginInstanceProcesses`: PluginManagerConfigProcessManagerResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginInstanceProcesses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginInstanceProcessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PluginManagerConfigProcessManagerResponse**](PluginManagerConfigProcessManagerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginInstances

> PluginInstanceListResponse GetPluginInstances(ctx).Execute()

Get plugin instances



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
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginInstances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginInstances`: PluginInstanceListResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginInstances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginInstancesRequest struct via the builder pattern


### Return type

[**PluginInstanceListResponse**](PluginInstanceListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginInstanceusers

> PluginInstanceUsersResponse GetPluginInstanceusers(ctx, id).Execute()

Get list of plugin instance users



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
    id := int32(56) // int32 | ID for plugin instance

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginInstanceusers(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginInstanceusers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginInstanceusers`: PluginInstanceUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginInstanceusers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginInstanceusersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PluginInstanceUsersResponse**](PluginInstanceUsersResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginManagerConfig

> PluginManagerConfigResponse GetPluginManagerConfig(ctx, id).Execute()

Get plugin manager config json



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
    id := int32(56) // int32 | ID for plugin (container image)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginManagerConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginManagerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginManagerConfig`: PluginManagerConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginManagerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin (container image) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginManagerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PluginManagerConfigResponse**](PluginManagerConfigResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginSystemExports

> PluginExportsResponse GetPluginSystemExports(ctx).Execute()

Get plugin exports



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
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPluginSystemExports(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPluginSystemExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPluginSystemExports`: PluginExportsResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPluginSystemExports`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginSystemExportsRequest struct via the builder pattern


### Return type

[**PluginExportsResponse**](PluginExportsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlugins

> PluginListResponse GetPlugins(ctx).Execute()

Get plugins



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
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetPlugins(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetPlugins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlugins`: PluginListResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetPlugins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginsRequest struct via the builder pattern


### Return type

[**PluginListResponse**](PluginListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunningContainers

> RunningContainersDetailResponse GetRunningContainers(ctx).ShowAll(showAll).Uuid(uuid).Execute()

Get running containers



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
    showAll := true // bool | Boolean for full list output of containers (optional) (default to true)
    uuid := "uuid_example" // string | UUID for resource (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.GetRunningContainers(context.Background()).ShowAll(showAll).Uuid(uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.GetRunningContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRunningContainers`: RunningContainersDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.GetRunningContainers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRunningContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showAll** | **bool** | Boolean for full list output of containers | [default to true]
 **uuid** | **string** | UUID for resource | 

### Return type

[**RunningContainersDetailResponse**](RunningContainersDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallPlugin

> PluginDetailResponse InstallPlugin(ctx).InstallPluginRequest(installPluginRequest).Execute()

Install new plugin from the catalog



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
    installPluginRequest := *openapiclient.NewInstallPluginRequest("Name_example", "ImageUrl_example") // InstallPluginRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.InstallPlugin(context.Background()).InstallPluginRequest(installPluginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.InstallPlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstallPlugin`: PluginDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.InstallPlugin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstallPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **installPluginRequest** | [**InstallPluginRequest**](InstallPluginRequest.md) |  | 

### Return type

[**PluginDetailResponse**](PluginDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostActionContainerSystem

> Object PostActionContainerSystem(ctx).ContainerSystemActionRequest(containerSystemActionRequest).Execute()

Take action on container system



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
    containerSystemActionRequest := *openapiclient.NewContainerSystemActionRequest("Action_example") // ContainerSystemActionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PostActionContainerSystem(context.Background()).ContainerSystemActionRequest(containerSystemActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PostActionContainerSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostActionContainerSystem`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PostActionContainerSystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostActionContainerSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerSystemActionRequest** | [**ContainerSystemActionRequest**](ContainerSystemActionRequest.md) |  | 

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


## PostCommitContainer

> CreateContainerImageResponse PostCommitContainer(ctx, uuid).CommitContainerRequest(commitContainerRequest).Execute()

Commit container



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
    commitContainerRequest := *openapiclient.NewCommitContainerRequest(false) // CommitContainerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PostCommitContainer(context.Background(), uuid).CommitContainerRequest(commitContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PostCommitContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCommitContainer`: CreateContainerImageResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PostCommitContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommitContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commitContainerRequest** | [**CommitContainerRequest**](CommitContainerRequest.md) |  | 

### Return type

[**CreateContainerImageResponse**](CreateContainerImageResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommitPluginInstance

> Object PostCommitPluginInstance(ctx, id).CommitPluginInstanceImageRequest(commitPluginInstanceImageRequest).Execute()

Commit plugin instance



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
    id := int32(56) // int32 | ID for plugin instance (container)
    commitPluginInstanceImageRequest := *openapiclient.NewCommitPluginInstanceImageRequest("Name_example") // CommitPluginInstanceImageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PostCommitPluginInstance(context.Background(), id).CommitPluginInstanceImageRequest(commitPluginInstanceImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PostCommitPluginInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCommitPluginInstance`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PostCommitPluginInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommitPluginInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commitPluginInstanceImageRequest** | [**CommitPluginInstanceImageRequest**](CommitPluginInstanceImageRequest.md) |  | 

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


## PostCreateContainerImage

> CreateImageDetailResponse PostCreateContainerImage(ctx).CreateContainerImageRequest(createContainerImageRequest).Execute()

Create container image



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
    createContainerImageRequest := openapiclient.CreateContainerImageRequest{Interface{}: new(interface{})} // CreateContainerImageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PostCreateContainerImage(context.Background()).CreateContainerImageRequest(createContainerImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PostCreateContainerImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateContainerImage`: CreateImageDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PostCreateContainerImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateContainerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContainerImageRequest** | [**CreateContainerImageRequest**](CreateContainerImageRequest.md) |  | 

### Return type

[**CreateImageDetailResponse**](CreateImageDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateManagerConfig

> Object PostCreateManagerConfig(ctx, id).Body(body).Execute()

Create Manager Config



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
    id := int32(56) // int32 | ID for plugin instance (container)
    body := Object(987) // Object |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PostCreateManagerConfig(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PostCreateManagerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateManagerConfig`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PostCreateManagerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateManagerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **Object** |  | 

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


## PostExportImage

> Object PostExportImage(ctx, uuid).CreateContainerImageExportRequest(createContainerImageExportRequest).Execute()

Create container image export



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
    createContainerImageExportRequest := *openapiclient.NewCreateContainerImageExportRequest("Name_example") // CreateContainerImageExportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PostExportImage(context.Background(), uuid).CreateContainerImageExportRequest(createContainerImageExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PostExportImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostExportImage`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PostExportImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostExportImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createContainerImageExportRequest** | [**CreateContainerImageExportRequest**](CreateContainerImageExportRequest.md) |  | 

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


## PostStartContainer

> RunContainerDetailResponse PostStartContainer(ctx).AllocateContainerRequest(allocateContainerRequest).Execute()

Start container



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
    allocateContainerRequest := openapiclient.AllocateContainerRequest{Interface{}: new(interface{})} // AllocateContainerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PostStartContainer(context.Background()).AllocateContainerRequest(allocateContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PostStartContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostStartContainer`: RunContainerDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PostStartContainer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStartContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allocateContainerRequest** | [**AllocateContainerRequest**](AllocateContainerRequest.md) |  | 

### Return type

[**RunContainerDetailResponse**](RunContainerDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutConfigureContainerSystem

> Object PutConfigureContainerSystem(ctx).UpdateConfigureContainerSystemRequest(updateConfigureContainerSystemRequest).Execute()

Update container system



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
    updateConfigureContainerSystemRequest := *openapiclient.NewUpdateConfigureContainerSystemRequest("Network_example") // UpdateConfigureContainerSystemRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PutConfigureContainerSystem(context.Background()).UpdateConfigureContainerSystemRequest(updateConfigureContainerSystemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PutConfigureContainerSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutConfigureContainerSystem`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PutConfigureContainerSystem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigureContainerSystemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateConfigureContainerSystemRequest** | [**UpdateConfigureContainerSystemRequest**](UpdateConfigureContainerSystemRequest.md) |  | 

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


## PutPluginInstanceFirewallRule

> Object PutPluginInstanceFirewallRule(ctx, id).CreatePluginInstancePresetFirewallRule(createPluginInstancePresetFirewallRule).Execute()

Create plugin instance firewall rule



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
    id := int32(56) // int32 | ID for plugin instance (container)
    createPluginInstancePresetFirewallRule := *openapiclient.NewCreatePluginInstancePresetFirewallRule("Preset_example") // CreatePluginInstancePresetFirewallRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PutPluginInstanceFirewallRule(context.Background(), id).CreatePluginInstancePresetFirewallRule(createPluginInstancePresetFirewallRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PutPluginInstanceFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPluginInstanceFirewallRule`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PutPluginInstanceFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPluginInstanceFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPluginInstancePresetFirewallRule** | [**CreatePluginInstancePresetFirewallRule**](CreatePluginInstancePresetFirewallRule.md) |  | 

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


## PutPluginInstanceUser

> Object PutPluginInstanceUser(ctx, id).PutPluginInstanceUserRequest(putPluginInstanceUserRequest).Execute()

Update or create a plugin instance user



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
    id := int32(56) // int32 | ID for plugin instance
    putPluginInstanceUserRequest := *openapiclient.NewPutPluginInstanceUserRequest() // PutPluginInstanceUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PutPluginInstanceUser(context.Background(), id).PutPluginInstanceUserRequest(putPluginInstanceUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PutPluginInstanceUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPluginInstanceUser`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PutPluginInstanceUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPluginInstanceUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putPluginInstanceUserRequest** | [**PutPluginInstanceUserRequest**](PutPluginInstanceUserRequest.md) |  | 

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


## PutStopContainer

> StopContainerDetailResponse PutStopContainer(ctx, uuid).Execute()

Stop container



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
    resp, r, err := apiClient.NetworkEdgePluginsApi.PutStopContainer(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PutStopContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutStopContainer`: StopContainerDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PutStopContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutStopContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StopContainerDetailResponse**](StopContainerDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateContainerImage

> UpdateContainerImageDetailResponse PutUpdateContainerImage(ctx, uuid).UpdateContainerImageRequest(updateContainerImageRequest).Execute()

Update container image



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
    updateContainerImageRequest := *openapiclient.NewUpdateContainerImageRequest("Name_example") // UpdateContainerImageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PutUpdateContainerImage(context.Background(), uuid).UpdateContainerImageRequest(updateContainerImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PutUpdateContainerImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateContainerImage`: UpdateContainerImageDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PutUpdateContainerImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | uuid of resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateContainerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateContainerImageRequest** | [**UpdateContainerImageRequest**](UpdateContainerImageRequest.md) |  | 

### Return type

[**UpdateContainerImageDetailResponse**](UpdateContainerImageDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateManagerConfig

> Object PutUpdateManagerConfig(ctx, id).Body(body).Execute()

Update Manager Config



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
    id := int32(56) // int32 | ID for plugin instance (container)
    body := interface{}(987) // interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PutUpdateManagerConfig(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PutUpdateManagerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateManagerConfig`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PutUpdateManagerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateManagerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

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


## PutUpdatePlugin

> Object PutUpdatePlugin(ctx, id).UpdatePluginRequest(updatePluginRequest).Execute()

Update plugin data



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
    id := int32(56) // int32 | ID for plugin (container image)
    updatePluginRequest := *openapiclient.NewUpdatePluginRequest() // UpdatePluginRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PutUpdatePlugin(context.Background(), id).UpdatePluginRequest(updatePluginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PutUpdatePlugin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdatePlugin`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PutUpdatePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin (container image) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdatePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePluginRequest** | [**UpdatePluginRequest**](UpdatePluginRequest.md) |  | 

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


## PutUpdatePluginInstanceConfigContent

> Object PutUpdatePluginInstanceConfigContent(ctx, id, slug).UpdatePluginInstanceConfigFileRequest(updatePluginInstanceConfigFileRequest).Execute()

Update plugin instance config file



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
    id := int32(56) // int32 | ID for plugin instance (container)
    slug := openapiclient.FirewallRuleTuple_inner{Int32: new(int32)} // FirewallRuleTupleInner | Either the index of the config file in the configuration  or the name of the config file. e.g. 0 or \"name\" 
    updatePluginInstanceConfigFileRequest := *openapiclient.NewUpdatePluginInstanceConfigFileRequest("Content_example") // UpdatePluginInstanceConfigFileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.PutUpdatePluginInstanceConfigContent(context.Background(), id, slug).UpdatePluginInstanceConfigFileRequest(updatePluginInstanceConfigFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.PutUpdatePluginInstanceConfigContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdatePluginInstanceConfigContent`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.PutUpdatePluginInstanceConfigContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 
**slug** | [**FirewallRuleTupleInner**](.md) | Either the index of the config file in the configuration  or the name of the config file. e.g. 0 or \&quot;name\&quot;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdatePluginInstanceConfigContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePluginInstanceConfigFileRequest** | [**UpdatePluginInstanceConfigFileRequest**](UpdatePluginInstanceConfigFileRequest.md) |  | 

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


## RevertPluginInstanceConfigFile

> Object RevertPluginInstanceConfigFile(ctx, id, slug).RevertPluginInstanceConfigFileRequest(revertPluginInstanceConfigFileRequest).Execute()

Revert plugin instance config file



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
    id := int32(56) // int32 | ID for plugin instance (container)
    slug := openapiclient.FirewallRuleTuple_inner{Int32: new(int32)} // FirewallRuleTupleInner | Either the index of the config file in the manager configuration  or the name of the confiig file. e.g. 0 or \"name\" 
    revertPluginInstanceConfigFileRequest := *openapiclient.NewRevertPluginInstanceConfigFileRequest(int32(123)) // RevertPluginInstanceConfigFileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.RevertPluginInstanceConfigFile(context.Background(), id, slug).RevertPluginInstanceConfigFileRequest(revertPluginInstanceConfigFileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.RevertPluginInstanceConfigFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevertPluginInstanceConfigFile`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.RevertPluginInstanceConfigFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 
**slug** | [**FirewallRuleTupleInner**](.md) | Either the index of the config file in the manager configuration  or the name of the confiig file. e.g. 0 or \&quot;name\&quot;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertPluginInstanceConfigFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revertPluginInstanceConfigFileRequest** | [**RevertPluginInstanceConfigFileRequest**](RevertPluginInstanceConfigFileRequest.md) |  | 

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


## RunPluginInstanceExecutableCommand

> Object RunPluginInstanceExecutableCommand(ctx, id).RunPluginInstanceExeCommandRequest(runPluginInstanceExeCommandRequest).Execute()

Run plugin instance executable command



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
    id := int32(56) // int32 | ID for plugin instance (container)
    runPluginInstanceExeCommandRequest := *openapiclient.NewRunPluginInstanceExeCommandRequest("Command_example") // RunPluginInstanceExeCommandRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.RunPluginInstanceExecutableCommand(context.Background(), id).RunPluginInstanceExeCommandRequest(runPluginInstanceExeCommandRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.RunPluginInstanceExecutableCommand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunPluginInstanceExecutableCommand`: Object
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.RunPluginInstanceExecutableCommand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunPluginInstanceExecutableCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runPluginInstanceExeCommandRequest** | [**RunPluginInstanceExeCommandRequest**](RunPluginInstanceExeCommandRequest.md) |  | 

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


## RunPluginInstanceProcessAction

> PluginManagerCommandResponse RunPluginInstanceProcessAction(ctx, id).RunPluginInstanceProcessActionRequest(runPluginInstanceProcessActionRequest).Execute()

Run plugin instance process action



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
    id := int32(56) // int32 | ID for plugin instance (container)
    runPluginInstanceProcessActionRequest := *openapiclient.NewRunPluginInstanceProcessActionRequest("Process_example", "Action_example") // RunPluginInstanceProcessActionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.RunPluginInstanceProcessAction(context.Background(), id).RunPluginInstanceProcessActionRequest(runPluginInstanceProcessActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.RunPluginInstanceProcessAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunPluginInstanceProcessAction`: PluginManagerCommandResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.RunPluginInstanceProcessAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID for plugin instance (container) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunPluginInstanceProcessActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **runPluginInstanceProcessActionRequest** | [**RunPluginInstanceProcessActionRequest**](RunPluginInstanceProcessActionRequest.md) |  | 

### Return type

[**PluginManagerCommandResponse**](PluginManagerCommandResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartPluginInstance

> PluginInstanceDetailResponse StartPluginInstance(ctx).StartPluginInstanceRequest(startPluginInstanceRequest).Execute()

Start new plugin instance from a plugin



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
    startPluginInstanceRequest := *openapiclient.NewStartPluginInstanceRequest("Name_example", int32(123)) // StartPluginInstanceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkEdgePluginsApi.StartPluginInstance(context.Background()).StartPluginInstanceRequest(startPluginInstanceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkEdgePluginsApi.StartPluginInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartPluginInstance`: PluginInstanceDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkEdgePluginsApi.StartPluginInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartPluginInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startPluginInstanceRequest** | [**StartPluginInstanceRequest**](StartPluginInstanceRequest.md) |  | 

### Return type

[**PluginInstanceDetailResponse**](PluginInstanceDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

