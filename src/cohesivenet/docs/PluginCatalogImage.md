# PluginCatalogImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Documentation** | Pointer to **string** |  | [optional] 
**Support** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Keyphrases** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ProviderCode** | Pointer to **string** |  | [optional] 
**Vns3Compatability** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Installed** | Pointer to **bool** |  | [optional] 
**InstalledImageId** | Pointer to **int32** |  | [optional] 

## Methods

### NewPluginCatalogImage

`func NewPluginCatalogImage() *PluginCatalogImage`

NewPluginCatalogImage instantiates a new PluginCatalogImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginCatalogImageWithDefaults

`func NewPluginCatalogImageWithDefaults() *PluginCatalogImage`

NewPluginCatalogImageWithDefaults instantiates a new PluginCatalogImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PluginCatalogImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginCatalogImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginCatalogImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginCatalogImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PluginCatalogImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PluginCatalogImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PluginCatalogImage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PluginCatalogImage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentation

`func (o *PluginCatalogImage) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *PluginCatalogImage) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *PluginCatalogImage) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *PluginCatalogImage) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetSupport

`func (o *PluginCatalogImage) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *PluginCatalogImage) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *PluginCatalogImage) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *PluginCatalogImage) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### GetImageUrl

`func (o *PluginCatalogImage) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PluginCatalogImage) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PluginCatalogImage) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PluginCatalogImage) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetCategories

`func (o *PluginCatalogImage) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PluginCatalogImage) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PluginCatalogImage) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PluginCatalogImage) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetTags

`func (o *PluginCatalogImage) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PluginCatalogImage) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PluginCatalogImage) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PluginCatalogImage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLogo

`func (o *PluginCatalogImage) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *PluginCatalogImage) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *PluginCatalogImage) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *PluginCatalogImage) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetKeyphrases

`func (o *PluginCatalogImage) GetKeyphrases() []string`

GetKeyphrases returns the Keyphrases field if non-nil, zero value otherwise.

### GetKeyphrasesOk

`func (o *PluginCatalogImage) GetKeyphrasesOk() (*[]string, bool)`

GetKeyphrasesOk returns a tuple with the Keyphrases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyphrases

`func (o *PluginCatalogImage) SetKeyphrases(v []string)`

SetKeyphrases sets Keyphrases field to given value.

### HasKeyphrases

`func (o *PluginCatalogImage) HasKeyphrases() bool`

HasKeyphrases returns a boolean if a field has been set.

### GetVersion

`func (o *PluginCatalogImage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PluginCatalogImage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PluginCatalogImage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PluginCatalogImage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetProviderCode

`func (o *PluginCatalogImage) GetProviderCode() string`

GetProviderCode returns the ProviderCode field if non-nil, zero value otherwise.

### GetProviderCodeOk

`func (o *PluginCatalogImage) GetProviderCodeOk() (*string, bool)`

GetProviderCodeOk returns a tuple with the ProviderCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCode

`func (o *PluginCatalogImage) SetProviderCode(v string)`

SetProviderCode sets ProviderCode field to given value.

### HasProviderCode

`func (o *PluginCatalogImage) HasProviderCode() bool`

HasProviderCode returns a boolean if a field has been set.

### GetVns3Compatability

`func (o *PluginCatalogImage) GetVns3Compatability() string`

GetVns3Compatability returns the Vns3Compatability field if non-nil, zero value otherwise.

### GetVns3CompatabilityOk

`func (o *PluginCatalogImage) GetVns3CompatabilityOk() (*string, bool)`

GetVns3CompatabilityOk returns a tuple with the Vns3Compatability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVns3Compatability

`func (o *PluginCatalogImage) SetVns3Compatability(v string)`

SetVns3Compatability sets Vns3Compatability field to given value.

### HasVns3Compatability

`func (o *PluginCatalogImage) HasVns3Compatability() bool`

HasVns3Compatability returns a boolean if a field has been set.

### GetId

`func (o *PluginCatalogImage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PluginCatalogImage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PluginCatalogImage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PluginCatalogImage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstalled

`func (o *PluginCatalogImage) GetInstalled() bool`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *PluginCatalogImage) GetInstalledOk() (*bool, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *PluginCatalogImage) SetInstalled(v bool)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *PluginCatalogImage) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetInstalledImageId

`func (o *PluginCatalogImage) GetInstalledImageId() int32`

GetInstalledImageId returns the InstalledImageId field if non-nil, zero value otherwise.

### GetInstalledImageIdOk

`func (o *PluginCatalogImage) GetInstalledImageIdOk() (*int32, bool)`

GetInstalledImageIdOk returns a tuple with the InstalledImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledImageId

`func (o *PluginCatalogImage) SetInstalledImageId(v int32)`

SetInstalledImageId sets InstalledImageId field to given value.

### HasInstalledImageId

`func (o *PluginCatalogImage) HasInstalledImageId() bool`

HasInstalledImageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


