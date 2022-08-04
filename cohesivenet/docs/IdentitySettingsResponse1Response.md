# IdentitySettingsResponse1Response

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
**Enabled** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**RedirectHostname** | Pointer to **string** |  | [optional] 
**AuthorizationEndpoint** | Pointer to **string** |  | [optional] 
**TokenEndpoint** | Pointer to **string** |  | [optional] 
**UserinfoEndpoint** | Pointer to **string** |  | [optional] 
**JwksUri** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Keys** | Pointer to [**IdentityOIDCSettingsKeys**](IdentityOIDCSettingsKeys.md) |  | [optional] 

## Methods

### NewIdentitySettingsResponse1Response

`func NewIdentitySettingsResponse1Response() *IdentitySettingsResponse1Response`

NewIdentitySettingsResponse1Response instantiates a new IdentitySettingsResponse1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySettingsResponse1ResponseWithDefaults

`func NewIdentitySettingsResponse1ResponseWithDefaults() *IdentitySettingsResponse1Response`

NewIdentitySettingsResponse1ResponseWithDefaults instantiates a new IdentitySettingsResponse1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *IdentitySettingsResponse1Response) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IdentitySettingsResponse1Response) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IdentitySettingsResponse1Response) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IdentitySettingsResponse1Response) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *IdentitySettingsResponse1Response) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *IdentitySettingsResponse1Response) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *IdentitySettingsResponse1Response) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *IdentitySettingsResponse1Response) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEncrypt

`func (o *IdentitySettingsResponse1Response) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *IdentitySettingsResponse1Response) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *IdentitySettingsResponse1Response) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *IdentitySettingsResponse1Response) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptLdaps

`func (o *IdentitySettingsResponse1Response) GetEncryptLdaps() bool`

GetEncryptLdaps returns the EncryptLdaps field if non-nil, zero value otherwise.

### GetEncryptLdapsOk

`func (o *IdentitySettingsResponse1Response) GetEncryptLdapsOk() (*bool, bool)`

GetEncryptLdapsOk returns a tuple with the EncryptLdaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptLdaps

`func (o *IdentitySettingsResponse1Response) SetEncryptLdaps(v bool)`

SetEncryptLdaps sets EncryptLdaps field to given value.

### HasEncryptLdaps

`func (o *IdentitySettingsResponse1Response) HasEncryptLdaps() bool`

HasEncryptLdaps returns a boolean if a field has been set.

### GetEncryptAuth

`func (o *IdentitySettingsResponse1Response) GetEncryptAuth() bool`

GetEncryptAuth returns the EncryptAuth field if non-nil, zero value otherwise.

### GetEncryptAuthOk

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthOk() (*bool, bool)`

GetEncryptAuthOk returns a tuple with the EncryptAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuth

`func (o *IdentitySettingsResponse1Response) SetEncryptAuth(v bool)`

SetEncryptAuth sets EncryptAuth field to given value.

### HasEncryptAuth

`func (o *IdentitySettingsResponse1Response) HasEncryptAuth() bool`

HasEncryptAuth returns a boolean if a field has been set.

### GetEncryptAuthKey

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthKey() bool`

GetEncryptAuthKey returns the EncryptAuthKey field if non-nil, zero value otherwise.

### GetEncryptAuthKeyOk

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthKeyOk() (*bool, bool)`

GetEncryptAuthKeyOk returns a tuple with the EncryptAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKey

`func (o *IdentitySettingsResponse1Response) SetEncryptAuthKey(v bool)`

SetEncryptAuthKey sets EncryptAuthKey field to given value.

### HasEncryptAuthKey

`func (o *IdentitySettingsResponse1Response) HasEncryptAuthKey() bool`

HasEncryptAuthKey returns a boolean if a field has been set.

### GetEncryptAuthCert

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthCert() bool`

GetEncryptAuthCert returns the EncryptAuthCert field if non-nil, zero value otherwise.

### GetEncryptAuthCertOk

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthCertOk() (*bool, bool)`

GetEncryptAuthCertOk returns a tuple with the EncryptAuthCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCert

`func (o *IdentitySettingsResponse1Response) SetEncryptAuthCert(v bool)`

SetEncryptAuthCert sets EncryptAuthCert field to given value.

### HasEncryptAuthCert

`func (o *IdentitySettingsResponse1Response) HasEncryptAuthCert() bool`

HasEncryptAuthCert returns a boolean if a field has been set.

### GetEncryptVerifyCa

`func (o *IdentitySettingsResponse1Response) GetEncryptVerifyCa() bool`

