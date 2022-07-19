# CreateAlertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Url** | **string** |  | 
**WebhookId** | **int32** |  | 
**Events** | Pointer to **[]string** |  | [optional] 
**WebhookName** | Pointer to **string** |  | [optional] 
**CustomProperties** | Pointer to [**[]CreateAlertRequestCustomPropertiesInner**](CreateAlertRequestCustomPropertiesInner.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCreateAlertRequest

`func NewCreateAlertRequest(name string, url string, webhookId int32, ) *CreateAlertRequest`

NewCreateAlertRequest instantiates a new CreateAlertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAlertRequestWithDefaults

`func NewCreateAlertRequestWithDefaults() *CreateAlertRequest`

NewCreateAlertRequestWithDefaults instantiates a new CreateAlertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAlertRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAlertRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAlertRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *CreateAlertRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateAlertRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateAlertRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWebhookId

`func (o *CreateAlertRequest) GetWebhookId() int32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *CreateAlertRequest) GetWebhookIdOk() (*int32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *CreateAlertRequest) SetWebhookId(v int32)`

SetWebhookId sets WebhookId field to given value.


### GetEvents

`func (o *CreateAlertRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateAlertRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateAlertRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *CreateAlertRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetWebhookName

`func (o *CreateAlertRequest) GetWebhookName() string`

GetWebhookName returns the WebhookName field if non-nil, zero value otherwise.

### GetWebhookNameOk

`func (o *CreateAlertRequest) GetWebhookNameOk() (*string, bool)`

GetWebhookNameOk returns a tuple with the WebhookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookName

`func (o *CreateAlertRequest) SetWebhookName(v string)`

SetWebhookName sets WebhookName field to given value.

### HasWebhookName

`func (o *CreateAlertRequest) HasWebhookName() bool`

HasWebhookName returns a boolean if a field has been set.

### GetCustomProperties

`func (o *CreateAlertRequest) GetCustomProperties() []CreateAlertRequestCustomPropertiesInner`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *CreateAlertRequest) GetCustomPropertiesOk() (*[]CreateAlertRequestCustomPropertiesInner, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *CreateAlertRequest) SetCustomProperties(v []CreateAlertRequestCustomPropertiesInner)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *CreateAlertRequest) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetEnabled

`func (o *CreateAlertRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateAlertRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateAlertRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CreateAlertRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


