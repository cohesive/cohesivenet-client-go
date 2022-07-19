# VNS3Controller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int32** |  | [optional] 
**ManagerId** | Pointer to **int32** |  | [optional] 
**OverlayIpaddress** | Pointer to [**TopologyClientsInner**](TopologyClientsInner.md) |  | [optional] 

## Methods

### NewVNS3Controller

`func NewVNS3Controller() *VNS3Controller`

NewVNS3Controller instantiates a new VNS3Controller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVNS3ControllerWithDefaults

`func NewVNS3ControllerWithDefaults() *VNS3Controller`

NewVNS3ControllerWithDefaults instantiates a new VNS3Controller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *VNS3Controller) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *VNS3Controller) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *VNS3Controller) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *VNS3Controller) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetManagerId

`func (o *VNS3Controller) GetManagerId() int32`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *VNS3Controller) GetManagerIdOk() (*int32, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *VNS3Controller) SetManagerId(v int32)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *VNS3Controller) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### GetOverlayIpaddress

`func (o *VNS3Controller) GetOverlayIpaddress() TopologyClientsInner`

GetOverlayIpaddress returns the OverlayIpaddress field if non-nil, zero value otherwise.

### GetOverlayIpaddressOk

`func (o *VNS3Controller) GetOverlayIpaddressOk() (*TopologyClientsInner, bool)`

GetOverlayIpaddressOk returns a tuple with the OverlayIpaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayIpaddress

`func (o *VNS3Controller) SetOverlayIpaddress(v TopologyClientsInner)`

SetOverlayIpaddress sets OverlayIpaddress field to given value.

### HasOverlayIpaddress

`func (o *VNS3Controller) HasOverlayIpaddress() bool`

HasOverlayIpaddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


