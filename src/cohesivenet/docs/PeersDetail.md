# PeersDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Peered** | Pointer to **bool** |  | [optional] 
**Managers** | Pointer to [**map[string]PeersDetailManagersValue**](PeersDetailManagersValue.md) |  | [optional] 
**Controllers** | Pointer to [**map[string]PeersDetailManagersValue**](PeersDetailManagersValue.md) |  | [optional] 

## Methods

### NewPeersDetail

`func NewPeersDetail() *PeersDetail`

NewPeersDetail instantiates a new PeersDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPeersDetailWithDefaults

`func NewPeersDetailWithDefaults() *PeersDetail`

NewPeersDetailWithDefaults instantiates a new PeersDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PeersDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PeersDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PeersDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PeersDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPeered

`func (o *PeersDetail) GetPeered() bool`

GetPeered returns the Peered field if non-nil, zero value otherwise.

### GetPeeredOk

`func (o *PeersDetail) GetPeeredOk() (*bool, bool)`

GetPeeredOk returns a tuple with the Peered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeered

`func (o *PeersDetail) SetPeered(v bool)`

SetPeered sets Peered field to given value.

### HasPeered

`func (o *PeersDetail) HasPeered() bool`

HasPeered returns a boolean if a field has been set.

### GetManagers

`func (o *PeersDetail) GetManagers() map[string]PeersDetailManagersValue`

GetManagers returns the Managers field if non-nil, zero value otherwise.

### GetManagersOk

`func (o *PeersDetail) GetManagersOk() (*map[string]PeersDetailManagersValue, bool)`

GetManagersOk returns a tuple with the Managers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagers

`func (o *PeersDetail) SetManagers(v map[string]PeersDetailManagersValue)`

SetManagers sets Managers field to given value.

### HasManagers

`func (o *PeersDetail) HasManagers() bool`

HasManagers returns a boolean if a field has been set.

### GetControllers

`func (o *PeersDetail) GetControllers() map[string]PeersDetailManagersValue`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *PeersDetail) GetControllersOk() (*map[string]PeersDetailManagersValue, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *PeersDetail) SetControllers(v map[string]PeersDetailManagersValue)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *PeersDetail) HasControllers() bool`

HasControllers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


