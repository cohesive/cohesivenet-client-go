# IdentityOIDCSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Provider** | **string** |  | 
**Identifier** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**RedirectHostname** | Pointer to **string** |  | [optional] 
**AuthorizationEndpoint** | Pointer to **string** |  | [optional] 
**TokenEndpoint** | Pointer to **string** |  | [optional] 
**UserinfoEndpoint** | Pointer to **string** |  | [optional] 
**JwksUri** | Pointer to **string** |  | [optional] 
**OtpUrl** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Keys** | Pointer to [**IdentityOIDCSettingsKeys**](IdentityOIDCSettingsKeys.md) |  | [optional] 

## Methods

### NewIdentityOIDCSettingsRequest

`func NewIdentityOIDCSettingsRequest(enabled bool, provider string, ) *IdentityOIDCSettingsRequest`

NewIdentityOIDCSettingsRequest instantiates a new IdentityOIDCSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityOIDCSettingsRequestWithDefaults

`func NewIdentityOIDCSettingsRequestWithDefaults() *IdentityOIDCSettingsRequest`

NewIdentityOIDCSettingsRequestWithDefaults instantiates a new IdentityOIDCSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *IdentityOIDCSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityOIDCSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityOIDCSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProvider

`func (o *IdentityOIDCSettingsRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IdentityOIDCSettingsRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IdentityOIDCSettingsRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetIdentifier

`func (o *IdentityOIDCSettingsRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *IdentityOIDCSettingsRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *IdentityOIDCSettingsRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *IdentityOIDCSettingsRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetSecret

`func (o *IdentityOIDCSettingsRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IdentityOIDCSettingsRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IdentityOIDCSettingsRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IdentityOIDCSettingsRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRedirectHostname

`func (o *IdentityOIDCSettingsRequest) GetRedirectHostname() string`

GetRedirectHostname returns the RedirectHostname field if non-nil, zero value otherwise.

### GetRedirectHostnameOk

`func (o *IdentityOIDCSettingsRequest) GetRedirectHostnameOk() (*string, bool)`

GetRedirectHostnameOk returns a tuple with the RedirectHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHostname

`func (o *IdentityOIDCSettingsRequest) SetRedirectHostname(v string)`

SetRedirectHostname sets RedirectHostname field to given value.

### HasRedirectHostname

`func (o *IdentityOIDCSettingsRequest) HasRedirectHostname() bool`

HasRedirectHostname returns a boolean if a field has been set.

### GetAuthorizationEndpoint

`func (o *IdentityOIDCSettingsRequest) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *IdentityOIDCSettingsRequest) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *IdentityOIDCSettingsRequest) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *IdentityOIDCSettingsRequest) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *IdentityOIDCSettingsRequest) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *IdentityOIDCSettingsRequest) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *IdentityOIDCSettingsRequest) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *IdentityOIDCSettingsRequest) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetUserinfoEndpoint

`func (o *IdentityOIDCSettingsRequest) GetUserinfoEndpoint() string`

GetUserinfoEndpoint returns the UserinfoEndpoint field if non-nil, zero value otherwise.

### GetUserinfoEndpointOk

`func (o *IdentityOIDCSettingsRequest) GetUserinfoEndpointOk() (*string, bool)`

GetUserinfoEndpointOk returns a tuple with the UserinfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserinfoEndpoint

`func (o *IdentityOIDCSettingsRequest) SetUserinfoEndpoint(v string)`

SetUserinfoEndpoint sets UserinfoEndpoint field to given value.

### HasUserinfoEndpoint

`func (o *IdentityOIDCSettingsRequest) HasUserinfoEndpoint() bool`

HasUserinfoEndpoint returns a boolean if a field has been set.

### GetJwksUri

`func (o *IdentityOIDCSettingsRequest) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *IdentityOIDCSettingsRequest) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *IdentityOIDCSettingsRequest) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *IdentityOIDCSettingsRequest) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetOtpUrl

`func (o *IdentityOIDCSettingsRequest) GetOtpUrl() string`

GetOtpUrl returns the OtpUrl field if non-nil, zero value otherwise.

### GetOtpUrlOk

`func (o *IdentityOIDCSettingsRequest) GetOtpUrlOk() (*string, bool)`

GetOtpUrlOk returns a tuple with the OtpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpUrl

`func (o *IdentityOIDCSettingsRequest) SetOtpUrl(v string)`

SetOtpUrl sets OtpUrl field to given value.

### HasOtpUrl

`func (o *IdentityOIDCSettingsRequest) HasOtpUrl() bool`

HasOtpUrl returns a boolean if a field has been set.

### GetIssuer

`func (o *IdentityOIDCSettingsRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *IdentityOIDCSettingsRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *IdentityOIDCSettingsRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *IdentityOIDCSettingsRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetKeys

`func (o *IdentityOIDCSettingsRequest) GetKeys() IdentityOIDCSettingsKeys`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *IdentityOIDCSettingsRequest) GetKeysOk() (*IdentityOIDCSettingsKeys, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *IdentityOIDCSettingsRequest) SetKeys(v IdentityOIDCSettingsKeys)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *IdentityOIDCSettingsRequest) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


