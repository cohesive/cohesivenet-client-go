# CreateContainerImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the image | [optional] 
**Url** | Pointer to **string** | URL of the image file to be imported | [optional] 
**Buildurl** | Pointer to **string** | URL of a dockerfile that will be used to build the image | [optional] 
**Localbuild** | Pointer to **string** | Local build file to create new image | [optional] 
**Localimage** | Pointer to **string** | Local image to tag | [optional] 
**Imagefile** | Pointer to ***os.File** | Upload image file | [optional] 
**Buildfile** | Pointer to ***os.File** | Upload docker file or zipped docker context directory | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateContainerImageRequest

`func NewCreateContainerImageRequest() *CreateContainerImageRequest`

NewCreateContainerImageRequest instantiates a new CreateContainerImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerImageRequestWithDefaults

`func NewCreateContainerImageRequestWithDefaults() *CreateContainerImageRequest`

NewCreateContainerImageRequestWithDefaults instantiates a new CreateContainerImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateContainerImageRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContainerImageRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContainerImageRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateContainerImageRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *CreateContainerImageRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateContainerImageRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateContainerImageRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateContainerImageRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBuildurl

`func (o *CreateContainerImageRequest) GetBuildurl() string`

GetBuildurl returns the Buildurl field if non-nil, zero value otherwise.

### GetBuildurlOk

`func (o *CreateContainerImageRequest) GetBuildurlOk() (*string, bool)`

GetBuildurlOk returns a tuple with the Buildurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildurl

`func (o *CreateContainerImageRequest) SetBuildurl(v string)`

SetBuildurl sets Buildurl field to given value.

### HasBuildurl

`func (o *CreateContainerImageRequest) HasBuildurl() bool`

HasBuildurl returns a boolean if a field has been set.

### GetLocalbuild

`func (o *CreateContainerImageRequest) GetLocalbuild() string`

GetLocalbuild returns the Localbuild field if non-nil, zero value otherwise.

### GetLocalbuildOk

`func (o *CreateContainerImageRequest) GetLocalbuildOk() (*string, bool)`

GetLocalbuildOk returns a tuple with the Localbuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalbuild

`func (o *CreateContainerImageRequest) SetLocalbuild(v string)`

SetLocalbuild sets Localbuild field to given value.

### HasLocalbuild

`func (o *CreateContainerImageRequest) HasLocalbuild() bool`

HasLocalbuild returns a boolean if a field has been set.

### GetLocalimage

`func (o *CreateContainerImageRequest) GetLocalimage() string`

GetLocalimage returns the Localimage field if non-nil, zero value otherwise.

### GetLocalimageOk

`func (o *CreateContainerImageRequest) GetLocalimageOk() (*string, bool)`

GetLocalimageOk returns a tuple with the Localimage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalimage

`func (o *CreateContainerImageRequest) SetLocalimage(v string)`

SetLocalimage sets Localimage field to given value.

### HasLocalimage

`func (o *CreateContainerImageRequest) HasLocalimage() bool`

HasLocalimage returns a boolean if a field has been set.

### GetImagefile

`func (o *CreateContainerImageRequest) GetImagefile() *os.File`

GetImagefile returns the Imagefile field if non-nil, zero value otherwise.

### GetImagefileOk

`func (o *CreateContainerImageRequest) GetImagefileOk() (**os.File, bool)`

GetImagefileOk returns a tuple with the Imagefile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagefile

`func (o *CreateContainerImageRequest) SetImagefile(v *os.File)`

SetImagefile sets Imagefile field to given value.

### HasImagefile

`func (o *CreateContainerImageRequest) HasImagefile() bool`

HasImagefile returns a boolean if a field has been set.

### GetBuildfile

`func (o *CreateContainerImageRequest) GetBuildfile() *os.File`

GetBuildfile returns the Buildfile field if non-nil, zero value otherwise.

### GetBuildfileOk

`func (o *CreateContainerImageRequest) GetBuildfileOk() (**os.File, bool)`

GetBuildfileOk returns a tuple with the Buildfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildfile

`func (o *CreateContainerImageRequest) SetBuildfile(v *os.File)`

SetBuildfile sets Buildfile field to given value.

### HasBuildfile

`func (o *CreateContainerImageRequest) HasBuildfile() bool`

HasBuildfile returns a boolean if a field has been set.

### GetDescription

`func (o *CreateContainerImageRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateContainerImageRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateContainerImageRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateContainerImageRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *CreateContainerImageRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateContainerImageRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateContainerImageRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateContainerImageRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


