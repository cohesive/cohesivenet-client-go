# IdentityLdapSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | IP address or resolvable hostname | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Encrypt** | Pointer to **bool** | Use SSL | [optional] 
**EncryptLdaps** | Pointer to **bool** | Use LDAPS or start TLS | [optional] 
**EncryptAuth** | Pointer to **bool** | Use certificates to authenticate via encrypted connection | [optional] 
**EncryptAuthKey** | Pointer to **bool** |  | [optional] 
**EncryptAuthCert** | Pointer to **bool** |  | [optional] 
**EncryptVerifyCa** | Pointer to **bool** | Verify certicate using authority | [optional] 
**EncryptCaCert** | Pointer to **bool** |  | [optional] 
**Binddn** | Pointer to **string** | Bind username | [optional] 
**Bindpw** | Pointer to **string** | Bind password | [optional] 
**EncryptAuthCertData** | Pointer to **string** | Authentication certificate text content to use, empty to delete | [optional] 
**EncryptAuthCertFilename** | Pointer to **string** | Authentication certificate filename | [optional] [default to "tls.cert"]
**EncryptAuthKeyData** | Pointer to **string** | Authentication key text content to use, empty to delete | [optional] 
**EncryptAuthKeyFilename** | Pointer to **string** | Authentication key filename | [optional] [default to "tls.key"]
**EncryptCaCertData** | Pointer to **string** | CA certificate text content to use, empty to delete | [optional] 
**EncryptCaCertFilename** | Pointer to **string** | CA certificate filename | [optional] [default to "ca.pem"]
**UserBase** | Pointer to **string** | Base DN from which to search for Users | [optional] 
**UserIdAttribute** | Pointer to **string** | Attribute type for the Users | [optional] 
**UserListFilter** | Pointer to **string** | Search filter for Users | [optional] 
**GroupBase** | Pointer to **string** | Base DN from which to search for Groups | [optional] 
**GroupIdAttribute** | Pointer to **string** | Attribute type for the Groups | [optional] 
**GroupListFilter** | Pointer to **string** | Search filter for Groups | [optional] 
**GroupMemberAttribute** | Pointer to **string** | Attribute used to search for a user within the Group | [optional] 
**GroupMemberAttrFormat** | Pointer to **string** | UserID or UserDN | [optional] 
**GroupSearchScope** | Pointer to **string** | base, single or subtree | [optional] 
**Otp** | Pointer to **bool** | Use OTP code | [optional] 
**OtpUrl** | Pointer to **string** |  | [optional] 
**Provider** | **string** |  | 
**Enabled** | **bool** |  | 

## Methods

### NewIdentityLdapSettingsRequest

`func NewIdentityLdapSettingsRequest(provider string, enabled bool, ) *IdentityLdapSettingsRequest`

NewIdentityLdapSettingsRequest instantiates a new IdentityLdapSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityLdapSettingsRequestWithDefaults

`func NewIdentityLdapSettingsRequestWithDefaults() *IdentityLdapSettingsRequest`

NewIdentityLdapSettingsRequestWithDefaults instantiates a new IdentityLdapSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *IdentityLdapSettingsRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IdentityLdapSettingsRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IdentityLdapSettingsRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IdentityLdapSettingsRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *IdentityLdapSettingsRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IdentityLdapSettingsRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IdentityLdapSettingsRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *IdentityLdapSettingsRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEncrypt

`func (o *IdentityLdapSettingsRequest) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *IdentityLdapSettingsRequest) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *IdentityLdapSettingsRequest) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *IdentityLdapSettingsRequest) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptLdaps

`func (o *IdentityLdapSettingsRequest) GetEncryptLdaps() bool`

GetEncryptLdaps returns the EncryptLdaps field if non-nil, zero value otherwise.

### GetEncryptLdapsOk

`func (o *IdentityLdapSettingsRequest) GetEncryptLdapsOk() (*bool, bool)`

GetEncryptLdapsOk returns a tuple with the EncryptLdaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptLdaps

`func (o *IdentityLdapSettingsRequest) SetEncryptLdaps(v bool)`

SetEncryptLdaps sets EncryptLdaps field to given value.

### HasEncryptLdaps

`func (o *IdentityLdapSettingsRequest) HasEncryptLdaps() bool`

HasEncryptLdaps returns a boolean if a field has been set.

### GetEncryptAuth

`func (o *IdentityLdapSettingsRequest) GetEncryptAuth() bool`

GetEncryptAuth returns the EncryptAuth field if non-nil, zero value otherwise.

### GetEncryptAuthOk

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthOk() (*bool, bool)`

