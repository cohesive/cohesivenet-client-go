# IdentityRadiusSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to **string** | IP address or resolvable hostname | [optional] 
**AuthPort** | Pointer to **int32** | Authentication port | [optional] [default to 1812]
**AccountingPort** | Pointer to **int32** |  | [optional] [default to 1812]
**Pass** | Pointer to **string** | Shared password | [optional] 

## Methods

### NewIdentityRadiusSettings

`func NewIdentityRadiusSettings() *IdentityRadiusSettings`

NewIdentityRadiusSettings instantiates a new IdentityRadiusSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityRadiusSettingsWithDefaults

`func NewIdentityRadiusSettingsWithDefaults() *IdentityRadiusSettings`

NewIdentityRadiusSettingsWithDefaults instantiates a new IdentityRadiusSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *IdentityRadiusSettings) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *IdentityRadiusSettings) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *IdentityRadiusSettings) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *IdentityRadiusSettings) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetAuthPort

`func (o *IdentityRadiusSettings) GetAuthPort() int32`

GetAuthPort returns the AuthPort field if non-nil, zero value otherwise.

### GetAuthPortOk

`func (o *IdentityRadiusSettings) GetAuthPortOk() (*int32, bool)`

GetAuthPortOk returns a tuple with the AuthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPort

`func (o *IdentityRadiusSettings) SetAuthPort(v int32)`

SetAuthPort sets AuthPort field to given value.

### HasAuthPort

`func (o *IdentityRadiusSettings) HasAuthPort() bool`

HasAuthPort returns a boolean if a field has been set.

### GetAccountingPort

`func (o *IdentityRadiusSettings) GetAccountingPort() int32`

GetAccountingPort returns the AccountingPort field if non-nil, zero value otherwise.

### GetAccountingPortOk

`func (o *IdentityRadiusSettings) GetAccountingPortOk() (*int32, bool)`

GetAccountingPortOk returns a tuple with the AccountingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingPort

`func (o *IdentityRadiusSettings) SetAccountingPort(v int32)`

SetAccountingPort sets AccountingPort field to given value.

### HasAccountingPort

`func (o *IdentityRadiusSettings) HasAccountingPort() bool`

HasAccountingPort returns a boolean if a field has been set.

### GetPass

`func (o *IdentityRadiusSettings) GetPass() string`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *IdentityRadiusSettings) GetPassOk() (*string, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *IdentityRadiusSettings) SetPass(v string)`

SetPass sets Pass field to given value.

### HasPass

`func (o *IdentityRadiusSettings) HasPass() bool`

HasPass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


