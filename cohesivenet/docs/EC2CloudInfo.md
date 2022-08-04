# EC2CloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**AvailabilityZone** | Pointer to **string** |  | [optional] 
**RamdiskId** | Pointer to **NullableString** |  | [optional] 
**KernelId** | Pointer to **NullableString** |  | [optional] 
**PendingTime** | Pointer to **string** |  | [optional] 
**Architecture** | Pointer to **string** |  | [optional] 
**PrivateIp** | Pointer to **string** |  | [optional] 
**DevpayProductCodes** | Pointer to **NullableString** |  | [optional] 
**MarketplaceProductCodes** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**BillingProducts** | Pointer to **NullableString** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 

## Methods

### NewEC2CloudInfo

`func NewEC2CloudInfo() *EC2CloudInfo`

NewEC2CloudInfo instantiates a new EC2CloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEC2CloudInfoWithDefaults

`func NewEC2CloudInfoWithDefaults() *EC2CloudInfo`

NewEC2CloudInfoWithDefaults instantiates a new EC2CloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *EC2CloudInfo) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EC2CloudInfo) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EC2CloudInfo) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *EC2CloudInfo) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *EC2CloudInfo) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *EC2CloudInfo) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *EC2CloudInfo) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *EC2CloudInfo) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetRamdiskId

`func (o *EC2CloudInfo) GetRamdiskId() string`

GetRamdiskId returns the RamdiskId field if non-nil, zero value otherwise.

### GetRamdiskIdOk

`func (o *EC2CloudInfo) GetRamdiskIdOk() (*string, bool)`

GetRamdiskIdOk returns a tuple with the RamdiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRamdiskId

`func (o *EC2CloudInfo) SetRamdiskId(v string)`

SetRamdiskId sets RamdiskId field to given value.

### HasRamdiskId

`func (o *EC2CloudInfo) HasRamdiskId() bool`

HasRamdiskId returns a boolean if a field has been set.

### SetRamdiskIdNil

`func (o *EC2CloudInfo) SetRamdiskIdNil(b bool)`

 SetRamdiskIdNil sets the value for RamdiskId to be an explicit nil

### UnsetRamdiskId
`func (o *EC2CloudInfo) UnsetRamdiskId()`

UnsetRamdiskId ensures that no value is present for RamdiskId, not even an explicit nil
### GetKernelId

`func (o *EC2CloudInfo) GetKernelId() string`

GetKernelId returns the KernelId field if non-nil, zero value otherwise.

### GetKernelIdOk

`func (o *EC2CloudInfo) GetKernelIdOk() (*string, bool)`

GetKernelIdOk returns a tuple with the KernelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelId

`func (o *EC2CloudInfo) SetKernelId(v string)`

SetKernelId sets KernelId field to given value.

### HasKernelId

`func (o *EC2CloudInfo) HasKernelId() bool`

HasKernelId returns a boolean if a field has been set.

### SetKernelIdNil

`func (o *EC2CloudInfo) SetKernelIdNil(b bool)`

 SetKernelIdNil sets the value for KernelId to be an explicit nil

### UnsetKernelId
`func (o *EC2CloudInfo) UnsetKernelId()`

UnsetKernelId ensures that no value is present for KernelId, not even an explicit nil
### GetPendingTime

`func (o *EC2CloudInfo) GetPendingTime() string`

GetPendingTime returns the PendingTime field if non-nil, zero value otherwise.

### GetPendingTimeOk

`func (o *EC2CloudInfo) GetPendingTimeOk() (*string, bool)`

GetPendingTimeOk returns a tuple with the PendingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTime

`func (o *EC2CloudInfo) SetPendingTime(v string)`

SetPendingTime sets PendingTime field to given value.

### HasPendingTime

`func (o *EC2CloudInfo) HasPendingTime() bool`

HasPendingTime returns a boolean if a field has been set.

### GetArchitecture

`func (o *EC2CloudInfo) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *EC2CloudInfo) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *EC2CloudInfo) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *EC2CloudInfo) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetPrivateIp

`func (o *EC2CloudInfo) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *EC2CloudInfo) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *EC2CloudInfo) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *EC2CloudInfo) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetDevpayProductCodes

