# DeleteContainerImageDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**ImageDeleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewDeleteContainerImageDetail

`func NewDeleteContainerImageDetail() *DeleteContainerImageDetail`

NewDeleteContainerImageDetail instantiates a new DeleteContainerImageDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteContainerImageDetailWithDefaults

`func NewDeleteContainerImageDetailWithDefaults() *DeleteContainerImageDetail`

NewDeleteContainerImageDetailWithDefaults instantiates a new DeleteContainerImageDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DeleteContainerImageDetail) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeleteContainerImageDetail) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeleteContainerImageDetail) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeleteContainerImageDetail) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetImageDeleted

`func (o *DeleteContainerImageDetail) GetImageDeleted() bool`

GetImageDeleted returns the ImageDeleted field if non-nil, zero value otherwise.

### GetImageDeletedOk

`func (o *DeleteContainerImageDetail) GetImageDeletedOk() (*bool, bool)`

GetImageDeletedOk returns a tuple with the ImageDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDeleted

`func (o *DeleteContainerImageDetail) SetImageDeleted(v bool)`

SetImageDeleted sets ImageDeleted field to given value.

### HasImageDeleted

`func (o *DeleteContainerImageDetail) HasImageDeleted() bool`

HasImageDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


