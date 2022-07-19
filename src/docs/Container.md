# Container

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**RestartCount** | Pointer to **int32** |  | [optional] 
**Args** | Pointer to **[]string** |  | [optional] 
**AppArmorProfile** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**NullableContainerConfig**](ContainerConfig.md) |  | [optional] 
**State** | Pointer to [**ContainerState**](ContainerState.md) |  | [optional] 
**NetworkSettings** | Pointer to [**ContainerNetworkSettings**](ContainerNetworkSettings.md) |  | [optional] 
**ResolvConfPath** | Pointer to **string** |  | [optional] 
**HostnamePath** | Pointer to **string** |  | [optional] 
**HostsPath** | Pointer to **string** |  | [optional] 
**LogPath** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 
**ExecDriver** | Pointer to **string** |  | [optional] 
**MountLabel** | Pointer to **string** |  | [optional] 
**ProcessLabel** | Pointer to **string** |  | [optional] 
**ExecIDs** | Pointer to **[]string** |  | [optional] 
**Volumes** | Pointer to **map[string]interface{}** |  | [optional] 
**VolumesRW** | Pointer to **map[string]interface{}** |  | [optional] 
**HostConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewContainer

`func NewContainer() *Container`

NewContainer instantiates a new Container object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerWithDefaults

`func NewContainerWithDefaults() *Container`

NewContainerWithDefaults instantiates a new Container object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Container) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Container) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Container) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Container) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreated

`func (o *Container) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Container) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Container) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Container) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetPath

`func (o *Container) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Container) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Container) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Container) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRestartCount

`func (o *Container) GetRestartCount() int32`

GetRestartCount returns the RestartCount field if non-nil, zero value otherwise.

### GetRestartCountOk

`func (o *Container) GetRestartCountOk() (*int32, bool)`

GetRestartCountOk returns a tuple with the RestartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartCount

`func (o *Container) SetRestartCount(v int32)`

SetRestartCount sets RestartCount field to given value.

### HasRestartCount

`func (o *Container) HasRestartCount() bool`

HasRestartCount returns a boolean if a field has been set.

### GetArgs

`func (o *Container) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *Container) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *Container) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *Container) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetAppArmorProfile

`func (o *Container) GetAppArmorProfile() string`

GetAppArmorProfile returns the AppArmorProfile field if non-nil, zero value otherwise.

### GetAppArmorProfileOk

`func (o *Container) GetAppArmorProfileOk() (*string, bool)`

GetAppArmorProfileOk returns a tuple with the AppArmorProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppArmorProfile

`func (o *Container) SetAppArmorProfile(v string)`

SetAppArmorProfile sets AppArmorProfile field to given value.

### HasAppArmorProfile

`func (o *Container) HasAppArmorProfile() bool`

HasAppArmorProfile returns a boolean if a field has been set.

### GetImage

`func (o *Container) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Container) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Container) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Container) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetConfig

`func (o *Container) GetConfig() ContainerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Container) GetConfigOk() (*ContainerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Container) SetConfig(v ContainerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Container) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *Container) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *Container) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetState

`func (o *Container) GetState() ContainerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Container) GetStateOk() (*ContainerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Container) SetState(v ContainerState)`

SetState sets State field to given value.

### HasState

`func (o *Container) HasState() bool`

HasState returns a boolean if a field has been set.

### GetNetworkSettings

`func (o *Container) GetNetworkSettings() ContainerNetworkSettings`

GetNetworkSettings returns the NetworkSettings field if non-nil, zero value otherwise.

### GetNetworkSettingsOk

`func (o *Container) GetNetworkSettingsOk() (*ContainerNetworkSettings, bool)`

GetNetworkSettingsOk returns a tuple with the NetworkSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSettings

`func (o *Container) SetNetworkSettings(v ContainerNetworkSettings)`

SetNetworkSettings sets NetworkSettings field to given value.

### HasNetworkSettings

`func (o *Container) HasNetworkSettings() bool`

HasNetworkSettings returns a boolean if a field has been set.

### GetResolvConfPath

`func (o *Container) GetResolvConfPath() string`

GetResolvConfPath returns the ResolvConfPath field if non-nil, zero value otherwise.

### GetResolvConfPathOk

`func (o *Container) GetResolvConfPathOk() (*string, bool)`

GetResolvConfPathOk returns a tuple with the ResolvConfPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvConfPath

`func (o *Container) SetResolvConfPath(v string)`

SetResolvConfPath sets ResolvConfPath field to given value.

### HasResolvConfPath

`func (o *Container) HasResolvConfPath() bool`

HasResolvConfPath returns a boolean if a field has been set.

### GetHostnamePath

`func (o *Container) GetHostnamePath() string`

GetHostnamePath returns the HostnamePath field if non-nil, zero value otherwise.

### GetHostnamePathOk

`func (o *Container) GetHostnamePathOk() (*string, bool)`

GetHostnamePathOk returns a tuple with the HostnamePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnamePath

`func (o *Container) SetHostnamePath(v string)`

SetHostnamePath sets HostnamePath field to given value.

### HasHostnamePath

`func (o *Container) HasHostnamePath() bool`

HasHostnamePath returns a boolean if a field has been set.

### GetHostsPath

`func (o *Container) GetHostsPath() string`

GetHostsPath returns the HostsPath field if non-nil, zero value otherwise.

### GetHostsPathOk

`func (o *Container) GetHostsPathOk() (*string, bool)`

GetHostsPathOk returns a tuple with the HostsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsPath

`func (o *Container) SetHostsPath(v string)`

SetHostsPath sets HostsPath field to given value.

### HasHostsPath

`func (o *Container) HasHostsPath() bool`

HasHostsPath returns a boolean if a field has been set.

### GetLogPath

`func (o *Container) GetLogPath() string`

GetLogPath returns the LogPath field if non-nil, zero value otherwise.

### GetLogPathOk

`func (o *Container) GetLogPathOk() (*string, bool)`

GetLogPathOk returns a tuple with the LogPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogPath

`func (o *Container) SetLogPath(v string)`

SetLogPath sets LogPath field to given value.

### HasLogPath

`func (o *Container) HasLogPath() bool`

HasLogPath returns a boolean if a field has been set.

### GetName

`func (o *Container) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Container) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Container) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Container) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *Container) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Container) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Container) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Container) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDriver

