# IdentityOIDCSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
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

### NewIdentityOIDCSettings

`func NewIdentityOIDCSettings() *IdentityOIDCSettings`

NewIdentityOIDCSettings instantiates a new IdentityOIDCSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityOIDCSettingsWithDefaults

`func NewIdentityOIDCSettingsWithDefaults() *IdentityOIDCSettings`

NewIdentityOIDCSettingsWithDefaults instantiates a new IdentityOIDCSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *IdentityOIDCSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityOIDCSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityOIDCSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdentityOIDCSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProvider

`func (o *IdentityOIDCSettings) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IdentityOIDCSettings) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IdentityOIDCSettings) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IdentityOIDCSettings) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetIdentifier

`func (o *IdentityOIDCSettings) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *IdentityOIDCSettings) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *IdentityOIDCSettings) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *IdentityOIDCSettings) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetSecret

`func (o *IdentityOIDCSettings) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IdentityOIDCSettings) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IdentityOIDCSettings) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IdentityOIDCSettings) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRedirectHostname

`func (o *IdentityOIDCSettings) GetRedirectHostname() string`

GetRedirectHostname returns the RedirectHostname field if non-nil, zero value otherwise.

### GetRedirectHostnameOk

`func (o *IdentityOIDCSettings) GetRedirectHostnameOk() (*string, bool)`

GetRedirectHostnameOk returns a tuple with the RedirectHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHostname

`func (o *IdentityOIDCSettings) SetRedirectHostname(v string)`

SetRedirectHostname sets RedirectHostname field to given value.

### HasRedirectHostname

`func (o *IdentityOIDCSettings) HasRedirectHostname() bool`

HasRedirectHostname returns a boolean if a field has been set.

### GetAuthorizationEndpoint

`func (o *IdentityOIDCSettings) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *IdentityOIDCSettings) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *IdentityOIDCSettings) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *IdentityOIDCSettings) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *IdentityOIDCSettings) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *IdentityOIDCSettings) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *IdentityOIDCSettings) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *IdentityOIDCSettings) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetUserinfoEndpoint

`func (o *IdentityOIDCSettings) GetUserinfoEndpoint() string`

GetUserinfoEndpoint returns the UserinfoEndpoint field if non-nil, zero value otherwise.

### GetUserinfoEndpointOk

`func (o *IdentityOIDCSettings) GetUserinfoEndpointOk() (*string, bool)`

GetUserinfoEndpointOk returns a tuple with the UserinfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserinfoEndpoint

`func (o *IdentityOIDCSettings) SetUserinfoEndpoint(v string)`

SetUserinfoEndpoint sets UserinfoEndpoint field to given value.

### HasUserinfoEndpoint

`func (o *IdentityOIDCSettings) HasUserinfoEndpoint() bool`

HasUserinfoEndpoint returns a boolean if a field has been set.

### GetJwksUri

`func (o *IdentityOIDCSettings) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *IdentityOIDCSettings) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *IdentityOIDCSettings) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *IdentityOIDCSettings) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetOtpUrl

`func (o *IdentityOIDCSettings) GetOtpUrl() string`

GetOtpUrl returns the OtpUrl field if non-nil, zero value otherwise.

### GetOtpUrlOk

`func (o *IdentityOIDCSettings) GetOtpUrlOk() (*string, bool)`

GetOtpUrlOk returns a tuple with the OtpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpUrl

`func (o *IdentityOIDCSettings) SetOtpUrl(v string)`

SetOtpUrl sets OtpUrl field to given value.

### HasOtpUrl

`func (o *IdentityOIDCSettings) HasOtpUrl() bool`

HasOtpUrl returns a boolean if a field has been set.

### GetIssuer

`func (o *IdentityOIDCSettings) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *IdentityOIDCSettings) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *IdentityOIDCSettings) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *IdentityOIDCSettings) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetKeys

`func (o *IdentityOIDCSettings) GetKeys() IdentityOIDCSettingsKeys`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *IdentityOIDCSettings) GetKeysOk() (*IdentityOIDCSettingsKeys, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *IdentityOIDCSettings) SetKeys(v IdentityOIDCSettingsKeys)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *IdentityOIDCSettings) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