GetEncryptAuthOk returns a tuple with the EncryptAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuth

`func (o *IdentityLdapSettingsRequest) SetEncryptAuth(v bool)`

SetEncryptAuth sets EncryptAuth field to given value.

### HasEncryptAuth

`func (o *IdentityLdapSettingsRequest) HasEncryptAuth() bool`

HasEncryptAuth returns a boolean if a field has been set.

### GetEncryptAuthKey

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthKey() bool`

GetEncryptAuthKey returns the EncryptAuthKey field if non-nil, zero value otherwise.

### GetEncryptAuthKeyOk

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthKeyOk() (*bool, bool)`

GetEncryptAuthKeyOk returns a tuple with the EncryptAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKey

`func (o *IdentityLdapSettingsRequest) SetEncryptAuthKey(v bool)`

SetEncryptAuthKey sets EncryptAuthKey field to given value.

### HasEncryptAuthKey

`func (o *IdentityLdapSettingsRequest) HasEncryptAuthKey() bool`

HasEncryptAuthKey returns a boolean if a field has been set.

### GetEncryptAuthCert

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthCert() bool`

GetEncryptAuthCert returns the EncryptAuthCert field if non-nil, zero value otherwise.

### GetEncryptAuthCertOk

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthCertOk() (*bool, bool)`

GetEncryptAuthCertOk returns a tuple with the EncryptAuthCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCert

`func (o *IdentityLdapSettingsRequest) SetEncryptAuthCert(v bool)`

SetEncryptAuthCert sets EncryptAuthCert field to given value.

### HasEncryptAuthCert

`func (o *IdentityLdapSettingsRequest) HasEncryptAuthCert() bool`

HasEncryptAuthCert returns a boolean if a field has been set.

### GetEncryptVerifyCa

`func (o *IdentityLdapSettingsRequest) GetEncryptVerifyCa() bool`

GetEncryptVerifyCa returns the EncryptVerifyCa field if non-nil, zero value otherwise.

### GetEncryptVerifyCaOk

`func (o *IdentityLdapSettingsRequest) GetEncryptVerifyCaOk() (*bool, bool)`

GetEncryptVerifyCaOk returns a tuple with the EncryptVerifyCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptVerifyCa

`func (o *IdentityLdapSettingsRequest) SetEncryptVerifyCa(v bool)`

SetEncryptVerifyCa sets EncryptVerifyCa field to given value.

### HasEncryptVerifyCa

`func (o *IdentityLdapSettingsRequest) HasEncryptVerifyCa() bool`

HasEncryptVerifyCa returns a boolean if a field has been set.

### GetEncryptCaCert

`func (o *IdentityLdapSettingsRequest) GetEncryptCaCert() bool`

GetEncryptCaCert returns the EncryptCaCert field if non-nil, zero value otherwise.

### GetEncryptCaCertOk

`func (o *IdentityLdapSettingsRequest) GetEncryptCaCertOk() (*bool, bool)`

GetEncryptCaCertOk returns a tuple with the EncryptCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCert

`func (o *IdentityLdapSettingsRequest) SetEncryptCaCert(v bool)`

SetEncryptCaCert sets EncryptCaCert field to given value.

### HasEncryptCaCert

`func (o *IdentityLdapSettingsRequest) HasEncryptCaCert() bool`

HasEncryptCaCert returns a boolean if a field has been set.

### GetBinddn

