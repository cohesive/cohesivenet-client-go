# LicenseParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to **string** |  | [optional] 
**Controllers** | Pointer to **[]string** | IP addresses of VNS3 controllers in topology | [optional] 
**Managers** | Pointer to **[]string** | IP addresses of VNS3 controllers in topology | [optional] 
**Clients** | Pointer to **[]string** | IP addresses of clients in topology | [optional] 
**Asns** | Pointer to **[]int32** | ASNs used by controllers in topology | [optional] 
**MyManagerVip** | Pointer to **string** |  | [optional] 

## Methods

### NewLicenseParameters

`func NewLicenseParameters() *LicenseParameters`

NewLicenseParameters instantiates a new LicenseParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseParametersWithDefaults

`func NewLicenseParametersWithDefaults() *LicenseParameters`

NewLicenseParametersWithDefaults instantiates a new LicenseParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *LicenseParameters) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *LicenseParameters) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *LicenseParameters) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *LicenseParameters) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetControllers

`func (o *LicenseParameters) GetControllers() []string`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *LicenseParameters) GetControllersOk() (*[]string, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *LicenseParameters) SetControllers(v []string)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *LicenseParameters) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetManagers

`func (o *LicenseParameters) GetManagers() []string`

GetManagers returns the Managers field if non-nil, zero value otherwise.

### GetManagersOk

`func (o *LicenseParameters) GetManagersOk() (*[]string, bool)`

GetManagersOk returns a tuple with the Managers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagers

`func (o *LicenseParameters) SetManagers(v []string)`

SetManagers sets Managers field to given value.

### HasManagers

`func (o *LicenseParameters) HasManagers() bool`

HasManagers returns a boolean if a field has been set.

### GetClients

`func (o *LicenseParameters) GetClients() []string`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *LicenseParameters) GetClientsOk() (*[]string, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *LicenseParameters) SetClients(v []string)`

SetClients sets Clients field to given value.

### HasClients

`func (o *LicenseParameters) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetAsns

`func (o *LicenseParameters) GetAsns() []int32`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *LicenseParameters) GetAsnsOk() (*[]int32, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *LicenseParameters) SetAsns(v []int32)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *LicenseParameters) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetMyManagerVip

`func (o *LicenseParameters) GetMyManagerVip() string`

GetMyManagerVip returns the MyManagerVip field if non-nil, zero value otherwise.

### GetMyManagerVipOk

`func (o *LicenseParameters) GetMyManagerVipOk() (*string, bool)`

GetMyManagerVipOk returns a tuple with the MyManagerVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyManagerVip

`func (o *LicenseParameters) SetMyManagerVip(v string)`

SetMyManagerVip sets MyManagerVip field to given value.

### HasMyManagerVip

`func (o *LicenseParameters) HasMyManagerVip() bool`

HasMyManagerVip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


