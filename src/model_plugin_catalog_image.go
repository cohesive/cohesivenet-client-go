/*
VNS3 Controller API

Cohesive networks VNS3 provides complete control of your network's addressing, routes, rules and edge enabling a secure, connected and reactive cloud network. 

API version: 6.0.0
Contact: support@cohesive.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cohesivenet

import (
	"encoding/json"
)

// PluginCatalogImage struct for PluginCatalogImage
type PluginCatalogImage struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Documentation *string `json:"documentation,omitempty"`
	Support *string `json:"support,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	Categories []string `json:"categories,omitempty"`
	Tags *string `json:"tags,omitempty"`
	Logo *string `json:"logo,omitempty"`
	Keyphrases []string `json:"keyphrases,omitempty"`
	Version *string `json:"version,omitempty"`
	ProviderCode *string `json:"provider_code,omitempty"`
	Vns3Compatability *string `json:"vns3_compatability,omitempty"`
	Id *string `json:"id,omitempty"`
	Installed *bool `json:"installed,omitempty"`
	InstalledImageId *int32 `json:"installed_image_id,omitempty"`
}

// NewPluginCatalogImage instantiates a new PluginCatalogImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginCatalogImage() *PluginCatalogImage {
	this := PluginCatalogImage{}
	return &this
}

// NewPluginCatalogImageWithDefaults instantiates a new PluginCatalogImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginCatalogImageWithDefaults() *PluginCatalogImage {
	this := PluginCatalogImage{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PluginCatalogImage) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PluginCatalogImage) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetDocumentation() string {
	if o == nil || o.Documentation == nil {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetDocumentationOk() (*string, bool) {
	if o == nil || o.Documentation == nil {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasDocumentation() bool {
	if o != nil && o.Documentation != nil {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *PluginCatalogImage) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetSupport() string {
	if o == nil || o.Support == nil {
		var ret string
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetSupportOk() (*string, bool) {
	if o == nil || o.Support == nil {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasSupport() bool {
	if o != nil && o.Support != nil {
		return true
	}

	return false
}

// SetSupport gets a reference to the given string and assigns it to the Support field.
func (o *PluginCatalogImage) SetSupport(v string) {
	o.Support = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *PluginCatalogImage) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetCategories() []string {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetCategoriesOk() ([]string, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *PluginCatalogImage) SetCategories(v []string) {
	o.Categories = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *PluginCatalogImage) SetTags(v string) {
	o.Tags = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *PluginCatalogImage) SetLogo(v string) {
	o.Logo = &v
}

// GetKeyphrases returns the Keyphrases field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetKeyphrases() []string {
	if o == nil || o.Keyphrases == nil {
		var ret []string
		return ret
	}
	return o.Keyphrases
}

// GetKeyphrasesOk returns a tuple with the Keyphrases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetKeyphrasesOk() ([]string, bool) {
	if o == nil || o.Keyphrases == nil {
		return nil, false
	}
	return o.Keyphrases, true
}

// HasKeyphrases returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasKeyphrases() bool {
	if o != nil && o.Keyphrases != nil {
		return true
	}

	return false
}

// SetKeyphrases gets a reference to the given []string and assigns it to the Keyphrases field.
func (o *PluginCatalogImage) SetKeyphrases(v []string) {
	o.Keyphrases = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PluginCatalogImage) SetVersion(v string) {
	o.Version = &v
}

// GetProviderCode returns the ProviderCode field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetProviderCode() string {
	if o == nil || o.ProviderCode == nil {
		var ret string
		return ret
	}
	return *o.ProviderCode
}

// GetProviderCodeOk returns a tuple with the ProviderCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetProviderCodeOk() (*string, bool) {
	if o == nil || o.ProviderCode == nil {
		return nil, false
	}
	return o.ProviderCode, true
}

// HasProviderCode returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasProviderCode() bool {
	if o != nil && o.ProviderCode != nil {
		return true
	}

	return false
}

// SetProviderCode gets a reference to the given string and assigns it to the ProviderCode field.
func (o *PluginCatalogImage) SetProviderCode(v string) {
	o.ProviderCode = &v
}

// GetVns3Compatability returns the Vns3Compatability field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetVns3Compatability() string {
	if o == nil || o.Vns3Compatability == nil {
		var ret string
		return ret
	}
	return *o.Vns3Compatability
}

// GetVns3CompatabilityOk returns a tuple with the Vns3Compatability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetVns3CompatabilityOk() (*string, bool) {
	if o == nil || o.Vns3Compatability == nil {
		return nil, false
	}
	return o.Vns3Compatability, true
}

// HasVns3Compatability returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasVns3Compatability() bool {
	if o != nil && o.Vns3Compatability != nil {
		return true
	}

	return false
}

// SetVns3Compatability gets a reference to the given string and assigns it to the Vns3Compatability field.
func (o *PluginCatalogImage) SetVns3Compatability(v string) {
	o.Vns3Compatability = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PluginCatalogImage) SetId(v string) {
	o.Id = &v
}

// GetInstalled returns the Installed field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetInstalled() bool {
	if o == nil || o.Installed == nil {
		var ret bool
		return ret
	}
	return *o.Installed
}

// GetInstalledOk returns a tuple with the Installed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetInstalledOk() (*bool, bool) {
	if o == nil || o.Installed == nil {
		return nil, false
	}
	return o.Installed, true
}

// HasInstalled returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasInstalled() bool {
	if o != nil && o.Installed != nil {
		return true
	}

	return false
}

// SetInstalled gets a reference to the given bool and assigns it to the Installed field.
func (o *PluginCatalogImage) SetInstalled(v bool) {
	o.Installed = &v
}

// GetInstalledImageId returns the InstalledImageId field value if set, zero value otherwise.
func (o *PluginCatalogImage) GetInstalledImageId() int32 {
	if o == nil || o.InstalledImageId == nil {
		var ret int32
		return ret
	}
	return *o.InstalledImageId
}

// GetInstalledImageIdOk returns a tuple with the InstalledImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginCatalogImage) GetInstalledImageIdOk() (*int32, bool) {
	if o == nil || o.InstalledImageId == nil {
		return nil, false
	}
	return o.InstalledImageId, true
}

// HasInstalledImageId returns a boolean if a field has been set.
func (o *PluginCatalogImage) HasInstalledImageId() bool {
	if o != nil && o.InstalledImageId != nil {
		return true
	}

	return false
}

// SetInstalledImageId gets a reference to the given int32 and assigns it to the InstalledImageId field.
func (o *PluginCatalogImage) SetInstalledImageId(v int32) {
	o.InstalledImageId = &v
}

func (o PluginCatalogImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Documentation != nil {
		toSerialize["documentation"] = o.Documentation
	}
	if o.Support != nil {
		toSerialize["support"] = o.Support
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.Keyphrases != nil {
		toSerialize["keyphrases"] = o.Keyphrases
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.ProviderCode != nil {
		toSerialize["provider_code"] = o.ProviderCode
	}
	if o.Vns3Compatability != nil {
		toSerialize["vns3_compatability"] = o.Vns3Compatability
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Installed != nil {
		toSerialize["installed"] = o.Installed
	}
	if o.InstalledImageId != nil {
		toSerialize["installed_image_id"] = o.InstalledImageId
	}
	return json.Marshal(toSerialize)
}

type NullablePluginCatalogImage struct {
	value *PluginCatalogImage
	isSet bool
}

func (v NullablePluginCatalogImage) Get() *PluginCatalogImage {
	return v.value
}

func (v *NullablePluginCatalogImage) Set(val *PluginCatalogImage) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginCatalogImage) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginCatalogImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginCatalogImage(val *PluginCatalogImage) *NullablePluginCatalogImage {
	return &NullablePluginCatalogImage{value: val, isSet: true}
}

func (v NullablePluginCatalogImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginCatalogImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

