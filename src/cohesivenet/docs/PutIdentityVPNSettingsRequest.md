# PutIdentityVPNSettingsRequest

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
**Identifier** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**RedirectHostname** | Pointer to **string** |  | [optional] 
**AuthorizationEndpoint** | Pointer to **string** |  | [optional] 
**TokenEndpoint** | Pointer to **string** |  | [optional] 
**UserinfoEndpoint** | Pointer to **string** |  | [optional] 
**JwksUri** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Keys** | Pointer to [**IdentityOIDCSettingsKeys**](IdentityOIDCSettingsKeys.md) |  | [optional] 
**Server** | Pointer to **string** | IP address or resolvable hostname | [optional] 
**AuthPort** | Pointer to **int32** | Authentication port | [optional] [default to 1812]
**AccountingPort** | Pointer to **int32** |  | [optional] [default to 1812]
**Pass** | Pointer to **string** | Shared password | [optional] 

## Methods

### NewPutIdentityVPNSettingsRequest

`func NewPutIdentityVPNSettingsRequest(provider string, enabled bool, ) *PutIdentityVPNSettingsRequest`

NewPutIdentityVPNSettingsRequest instantiates a new PutIdentityVPNSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutIdentityVPNSettingsRequestWithDefaults

`func NewPutIdentityVPNSettingsRequestWithDefaults() *PutIdentityVPNSettingsRequest`

NewPutIdentityVPNSettingsRequestWithDefaults instantiates a new PutIdentityVPNSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *PutIdentityVPNSettingsRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PutIdentityVPNSettingsRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PutIdentityVPNSettingsRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PutIdentityVPNSettingsRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *PutIdentityVPNSettingsRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PutIdentityVPNSettingsRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PutIdentityVPNSettingsRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PutIdentityVPNSettingsRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEncrypt

`func (o *PutIdentityVPNSettingsRequest) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *PutIdentityVPNSettingsRequest) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *PutIdentityVPNSettingsRequest) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptLdaps

`func (o *PutIdentityVPNSettingsRequest) GetEncryptLdaps() bool`

GetEncryptLdaps returns the EncryptLdaps field if non-nil, zero value otherwise.

### GetEncryptLdapsOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptLdapsOk() (*bool, bool)`

GetEncryptLdapsOk returns a tuple with the EncryptLdaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptLdaps

`func (o *PutIdentityVPNSettingsRequest) SetEncryptLdaps(v bool)`

SetEncryptLdaps sets EncryptLdaps field to given value.

### HasEncryptLdaps

`func (o *PutIdentityVPNSettingsRequest) HasEncryptLdaps() bool`

HasEncryptLdaps returns a boolean if a field has been set.

### GetEncryptAuth

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuth() bool`

GetEncryptAuth returns the EncryptAuth field if non-nil, zero value otherwise.

### GetEncryptAuthOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthOk() (*bool, bool)`

GetEncryptAuthOk returns a tuple with the EncryptAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuth

`func (o *PutIdentityVPNSettingsRequest) SetEncryptAuth(v bool)`

SetEncryptAuth sets EncryptAuth field to given value.

### HasEncryptAuth

`func (o *PutIdentityVPNSettingsRequest) HasEncryptAuth() bool`

HasEncryptAuth returns a boolean if a field has been set.

### GetEncryptAuthKey

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthKey() bool`

GetEncryptAuthKey returns the EncryptAuthKey field if non-nil, zero value otherwise.

### GetEncryptAuthKeyOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthKeyOk() (*bool, bool)`

GetEncryptAuthKeyOk returns a tuple with the EncryptAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKey

`func (o *PutIdentityVPNSettingsRequest) SetEncryptAuthKey(v bool)`

SetEncryptAuthKey sets EncryptAuthKey field to given value.

### HasEncryptAuthKey

`func (o *PutIdentityVPNSettingsRequest) HasEncryptAuthKey() bool`

HasEncryptAuthKey returns a boolean if a field has been set.

### GetEncryptAuthCert

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthCert() bool`

GetEncryptAuthCert returns the EncryptAuthCert field if non-nil, zero value otherwise.

### GetEncryptAuthCertOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthCertOk() (*bool, bool)`