`func (o *IdentityLdapSettingsRequest) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *IdentityLdapSettingsRequest) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *IdentityLdapSettingsRequest) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *IdentityLdapSettingsRequest) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.

### GetBindpw

`func (o *IdentityLdapSettingsRequest) GetBindpw() string`

GetBindpw returns the Bindpw field if non-nil, zero value otherwise.

### GetBindpwOk

`func (o *IdentityLdapSettingsRequest) GetBindpwOk() (*string, bool)`

GetBindpwOk returns a tuple with the Bindpw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpw

`func (o *IdentityLdapSettingsRequest) SetBindpw(v string)`

SetBindpw sets Bindpw field to given value.

### HasBindpw

`func (o *IdentityLdapSettingsRequest) HasBindpw() bool`

HasBindpw returns a boolean if a field has been set.

### GetEncryptAuthCertData

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthCertData() string`

GetEncryptAuthCertData returns the EncryptAuthCertData field if non-nil, zero value otherwise.

### GetEncryptAuthCertDataOk

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthCertDataOk() (*string, bool)`

GetEncryptAuthCertDataOk returns a tuple with the EncryptAuthCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCertData

`func (o *IdentityLdapSettingsRequest) SetEncryptAuthCertData(v string)`

SetEncryptAuthCertData sets EncryptAuthCertData field to given value.

### HasEncryptAuthCertData

`func (o *IdentityLdapSettingsRequest) HasEncryptAuthCertData() bool`

HasEncryptAuthCertData returns a boolean if a field has been set.

### GetEncryptAuthCertFilename

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthCertFilename() string`

GetEncryptAuthCertFilename returns the EncryptAuthCertFilename field if non-nil, zero value otherwise.

### GetEncryptAuthCertFilenameOk

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthCertFilenameOk() (*string, bool)`

GetEncryptAuthCertFilenameOk returns a tuple with the EncryptAuthCertFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCertFilename

`func (o *IdentityLdapSettingsRequest) SetEncryptAuthCertFilename(v string)`

SetEncryptAuthCertFilename sets EncryptAuthCertFilename field to given value.

### HasEncryptAuthCertFilename

`func (o *IdentityLdapSettingsRequest) HasEncryptAuthCertFilename() bool`

HasEncryptAuthCertFilename returns a boolean if a field has been set.

### GetEncryptAuthKeyData

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthKeyData() string`

GetEncryptAuthKeyData returns the EncryptAuthKeyData field if non-nil, zero value otherwise.

### GetEncryptAuthKeyDataOk

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthKeyDataOk() (*string, bool)`

GetEncryptAuthKeyDataOk returns a tuple with the EncryptAuthKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKeyData

`func (o *IdentityLdapSettingsRequest) SetEncryptAuthKeyData(v string)`

SetEncryptAuthKeyData sets EncryptAuthKeyData field to given value.

### HasEncryptAuthKeyData

`func (o *IdentityLdapSettingsRequest) HasEncryptAuthKeyData() bool`

HasEncryptAuthKeyData returns a boolean if a field has been set.

### GetEncryptAuthKeyFilename

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthKeyFilename() string`

GetEncryptAuthKeyFilename returns the EncryptAuthKeyFilename field if non-nil, zero value otherwise.

### GetEncryptAuthKeyFilenameOk

`func (o *IdentityLdapSettingsRequest) GetEncryptAuthKeyFilenameOk() (*string, bool)`

GetEncryptAuthKeyFilenameOk returns a tuple with the EncryptAuthKeyFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKeyFilename

`func (o *IdentityLdapSettingsRequest) SetEncryptAuthKeyFilename(v string)`

SetEncryptAuthKeyFilename sets EncryptAuthKeyFilename field to given value.

### HasEncryptAuthKeyFilename

`func (o *IdentityLdapSettingsRequest) HasEncryptAuthKeyFilename() bool`

HasEncryptAuthKeyFilename returns a boolean if a field has been set.

### GetEncryptCaCertData

`func (o *IdentityLdapSettingsRequest) GetEncryptCaCertData() string`

GetEncryptCaCertData returns the EncryptCaCertData field if non-nil, zero value otherwise.

### GetEncryptCaCertDataOk

`func (o *IdentityLdapSettingsRequest) GetEncryptCaCertDataOk() (*string, bool)`

GetEncryptCaCertDataOk returns a tuple with the EncryptCaCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCertData

`func (o *IdentityLdapSettingsRequest) SetEncryptCaCertData(v string)`

