# ClientpackList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**OverlayIpaddress** | Pointer to **string** |  | [optional] 
**LinuxOnefile** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ConfSha1** | Pointer to **string** |  | [optional] 
**WindowsOnefile** | Pointer to **string** |  | [optional] 
**OvpnSha1** | Pointer to **string** |  | [optional] 
**TarballFile** | Pointer to **string** |  | [optional] 
**TarballSha1** | Pointer to **string** |  | [optional] 
**SequentialId** | Pointer to **int32** |  | [optional] 
**CheckedOut** | Pointer to **bool** |  | [optional] 
**ZipSha1** | Pointer to **string** |  | [optional] 
**ZipFile** | Pointer to **string** |  | [optional] 
**LastConnect** | Pointer to **string** |  | [optional] 
**LastDisconnect** | Pointer to **string** |  | [optional] 
**Compression** | Pointer to **string** | VPN compression status. on, off or error message. | [optional] 
**ManagerId** | Pointer to **string** |  | [optional] 
**Ipaddress** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Connected** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]string** | Key, value object of tags | [optional] 

## Methods

### NewClientpackList

`func NewClientpackList() *ClientpackList`

NewClientpackList instantiates a new ClientpackList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientpackListWithDefaults

`func NewClientpackListWithDefaults() *ClientpackList`

NewClientpackListWithDefaults instantiates a new ClientpackList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientpackList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientpackList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientpackList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientpackList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverlayIpaddress

`func (o *ClientpackList) GetOverlayIpaddress() string`

GetOverlayIpaddress returns the OverlayIpaddress field if non-nil, zero value otherwise.

### GetOverlayIpaddressOk

`func (o *ClientpackList) GetOverlayIpaddressOk() (*string, bool)`

GetOverlayIpaddressOk returns a tuple with the OverlayIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayIpaddress

`func (o *ClientpackList) SetOverlayIpaddress(v string)`

SetOverlayIpaddress sets OverlayIpaddress field to given value.

### HasOverlayIpaddress

`func (o *ClientpackList) HasOverlayIpaddress() bool`

HasOverlayIpaddress returns a boolean if a field has been set.

### GetLinuxOnefile

`func (o *ClientpackList) GetLinuxOnefile() string`

GetLinuxOnefile returns the LinuxOnefile field if non-nil, zero value otherwise.

### GetLinuxOnefileOk

`func (o *ClientpackList) GetLinuxOnefileOk() (*string, bool)`

GetLinuxOnefileOk returns a tuple with the LinuxOnefile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxOnefile

`func (o *ClientpackList) SetLinuxOnefile(v string)`

SetLinuxOnefile sets LinuxOnefile field to given value.

### HasLinuxOnefile

`func (o *ClientpackList) HasLinuxOnefile() bool`

HasLinuxOnefile returns a boolean if a field has been set.

### GetEnabled

