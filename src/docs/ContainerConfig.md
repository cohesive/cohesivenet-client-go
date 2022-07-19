# ContainerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entrypoint** | Pointer to **NullableString** |  | [optional] 
**Dns** | Pointer to **NullableString** |  | [optional] 
**OpenStdin** | Pointer to **bool** |  | [optional] 
**StdinOnce** | Pointer to **bool** |  | [optional] 
**AttachStderr** | Pointer to **bool** |  | [optional] 
**AttachStdout** | Pointer to **bool** |  | [optional] 
**AttachStdin** | Pointer to **bool** |  | [optional] 
**Env** | Pointer to **map[string]interface{}** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Tty** | Pointer to **bool** |  | [optional] 
**ExposedPorts** | Pointer to **map[string]interface{}** |  | [optional] 
**Memory** | Pointer to **int32** |  | [optional] 
**MemorySwap** | Pointer to **int32** |  | [optional] 
**VolumesFrom** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to **NullableString** |  | [optional] 
**Cmd** | Pointer to **NullableString** |  | [optional] 
**PortSpecs** | Pointer to **map[string]interface{}** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**WorkingDir** | Pointer to **string** |  | [optional] 
**CpuShares** | Pointer to **int32** |  | [optional] 
**NetworkDisabled** | Pointer to **bool** |  | [optional] 
**Domainname** | Pointer to **string** |  | [optional] 
**OnBuild** | Pointer to **NullableString** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 

## Methods

### NewContainerConfig

`func NewContainerConfig() *ContainerConfig`

NewContainerConfig instantiates a new ContainerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerConfigWithDefaults

`func NewContainerConfigWithDefaults() *ContainerConfig`

NewContainerConfigWithDefaults instantiates a new ContainerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntrypoint

`func (o *ContainerConfig) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *ContainerConfig) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *ContainerConfig) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *ContainerConfig) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### SetEntrypointNil

`func (o *ContainerConfig) SetEntrypointNil(b bool)`

 SetEntrypointNil sets the value for Entrypoint to be an explicit nil

### UnsetEntrypoint
`func (o *ContainerConfig) UnsetEntrypoint()`

UnsetEntrypoint ensures that no value is present for Entrypoint, not even an explicit nil
### GetDns

`func (o *ContainerConfig) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *ContainerConfig) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *ContainerConfig) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *ContainerConfig) HasDns() bool`

HasDns returns a boolean if a field has been set.

### SetDnsNil

`func (o *ContainerConfig) SetDnsNil(b bool)`

 SetDnsNil sets the value for Dns to be an explicit nil

### UnsetDns
`func (o *ContainerConfig) UnsetDns()`

UnsetDns ensures that no value is present for Dns, not even an explicit nil
### GetOpenStdin

`func (o *ContainerConfig) GetOpenStdin() bool`

GetOpenStdin returns the OpenStdin field if non-nil, zero value otherwise.

### GetOpenStdinOk

`func (o *ContainerConfig) GetOpenStdinOk() (*bool, bool)`

GetOpenStdinOk returns a tuple with the OpenStdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStdin

`func (o *ContainerConfig) SetOpenStdin(v bool)`

SetOpenStdin sets OpenStdin field to given value.

### HasOpenStdin

`func (o *ContainerConfig) HasOpenStdin() bool`

HasOpenStdin returns a boolean if a field has been set.

### GetStdinOnce

`func (o *ContainerConfig) GetStdinOnce() bool`

GetStdinOnce returns the StdinOnce field if non-nil, zero value otherwise.

### GetStdinOnceOk

`func (o *ContainerConfig) GetStdinOnceOk() (*bool, bool)`

GetStdinOnceOk returns a tuple with the StdinOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdinOnce

`func (o *ContainerConfig) SetStdinOnce(v bool)`

SetStdinOnce sets StdinOnce field to given value.

### HasStdinOnce

`func (o *ContainerConfig) HasStdinOnce() bool`

HasStdinOnce returns a boolean if a field has been set.

### GetAttachStderr

`func (o *ContainerConfig) GetAttachStderr() bool`

GetAttachStderr returns the AttachStderr field if non-nil, zero value otherwise.

### GetAttachStderrOk

`func (o *ContainerConfig) GetAttachStderrOk() (*bool, bool)`

GetAttachStderrOk returns a tuple with the AttachStderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStderr

`func (o *ContainerConfig) SetAttachStderr(v bool)`

SetAttachStderr sets AttachStderr field to given value.

### HasAttachStderr

`func (o *ContainerConfig) HasAttachStderr() bool`

HasAttachStderr returns a boolean if a field has been set.

### GetAttachStdout

`func (o *ContainerConfig) GetAttachStdout() bool`

GetAttachStdout returns the AttachStdout field if non-nil, zero value otherwise.

### GetAttachStdoutOk

`func (o *ContainerConfig) GetAttachStdoutOk() (*bool, bool)`

GetAttachStdoutOk returns a tuple with the AttachStdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStdout

`func (o *ContainerConfig) SetAttachStdout(v bool)`

SetAttachStdout sets AttachStdout field to given value.

### HasAttachStdout

`func (o *ContainerConfig) HasAttachStdout() bool`

HasAttachStdout returns a boolean if a field has been set.

### GetAttachStdin

`func (o *ContainerConfig) GetAttachStdin() bool`

GetAttachStdin returns the AttachStdin field if non-nil, zero value otherwise.

### GetAttachStdinOk

`func (o *ContainerConfig) GetAttachStdinOk() (*bool, bool)`

GetAttachStdinOk returns a tuple with the AttachStdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStdin

`func (o *ContainerConfig) SetAttachStdin(v bool)`

SetAttachStdin sets AttachStdin field to given value.

### HasAttachStdin

`func (o *ContainerConfig) HasAttachStdin() bool`

HasAttachStdin returns a boolean if a field has been set.

### GetEnv

`func (o *ContainerConfig) GetEnv() map[string]interface{}`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ContainerConfig) GetEnvOk() (*map[string]interface{}, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ContainerConfig) SetEnv(v map[string]interface{})`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ContainerConfig) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### SetEnvNil

