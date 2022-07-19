# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** | Features available such as eBGP, CloudWAN etc. | [optional] 
**Finalized** | Pointer to **bool** |  | [optional] 
**MyManagerVip** | Pointer to **string** |  | [optional] 
**License** | Pointer to **string** | State of license, accepted, in-progress, failed | [optional] 
**LicensePresent** | Pointer to **bool** |  | [optional] 
**Sha1Checksum** | Pointer to **string** |  | [optional] 
**UploadedAt** | Pointer to **string** |  | [optional] 
**CustomAddressing** | Pointer to **bool** |  | [optional] 
**UploadedAtI** | Pointer to **int32** |  | [optional] 
**ContainerDetails** | Pointer to [**LicenseContainerDetails**](LicenseContainerDetails.md) |  | [optional] 
**Topology** | Pointer to [**Topology**](Topology.md) |  | [optional] 

## Methods

### NewLicense

`func NewLicense() *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseWithDefaults

`func NewLicenseWithDefaults() *License`

NewLicenseWithDefaults instantiates a new License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *License) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *License) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *License) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *License) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetFinalized

`func (o *License) GetFinalized() bool`

GetFinalized returns the Finalized field if non-nil, zero value otherwise.

### GetFinalizedOk

`func (o *License) GetFinalizedOk() (*bool, bool)`

GetFinalizedOk returns a tuple with the Finalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalized

`func (o *License) SetFinalized(v bool)`

SetFinalized sets Finalized field to given value.

### HasFinalized

`func (o *License) HasFinalized() bool`

HasFinalized returns a boolean if a field has been set.

### GetMyManagerVip

`func (o *License) GetMyManagerVip() string`

GetMyManagerVip returns the MyManagerVip field if non-nil, zero value otherwise.

### GetMyManagerVipOk

`func (o *License) GetMyManagerVipOk() (*string, bool)`

GetMyManagerVipOk returns a tuple with the MyManagerVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyManagerVip

`func (o *License) SetMyManagerVip(v string)`

SetMyManagerVip sets MyManagerVip field to given value.

### HasMyManagerVip

`func (o *License) HasMyManagerVip() bool`

HasMyManagerVip returns a boolean if a field has been set.

### GetLicense

`func (o *License) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *License) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *License) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *License) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetLicensePresent

`func (o *License) GetLicensePresent() bool`

GetLicensePresent returns the LicensePresent field if non-nil, zero value otherwise.

### GetLicensePresentOk

`func (o *License) GetLicensePresentOk() (*bool, bool)`

GetLicensePresentOk returns a tuple with the LicensePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePresent

`func (o *License) SetLicensePresent(v bool)`

SetLicensePresent sets LicensePresent field to given value.

### HasLicensePresent

`func (o *License) HasLicensePresent() bool`

HasLicensePresent returns a boolean if a field has been set.

### GetSha1Checksum

`func (o *License) GetSha1Checksum() string`

GetSha1Checksum returns the Sha1Checksum field if non-nil, zero value otherwise.

### GetSha1ChecksumOk

`func (o *License) GetSha1ChecksumOk() (*string, bool)`

GetSha1ChecksumOk returns a tuple with the Sha1Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha1Checksum

`func (o *License) SetSha1Checksum(v string)`

SetSha1Checksum sets Sha1Checksum field to given value.

### HasSha1Checksum

`func (o *License) HasSha1Checksum() bool`

HasSha1Checksum returns a boolean if a field has been set.

### GetUploadedAt

`func (o *License) GetUploadedAt() string`

GetUploadedAt returns the UploadedAt field if non-nil, zero value otherwise.

### GetUploadedAtOk

`func (o *License) GetUploadedAtOk() (*string, bool)`

GetUploadedAtOk returns a tuple with the UploadedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAt

`func (o *License) SetUploadedAt(v string)`

SetUploadedAt sets UploadedAt field to given value.

### HasUploadedAt

`func (o *License) HasUploadedAt() bool`

HasUploadedAt returns a boolean if a field has been set.

### GetCustomAddressing

`func (o *License) GetCustomAddressing() bool`

GetCustomAddressing returns the CustomAddressing field if non-nil, zero value otherwise.

### GetCustomAddressingOk

`func (o *License) GetCustomAddressingOk() (*bool, bool)`

GetCustomAddressingOk returns a tuple with the CustomAddressing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAddressing

`func (o *License) SetCustomAddressing(v bool)`

SetCustomAddressing sets CustomAddressing field to given value.

### HasCustomAddressing

`func (o *License) HasCustomAddressing() bool`

HasCustomAddressing returns a boolean if a field has been set.

### GetUploadedAtI

`func (o *License) GetUploadedAtI() int32`

GetUploadedAtI returns the UploadedAtI field if non-nil, zero value otherwise.

### GetUploadedAtIOk

`func (o *License) GetUploadedAtIOk() (*int32, bool)`

GetUploadedAtIOk returns a tuple with the UploadedAtI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedAtI

`func (o *License) SetUploadedAtI(v int32)`

SetUploadedAtI sets UploadedAtI field to given value.

### HasUploadedAtI

`func (o *License) HasUploadedAtI() bool`

HasUploadedAtI returns a boolean if a field has been set.

### GetContainerDetails

`func (o *License) GetContainerDetails() LicenseContainerDetails`

GetContainerDetails returns the ContainerDetails field if non-nil, zero value otherwise.

### GetContainerDetailsOk

`func (o *License) GetContainerDetailsOk() (*LicenseContainerDetails, bool)`

GetContainerDetailsOk returns a tuple with the ContainerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDetails

`func (o *License) SetContainerDetails(v LicenseContainerDetails)`

SetContainerDetails sets ContainerDetails field to given value.

### HasContainerDetails

`func (o *License) HasContainerDetails() bool`

HasContainerDetails returns a boolean if a field has been set.

### GetTopology

`func (o *License) GetTopology() Topology`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *License) GetTopologyOk() (*Topology, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *License) SetTopology(v Topology)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *License) HasTopology() bool`

HasTopology returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