`func (o *ClientpackList) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClientpackList) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClientpackList) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClientpackList) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetConfSha1

`func (o *ClientpackList) GetConfSha1() string`

GetConfSha1 returns the ConfSha1 field if non-nil, zero value otherwise.

### GetConfSha1Ok

`func (o *ClientpackList) GetConfSha1Ok() (*string, bool)`

GetConfSha1Ok returns a tuple with the ConfSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfSha1

`func (o *ClientpackList) SetConfSha1(v string)`

SetConfSha1 sets ConfSha1 field to given value.

### HasConfSha1

`func (o *ClientpackList) HasConfSha1() bool`

HasConfSha1 returns a boolean if a field has been set.

### GetWindowsOnefile

`func (o *ClientpackList) GetWindowsOnefile() string`

GetWindowsOnefile returns the WindowsOnefile field if non-nil, zero value otherwise.

### GetWindowsOnefileOk

`func (o *ClientpackList) GetWindowsOnefileOk() (*string, bool)`

GetWindowsOnefileOk returns a tuple with the WindowsOnefile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsOnefile

`func (o *ClientpackList) SetWindowsOnefile(v string)`

SetWindowsOnefile sets WindowsOnefile field to given value.

### HasWindowsOnefile

`func (o *ClientpackList) HasWindowsOnefile() bool`

HasWindowsOnefile returns a boolean if a field has been set.

### GetOvpnSha1

`func (o *ClientpackList) GetOvpnSha1() string`

GetOvpnSha1 returns the OvpnSha1 field if non-nil, zero value otherwise.

### GetOvpnSha1Ok

`func (o *ClientpackList) GetOvpnSha1Ok() (*string, bool)`

GetOvpnSha1Ok returns a tuple with the OvpnSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvpnSha1

`func (o *ClientpackList) SetOvpnSha1(v string)`

SetOvpnSha1 sets OvpnSha1 field to given value.

### HasOvpnSha1

`func (o *ClientpackList) HasOvpnSha1() bool`

HasOvpnSha1 returns a boolean if a field has been set.

### GetTarballFile

`func (o *ClientpackList) GetTarballFile() string`

GetTarballFile returns the TarballFile field if non-nil, zero value otherwise.

### GetTarballFileOk

`func (o *ClientpackList) GetTarballFileOk() (*string, bool)`

GetTarballFileOk returns a tuple with the TarballFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarballFile

`func (o *ClientpackList) SetTarballFile(v string)`

SetTarballFile sets TarballFile field to given value.

### HasTarballFile

`func (o *ClientpackList) HasTarballFile() bool`

HasTarballFile returns a boolean if a field has been set.

### GetTarballSha1

`func (o *ClientpackList) GetTarballSha1() string`

GetTarballSha1 returns the TarballSha1 field if non-nil, zero value otherwise.

### GetTarballSha1Ok

`func (o *ClientpackList) GetTarballSha1Ok() (*string, bool)`

GetTarballSha1Ok returns a tuple with the TarballSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarballSha1

`func (o *ClientpackList) SetTarballSha1(v string)`

SetTarballSha1 sets TarballSha1 field to given value.

### HasTarballSha1

`func (o *ClientpackList) HasTarballSha1() bool`

HasTarballSha1 returns a boolean if a field has been set.

### GetSequentialId

`func (o *ClientpackList) GetSequentialId() int32`

GetSequentialId returns the SequentialId field if non-nil, zero value otherwise.

### GetSequentialIdOk

`func (o *ClientpackList) GetSequentialIdOk() (*int32, bool)`

GetSequentialIdOk returns a tuple with the SequentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequentialId

`func (o *ClientpackList) SetSequentialId(v int32)`

SetSequentialId sets SequentialId field to given value.

### HasSequentialId

`func (o *ClientpackList) HasSequentialId() bool`

HasSequentialId returns a boolean if a field has been set.

### GetCheckedOut

`func (o *ClientpackList) GetCheckedOut() bool`

GetCheckedOut returns the CheckedOut field if non-nil, zero value otherwise.

### GetCheckedOutOk

`func (o *ClientpackList) GetCheckedOutOk() (*bool, bool)`

GetCheckedOutOk returns a tuple with the CheckedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedOut

`func (o *ClientpackList) SetCheckedOut(v bool)`

SetCheckedOut sets CheckedOut field to given value.

### HasCheckedOut

`func (o *ClientpackList) HasCheckedOut() bool`

HasCheckedOut returns a boolean if a field has been set.

### GetZipSha1

`func (o *ClientpackList) GetZipSha1() string`

GetZipSha1 returns the ZipSha1 field if non-nil, zero value otherwise.

### GetZipSha1Ok

`func (o *ClientpackList) GetZipSha1Ok() (*string, bool)`

GetZipSha1Ok returns a tuple with the ZipSha1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipSha1

`func (o *ClientpackList) SetZipSha1(v string)`

SetZipSha1 sets ZipSha1 field to given value.

### HasZipSha1

`func (o *ClientpackList) HasZipSha1() bool`

HasZipSha1 returns a boolean if a field has been set.

### GetZipFile

`func (o *ClientpackList) GetZipFile() string`

GetZipFile returns the ZipFile field if non-nil, zero value otherwise.

### GetZipFileOk

`func (o *ClientpackList) GetZipFileOk() (*string, bool)`

GetZipFileOk returns a tuple with the ZipFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipFile

`func (o *ClientpackList) SetZipFile(v string)`

SetZipFile sets ZipFile field to given value.

### HasZipFile

`func (o *ClientpackList) HasZipFile() bool`

HasZipFile returns a boolean if a field has been set.

### GetLastConnect

`func (o *ClientpackList) GetLastConnect() string`

GetLastConnect returns the LastConnect field if non-nil, zero value otherwise.

### GetLastConnectOk

`func (o *ClientpackList) GetLastConnectOk() (*string, bool)`

GetLastConnectOk returns a tuple with the LastConnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnect

`func (o *ClientpackList) SetLastConnect(v string)`

SetLastConnect sets LastConnect field to given value.

### HasLastConnect

`func (o *ClientpackList) HasLastConnect() bool`

HasLastConnect returns a boolean if a field has been set.

### GetLastDisconnect

`func (o *ClientpackList) GetLastDisconnect() string`

GetLastDisconnect returns the LastDisconnect field if non-nil, zero value otherwise.

### GetLastDisconnectOk

`func (o *ClientpackList) GetLastDisconnectOk() (*string, bool)`

GetLastDisconnectOk returns a tuple with the LastDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDisconnect

`func (o *ClientpackList) SetLastDisconnect(v string)`

SetLastDisconnect sets LastDisconnect field to given value.

### HasLastDisconnect

`func (o *ClientpackList) HasLastDisconnect() bool`

HasLastDisconnect returns a boolean if a field has been set.

### GetCompression

`func (o *ClientpackList) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *ClientpackList) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *ClientpackList) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *ClientpackList) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetManagerId

`func (o *ClientpackList) GetManagerId() string`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *ClientpackList) GetManagerIdOk() (*string, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *ClientpackList) SetManagerId(v string)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *ClientpackList) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### GetIpaddress

`func (o *ClientpackList) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *ClientpackList) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *ClientpackList) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *ClientpackList) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetStatus

`func (o *ClientpackList) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClientpackList) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClientpackList) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClientpackList) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetConnected

`func (o *ClientpackList) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ClientpackList) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ClientpackList) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ClientpackList) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetTags

`func (o *ClientpackList) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClientpackList) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClientpackList) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClientpackList) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


