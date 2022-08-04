# RunPluginInstanceExeCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | The command to run. (A key in the commands object) | 
**ExecutablePath** | Pointer to **string** | Path to the executable. This is required if more  than 1 executable is defined in plugin configuration.  | [optional] 
**Timeout** | Pointer to **int32** | Number of seconds to wait before timing out. | [optional] [default to 30]

## Methods

### NewRunPluginInstanceExeCommandRequest

`func NewRunPluginInstanceExeCommandRequest(command string, ) *RunPluginInstanceExeCommandRequest`

NewRunPluginInstanceExeCommandRequest instantiates a new RunPluginInstanceExeCommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunPluginInstanceExeCommandRequestWithDefaults

`func NewRunPluginInstanceExeCommandRequestWithDefaults() *RunPluginInstanceExeCommandRequest`

NewRunPluginInstanceExeCommandRequestWithDefaults instantiates a new RunPluginInstanceExeCommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *RunPluginInstanceExeCommandRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *RunPluginInstanceExeCommandRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *RunPluginInstanceExeCommandRequest) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetExecutablePath

`func (o *RunPluginInstanceExeCommandRequest) GetExecutablePath() string`

GetExecutablePath returns the ExecutablePath field if non-nil, zero value otherwise.

### GetExecutablePathOk

`func (o *RunPluginInstanceExeCommandRequest) GetExecutablePathOk() (*string, bool)`

GetExecutablePathOk returns a tuple with the ExecutablePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutablePath

`func (o *RunPluginInstanceExeCommandRequest) SetExecutablePath(v string)`

SetExecutablePath sets ExecutablePath field to given value.

### HasExecutablePath

`func (o *RunPluginInstanceExeCommandRequest) HasExecutablePath() bool`

HasExecutablePath returns a boolean if a field has been set.

### GetTimeout

`func (o *RunPluginInstanceExeCommandRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *RunPluginInstanceExeCommandRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *RunPluginInstanceExeCommandRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *RunPluginInstanceExeCommandRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


