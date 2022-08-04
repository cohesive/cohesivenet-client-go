# UpdateServerSSLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cert** | **string** |  | 
**Key** | **string** |  | 

## Methods

### NewUpdateServerSSLRequest

`func NewUpdateServerSSLRequest(cert string, key string, ) *UpdateServerSSLRequest`

NewUpdateServerSSLRequest instantiates a new UpdateServerSSLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServerSSLRequestWithDefaults

`func NewUpdateServerSSLRequestWithDefaults() *UpdateServerSSLRequest`

NewUpdateServerSSLRequestWithDefaults instantiates a new UpdateServerSSLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCert

`func (o *UpdateServerSSLRequest) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *UpdateServerSSLRequest) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *UpdateServerSSLRequest) SetCert(v string)`

SetCert sets Cert field to given value.


### GetKey

`func (o *UpdateServerSSLRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateServerSSLRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateServerSSLRequest) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


