# ContainerLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Logs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewContainerLogs

`func NewContainerLogs() *ContainerLogs`

NewContainerLogs instantiates a new ContainerLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerLogsWithDefaults

`func NewContainerLogsWithDefaults() *ContainerLogs`

NewContainerLogsWithDefaults instantiates a new ContainerLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ContainerLogs) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ContainerLogs) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ContainerLogs) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ContainerLogs) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetLogs

`func (o *ContainerLogs) GetLogs() []string`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ContainerLogs) GetLogsOk() (*[]string, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ContainerLogs) SetLogs(v []string)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ContainerLogs) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


