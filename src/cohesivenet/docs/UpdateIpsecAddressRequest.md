# UpdateIpsecAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpsecLocalIpaddress** | **string** | This is effectively a \&quot;cloud NAT\&quot; address, since you don&#39;t know what your LAN address  will be between invocations in a cloud, this address can be used by remote endpoints  as your \&quot;behind a NAT\&quot; address, sometimes referred to Peer or IKE ID, if needed (e.g. Watchguard or Juniper). It can ALSO be thought of even more simply as an IPsec \&quot;loopback\&quot; interface that you can use to terminate traffic.  | 
**Async** | Pointer to **bool** | Return a task token waiting for IPsec configuration update, default is false, meaning the request will block | [optional] [default to false]

## Methods

### NewUpdateIpsecAddressRequest

`func NewUpdateIpsecAddressRequest(ipsecLocalIpaddress string, ) *UpdateIpsecAddressRequest`

NewUpdateIpsecAddressRequest instantiates a new UpdateIpsecAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIpsecAddressRequestWithDefaults

`func NewUpdateIpsecAddressRequestWithDefaults() *UpdateIpsecAddressRequest`

NewUpdateIpsecAddressRequestWithDefaults instantiates a new UpdateIpsecAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpsecLocalIpaddress

`func (o *UpdateIpsecAddressRequest) GetIpsecLocalIpaddress() string`

GetIpsecLocalIpaddress returns the IpsecLocalIpaddress field if non-nil, zero value otherwise.

### GetIpsecLocalIpaddressOk

`func (o *UpdateIpsecAddressRequest) GetIpsecLocalIpaddressOk() (*string, bool)`

GetIpsecLocalIpaddressOk returns a tuple with the IpsecLocalIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsecLocalIpaddress

`func (o *UpdateIpsecAddressRequest) SetIpsecLocalIpaddress(v string)`

SetIpsecLocalIpaddress sets IpsecLocalIpaddress field to given value.


### GetAsync

`func (o *UpdateIpsecAddressRequest) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *UpdateIpsecAddressRequest) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *UpdateIpsecAddressRequest) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *UpdateIpsecAddressRequest) HasAsync() bool`

HasAsync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


