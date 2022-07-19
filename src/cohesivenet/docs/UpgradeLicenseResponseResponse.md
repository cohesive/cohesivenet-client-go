# UpgradeLicenseResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Finalized** | Pointer to **bool** |  | [optional] 
**Uniq** | Pointer to **string** | new sha1 hash of license | [optional] 
**License** | Pointer to **string** | State of license, accepted, in-progress, failed | [optional] 
**NewClientpacks** | Pointer to **int32** |  | [optional] 
**NewManagers** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpgradeLicenseResponseResponse

`func NewUpgradeLicenseResponseResponse() *UpgradeLicenseResponseResponse`

NewUpgradeLicenseResponseResponse instantiates a new UpgradeLicenseResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeLicenseResponseResponseWithDefaults

`func NewUpgradeLicenseResponseResponseWithDefaults() *UpgradeLicenseResponseResponse`

NewUpgradeLicenseResponseResponseWithDefaults instantiates a new UpgradeLicenseResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinalized

`func (o *UpgradeLicenseResponseResponse) GetFinalized() bool`

GetFinalized returns the Finalized field if non-nil, zero value otherwise.

### GetFinalizedOk

`func (o *UpgradeLicenseResponseResponse) GetFinalizedOk() (*bool, bool)`

GetFinalizedOk returns a tuple with the Finalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalized

`func (o *UpgradeLicenseResponseResponse) SetFinalized(v bool)`

SetFinalized sets Finalized field to given value.

### HasFinalized

`func (o *UpgradeLicenseResponseResponse) HasFinalized() bool`

HasFinalized returns a boolean if a field has been set.

### GetUniq

`func (o *UpgradeLicenseResponseResponse) GetUniq() string`

GetUniq returns the Uniq field if non-nil, zero value otherwise.

### GetUniqOk

`func (o *UpgradeLicenseResponseResponse) GetUniqOk() (*string, bool)`

GetUniqOk returns a tuple with the Uniq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniq

`func (o *UpgradeLicenseResponseResponse) SetUniq(v string)`

SetUniq sets Uniq field to given value.

### HasUniq

`func (o *UpgradeLicenseResponseResponse) HasUniq() bool`

HasUniq returns a boolean if a field has been set.

### GetLicense

`func (o *UpgradeLicenseResponseResponse) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *UpgradeLicenseResponseResponse) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *UpgradeLicenseResponseResponse) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *UpgradeLicenseResponseResponse) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetNewClientpacks

`func (o *UpgradeLicenseResponseResponse) GetNewClientpacks() int32`

GetNewClientpacks returns the NewClientpacks field if non-nil, zero value otherwise.

### GetNewClientpacksOk

`func (o *UpgradeLicenseResponseResponse) GetNewClientpacksOk() (*int32, bool)`

GetNewClientpacksOk returns a tuple with the NewClientpacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientpacks

`func (o *UpgradeLicenseResponseResponse) SetNewClientpacks(v int32)`

SetNewClientpacks sets NewClientpacks field to given value.

### HasNewClientpacks

`func (o *UpgradeLicenseResponseResponse) HasNewClientpacks() bool`

HasNewClientpacks returns a boolean if a field has been set.

### GetNewManagers

`func (o *UpgradeLicenseResponseResponse) GetNewManagers() int32`

GetNewManagers returns the NewManagers field if non-nil, zero value otherwise.

### GetNewManagersOk

`func (o *UpgradeLicenseResponseResponse) GetNewManagersOk() (*int32, bool)`

GetNewManagersOk returns a tuple with the NewManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewManagers

`func (o *UpgradeLicenseResponseResponse) SetNewManagers(v int32)`

SetNewManagers sets NewManagers field to given value.

### HasNewManagers

`func (o *UpgradeLicenseResponseResponse) HasNewManagers() bool`

HasNewManagers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


