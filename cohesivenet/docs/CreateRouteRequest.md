# CreateRouteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** | CIDR of a route that the VNS3 Controller has access  to that it wants to publish throughout the  Routing tables of the overlay network  | 
**Description** | Pointer to **string** |  | [optional] 
**Interface** | Pointer to **string** | Sets the interface where this route will be advertised. | [optional] 
**Gateway** | Pointer to **string** | If interface is set, a specific gateway address reachable from that interface | [optional] 
**Table** | Pointer to **string** | Table to create route for | [optional] [default to "main"]
**Tunnel** | Pointer to **int32** | numerical reference for the GRE endpoint id (must provide either tunnel OR interface) | [optional] 
**Advertise** | Pointer to **bool** | advertise route to overlay network | [optional] 
**Metric** | Pointer to **int32** | weight for route | [optional] 

## Methods

### NewCreateRouteRequest

`func NewCreateRouteRequest(cidr string, ) *CreateRouteRequest`

NewCreateRouteRequest instantiates a new CreateRouteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteRequestWithDefaults

`func NewCreateRouteRequestWithDefaults() *CreateRouteRequest`

NewCreateRouteRequestWithDefaults instantiates a new CreateRouteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *CreateRouteRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateRouteRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateRouteRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetDescription

`func (o *CreateRouteRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRouteRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRouteRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRouteRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInterface

`func (o *CreateRouteRequest) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *CreateRouteRequest) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *CreateRouteRequest) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *CreateRouteRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetGateway

`func (o *CreateRouteRequest) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *CreateRouteRequest) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *CreateRouteRequest) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *CreateRouteRequest) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetTable

`func (o *CreateRouteRequest) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### GetTableOk

`func (o *CreateRouteRequest) GetTableOk() (*string, bool)`

GetTableOk returns a tuple with the Table field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTable

`func (o *CreateRouteRequest) SetTable(v string)`

SetTable sets Table field to given value.

### HasTable

`func (o *CreateRouteRequest) HasTable() bool`

HasTable returns a boolean if a field has been set.

### GetTunnel

`func (o *CreateRouteRequest) GetTunnel() int32`

GetTunnel returns the Tunnel field if non-nil, zero value otherwise.

### GetTunnelOk

`func (o *CreateRouteRequest) GetTunnelOk() (*int32, bool)`

GetTunnelOk returns a tuple with the Tunnel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnel

`func (o *CreateRouteRequest) SetTunnel(v int32)`

SetTunnel sets Tunnel field to given value.

### HasTunnel

`func (o *CreateRouteRequest) HasTunnel() bool`

HasTunnel returns a boolean if a field has been set.

### GetAdvertise

`func (o *CreateRouteRequest) GetAdvertise() bool`

GetAdvertise returns the Advertise field if non-nil, zero value otherwise.

### GetAdvertiseOk

`func (o *CreateRouteRequest) GetAdvertiseOk() (*bool, bool)`

GetAdvertiseOk returns a tuple with the Advertise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertise

`func (o *CreateRouteRequest) SetAdvertise(v bool)`

SetAdvertise sets Advertise field to given value.

### HasAdvertise

`func (o *CreateRouteRequest) HasAdvertise() bool`

HasAdvertise returns a boolean if a field has been set.

### GetMetric

`func (o *CreateRouteRequest) GetMetric() int32`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *CreateRouteRequest) GetMetricOk() (*int32, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *CreateRouteRequest) SetMetric(v int32)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *CreateRouteRequest) HasMetric() bool`

HasMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


