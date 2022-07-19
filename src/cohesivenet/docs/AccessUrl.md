# AccessUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedIp** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Lifetime** | Pointer to **string** |  | [optional] 
**Expired** | Pointer to **bool** |  | [optional] 
**LastAccessedAt** | Pointer to **NullableString** |  | [optional] 
**LastAccessedIp** | Pointer to **NullableString** |  | [optional] 
**Access** | Pointer to **string** | Type of access, remote support (rs) or clientpack (cp:100_1_64_0) | [optional] [default to "rs"]

## Methods

### NewAccessUrl

`func NewAccessUrl() *AccessUrl`

NewAccessUrl instantiates a new AccessUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessUrlWithDefaults

`func NewAccessUrlWithDefaults() *AccessUrl`

NewAccessUrlWithDefaults instantiates a new AccessUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessUrl) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessUrl) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessUrl) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccessUrl) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *AccessUrl) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AccessUrl) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AccessUrl) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AccessUrl) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccessUrl) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessUrl) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessUrl) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccessUrl) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedIp

`func (o *AccessUrl) GetCreatedIp() string`

GetCreatedIp returns the CreatedIp field if non-nil, zero value otherwise.

### GetCreatedIpOk

`func (o *AccessUrl) GetCreatedIpOk() (*string, bool)`

GetCreatedIpOk returns a tuple with the CreatedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedIp

`func (o *AccessUrl) SetCreatedIp(v string)`

SetCreatedIp sets CreatedIp field to given value.

### HasCreatedIp

`func (o *AccessUrl) HasCreatedIp() bool`

HasCreatedIp returns a boolean if a field has been set.

### GetName

`func (o *AccessUrl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessUrl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessUrl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessUrl) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AccessUrl) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AccessUrl) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AccessUrl) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AccessUrl) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetLifetime

`func (o *AccessUrl) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *AccessUrl) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *AccessUrl) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *AccessUrl) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetExpired

`func (o *AccessUrl) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *AccessUrl) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *AccessUrl) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *AccessUrl) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetLastAccessedAt

`func (o *AccessUrl) GetLastAccessedAt() string`

GetLastAccessedAt returns the LastAccessedAt field if non-nil, zero value otherwise.

### GetLastAccessedAtOk

`func (o *AccessUrl) GetLastAccessedAtOk() (*string, bool)`

GetLastAccessedAtOk returns a tuple with the LastAccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedAt

`func (o *AccessUrl) SetLastAccessedAt(v string)`

SetLastAccessedAt sets LastAccessedAt field to given value.

### HasLastAccessedAt

`func (o *AccessUrl) HasLastAccessedAt() bool`

HasLastAccessedAt returns a boolean if a field has been set.

### SetLastAccessedAtNil

`func (o *AccessUrl) SetLastAccessedAtNil(b bool)`

 SetLastAccessedAtNil sets the value for LastAccessedAt to be an explicit nil

### UnsetLastAccessedAt
`func (o *AccessUrl) UnsetLastAccessedAt()`

UnsetLastAccessedAt ensures that no value is present for LastAccessedAt, not even an explicit nil
### GetLastAccessedIp

`func (o *AccessUrl) GetLastAccessedIp() string`

GetLastAccessedIp returns the LastAccessedIp field if non-nil, zero value otherwise.

### GetLastAccessedIpOk

`func (o *AccessUrl) GetLastAccessedIpOk() (*string, bool)`

GetLastAccessedIpOk returns a tuple with the LastAccessedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedIp

`func (o *AccessUrl) SetLastAccessedIp(v string)`

SetLastAccessedIp sets LastAccessedIp field to given value.

### HasLastAccessedIp

`func (o *AccessUrl) HasLastAccessedIp() bool`

HasLastAccessedIp returns a boolean if a field has been set.

### SetLastAccessedIpNil

`func (o *AccessUrl) SetLastAccessedIpNil(b bool)`

 SetLastAccessedIpNil sets the value for LastAccessedIp to be an explicit nil

### UnsetLastAccessedIp
`func (o *AccessUrl) UnsetLastAccessedIp()`

UnsetLastAccessedIp ensures that no value is present for LastAccessedIp, not even an explicit nil
### GetAccess

`func (o *AccessUrl) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AccessUrl) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AccessUrl) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AccessUrl) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


