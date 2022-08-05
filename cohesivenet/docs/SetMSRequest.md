# SetMSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | VNS3 Management system endpoint IP address | [optional] 
**AlertEnabled** | **bool** | Disable/Enable sending alerts to VNS3:ms | 

## Methods

### NewSetMSRequest

`func NewSetMSRequest(alertEnabled bool, ) *SetMSRequest`

NewSetMSRequest instantiates a new SetMSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetMSRequestWithDefaults

`func NewSetMSRequestWithDefaults() *SetMSRequest`

NewSetMSRequestWithDefaults instantiates a new SetMSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *SetMSRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SetMSRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SetMSRequest) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SetMSRequest) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAlertEnabled

`func (o *SetMSRequest) GetAlertEnabled() bool`

GetAlertEnabled returns the AlertEnabled field if non-nil, zero value otherwise.

### GetAlertEnabledOk

`func (o *SetMSRequest) GetAlertEnabledOk() (*bool, bool)`

GetAlertEnabledOk returns a tuple with the AlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEnabled

`func (o *SetMSRequest) SetAlertEnabled(v bool)`

SetAlertEnabled sets AlertEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