SetEncryptCaCertData sets EncryptCaCertData field to given value.

### HasEncryptCaCertData

`func (o *IdentityLdapSettingsRequest) HasEncryptCaCertData() bool`

HasEncryptCaCertData returns a boolean if a field has been set.

### GetEncryptCaCertFilename

`func (o *IdentityLdapSettingsRequest) GetEncryptCaCertFilename() string`

GetEncryptCaCertFilename returns the EncryptCaCertFilename field if non-nil, zero value otherwise.

### GetEncryptCaCertFilenameOk

`func (o *IdentityLdapSettingsRequest) GetEncryptCaCertFilenameOk() (*string, bool)`

GetEncryptCaCertFilenameOk returns a tuple with the EncryptCaCertFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCertFilename

`func (o *IdentityLdapSettingsRequest) SetEncryptCaCertFilename(v string)`

SetEncryptCaCertFilename sets EncryptCaCertFilename field to given value.

### HasEncryptCaCertFilename

`func (o *IdentityLdapSettingsRequest) HasEncryptCaCertFilename() bool`

HasEncryptCaCertFilename returns a boolean if a field has been set.

### GetUserBase

`func (o *IdentityLdapSettingsRequest) GetUserBase() string`

GetUserBase returns the UserBase field if non-nil, zero value otherwise.

### GetUserBaseOk

`func (o *IdentityLdapSettingsRequest) GetUserBaseOk() (*string, bool)`

GetUserBaseOk returns a tuple with the UserBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBase

`func (o *IdentityLdapSettingsRequest) SetUserBase(v string)`

SetUserBase sets UserBase field to given value.

### HasUserBase

`func (o *IdentityLdapSettingsRequest) HasUserBase() bool`

HasUserBase returns a boolean if a field has been set.

### GetUserIdAttribute

`func (o *IdentityLdapSettingsRequest) GetUserIdAttribute() string`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *IdentityLdapSettingsRequest) GetUserIdAttributeOk() (*string, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *IdentityLdapSettingsRequest) SetUserIdAttribute(v string)`

SetUserIdAttribute sets UserIdAttribute field to given value.

### HasUserIdAttribute

`func (o *IdentityLdapSettingsRequest) HasUserIdAttribute() bool`

HasUserIdAttribute returns a boolean if a field has been set.

### GetUserListFilter

`func (o *IdentityLdapSettingsRequest) GetUserListFilter() string`

GetUserListFilter returns the UserListFilter field if non-nil, zero value otherwise.

### GetUserListFilterOk

`func (o *IdentityLdapSettingsRequest) GetUserListFilterOk() (*string, bool)`

GetUserListFilterOk returns a tuple with the UserListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListFilter

`func (o *IdentityLdapSettingsRequest) SetUserListFilter(v string)`

SetUserListFilter sets UserListFilter field to given value.

### HasUserListFilter

`func (o *IdentityLdapSettingsRequest) HasUserListFilter() bool`

HasUserListFilter returns a boolean if a field has been set.

### GetGroupBase

`func (o *IdentityLdapSettingsRequest) GetGroupBase() string`

GetGroupBase returns the GroupBase field if non-nil, zero value otherwise.

### GetGroupBaseOk

`func (o *IdentityLdapSettingsRequest) GetGroupBaseOk() (*string, bool)`

GetGroupBaseOk returns a tuple with the GroupBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBase

`func (o *IdentityLdapSettingsRequest) SetGroupBase(v string)`

SetGroupBase sets GroupBase field to given value.

### HasGroupBase

`func (o *IdentityLdapSettingsRequest) HasGroupBase() bool`

HasGroupBase returns a boolean if a field has been set.

### GetGroupIdAttribute

`func (o *IdentityLdapSettingsRequest) GetGroupIdAttribute() string`

GetGroupIdAttribute returns the GroupIdAttribute field if non-nil, zero value otherwise.

### GetGroupIdAttributeOk

`func (o *IdentityLdapSettingsRequest) GetGroupIdAttributeOk() (*string, bool)`

GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdAttribute

`func (o *IdentityLdapSettingsRequest) SetGroupIdAttribute(v string)`

SetGroupIdAttribute sets GroupIdAttribute field to given value.

### HasGroupIdAttribute

`func (o *IdentityLdapSettingsRequest) HasGroupIdAttribute() bool`

HasGroupIdAttribute returns a boolean if a field has been set.

### GetGroupListFilter

`func (o *IdentityLdapSettingsRequest) GetGroupListFilter() string`

GetGroupListFilter returns the GroupListFilter field if non-nil, zero value otherwise.

### GetGroupListFilterOk

`func (o *IdentityLdapSettingsRequest) GetGroupListFilterOk() (*string, bool)`

GetGroupListFilterOk returns a tuple with the GroupListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupListFilter

`func (o *IdentityLdapSettingsRequest) SetGroupListFilter(v string)`

SetGroupListFilter sets GroupListFilter field to given value.

### HasGroupListFilter

`func (o *IdentityLdapSettingsRequest) HasGroupListFilter() bool`

HasGroupListFilter returns a boolean if a field has been set.

### GetGroupMemberAttribute

`func (o *IdentityLdapSettingsRequest) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *IdentityLdapSettingsRequest) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *IdentityLdapSettingsRequest) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.

