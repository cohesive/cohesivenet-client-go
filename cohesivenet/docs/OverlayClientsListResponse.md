# OverlayClientsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**map[string]RuntimeStatusConnectedClientsValue**](RuntimeStatusConnectedClientsValue.md) | Client details with IPs as keys | [optional] 

## Methods

### NewOverlayClientsListResponse

`func NewOverlayClientsListResponse() *OverlayClientsListResponse`

NewOverlayClientsListResponse instantiates a new OverlayClientsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverlayClientsListResponseWithDefaults

`func NewOverlayClientsListResponseWithDefaults() *OverlayClientsListResponse`

NewOverlayClientsListResponseWithDefaults instantiates a new OverlayClientsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *OverlayClientsListResponse) GetResponse() map[string]RuntimeStatusConnectedClientsValue`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *OverlayClientsListResponse) GetResponseOk() (*map[string]RuntimeStatusConnectedClientsValue, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *OverlayClientsListResponse) SetResponse(v map[string]RuntimeStatusConnectedClientsValue)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *OverlayClientsListResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


