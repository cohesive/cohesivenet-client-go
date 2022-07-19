# IdentityRadiusSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to **string** | IP address or resolvable hostname | [optional] 
**AuthPort** | Pointer to **int32** | Authentication port | [optional] [default to 1812]
**AccountingPort** | Pointer to **int32** |  | [optional] [default to 1812]
**Pass** | Pointer to **string** | Shared password | [optional] 
**Provider** | **string** |  | 
**Enabled** | **bool** |  | 

## Methods

### NewIdentityRadiusSettingsRequest

`func NewIdentityRadiusSettingsRequest(provider string, enabled bool, ) *IdentityRadiusSettingsRequest`

NewIdentityRadiusSettingsRequest instantiates a new IdentityRadiusSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityRadiusSettingsRequestWithDefaults

`func NewIdentityRadiusSettingsRequestWithDefaults() *IdentityRadiusSettingsRequest`

NewIdentityRadiusSettingsRequestWithDefaults instantiates a new IdentityRadiusSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *IdentityRadiusSettingsRequest) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *IdentityRadiusSettingsRequest) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *IdentityRadiusSettingsRequest) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *IdentityRadiusSettingsRequest) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetAuthPort

`func (o *IdentityRadiusSettingsRequest) GetAuthPort() int32`

GetAuthPort returns the AuthPort field if non-nil, zero value otherwise.

### GetAuthPortOk

`func (o *IdentityRadiusSettingsRequest) GetAuthPortOk() (*int32, bool)`

GetAuthPortOk returns a tuple with the AuthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPort

`func (o *IdentityRadiusSettingsRequest) SetAuthPort(v int32)`

SetAuthPort sets AuthPort field to given value.

### HasAuthPort

`func (o *IdentityRadiusSettingsRequest) HasAuthPort() bool`

HasAuthPort returns a boolean if a field has been set.

### GetAccountingPort

`func (o *IdentityRadiusSettingsRequest) GetAccountingPort() int32`

GetAccountingPort returns the AccountingPort field if non-nil, zero value otherwise.

### GetAccountingPortOk

`func (o *IdentityRadiusSettingsRequest) GetAccountingPortOk() (*int32, bool)`

GetAccountingPortOk returns a tuple with the AccountingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingPort

`func (o *IdentityRadiusSettingsRequest) SetAccountingPort(v int32)`

SetAccountingPort sets AccountingPort field to given value.

### HasAccountingPort

`func (o *IdentityRadiusSettingsRequest) HasAccountingPort() bool`

HasAccountingPort returns a boolean if a field has been set.

### GetPass

`func (o *IdentityRadiusSettingsRequest) GetPass() string`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *IdentityRadiusSettingsRequest) GetPassOk() (*string, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *IdentityRadiusSettingsRequest) SetPass(v string)`

SetPass sets Pass field to given value.

### HasPass

`func (o *IdentityRadiusSettingsRequest) HasPass() bool`

HasPass returns a boolean if a field has been set.

### GetProvider

`func (o *IdentityRadiusSettingsRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IdentityRadiusSettingsRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IdentityRadiusSettingsRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetEnabled

`func (o *IdentityRadiusSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityRadiusSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityRadiusSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


