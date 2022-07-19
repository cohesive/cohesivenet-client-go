# DeleteAccessURLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessUrlId** | Pointer to **int32** | ID of access URL | [optional] 
**AccessUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewDeleteAccessURLRequest

`func NewDeleteAccessURLRequest() *DeleteAccessURLRequest`

NewDeleteAccessURLRequest instantiates a new DeleteAccessURLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAccessURLRequestWithDefaults

`func NewDeleteAccessURLRequestWithDefaults() *DeleteAccessURLRequest`

NewDeleteAccessURLRequestWithDefaults instantiates a new DeleteAccessURLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessUrlId

`func (o *DeleteAccessURLRequest) GetAccessUrlId() int32`

GetAccessUrlId returns the AccessUrlId field if non-nil, zero value otherwise.

### GetAccessUrlIdOk

`func (o *DeleteAccessURLRequest) GetAccessUrlIdOk() (*int32, bool)`

GetAccessUrlIdOk returns a tuple with the AccessUrlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessUrlId

`func (o *DeleteAccessURLRequest) SetAccessUrlId(v int32)`

SetAccessUrlId sets AccessUrlId field to given value.

### HasAccessUrlId

`func (o *DeleteAccessURLRequest) HasAccessUrlId() bool`

HasAccessUrlId returns a boolean if a field has been set.

### GetAccessUrl

`func (o *DeleteAccessURLRequest) GetAccessUrl() string`

GetAccessUrl returns the AccessUrl field if non-nil, zero value otherwise.

### GetAccessUrlOk

`func (o *DeleteAccessURLRequest) GetAccessUrlOk() (*string, bool)`

GetAccessUrlOk returns a tuple with the AccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessUrl

`func (o *DeleteAccessURLRequest) SetAccessUrl(v string)`

SetAccessUrl sets AccessUrl field to given value.

### HasAccessUrl

`func (o *DeleteAccessURLRequest) HasAccessUrl() bool`

HasAccessUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


