# RunningContainersDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to [**[]Container**](Container.md) |  | [optional] 

## Methods

### NewRunningContainersDetail

`func NewRunningContainersDetail() *RunningContainersDetail`

NewRunningContainersDetail instantiates a new RunningContainersDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunningContainersDetailWithDefaults

`func NewRunningContainersDetailWithDefaults() *RunningContainersDetail`

NewRunningContainersDetailWithDefaults instantiates a new RunningContainersDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *RunningContainersDetail) GetContainers() []Container`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *RunningContainersDetail) GetContainersOk() (*[]Container, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *RunningContainersDetail) SetContainers(v []Container)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *RunningContainersDetail) HasContainers() bool`

HasContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