GetEncryptVerifyCa returns the EncryptVerifyCa field if non-nil, zero value otherwise.

### GetEncryptVerifyCaOk

`func (o *IdentitySettingsResponse1Response) GetEncryptVerifyCaOk() (*bool, bool)`

GetEncryptVerifyCaOk returns a tuple with the EncryptVerifyCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptVerifyCa

`func (o *IdentitySettingsResponse1Response) SetEncryptVerifyCa(v bool)`

SetEncryptVerifyCa sets EncryptVerifyCa field to given value.

### HasEncryptVerifyCa

`func (o *IdentitySettingsResponse1Response) HasEncryptVerifyCa() bool`

HasEncryptVerifyCa returns a boolean if a field has been set.

### GetEncryptCaCert

`func (o *IdentitySettingsResponse1Response) GetEncryptCaCert() bool`

GetEncryptCaCert returns the EncryptCaCert field if non-nil, zero value otherwise.

### GetEncryptCaCertOk

`func (o *IdentitySettingsResponse1Response) GetEncryptCaCertOk() (*bool, bool)`

GetEncryptCaCertOk returns a tuple with the EncryptCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCert

`func (o *IdentitySettingsResponse1Response) SetEncryptCaCert(v bool)`

SetEncryptCaCert sets EncryptCaCert field to given value.

### HasEncryptCaCert

`func (o *IdentitySettingsResponse1Response) HasEncryptCaCert() bool`

HasEncryptCaCert returns a boolean if a field has been set.

### GetBinddn

`func (o *IdentitySettingsResponse1Response) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *IdentitySettingsResponse1Response) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *IdentitySettingsResponse1Response) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *IdentitySettingsResponse1Response) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.

### GetBindpw

`func (o *IdentitySettingsResponse1Response) GetBindpw() string`

GetBindpw returns the Bindpw field if non-nil, zero value otherwise.

### GetBindpwOk

`func (o *IdentitySettingsResponse1Response) GetBindpwOk() (*string, bool)`

GetBindpwOk returns a tuple with the Bindpw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpw

`func (o *IdentitySettingsResponse1Response) SetBindpw(v string)`

SetBindpw sets Bindpw field to given value.

### HasBindpw

`func (o *IdentitySettingsResponse1Response) HasBindpw() bool`

HasBindpw returns a boolean if a field has been set.

### GetEncryptAuthCertData

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthCertData() string`

GetEncryptAuthCertData returns the EncryptAuthCertData field if non-nil, zero value otherwise.

### GetEncryptAuthCertDataOk

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthCertDataOk() (*string, bool)`

GetEncryptAuthCertDataOk returns a tuple with the EncryptAuthCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCertData

`func (o *IdentitySettingsResponse1Response) SetEncryptAuthCertData(v string)`

SetEncryptAuthCertData sets EncryptAuthCertData field to given value.

### HasEncryptAuthCertData

`func (o *IdentitySettingsResponse1Response) HasEncryptAuthCertData() bool`

HasEncryptAuthCertData returns a boolean if a field has been set.

### GetEncryptAuthCertFilename

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthCertFilename() string`

GetEncryptAuthCertFilename returns the EncryptAuthCertFilename field if non-nil, zero value otherwise.

### GetEncryptAuthCertFilenameOk

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthCertFilenameOk() (*string, bool)`

GetEncryptAuthCertFilenameOk returns a tuple with the EncryptAuthCertFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthCertFilename

`func (o *IdentitySettingsResponse1Response) SetEncryptAuthCertFilename(v string)`

SetEncryptAuthCertFilename sets EncryptAuthCertFilename field to given value.

### HasEncryptAuthCertFilename

`func (o *IdentitySettingsResponse1Response) HasEncryptAuthCertFilename() bool`

HasEncryptAuthCertFilename returns a boolean if a field has been set.

### GetEncryptAuthKeyData

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthKeyData() string`

GetEncryptAuthKeyData returns the EncryptAuthKeyData field if non-nil, zero value otherwise.

### GetEncryptAuthKeyDataOk

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthKeyDataOk() (*string, bool)`

GetEncryptAuthKeyDataOk returns a tuple with the EncryptAuthKeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKeyData

`func (o *IdentitySettingsResponse1Response) SetEncryptAuthKeyData(v string)`

SetEncryptAuthKeyData sets EncryptAuthKeyData field to given value.

### HasEncryptAuthKeyData

`func (o *IdentitySettingsResponse1Response) HasEncryptAuthKeyData() bool`

HasEncryptAuthKeyData returns a boolean if a field has been set.

### GetEncryptAuthKeyFilename

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthKeyFilename() string`

