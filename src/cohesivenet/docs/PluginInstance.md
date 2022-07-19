# PluginInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**IpAddress** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**Firewall** | Pointer to [**[]PluginInstanceFirewallRule**](PluginInstanceFirewallRule.md) |  | [optional] 
**PortMaps** | Pointer to **map[string]int32** | Map of plugin ports -&gt; VNS3 ports | [optional] 
**Manager** | Pointer to [**PluginManagerConfig**](PluginManagerConfig.md) |  | [optional] 
**Data** | Pointer to [**Container**](Container.md) |  | [optional] 

## Methods

### NewPluginInstance

`func NewPluginInstance() *PluginInstance`

NewPluginInstance instantiates a new PluginInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginInstanceWithDefaults

`func NewPluginInstanceWithDefaults() *PluginInstance`

NewPluginInstanceWithDefaults instantiates a new PluginInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PluginInstance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PluginInstance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PluginInstance) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PluginInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PluginInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PluginInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PluginInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PluginInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PluginInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PluginInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PluginInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PluginInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *PluginInstance) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PluginInstance) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PluginInstance) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PluginInstance) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCommand

`func (o *PluginInstance) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PluginInstance) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PluginInstance) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *PluginInstance) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetImageId

`func (o *PluginInstance) GetImageId() int32`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *PluginInstance) GetImageIdOk() (*int32, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *PluginInstance) SetImageId(v int32)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *PluginInstance) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetUuid

`func (o *PluginInstance) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PluginInstance) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PluginInstance) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PluginInstance) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStatus

`func (o *PluginInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PluginInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PluginInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PluginInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIpAddress

`func (o *PluginInstance) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PluginInstance) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PluginInstance) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *PluginInstance) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PluginInstance) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PluginInstance) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PluginInstance) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PluginInstance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PluginInstance) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PluginInstance) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PluginInstance) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PluginInstance) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatusMessage

`func (o *PluginInstance) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *PluginInstance) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *PluginInstance) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *PluginInstance) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetEnvironment

`func (o *PluginInstance) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PluginInstance) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PluginInstance) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PluginInstance) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetImageName

`func (o *PluginInstance) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *PluginInstance) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *PluginInstance) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *PluginInstance) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetFirewall

`func (o *PluginInstance) GetFirewall() []PluginInstanceFirewallRule`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *PluginInstance) GetFirewallOk() (*[]PluginInstanceFirewallRule, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *PluginInstance) SetFirewall(v []PluginInstanceFirewallRule)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *PluginInstance) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetPortMaps

`func (o *PluginInstance) GetPortMaps() map[string]int32`

GetPortMaps returns the PortMaps field if non-nil, zero value otherwise.

### GetPortMapsOk

`func (o *PluginInstance) GetPortMapsOk() (*map[string]int32, bool)`

GetPortMapsOk returns a tuple with the PortMaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMaps

`func (o *PluginInstance) SetPortMaps(v map[string]int32)`

SetPortMaps sets PortMaps field to given value.

### HasPortMaps

`func (o *PluginInstance) HasPortMaps() bool`

HasPortMaps returns a boolean if a field has been set.

### GetManager

`func (o *PluginInstance) GetManager() PluginManagerConfig`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *PluginInstance) GetManagerOk() (*PluginManagerConfig, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *PluginInstance) SetManager(v PluginManagerConfig)`

SetManager sets Manager field to given value.

### HasManager

`func (o *PluginInstance) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetData

`func (o *PluginInstance) GetData() Container`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PluginInstance) GetDataOk() (*Container, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PluginInstance) SetData(v Container)`

SetData sets Data field to given value.

### HasData

`func (o *PluginInstance) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


