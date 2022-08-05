# PluginManagerConfigConfigFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**PreviousVersions** | Pointer to [**[]PluginManagerConfigConfigFilePreviousVersionsInner**](PluginManagerConfigConfigFilePreviousVersionsInner.md) |  | [optional] 

## Methods

### NewPluginManagerConfigConfigFile

`func NewPluginManagerConfigConfigFile() *PluginManagerConfigConfigFile`

NewPluginManagerConfigConfigFile instantiates a new PluginManagerConfigConfigFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginManagerConfigConfigFileWithDefaults

`func NewPluginManagerConfigConfigFileWithDefaults() *PluginManagerConfigConfigFile`

NewPluginManagerConfigConfigFileWithDefaults instantiates a new PluginManagerConfigConfigFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PluginManagerConfigConfigFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginManagerConfigConfigFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginManagerConfigConfigFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginManagerConfigConfigFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *PluginManagerConfigConfigFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PluginManagerConfigConfigFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PluginManagerConfigConfigFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PluginManagerConfigConfigFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetDescription

`func (o *PluginManagerConfigConfigFile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PluginManagerConfigConfigFile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PluginManagerConfigConfigFile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PluginManagerConfigConfigFile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *PluginManagerConfigConfigFile) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PluginManagerConfigConfigFile) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PluginManagerConfigConfigFile) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PluginManagerConfigConfigFile) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PluginManagerConfigConfigFile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PluginManagerConfigConfigFile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PluginManagerConfigConfigFile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PluginManagerConfigConfigFile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPreviousVersions

`func (o *PluginManagerConfigConfigFile) GetPreviousVersions() []PluginManagerConfigConfigFilePreviousVersionsInner`

GetPreviousVersions returns the PreviousVersions field if non-nil, zero value otherwise.

### GetPreviousVersionsOk

`func (o *PluginManagerConfigConfigFile) GetPreviousVersionsOk() (*[]PluginManagerConfigConfigFilePreviousVersionsInner, bool)`

GetPreviousVersionsOk returns a tuple with the PreviousVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersions

`func (o *PluginManagerConfigConfigFile) SetPreviousVersions(v []PluginManagerConfigConfigFilePreviousVersionsInner)`

SetPreviousVersions sets PreviousVersions field to given value.

### HasPreviousVersions

`func (o *PluginManagerConfigConfigFile) HasPreviousVersions() bool`

HasPreviousVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


