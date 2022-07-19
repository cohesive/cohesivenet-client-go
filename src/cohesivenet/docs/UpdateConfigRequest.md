# UpdateConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopologyName** | Pointer to **string** | Specifies a text name to display at the top of the web ui and in the desc_config API response | [optional] 
**ControllerName** | Pointer to **string** | Specifies a text name for this controller | [optional] 
**NtpHosts** | Pointer to **string** | Single or space separated list of ntp server IPs or dns names.  Using this argument overwrites the existing Configuration.  | [optional] 

## Methods

### NewUpdateConfigRequest

`func NewUpdateConfigRequest() *UpdateConfigRequest`

NewUpdateConfigRequest instantiates a new UpdateConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConfigRequestWithDefaults

`func NewUpdateConfigRequestWithDefaults() *UpdateConfigRequest`

NewUpdateConfigRequestWithDefaults instantiates a new UpdateConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopologyName

`func (o *UpdateConfigRequest) GetTopologyName() string`

GetTopologyName returns the TopologyName field if non-nil, zero value otherwise.

### GetTopologyNameOk

`func (o *UpdateConfigRequest) GetTopologyNameOk() (*string, bool)`

GetTopologyNameOk returns a tuple with the TopologyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyName

`func (o *UpdateConfigRequest) SetTopologyName(v string)`

SetTopologyName sets TopologyName field to given value.

### HasTopologyName

`func (o *UpdateConfigRequest) HasTopologyName() bool`

HasTopologyName returns a boolean if a field has been set.

### GetControllerName

`func (o *UpdateConfigRequest) GetControllerName() string`

GetControllerName returns the ControllerName field if non-nil, zero value otherwise.

### GetControllerNameOk

`func (o *UpdateConfigRequest) GetControllerNameOk() (*string, bool)`

GetControllerNameOk returns a tuple with the ControllerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerName

`func (o *UpdateConfigRequest) SetControllerName(v string)`

SetControllerName sets ControllerName field to given value.

### HasControllerName

`func (o *UpdateConfigRequest) HasControllerName() bool`

HasControllerName returns a boolean if a field has been set.

### GetNtpHosts

`func (o *UpdateConfigRequest) GetNtpHosts() string`

GetNtpHosts returns the NtpHosts field if non-nil, zero value otherwise.

### GetNtpHostsOk

`func (o *UpdateConfigRequest) GetNtpHostsOk() (*string, bool)`

GetNtpHostsOk returns a tuple with the NtpHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpHosts

`func (o *UpdateConfigRequest) SetNtpHosts(v string)`

SetNtpHosts sets NtpHosts field to given value.

### HasNtpHosts

`func (o *UpdateConfigRequest) HasNtpHosts() bool`

HasNtpHosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


