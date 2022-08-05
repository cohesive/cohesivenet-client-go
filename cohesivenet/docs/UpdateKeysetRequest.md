# UpdateKeysetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | If provided, fetches keyset from source manager | [optional] 
**Token** | **string** | Arbitrary key used to help randomize the checksum, it must be identical for each manager in a topology. | 
**TopologyName** | Pointer to **string** | Name for the topology | [optional] 

## Methods

### NewUpdateKeysetRequest

`func NewUpdateKeysetRequest(token string, ) *UpdateKeysetRequest`

NewUpdateKeysetRequest instantiates a new UpdateKeysetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateKeysetRequestWithDefaults

`func NewUpdateKeysetRequestWithDefaults() *UpdateKeysetRequest`

NewUpdateKeysetRequestWithDefaults instantiates a new UpdateKeysetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *UpdateKeysetRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UpdateKeysetRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UpdateKeysetRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UpdateKeysetRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetToken

`func (o *UpdateKeysetRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateKeysetRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateKeysetRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetTopologyName

`func (o *UpdateKeysetRequest) GetTopologyName() string`

GetTopologyName returns the TopologyName field if non-nil, zero value otherwise.

### GetTopologyNameOk

`func (o *UpdateKeysetRequest) GetTopologyNameOk() (*string, bool)`

GetTopologyNameOk returns a tuple with the TopologyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyName

`func (o *UpdateKeysetRequest) SetTopologyName(v string)`

SetTopologyName sets TopologyName field to given value.

### HasTopologyName

`func (o *UpdateKeysetRequest) HasTopologyName() bool`

HasTopologyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


