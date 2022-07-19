# AccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CreatedIp** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Lifetime** | Pointer to **string** |  | [optional] 
**Refreshes** | Pointer to **bool** |  | [optional] 
**Expired** | Pointer to **bool** |  | [optional] 
**LastAccessedAt** | Pointer to **NullableString** |  | [optional] 
**LastAccessedIp** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAccessToken

`func NewAccessToken() *AccessToken`

NewAccessToken instantiates a new AccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenWithDefaults

`func NewAccessTokenWithDefaults() *AccessToken`

NewAccessTokenWithDefaults instantiates a new AccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessToken) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessToken) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessToken) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccessToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccessToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccessToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetToken

`func (o *AccessToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AccessToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AccessToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AccessToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetName

`func (o *AccessToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedIp

`func (o *AccessToken) GetCreatedIp() string`

GetCreatedIp returns the CreatedIp field if non-nil, zero value otherwise.

### GetCreatedIpOk

`func (o *AccessToken) GetCreatedIpOk() (*string, bool)`

GetCreatedIpOk returns a tuple with the CreatedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedIp

`func (o *AccessToken) SetCreatedIp(v string)`

SetCreatedIp sets CreatedIp field to given value.

### HasCreatedIp

`func (o *AccessToken) HasCreatedIp() bool`

HasCreatedIp returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AccessToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AccessToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AccessToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AccessToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetLifetime

`func (o *AccessToken) GetLifetime() string`

GetLifetime returns the Lifetime field if non-nil, zero value otherwise.

### GetLifetimeOk

`func (o *AccessToken) GetLifetimeOk() (*string, bool)`

GetLifetimeOk returns a tuple with the Lifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifetime

`func (o *AccessToken) SetLifetime(v string)`

SetLifetime sets Lifetime field to given value.

### HasLifetime

`func (o *AccessToken) HasLifetime() bool`

HasLifetime returns a boolean if a field has been set.

### GetRefreshes

`func (o *AccessToken) GetRefreshes() bool`

GetRefreshes returns the Refreshes field if non-nil, zero value otherwise.

### GetRefreshesOk

`func (o *AccessToken) GetRefreshesOk() (*bool, bool)`

GetRefreshesOk returns a tuple with the Refreshes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshes

`func (o *AccessToken) SetRefreshes(v bool)`

SetRefreshes sets Refreshes field to given value.

### HasRefreshes

`func (o *AccessToken) HasRefreshes() bool`

HasRefreshes returns a boolean if a field has been set.

### GetExpired

`func (o *AccessToken) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *AccessToken) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *AccessToken) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *AccessToken) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetLastAccessedAt

`func (o *AccessToken) GetLastAccessedAt() string`

GetLastAccessedAt returns the LastAccessedAt field if non-nil, zero value otherwise.

### GetLastAccessedAtOk

`func (o *AccessToken) GetLastAccessedAtOk() (*string, bool)`

GetLastAccessedAtOk returns a tuple with the LastAccessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedAt

`func (o *AccessToken) SetLastAccessedAt(v string)`

SetLastAccessedAt sets LastAccessedAt field to given value.

### HasLastAccessedAt

`func (o *AccessToken) HasLastAccessedAt() bool`

HasLastAccessedAt returns a boolean if a field has been set.

### SetLastAccessedAtNil

`func (o *AccessToken) SetLastAccessedAtNil(b bool)`

 SetLastAccessedAtNil sets the value for LastAccessedAt to be an explicit nil

### UnsetLastAccessedAt
`func (o *AccessToken) UnsetLastAccessedAt()`

UnsetLastAccessedAt ensures that no value is present for LastAccessedAt, not even an explicit nil
### GetLastAccessedIp

`func (o *AccessToken) GetLastAccessedIp() string`

GetLastAccessedIp returns the LastAccessedIp field if non-nil, zero value otherwise.

### GetLastAccessedIpOk

`func (o *AccessToken) GetLastAccessedIpOk() (*string, bool)`

GetLastAccessedIpOk returns a tuple with the LastAccessedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessedIp

`func (o *AccessToken) SetLastAccessedIp(v string)`

SetLastAccessedIp sets LastAccessedIp field to given value.

### HasLastAccessedIp

`func (o *AccessToken) HasLastAccessedIp() bool`

HasLastAccessedIp returns a boolean if a field has been set.

### SetLastAccessedIpNil

`func (o *AccessToken) SetLastAccessedIpNil(b bool)`

 SetLastAccessedIpNil sets the value for LastAccessedIp to be an explicit nil

### UnsetLastAccessedIp
`func (o *AccessToken) UnsetLastAccessedIp()`

UnsetLastAccessedIp ensures that no value is present for LastAccessedIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