GetEncryptAuthKeyFilename returns the EncryptAuthKeyFilename field if non-nil, zero value otherwise.

### GetEncryptAuthKeyFilenameOk

`func (o *IdentitySettingsResponse1Response) GetEncryptAuthKeyFilenameOk() (*string, bool)`

GetEncryptAuthKeyFilenameOk returns a tuple with the EncryptAuthKeyFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptAuthKeyFilename

`func (o *IdentitySettingsResponse1Response) SetEncryptAuthKeyFilename(v string)`

SetEncryptAuthKeyFilename sets EncryptAuthKeyFilename field to given value.

### HasEncryptAuthKeyFilename

`func (o *IdentitySettingsResponse1Response) HasEncryptAuthKeyFilename() bool`

HasEncryptAuthKeyFilename returns a boolean if a field has been set.

### GetEncryptCaCertData

`func (o *IdentitySettingsResponse1Response) GetEncryptCaCertData() string`

GetEncryptCaCertData returns the EncryptCaCertData field if non-nil, zero value otherwise.

### GetEncryptCaCertDataOk

`func (o *IdentitySettingsResponse1Response) GetEncryptCaCertDataOk() (*string, bool)`

GetEncryptCaCertDataOk returns a tuple with the EncryptCaCertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCertData

`func (o *IdentitySettingsResponse1Response) SetEncryptCaCertData(v string)`

SetEncryptCaCertData sets EncryptCaCertData field to given value.

### HasEncryptCaCertData

`func (o *IdentitySettingsResponse1Response) HasEncryptCaCertData() bool`

HasEncryptCaCertData returns a boolean if a field has been set.

### GetEncryptCaCertFilename

`func (o *IdentitySettingsResponse1Response) GetEncryptCaCertFilename() string`

GetEncryptCaCertFilename returns the EncryptCaCertFilename field if non-nil, zero value otherwise.

### GetEncryptCaCertFilenameOk

`func (o *IdentitySettingsResponse1Response) GetEncryptCaCertFilenameOk() (*string, bool)`

GetEncryptCaCertFilenameOk returns a tuple with the EncryptCaCertFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptCaCertFilename

`func (o *IdentitySettingsResponse1Response) SetEncryptCaCertFilename(v string)`

SetEncryptCaCertFilename sets EncryptCaCertFilename field to given value.

### HasEncryptCaCertFilename

`func (o *IdentitySettingsResponse1Response) HasEncryptCaCertFilename() bool`

HasEncryptCaCertFilename returns a boolean if a field has been set.

### GetUserBase

`func (o *IdentitySettingsResponse1Response) GetUserBase() string`

GetUserBase returns the UserBase field if non-nil, zero value otherwise.

### GetUserBaseOk

`func (o *IdentitySettingsResponse1Response) GetUserBaseOk() (*string, bool)`

GetUserBaseOk returns a tuple with the UserBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBase

`func (o *IdentitySettingsResponse1Response) SetUserBase(v string)`

SetUserBase sets UserBase field to given value.

### HasUserBase

`func (o *IdentitySettingsResponse1Response) HasUserBase() bool`

HasUserBase returns a boolean if a field has been set.

### GetUserIdAttribute

