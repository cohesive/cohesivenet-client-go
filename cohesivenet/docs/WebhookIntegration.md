# WebhookIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ValidateCert** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**System** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Body** | Pointer to **NullableString** |  | [optional] 
**CustomProperties** | Pointer to [**[]WebhookCustomProperty**](WebhookCustomProperty.md) |  | [optional] 
**Headers** | Pointer to [**[]WebhookHeader**](WebhookHeader.md) |  | [optional] 
**Parameters** | Pointer to [**[]WebhookParameter**](WebhookParameter.md) |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] 

## Methods

### NewWebhookIntegration

`func NewWebhookIntegration() *WebhookIntegration`

NewWebhookIntegration instantiates a new WebhookIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookIntegrationWithDefaults

`func NewWebhookIntegrationWithDefaults() *WebhookIntegration`

NewWebhookIntegrationWithDefaults instantiates a new WebhookIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookIntegration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookIntegration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookIntegration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WebhookIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookIntegration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookIntegration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValidateCert

`func (o *WebhookIntegration) GetValidateCert() bool`

GetValidateCert returns the ValidateCert field if non-nil, zero value otherwise.

### GetValidateCertOk

`func (o *WebhookIntegration) GetValidateCertOk() (*bool, bool)`

GetValidateCertOk returns a tuple with the ValidateCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateCert

`func (o *WebhookIntegration) SetValidateCert(v bool)`

SetValidateCert sets ValidateCert field to given value.

### HasValidateCert

`func (o *WebhookIntegration) HasValidateCert() bool`

HasValidateCert returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookIntegration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookIntegration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookIntegration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhookIntegration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WebhookIntegration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookIntegration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookIntegration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WebhookIntegration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSystem

`func (o *WebhookIntegration) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *WebhookIntegration) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *WebhookIntegration) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *WebhookIntegration) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WebhookIntegration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *WebhookIntegration) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *WebhookIntegration) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetBody

`func (o *WebhookIntegration) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WebhookIntegration) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WebhookIntegration) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *WebhookIntegration) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *WebhookIntegration) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *WebhookIntegration) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetCustomProperties

`func (o *WebhookIntegration) GetCustomProperties() []WebhookCustomProperty`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *WebhookIntegration) GetCustomPropertiesOk() (*[]WebhookCustomProperty, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *WebhookIntegration) SetCustomProperties(v []WebhookCustomProperty)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *WebhookIntegration) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetHeaders

`func (o *WebhookIntegration) GetHeaders() []WebhookHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookIntegration) GetHeadersOk() (*[]WebhookHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookIntegration) SetHeaders(v []WebhookHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebhookIntegration) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetParameters

`func (o *WebhookIntegration) GetParameters() []WebhookParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WebhookIntegration) GetParametersOk() (*[]WebhookParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WebhookIntegration) SetParameters(v []WebhookParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *WebhookIntegration) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetEvents

`func (o *WebhookIntegration) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookIntegration) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookIntegration) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookIntegration) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


