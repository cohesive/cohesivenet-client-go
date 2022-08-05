# IdentityLdapSettings

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

## Methods

### NewIdentityLdapSettings

`func NewIdentityLdapSettings() *IdentityLdapSettings`

NewIdentityLdapSettings instantiates a new IdentityLdapSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityLdapSettingsWithDefaults

`func NewIdentityLdapSettingsWithDefaults() *IdentityLdapSettings`

NewIdentityLdapSettingsWithDefaults instantiates a new IdentityLdapSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *IdentityLdapSettings) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IdentityLdapSettings) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IdentityLdapSettings) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IdentityLdapSettings) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *IdentityLdapSettings) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IdentityLdapSettings) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IdentityLdapSettings) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *IdentityLdapSettings) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEncrypt

`func (o *IdentityLdapSettings) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *IdentityLdapSettings) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *IdentityLdapSettings) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *IdentityLdapSettings) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptLdaps

`func (o *IdentityLdapSettings) GetEncryptLdaps() bool`

GetEncryptLdaps returns the EncryptLdaps field if non-nil, zero value otherwise.

### GetEncryptLdapsOk

`func (o *IdentityLdapSettings) GetEncryptLdapsOk() (*bool, bool)`

GetEncryptLdapsOk returns a tuple with the EncryptLdaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptLdaps

`func (o *IdentityLdapSettings) SetEncryptLdaps(v bool)`

SetEncryptLdaps sets EncryptLdaps field to given value.

### HasEncryptLdaps

`func (o *IdentityLdapSettings) HasEncryptLdaps() bool`

HasEncryptLdaps returns a boolean if a field has been set.

### GetEncryptAuth

`func (o *IdentityLdapSettings) GetEncryptAuth() bool`

GetEncryptAuth returns the EncryptAuth field if non-nil, zero value otherwise.

### GetEncryptAuthOk

`func (o *IdentityLdapSettings) GetEncryptAuthOk() (*bool, bool)`

GetEncryptAuthOk returns a tuple with the EncryptAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuth

`func (o *IdentityLdapSettings) SetEncryptAuth(v bool)`

SetEncryptAuth sets EncryptAuth field to given value.

### HasEncryptAuth

`func (o *IdentityLdapSettings) HasEncryptAuth() bool`

HasEncryptAuth returns a boolean if a field has been set.

### GetEncryptAuthKey

`func (o *IdentityLdapSettings) GetEncryptAuthKey() bool`

GetEncryptAuthKey returns the EncryptAuthKey field if non-nil, zero value otherwise.

### GetEncryptAuthKeyOk

`func (o *IdentityLdapSettings) GetEncryptAuthKeyOk() (*bool, bool)`

GetEncryptAuthKeyOk returns a tuple with the EncryptAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKey

`func (o *IdentityLdapSettings) SetEncryptAuthKey(v bool)`

SetEncryptAuthKey sets EncryptAuthKey field to given value.

### HasEncryptAuthKey

`func (o *IdentityLdapSettings) HasEncryptAuthKey() bool`

HasEncryptAuthKey returns a boolean if a field has been set.

### GetEncryptAuthCert

`func (o *IdentityLdapSettings) GetEncryptAuthCert() bool`

GetEncryptAuthCert returns the EncryptAuthCert field if non-nil, zero value otherwise.

### GetEncryptAuthCertOk

`func (o *IdentityLdapSettings) GetEncryptAuthCertOk() (*bool, bool)`

GetEncryptAuthCertOk returns a tuple with the EncryptAuthCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCert

`func (o *IdentityLdapSettings) SetEncryptAuthCert(v bool)`

SetEncryptAuthCert sets EncryptAuthCert field to given value.

### HasEncryptAuthCert

`func (o *IdentityLdapSettings) HasEncryptAuthCert() bool`

HasEncryptAuthCert returns a boolean if a field has been set.

### GetEncryptVerifyCa

`func (o *IdentityLdapSettings) GetEncryptVerifyCa() bool`

GetEncryptVerifyCa returns the EncryptVerifyCa field if non-nil, zero value otherwise.

### GetEncryptVerifyCaOk

`func (o *IdentityLdapSettings) GetEncryptVerifyCaOk() (*bool, bool)`

GetEncryptVerifyCaOk returns a tuple with the EncryptVerifyCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptVerifyCa

`func (o *IdentityLdapSettings) SetEncryptVerifyCa(v bool)`

SetEncryptVerifyCa sets EncryptVerifyCa field to given value.

### HasEncryptVerifyCa

`func (o *IdentityLdapSettings) HasEncryptVerifyCa() bool`

HasEncryptVerifyCa returns a boolean if a field has been set.

### GetEncryptCaCert

`func (o *IdentityLdapSettings) GetEncryptCaCert() bool`

GetEncryptCaCert returns the EncryptCaCert field if non-nil, zero value otherwise.

### GetEncryptCaCertOk

`func (o *IdentityLdapSettings) GetEncryptCaCertOk() (*bool, bool)`

GetEncryptCaCertOk returns a tuple with the EncryptCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCert

`func (o *IdentityLdapSettings) SetEncryptCaCert(v bool)`

SetEncryptCaCert sets EncryptCaCert field to given value.

### HasEncryptCaCert

`func (o *IdentityLdapSettings) HasEncryptCaCert() bool`

HasEncryptCaCert returns a boolean if a field has been set.

### GetBinddn

`func (o *IdentityLdapSettings) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *IdentityLdapSettings) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *IdentityLdapSettings) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *IdentityLdapSettings) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.

