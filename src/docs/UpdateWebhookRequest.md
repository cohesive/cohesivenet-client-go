# UpdateWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Events** | Pointer to **[]string** |  | [optional] [default to ["tunnel_up","tunnel_down"]]
**Body** | Pointer to **string** | Serialized HTTP Post request body | [optional] 
**ValidateCert** | Pointer to **bool** |  | [optional] 
**CustomProperties** | Pointer to [**[]CreateWebookRequestCustomPropertiesInner**](CreateWebookRequestCustomPropertiesInner.md) |  | [optional] 
**Headers** | Pointer to [**[]CreateWebookRequestHeadersInner**](CreateWebookRequestHeadersInner.md) |  | [optional] 
**Parameters** | Pointer to [**[]CreateWebookRequestHeadersInner**](CreateWebookRequestHeadersInner.md) |  | [optional] 

## Methods

### NewUpdateWebhookRequest

`func NewUpdateWebhookRequest() *UpdateWebhookRequest`

NewUpdateWebhookRequest instantiates a new UpdateWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookRequestWithDefaults

`func NewUpdateWebhookRequestWithDefaults() *UpdateWebhookRequest`

NewUpdateWebhookRequestWithDefaults instantiates a new UpdateWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateWebhookRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateWebhookRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateWebhookRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateWebhookRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEvents

`func (o *UpdateWebhookRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UpdateWebhookRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UpdateWebhookRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *UpdateWebhookRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetBody

`func (o *UpdateWebhookRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *UpdateWebhookRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *UpdateWebhookRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *UpdateWebhookRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetValidateCert

`func (o *UpdateWebhookRequest) GetValidateCert() bool`

GetValidateCert returns the ValidateCert field if non-nil, zero value otherwise.

### GetValidateCertOk

`func (o *UpdateWebhookRequest) GetValidateCertOk() (*bool, bool)`

GetValidateCertOk returns a tuple with the ValidateCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateCert

`func (o *UpdateWebhookRequest) SetValidateCert(v bool)`

SetValidateCert sets ValidateCert field to given value.

### HasValidateCert

`func (o *UpdateWebhookRequest) HasValidateCert() bool`

HasValidateCert returns a boolean if a field has been set.

### GetCustomProperties

`func (o *UpdateWebhookRequest) GetCustomProperties() []CreateWebookRequestCustomPropertiesInner`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *UpdateWebhookRequest) GetCustomPropertiesOk() (*[]CreateWebookRequestCustomPropertiesInner, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *UpdateWebhookRequest) SetCustomProperties(v []CreateWebookRequestCustomPropertiesInner)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *UpdateWebhookRequest) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetHeaders

`func (o *UpdateWebhookRequest) GetHeaders() []CreateWebookRequestHeadersInner`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *UpdateWebhookRequest) GetHeadersOk() (*[]CreateWebookRequestHeadersInner, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *UpdateWebhookRequest) SetHeaders(v []CreateWebookRequestHeadersInner)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *UpdateWebhookRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetParameters

`func (o *UpdateWebhookRequest) GetParameters() []CreateWebookRequestHeadersInner`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *UpdateWebhookRequest) GetParametersOk() (*[]CreateWebookRequestHeadersInner, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *UpdateWebhookRequest) SetParameters(v []CreateWebookRequestHeadersInner)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *UpdateWebhookRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


