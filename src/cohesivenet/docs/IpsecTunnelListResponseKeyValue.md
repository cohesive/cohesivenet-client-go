# IpsecTunnelListResponseKeyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**map[string]RuntimeStatusIpsecValue**](RuntimeStatusIpsecValue.md) | Ipsec tunnel details keyed by ID | [optional] 

## Methods

### NewIpsecTunnelListResponseKeyValue

`func NewIpsecTunnelListResponseKeyValue() *IpsecTunnelListResponseKeyValue`

NewIpsecTunnelListResponseKeyValue instantiates a new IpsecTunnelListResponseKeyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsecTunnelListResponseKeyValueWithDefaults

`func NewIpsecTunnelListResponseKeyValueWithDefaults() *IpsecTunnelListResponseKeyValue`

NewIpsecTunnelListResponseKeyValueWithDefaults instantiates a new IpsecTunnelListResponseKeyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *IpsecTunnelListResponseKeyValue) GetResponse() map[string]RuntimeStatusIpsecValue`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *IpsecTunnelListResponseKeyValue) GetResponseOk() (*map[string]RuntimeStatusIpsecValue, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *IpsecTunnelListResponseKeyValue) SetResponse(v map[string]RuntimeStatusIpsecValue)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *IpsecTunnelListResponseKeyValue) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


