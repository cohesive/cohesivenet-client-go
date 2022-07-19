# CreateAPITokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **int32** | Number of seconds before expiration | [optional] [default to 3600]
**Name** | Pointer to **string** | Optional name of token | [optional] 
**TokenName** | Pointer to **string** | Optional name of token (deprecated) | [optional] 
**Refreshes** | Pointer to **bool** | Token lifetime refreshes when used | [optional] 

## Methods

### NewCreateAPITokenRequest

`func NewCreateAPITokenRequest() *CreateAPITokenRequest`

NewCreateAPITokenRequest instantiates a new CreateAPITokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAPITokenRequestWithDefaults

`func NewCreateAPITokenRequestWithDefaults() *CreateAPITokenRequest`

NewCreateAPITokenRequestWithDefaults instantiates a new CreateAPITokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *CreateAPITokenRequest) GetExpires() int32`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CreateAPITokenRequest) GetExpiresOk() (*int32, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CreateAPITokenRequest) SetExpires(v int32)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CreateAPITokenRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetName

`func (o *CreateAPITokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAPITokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAPITokenRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAPITokenRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTokenName

`func (o *CreateAPITokenRequest) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *CreateAPITokenRequest) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *CreateAPITokenRequest) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *CreateAPITokenRequest) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### GetRefreshes

`func (o *CreateAPITokenRequest) GetRefreshes() bool`

GetRefreshes returns the Refreshes field if non-nil, zero value otherwise.

### GetRefreshesOk

`func (o *CreateAPITokenRequest) GetRefreshesOk() (*bool, bool)`

GetRefreshesOk returns a tuple with the Refreshes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshes

`func (o *CreateAPITokenRequest) SetRefreshes(v bool)`

SetRefreshes sets Refreshes field to given value.

### HasRefreshes

`func (o *CreateAPITokenRequest) HasRefreshes() bool`

HasRefreshes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


