# AllocateContainerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Id of container if present | [optional] 
**ImageUuid** | Pointer to **string** | Container image from which to allocate container | [optional] 
**Name** | Pointer to **string** | Name for the allocated container | [optional] 
**Ipaddress** | Pointer to **string** | IP address to use for container | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Command** | Pointer to **string** | Command used to run container | [optional] 
**Environment** | Pointer to **string** | Environment string to pass to container | [optional] 

## Methods

### NewAllocateContainerRequest

`func NewAllocateContainerRequest() *AllocateContainerRequest`

NewAllocateContainerRequest instantiates a new AllocateContainerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocateContainerRequestWithDefaults

`func NewAllocateContainerRequestWithDefaults() *AllocateContainerRequest`

NewAllocateContainerRequestWithDefaults instantiates a new AllocateContainerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AllocateContainerRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AllocateContainerRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AllocateContainerRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AllocateContainerRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetImageUuid

`func (o *AllocateContainerRequest) GetImageUuid() string`

GetImageUuid returns the ImageUuid field if non-nil, zero value otherwise.

### GetImageUuidOk

`func (o *AllocateContainerRequest) GetImageUuidOk() (*string, bool)`

GetImageUuidOk returns a tuple with the ImageUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUuid

`func (o *AllocateContainerRequest) SetImageUuid(v string)`

SetImageUuid sets ImageUuid field to given value.

### HasImageUuid

`func (o *AllocateContainerRequest) HasImageUuid() bool`

HasImageUuid returns a boolean if a field has been set.

### GetName

`func (o *AllocateContainerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AllocateContainerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AllocateContainerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AllocateContainerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIpaddress

`func (o *AllocateContainerRequest) GetIpaddress() string`

GetIpaddress returns the Ipaddress field if non-nil, zero value otherwise.

### GetIpaddressOk

`func (o *AllocateContainerRequest) GetIpaddressOk() (*string, bool)`

GetIpaddressOk returns a tuple with the Ipaddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddress

`func (o *AllocateContainerRequest) SetIpaddress(v string)`

SetIpaddress sets Ipaddress field to given value.

### HasIpaddress

`func (o *AllocateContainerRequest) HasIpaddress() bool`

HasIpaddress returns a boolean if a field has been set.

### GetDescription

`func (o *AllocateContainerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AllocateContainerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AllocateContainerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AllocateContainerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCommand

`func (o *AllocateContainerRequest) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *AllocateContainerRequest) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *AllocateContainerRequest) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *AllocateContainerRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEnvironment

`func (o *AllocateContainerRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AllocateContainerRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AllocateContainerRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AllocateContainerRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


