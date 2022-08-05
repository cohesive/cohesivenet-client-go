# GREEndpointAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**InterfaceId** | Pointer to **int32** |  | [optional] 
**IpExternal** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InterfaceType** | Pointer to **string** | system or system_virtual | [optional] 
**Status** | Pointer to **string** | Availability of interface, Up or Down | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGREEndpointAllOf

`func NewGREEndpointAllOf() *GREEndpointAllOf`

NewGREEndpointAllOf instantiates a new GREEndpointAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGREEndpointAllOfWithDefaults

`func NewGREEndpointAllOfWithDefaults() *GREEndpointAllOf`

NewGREEndpointAllOfWithDefaults instantiates a new GREEndpointAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GREEndpointAllOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GREEndpointAllOf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GREEndpointAllOf) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GREEndpointAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *GREEndpointAllOf) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *GREEndpointAllOf) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *GREEndpointAllOf) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *GREEndpointAllOf) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetIpExternal

`func (o *GREEndpointAllOf) GetIpExternal() string`

GetIpExternal returns the IpExternal field if non-nil, zero value otherwise.

### GetIpExternalOk

`func (o *GREEndpointAllOf) GetIpExternalOk() (*string, bool)`

GetIpExternalOk returns a tuple with the IpExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpExternal

`func (o *GREEndpointAllOf) SetIpExternal(v string)`

SetIpExternal sets IpExternal field to given value.

### HasIpExternal

`func (o *GREEndpointAllOf) HasIpExternal() bool`

HasIpExternal returns a boolean if a field has been set.

### SetIpExternalNil

`func (o *GREEndpointAllOf) SetIpExternalNil(b bool)`

 SetIpExternalNil sets the value for IpExternal to be an explicit nil

### UnsetIpExternal
`func (o *GREEndpointAllOf) UnsetIpExternal()`

UnsetIpExternal ensures that no value is present for IpExternal, not even an explicit nil
### GetName

`func (o *GREEndpointAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GREEndpointAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GREEndpointAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GREEndpointAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *GREEndpointAllOf) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *GREEndpointAllOf) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *GREEndpointAllOf) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *GREEndpointAllOf) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetStatus

`func (o *GREEndpointAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GREEndpointAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GREEndpointAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GREEndpointAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GREEndpointAllOf) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GREEndpointAllOf) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GREEndpointAllOf) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GREEndpointAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


