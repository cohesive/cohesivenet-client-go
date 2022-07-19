# Plugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TagName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMsg** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**ImportId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**CatalogId** | Pointer to **string** |  | [optional] 
**Containers** | Pointer to [**[]PluginContainersInner**](PluginContainersInner.md) |  | [optional] 
**RunningContainers** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**PluginData**](PluginData.md) |  | [optional] 

## Methods

### NewPlugin

`func NewPlugin() *Plugin`

NewPlugin instantiates a new Plugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginWithDefaults

`func NewPluginWithDefaults() *Plugin`

NewPluginWithDefaults instantiates a new Plugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plugin) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plugin) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plugin) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Plugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *Plugin) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Plugin) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Plugin) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Plugin) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *Plugin) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *Plugin) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetName

`func (o *Plugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Plugin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTagName

`func (o *Plugin) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *Plugin) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *Plugin) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *Plugin) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetStatus

`func (o *Plugin) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Plugin) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Plugin) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Plugin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMsg

`func (o *Plugin) GetStatusMsg() string`

GetStatusMsg returns the StatusMsg field if non-nil, zero value otherwise.

### GetStatusMsgOk

`func (o *Plugin) GetStatusMsgOk() (*string, bool)`

GetStatusMsgOk returns a tuple with the StatusMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMsg

`func (o *Plugin) SetStatusMsg(v string)`

SetStatusMsg sets StatusMsg field to given value.

### HasStatusMsg

`func (o *Plugin) HasStatusMsg() bool`

HasStatusMsg returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Plugin) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Plugin) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Plugin) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Plugin) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetImportId

`func (o *Plugin) GetImportId() string`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *Plugin) GetImportIdOk() (*string, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *Plugin) SetImportId(v string)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *Plugin) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Plugin) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Plugin) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Plugin) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Plugin) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Plugin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Plugin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Plugin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Plugin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *Plugin) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Plugin) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Plugin) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Plugin) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCatalogId

`func (o *Plugin) GetCatalogId() string`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *Plugin) GetCatalogIdOk() (*string, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *Plugin) SetCatalogId(v string)`

SetCatalogId sets CatalogId field to given value.

### HasCatalogId

`func (o *Plugin) HasCatalogId() bool`

HasCatalogId returns a boolean if a field has been set.

### GetContainers

`func (o *Plugin) GetContainers() []PluginContainersInner`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *Plugin) GetContainersOk() (*[]PluginContainersInner, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *Plugin) SetContainers(v []PluginContainersInner)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *Plugin) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetRunningContainers

`func (o *Plugin) GetRunningContainers() int32`

GetRunningContainers returns the RunningContainers field if non-nil, zero value otherwise.

### GetRunningContainersOk

`func (o *Plugin) GetRunningContainersOk() (*int32, bool)`

GetRunningContainersOk returns a tuple with the RunningContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningContainers

`func (o *Plugin) SetRunningContainers(v int32)`

SetRunningContainers sets RunningContainers field to given value.

### HasRunningContainers

`func (o *Plugin) HasRunningContainers() bool`

HasRunningContainers returns a boolean if a field has been set.

### GetData

`func (o *Plugin) GetData() PluginData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Plugin) GetDataOk() (*PluginData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Plugin) SetData(v PluginData)`

SetData sets Data field to given value.

### HasData

`func (o *Plugin) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


