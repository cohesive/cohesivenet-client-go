# KeysetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InProgress** | Pointer to **bool** |  | [optional] 
**Running** | Pointer to **int32** |  | [optional] 
**KeysetPresent** | Pointer to **bool** |  | [optional] 
**Checksum** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedAtI** | Pointer to **int32** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**StartedAtI** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewKeysetStatus

`func NewKeysetStatus() *KeysetStatus`

NewKeysetStatus instantiates a new KeysetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeysetStatusWithDefaults

`func NewKeysetStatusWithDefaults() *KeysetStatus`

NewKeysetStatusWithDefaults instantiates a new KeysetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInProgress

`func (o *KeysetStatus) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *KeysetStatus) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *KeysetStatus) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *KeysetStatus) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetRunning

`func (o *KeysetStatus) GetRunning() int32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *KeysetStatus) GetRunningOk() (*int32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *KeysetStatus) SetRunning(v int32)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *KeysetStatus) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetKeysetPresent

`func (o *KeysetStatus) GetKeysetPresent() bool`

GetKeysetPresent returns the KeysetPresent field if non-nil, zero value otherwise.

### GetKeysetPresentOk

`func (o *KeysetStatus) GetKeysetPresentOk() (*bool, bool)`

GetKeysetPresentOk returns a tuple with the KeysetPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeysetPresent

`func (o *KeysetStatus) SetKeysetPresent(v bool)`

SetKeysetPresent sets KeysetPresent field to given value.

### HasKeysetPresent

`func (o *KeysetStatus) HasKeysetPresent() bool`

HasKeysetPresent returns a boolean if a field has been set.

### GetChecksum

`func (o *KeysetStatus) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *KeysetStatus) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *KeysetStatus) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *KeysetStatus) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KeysetStatus) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KeysetStatus) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KeysetStatus) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KeysetStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedAtI

`func (o *KeysetStatus) GetCreatedAtI() int32`

GetCreatedAtI returns the CreatedAtI field if non-nil, zero value otherwise.

### GetCreatedAtIOk

`func (o *KeysetStatus) GetCreatedAtIOk() (*int32, bool)`

GetCreatedAtIOk returns a tuple with the CreatedAtI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtI

`func (o *KeysetStatus) SetCreatedAtI(v int32)`

SetCreatedAtI sets CreatedAtI field to given value.

### HasCreatedAtI

`func (o *KeysetStatus) HasCreatedAtI() bool`

HasCreatedAtI returns a boolean if a field has been set.

### GetStartedAt

`func (o *KeysetStatus) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *KeysetStatus) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *KeysetStatus) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *KeysetStatus) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStartedAtI

`func (o *KeysetStatus) GetStartedAtI() int32`

GetStartedAtI returns the StartedAtI field if non-nil, zero value otherwise.

### GetStartedAtIOk

`func (o *KeysetStatus) GetStartedAtIOk() (*int32, bool)`

GetStartedAtIOk returns a tuple with the StartedAtI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtI

`func (o *KeysetStatus) SetStartedAtI(v int32)`

SetStartedAtI sets StartedAtI field to given value.

### HasStartedAtI

`func (o *KeysetStatus) HasStartedAtI() bool`

HasStartedAtI returns a boolean if a field has been set.

### GetUuid

`func (o *KeysetStatus) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *KeysetStatus) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *KeysetStatus) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *KeysetStatus) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