GetEncryptAuthCertOk returns a tuple with the EncryptAuthCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCert

`func (o *PutIdentityVPNSettingsRequest) SetEncryptAuthCert(v bool)`

SetEncryptAuthCert sets EncryptAuthCert field to given value.

### HasEncryptAuthCert

`func (o *PutIdentityVPNSettingsRequest) HasEncryptAuthCert() bool`

HasEncryptAuthCert returns a boolean if a field has been set.

### GetEncryptVerifyCa

`func (o *PutIdentityVPNSettingsRequest) GetEncryptVerifyCa() bool`

GetEncryptVerifyCa returns the EncryptVerifyCa field if non-nil, zero value otherwise.

### GetEncryptVerifyCaOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptVerifyCaOk() (*bool, bool)`

GetEncryptVerifyCaOk returns a tuple with the EncryptVerifyCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptVerifyCa

`func (o *PutIdentityVPNSettingsRequest) SetEncryptVerifyCa(v bool)`

SetEncryptVerifyCa sets EncryptVerifyCa field to given value.

### HasEncryptVerifyCa

`func (o *PutIdentityVPNSettingsRequest) HasEncryptVerifyCa() bool`

HasEncryptVerifyCa returns a boolean if a field has been set.

### GetEncryptCaCert

`func (o *PutIdentityVPNSettingsRequest) GetEncryptCaCert() bool`

GetEncryptCaCert returns the EncryptCaCert field if non-nil, zero value otherwise.

### GetEncryptCaCertOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptCaCertOk() (*bool, bool)`

GetEncryptCaCertOk returns a tuple with the EncryptCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCert

`func (o *PutIdentityVPNSettingsRequest) SetEncryptCaCert(v bool)`

SetEncryptCaCert sets EncryptCaCert field to given value.

### HasEncryptCaCert

`func (o *PutIdentityVPNSettingsRequest) HasEncryptCaCert() bool`

HasEncryptCaCert returns a boolean if a field has been set.

### GetBinddn

`func (o *PutIdentityVPNSettingsRequest) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *PutIdentityVPNSettingsRequest) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *PutIdentityVPNSettingsRequest) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *PutIdentityVPNSettingsRequest) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.

### GetBindpw

`func (o *PutIdentityVPNSettingsRequest) GetBindpw() string`

GetBindpw returns the Bindpw field if non-nil, zero value otherwise.

### GetBindpwOk

`func (o *PutIdentityVPNSettingsRequest) GetBindpwOk() (*string, bool)`

GetBindpwOk returns a tuple with the Bindpw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpw

`func (o *PutIdentityVPNSettingsRequest) SetBindpw(v string)`

SetBindpw sets Bindpw field to given value.

### HasBindpw

`func (o *PutIdentityVPNSettingsRequest) HasBindpw() bool`

HasBindpw returns a boolean if a field has been set.

### GetEncryptAuthCertData

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthCertData() string`

GetEncryptAuthCertData returns the EncryptAuthCertData field if non-nil, zero value otherwise.

### GetEncryptAuthCertDataOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthCertDataOk() (*string, bool)`

GetEncryptAuthCertDataOk returns a tuple with the EncryptAuthCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCertData

`func (o *PutIdentityVPNSettingsRequest) SetEncryptAuthCertData(v string)`

SetEncryptAuthCertData sets EncryptAuthCertData field to given value.

### HasEncryptAuthCertData

`func (o *PutIdentityVPNSettingsRequest) HasEncryptAuthCertData() bool`

HasEncryptAuthCertData returns a boolean if a field has been set.

### GetEncryptAuthCertFilename

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthCertFilename() string`

GetEncryptAuthCertFilename returns the EncryptAuthCertFilename field if non-nil, zero value otherwise.

### GetEncryptAuthCertFilenameOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthCertFilenameOk() (*string, bool)`

GetEncryptAuthCertFilenameOk returns a tuple with the EncryptAuthCertFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCertFilename

`func (o *PutIdentityVPNSettingsRequest) SetEncryptAuthCertFilename(v string)`

