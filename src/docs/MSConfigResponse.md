# MSConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | IP address of VNS3 Management Systems | [optional] 
**AlertEnabled** | Pointer to **bool** | Enable alerting to MS | [optional] 

## Methods

### NewMSConfigResponse

`func NewMSConfigResponse() *MSConfigResponse`

NewMSConfigResponse instantiates a new MSConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSConfigResponseWithDefaults

`func NewMSConfigResponseWithDefaults() *MSConfigResponse`

NewMSConfigResponseWithDefaults instantiates a new MSConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *MSConfigResponse) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *MSConfigResponse) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *MSConfigResponse) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *MSConfigResponse) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetAlertEnabled

`func (o *MSConfigResponse) GetAlertEnabled() bool`

GetAlertEnabled returns the AlertEnabled field if non-nil, zero value otherwise.

### GetAlertEnabledOk

`func (o *MSConfigResponse) GetAlertEnabledOk() (*bool, bool)`

GetAlertEnabledOk returns a tuple with the AlertEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEnabled

`func (o *MSConfigResponse) SetAlertEnabled(v bool)`

SetAlertEnabled sets AlertEnabled field to given value.

### HasAlertEnabled

`func (o *MSConfigResponse) HasAlertEnabled() bool`

HasAlertEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


