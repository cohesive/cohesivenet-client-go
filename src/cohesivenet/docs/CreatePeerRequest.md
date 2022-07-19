# CreatePeerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Manager ID as an integer of the the manager you are peering with,  NOT the id of the one you are calling from  | 
**Name** | **string** | IP address or host name of the one you are peering with. | 
**OverlayMtu** | Pointer to **string** | link MTU between 500 and 4800. Defaults to 1500 | [optional] 
**Force** | Pointer to **bool** | Setting false will NOT finalize the peering operation.  A peer \&quot;reconfigure\&quot; call would then be required. Default is true  | [optional] 

## Methods

### NewCreatePeerRequest

`func NewCreatePeerRequest(id int32, name string, ) *CreatePeerRequest`

NewCreatePeerRequest instantiates a new CreatePeerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePeerRequestWithDefaults

`func NewCreatePeerRequestWithDefaults() *CreatePeerRequest`

NewCreatePeerRequestWithDefaults instantiates a new CreatePeerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreatePeerRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatePeerRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatePeerRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CreatePeerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePeerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePeerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOverlayMtu

`func (o *CreatePeerRequest) GetOverlayMtu() string`

GetOverlayMtu returns the OverlayMtu field if non-nil, zero value otherwise.

### GetOverlayMtuOk

`func (o *CreatePeerRequest) GetOverlayMtuOk() (*string, bool)`

GetOverlayMtuOk returns a tuple with the OverlayMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayMtu

`func (o *CreatePeerRequest) SetOverlayMtu(v string)`

SetOverlayMtu sets OverlayMtu field to given value.

### HasOverlayMtu

`func (o *CreatePeerRequest) HasOverlayMtu() bool`

HasOverlayMtu returns a boolean if a field has been set.

### GetForce

`func (o *CreatePeerRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *CreatePeerRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *CreatePeerRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *CreatePeerRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


