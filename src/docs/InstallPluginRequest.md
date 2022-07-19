# InstallPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new plugin | 
**ImageUrl** | **string** | Plugin Image URL to a tar.gz installable image. | 
**Description** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** | Plugin start command | [optional] 
**DocumentationUrl** | Pointer to **string** | URL to documentation for the plugin provided by plugin developer | [optional] 
**SupportUrl** | Pointer to **string** | URL to support site provided by plugin developer | [optional] 
**CatalogId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] [default to "1"]
**Tags** | Pointer to **map[string]string** | Key Value pairs associated with plugin | [optional] 
**Metadata** | Pointer to **map[string]string** | Optional additional data to associate with metadata | [optional] 

## Methods

### NewInstallPluginRequest

`func NewInstallPluginRequest(name string, imageUrl string, ) *InstallPluginRequest`

NewInstallPluginRequest instantiates a new InstallPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallPluginRequestWithDefaults

`func NewInstallPluginRequestWithDefaults() *InstallPluginRequest`

NewInstallPluginRequestWithDefaults instantiates a new InstallPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstallPluginRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstallPluginRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstallPluginRequest) SetName(v string)`

SetName sets Name field to given value.


### GetImageUrl

`func (o *InstallPluginRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *InstallPluginRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *InstallPluginRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetDescription

`func (o *InstallPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstallPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstallPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstallPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCommand

`func (o *InstallPluginRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *InstallPluginRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *InstallPluginRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *InstallPluginRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetDocumentationUrl

`func (o *InstallPluginRequest) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *InstallPluginRequest) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *InstallPluginRequest) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *InstallPluginRequest) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetSupportUrl

`func (o *InstallPluginRequest) GetSupportUrl() string`

GetSupportUrl returns the SupportUrl field if non-nil, zero value otherwise.

### GetSupportUrlOk

`func (o *InstallPluginRequest) GetSupportUrlOk() (*string, bool)`

GetSupportUrlOk returns a tuple with the SupportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportUrl

`func (o *InstallPluginRequest) SetSupportUrl(v string)`

SetSupportUrl sets SupportUrl field to given value.

### HasSupportUrl

`func (o *InstallPluginRequest) HasSupportUrl() bool`

HasSupportUrl returns a boolean if a field has been set.

### GetCatalogId

`func (o *InstallPluginRequest) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *InstallPluginRequest) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *InstallPluginRequest) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *InstallPluginRequest) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetVersion

`func (o *InstallPluginRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstallPluginRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstallPluginRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InstallPluginRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetTags

`func (o *InstallPluginRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstallPluginRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstallPluginRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstallPluginRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *InstallPluginRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstallPluginRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstallPluginRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InstallPluginRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


