# \FirewallApi

All URIs are relative to *https://vns3-host:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEntryToFwset**](FirewallApi.md#AddEntryToFwset) | **Post** /v2/firewall/fwsets/{name}/entries | V2 Add entry to firewall fwset
[**AddRuleToGroup**](FirewallApi.md#AddRuleToGroup) | **Post** /v2/firewall/rules-groups/{name}/rules | V2 Add rule to existing rule group
[**AddRuleToSubtable**](FirewallApi.md#AddRuleToSubtable) | **Post** /v2/firewall/subtables/{name}/rules | V2 Add rule to firewall subtable
[**DeleteEntryFromFwset**](FirewallApi.md#DeleteEntryFromFwset) | **Delete** /v2/firewall/fwsets/{name}/entries | V2 Delete entry from firewall fwset
[**DeleteFirewallFwset**](FirewallApi.md#DeleteFirewallFwset) | **Delete** /v2/firewall/fwsets/{name} | V2 Delete firewall fwset
[**DeleteFirewallFwsetV1**](FirewallApi.md#DeleteFirewallFwsetV1) | **Delete** /firewall/fwsets | V1 Delete firewall Fwset
[**DeleteFirewallRule**](FirewallApi.md#DeleteFirewallRule) | **Delete** /v2/firewall/rules/{id} | V2 delete firewall rule
[**DeleteFirewallRuleByPositionV1**](FirewallApi.md#DeleteFirewallRuleByPositionV1) | **Delete** /firewall/rules/{position} | V1 Delete firewall rule by position
[**DeleteFirewallRuleByRule**](FirewallApi.md#DeleteFirewallRuleByRule) | **Delete** /firewall/rules | V1 Delete firewall rule
[**DeleteFirewallRuleGroup**](FirewallApi.md#DeleteFirewallRuleGroup) | **Delete** /v2/firewall/rules-groups/{name} | V2 Delete firewall rule group
[**DeleteFirewallSubgroupV1**](FirewallApi.md#DeleteFirewallSubgroupV1) | **Delete** /firewall/rules/subgroup | V1 Delete firewall subgroup
[**DeleteFirewallSubtable**](FirewallApi.md#DeleteFirewallSubtable) | **Delete** /v2/firewall/subtables/{name} | V2 Delete firewall subtable
[**GetFirewallFwset**](FirewallApi.md#GetFirewallFwset) | **Get** /v2/firewall/fwsets/{name} | V2 Read firewall fwset data
[**GetFirewallFwsets**](FirewallApi.md#GetFirewallFwsets) | **Get** /v2/firewall/fwsets | V2 Get firewall fwsets
[**GetFirewallFwsetsV1**](FirewallApi.md#GetFirewallFwsetsV1) | **Get** /firewall/fwsets | V1 Get firewall Fwsets
[**GetFirewallRuleGroup**](FirewallApi.md#GetFirewallRuleGroup) | **Get** /v2/firewall/rules-groups/{name} | V2 Read rule group details
[**GetFirewallRuleGroups**](FirewallApi.md#GetFirewallRuleGroups) | **Get** /v2/firewall/rules-groups | V2 Get firewall rule groups
[**GetFirewallRuleSubgroupsV1**](FirewallApi.md#GetFirewallRuleSubgroupsV1) | **Get** /firewall/rules/subgroup | V1 Get firewall subgroups
[**GetFirewallRules**](FirewallApi.md#GetFirewallRules) | **Get** /v2/firewall/rules | V2 Get firewall rules
[**GetFirewallRulesV1**](FirewallApi.md#GetFirewallRulesV1) | **Get** /firewall/rules | V1 Get firewall rules
[**GetFirewallSubtable**](FirewallApi.md#GetFirewallSubtable) | **Get** /v2/firewall/subtables/{name} | V2 Read firewall subtable data
[**GetFirewallSubtables**](FirewallApi.md#GetFirewallSubtables) | **Get** /v2/firewall/subtables | V2 Get firewall subtables
[**ImportFirewallRules**](FirewallApi.md#ImportFirewallRules) | **Put** /v2/firewall/rules/import | V2 Import firewall rules
[**PostCreateFirewallFwset**](FirewallApi.md#PostCreateFirewallFwset) | **Post** /v2/firewall/fwsets | V2 Create firewall fwset
[**PostCreateFirewallFwsetV1**](FirewallApi.md#PostCreateFirewallFwsetV1) | **Post** /firewall/fwsets | V1 Create firewall FwSet
[**PostCreateFirewallRule**](FirewallApi.md#PostCreateFirewallRule) | **Post** /v2/firewall/rules | V2 Create firewall rule
[**PostCreateFirewallRuleGroup**](FirewallApi.md#PostCreateFirewallRuleGroup) | **Post** /v2/firewall/rules-groups | V2 Create firewall rule group
[**PostCreateFirewallRuleV1**](FirewallApi.md#PostCreateFirewallRuleV1) | **Post** /firewall/rules | V1 Create firewall rule
[**PostCreateFirewallSubgroupV1**](FirewallApi.md#PostCreateFirewallSubgroupV1) | **Post** /firewall/rules/subgroup | V1 Create firewall subgroup
[**PostCreateFirewallSubtable**](FirewallApi.md#PostCreateFirewallSubtable) | **Post** /v2/firewall/subtables | V2 Create firewall subtable
[**PutFirewallAction**](FirewallApi.md#PutFirewallAction) | **Put** /v2/firewall/actions | V2 Put firewall action
[**PutFirewallActionV1**](FirewallApi.md#PutFirewallActionV1) | **Put** /firewall/actions | V1 Put firewall action
[**PutOverwriteFirewall**](FirewallApi.md#PutOverwriteFirewall) | **Put** /v2/firewall/rules | V2 Put firewall
[**PutOverwriteFirewallV1**](FirewallApi.md#PutOverwriteFirewallV1) | **Put** /firewall | V1 Put firewall
[**PutReinitializeFirewallSubgroupsV1**](FirewallApi.md#PutReinitializeFirewallSubgroupsV1) | **Put** /firewall/rules/subgroup | V1 Reload firewall subgroups
[**PutReinitializeFwsetsV1**](FirewallApi.md#PutReinitializeFwsetsV1) | **Put** /firewall/fwsets | V1 Reload all firewall Fwsets
[**PutUpdateFirewallFwset**](FirewallApi.md#PutUpdateFirewallFwset) | **Put** /v2/firewall/fwsets/{name} | V2 Update firewall fwset
[**PutUpdateFirewallRule**](FirewallApi.md#PutUpdateFirewallRule) | **Put** /v2/firewall/rules/{id} | V2 Update firewall rule
[**PutUpdateFirewallRuleGroup**](FirewallApi.md#PutUpdateFirewallRuleGroup) | **Put** /v2/firewall/rules-groups/{name} | V2 Update rule group data
[**PutUpdateFirewallSubtable**](FirewallApi.md#PutUpdateFirewallSubtable) | **Put** /v2/firewall/subtables/{name} | V2 Update firewall subtable



## AddEntryToFwset

> Object AddEntryToFwset(ctx, name).AddFirewallEntryToFwsetRequest(addFirewallEntryToFwsetRequest).Execute()

V2 Add entry to firewall fwset



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
    name := "name_example" // string | unique fwset name
    addFirewallEntryToFwsetRequest := *openapiclient.NewAddFirewallEntryToFwsetRequest("Entry_example") // AddFirewallEntryToFwsetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.AddEntryToFwset(context.Background(), name).AddFirewallEntryToFwsetRequest(addFirewallEntryToFwsetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.AddEntryToFwset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEntryToFwset`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.AddEntryToFwset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique fwset name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEntryToFwsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addFirewallEntryToFwsetRequest** | [**AddFirewallEntryToFwsetRequest**](AddFirewallEntryToFwsetRequest.md) |  | 

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


## AddRuleToGroup

> Object AddRuleToGroup(ctx, name).AddFirewallRuleToGroupRequest(addFirewallRuleToGroupRequest).Execute()

V2 Add rule to existing rule group



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
    name := "name_example" // string | unique rule group name
    addFirewallRuleToGroupRequest := *openapiclient.NewAddFirewallRuleToGroupRequest("RuleId_example") // AddFirewallRuleToGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.AddRuleToGroup(context.Background(), name).AddFirewallRuleToGroupRequest(addFirewallRuleToGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.AddRuleToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRuleToGroup`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.AddRuleToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique rule group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRuleToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addFirewallRuleToGroupRequest** | [**AddFirewallRuleToGroupRequest**](AddFirewallRuleToGroupRequest.md) |  | 

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


## AddRuleToSubtable

> Object AddRuleToSubtable(ctx, name).AddFirewallRuleToSubtableRequest(addFirewallRuleToSubtableRequest).Execute()

V2 Add rule to firewall subtable



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
    name := "name_example" // string | unique subtable name
    addFirewallRuleToSubtableRequest := *openapiclient.NewAddFirewallRuleToSubtableRequest("Rule_example") // AddFirewallRuleToSubtableRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.AddRuleToSubtable(context.Background(), name).AddFirewallRuleToSubtableRequest(addFirewallRuleToSubtableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.AddRuleToSubtable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRuleToSubtable`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.AddRuleToSubtable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique subtable name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRuleToSubtableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addFirewallRuleToSubtableRequest** | [**AddFirewallRuleToSubtableRequest**](AddFirewallRuleToSubtableRequest.md) |  | 

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


## DeleteEntryFromFwset

> Object DeleteEntryFromFwset(ctx, name).RemoveFirewallEntryToFwsetRequest(removeFirewallEntryToFwsetRequest).Execute()

V2 Delete entry from firewall fwset



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
    name := "name_example" // string | unique fwset name
    removeFirewallEntryToFwsetRequest := *openapiclient.NewRemoveFirewallEntryToFwsetRequest("Entry_example") // RemoveFirewallEntryToFwsetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteEntryFromFwset(context.Background(), name).RemoveFirewallEntryToFwsetRequest(removeFirewallEntryToFwsetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteEntryFromFwset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEntryFromFwset`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteEntryFromFwset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique fwset name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntryFromFwsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeFirewallEntryToFwsetRequest** | [**RemoveFirewallEntryToFwsetRequest**](RemoveFirewallEntryToFwsetRequest.md) |  | 

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


## DeleteFirewallFwset

> Object DeleteFirewallFwset(ctx, name).Force(force).Execute()

V2 Delete firewall fwset



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
    name := "name_example" // string | unique fwset name
    force := true // bool | Force delete if fwset is associated with rules or other fwsets (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallFwset(context.Background(), name).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallFwset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallFwset`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallFwset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique fwset name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallFwsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force delete if fwset is associated with rules or other fwsets | [default to false]

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


## DeleteFirewallFwsetV1

> Object DeleteFirewallFwsetV1(ctx).FirewallFwsetDeleteRequest(firewallFwsetDeleteRequest).Execute()

V1 Delete firewall Fwset



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
    firewallFwsetDeleteRequest := *openapiclient.NewFirewallFwsetDeleteRequest() // FirewallFwsetDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallFwsetV1(context.Background()).FirewallFwsetDeleteRequest(firewallFwsetDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallFwsetV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallFwsetV1`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallFwsetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallFwsetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallFwsetDeleteRequest** | [**FirewallFwsetDeleteRequest**](FirewallFwsetDeleteRequest.md) |  | 

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


## DeleteFirewallRule

> Object DeleteFirewallRule(ctx, id).Execute()

V2 delete firewall rule



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
    id := "id_example" // string | Rule ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallRule`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleRequest struct via the builder pattern


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


## DeleteFirewallRuleByPositionV1

> Object DeleteFirewallRuleByPositionV1(ctx, position).Execute()

V1 Delete firewall rule by position



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
    position := int32(56) // int32 | index position for firewall rule, 0 is first

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallRuleByPositionV1(context.Background(), position).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallRuleByPositionV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallRuleByPositionV1`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallRuleByPositionV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**position** | **int32** | index position for firewall rule, 0 is first | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleByPositionV1Request struct via the builder pattern


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


## DeleteFirewallRuleByRule

> Object DeleteFirewallRuleByRule(ctx).DeleteFirewallRuleRequest(deleteFirewallRuleRequest).Execute()

V1 Delete firewall rule



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
    deleteFirewallRuleRequest := *openapiclient.NewDeleteFirewallRuleRequest("Rule_example") // DeleteFirewallRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallRuleByRule(context.Background()).DeleteFirewallRuleRequest(deleteFirewallRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallRuleByRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallRuleByRule`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallRuleByRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleByRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteFirewallRuleRequest** | [**DeleteFirewallRuleRequest**](DeleteFirewallRuleRequest.md) |  | 

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


## DeleteFirewallRuleGroup

> Object DeleteFirewallRuleGroup(ctx, name).Force(force).Execute()

V2 Delete firewall rule group



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
    name := "name_example" // string | unique rule group name
    force := true // bool | Force delete if rule group has rules associated (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallRuleGroup(context.Background(), name).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallRuleGroup`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique rule group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force delete if rule group has rules associated | [default to false]

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


## DeleteFirewallSubgroupV1

> Object DeleteFirewallSubgroupV1(ctx).FirewallSubgroupDeleteRequest(firewallSubgroupDeleteRequest).Execute()

V1 Delete firewall subgroup



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
    firewallSubgroupDeleteRequest := *openapiclient.NewFirewallSubgroupDeleteRequest() // FirewallSubgroupDeleteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallSubgroupV1(context.Background()).FirewallSubgroupDeleteRequest(firewallSubgroupDeleteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallSubgroupV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallSubgroupV1`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallSubgroupV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallSubgroupV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallSubgroupDeleteRequest** | [**FirewallSubgroupDeleteRequest**](FirewallSubgroupDeleteRequest.md) |  | 

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


## DeleteFirewallSubtable

> Object DeleteFirewallSubtable(ctx, name).Force(force).Execute()

V2 Delete firewall subtable



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
    name := "name_example" // string | unique subtable name
    force := true // bool | Force delete if subtable is associated with rules in main firewall policy (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.DeleteFirewallSubtable(context.Background(), name).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.DeleteFirewallSubtable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFirewallSubtable`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.DeleteFirewallSubtable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique subtable name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallSubtableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | Force delete if subtable is associated with rules in main firewall policy | [default to false]

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


## GetFirewallFwset

> interface{} GetFirewallFwset(ctx, name).Osview(osview).Execute()

V2 Read firewall fwset data



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
    name := "name_example" // string | unique fwset name
    osview := true // bool | Show operating system level entries (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetFirewallFwset(context.Background(), name).Osview(osview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallFwset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallFwset`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallFwset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique fwset name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallFwsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **osview** | **bool** | Show operating system level entries | [default to false]

### Return type

**interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallFwsets

> interface{} GetFirewallFwsets(ctx).Execute()

V2 Get firewall fwsets



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
    resp, r, err := apiClient.FirewallApi.GetFirewallFwsets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallFwsets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallFwsets`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallFwsets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallFwsetsRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallFwsetsV1

> FirewallFwSetListResponse GetFirewallFwsetsV1(ctx).Name(name).Verbose(verbose).Execute()

V1 Get firewall Fwsets



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
    name := "name_example" // string | name of resource (optional)
    verbose := true // bool | True for verbose output (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetFirewallFwsetsV1(context.Background()).Name(name).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallFwsetsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallFwsetsV1`: FirewallFwSetListResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallFwsetsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallFwsetsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | name of resource | 
 **verbose** | **bool** | True for verbose output | [default to true]

### Return type

[**FirewallFwSetListResponse**](FirewallFwSetListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRuleGroup

> Object GetFirewallRuleGroup(ctx, name).Execute()

V2 Read rule group details



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
    name := "name_example" // string | unique rule group name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetFirewallRuleGroup(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallRuleGroup`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique rule group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRuleGroupRequest struct via the builder pattern


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


## GetFirewallRuleGroups

> RuleGroupsListResponse GetFirewallRuleGroups(ctx).Execute()

V2 Get firewall rule groups



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
    resp, r, err := apiClient.FirewallApi.GetFirewallRuleGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallRuleGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallRuleGroups`: RuleGroupsListResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallRuleGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRuleGroupsRequest struct via the builder pattern


### Return type

[**RuleGroupsListResponse**](RuleGroupsListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRuleSubgroupsV1

> FirewallSubgroupListResponse GetFirewallRuleSubgroupsV1(ctx).Name(name).Verbose(verbose).Execute()

V1 Get firewall subgroups



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
    name := "name_example" // string | name of resource (optional)
    verbose := true // bool | True for verbose output (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetFirewallRuleSubgroupsV1(context.Background()).Name(name).Verbose(verbose).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallRuleSubgroupsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallRuleSubgroupsV1`: FirewallSubgroupListResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallRuleSubgroupsV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRuleSubgroupsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | name of resource | 
 **verbose** | **bool** | True for verbose output | [default to true]

### Return type

[**FirewallSubgroupListResponse**](FirewallSubgroupListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRules

> FirewallRuleV2ListResponse GetFirewallRules(ctx).State(state).Groups(groups).Osview(osview).Tables(tables).Execute()

V2 Get firewall rules



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
    state := "state_example" // string | Filter rules by state, active or disabled (optional)
    groups := "groups_example" // string | Filter by groups. Accepts csv. (optional)
    osview := true // bool | Show operating system level rules (optional) (default to false)
    tables := "tables_example" // string | Filter by tables. Accepts csv. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetFirewallRules(context.Background()).State(state).Groups(groups).Osview(osview).Tables(tables).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallRules`: FirewallRuleV2ListResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | Filter rules by state, active or disabled | 
 **groups** | **string** | Filter by groups. Accepts csv. | 
 **osview** | **bool** | Show operating system level rules | [default to false]
 **tables** | **string** | Filter by tables. Accepts csv. | 

### Return type

[**FirewallRuleV2ListResponse**](FirewallRuleV2ListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRulesV1

> FirewallRuleListResponse GetFirewallRulesV1(ctx).Execute()

V1 Get firewall rules



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
    resp, r, err := apiClient.FirewallApi.GetFirewallRulesV1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallRulesV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallRulesV1`: FirewallRuleListResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallRulesV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRulesV1Request struct via the builder pattern


### Return type

[**FirewallRuleListResponse**](FirewallRuleListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallSubtable

> FirewallSubtableDetailResponse GetFirewallSubtable(ctx, name).Osview(osview).Execute()

V2 Read firewall subtable data



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
    name := "name_example" // string | unique subtable name
    osview := true // bool | Show operating system level rules (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.GetFirewallSubtable(context.Background(), name).Osview(osview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallSubtable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallSubtable`: FirewallSubtableDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallSubtable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique subtable name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallSubtableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **osview** | **bool** | Show operating system level rules | [default to false]

### Return type

[**FirewallSubtableDetailResponse**](FirewallSubtableDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallSubtables

> FirewallSubtableListResponse GetFirewallSubtables(ctx).Execute()

V2 Get firewall subtables



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
    resp, r, err := apiClient.FirewallApi.GetFirewallSubtables(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.GetFirewallSubtables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallSubtables`: FirewallSubtableListResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.GetFirewallSubtables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallSubtablesRequest struct via the builder pattern


### Return type

[**FirewallSubtableListResponse**](FirewallSubtableListResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportFirewallRules

> Object ImportFirewallRules(ctx).ImportFirewallRulesRequest(importFirewallRulesRequest).Execute()

V2 Import firewall rules



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
    importFirewallRulesRequest := *openapiclient.NewImportFirewallRulesRequest("Rules_example") // ImportFirewallRulesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.ImportFirewallRules(context.Background()).ImportFirewallRulesRequest(importFirewallRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.ImportFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportFirewallRules`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.ImportFirewallRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importFirewallRulesRequest** | [**ImportFirewallRulesRequest**](ImportFirewallRulesRequest.md) |  | 

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


## PostCreateFirewallFwset

> FirewallFwsetSaveResponse PostCreateFirewallFwset(ctx).CreateFirewallFwsetRequest(createFirewallFwsetRequest).Execute()

V2 Create firewall fwset



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
    createFirewallFwsetRequest := *openapiclient.NewCreateFirewallFwsetRequest("Name_example", "Type_example") // CreateFirewallFwsetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PostCreateFirewallFwset(context.Background()).CreateFirewallFwsetRequest(createFirewallFwsetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PostCreateFirewallFwset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateFirewallFwset`: FirewallFwsetSaveResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PostCreateFirewallFwset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateFirewallFwsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirewallFwsetRequest** | [**CreateFirewallFwsetRequest**](CreateFirewallFwsetRequest.md) |  | 

### Return type

[**FirewallFwsetSaveResponse**](FirewallFwsetSaveResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateFirewallFwsetV1

> PostCreateFirewallFwsetV1200Response PostCreateFirewallFwsetV1(ctx).CreateFwsetRequest(createFwsetRequest).Execute()

V1 Create firewall FwSet



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
    createFwsetRequest := *openapiclient.NewCreateFwsetRequest() // CreateFwsetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PostCreateFirewallFwsetV1(context.Background()).CreateFwsetRequest(createFwsetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PostCreateFirewallFwsetV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateFirewallFwsetV1`: PostCreateFirewallFwsetV1200Response
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PostCreateFirewallFwsetV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateFirewallFwsetV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFwsetRequest** | [**CreateFwsetRequest**](CreateFwsetRequest.md) |  | 

### Return type

[**PostCreateFirewallFwsetV1200Response**](PostCreateFirewallFwsetV1200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateFirewallRule

> PostCreateFirewallRule201Response PostCreateFirewallRule(ctx).CreateFirewallRuleRequestV2(createFirewallRuleRequestV2).Execute()

V2 Create firewall rule



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
    createFirewallRuleRequestV2 := openapiclient.CreateFirewallRuleRequestV2{Interface{}: new(interface{})} // CreateFirewallRuleRequestV2 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PostCreateFirewallRule(context.Background()).CreateFirewallRuleRequestV2(createFirewallRuleRequestV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PostCreateFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateFirewallRule`: PostCreateFirewallRule201Response
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PostCreateFirewallRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirewallRuleRequestV2** | [**CreateFirewallRuleRequestV2**](CreateFirewallRuleRequestV2.md) |  | 

### Return type

[**PostCreateFirewallRule201Response**](PostCreateFirewallRule201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateFirewallRuleGroup

> RuleGroupDetailResponse PostCreateFirewallRuleGroup(ctx).CreateFirewallRuleGroupRequest(createFirewallRuleGroupRequest).Execute()

V2 Create firewall rule group



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
    createFirewallRuleGroupRequest := *openapiclient.NewCreateFirewallRuleGroupRequest("Name_example") // CreateFirewallRuleGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PostCreateFirewallRuleGroup(context.Background()).CreateFirewallRuleGroupRequest(createFirewallRuleGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PostCreateFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateFirewallRuleGroup`: RuleGroupDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PostCreateFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirewallRuleGroupRequest** | [**CreateFirewallRuleGroupRequest**](CreateFirewallRuleGroupRequest.md) |  | 

### Return type

[**RuleGroupDetailResponse**](RuleGroupDetailResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateFirewallRuleV1

> FirewallRuleOperationResponse PostCreateFirewallRuleV1(ctx).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()

V1 Create firewall rule



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
    createFirewallRuleRequest := *openapiclient.NewCreateFirewallRuleRequest("Rule_example") // CreateFirewallRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PostCreateFirewallRuleV1(context.Background()).CreateFirewallRuleRequest(createFirewallRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PostCreateFirewallRuleV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateFirewallRuleV1`: FirewallRuleOperationResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PostCreateFirewallRuleV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateFirewallRuleV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirewallRuleRequest** | [**CreateFirewallRuleRequest**](CreateFirewallRuleRequest.md) |  | 

### Return type

[**FirewallRuleOperationResponse**](FirewallRuleOperationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateFirewallSubgroupV1

> PostCreateFirewallSubgroupV1200Response PostCreateFirewallSubgroupV1(ctx).CreateSubgroupRequest(createSubgroupRequest).Execute()

V1 Create firewall subgroup



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
    createSubgroupRequest := openapiclient.CreateSubgroupRequest{Interface{}: new(interface{})} // CreateSubgroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PostCreateFirewallSubgroupV1(context.Background()).CreateSubgroupRequest(createSubgroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PostCreateFirewallSubgroupV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateFirewallSubgroupV1`: PostCreateFirewallSubgroupV1200Response
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PostCreateFirewallSubgroupV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateFirewallSubgroupV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubgroupRequest** | [**CreateSubgroupRequest**](CreateSubgroupRequest.md) |  | 

### Return type

[**PostCreateFirewallSubgroupV1200Response**](PostCreateFirewallSubgroupV1200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCreateFirewallSubtable

> FirewallSubtableCreateResponse PostCreateFirewallSubtable(ctx).CreateFirewallSubtableRequest(createFirewallSubtableRequest).Execute()

V2 Create firewall subtable



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
    createFirewallSubtableRequest := *openapiclient.NewCreateFirewallSubtableRequest("Name_example", "Type_example") // CreateFirewallSubtableRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PostCreateFirewallSubtable(context.Background()).CreateFirewallSubtableRequest(createFirewallSubtableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PostCreateFirewallSubtable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCreateFirewallSubtable`: FirewallSubtableCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PostCreateFirewallSubtable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCreateFirewallSubtableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFirewallSubtableRequest** | [**CreateFirewallSubtableRequest**](CreateFirewallSubtableRequest.md) |  | 

### Return type

[**FirewallSubtableCreateResponse**](FirewallSubtableCreateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFirewallAction

> Object PutFirewallAction(ctx).FirewallActionRequest(firewallActionRequest).Execute()

V2 Put firewall action



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
    firewallActionRequest := *openapiclient.NewFirewallActionRequest("Action_example") // FirewallActionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutFirewallAction(context.Background()).FirewallActionRequest(firewallActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutFirewallAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFirewallAction`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PutFirewallAction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutFirewallActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallActionRequest** | [**FirewallActionRequest**](FirewallActionRequest.md) |  | 

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


## PutFirewallActionV1

> Object PutFirewallActionV1(ctx).FirewallActionRequest(firewallActionRequest).Execute()

V1 Put firewall action



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
    firewallActionRequest := *openapiclient.NewFirewallActionRequest("Action_example") // FirewallActionRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutFirewallActionV1(context.Background()).FirewallActionRequest(firewallActionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutFirewallActionV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutFirewallActionV1`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PutFirewallActionV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutFirewallActionV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **firewallActionRequest** | [**FirewallActionRequest**](FirewallActionRequest.md) |  | 

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


## PutOverwriteFirewall

> FirewallBulkWriteResponse PutOverwriteFirewall(ctx).OverwriteRequestV2(overwriteRequestV2).Execute()

V2 Put firewall



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
    overwriteRequestV2 := *openapiclient.NewOverwriteRequestV2(*openapiclient.NewOverwriteRequestV2Rules("Rule_example")) // OverwriteRequestV2 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutOverwriteFirewall(context.Background()).OverwriteRequestV2(overwriteRequestV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutOverwriteFirewall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutOverwriteFirewall`: FirewallBulkWriteResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PutOverwriteFirewall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutOverwriteFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **overwriteRequestV2** | [**OverwriteRequestV2**](OverwriteRequestV2.md) |  | 

### Return type

[**FirewallBulkWriteResponse**](FirewallBulkWriteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOverwriteFirewallV1

> TaskTokenResponse PutOverwriteFirewallV1(ctx).OverwriteRequest(overwriteRequest).Execute()

V1 Put firewall



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
    overwriteRequest := *openapiclient.NewOverwriteRequest("Rules_example") // OverwriteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutOverwriteFirewallV1(context.Background()).OverwriteRequest(overwriteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutOverwriteFirewallV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutOverwriteFirewallV1`: TaskTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PutOverwriteFirewallV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutOverwriteFirewallV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **overwriteRequest** | [**OverwriteRequest**](OverwriteRequest.md) |  | 

### Return type

[**TaskTokenResponse**](TaskTokenResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutReinitializeFirewallSubgroupsV1

> PutReinitializeFirewallSubgroupsV1(ctx).ReinitRequest(reinitRequest).Execute()

V1 Reload firewall subgroups



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
    reinitRequest := *openapiclient.NewReinitRequest() // ReinitRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutReinitializeFirewallSubgroupsV1(context.Background()).ReinitRequest(reinitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutReinitializeFirewallSubgroupsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutReinitializeFirewallSubgroupsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reinitRequest** | [**ReinitRequest**](ReinitRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutReinitializeFwsetsV1

> PutReinitializeFwsetsV1(ctx).ReinitRequest(reinitRequest).Execute()

V1 Reload all firewall Fwsets



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
    reinitRequest := *openapiclient.NewReinitRequest() // ReinitRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutReinitializeFwsetsV1(context.Background()).ReinitRequest(reinitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutReinitializeFwsetsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutReinitializeFwsetsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reinitRequest** | [**ReinitRequest**](ReinitRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateFirewallFwset

> Object PutUpdateFirewallFwset(ctx, name).UpdateFirewallFwsetRequest(updateFirewallFwsetRequest).Execute()

V2 Update firewall fwset



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
    name := "name_example" // string | unique fwset name
    updateFirewallFwsetRequest := *openapiclient.NewUpdateFirewallFwsetRequest() // UpdateFirewallFwsetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutUpdateFirewallFwset(context.Background(), name).UpdateFirewallFwsetRequest(updateFirewallFwsetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutUpdateFirewallFwset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateFirewallFwset`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PutUpdateFirewallFwset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique fwset name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateFirewallFwsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFirewallFwsetRequest** | [**UpdateFirewallFwsetRequest**](UpdateFirewallFwsetRequest.md) |  | 

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


## PutUpdateFirewallRule

> FirewallRuleV2UpdateResponse PutUpdateFirewallRule(ctx, id).UpdateFirewallRuleRequest(updateFirewallRuleRequest).Execute()

V2 Update firewall rule



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
    id := "id_example" // string | Rule ID
    updateFirewallRuleRequest := *openapiclient.NewUpdateFirewallRuleRequest() // UpdateFirewallRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutUpdateFirewallRule(context.Background(), id).UpdateFirewallRuleRequest(updateFirewallRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutUpdateFirewallRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateFirewallRule`: FirewallRuleV2UpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PutUpdateFirewallRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Rule ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateFirewallRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFirewallRuleRequest** | [**UpdateFirewallRuleRequest**](UpdateFirewallRuleRequest.md) |  | 

### Return type

[**FirewallRuleV2UpdateResponse**](FirewallRuleV2UpdateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUpdateFirewallRuleGroup

> Object PutUpdateFirewallRuleGroup(ctx, name).UpdateFirewallRuleGroupRequest(updateFirewallRuleGroupRequest).Execute()

V2 Update rule group data



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
    name := "name_example" // string | unique rule group name
    updateFirewallRuleGroupRequest := *openapiclient.NewUpdateFirewallRuleGroupRequest() // UpdateFirewallRuleGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutUpdateFirewallRuleGroup(context.Background(), name).UpdateFirewallRuleGroupRequest(updateFirewallRuleGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutUpdateFirewallRuleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateFirewallRuleGroup`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PutUpdateFirewallRuleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique rule group name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateFirewallRuleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFirewallRuleGroupRequest** | [**UpdateFirewallRuleGroupRequest**](UpdateFirewallRuleGroupRequest.md) |  | 

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


## PutUpdateFirewallSubtable

> Object PutUpdateFirewallSubtable(ctx, name).UpdateFirewallSubtableRequest(updateFirewallSubtableRequest).Execute()

V2 Update firewall subtable



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
    name := "name_example" // string | unique subtable name
    updateFirewallSubtableRequest := *openapiclient.NewUpdateFirewallSubtableRequest() // UpdateFirewallSubtableRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallApi.PutUpdateFirewallSubtable(context.Background(), name).UpdateFirewallSubtableRequest(updateFirewallSubtableRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallApi.PutUpdateFirewallSubtable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUpdateFirewallSubtable`: Object
    fmt.Fprintf(os.Stdout, "Response from `FirewallApi.PutUpdateFirewallSubtable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | unique subtable name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUpdateFirewallSubtableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFirewallSubtableRequest** | [**UpdateFirewallSubtableRequest**](UpdateFirewallSubtableRequest.md) |  | 

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

