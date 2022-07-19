# UpdateRemoteSupportConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | True if remote support should be enabled | [optional] 
**RevokeCredential** | Pointer to **bool** | True if remote support credential should be revoked | [optional] 

## Methods

### NewUpdateRemoteSupportConfigRequest

`func NewUpdateRemoteSupportConfigRequest() *UpdateRemoteSupportConfigRequest`

NewUpdateRemoteSupportConfigRequest instantiates a new UpdateRemoteSupportConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRemoteSupportConfigRequestWithDefaults

`func NewUpdateRemoteSupportConfigRequestWithDefaults() *UpdateRemoteSupportConfigRequest`

NewUpdateRemoteSupportConfigRequestWithDefaults instantiates a new UpdateRemoteSupportConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UpdateRemoteSupportConfigRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateRemoteSupportConfigRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateRemoteSupportConfigRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateRemoteSupportConfigRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRevokeCredential

`func (o *UpdateRemoteSupportConfigRequest) GetRevokeCredential() bool`

GetRevokeCredential returns the RevokeCredential field if non-nil, zero value otherwise.

### GetRevokeCredentialOk

`func (o *UpdateRemoteSupportConfigRequest) GetRevokeCredentialOk() (*bool, bool)`

GetRevokeCredentialOk returns a tuple with the RevokeCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeCredential

`func (o *UpdateRemoteSupportConfigRequest) SetRevokeCredential(v bool)`

SetRevokeCredential sets RevokeCredential field to given value.

### HasRevokeCredential

`func (o *UpdateRemoteSupportConfigRequest) HasRevokeCredential() bool`

HasRevokeCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