`func (o *EC2CloudInfo) GetDevpayProductCodes() string`

GetDevpayProductCodes returns the DevpayProductCodes field if non-nil, zero value otherwise.

### GetDevpayProductCodesOk

`func (o *EC2CloudInfo) GetDevpayProductCodesOk() (*string, bool)`

GetDevpayProductCodesOk returns a tuple with the DevpayProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevpayProductCodes

`func (o *EC2CloudInfo) SetDevpayProductCodes(v string)`

SetDevpayProductCodes sets DevpayProductCodes field to given value.

### HasDevpayProductCodes

`func (o *EC2CloudInfo) HasDevpayProductCodes() bool`

HasDevpayProductCodes returns a boolean if a field has been set.

### SetDevpayProductCodesNil

`func (o *EC2CloudInfo) SetDevpayProductCodesNil(b bool)`

 SetDevpayProductCodesNil sets the value for DevpayProductCodes to be an explicit nil

### UnsetDevpayProductCodes
`func (o *EC2CloudInfo) UnsetDevpayProductCodes()`

UnsetDevpayProductCodes ensures that no value is present for DevpayProductCodes, not even an explicit nil
### GetMarketplaceProductCodes

`func (o *EC2CloudInfo) GetMarketplaceProductCodes() string`

GetMarketplaceProductCodes returns the MarketplaceProductCodes field if non-nil, zero value otherwise.

### GetMarketplaceProductCodesOk

`func (o *EC2CloudInfo) GetMarketplaceProductCodesOk() (*string, bool)`

GetMarketplaceProductCodesOk returns a tuple with the MarketplaceProductCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceProductCodes

`func (o *EC2CloudInfo) SetMarketplaceProductCodes(v string)`

SetMarketplaceProductCodes sets MarketplaceProductCodes field to given value.

### HasMarketplaceProductCodes

`func (o *EC2CloudInfo) HasMarketplaceProductCodes() bool`

HasMarketplaceProductCodes returns a boolean if a field has been set.

### SetMarketplaceProductCodesNil

`func (o *EC2CloudInfo) SetMarketplaceProductCodesNil(b bool)`

 SetMarketplaceProductCodesNil sets the value for MarketplaceProductCodes to be an explicit nil

### UnsetMarketplaceProductCodes
`func (o *EC2CloudInfo) UnsetMarketplaceProductCodes()`

UnsetMarketplaceProductCodes ensures that no value is present for MarketplaceProductCodes, not even an explicit nil
### GetVersion

`func (o *EC2CloudInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EC2CloudInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EC2CloudInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EC2CloudInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRegion

`func (o *EC2CloudInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EC2CloudInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EC2CloudInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *EC2CloudInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetImageId

`func (o *EC2CloudInfo) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *EC2CloudInfo) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *EC2CloudInfo) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *EC2CloudInfo) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetBillingProducts

`func (o *EC2CloudInfo) GetBillingProducts() string`

GetBillingProducts returns the BillingProducts field if non-nil, zero value otherwise.

### GetBillingProductsOk

`func (o *EC2CloudInfo) GetBillingProductsOk() (*string, bool)`

GetBillingProductsOk returns a tuple with the BillingProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProducts

`func (o *EC2CloudInfo) SetBillingProducts(v string)`

SetBillingProducts sets BillingProducts field to given value.

### HasBillingProducts

`func (o *EC2CloudInfo) HasBillingProducts() bool`

HasBillingProducts returns a boolean if a field has been set.

### SetBillingProductsNil

`func (o *EC2CloudInfo) SetBillingProductsNil(b bool)`

 SetBillingProductsNil sets the value for BillingProducts to be an explicit nil

### UnsetBillingProducts
`func (o *EC2CloudInfo) UnsetBillingProducts()`

UnsetBillingProducts ensures that no value is present for BillingProducts, not even an explicit nil
### GetInstanceId

`func (o *EC2CloudInfo) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *EC2CloudInfo) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *EC2CloudInfo) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *EC2CloudInfo) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceType

`func (o *EC2CloudInfo) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *EC2CloudInfo) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *EC2CloudInfo) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *EC2CloudInfo) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