`func (o *ContainerConfig) SetEnvNil(b bool)`

 SetEnvNil sets the value for Env to be an explicit nil

### UnsetEnv
`func (o *ContainerConfig) UnsetEnv()`

UnsetEnv ensures that no value is present for Env, not even an explicit nil
### GetUser

`func (o *ContainerConfig) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ContainerConfig) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ContainerConfig) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ContainerConfig) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTty

`func (o *ContainerConfig) GetTty() bool`

GetTty returns the Tty field if non-nil, zero value otherwise.

### GetTtyOk

`func (o *ContainerConfig) GetTtyOk() (*bool, bool)`

GetTtyOk returns a tuple with the Tty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTty

`func (o *ContainerConfig) SetTty(v bool)`

SetTty sets Tty field to given value.

### HasTty

`func (o *ContainerConfig) HasTty() bool`

HasTty returns a boolean if a field has been set.

### GetExposedPorts

`func (o *ContainerConfig) GetExposedPorts() map[string]interface{}`

GetExposedPorts returns the ExposedPorts field if non-nil, zero value otherwise.

### GetExposedPortsOk

`func (o *ContainerConfig) GetExposedPortsOk() (*map[string]interface{}, bool)`

GetExposedPortsOk returns a tuple with the ExposedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedPorts

`func (o *ContainerConfig) SetExposedPorts(v map[string]interface{})`

SetExposedPorts sets ExposedPorts field to given value.

### HasExposedPorts

`func (o *ContainerConfig) HasExposedPorts() bool`

HasExposedPorts returns a boolean if a field has been set.

### SetExposedPortsNil

`func (o *ContainerConfig) SetExposedPortsNil(b bool)`

 SetExposedPortsNil sets the value for ExposedPorts to be an explicit nil

### UnsetExposedPorts
`func (o *ContainerConfig) UnsetExposedPorts()`

UnsetExposedPorts ensures that no value is present for ExposedPorts, not even an explicit nil
### GetMemory

`func (o *ContainerConfig) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ContainerConfig) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ContainerConfig) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ContainerConfig) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMemorySwap

`func (o *ContainerConfig) GetMemorySwap() int32`

GetMemorySwap returns the MemorySwap field if non-nil, zero value otherwise.

### GetMemorySwapOk

`func (o *ContainerConfig) GetMemorySwapOk() (*int32, bool)`

GetMemorySwapOk returns a tuple with the MemorySwap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySwap

`func (o *ContainerConfig) SetMemorySwap(v int32)`

SetMemorySwap sets MemorySwap field to given value.

### HasMemorySwap

`func (o *ContainerConfig) HasMemorySwap() bool`

HasMemorySwap returns a boolean if a field has been set.

### GetVolumesFrom

`func (o *ContainerConfig) GetVolumesFrom() string`

GetVolumesFrom returns the VolumesFrom field if non-nil, zero value otherwise.

### GetVolumesFromOk

`func (o *ContainerConfig) GetVolumesFromOk() (*string, bool)`

GetVolumesFromOk returns a tuple with the VolumesFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumesFrom

`func (o *ContainerConfig) SetVolumesFrom(v string)`

SetVolumesFrom sets VolumesFrom field to given value.

### HasVolumesFrom

`func (o *ContainerConfig) HasVolumesFrom() bool`

HasVolumesFrom returns a boolean if a field has been set.

### GetVolumes

