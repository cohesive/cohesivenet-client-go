# ConnectedSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**CidrMask** | Pointer to **string** |  | [optional] 
**Managerid** | Pointer to **int32** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** | ipsec, local_no_encryption, remote_manager, or ebgp | [optional] 

## Methods

### NewConnectedSubnet

`func NewConnectedSubnet() *ConnectedSubnet`

NewConnectedSubnet instantiates a new ConnectedSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedSubnetWithDefaults

`func NewConnectedSubnetWithDefaults() *ConnectedSubnet`

NewConnectedSubnetWithDefaults instantiates a new ConnectedSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *ConnectedSubnet) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *ConnectedSubnet) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *ConnectedSubnet) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *ConnectedSubnet) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetNetwork

`func (o *ConnectedSubnet) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ConnectedSubnet) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ConnectedSubnet) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ConnectedSubnet) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCidrMask

`func (o *ConnectedSubnet) GetCidrMask() string`

GetCidrMask returns the CidrMask field if non-nil, zero value otherwise.

### GetCidrMaskOk

`func (o *ConnectedSubnet) GetCidrMaskOk() (*string, bool)`

GetCidrMaskOk returns a tuple with the CidrMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrMask

`func (o *ConnectedSubnet) SetCidrMask(v string)`

SetCidrMask sets CidrMask field to given value.

### HasCidrMask

`func (o *ConnectedSubnet) HasCidrMask() bool`

HasCidrMask returns a boolean if a field has been set.

### GetManagerid

`func (o *ConnectedSubnet) GetManagerid() int32`

GetManagerid returns the Managerid field if non-nil, zero value otherwise.

### GetManageridOk

`func (o *ConnectedSubnet) GetManageridOk() (*int32, bool)`

GetManageridOk returns a tuple with the Managerid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerid

`func (o *ConnectedSubnet) SetManagerid(v int32)`

SetManagerid sets Managerid field to given value.

### HasManagerid

`func (o *ConnectedSubnet) HasManagerid() bool`

HasManagerid returns a boolean if a field has been set.

### GetNetmask

`func (o *ConnectedSubnet) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *ConnectedSubnet) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *ConnectedSubnet) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *ConnectedSubnet) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetOrigin

`func (o *ConnectedSubnet) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ConnectedSubnet) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ConnectedSubnet) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ConnectedSubnet) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


