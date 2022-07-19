# IpsecSystemDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThisEndpoint** | Pointer to [**IpsecLocalEndpoint**](IpsecLocalEndpoint.md) |  | [optional] 
**RemoteEndpoints** | Pointer to [**map[string]IpsecSystemDetailResponseRemoteEndpointsValue**](IpsecSystemDetailResponseRemoteEndpointsValue.md) |  | [optional] 

## Methods

### NewIpsecSystemDetailResponse

`func NewIpsecSystemDetailResponse() *IpsecSystemDetailResponse`

NewIpsecSystemDetailResponse instantiates a new IpsecSystemDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsecSystemDetailResponseWithDefaults

`func NewIpsecSystemDetailResponseWithDefaults() *IpsecSystemDetailResponse`

NewIpsecSystemDetailResponseWithDefaults instantiates a new IpsecSystemDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThisEndpoint

`func (o *IpsecSystemDetailResponse) GetThisEndpoint() IpsecLocalEndpoint`

GetThisEndpoint returns the ThisEndpoint field if non-nil, zero value otherwise.

### GetThisEndpointOk

`func (o *IpsecSystemDetailResponse) GetThisEndpointOk() (*IpsecLocalEndpoint, bool)`

GetThisEndpointOk returns a tuple with the ThisEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThisEndpoint

`func (o *IpsecSystemDetailResponse) SetThisEndpoint(v IpsecLocalEndpoint)`

SetThisEndpoint sets ThisEndpoint field to given value.

### HasThisEndpoint

`func (o *IpsecSystemDetailResponse) HasThisEndpoint() bool`

HasThisEndpoint returns a boolean if a field has been set.

### GetRemoteEndpoints

`func (o *IpsecSystemDetailResponse) GetRemoteEndpoints() map[string]IpsecSystemDetailResponseRemoteEndpointsValue`

GetRemoteEndpoints returns the RemoteEndpoints field if non-nil, zero value otherwise.

### GetRemoteEndpointsOk

`func (o *IpsecSystemDetailResponse) GetRemoteEndpointsOk() (*map[string]IpsecSystemDetailResponseRemoteEndpointsValue, bool)`

GetRemoteEndpointsOk returns a tuple with the RemoteEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndpoints

`func (o *IpsecSystemDetailResponse) SetRemoteEndpoints(v map[string]IpsecSystemDetailResponseRemoteEndpointsValue)`

SetRemoteEndpoints sets RemoteEndpoints field to given value.

### HasRemoteEndpoints

`func (o *IpsecSystemDetailResponse) HasRemoteEndpoints() bool`

HasRemoteEndpoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