SetEncryptAuthCertFilename sets EncryptAuthCertFilename field to given value.

### HasEncryptAuthCertFilename

`func (o *PutIdentityVPNSettingsRequest) HasEncryptAuthCertFilename() bool`

HasEncryptAuthCertFilename returns a boolean if a field has been set.

### GetEncryptAuthKeyData

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthKeyData() string`

GetEncryptAuthKeyData returns the EncryptAuthKeyData field if non-nil, zero value otherwise.

### GetEncryptAuthKeyDataOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthKeyDataOk() (*string, bool)`

GetEncryptAuthKeyDataOk returns a tuple with the EncryptAuthKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKeyData

`func (o *PutIdentityVPNSettingsRequest) SetEncryptAuthKeyData(v string)`

SetEncryptAuthKeyData sets EncryptAuthKeyData field to given value.

### HasEncryptAuthKeyData

`func (o *PutIdentityVPNSettingsRequest) HasEncryptAuthKeyData() bool`

HasEncryptAuthKeyData returns a boolean if a field has been set.

### GetEncryptAuthKeyFilename

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthKeyFilename() string`

GetEncryptAuthKeyFilename returns the EncryptAuthKeyFilename field if non-nil, zero value otherwise.

### GetEncryptAuthKeyFilenameOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptAuthKeyFilenameOk() (*string, bool)`

GetEncryptAuthKeyFilenameOk returns a tuple with the EncryptAuthKeyFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKeyFilename

`func (o *PutIdentityVPNSettingsRequest) SetEncryptAuthKeyFilename(v string)`

SetEncryptAuthKeyFilename sets EncryptAuthKeyFilename field to given value.

### HasEncryptAuthKeyFilename

`func (o *PutIdentityVPNSettingsRequest) HasEncryptAuthKeyFilename() bool`

HasEncryptAuthKeyFilename returns a boolean if a field has been set.

### GetEncryptCaCertData

`func (o *PutIdentityVPNSettingsRequest) GetEncryptCaCertData() string`

GetEncryptCaCertData returns the EncryptCaCertData field if non-nil, zero value otherwise.

### GetEncryptCaCertDataOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptCaCertDataOk() (*string, bool)`

GetEncryptCaCertDataOk returns a tuple with the EncryptCaCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCertData

`func (o *PutIdentityVPNSettingsRequest) SetEncryptCaCertData(v string)`

SetEncryptCaCertData sets EncryptCaCertData field to given value.

### HasEncryptCaCertData

`func (o *PutIdentityVPNSettingsRequest) HasEncryptCaCertData() bool`

HasEncryptCaCertData returns a boolean if a field has been set.

### GetEncryptCaCertFilename

`func (o *PutIdentityVPNSettingsRequest) GetEncryptCaCertFilename() string`

GetEncryptCaCertFilename returns the EncryptCaCertFilename field if non-nil, zero value otherwise.

### GetEncryptCaCertFilenameOk

`func (o *PutIdentityVPNSettingsRequest) GetEncryptCaCertFilenameOk() (*string, bool)`

GetEncryptCaCertFilenameOk returns a tuple with the EncryptCaCertFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCertFilename

`func (o *PutIdentityVPNSettingsRequest) SetEncryptCaCertFilename(v string)`

SetEncryptCaCertFilename sets EncryptCaCertFilename field to given value.

### HasEncryptCaCertFilename

`func (o *PutIdentityVPNSettingsRequest) HasEncryptCaCertFilename() bool`

HasEncryptCaCertFilename returns a boolean if a field has been set.

### GetUserBase

`func (o *PutIdentityVPNSettingsRequest) GetUserBase() string`

GetUserBase returns the UserBase field if non-nil, zero value otherwise.

### GetUserBaseOk

`func (o *PutIdentityVPNSettingsRequest) GetUserBaseOk() (*string, bool)`

GetUserBaseOk returns a tuple with the UserBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBase

`func (o *PutIdentityVPNSettingsRequest) SetUserBase(v string)`

SetUserBase sets UserBase field to given value.

### HasUserBase

