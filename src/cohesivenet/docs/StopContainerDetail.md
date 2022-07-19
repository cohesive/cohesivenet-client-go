# StopContainerDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Container system status for allocated container | [optional] 

## Methods

### NewStopContainerDetail

`func NewStopContainerDetail() *StopContainerDetail`

NewStopContainerDetail instantiates a new StopContainerDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopContainerDetailWithDefaults

`func NewStopContainerDetailWithDefaults() *StopContainerDetail`

NewStopContainerDetailWithDefaults instantiates a new StopContainerDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *StopContainerDetail) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StopContainerDetail) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StopContainerDetail) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StopContainerDetail) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStatus

`func (o *StopContainerDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StopContainerDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StopContainerDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StopContainerDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


