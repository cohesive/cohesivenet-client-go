# UpdateAlertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**WebhookId** | Pointer to **int32** |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 
**CustomProperties** | Pointer to [**[]CreateWebookRequestHeadersInner**](CreateWebookRequestHeadersInner.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewUpdateAlertRequest

`func NewUpdateAlertRequest() *UpdateAlertRequest`

NewUpdateAlertRequest instantiates a new UpdateAlertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAlertRequestWithDefaults

`func NewUpdateAlertRequestWithDefaults() *UpdateAlertRequest`

NewUpdateAlertRequestWithDefaults instantiates a new UpdateAlertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAlertRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAlertRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAlertRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAlertRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateAlertRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateAlertRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateAlertRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateAlertRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWebhookId

`func (o *UpdateAlertRequest) GetWebhookId() int32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *UpdateAlertRequest) GetWebhookIdOk() (*int32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *UpdateAlertRequest) SetWebhookId(v int32)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *UpdateAlertRequest) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetEvents

`func (o *UpdateAlertRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UpdateAlertRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UpdateAlertRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *UpdateAlertRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetCustomProperties

`func (o *UpdateAlertRequest) GetCustomProperties() []CreateWebookRequestHeadersInner`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *UpdateAlertRequest) GetCustomPropertiesOk() (*[]CreateWebookRequestHeadersInner, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *UpdateAlertRequest) SetCustomProperties(v []CreateWebookRequestHeadersInner)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *UpdateAlertRequest) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateAlertRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateAlertRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateAlertRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateAlertRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


