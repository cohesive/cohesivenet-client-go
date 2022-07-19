# AddClientpackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestedIps** | **string** | CSV of IP addresses to be used for new clientpacks | 

## Methods

### NewAddClientpackRequest

`func NewAddClientpackRequest(requestedIps string, ) *AddClientpackRequest`

NewAddClientpackRequest instantiates a new AddClientpackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClientpackRequestWithDefaults

`func NewAddClientpackRequestWithDefaults() *AddClientpackRequest`

NewAddClientpackRequestWithDefaults instantiates a new AddClientpackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestedIps

`func (o *AddClientpackRequest) GetRequestedIps() string`

GetRequestedIps returns the RequestedIps field if non-nil, zero value otherwise.

### GetRequestedIpsOk

`func (o *AddClientpackRequest) GetRequestedIpsOk() (*string, bool)`

GetRequestedIpsOk returns a tuple with the RequestedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedIps

`func (o *AddClientpackRequest) SetRequestedIps(v string)`

SetRequestedIps sets RequestedIps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


