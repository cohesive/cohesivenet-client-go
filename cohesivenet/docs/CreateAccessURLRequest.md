# CreateAccessURLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **int32** | Number of seconds before expiration | [optional] [default to 3600]
**Name** | Pointer to **string** | Optional name | [optional] 
**Description** | Pointer to **string** | Optional name (deprecated) | [optional] 
**Access** | Pointer to **string** | Type of access, remote support (rs) or clientpack (cp:100_1_64_0) | [optional] [default to "rs"]

## Methods

### NewCreateAccessURLRequest

`func NewCreateAccessURLRequest() *CreateAccessURLRequest`

NewCreateAccessURLRequest instantiates a new CreateAccessURLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessURLRequestWithDefaults

`func NewCreateAccessURLRequestWithDefaults() *CreateAccessURLRequest`

NewCreateAccessURLRequestWithDefaults instantiates a new CreateAccessURLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *CreateAccessURLRequest) GetExpires() int32`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CreateAccessURLRequest) GetExpiresOk() (*int32, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CreateAccessURLRequest) SetExpires(v int32)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CreateAccessURLRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetName

`func (o *CreateAccessURLRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccessURLRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccessURLRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAccessURLRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAccessURLRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAccessURLRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAccessURLRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAccessURLRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccess

`func (o *CreateAccessURLRequest) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CreateAccessURLRequest) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CreateAccessURLRequest) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *CreateAccessURLRequest) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