### GetBindpw

`func (o *IdentityLdapSettings) GetBindpw() string`

GetBindpw returns the Bindpw field if non-nil, zero value otherwise.

### GetBindpwOk

`func (o *IdentityLdapSettings) GetBindpwOk() (*string, bool)`

GetBindpwOk returns a tuple with the Bindpw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpw

`func (o *IdentityLdapSettings) SetBindpw(v string)`

SetBindpw sets Bindpw field to given value.

### HasBindpw

`func (o *IdentityLdapSettings) HasBindpw() bool`

HasBindpw returns a boolean if a field has been set.

### GetEncryptAuthCertData

`func (o *IdentityLdapSettings) GetEncryptAuthCertData() string`

GetEncryptAuthCertData returns the EncryptAuthCertData field if non-nil, zero value otherwise.

### GetEncryptAuthCertDataOk

`func (o *IdentityLdapSettings) GetEncryptAuthCertDataOk() (*string, bool)`

GetEncryptAuthCertDataOk returns a tuple with the EncryptAuthCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCertData

`func (o *IdentityLdapSettings) SetEncryptAuthCertData(v string)`

SetEncryptAuthCertData sets EncryptAuthCertData field to given value.

### HasEncryptAuthCertData

`func (o *IdentityLdapSettings) HasEncryptAuthCertData() bool`

HasEncryptAuthCertData returns a boolean if a field has been set.

### GetEncryptAuthCertFilename

`func (o *IdentityLdapSettings) GetEncryptAuthCertFilename() string`

GetEncryptAuthCertFilename returns the EncryptAuthCertFilename field if non-nil, zero value otherwise.

### GetEncryptAuthCertFilenameOk

`func (o *IdentityLdapSettings) GetEncryptAuthCertFilenameOk() (*string, bool)`

GetEncryptAuthCertFilenameOk returns a tuple with the EncryptAuthCertFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCertFilename

`func (o *IdentityLdapSettings) SetEncryptAuthCertFilename(v string)`

SetEncryptAuthCertFilename sets EncryptAuthCertFilename field to given value.

### HasEncryptAuthCertFilename

`func (o *IdentityLdapSettings) HasEncryptAuthCertFilename() bool`

HasEncryptAuthCertFilename returns a boolean if a field has been set.

### GetEncryptAuthKeyData

`func (o *IdentityLdapSettings) GetEncryptAuthKeyData() string`

GetEncryptAuthKeyData returns the EncryptAuthKeyData field if non-nil, zero value otherwise.

### GetEncryptAuthKeyDataOk

`func (o *IdentityLdapSettings) GetEncryptAuthKeyDataOk() (*string, bool)`

GetEncryptAuthKeyDataOk returns a tuple with the EncryptAuthKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKeyData

`func (o *IdentityLdapSettings) SetEncryptAuthKeyData(v string)`

SetEncryptAuthKeyData sets EncryptAuthKeyData field to given value.

### HasEncryptAuthKeyData

`func (o *IdentityLdapSettings) HasEncryptAuthKeyData() bool`

HasEncryptAuthKeyData returns a boolean if a field has been set.

### GetEncryptAuthKeyFilename

`func (o *IdentityLdapSettings) GetEncryptAuthKeyFilename() string`