`func (o *ContainerConfig) GetVolumes() string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ContainerConfig) GetVolumesOk() (*string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ContainerConfig) SetVolumes(v string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ContainerConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *ContainerConfig) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *ContainerConfig) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil
### GetCmd

`func (o *ContainerConfig) GetCmd() string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *ContainerConfig) GetCmdOk() (*string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *ContainerConfig) SetCmd(v string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *ContainerConfig) HasCmd() bool`

HasCmd returns a boolean if a field has been set.

### SetCmdNil

`func (o *ContainerConfig) SetCmdNil(b bool)`

 SetCmdNil sets the value for Cmd to be an explicit nil

### UnsetCmd
`func (o *ContainerConfig) UnsetCmd()`

UnsetCmd ensures that no value is present for Cmd, not even an explicit nil
### GetPortSpecs

`func (o *ContainerConfig) GetPortSpecs() map[string]interface{}`

GetPortSpecs returns the PortSpecs field if non-nil, zero value otherwise.

### GetPortSpecsOk

`func (o *ContainerConfig) GetPortSpecsOk() (*map[string]interface{}, bool)`

GetPortSpecsOk returns a tuple with the PortSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpecs

`func (o *ContainerConfig) SetPortSpecs(v map[string]interface{})`

SetPortSpecs sets PortSpecs field to given value.

### HasPortSpecs

`func (o *ContainerConfig) HasPortSpecs() bool`

HasPortSpecs returns a boolean if a field has been set.

### SetPortSpecsNil

`func (o *ContainerConfig) SetPortSpecsNil(b bool)`

 SetPortSpecsNil sets the value for PortSpecs to be an explicit nil

### UnsetPortSpecs
`func (o *ContainerConfig) UnsetPortSpecs()`

UnsetPortSpecs ensures that no value is present for PortSpecs, not even an explicit nil
### GetImage

`func (o *ContainerConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ContainerConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ContainerConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *ContainerConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetWorkingDir

`func (o *ContainerConfig) GetWorkingDir() string`

GetWorkingDir returns the WorkingDir field if non-nil, zero value otherwise.

### GetWorkingDirOk

`func (o *ContainerConfig) GetWorkingDirOk() (*string, bool)`

GetWorkingDirOk returns a tuple with the WorkingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDir

`func (o *ContainerConfig) SetWorkingDir(v string)`

SetWorkingDir sets WorkingDir field to given value.

### HasWorkingDir

`func (o *ContainerConfig) HasWorkingDir() bool`

HasWorkingDir returns a boolean if a field has been set.

### GetCpuShares

`func (o *ContainerConfig) GetCpuShares() int32`

GetCpuShares returns the CpuShares field if non-nil, zero value otherwise.

### GetCpuSharesOk

`func (o *ContainerConfig) GetCpuSharesOk() (*int32, bool)`

GetCpuSharesOk returns a tuple with the CpuShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuShares

`func (o *ContainerConfig) SetCpuShares(v int32)`

SetCpuShares sets CpuShares field to given value.

### HasCpuShares

`func (o *ContainerConfig) HasCpuShares() bool`

HasCpuShares returns a boolean if a field has been set.

### GetNetworkDisabled

`func (o *ContainerConfig) GetNetworkDisabled() bool`

GetNetworkDisabled returns the NetworkDisabled field if non-nil, zero value otherwise.

### GetNetworkDisabledOk

`func (o *ContainerConfig) GetNetworkDisabledOk() (*bool, bool)`

GetNetworkDisabledOk returns a tuple with the NetworkDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDisabled

`func (o *ContainerConfig) SetNetworkDisabled(v bool)`

SetNetworkDisabled sets NetworkDisabled field to given value.

### HasNetworkDisabled

`func (o *ContainerConfig) HasNetworkDisabled() bool`

HasNetworkDisabled returns a boolean if a field has been set.

### GetDomainname

`func (o *ContainerConfig) GetDomainname() string`

GetDomainname returns the Domainname field if non-nil, zero value otherwise.

### GetDomainnameOk

`func (o *ContainerConfig) GetDomainnameOk() (*string, bool)`

GetDomainnameOk returns a tuple with the Domainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainname

`func (o *ContainerConfig) SetDomainname(v string)`

SetDomainname sets Domainname field to given value.

### HasDomainname

`func (o *ContainerConfig) HasDomainname() bool`

HasDomainname returns a boolean if a field has been set.

### GetOnBuild

`func (o *ContainerConfig) GetOnBuild() string`

GetOnBuild returns the OnBuild field if non-nil, zero value otherwise.

### GetOnBuildOk

`func (o *ContainerConfig) GetOnBuildOk() (*string, bool)`

GetOnBuildOk returns a tuple with the OnBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBuild

`func (o *ContainerConfig) SetOnBuild(v string)`

SetOnBuild sets OnBuild field to given value.

### HasOnBuild

`func (o *ContainerConfig) HasOnBuild() bool`

HasOnBuild returns a boolean if a field has been set.

### SetOnBuildNil

`func (o *ContainerConfig) SetOnBuildNil(b bool)`

 SetOnBuildNil sets the value for OnBuild to be an explicit nil

### UnsetOnBuild
`func (o *ContainerConfig) UnsetOnBuild()`

UnsetOnBuild ensures that no value is present for OnBuild, not even an explicit nil
### GetHostname

`func (o *ContainerConfig) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ContainerConfig) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ContainerConfig) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ContainerConfig) HasHostname() bool`

HasHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


