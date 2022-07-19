# CalculateNextClientpackRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LowIp** | Pointer to **string** | Set the lower bound for the resulting IP | [optional] 
**HighIp** | Pointer to **string** | Set the upper bound for the resulting IP | [optional] 
**IncludeDisabled** | Pointer to **bool** | Allows clientpack allocation from the disabled pool, for workflows where all clientpacks are disabled at the start. | [optional] [default to false]

## Methods

### NewCalculateNextClientpackRequest

`func NewCalculateNextClientpackRequest() *CalculateNextClientpackRequest`

NewCalculateNextClientpackRequest instantiates a new CalculateNextClientpackRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculateNextClientpackRequestWithDefaults

`func NewCalculateNextClientpackRequestWithDefaults() *CalculateNextClientpackRequest`

NewCalculateNextClientpackRequestWithDefaults instantiates a new CalculateNextClientpackRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLowIp

`func (o *CalculateNextClientpackRequest) GetLowIp() string`

GetLowIp returns the LowIp field if non-nil, zero value otherwise.

### GetLowIpOk

`func (o *CalculateNextClientpackRequest) GetLowIpOk() (*string, bool)`

GetLowIpOk returns a tuple with the LowIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowIp

`func (o *CalculateNextClientpackRequest) SetLowIp(v string)`

SetLowIp sets LowIp field to given value.

### HasLowIp

`func (o *CalculateNextClientpackRequest) HasLowIp() bool`

HasLowIp returns a boolean if a field has been set.

### GetHighIp

`func (o *CalculateNextClientpackRequest) GetHighIp() string`

GetHighIp returns the HighIp field if non-nil, zero value otherwise.

### GetHighIpOk

`func (o *CalculateNextClientpackRequest) GetHighIpOk() (*string, bool)`

GetHighIpOk returns a tuple with the HighIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighIp

`func (o *CalculateNextClientpackRequest) SetHighIp(v string)`

SetHighIp sets HighIp field to given value.

### HasHighIp

`func (o *CalculateNextClientpackRequest) HasHighIp() bool`

HasHighIp returns a boolean if a field has been set.

### GetIncludeDisabled

`func (o *CalculateNextClientpackRequest) GetIncludeDisabled() bool`

GetIncludeDisabled returns the IncludeDisabled field if non-nil, zero value otherwise.

### GetIncludeDisabledOk

`func (o *CalculateNextClientpackRequest) GetIncludeDisabledOk() (*bool, bool)`

GetIncludeDisabledOk returns a tuple with the IncludeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisabled

`func (o *CalculateNextClientpackRequest) SetIncludeDisabled(v bool)`

SetIncludeDisabled sets IncludeDisabled field to given value.

### HasIncludeDisabled

`func (o *CalculateNextClientpackRequest) HasIncludeDisabled() bool`

HasIncludeDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


