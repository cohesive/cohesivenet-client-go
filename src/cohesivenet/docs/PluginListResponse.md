# PluginListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**[]Plugin**](Plugin.md) |  | [optional] 

## Methods

### NewPluginListResponse

`func NewPluginListResponse() *PluginListResponse`

NewPluginListResponse instantiates a new PluginListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginListResponseWithDefaults

`func NewPluginListResponseWithDefaults() *PluginListResponse`

NewPluginListResponseWithDefaults instantiates a new PluginListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *PluginListResponse) GetResponse() []Plugin`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *PluginListResponse) GetResponseOk() (*[]Plugin, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *PluginListResponse) SetResponse(v []Plugin)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *PluginListResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


