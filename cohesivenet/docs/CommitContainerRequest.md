# CommitContainerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **bool** | Name of new image | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewCommitContainerRequest

`func NewCommitContainerRequest(name bool, ) *CommitContainerRequest`

NewCommitContainerRequest instantiates a new CommitContainerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitContainerRequestWithDefaults

`func NewCommitContainerRequestWithDefaults() *CommitContainerRequest`

NewCommitContainerRequestWithDefaults instantiates a new CommitContainerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommitContainerRequest) GetName() bool`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommitContainerRequest) GetNameOk() (*bool, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommitContainerRequest) SetName(v bool)`

SetName sets Name field to given value.


### GetDescription

`func (o *CommitContainerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommitContainerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommitContainerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommitContainerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


