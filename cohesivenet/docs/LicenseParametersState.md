# LicenseParametersState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**License** | Pointer to **string** |  | [optional] 
**Finalized** | Pointer to **bool** |  | [optional] 
**Parameters** | Pointer to [**LicenseParameters**](LicenseParameters.md) |  | [optional] 

## Methods

### NewLicenseParametersState

`func NewLicenseParametersState() *LicenseParametersState`

NewLicenseParametersState instantiates a new LicenseParametersState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseParametersStateWithDefaults

`func NewLicenseParametersStateWithDefaults() *LicenseParametersState`

NewLicenseParametersStateWithDefaults instantiates a new LicenseParametersState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicense

`func (o *LicenseParametersState) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *LicenseParametersState) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *LicenseParametersState) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *LicenseParametersState) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetFinalized

`func (o *LicenseParametersState) GetFinalized() bool`

GetFinalized returns the Finalized field if non-nil, zero value otherwise.

### GetFinalizedOk

`func (o *LicenseParametersState) GetFinalizedOk() (*bool, bool)`

GetFinalizedOk returns a tuple with the Finalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalized

`func (o *LicenseParametersState) SetFinalized(v bool)`

SetFinalized sets Finalized field to given value.

### HasFinalized

`func (o *LicenseParametersState) HasFinalized() bool`

HasFinalized returns a boolean if a field has been set.

### GetParameters

`func (o *LicenseParametersState) GetParameters() LicenseParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *LicenseParametersState) GetParametersOk() (*LicenseParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *LicenseParametersState) SetParameters(v LicenseParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *LicenseParametersState) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


