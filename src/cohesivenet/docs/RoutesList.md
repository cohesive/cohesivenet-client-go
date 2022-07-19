# RoutesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Cidr** | Pointer to **string** |  | [optional] 
**Interface** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Advertise** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Metric** | Pointer to **int32** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Table** | Pointer to **string** |  | [optional] 
**RouteBasedInterface** | Pointer to **string** |  | [optional] 
**RouteBasedGateway** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** | message about state of route | [optional] 
**Tunnel** | Pointer to [**RuntimeStatusIpsecValue**](RuntimeStatusIpsecValue.md) |  | [optional] 
**TrafficPair** | Pointer to [**IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue**](IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue.md) |  | [optional] 

## Methods

### NewRoutesList

`func NewRoutesList() *RoutesList`

NewRoutesList instantiates a new RoutesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutesListWithDefaults

`func NewRoutesListWithDefaults() *RoutesList`

NewRoutesListWithDefaults instantiates a new RoutesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoutesList) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoutesList) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoutesList) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RoutesList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCidr

`func (o *RoutesList) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *RoutesList) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *RoutesList) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *RoutesList) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetInterface

`func (o *RoutesList) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *RoutesList) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *RoutesList) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *RoutesList) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetNetmask

`func (o *RoutesList) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *RoutesList) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *RoutesList) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *RoutesList) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetDescription

`func (o *RoutesList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutesList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutesList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutesList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdvertise

`func (o *RoutesList) GetAdvertise() bool`

GetAdvertise returns the Advertise field if non-nil, zero value otherwise.

### GetAdvertiseOk

`func (o *RoutesList) GetAdvertiseOk() (*bool, bool)`

GetAdvertiseOk returns a tuple with the Advertise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertise

`func (o *RoutesList) SetAdvertise(v bool)`

SetAdvertise sets Advertise field to given value.

### HasAdvertise

`func (o *RoutesList) HasAdvertise() bool`

HasAdvertise returns a boolean if a field has been set.

### GetEnabled

`func (o *RoutesList) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RoutesList) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RoutesList) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RoutesList) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEditable

`func (o *RoutesList) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *RoutesList) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *RoutesList) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *RoutesList) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetMetric

`func (o *RoutesList) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *RoutesList) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *RoutesList) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *RoutesList) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetGateway

`func (o *RoutesList) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *RoutesList) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *RoutesList) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *RoutesList) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetTable

`func (o *RoutesList) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *RoutesList) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *RoutesList) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *RoutesList) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetRouteBasedInterface

`func (o *RoutesList) GetRouteBasedInterface() string`

GetRouteBasedInterface returns the RouteBasedInterface field if non-nil, zero value otherwise.

### GetRouteBasedInterfaceOk

`func (o *RoutesList) GetRouteBasedInterfaceOk() (*string, bool)`

GetRouteBasedInterfaceOk returns a tuple with the RouteBasedInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedInterface

`func (o *RoutesList) SetRouteBasedInterface(v string)`

SetRouteBasedInterface sets RouteBasedInterface field to given value.

### HasRouteBasedInterface

`func (o *RoutesList) HasRouteBasedInterface() bool`

HasRouteBasedInterface returns a boolean if a field has been set.

### GetRouteBasedGateway

`func (o *RoutesList) GetRouteBasedGateway() string`

GetRouteBasedGateway returns the RouteBasedGateway field if non-nil, zero value otherwise.

### GetRouteBasedGatewayOk

`func (o *RoutesList) GetRouteBasedGatewayOk() (*string, bool)`

GetRouteBasedGatewayOk returns a tuple with the RouteBasedGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteBasedGateway

`func (o *RoutesList) SetRouteBasedGateway(v string)`

SetRouteBasedGateway sets RouteBasedGateway field to given value.

### HasRouteBasedGateway

`func (o *RoutesList) HasRouteBasedGateway() bool`

HasRouteBasedGateway returns a boolean if a field has been set.

### GetMsg

`func (o *RoutesList) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *RoutesList) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *RoutesList) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *RoutesList) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetTunnel

`func (o *RoutesList) GetTunnel() RuntimeStatusIpsecValue`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *RoutesList) GetTunnelOk() (*RuntimeStatusIpsecValue, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *RoutesList) SetTunnel(v RuntimeStatusIpsecValue)`

SetTunnel sets Tunnel field to given value.

### HasTunnel

`func (o *RoutesList) HasTunnel() bool`

HasTunnel returns a boolean if a field has been set.

### GetTrafficPair

`func (o *RoutesList) GetTrafficPair() IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue`

GetTrafficPair returns the TrafficPair field if non-nil, zero value otherwise.

### GetTrafficPairOk

`func (o *RoutesList) GetTrafficPairOk() (*IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue, bool)`

GetTrafficPairOk returns a tuple with the TrafficPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPair

`func (o *RoutesList) SetTrafficPair(v IpsecSystemDetailResponseRemoteEndpointsValueTrafficPairsValue)`

SetTrafficPair sets TrafficPair field to given value.

### HasTrafficPair

`func (o *RoutesList) HasTrafficPair() bool`

HasTrafficPair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


