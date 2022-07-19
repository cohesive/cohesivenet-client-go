# PluginData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**ContainerConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPluginData

`func NewPluginData() *PluginData`

NewPluginData instantiates a new PluginData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginDataWithDefaults

`func NewPluginDataWithDefaults() *PluginData`

NewPluginDataWithDefaults instantiates a new PluginData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PluginData) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PluginData) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PluginData) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PluginData) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContainerConfig

`func (o *PluginData) GetContainerConfig() map[string]interface{}`

GetContainerConfig returns the ContainerConfig field if non-nil, zero value otherwise.

### GetContainerConfigOk

`func (o *PluginData) GetContainerConfigOk() (*map[string]interface{}, bool)`

GetContainerConfigOk returns a tuple with the ContainerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerConfig

`func (o *PluginData) SetContainerConfig(v map[string]interface{})`

SetContainerConfig sets ContainerConfig field to given value.

### HasContainerConfig

`func (o *PluginData) HasContainerConfig() bool`

HasContainerConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


