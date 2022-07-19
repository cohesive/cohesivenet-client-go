# UpdatePluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** | Plugin start command | [optional] 
**DocumentationUrl** | Pointer to **string** | URL to documentation for the plugin provided by plugin developer | [optional] 
**SupportUrl** | Pointer to **string** | URL to support site provided by plugin developer | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]string** | Key Value pairs associated with plugin | [optional] 

## Methods

### NewUpdatePluginRequest

`func NewUpdatePluginRequest() *UpdatePluginRequest`

NewUpdatePluginRequest instantiates a new UpdatePluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePluginRequestWithDefaults

`func NewUpdatePluginRequestWithDefaults() *UpdatePluginRequest`

NewUpdatePluginRequestWithDefaults instantiates a new UpdatePluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePluginRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePluginRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePluginRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePluginRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCommand

`func (o *UpdatePluginRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *UpdatePluginRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *UpdatePluginRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *UpdatePluginRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *UpdatePluginRequest) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *UpdatePluginRequest) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *UpdatePluginRequest) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *UpdatePluginRequest) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetSupportUrl

`func (o *UpdatePluginRequest) GetSupportUrl() string`

GetSupportUrl returns the SupportUrl field if non-nil, zero value otherwise.

### GetSupportUrlOk

`func (o *UpdatePluginRequest) GetSupportUrlOk() (*string, bool)`

GetSupportUrlOk returns a tuple with the SupportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportUrl

`func (o *UpdatePluginRequest) SetSupportUrl(v string)`

SetSupportUrl sets SupportUrl field to given value.

### HasSupportUrl

`func (o *UpdatePluginRequest) HasSupportUrl() bool`

HasSupportUrl returns a boolean if a field has been set.

### GetVersion

`func (o *UpdatePluginRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdatePluginRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdatePluginRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdatePluginRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTags

`func (o *UpdatePluginRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdatePluginRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdatePluginRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdatePluginRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