`func (o *IdentitySettingsResponse1Response) GetUserIdAttribute() string`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *IdentitySettingsResponse1Response) GetUserIdAttributeOk() (*string, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *IdentitySettingsResponse1Response) SetUserIdAttribute(v string)`

SetUserIdAttribute sets UserIdAttribute field to given value.

### HasUserIdAttribute

`func (o *IdentitySettingsResponse1Response) HasUserIdAttribute() bool`

HasUserIdAttribute returns a boolean if a field has been set.

### GetUserListFilter

`func (o *IdentitySettingsResponse1Response) GetUserListFilter() string`

GetUserListFilter returns the UserListFilter field if non-nil, zero value otherwise.

### GetUserListFilterOk

`func (o *IdentitySettingsResponse1Response) GetUserListFilterOk() (*string, bool)`

GetUserListFilterOk returns a tuple with the UserListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserListFilter

`func (o *IdentitySettingsResponse1Response) SetUserListFilter(v string)`

SetUserListFilter sets UserListFilter field to given value.

### HasUserListFilter

`func (o *IdentitySettingsResponse1Response) HasUserListFilter() bool`

HasUserListFilter returns a boolean if a field has been set.

### GetGroupBase

`func (o *IdentitySettingsResponse1Response) GetGroupBase() string`

GetGroupBase returns the GroupBase field if non-nil, zero value otherwise.

### GetGroupBaseOk

`func (o *IdentitySettingsResponse1Response) GetGroupBaseOk() (*string, bool)`

GetGroupBaseOk returns a tuple with the GroupBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBase

`func (o *IdentitySettingsResponse1Response) SetGroupBase(v string)`

SetGroupBase sets GroupBase field to given value.

### HasGroupBase

`func (o *IdentitySettingsResponse1Response) HasGroupBase() bool`

HasGroupBase returns a boolean if a field has been set.

### GetGroupIdAttribute

`func (o *IdentitySettingsResponse1Response) GetGroupIdAttribute() string`

GetGroupIdAttribute returns the GroupIdAttribute field if non-nil, zero value otherwise.

### GetGroupIdAttributeOk

`func (o *IdentitySettingsResponse1Response) GetGroupIdAttributeOk() (*string, bool)`

GetGroupIdAttributeOk returns a tuple with the GroupIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdAttribute

`func (o *IdentitySettingsResponse1Response) SetGroupIdAttribute(v string)`

SetGroupIdAttribute sets GroupIdAttribute field to given value.

### HasGroupIdAttribute

`func (o *IdentitySettingsResponse1Response) HasGroupIdAttribute() bool`

HasGroupIdAttribute returns a boolean if a field has been set.

### GetGroupListFilter

`func (o *IdentitySettingsResponse1Response) GetGroupListFilter() string`

GetGroupListFilter returns the GroupListFilter field if non-nil, zero value otherwise.

### GetGroupListFilterOk

`func (o *IdentitySettingsResponse1Response) GetGroupListFilterOk() (*string, bool)`

GetGroupListFilterOk returns a tuple with the GroupListFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupListFilter

`func (o *IdentitySettingsResponse1Response) SetGroupListFilter(v string)`

SetGroupListFilter sets GroupListFilter field to given value.

### HasGroupListFilter

`func (o *IdentitySettingsResponse1Response) HasGroupListFilter() bool`

HasGroupListFilter returns a boolean if a field has been set.

### GetGroupMemberAttribute

`func (o *IdentitySettingsResponse1Response) GetGroupMemberAttribute() string`

GetGroupMemberAttribute returns the GroupMemberAttribute field if non-nil, zero value otherwise.

### GetGroupMemberAttributeOk

`func (o *IdentitySettingsResponse1Response) GetGroupMemberAttributeOk() (*string, bool)`

GetGroupMemberAttributeOk returns a tuple with the GroupMemberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttribute

`func (o *IdentitySettingsResponse1Response) SetGroupMemberAttribute(v string)`

SetGroupMemberAttribute sets GroupMemberAttribute field to given value.

### HasGroupMemberAttribute

`func (o *IdentitySettingsResponse1Response) HasGroupMemberAttribute() bool`

HasGroupMemberAttribute returns a boolean if a field has been set.

### GetGroupMemberAttrFormat

`func (o *IdentitySettingsResponse1Response) GetGroupMemberAttrFormat() string`

GetGroupMemberAttrFormat returns the GroupMemberAttrFormat field if non-nil, zero value otherwise.

### GetGroupMemberAttrFormatOk

`func (o *IdentitySettingsResponse1Response) GetGroupMemberAttrFormatOk() (*string, bool)`

GetGroupMemberAttrFormatOk returns a tuple with the GroupMemberAttrFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMemberAttrFormat

`func (o *IdentitySettingsResponse1Response) SetGroupMemberAttrFormat(v string)`

SetGroupMemberAttrFormat sets GroupMemberAttrFormat field to given value.

### HasGroupMemberAttrFormat

`func (o *IdentitySettingsResponse1Response) HasGroupMemberAttrFormat() bool`

HasGroupMemberAttrFormat returns a boolean if a field has been set.

### GetGroupSearchScope

`func (o *IdentitySettingsResponse1Response) GetGroupSearchScope() string`

GetGroupSearchScope returns the GroupSearchScope field if non-nil, zero value otherwise.

### GetGroupSearchScopeOk

`func (o *IdentitySettingsResponse1Response) GetGroupSearchScopeOk() (*string, bool)`

GetGroupSearchScopeOk returns a tuple with the GroupSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchScope

`func (o *IdentitySettingsResponse1Response) SetGroupSearchScope(v string)`

SetGroupSearchScope sets GroupSearchScope field to given value.

### HasGroupSearchScope

`func (o *IdentitySettingsResponse1Response) HasGroupSearchScope() bool`

HasGroupSearchScope returns a boolean if a field has been set.

### GetOtp

`func (o *IdentitySettingsResponse1Response) GetOtp() bool`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *IdentitySettingsResponse1Response) GetOtpOk() (*bool, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *IdentitySettingsResponse1Response) SetOtp(v bool)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *IdentitySettingsResponse1Response) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetOtpUrl

`func (o *IdentitySettingsResponse1Response) GetOtpUrl() string`

GetOtpUrl returns the OtpUrl field if non-nil, zero value otherwise.

### GetOtpUrlOk

`func (o *IdentitySettingsResponse1Response) GetOtpUrlOk() (*string, bool)`

GetOtpUrlOk returns a tuple with the OtpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpUrl

`func (o *IdentitySettingsResponse1Response) SetOtpUrl(v string)`

SetOtpUrl sets OtpUrl field to given value.

### HasOtpUrl

`func (o *IdentitySettingsResponse1Response) HasOtpUrl() bool`

HasOtpUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentitySettingsResponse1Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentitySettingsResponse1Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentitySettingsResponse1Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IdentitySettingsResponse1Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetProvider

`func (o *IdentitySettingsResponse1Response) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *IdentitySettingsResponse1Response) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *IdentitySettingsResponse1Response) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *IdentitySettingsResponse1Response) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetIdentifier

