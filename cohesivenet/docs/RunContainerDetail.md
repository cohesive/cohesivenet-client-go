# RunContainerDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**ContainerStarted** | Pointer to **bool** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Container system status for allocated container | [optional] 

## Methods

### NewRunContainerDetail

`func NewRunContainerDetail() *RunContainerDetail`

NewRunContainerDetail instantiates a new RunContainerDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunContainerDetailWithDefaults

`func NewRunContainerDetailWithDefaults() *RunContainerDetail`

NewRunContainerDetailWithDefaults instantiates a new RunContainerDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RunContainerDetail) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RunContainerDetail) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RunContainerDetail) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RunContainerDetail) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetContainerStarted

`func (o *RunContainerDetail) GetContainerStarted() bool`

GetContainerStarted returns the ContainerStarted field if non-nil, zero value otherwise.

### GetContainerStartedOk

`func (o *RunContainerDetail) GetContainerStartedOk() (*bool, bool)`

GetContainerStartedOk returns a tuple with the ContainerStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerStarted

`func (o *RunContainerDetail) SetContainerStarted(v bool)`

SetContainerStarted sets ContainerStarted field to given value.

### HasContainerStarted

`func (o *RunContainerDetail) HasContainerStarted() bool`

HasContainerStarted returns a boolean if a field has been set.

### GetIpAddress

`func (o *RunContainerDetail) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *RunContainerDetail) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *RunContainerDetail) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *RunContainerDetail) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetStatus

`func (o *RunContainerDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RunContainerDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RunContainerDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RunContainerDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