`func (o *PutIdentityVPNSettingsRequest) HasUserBase() bool`

HasUserBase returns a boolean if a field has been set.

### GetUserIdAttribute

`func (o *PutIdentityVPNSettingsRequest) GetUserIdAttribute() string`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *PutIdentityVPNSettingsRequest) GetUserIdAttributeOk() (*string, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *PutIdentityVPNSettingsRequest) SetUserIdAttribute(v string)`

SetUserIdAttribute sets UserIdAttribute field to given value.

### HasUserIdAttribute

`func (o *PutIdentityVPNSettingsRequest) HasUserIdAttribute() bool`

HasUserIdAttribute returns a boolean if a field has been set.

### GetUserListFilter

`func (o *PutIdentityVPNSettingsRequest) GetUserListFilter() string`

GetUserListFilter returns the UserListFilter field if non-nil, zero value otherwise.

### GetUserListFilterOk

`func (o *PutIdentityVPNSettingsRequest) GetUserListFilterOk() (*string, bool)`

GetUserListFilterOk returns a tuple with the UserListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListFilter

`func (o *PutIdentityVPNSettingsRequest) SetUserListFilter(v string)`

SetUserListFilter sets UserListFilter field to given value.

### HasUserListFilter

`func (o *PutIdentityVPNSettingsRequest) HasUserListFilter() bool`

HasUserListFilter returns a boolean if a field has been set.

### GetGroupBase

`func (o *PutIdentityVPNSettingsRequest) GetGroupBase() string`

GetGroupBase returns the GroupBase field if non-nil, zero value otherwise.

### GetGroupBaseOk

`func (o *PutIdentityVPNSettingsRequest) GetGroupBaseOk() (*string, bool)`

GetGroupBaseOk returns a tuple with the GroupBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBase

`func (o *PutIdentityVPNSettingsRequest) SetGroupBase(v string)`

SetGroupBase sets GroupBase field to given value.

### HasGroupBase

`func (o *PutIdentityVPNSettingsRequest) HasGroupBase() bool`

HasGroupBase returns a boolean if a field has been set.

### GetGroupIdAttribute

`func (o *PutIdentityVPNSettingsRequest) GetGroupIdAttribute() string`

GetGroupIdAttribute returns the GroupIdAttribute field if non-nil, zero value otherwise.

### GetGroupIdAttributeOk

`func (o *PutIdentityVPNSettingsRequest) GetGroupIdAttributeOk() (*string, bool)`

GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdAttribute

`func (o *PutIdentityVPNSettingsRequest) SetGroupIdAttribute(v string)`

SetGroupIdAttribute sets GroupIdAttribute field to given value.

### HasGroupIdAttribute

`func (o *PutIdentityVPNSettingsRequest) HasGroupIdAttribute() bool`

HasGroupIdAttribute returns a boolean if a field has been set.

### GetGroupListFilter

`func (o *PutIdentityVPNSettingsRequest) GetGroupListFilter() string`

GetGroupListFilter returns the GroupListFilter field if non-nil, zero value otherwise.

### GetGroupListFilterOk

`func (o *PutIdentityVPNSettingsRequest) GetGroupListFilterOk() (*string, bool)`

GetGroupListFilterOk returns a tuple with the GroupListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupListFilter

`func (o *PutIdentityVPNSettingsRequest) SetGroupListFilter(v string)`

SetGroupListFilter sets GroupListFilter field to given value.

### HasGroupListFilter

`func (o *PutIdentityVPNSettingsRequest) HasGroupListFilter() bool`

HasGroupListFilter returns a boolean if a field has been set.

### GetGroupMemberAttribute

`func (o *PutIdentityVPNSettingsRequest) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *PutIdentityVPNSettingsRequest) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *PutIdentityVPNSettingsRequest) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.

### HasGroupMemberAttribute

`func (o *PutIdentityVPNSettingsRequest) HasGroupMemberAttribute() bool`

HasGroupMemberAttribute returns a boolean if a field has been set.

### GetGroupMemberAttrFormat

`func (o *PutIdentityVPNSettingsRequest) GetGroupMemberAttrFormat() string`

