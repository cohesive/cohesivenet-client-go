# ContainerSystemStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to **string** | Local network in CIDR notation container system is running on | [optional] 
**Running** | Pointer to [**ContainerSystemStatusRunning**](ContainerSystemStatusRunning.md) |  | [optional] 

## Methods

### NewContainerSystemStatus

`func NewContainerSystemStatus() *ContainerSystemStatus`

NewContainerSystemStatus instantiates a new ContainerSystemStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerSystemStatusWithDefaults

`func NewContainerSystemStatusWithDefaults() *ContainerSystemStatus`

NewContainerSystemStatusWithDefaults instantiates a new ContainerSystemStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *ContainerSystemStatus) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ContainerSystemStatus) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ContainerSystemStatus) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ContainerSystemStatus) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRunning

`func (o *ContainerSystemStatus) GetRunning() ContainerSystemStatusRunning`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ContainerSystemStatus) GetRunningOk() (*ContainerSystemStatusRunning, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ContainerSystemStatus) SetRunning(v ContainerSystemStatusRunning)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *ContainerSystemStatus) HasRunning() bool`

HasRunning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


