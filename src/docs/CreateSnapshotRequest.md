# CreateSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of file. Defaults to a timestamped name | [optional] 
**Async** | Pointer to **bool** | If true, will return a task id rather than wait for snapshot generation | [optional] 

## Methods

### NewCreateSnapshotRequest

`func NewCreateSnapshotRequest() *CreateSnapshotRequest`

NewCreateSnapshotRequest instantiates a new CreateSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotRequestWithDefaults

`func NewCreateSnapshotRequestWithDefaults() *CreateSnapshotRequest`

NewCreateSnapshotRequestWithDefaults instantiates a new CreateSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSnapshotRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSnapshotRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSnapshotRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSnapshotRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAsync

`func (o *CreateSnapshotRequest) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *CreateSnapshotRequest) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *CreateSnapshotRequest) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *CreateSnapshotRequest) HasAsync() bool`

HasAsync returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


