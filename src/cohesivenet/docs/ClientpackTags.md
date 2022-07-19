# ClientpackTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Clientpack name | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewClientpackTags

`func NewClientpackTags() *ClientpackTags`

NewClientpackTags instantiates a new ClientpackTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientpackTagsWithDefaults

`func NewClientpackTagsWithDefaults() *ClientpackTags`

NewClientpackTagsWithDefaults instantiates a new ClientpackTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientpackTags) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientpackTags) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientpackTags) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientpackTags) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *ClientpackTags) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClientpackTags) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClientpackTags) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClientpackTags) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


