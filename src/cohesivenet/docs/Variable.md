# Variable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Variable** | Pointer to **string** | resolvable variable name | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** | value that the variable resolves to | [optional] 

## Methods

### NewVariable

`func NewVariable() *Variable`

NewVariable instantiates a new Variable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableWithDefaults

`func NewVariableWithDefaults() *Variable`

NewVariableWithDefaults instantiates a new Variable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariable

`func (o *Variable) GetVariable() string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *Variable) GetVariableOk() (*string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *Variable) SetVariable(v string)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *Variable) HasVariable() bool`

HasVariable returns a boolean if a field has been set.

### GetDescription

`func (o *Variable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Variable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Variable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Variable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetValue

`func (o *Variable) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Variable) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Variable) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Variable) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


