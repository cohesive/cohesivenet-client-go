# PostTestLdapSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Currently only ldap is supported for testing | 
**Host** | Pointer to **string** | IP address or resolvable hostname of LDAP server | [optional] 
**Port** | Pointer to **int32** | Port for LDAP | [optional] [default to 389]
**Encrypt** | Pointer to **bool** | Use SSL | [optional] [default to false]
**EncryptLdaps** | Pointer to **bool** | Use LDAPS or start TLS (default)? | [optional] [default to true]
**EncryptAuth** | Pointer to **bool** | Use certificates to authenticate via encrypted connection | [optional] [default to false]
**EncryptVerifyCa** | Pointer to **bool** | Verify certicate using authority | [optional] [default to false]
**Binddn** | Pointer to **string** | Bind Username | [optional] 
**Bindpw** | Pointer to **string** | Bind Password | [optional] 
**AuthCert** | Pointer to **string** | Authentication certificate text content to use | [optional] 
**AuthCertCurrent** | Pointer to **bool** | Test with current uploaded authentication certificate? | [optional] [default to false]
**AuthKey** | Pointer to **string** | Authentication key text content to use | [optional] 
**AuthKeyCurrent** | Pointer to **bool** | Test with current uploaded authentication key? | [optional] [default to false]
**CaCert** | Pointer to **string** | CA certificate text content to use | [optional] 
**CaCertCurrent** | Pointer to **bool** | Test with current uploaded CA certificate? | [optional] [default to false]

## Methods

### NewPostTestLdapSettingsRequest

`func NewPostTestLdapSettingsRequest(provider string, ) *PostTestLdapSettingsRequest`

NewPostTestLdapSettingsRequest instantiates a new PostTestLdapSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTestLdapSettingsRequestWithDefaults

`func NewPostTestLdapSettingsRequestWithDefaults() *PostTestLdapSettingsRequest`

NewPostTestLdapSettingsRequestWithDefaults instantiates a new PostTestLdapSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *PostTestLdapSettingsRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PostTestLdapSettingsRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PostTestLdapSettingsRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetHost

`func (o *PostTestLdapSettingsRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PostTestLdapSettingsRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PostTestLdapSettingsRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PostTestLdapSettingsRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *PostTestLdapSettingsRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostTestLdapSettingsRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostTestLdapSettingsRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PostTestLdapSettingsRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEncrypt

`func (o *PostTestLdapSettingsRequest) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *PostTestLdapSettingsRequest) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *PostTestLdapSettingsRequest) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *PostTestLdapSettingsRequest) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptLdaps

`func (o *PostTestLdapSettingsRequest) GetEncryptLdaps() bool`

GetEncryptLdaps returns the EncryptLdaps field if non-nil, zero value otherwise.

### GetEncryptLdapsOk

`func (o *PostTestLdapSettingsRequest) GetEncryptLdapsOk() (*bool, bool)`

GetEncryptLdapsOk returns a tuple with the EncryptLdaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptLdaps

`func (o *PostTestLdapSettingsRequest) SetEncryptLdaps(v bool)`

SetEncryptLdaps sets EncryptLdaps field to given value.

### HasEncryptLdaps

`func (o *PostTestLdapSettingsRequest) HasEncryptLdaps() bool`

HasEncryptLdaps returns a boolean if a field has been set.

### GetEncryptAuth

`func (o *PostTestLdapSettingsRequest) GetEncryptAuth() bool`

GetEncryptAuth returns the EncryptAuth field if non-nil, zero value otherwise.

### GetEncryptAuthOk

`func (o *PostTestLdapSettingsRequest) GetEncryptAuthOk() (*bool, bool)`

GetEncryptAuthOk returns a tuple with the EncryptAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuth

`func (o *PostTestLdapSettingsRequest) SetEncryptAuth(v bool)`

SetEncryptAuth sets EncryptAuth field to given value.

### HasEncryptAuth

`func (o *PostTestLdapSettingsRequest) HasEncryptAuth() bool`

HasEncryptAuth returns a boolean if a field has been set.

### GetEncryptVerifyCa

`func (o *PostTestLdapSettingsRequest) GetEncryptVerifyCa() bool`

GetEncryptVerifyCa returns the EncryptVerifyCa field if non-nil, zero value otherwise.

### GetEncryptVerifyCaOk

`func (o *PostTestLdapSettingsRequest) GetEncryptVerifyCaOk() (*bool, bool)`

GetEncryptVerifyCaOk returns a tuple with the EncryptVerifyCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptVerifyCa

`func (o *PostTestLdapSettingsRequest) SetEncryptVerifyCa(v bool)`

