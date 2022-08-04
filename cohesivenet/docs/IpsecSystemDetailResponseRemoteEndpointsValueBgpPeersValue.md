# IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Asn** | Pointer to **int32** |  | [optional] 
**Ipaddress** | Pointer to **string** |  | [optional] 
**AccessList** | Pointer to **string** | List of \&quot;in permit CIDR\&quot; and/or \&quot;out permit CIDR\&quot; statements in a string delimited by \&quot;\\n\&quot; | [optional] 
**LocalAsnAlias** | Pointer to **int32** | Allow BGP configuration to use any ASN required by peer | [optional] 
**KeepaliveInterval** | Pointer to **int32** | Interval for checking if BGP peers are still alive | [optional] 
**HoldTime** | Pointer to **int32** | The length of inactive time after which BGP session is torn down. The timer is reset after updates and keepalives | [optional] 
**BgpPassword** | Pointer to **string** |  | [optional] 
**AddNetworkDistance** | Pointer to **bool** |  | [optional] 
**AddNetworkDistanceDirection** | Pointer to **string** | in or out | [optional] 
**AddNetworkDistanceHops** | Pointer to **int32** |  | [optional] 
**ConnectionDetail** | Pointer to **string** |  | [optional] 

## Methods

### NewIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue

`func NewIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue() *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue`

NewIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue instantiates a new IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValueWithDefaults

`func NewIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValueWithDefaults() *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue`

NewIpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValueWithDefaults instantiates a new IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAsn

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetAccessList

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAccessList() string`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAccessListOk() (*string, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetAccessList(v string)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetLocalAsnAlias

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetLocalAsnAlias() int32`

GetLocalAsnAlias returns the LocalAsnAlias field if non-nil, zero value otherwise.

### GetLocalAsnAliasOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetLocalAsnAliasOk() (*int32, bool)`

GetLocalAsnAliasOk returns a tuple with the LocalAsnAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsnAlias

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetLocalAsnAlias(v int32)`

SetLocalAsnAlias sets LocalAsnAlias field to given value.

### HasLocalAsnAlias

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasLocalAsnAlias() bool`

HasLocalAsnAlias returns a boolean if a field has been set.

### GetKeepaliveInterval

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetKeepaliveInterval() int32`

GetKeepaliveInterval returns the KeepaliveInterval field if non-nil, zero value otherwise.

### GetKeepaliveIntervalOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetKeepaliveIntervalOk() (*int32, bool)`

GetKeepaliveIntervalOk returns a tuple with the KeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveInterval

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetKeepaliveInterval(v int32)`

SetKeepaliveInterval sets KeepaliveInterval field to given value.

### HasKeepaliveInterval

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasKeepaliveInterval() bool`

HasKeepaliveInterval returns a boolean if a field has been set.

### GetHoldTime

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.

### GetBgpPassword

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetBgpPassword() string`

GetBgpPassword returns the BgpPassword field if non-nil, zero value otherwise.

### GetBgpPasswordOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetBgpPasswordOk() (*string, bool)`

GetBgpPasswordOk returns a tuple with the BgpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpPassword

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetBgpPassword(v string)`

SetBgpPassword sets BgpPassword field to given value.

### HasBgpPassword

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasBgpPassword() bool`

HasBgpPassword returns a boolean if a field has been set.

### GetAddNetworkDistance

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistance() bool`

GetAddNetworkDistance returns the AddNetworkDistance field if non-nil, zero value otherwise.

### GetAddNetworkDistanceOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistanceOk() (*bool, bool)`

GetAddNetworkDistanceOk returns a tuple with the AddNetworkDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNetworkDistance

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetAddNetworkDistance(v bool)`

SetAddNetworkDistance sets AddNetworkDistance field to given value.

### HasAddNetworkDistance

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasAddNetworkDistance() bool`

HasAddNetworkDistance returns a boolean if a field has been set.

### GetAddNetworkDistanceDirection

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistanceDirection() string`

GetAddNetworkDistanceDirection returns the AddNetworkDistanceDirection field if non-nil, zero value otherwise.

### GetAddNetworkDistanceDirectionOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistanceDirectionOk() (*string, bool)`

GetAddNetworkDistanceDirectionOk returns a tuple with the AddNetworkDistanceDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNetworkDistanceDirection

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetAddNetworkDistanceDirection(v string)`

SetAddNetworkDistanceDirection sets AddNetworkDistanceDirection field to given value.

### HasAddNetworkDistanceDirection

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasAddNetworkDistanceDirection() bool`

HasAddNetworkDistanceDirection returns a boolean if a field has been set.

### GetAddNetworkDistanceHops

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistanceHops() int32`

GetAddNetworkDistanceHops returns the AddNetworkDistanceHops field if non-nil, zero value otherwise.

### GetAddNetworkDistanceHopsOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetAddNetworkDistanceHopsOk() (*int32, bool)`

GetAddNetworkDistanceHopsOk returns a tuple with the AddNetworkDistanceHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNetworkDistanceHops

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetAddNetworkDistanceHops(v int32)`

SetAddNetworkDistanceHops sets AddNetworkDistanceHops field to given value.

### HasAddNetworkDistanceHops

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasAddNetworkDistanceHops() bool`

HasAddNetworkDistanceHops returns a boolean if a field has been set.

### GetConnectionDetail

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetConnectionDetail() string`

GetConnectionDetail returns the ConnectionDetail field if non-nil, zero value otherwise.

### GetConnectionDetailOk

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) GetConnectionDetailOk() (*string, bool)`

GetConnectionDetailOk returns a tuple with the ConnectionDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDetail

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) SetConnectionDetail(v string)`

SetConnectionDetail sets ConnectionDetail field to given value.

### HasConnectionDetail

`func (o *IpsecSystemDetailResponseRemoteEndpointsValueBgpPeersValue) HasConnectionDetail() bool`

HasConnectionDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