GetGroupMemberAttrFormat returns the GroupMemberAttrFormat field if non-nil, zero value otherwise.

### GetGroupMemberAttrFormatOk

`func (o *PutIdentityVPNSettingsRequest) GetGroupMemberAttrFormatOk() (*string, bool)`

GetGroupMemberAttrFormatOk returns a tuple with the GroupMemberAttrFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttrFormat

`func (o *PutIdentityVPNSettingsRequest) SetGroupMemberAttrFormat(v string)`

SetGroupMemberAttrFormat sets GroupMemberAttrFormat field to given value.

### HasGroupMemberAttrFormat

`func (o *PutIdentityVPNSettingsRequest) HasGroupMemberAttrFormat() bool`

HasGroupMemberAttrFormat returns a boolean if a field has been set.

### GetGroupSearchScope

`func (o *PutIdentityVPNSettingsRequest) GetGroupSearchScope() string`

GetGroupSearchScope returns the GroupSearchScope field if non-nil, zero value otherwise.

### GetGroupSearchScopeOk

`func (o *PutIdentityVPNSettingsRequest) GetGroupSearchScopeOk() (*string, bool)`

GetGroupSearchScopeOk returns a tuple with the GroupSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchScope

`func (o *PutIdentityVPNSettingsRequest) SetGroupSearchScope(v string)`

SetGroupSearchScope sets GroupSearchScope field to given value.

### HasGroupSearchScope

`func (o *PutIdentityVPNSettingsRequest) HasGroupSearchScope() bool`

HasGroupSearchScope returns a boolean if a field has been set.

### GetOtp

`func (o *PutIdentityVPNSettingsRequest) GetOtp() bool`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *PutIdentityVPNSettingsRequest) GetOtpOk() (*bool, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *PutIdentityVPNSettingsRequest) SetOtp(v bool)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *PutIdentityVPNSettingsRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetOtpUrl

`func (o *PutIdentityVPNSettingsRequest) GetOtpUrl() string`

GetOtpUrl returns the OtpUrl field if non-nil, zero value otherwise.

### GetOtpUrlOk

`func (o *PutIdentityVPNSettingsRequest) GetOtpUrlOk() (*string, bool)`

GetOtpUrlOk returns a tuple with the OtpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpUrl

`func (o *PutIdentityVPNSettingsRequest) SetOtpUrl(v string)`

SetOtpUrl sets OtpUrl field to given value.

### HasOtpUrl

`func (o *PutIdentityVPNSettingsRequest) HasOtpUrl() bool`

HasOtpUrl returns a boolean if a field has been set.

### GetProvider

`func (o *PutIdentityVPNSettingsRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PutIdentityVPNSettingsRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PutIdentityVPNSettingsRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetEnabled

`func (o *PutIdentityVPNSettingsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PutIdentityVPNSettingsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PutIdentityVPNSettingsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIdentifier

`func (o *PutIdentityVPNSettingsRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *PutIdentityVPNSettingsRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *PutIdentityVPNSettingsRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *PutIdentityVPNSettingsRequest) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetSecret

`func (o *PutIdentityVPNSettingsRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PutIdentityVPNSettingsRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PutIdentityVPNSettingsRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PutIdentityVPNSettingsRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRedirectHostname

`func (o *PutIdentityVPNSettingsRequest) GetRedirectHostname() string`

GetRedirectHostname returns the RedirectHostname field if non-nil, zero value otherwise.

### GetRedirectHostnameOk

`func (o *PutIdentityVPNSettingsRequest) GetRedirectHostnameOk() (*string, bool)`

GetRedirectHostnameOk returns a tuple with the RedirectHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHostname

`func (o *PutIdentityVPNSettingsRequest) SetRedirectHostname(v string)`

SetRedirectHostname sets RedirectHostname field to given value.

### HasRedirectHostname

`func (o *PutIdentityVPNSettingsRequest) HasRedirectHostname() bool`

HasRedirectHostname returns a boolean if a field has been set.

### GetAuthorizationEndpoint

`func (o *PutIdentityVPNSettingsRequest) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *PutIdentityVPNSettingsRequest) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *PutIdentityVPNSettingsRequest) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *PutIdentityVPNSettingsRequest) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *PutIdentityVPNSettingsRequest) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *PutIdentityVPNSettingsRequest) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *PutIdentityVPNSettingsRequest) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *PutIdentityVPNSettingsRequest) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetUserinfoEndpoint

`func (o *PutIdentityVPNSettingsRequest) GetUserinfoEndpoint() string`

GetUserinfoEndpoint returns the UserinfoEndpoint field if non-nil, zero value otherwise.

### GetUserinfoEndpointOk

`func (o *PutIdentityVPNSettingsRequest) GetUserinfoEndpointOk() (*string, bool)`

GetUserinfoEndpointOk returns a tuple with the UserinfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserinfoEndpoint

`func (o *PutIdentityVPNSettingsRequest) SetUserinfoEndpoint(v string)`

SetUserinfoEndpoint sets UserinfoEndpoint field to given value.

### HasUserinfoEndpoint

`func (o *PutIdentityVPNSettingsRequest) HasUserinfoEndpoint() bool`

HasUserinfoEndpoint returns a boolean if a field has been set.

### GetJwksUri

`func (o *PutIdentityVPNSettingsRequest) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *PutIdentityVPNSettingsRequest) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *PutIdentityVPNSettingsRequest) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *PutIdentityVPNSettingsRequest) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetIssuer

`func (o *PutIdentityVPNSettingsRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *PutIdentityVPNSettingsRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *PutIdentityVPNSettingsRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *PutIdentityVPNSettingsRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetKeys

`func (o *PutIdentityVPNSettingsRequest) GetKeys() IdentityOIDCSettingsKeys`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *PutIdentityVPNSettingsRequest) GetKeysOk() (*IdentityOIDCSettingsKeys, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *PutIdentityVPNSettingsRequest) SetKeys(v IdentityOIDCSettingsKeys)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *PutIdentityVPNSettingsRequest) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetServer

`func (o *PutIdentityVPNSettingsRequest) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PutIdentityVPNSettingsRequest) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PutIdentityVPNSettingsRequest) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *PutIdentityVPNSettingsRequest) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetAuthPort

`func (o *PutIdentityVPNSettingsRequest) GetAuthPort() int32`

GetAuthPort returns the AuthPort field if non-nil, zero value otherwise.

### GetAuthPortOk

`func (o *PutIdentityVPNSettingsRequest) GetAuthPortOk() (*int32, bool)`

GetAuthPortOk returns a tuple with the AuthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPort

`func (o *PutIdentityVPNSettingsRequest) SetAuthPort(v int32)`

SetAuthPort sets AuthPort field to given value.

### HasAuthPort

`func (o *PutIdentityVPNSettingsRequest) HasAuthPort() bool`

HasAuthPort returns a boolean if a field has been set.

### GetAccountingPort

`func (o *PutIdentityVPNSettingsRequest) GetAccountingPort() int32`

GetAccountingPort returns the AccountingPort field if non-nil, zero value otherwise.

### GetAccountingPortOk

`func (o *PutIdentityVPNSettingsRequest) GetAccountingPortOk() (*int32, bool)`

GetAccountingPortOk returns a tuple with the AccountingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingPort

`func (o *PutIdentityVPNSettingsRequest) SetAccountingPort(v int32)`

SetAccountingPort sets AccountingPort field to given value.

### HasAccountingPort

`func (o *PutIdentityVPNSettingsRequest) HasAccountingPort() bool`

HasAccountingPort returns a boolean if a field has been set.

### GetPass

`func (o *PutIdentityVPNSettingsRequest) GetPass() string`

GetPass returns the Pass field if non-nil, zero value otherwise.

### GetPassOk

`func (o *PutIdentityVPNSettingsRequest) GetPassOk() (*string, bool)`

GetPassOk returns a tuple with the Pass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPass

`func (o *PutIdentityVPNSettingsRequest) SetPass(v string)`

SetPass sets Pass field to given value.

### HasPass

`func (o *PutIdentityVPNSettingsRequest) HasPass() bool`

HasPass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


