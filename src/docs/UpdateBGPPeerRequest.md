# UpdateBGPPeerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipaddress** | **string** | IP address of the desired BGP peer. | 
**Asn** | **int32** | Autonomous system number assigned to device at ipaddress | 
**LocalAsnAlias** | Pointer to **int32** | Autonomous system number alias | [optional] 
**AccessList** | Pointer to **string** | List of \&quot;in permit CIDR\&quot; and/or \&quot;out permit CIDR\&quot; statements in a string delimited by \&quot;\\n\&quot; | [optional] 
**BgpPassword** | Pointer to **string** | String to be agreed upon by both peers as a simple password. | [optional] 
**AddNetworkDistance** | Pointer to **bool** | Enable network distance for BGP peer | [optional] 
**AddNetworkDistanceDirection** | Pointer to **string** | Add distance direction for BGP peer, in or out | [optional] 
**AddNetworkDistanceHops** | Pointer to **int32** | Distance metric weight indicating distance in hops for BGP peer | [optional] 
**HoldTime** | Pointer to **int32** |  | [optional] 
**KeepaliveInterval** | Pointer to **int32** | Distance metric weight indicating distance in hops for BGP peer | [optional] 

## Methods

### NewUpdateBGPPeerRequest

`func NewUpdateBGPPeerRequest(ipaddress string, asn int32, ) *UpdateBGPPeerRequest`

NewUpdateBGPPeerRequest instantiates a new UpdateBGPPeerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBGPPeerRequestWithDefaults

`func NewUpdateBGPPeerRequestWithDefaults() *UpdateBGPPeerRequest`

NewUpdateBGPPeerRequestWithDefaults instantiates a new UpdateBGPPeerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpaddress

