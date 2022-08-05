# SnapshotsListSnapshotsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha1Checksum** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CreatedAtI** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | Status of snapshot, pending, finished_ok, finished_failed or invalid | [optional] 
**Token** | Pointer to **string** | Token if snapshot task is still pending for polling | [optional] 

## Methods

### NewSnapshotsListSnapshotsValue

`func NewSnapshotsListSnapshotsValue() *SnapshotsListSnapshotsValue`

NewSnapshotsListSnapshotsValue instantiates a new SnapshotsListSnapshotsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotsListSnapshotsValueWithDefaults

`func NewSnapshotsListSnapshotsValueWithDefaults() *SnapshotsListSnapshotsValue`

NewSnapshotsListSnapshotsValueWithDefaults instantiates a new SnapshotsListSnapshotsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha1Checksum

`func (o *SnapshotsListSnapshotsValue) GetSha1Checksum() string`

GetSha1Checksum returns the Sha1Checksum field if non-nil, zero value otherwise.

### GetSha1ChecksumOk

`func (o *SnapshotsListSnapshotsValue) GetSha1ChecksumOk() (*string, bool)`

GetSha1ChecksumOk returns a tuple with the Sha1Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Checksum

`func (o *SnapshotsListSnapshotsValue) SetSha1Checksum(v string)`

SetSha1Checksum sets Sha1Checksum field to given value.

### HasSha1Checksum

`func (o *SnapshotsListSnapshotsValue) HasSha1Checksum() bool`

HasSha1Checksum returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SnapshotsListSnapshotsValue) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SnapshotsListSnapshotsValue) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SnapshotsListSnapshotsValue) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SnapshotsListSnapshotsValue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedAtI

`func (o *SnapshotsListSnapshotsValue) GetCreatedAtI() int32`

GetCreatedAtI returns the CreatedAtI field if non-nil, zero value otherwise.

### GetCreatedAtIOk

`func (o *SnapshotsListSnapshotsValue) GetCreatedAtIOk() (*int32, bool)`

GetCreatedAtIOk returns a tuple with the CreatedAtI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtI

`func (o *SnapshotsListSnapshotsValue) SetCreatedAtI(v int32)`

SetCreatedAtI sets CreatedAtI field to given value.

### HasCreatedAtI

`func (o *SnapshotsListSnapshotsValue) HasCreatedAtI() bool`

HasCreatedAtI returns a boolean if a field has been set.

### GetSize

`func (o *SnapshotsListSnapshotsValue) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SnapshotsListSnapshotsValue) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SnapshotsListSnapshotsValue) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SnapshotsListSnapshotsValue) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *SnapshotsListSnapshotsValue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotsListSnapshotsValue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotsListSnapshotsValue) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SnapshotsListSnapshotsValue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *SnapshotsListSnapshotsValue) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SnapshotsListSnapshotsValue) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SnapshotsListSnapshotsValue) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SnapshotsListSnapshotsValue) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