GetEncryptAuthKeyFilename returns the EncryptAuthKeyFilename field if non-nil, zero value otherwise.

### GetEncryptAuthKeyFilenameOk

`func (o *IdentityLdapSettings) GetEncryptAuthKeyFilenameOk() (*string, bool)`

GetEncryptAuthKeyFilenameOk returns a tuple with the EncryptAuthKeyFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKeyFilename

`func (o *IdentityLdapSettings) SetEncryptAuthKeyFilename(v string)`

SetEncryptAuthKeyFilename sets EncryptAuthKeyFilename field to given value.

### HasEncryptAuthKeyFilename

`func (o *IdentityLdapSettings) HasEncryptAuthKeyFilename() bool`

HasEncryptAuthKeyFilename returns a boolean if a field has been set.

### GetEncryptCaCertData

`func (o *IdentityLdapSettings) GetEncryptCaCertData() string`

GetEncryptCaCertData returns the EncryptCaCertData field if non-nil, zero value otherwise.

### GetEncryptCaCertDataOk

`func (o *IdentityLdapSettings) GetEncryptCaCertDataOk() (*string, bool)`

GetEncryptCaCertDataOk returns a tuple with the EncryptCaCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCertData

`func (o *IdentityLdapSettings) SetEncryptCaCertData(v string)`

SetEncryptCaCertData sets EncryptCaCertData field to given value.

### HasEncryptCaCertData

`func (o *IdentityLdapSettings) HasEncryptCaCertData() bool`

HasEncryptCaCertData returns a boolean if a field has been set.

### GetEncryptCaCertFilename

`func (o *IdentityLdapSettings) GetEncryptCaCertFilename() string`

GetEncryptCaCertFilename returns the EncryptCaCertFilename field if non-nil, zero value otherwise.

### GetEncryptCaCertFilenameOk

`func (o *IdentityLdapSettings) GetEncryptCaCertFilenameOk() (*string, bool)`

GetEncryptCaCertFilenameOk returns a tuple with the EncryptCaCertFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCertFilename

`func (o *IdentityLdapSettings) SetEncryptCaCertFilename(v string)`

SetEncryptCaCertFilename sets EncryptCaCertFilename field to given value.

### HasEncryptCaCertFilename

`func (o *IdentityLdapSettings) HasEncryptCaCertFilename() bool`

HasEncryptCaCertFilename returns a boolean if a field has been set.

### GetUserBase

`func (o *IdentityLdapSettings) GetUserBase() string`

GetUserBase returns the UserBase field if non-nil, zero value otherwise.

### GetUserBaseOk

`func (o *IdentityLdapSettings) GetUserBaseOk() (*string, bool)`

GetUserBaseOk returns a tuple with the UserBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBase

`func (o *IdentityLdapSettings) SetUserBase(v string)`

SetUserBase sets UserBase field to given value.

### HasUserBase

`func (o *IdentityLdapSettings) HasUserBase() bool`

HasUserBase returns a boolean if a field has been set.

### GetUserIdAttribute

