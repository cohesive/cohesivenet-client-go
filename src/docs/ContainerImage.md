# ContainerImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMsg** | Pointer to **NullableString** |  | [optional] 
**ImportId** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TagName** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**ContainerConfig** | Pointer to [**NullableContainerConfig**](ContainerConfig.md) |  | [optional] 

## Methods

### NewContainerImage

`func NewContainerImage() *ContainerImage`

NewContainerImage instantiates a new ContainerImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageWithDefaults

`func NewContainerImageWithDefaults() *ContainerImage`

NewContainerImageWithDefaults instantiates a new ContainerImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerImage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerImage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerImage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContainerImage) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ContainerImage) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ContainerImage) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetImageName

`func (o *ContainerImage) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ContainerImage) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ContainerImage) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *ContainerImage) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetStatus

`func (o *ContainerImage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContainerImage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContainerImage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContainerImage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMsg

`func (o *ContainerImage) GetStatusMsg() string`

GetStatusMsg returns the StatusMsg field if non-nil, zero value otherwise.

### GetStatusMsgOk

`func (o *ContainerImage) GetStatusMsgOk() (*string, bool)`

GetStatusMsgOk returns a tuple with the StatusMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMsg

`func (o *ContainerImage) SetStatusMsg(v string)`

SetStatusMsg sets StatusMsg field to given value.

### HasStatusMsg

`func (o *ContainerImage) HasStatusMsg() bool`

HasStatusMsg returns a boolean if a field has been set.

### SetStatusMsgNil

`func (o *ContainerImage) SetStatusMsgNil(b bool)`

 SetStatusMsgNil sets the value for StatusMsg to be an explicit nil

### UnsetStatusMsg
`func (o *ContainerImage) UnsetStatusMsg()`

UnsetStatusMsg ensures that no value is present for StatusMsg, not even an explicit nil
### GetImportId

`func (o *ContainerImage) GetImportId() string`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *ContainerImage) GetImportIdOk() (*string, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *ContainerImage) SetImportId(v string)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *ContainerImage) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### GetCreated

`func (o *ContainerImage) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ContainerImage) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ContainerImage) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ContainerImage) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *ContainerImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerImage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerImage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTagName

`func (o *ContainerImage) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *ContainerImage) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *ContainerImage) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *ContainerImage) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetComment

`func (o *ContainerImage) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ContainerImage) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ContainerImage) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ContainerImage) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ContainerImage) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ContainerImage) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetContainerConfig

`func (o *ContainerImage) GetContainerConfig() ContainerConfig`

GetContainerConfig returns the ContainerConfig field if non-nil, zero value otherwise.

### GetContainerConfigOk

`func (o *ContainerImage) GetContainerConfigOk() (*ContainerConfig, bool)`

GetContainerConfigOk returns a tuple with the ContainerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerConfig

`func (o *ContainerImage) SetContainerConfig(v ContainerConfig)`

SetContainerConfig sets ContainerConfig field to given value.

### HasContainerConfig

`func (o *ContainerImage) HasContainerConfig() bool`

HasContainerConfig returns a boolean if a field has been set.

### SetContainerConfigNil

`func (o *ContainerImage) SetContainerConfigNil(b bool)`

 SetContainerConfigNil sets the value for ContainerConfig to be an explicit nil

### UnsetContainerConfig
`func (o *ContainerImage) UnsetContainerConfig()`

UnsetContainerConfig ensures that no value is present for ContainerConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


