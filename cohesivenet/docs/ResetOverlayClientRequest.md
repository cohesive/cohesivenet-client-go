# ResetOverlayClientRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the clientpack as returned by the \&quot;desc_clientpacks\&quot; call | 
**Disconnect** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewResetOverlayClientRequest

`func NewResetOverlayClientRequest(name string, ) *ResetOverlayClientRequest`

NewResetOverlayClientRequest instantiates a new ResetOverlayClientRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetOverlayClientRequestWithDefaults

`func NewResetOverlayClientRequestWithDefaults() *ResetOverlayClientRequest`

NewResetOverlayClientRequestWithDefaults instantiates a new ResetOverlayClientRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ResetOverlayClientRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResetOverlayClientRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResetOverlayClientRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDisconnect

`func (o *ResetOverlayClientRequest) GetDisconnect() bool`

GetDisconnect returns the Disconnect field if non-nil, zero value otherwise.

### GetDisconnectOk

`func (o *ResetOverlayClientRequest) GetDisconnectOk() (*bool, bool)`

GetDisconnectOk returns a tuple with the Disconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnect

`func (o *ResetOverlayClientRequest) SetDisconnect(v bool)`

SetDisconnect sets Disconnect field to given value.

### HasDisconnect

`func (o *ResetOverlayClientRequest) HasDisconnect() bool`

HasDisconnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


