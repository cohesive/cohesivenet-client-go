# DeletePluginInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Force** | Pointer to **bool** | Stop and delete if running | [optional] 

## Methods

### NewDeletePluginInstanceRequest

`func NewDeletePluginInstanceRequest() *DeletePluginInstanceRequest`

NewDeletePluginInstanceRequest instantiates a new DeletePluginInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePluginInstanceRequestWithDefaults

`func NewDeletePluginInstanceRequestWithDefaults() *DeletePluginInstanceRequest`

NewDeletePluginInstanceRequestWithDefaults instantiates a new DeletePluginInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForce

`func (o *DeletePluginInstanceRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *DeletePluginInstanceRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *DeletePluginInstanceRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *DeletePluginInstanceRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


