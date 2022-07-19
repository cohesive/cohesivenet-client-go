# SystemStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] 
**TimestampI** | Pointer to **int32** |  | [optional] 
**Vns3Version** | Pointer to **string** |  | [optional] 
**KernelVersion** | Pointer to **string** |  | [optional] 
**Uptime** | Pointer to **int32** |  | [optional] 
**Loadavg** | Pointer to **[]string** |  | [optional] 
**Diskinfo** | Pointer to **[][]string** |  | [optional] 
**Meminfo** | Pointer to **[]string** |  | [optional] 
**Swapinfo** | Pointer to **[]string** |  | [optional] 
**ContainerSystem** | Pointer to [**ContainerSystem**](ContainerSystem.md) |  | [optional] 
**Data** | Pointer to [**SystemStatusData**](SystemStatusData.md) |  | [optional] 

## Methods

### NewSystemStatus

`func NewSystemStatus() *SystemStatus`

NewSystemStatus instantiates a new SystemStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemStatusWithDefaults

`func NewSystemStatusWithDefaults() *SystemStatus`

NewSystemStatusWithDefaults instantiates a new SystemStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *SystemStatus) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SystemStatus) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SystemStatus) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SystemStatus) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTimestampI

`func (o *SystemStatus) GetTimestampI() int32`

GetTimestampI returns the TimestampI field if non-nil, zero value otherwise.

### GetTimestampIOk

`func (o *SystemStatus) GetTimestampIOk() (*int32, bool)`

GetTimestampIOk returns a tuple with the TimestampI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampI

`func (o *SystemStatus) SetTimestampI(v int32)`

SetTimestampI sets TimestampI field to given value.

### HasTimestampI

`func (o *SystemStatus) HasTimestampI() bool`

HasTimestampI returns a boolean if a field has been set.

### GetVns3Version

`func (o *SystemStatus) GetVns3Version() string`

GetVns3Version returns the Vns3Version field if non-nil, zero value otherwise.

### GetVns3VersionOk

`func (o *SystemStatus) GetVns3VersionOk() (*string, bool)`

GetVns3VersionOk returns a tuple with the Vns3Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVns3Version

`func (o *SystemStatus) SetVns3Version(v string)`

SetVns3Version sets Vns3Version field to given value.

### HasVns3Version

`func (o *SystemStatus) HasVns3Version() bool`

HasVns3Version returns a boolean if a field has been set.

### GetKernelVersion

`func (o *SystemStatus) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *SystemStatus) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *SystemStatus) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.

### HasKernelVersion

`func (o *SystemStatus) HasKernelVersion() bool`

HasKernelVersion returns a boolean if a field has been set.

### GetUptime

`func (o *SystemStatus) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *SystemStatus) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *SystemStatus) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *SystemStatus) HasUptime() bool`

HasUptime returns a boolean if a field has been set.

### GetLoadavg

`func (o *SystemStatus) GetLoadavg() []string`

GetLoadavg returns the Loadavg field if non-nil, zero value otherwise.

### GetLoadavgOk

`func (o *SystemStatus) GetLoadavgOk() (*[]string, bool)`

GetLoadavgOk returns a tuple with the Loadavg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadavg

`func (o *SystemStatus) SetLoadavg(v []string)`

SetLoadavg sets Loadavg field to given value.

### HasLoadavg

`func (o *SystemStatus) HasLoadavg() bool`

HasLoadavg returns a boolean if a field has been set.

### GetDiskinfo

`func (o *SystemStatus) GetDiskinfo() [][]string`

GetDiskinfo returns the Diskinfo field if non-nil, zero value otherwise.

### GetDiskinfoOk

`func (o *SystemStatus) GetDiskinfoOk() (*[][]string, bool)`

GetDiskinfoOk returns a tuple with the Diskinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskinfo

`func (o *SystemStatus) SetDiskinfo(v [][]string)`

SetDiskinfo sets Diskinfo field to given value.

### HasDiskinfo

`func (o *SystemStatus) HasDiskinfo() bool`

HasDiskinfo returns a boolean if a field has been set.

### GetMeminfo

`func (o *SystemStatus) GetMeminfo() []string`

GetMeminfo returns the Meminfo field if non-nil, zero value otherwise.

### GetMeminfoOk

`func (o *SystemStatus) GetMeminfoOk() (*[]string, bool)`

GetMeminfoOk returns a tuple with the Meminfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeminfo

`func (o *SystemStatus) SetMeminfo(v []string)`

SetMeminfo sets Meminfo field to given value.

### HasMeminfo

`func (o *SystemStatus) HasMeminfo() bool`

HasMeminfo returns a boolean if a field has been set.

### GetSwapinfo

`func (o *SystemStatus) GetSwapinfo() []string`

GetSwapinfo returns the Swapinfo field if non-nil, zero value otherwise.

### GetSwapinfoOk

`func (o *SystemStatus) GetSwapinfoOk() (*[]string, bool)`

GetSwapinfoOk returns a tuple with the Swapinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapinfo

`func (o *SystemStatus) SetSwapinfo(v []string)`

SetSwapinfo sets Swapinfo field to given value.

### HasSwapinfo

`func (o *SystemStatus) HasSwapinfo() bool`

HasSwapinfo returns a boolean if a field has been set.

### GetContainerSystem

`func (o *SystemStatus) GetContainerSystem() ContainerSystem`

GetContainerSystem returns the ContainerSystem field if non-nil, zero value otherwise.

### GetContainerSystemOk

`func (o *SystemStatus) GetContainerSystemOk() (*ContainerSystem, bool)`

GetContainerSystemOk returns a tuple with the ContainerSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSystem

`func (o *SystemStatus) SetContainerSystem(v ContainerSystem)`

SetContainerSystem sets ContainerSystem field to given value.

### HasContainerSystem

`func (o *SystemStatus) HasContainerSystem() bool`

HasContainerSystem returns a boolean if a field has been set.

### GetData

`func (o *SystemStatus) GetData() SystemStatusData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SystemStatus) GetDataOk() (*SystemStatusData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SystemStatus) SetData(v SystemStatusData)`

SetData sets Data field to given value.

### HasData

`func (o *SystemStatus) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


