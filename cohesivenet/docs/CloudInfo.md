# CloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | Pointer to **string** | ec2, gce, azure, hpcloud, centurylink | [optional] 
**CloudData** | Pointer to [**CloudInfoCloudData**](CloudInfoCloudData.md) |  | [optional] 

## Methods

### NewCloudInfo

`func NewCloudInfo() *CloudInfo`

NewCloudInfo instantiates a new CloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudInfoWithDefaults

`func NewCloudInfoWithDefaults() *CloudInfo`

NewCloudInfoWithDefaults instantiates a new CloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudType

`func (o *CloudInfo) GetCloudType() string`

GetCloudType returns the CloudType field if non-nil, zero value otherwise.

### GetCloudTypeOk

`func (o *CloudInfo) GetCloudTypeOk() (*string, bool)`

GetCloudTypeOk returns a tuple with the CloudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudType

`func (o *CloudInfo) SetCloudType(v string)`

SetCloudType sets CloudType field to given value.

### HasCloudType

`func (o *CloudInfo) HasCloudType() bool`

HasCloudType returns a boolean if a field has been set.

### GetCloudData

`func (o *CloudInfo) GetCloudData() CloudInfoCloudData`

GetCloudData returns the CloudData field if non-nil, zero value otherwise.

### GetCloudDataOk

`func (o *CloudInfo) GetCloudDataOk() (*CloudInfoCloudData, bool)`

GetCloudDataOk returns a tuple with the CloudData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudData

`func (o *CloudInfo) SetCloudData(v CloudInfoCloudData)`

SetCloudData sets CloudData field to given value.

### HasCloudData

`func (o *CloudInfo) HasCloudData() bool`

HasCloudData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


