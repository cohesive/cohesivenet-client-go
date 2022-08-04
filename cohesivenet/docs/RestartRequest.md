# RestartRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restart** | **bool** | Restarts target system on server if true | 

## Methods

### NewRestartRequest

`func NewRestartRequest(restart bool, ) *RestartRequest`

NewRestartRequest instantiates a new RestartRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartRequestWithDefaults

`func NewRestartRequestWithDefaults() *RestartRequest`

NewRestartRequestWithDefaults instantiates a new RestartRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestart

`func (o *RestartRequest) GetRestart() bool`

GetRestart returns the Restart field if non-nil, zero value otherwise.

### GetRestartOk

`func (o *RestartRequest) GetRestartOk() (*bool, bool)`

GetRestartOk returns a tuple with the Restart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestart

`func (o *RestartRequest) SetRestart(v bool)`

SetRestart sets Restart field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


