# SnapshotImportStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshot** | Pointer to **string** | Status of import | [optional] 
**LogLines** | Pointer to **[]string** | Log details if failed | [optional] 

## Methods

### NewSnapshotImportStatus

`func NewSnapshotImportStatus() *SnapshotImportStatus`

NewSnapshotImportStatus instantiates a new SnapshotImportStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotImportStatusWithDefaults

`func NewSnapshotImportStatusWithDefaults() *SnapshotImportStatus`

NewSnapshotImportStatusWithDefaults instantiates a new SnapshotImportStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshot

`func (o *SnapshotImportStatus) GetSnapshot() string`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *SnapshotImportStatus) GetSnapshotOk() (*string, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *SnapshotImportStatus) SetSnapshot(v string)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *SnapshotImportStatus) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetLogLines

`func (o *SnapshotImportStatus) GetLogLines() []string`

GetLogLines returns the LogLines field if non-nil, zero value otherwise.

### GetLogLinesOk

`func (o *SnapshotImportStatus) GetLogLinesOk() (*[]string, bool)`

GetLogLinesOk returns a tuple with the LogLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLines

`func (o *SnapshotImportStatus) SetLogLines(v []string)`

SetLogLines sets LogLines field to given value.

### HasLogLines

`func (o *SnapshotImportStatus) HasLogLines() bool`

HasLogLines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


