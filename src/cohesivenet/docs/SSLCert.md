# SSLCert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Before** | Pointer to **string** |  | [optional] 
**After** | Pointer to **string** |  | [optional] 
**Algorithm** | Pointer to **string** |  | [optional] 
**Sha1Fingerprint** | Pointer to **string** |  | [optional] 
**Sha256Fingerprint** | Pointer to **string** |  | [optional] 

## Methods

### NewSSLCert

`func NewSSLCert() *SSLCert`

NewSSLCert instantiates a new SSLCert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLCertWithDefaults

`func NewSSLCertWithDefaults() *SSLCert`

NewSSLCertWithDefaults instantiates a new SSLCert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *SSLCert) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SSLCert) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SSLCert) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SSLCert) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetIssuer

`func (o *SSLCert) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SSLCert) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SSLCert) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SSLCert) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetBefore

`func (o *SSLCert) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *SSLCert) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *SSLCert) SetBefore(v string)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *SSLCert) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetAfter

`func (o *SSLCert) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *SSLCert) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *SSLCert) SetAfter(v string)`

SetAfter sets After field to given value.

### HasAfter

`func (o *SSLCert) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetAlgorithm

`func (o *SSLCert) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *SSLCert) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *SSLCert) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *SSLCert) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetSha1Fingerprint

`func (o *SSLCert) GetSha1Fingerprint() string`

GetSha1Fingerprint returns the Sha1Fingerprint field if non-nil, zero value otherwise.

### GetSha1FingerprintOk

`func (o *SSLCert) GetSha1FingerprintOk() (*string, bool)`

GetSha1FingerprintOk returns a tuple with the Sha1Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Fingerprint

`func (o *SSLCert) SetSha1Fingerprint(v string)`

SetSha1Fingerprint sets Sha1Fingerprint field to given value.

### HasSha1Fingerprint

`func (o *SSLCert) HasSha1Fingerprint() bool`

HasSha1Fingerprint returns a boolean if a field has been set.

### GetSha256Fingerprint

`func (o *SSLCert) GetSha256Fingerprint() string`

GetSha256Fingerprint returns the Sha256Fingerprint field if non-nil, zero value otherwise.

### GetSha256FingerprintOk

`func (o *SSLCert) GetSha256FingerprintOk() (*string, bool)`

GetSha256FingerprintOk returns a tuple with the Sha256Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Fingerprint

`func (o *SSLCert) SetSha256Fingerprint(v string)`

SetSha256Fingerprint sets Sha256Fingerprint field to given value.

### HasSha256Fingerprint

`func (o *SSLCert) HasSha256Fingerprint() bool`

HasSha256Fingerprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