`func (o *IdentitySettingsResponse1Response) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *IdentitySettingsResponse1Response) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *IdentitySettingsResponse1Response) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *IdentitySettingsResponse1Response) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetSecret

`func (o *IdentitySettingsResponse1Response) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IdentitySettingsResponse1Response) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IdentitySettingsResponse1Response) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IdentitySettingsResponse1Response) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRedirectHostname

`func (o *IdentitySettingsResponse1Response) GetRedirectHostname() string`

GetRedirectHostname returns the RedirectHostname field if non-nil, zero value otherwise.

### GetRedirectHostnameOk

`func (o *IdentitySettingsResponse1Response) GetRedirectHostnameOk() (*string, bool)`

GetRedirectHostnameOk returns a tuple with the RedirectHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHostname

`func (o *IdentitySettingsResponse1Response) SetRedirectHostname(v string)`

SetRedirectHostname sets RedirectHostname field to given value.

### HasRedirectHostname

`func (o *IdentitySettingsResponse1Response) HasRedirectHostname() bool`

HasRedirectHostname returns a boolean if a field has been set.

### GetAuthorizationEndpoint

`func (o *IdentitySettingsResponse1Response) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *IdentitySettingsResponse1Response) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *IdentitySettingsResponse1Response) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.

### HasAuthorizationEndpoint

`func (o *IdentitySettingsResponse1Response) HasAuthorizationEndpoint() bool`

HasAuthorizationEndpoint returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *IdentitySettingsResponse1Response) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *IdentitySettingsResponse1Response) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *IdentitySettingsResponse1Response) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *IdentitySettingsResponse1Response) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetUserinfoEndpoint

`func (o *IdentitySettingsResponse1Response) GetUserinfoEndpoint() string`

GetUserinfoEndpoint returns the UserinfoEndpoint field if non-nil, zero value otherwise.

### GetUserinfoEndpointOk

`func (o *IdentitySettingsResponse1Response) GetUserinfoEndpointOk() (*string, bool)`

GetUserinfoEndpointOk returns a tuple with the UserinfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserinfoEndpoint

`func (o *IdentitySettingsResponse1Response) SetUserinfoEndpoint(v string)`

SetUserinfoEndpoint sets UserinfoEndpoint field to given value.

### HasUserinfoEndpoint

`func (o *IdentitySettingsResponse1Response) HasUserinfoEndpoint() bool`

HasUserinfoEndpoint returns a boolean if a field has been set.

### GetJwksUri

`func (o *IdentitySettingsResponse1Response) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *IdentitySettingsResponse1Response) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *IdentitySettingsResponse1Response) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *IdentitySettingsResponse1Response) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetIssuer

`func (o *IdentitySettingsResponse1Response) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *IdentitySettingsResponse1Response) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *IdentitySettingsResponse1Response) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *IdentitySettingsResponse1Response) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetKeys

`func (o *IdentitySettingsResponse1Response) GetKeys() IdentityOIDCSettingsKeys`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *IdentitySettingsResponse1Response) GetKeysOk() (*IdentityOIDCSettingsKeys, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *IdentitySettingsResponse1Response) SetKeys(v IdentityOIDCSettingsKeys)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *IdentitySettingsResponse1Response) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


