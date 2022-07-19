# LicenseInitial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** | Features available such as eBGP, CloudWAN, Containers etc. | [optional] 
**Finalized** | Pointer to **bool** |  | [optional] 
**License** | Pointer to **string** | State of license, accepted, in-progress, failed | [optional] 
**LicensePresent** | Pointer to **bool** |  | [optional] 
**DefaultTopology** | Pointer to [**Topology**](Topology.md) |  | [optional] 

## Methods

### NewLicenseInitial

`func NewLicenseInitial() *LicenseInitial`

NewLicenseInitial instantiates a new LicenseInitial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseInitialWithDefaults

`func NewLicenseInitialWithDefaults() *LicenseInitial`

NewLicenseInitialWithDefaults instantiates a new LicenseInitial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *LicenseInitial) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *LicenseInitial) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *LicenseInitial) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *LicenseInitial) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetFinalized

`func (o *LicenseInitial) GetFinalized() bool`

GetFinalized returns the Finalized field if non-nil, zero value otherwise.

### GetFinalizedOk

`func (o *LicenseInitial) GetFinalizedOk() (*bool, bool)`

GetFinalizedOk returns a tuple with the Finalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalized

`func (o *LicenseInitial) SetFinalized(v bool)`

SetFinalized sets Finalized field to given value.

### HasFinalized

`func (o *LicenseInitial) HasFinalized() bool`

HasFinalized returns a boolean if a field has been set.

### GetLicense

`func (o *LicenseInitial) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *LicenseInitial) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *LicenseInitial) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *LicenseInitial) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetLicensePresent

`func (o *LicenseInitial) GetLicensePresent() bool`

GetLicensePresent returns the LicensePresent field if non-nil, zero value otherwise.

### GetLicensePresentOk

`func (o *LicenseInitial) GetLicensePresentOk() (*bool, bool)`

GetLicensePresentOk returns a tuple with the LicensePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePresent

`func (o *LicenseInitial) SetLicensePresent(v bool)`

SetLicensePresent sets LicensePresent field to given value.

### HasLicensePresent

`func (o *LicenseInitial) HasLicensePresent() bool`

HasLicensePresent returns a boolean if a field has been set.

### GetDefaultTopology

`func (o *LicenseInitial) GetDefaultTopology() Topology`

GetDefaultTopology returns the DefaultTopology field if non-nil, zero value otherwise.

### GetDefaultTopologyOk

`func (o *LicenseInitial) GetDefaultTopologyOk() (*Topology, bool)`

GetDefaultTopologyOk returns a tuple with the DefaultTopology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTopology

`func (o *LicenseInitial) SetDefaultTopology(v Topology)`

SetDefaultTopology sets DefaultTopology field to given value.

### HasDefaultTopology

`func (o *LicenseInitial) HasDefaultTopology() bool`

HasDefaultTopology returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


