# DeleteContainerDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**ContainerDeleted** | Pointer to **bool** |  | [optional] 

## Methods

### NewDeleteContainerDetail

`func NewDeleteContainerDetail() *DeleteContainerDetail`

NewDeleteContainerDetail instantiates a new DeleteContainerDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteContainerDetailWithDefaults

`func NewDeleteContainerDetailWithDefaults() *DeleteContainerDetail`

NewDeleteContainerDetailWithDefaults instantiates a new DeleteContainerDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DeleteContainerDetail) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeleteContainerDetail) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeleteContainerDetail) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeleteContainerDetail) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetContainerDeleted

`func (o *DeleteContainerDetail) GetContainerDeleted() bool`

GetContainerDeleted returns the ContainerDeleted field if non-nil, zero value otherwise.

### GetContainerDeletedOk

`func (o *DeleteContainerDetail) GetContainerDeletedOk() (*bool, bool)`

GetContainerDeletedOk returns a tuple with the ContainerDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDeleted

`func (o *DeleteContainerDetail) SetContainerDeleted(v bool)`

SetContainerDeleted sets ContainerDeleted field to given value.

### HasContainerDeleted

`func (o *DeleteContainerDetail) HasContainerDeleted() bool`

HasContainerDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


