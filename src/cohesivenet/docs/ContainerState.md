# ContainerState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Running** | Pointer to **bool** |  | [optional] 
**Paused** | Pointer to **bool** |  | [optional] 
**Restarting** | Pointer to **bool** |  | [optional] 
**OOMKilled** | Pointer to **bool** |  | [optional] 
**Pid** | Pointer to **int32** |  | [optional] 
**ExitCode** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**FinishedAt** | Pointer to **string** |  | [optional] 
**Ghost** | Pointer to **bool** |  | [optional] 
**Dead** | Pointer to **bool** |  | [optional] 

## Methods

### NewContainerState

`func NewContainerState() *ContainerState`

NewContainerState instantiates a new ContainerState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerStateWithDefaults

`func NewContainerStateWithDefaults() *ContainerState`

NewContainerStateWithDefaults instantiates a new ContainerState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ContainerState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContainerState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContainerState) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContainerState) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRunning

`func (o *ContainerState) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *ContainerState) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *ContainerState) SetRunning(v bool)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *ContainerState) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetPaused

`func (o *ContainerState) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *ContainerState) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *ContainerState) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *ContainerState) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetRestarting

`func (o *ContainerState) GetRestarting() bool`

GetRestarting returns the Restarting field if non-nil, zero value otherwise.

### GetRestartingOk

`func (o *ContainerState) GetRestartingOk() (*bool, bool)`

GetRestartingOk returns a tuple with the Restarting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarting

`func (o *ContainerState) SetRestarting(v bool)`

SetRestarting sets Restarting field to given value.

### HasRestarting

`func (o *ContainerState) HasRestarting() bool`

HasRestarting returns a boolean if a field has been set.

### GetOOMKilled

`func (o *ContainerState) GetOOMKilled() bool`

GetOOMKilled returns the OOMKilled field if non-nil, zero value otherwise.

### GetOOMKilledOk

`func (o *ContainerState) GetOOMKilledOk() (*bool, bool)`

GetOOMKilledOk returns a tuple with the OOMKilled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOOMKilled

`func (o *ContainerState) SetOOMKilled(v bool)`

SetOOMKilled sets OOMKilled field to given value.

### HasOOMKilled

`func (o *ContainerState) HasOOMKilled() bool`

HasOOMKilled returns a boolean if a field has been set.

### GetPid

`func (o *ContainerState) GetPid() int32`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ContainerState) GetPidOk() (*int32, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ContainerState) SetPid(v int32)`

SetPid sets Pid field to given value.

### HasPid

`func (o *ContainerState) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetExitCode

`func (o *ContainerState) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ContainerState) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ContainerState) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *ContainerState) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetError

`func (o *ContainerState) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ContainerState) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ContainerState) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ContainerState) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartedAt

`func (o *ContainerState) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ContainerState) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ContainerState) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ContainerState) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *ContainerState) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ContainerState) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ContainerState) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ContainerState) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetGhost

`func (o *ContainerState) GetGhost() bool`

GetGhost returns the Ghost field if non-nil, zero value otherwise.

### GetGhostOk

`func (o *ContainerState) GetGhostOk() (*bool, bool)`

GetGhostOk returns a tuple with the Ghost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhost

`func (o *ContainerState) SetGhost(v bool)`

SetGhost sets Ghost field to given value.

### HasGhost

`func (o *ContainerState) HasGhost() bool`

HasGhost returns a boolean if a field has been set.

### GetDead

`func (o *ContainerState) GetDead() bool`

GetDead returns the Dead field if non-nil, zero value otherwise.

### GetDeadOk

`func (o *ContainerState) GetDeadOk() (*bool, bool)`

GetDeadOk returns a tuple with the Dead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDead

`func (o *ContainerState) SetDead(v bool)`

SetDead sets Dead field to given value.

### HasDead

`func (o *ContainerState) HasDead() bool`

HasDead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


