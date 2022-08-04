# UpdatePeerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | IP address or host name of the one you are peering with. | [optional] 
**OverlayMtu** | Pointer to **string** | link MTU between 500 and 4800 | [optional] 
**Force** | Pointer to **bool** | Setting false will NOT finalize the peering operation.  A peer \&quot;reconfigure\&quot; call would then be required. Default is true  | [optional] 

## Methods

### NewUpdatePeerRequest

`func NewUpdatePeerRequest() *UpdatePeerRequest`

NewUpdatePeerRequest instantiates a new UpdatePeerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePeerRequestWithDefaults

`func NewUpdatePeerRequestWithDefaults() *UpdatePeerRequest`

NewUpdatePeerRequestWithDefaults instantiates a new UpdatePeerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePeerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePeerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePeerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePeerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverlayMtu

`func (o *UpdatePeerRequest) GetOverlayMtu() string`

GetOverlayMtu returns the OverlayMtu field if non-nil, zero value otherwise.

### GetOverlayMtuOk

`func (o *UpdatePeerRequest) GetOverlayMtuOk() (*string, bool)`

GetOverlayMtuOk returns a tuple with the OverlayMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayMtu

`func (o *UpdatePeerRequest) SetOverlayMtu(v string)`

SetOverlayMtu sets OverlayMtu field to given value.

### HasOverlayMtu

`func (o *UpdatePeerRequest) HasOverlayMtu() bool`

HasOverlayMtu returns a boolean if a field has been set.

### GetForce

`func (o *UpdatePeerRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *UpdatePeerRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *UpdatePeerRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *UpdatePeerRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