SetEncryptVerifyCa sets EncryptVerifyCa field to given value.

### HasEncryptVerifyCa

`func (o *PostTestLdapSettingsRequest) HasEncryptVerifyCa() bool`

HasEncryptVerifyCa returns a boolean if a field has been set.

### GetBinddn

`func (o *PostTestLdapSettingsRequest) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *PostTestLdapSettingsRequest) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *PostTestLdapSettingsRequest) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *PostTestLdapSettingsRequest) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.

### GetBindpw

`func (o *PostTestLdapSettingsRequest) GetBindpw() string`

GetBindpw returns the Bindpw field if non-nil, zero value otherwise.

### GetBindpwOk

`func (o *PostTestLdapSettingsRequest) GetBindpwOk() (*string, bool)`

GetBindpwOk returns a tuple with the Bindpw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpw

`func (o *PostTestLdapSettingsRequest) SetBindpw(v string)`

SetBindpw sets Bindpw field to given value.

### HasBindpw

`func (o *PostTestLdapSettingsRequest) HasBindpw() bool`

HasBindpw returns a boolean if a field has been set.

### GetAuthCert

`func (o *PostTestLdapSettingsRequest) GetAuthCert() string`

GetAuthCert returns the AuthCert field if non-nil, zero value otherwise.

### GetAuthCertOk

`func (o *PostTestLdapSettingsRequest) GetAuthCertOk() (*string, bool)`

GetAuthCertOk returns a tuple with the AuthCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCert

`func (o *PostTestLdapSettingsRequest) SetAuthCert(v string)`

SetAuthCert sets AuthCert field to given value.

### HasAuthCert

`func (o *PostTestLdapSettingsRequest) HasAuthCert() bool`

HasAuthCert returns a boolean if a field has been set.

### GetAuthCertCurrent

`func (o *PostTestLdapSettingsRequest) GetAuthCertCurrent() bool`

GetAuthCertCurrent returns the AuthCertCurrent field if non-nil, zero value otherwise.

### GetAuthCertCurrentOk

`func (o *PostTestLdapSettingsRequest) GetAuthCertCurrentOk() (*bool, bool)`

GetAuthCertCurrentOk returns a tuple with the AuthCertCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertCurrent

`func (o *PostTestLdapSettingsRequest) SetAuthCertCurrent(v bool)`

SetAuthCertCurrent sets AuthCertCurrent field to given value.

### HasAuthCertCurrent

`func (o *PostTestLdapSettingsRequest) HasAuthCertCurrent() bool`

HasAuthCertCurrent returns a boolean if a field has been set.

### GetAuthKey

`func (o *PostTestLdapSettingsRequest) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *PostTestLdapSettingsRequest) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *PostTestLdapSettingsRequest) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.

### HasAuthKey

`func (o *PostTestLdapSettingsRequest) HasAuthKey() bool`

HasAuthKey returns a boolean if a field has been set.

### GetAuthKeyCurrent

`func (o *PostTestLdapSettingsRequest) GetAuthKeyCurrent() bool`

GetAuthKeyCurrent returns the AuthKeyCurrent field if non-nil, zero value otherwise.

### GetAuthKeyCurrentOk

`func (o *PostTestLdapSettingsRequest) GetAuthKeyCurrentOk() (*bool, bool)`

GetAuthKeyCurrentOk returns a tuple with the AuthKeyCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKeyCurrent

`func (o *PostTestLdapSettingsRequest) SetAuthKeyCurrent(v bool)`

SetAuthKeyCurrent sets AuthKeyCurrent field to given value.

### HasAuthKeyCurrent

`func (o *PostTestLdapSettingsRequest) HasAuthKeyCurrent() bool`

HasAuthKeyCurrent returns a boolean if a field has been set.

### GetCaCert

`func (o *PostTestLdapSettingsRequest) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *PostTestLdapSettingsRequest) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *PostTestLdapSettingsRequest) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *PostTestLdapSettingsRequest) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetCaCertCurrent

`func (o *PostTestLdapSettingsRequest) GetCaCertCurrent() bool`

GetCaCertCurrent returns the CaCertCurrent field if non-nil, zero value otherwise.

### GetCaCertCurrentOk

`func (o *PostTestLdapSettingsRequest) GetCaCertCurrentOk() (*bool, bool)`

GetCaCertCurrentOk returns a tuple with the CaCertCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertCurrent

`func (o *PostTestLdapSettingsRequest) SetCaCertCurrent(v bool)`

SetCaCertCurrent sets CaCertCurrent field to given value.

### HasCaCertCurrent

`func (o *PostTestLdapSettingsRequest) HasCaCertCurrent() bool`

HasCaCertCurrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


