# PutPluginInstanceUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** | password for ssh login | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**RequirePublicKey** | Pointer to **bool** |  | [optional] 

## Methods

### NewPutPluginInstanceUserRequest

`func NewPutPluginInstanceUserRequest() *PutPluginInstanceUserRequest`

NewPutPluginInstanceUserRequest instantiates a new PutPluginInstanceUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutPluginInstanceUserRequestWithDefaults

`func NewPutPluginInstanceUserRequestWithDefaults() *PutPluginInstanceUserRequest`

NewPutPluginInstanceUserRequestWithDefaults instantiates a new PutPluginInstanceUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *PutPluginInstanceUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PutPluginInstanceUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PutPluginInstanceUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PutPluginInstanceUserRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PutPluginInstanceUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PutPluginInstanceUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PutPluginInstanceUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PutPluginInstanceUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPublicKey

`func (o *PutPluginInstanceUserRequest) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *PutPluginInstanceUserRequest) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *PutPluginInstanceUserRequest) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *PutPluginInstanceUserRequest) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRequirePublicKey

`func (o *PutPluginInstanceUserRequest) GetRequirePublicKey() bool`

GetRequirePublicKey returns the RequirePublicKey field if non-nil, zero value otherwise.

### GetRequirePublicKeyOk

`func (o *PutPluginInstanceUserRequest) GetRequirePublicKeyOk() (*bool, bool)`

GetRequirePublicKeyOk returns a tuple with the RequirePublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePublicKey

`func (o *PutPluginInstanceUserRequest) SetRequirePublicKey(v bool)`

SetRequirePublicKey sets RequirePublicKey field to given value.

### HasRequirePublicKey

`func (o *PutPluginInstanceUserRequest) HasRequirePublicKey() bool`

HasRequirePublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