### HasGroupMemberAttribute

`func (o *IdentityLdapSettingsRequest) HasGroupMemberAttribute() bool`

HasGroupMemberAttribute returns a boolean if a field has been set.

### GetGroupMemberAttrFormat

`func (o *IdentityLdapSettingsRequest) GetGroupMemberAttrFormat() string`

GetGroupMemberAttrFormat returns the GroupMemberAttrFormat field if non-nil, zero value otherwise.

### GetGroupMemberAttrFormatOk

`func (o *IdentityLdapSettingsRequest) GetGroupMemberAttrFormatOk() (*string, bool)`

GetGroupMemberAttrFormatOk returns a tuple with the GroupMemberAttrFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttrFormat

`func (o *IdentityLdapSettingsRequest) SetGroupMemberAttrFormat(v string)`

SetGroupMemberAttrFormat sets GroupMemberAttrFormat field to given value.

### HasGroupMemberAttrFormat

`func (o *IdentityLdapSettingsRequest) HasGroupMemberAttrFormat() bool`

HasGroupMemberAttrFormat returns a boolean if a field has been set.

### GetGroupSearchScope

`func (o *IdentityLdapSettingsRequest) GetGroupSearchScope() string`

GetGroupSearchScope returns the GroupSearchScope field if non-nil, zero value otherwise.

### GetGroupSearchScopeOk

`func (o *IdentityLdapSettingsRequest) GetGroupSearchScopeOk() (*string, bool)`

GetGroupSearchScopeOk returns a tuple with the GroupSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchScope

`func (o *IdentityLdapSettingsRequest) SetGroupSearchScope(v string)`

SetGroupSearchScope sets GroupSearchScope field to given value.

### HasGroupSearchScope

`func (o *IdentityLdapSettingsRequest) HasGroupSearchScope() bool`

HasGroupSearchScope returns a boolean if a field has been set.

### GetOtp

`func (o *IdentityLdapSettingsRequest) GetOtp() bool`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *IdentityLdapSettingsRequest) GetOtpOk() (*bool, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *IdentityLdapSettingsRequest) SetOtp(v bool)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *IdentityLdapSettingsRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetOtpUrl

`func (o *IdentityLdapSettingsRequest) GetOtpUrl() string`

GetOtpUrl returns the OtpUrl field if non-nil, zero value otherwise.

### GetOtpUrlOk

`func (o *IdentityLdapSettingsRequest) GetOtpUrlOk() (*string, bool)`

GetOtpUrlOk returns a tuple with the OtpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpUrl

`func (o *IdentityLdapSettingsRequest) SetOtpUrl(v string)`

SetOtpUrl sets OtpUrl field to given value.

### HasOtpUrl

`func (o *IdentityLdapSettingsRequest) HasOtpUrl() bool`

HasOtpUrl returns a boolean if a field has been set.

### GetProvider

`func (o *IdentityLdapSettingsRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IdentityLdapSettingsRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IdentityLdapSettingsRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetEnabled

`func (o *IdentityLdapSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityLdapSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityLdapSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