`func (o *Container) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *Container) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *Container) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *Container) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetExecDriver

`func (o *Container) GetExecDriver() string`

GetExecDriver returns the ExecDriver field if non-nil, zero value otherwise.

### GetExecDriverOk

`func (o *Container) GetExecDriverOk() (*string, bool)`

GetExecDriverOk returns a tuple with the ExecDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecDriver

`func (o *Container) SetExecDriver(v string)`

SetExecDriver sets ExecDriver field to given value.

### HasExecDriver

`func (o *Container) HasExecDriver() bool`

HasExecDriver returns a boolean if a field has been set.

### GetMountLabel

`func (o *Container) GetMountLabel() string`

GetMountLabel returns the MountLabel field if non-nil, zero value otherwise.

### GetMountLabelOk

`func (o *Container) GetMountLabelOk() (*string, bool)`

GetMountLabelOk returns a tuple with the MountLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountLabel

`func (o *Container) SetMountLabel(v string)`

SetMountLabel sets MountLabel field to given value.

### HasMountLabel

`func (o *Container) HasMountLabel() bool`

HasMountLabel returns a boolean if a field has been set.

### GetProcessLabel

`func (o *Container) GetProcessLabel() string`

GetProcessLabel returns the ProcessLabel field if non-nil, zero value otherwise.

### GetProcessLabelOk

`func (o *Container) GetProcessLabelOk() (*string, bool)`

GetProcessLabelOk returns a tuple with the ProcessLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessLabel

`func (o *Container) SetProcessLabel(v string)`

SetProcessLabel sets ProcessLabel field to given value.

### HasProcessLabel

`func (o *Container) HasProcessLabel() bool`

HasProcessLabel returns a boolean if a field has been set.

### GetExecIDs

`func (o *Container) GetExecIDs() []string`

GetExecIDs returns the ExecIDs field if non-nil, zero value otherwise.

### GetExecIDsOk

`func (o *Container) GetExecIDsOk() (*[]string, bool)`

GetExecIDsOk returns a tuple with the ExecIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecIDs

`func (o *Container) SetExecIDs(v []string)`

SetExecIDs sets ExecIDs field to given value.

### HasExecIDs

`func (o *Container) HasExecIDs() bool`

HasExecIDs returns a boolean if a field has been set.

### SetExecIDsNil

`func (o *Container) SetExecIDsNil(b bool)`

 SetExecIDsNil sets the value for ExecIDs to be an explicit nil

### UnsetExecIDs
`func (o *Container) UnsetExecIDs()`

UnsetExecIDs ensures that no value is present for ExecIDs, not even an explicit nil
### GetVolumes

`func (o *Container) GetVolumes() map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *Container) GetVolumesOk() (*map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *Container) SetVolumes(v map[string]interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *Container) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *Container) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *Container) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil
### GetVolumesRW

`func (o *Container) GetVolumesRW() map[string]interface{}`

GetVolumesRW returns the VolumesRW field if non-nil, zero value otherwise.

### GetVolumesRWOk

`func (o *Container) GetVolumesRWOk() (*map[string]interface{}, bool)`

GetVolumesRWOk returns a tuple with the VolumesRW field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesRW

`func (o *Container) SetVolumesRW(v map[string]interface{})`

SetVolumesRW sets VolumesRW field to given value.

### HasVolumesRW

`func (o *Container) HasVolumesRW() bool`

HasVolumesRW returns a boolean if a field has been set.

### SetVolumesRWNil

`func (o *Container) SetVolumesRWNil(b bool)`

 SetVolumesRWNil sets the value for VolumesRW to be an explicit nil

### UnsetVolumesRW
`func (o *Container) UnsetVolumesRW()`

UnsetVolumesRW ensures that no value is present for VolumesRW, not even an explicit nil
### GetHostConfig

`func (o *Container) GetHostConfig() map[string]interface{}`

GetHostConfig returns the HostConfig field if non-nil, zero value otherwise.

### GetHostConfigOk

`func (o *Container) GetHostConfigOk() (*map[string]interface{}, bool)`

GetHostConfigOk returns a tuple with the HostConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostConfig

`func (o *Container) SetHostConfig(v map[string]interface{})`

SetHostConfig sets HostConfig field to given value.

### HasHostConfig

`func (o *Container) HasHostConfig() bool`

HasHostConfig returns a boolean if a field has been set.

### SetHostConfigNil

`func (o *Container) SetHostConfigNil(b bool)`

 SetHostConfigNil sets the value for HostConfig to be an explicit nil

### UnsetHostConfig
`func (o *Container) UnsetHostConfig()`

UnsetHostConfig ensures that no value is present for HostConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


