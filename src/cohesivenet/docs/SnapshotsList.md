# SnapshotsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatestSnapshot** | Pointer to **string** | Name of the latest snapshot taken | [optional] 
**Snapshots** | Pointer to [**map[string]SnapshotsListSnapshotsValue**](SnapshotsListSnapshotsValue.md) |  | [optional] 

## Methods

### NewSnapshotsList

`func NewSnapshotsList() *SnapshotsList`

NewSnapshotsList instantiates a new SnapshotsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotsListWithDefaults

`func NewSnapshotsListWithDefaults() *SnapshotsList`

NewSnapshotsListWithDefaults instantiates a new SnapshotsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatestSnapshot

`func (o *SnapshotsList) GetLatestSnapshot() string`

GetLatestSnapshot returns the LatestSnapshot field if non-nil, zero value otherwise.

### GetLatestSnapshotOk

`func (o *SnapshotsList) GetLatestSnapshotOk() (*string, bool)`

GetLatestSnapshotOk returns a tuple with the LatestSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSnapshot

`func (o *SnapshotsList) SetLatestSnapshot(v string)`

SetLatestSnapshot sets LatestSnapshot field to given value.

### HasLatestSnapshot

`func (o *SnapshotsList) HasLatestSnapshot() bool`

HasLatestSnapshot returns a boolean if a field has been set.

### GetSnapshots

`func (o *SnapshotsList) GetSnapshots() map[string]SnapshotsListSnapshotsValue`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *SnapshotsList) GetSnapshotsOk() (*map[string]SnapshotsListSnapshotsValue, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *SnapshotsList) SetSnapshots(v map[string]SnapshotsListSnapshotsValue)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *SnapshotsList) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


