# PutLicenseParametersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to **string** | Specifies the CIDR of the virtual network created for use with the VNS3 Manager | [optional] 
**Managers** | Pointer to **string** | Whitespace delimited address string in the subnet to use for the VNS3 Controllers&#39; virtual interfaces. | [optional] 
**Asns** | Pointer to **string** | Whitespace delimited string of ASNs (autonomous system numbers) corresponding to the order of the controllers | [optional] 
**Clients** | Pointer to **string** | Comma delimited, or hyphenated sequence of addresses for use as client addresses in the virtual network. | [optional] 
**MyManagerVip** | Pointer to **string** | IPAddress that must be allocated from the subnet, and be the same for all controllers | [optional] 
**Default** | Pointer to **bool** | Specifices whether to use defualt topology addressing as specified by the license | [optional] [default to false]

## Methods

### NewPutLicenseParametersRequest

`func NewPutLicenseParametersRequest() *PutLicenseParametersRequest`

NewPutLicenseParametersRequest instantiates a new PutLicenseParametersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLicenseParametersRequestWithDefaults

`func NewPutLicenseParametersRequestWithDefaults() *PutLicenseParametersRequest`

NewPutLicenseParametersRequestWithDefaults instantiates a new PutLicenseParametersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *PutLicenseParametersRequest) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *PutLicenseParametersRequest) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *PutLicenseParametersRequest) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *PutLicenseParametersRequest) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetManagers

`func (o *PutLicenseParametersRequest) GetManagers() string`

GetManagers returns the Managers field if non-nil, zero value otherwise.

### GetManagersOk

`func (o *PutLicenseParametersRequest) GetManagersOk() (*string, bool)`

GetManagersOk returns a tuple with the Managers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagers

`func (o *PutLicenseParametersRequest) SetManagers(v string)`

SetManagers sets Managers field to given value.

### HasManagers

`func (o *PutLicenseParametersRequest) HasManagers() bool`

HasManagers returns a boolean if a field has been set.

### GetAsns

`func (o *PutLicenseParametersRequest) GetAsns() string`

GetAsns returns the Asns field if non-nil, zero value otherwise.

### GetAsnsOk

`func (o *PutLicenseParametersRequest) GetAsnsOk() (*string, bool)`

GetAsnsOk returns a tuple with the Asns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsns

`func (o *PutLicenseParametersRequest) SetAsns(v string)`

SetAsns sets Asns field to given value.

### HasAsns

`func (o *PutLicenseParametersRequest) HasAsns() bool`

HasAsns returns a boolean if a field has been set.

### GetClients

`func (o *PutLicenseParametersRequest) GetClients() string`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *PutLicenseParametersRequest) GetClientsOk() (*string, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *PutLicenseParametersRequest) SetClients(v string)`

SetClients sets Clients field to given value.

### HasClients

`func (o *PutLicenseParametersRequest) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetMyManagerVip

`func (o *PutLicenseParametersRequest) GetMyManagerVip() string`

GetMyManagerVip returns the MyManagerVip field if non-nil, zero value otherwise.

### GetMyManagerVipOk

`func (o *PutLicenseParametersRequest) GetMyManagerVipOk() (*string, bool)`

GetMyManagerVipOk returns a tuple with the MyManagerVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyManagerVip

`func (o *PutLicenseParametersRequest) SetMyManagerVip(v string)`

SetMyManagerVip sets MyManagerVip field to given value.

### HasMyManagerVip

`func (o *PutLicenseParametersRequest) HasMyManagerVip() bool`

HasMyManagerVip returns a boolean if a field has been set.

### GetDefault

`func (o *PutLicenseParametersRequest) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PutLicenseParametersRequest) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PutLicenseParametersRequest) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PutLicenseParametersRequest) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