`func (o *IdentityLdapSettings) GetUserIdAttribute() string`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *IdentityLdapSettings) GetUserIdAttributeOk() (*string, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *IdentityLdapSettings) SetUserIdAttribute(v string)`

SetUserIdAttribute sets UserIdAttribute field to given value.

### HasUserIdAttribute

`func (o *IdentityLdapSettings) HasUserIdAttribute() bool`

HasUserIdAttribute returns a boolean if a field has been set.

### GetUserListFilter

`func (o *IdentityLdapSettings) GetUserListFilter() string`

GetUserListFilter returns the UserListFilter field if non-nil, zero value otherwise.

### GetUserListFilterOk

`func (o *IdentityLdapSettings) GetUserListFilterOk() (*string, bool)`

GetUserListFilterOk returns a tuple with the UserListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListFilter

`func (o *IdentityLdapSettings) SetUserListFilter(v string)`

SetUserListFilter sets UserListFilter field to given value.

### HasUserListFilter

`func (o *IdentityLdapSettings) HasUserListFilter() bool`

HasUserListFilter returns a boolean if a field has been set.

### GetGroupBase

`func (o *IdentityLdapSettings) GetGroupBase() string`

GetGroupBase returns the GroupBase field if non-nil, zero value otherwise.

### GetGroupBaseOk

`func (o *IdentityLdapSettings) GetGroupBaseOk() (*string, bool)`

GetGroupBaseOk returns a tuple with the GroupBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBase

`func (o *IdentityLdapSettings) SetGroupBase(v string)`

SetGroupBase sets GroupBase field to given value.

### HasGroupBase

`func (o *IdentityLdapSettings) HasGroupBase() bool`

HasGroupBase returns a boolean if a field has been set.

### GetGroupIdAttribute

`func (o *IdentityLdapSettings) GetGroupIdAttribute() string`

GetGroupIdAttribute returns the GroupIdAttribute field if non-nil, zero value otherwise.

### GetGroupIdAttributeOk

`func (o *IdentityLdapSettings) GetGroupIdAttributeOk() (*string, bool)`

GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdAttribute

`func (o *IdentityLdapSettings) SetGroupIdAttribute(v string)`

SetGroupIdAttribute sets GroupIdAttribute field to given value.

### HasGroupIdAttribute

`func (o *IdentityLdapSettings) HasGroupIdAttribute() bool`

HasGroupIdAttribute returns a boolean if a field has been set.

### GetGroupListFilter

`func (o *IdentityLdapSettings) GetGroupListFilter() string`

GetGroupListFilter returns the GroupListFilter field if non-nil, zero value otherwise.

### GetGroupListFilterOk

`func (o *IdentityLdapSettings) GetGroupListFilterOk() (*string, bool)`

GetGroupListFilterOk returns a tuple with the GroupListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupListFilter

`func (o *IdentityLdapSettings) SetGroupListFilter(v string)`

SetGroupListFilter sets GroupListFilter field to given value.

### HasGroupListFilter

`func (o *IdentityLdapSettings) HasGroupListFilter() bool`

HasGroupListFilter returns a boolean if a field has been set.

### GetGroupMemberAttribute

`func (o *IdentityLdapSettings) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *IdentityLdapSettings) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *IdentityLdapSettings) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.

### HasGroupMemberAttribute

`func (o *IdentityLdapSettings) HasGroupMemberAttribute() bool`

HasGroupMemberAttribute returns a boolean if a field has been set.

### GetGroupMemberAttrFormat

`func (o *IdentityLdapSettings) GetGroupMemberAttrFormat() string`

GetGroupMemberAttrFormat returns the GroupMemberAttrFormat field if non-nil, zero value otherwise.

### GetGroupMemberAttrFormatOk

`func (o *IdentityLdapSettings) GetGroupMemberAttrFormatOk() (*string, bool)`

GetGroupMemberAttrFormatOk returns a tuple with the GroupMemberAttrFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttrFormat

`func (o *IdentityLdapSettings) SetGroupMemberAttrFormat(v string)`

SetGroupMemberAttrFormat sets GroupMemberAttrFormat field to given value.

### HasGroupMemberAttrFormat

`func (o *IdentityLdapSettings) HasGroupMemberAttrFormat() bool`

HasGroupMemberAttrFormat returns a boolean if a field has been set.

### GetGroupSearchScope

`func (o *IdentityLdapSettings) GetGroupSearchScope() string`

GetGroupSearchScope returns the GroupSearchScope field if non-nil, zero value otherwise.

### GetGroupSearchScopeOk

`func (o *IdentityLdapSettings) GetGroupSearchScopeOk() (*string, bool)`

GetGroupSearchScopeOk returns a tuple with the GroupSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchScope

`func (o *IdentityLdapSettings) SetGroupSearchScope(v string)`

SetGroupSearchScope sets GroupSearchScope field to given value.

### HasGroupSearchScope

`func (o *IdentityLdapSettings) HasGroupSearchScope() bool`

HasGroupSearchScope returns a boolean if a field has been set.

### GetOtp

`func (o *IdentityLdapSettings) GetOtp() bool`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *IdentityLdapSettings) GetOtpOk() (*bool, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *IdentityLdapSettings) SetOtp(v bool)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *IdentityLdapSettings) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetOtpUrl

`func (o *IdentityLdapSettings) GetOtpUrl() string`

GetOtpUrl returns the OtpUrl field if non-nil, zero value otherwise.

### GetOtpUrlOk

`func (o *IdentityLdapSettings) GetOtpUrlOk() (*string, bool)`

GetOtpUrlOk returns a tuple with the OtpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpUrl

`func (o *IdentityLdapSettings) SetOtpUrl(v string)`

SetOtpUrl sets OtpUrl field to given value.

### HasOtpUrl

`func (o *IdentityLdapSettings) HasOtpUrl() bool`

HasOtpUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


