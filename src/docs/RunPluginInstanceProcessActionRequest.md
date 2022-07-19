# RunPluginInstanceProcessActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Process** | **string** | Name of the process. Should be listed in  subprocesses list of config.  | 
**Action** | **string** | Action to take. See documentation for supported actions for your process manager.  | 
**Timeout** | Pointer to **int32** | Number of seconds to wait before timing out. | [optional] [default to 20]

## Methods

### NewRunPluginInstanceProcessActionRequest

`func NewRunPluginInstanceProcessActionRequest(process string, action string, ) *RunPluginInstanceProcessActionRequest`

NewRunPluginInstanceProcessActionRequest instantiates a new RunPluginInstanceProcessActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunPluginInstanceProcessActionRequestWithDefaults

`func NewRunPluginInstanceProcessActionRequestWithDefaults() *RunPluginInstanceProcessActionRequest`

NewRunPluginInstanceProcessActionRequestWithDefaults instantiates a new RunPluginInstanceProcessActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcess

`func (o *RunPluginInstanceProcessActionRequest) GetProcess() string`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *RunPluginInstanceProcessActionRequest) GetProcessOk() (*string, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *RunPluginInstanceProcessActionRequest) SetProcess(v string)`

SetProcess sets Process field to given value.


### GetAction

`func (o *RunPluginInstanceProcessActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RunPluginInstanceProcessActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RunPluginInstanceProcessActionRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetTimeout

`func (o *RunPluginInstanceProcessActionRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *RunPluginInstanceProcessActionRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *RunPluginInstanceProcessActionRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *RunPluginInstanceProcessActionRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