`func (o *UpdateBGPPeerRequest) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *UpdateBGPPeerRequest) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *UpdateBGPPeerRequest) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.


### GetAsn

`func (o *UpdateBGPPeerRequest) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *UpdateBGPPeerRequest) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *UpdateBGPPeerRequest) SetAsn(v int32)`

SetAsn sets Asn field to given value.


### GetLocalAsnAlias

`func (o *UpdateBGPPeerRequest) GetLocalAsnAlias() int32`

GetLocalAsnAlias returns the LocalAsnAlias field if non-nil, zero value otherwise.

### GetLocalAsnAliasOk

`func (o *UpdateBGPPeerRequest) GetLocalAsnAliasOk() (*int32, bool)`

GetLocalAsnAliasOk returns a tuple with the LocalAsnAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAsnAlias

`func (o *UpdateBGPPeerRequest) SetLocalAsnAlias(v int32)`

SetLocalAsnAlias sets LocalAsnAlias field to given value.

### HasLocalAsnAlias

`func (o *UpdateBGPPeerRequest) HasLocalAsnAlias() bool`

HasLocalAsnAlias returns a boolean if a field has been set.

### GetAccessList

`func (o *UpdateBGPPeerRequest) GetAccessList() string`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *UpdateBGPPeerRequest) GetAccessListOk() (*string, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *UpdateBGPPeerRequest) SetAccessList(v string)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *UpdateBGPPeerRequest) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetBgpPassword

`func (o *UpdateBGPPeerRequest) GetBgpPassword() string`

GetBgpPassword returns the BgpPassword field if non-nil, zero value otherwise.

### GetBgpPasswordOk

`func (o *UpdateBGPPeerRequest) GetBgpPasswordOk() (*string, bool)`

GetBgpPasswordOk returns a tuple with the BgpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpPassword

`func (o *UpdateBGPPeerRequest) SetBgpPassword(v string)`

SetBgpPassword sets BgpPassword field to given value.

### HasBgpPassword

`func (o *UpdateBGPPeerRequest) HasBgpPassword() bool`

HasBgpPassword returns a boolean if a field has been set.

### GetAddNetworkDistance

`func (o *UpdateBGPPeerRequest) GetAddNetworkDistance() bool`

GetAddNetworkDistance returns the AddNetworkDistance field if non-nil, zero value otherwise.

### GetAddNetworkDistanceOk

`func (o *UpdateBGPPeerRequest) GetAddNetworkDistanceOk() (*bool, bool)`

GetAddNetworkDistanceOk returns a tuple with the AddNetworkDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNetworkDistance

`func (o *UpdateBGPPeerRequest) SetAddNetworkDistance(v bool)`

SetAddNetworkDistance sets AddNetworkDistance field to given value.

### HasAddNetworkDistance

`func (o *UpdateBGPPeerRequest) HasAddNetworkDistance() bool`

HasAddNetworkDistance returns a boolean if a field has been set.

### GetAddNetworkDistanceDirection

`func (o *UpdateBGPPeerRequest) GetAddNetworkDistanceDirection() string`

GetAddNetworkDistanceDirection returns the AddNetworkDistanceDirection field if non-nil, zero value otherwise.

### GetAddNetworkDistanceDirectionOk

`func (o *UpdateBGPPeerRequest) GetAddNetworkDistanceDirectionOk() (*string, bool)`

GetAddNetworkDistanceDirectionOk returns a tuple with the AddNetworkDistanceDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNetworkDistanceDirection

`func (o *UpdateBGPPeerRequest) SetAddNetworkDistanceDirection(v string)`

SetAddNetworkDistanceDirection sets AddNetworkDistanceDirection field to given value.

### HasAddNetworkDistanceDirection

`func (o *UpdateBGPPeerRequest) HasAddNetworkDistanceDirection() bool`

HasAddNetworkDistanceDirection returns a boolean if a field has been set.

### GetAddNetworkDistanceHops

`func (o *UpdateBGPPeerRequest) GetAddNetworkDistanceHops() int32`

GetAddNetworkDistanceHops returns the AddNetworkDistanceHops field if non-nil, zero value otherwise.

### GetAddNetworkDistanceHopsOk

`func (o *UpdateBGPPeerRequest) GetAddNetworkDistanceHopsOk() (*int32, bool)`

GetAddNetworkDistanceHopsOk returns a tuple with the AddNetworkDistanceHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNetworkDistanceHops

`func (o *UpdateBGPPeerRequest) SetAddNetworkDistanceHops(v int32)`

SetAddNetworkDistanceHops sets AddNetworkDistanceHops field to given value.

### HasAddNetworkDistanceHops

`func (o *UpdateBGPPeerRequest) HasAddNetworkDistanceHops() bool`

HasAddNetworkDistanceHops returns a boolean if a field has been set.

### GetHoldTime

`func (o *UpdateBGPPeerRequest) GetHoldTime() int32`

GetHoldTime returns the HoldTime field if non-nil, zero value otherwise.

### GetHoldTimeOk

`func (o *UpdateBGPPeerRequest) GetHoldTimeOk() (*int32, bool)`

GetHoldTimeOk returns a tuple with the HoldTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldTime

`func (o *UpdateBGPPeerRequest) SetHoldTime(v int32)`

SetHoldTime sets HoldTime field to given value.

### HasHoldTime

`func (o *UpdateBGPPeerRequest) HasHoldTime() bool`

HasHoldTime returns a boolean if a field has been set.

### GetKeepaliveInterval

`func (o *UpdateBGPPeerRequest) GetKeepaliveInterval() int32`

GetKeepaliveInterval returns the KeepaliveInterval field if non-nil, zero value otherwise.

### GetKeepaliveIntervalOk

`func (o *UpdateBGPPeerRequest) GetKeepaliveIntervalOk() (*int32, bool)`

GetKeepaliveIntervalOk returns a tuple with the KeepaliveInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepaliveInterval

`func (o *UpdateBGPPeerRequest) SetKeepaliveInterval(v int32)`

SetKeepaliveInterval sets KeepaliveInterval field to given value.

### HasKeepaliveInterval

`func (o *UpdateBGPPeerRequest) HasKeepaliveInterval() bool`

HasKeepaliveInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


