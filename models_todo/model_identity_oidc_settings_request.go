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

// IdentityOIDCSettingsRequest struct for IdentityOIDCSettingsRequest
type IdentityOIDCSettingsRequest struct {
	Enabled bool `json:"enabled"`
	Provider string `json:"provider"`
	Identifier *string `json:"identifier,omitempty"`
	Secret *string `json:"secret,omitempty"`
	RedirectHostname *string `json:"redirect_hostname,omitempty"`
	AuthorizationEndpoint *string `json:"authorization_endpoint,omitempty"`
	TokenEndpoint *string `json:"token_endpoint,omitempty"`
	UserinfoEndpoint *string `json:"userinfo_endpoint,omitempty"`
	JwksUri *string `json:"jwks_uri,omitempty"`
	OtpUrl *string `json:"otp_url,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	Keys *IdentityOIDCSettingsKeys `json:"keys,omitempty"`
}

// NewIdentityOIDCSettingsRequest instantiates a new IdentityOIDCSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityOIDCSettingsRequest(enabled bool, provider string) *IdentityOIDCSettingsRequest {
	this := IdentityOIDCSettingsRequest{}
	this.Enabled = enabled
	this.Provider = provider
	return &this
}

// NewIdentityOIDCSettingsRequestWithDefaults instantiates a new IdentityOIDCSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityOIDCSettingsRequestWithDefaults() *IdentityOIDCSettingsRequest {
	this := IdentityOIDCSettingsRequest{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *IdentityOIDCSettingsRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *IdentityOIDCSettingsRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetProvider returns the Provider field value
func (o *IdentityOIDCSettingsRequest) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *IdentityOIDCSettingsRequest) SetProvider(v string) {
	o.Provider = v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsRequest) GetIdentifier() string {
	if o == nil || o.Identifier == nil {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetIdentifierOk() (*string, bool) {
	if o == nil || o.Identifier == nil {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsRequest) HasIdentifier() bool {
	if o != nil && o.Identifier != nil {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *IdentityOIDCSettingsRequest) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *IdentityOIDCSettingsRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetRedirectHostname returns the RedirectHostname field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsRequest) GetRedirectHostname() string {
	if o == nil || o.RedirectHostname == nil {
		var ret string
		return ret
	}
	return *o.RedirectHostname
}

// GetRedirectHostnameOk returns a tuple with the RedirectHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetRedirectHostnameOk() (*string, bool) {
	if o == nil || o.RedirectHostname == nil {
		return nil, false
	}
	return o.RedirectHostname, true
}

// HasRedirectHostname returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsRequest) HasRedirectHostname() bool {
	if o != nil && o.RedirectHostname != nil {
		return true
	}

	return false
}

// SetRedirectHostname gets a reference to the given string and assigns it to the RedirectHostname field.
func (o *IdentityOIDCSettingsRequest) SetRedirectHostname(v string) {
	o.RedirectHostname = &v
}

// GetAuthorizationEndpoint returns the AuthorizationEndpoint field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsRequest) GetAuthorizationEndpoint() string {
	if o == nil || o.AuthorizationEndpoint == nil {
		var ret string
		return ret
	}
	return *o.AuthorizationEndpoint
}

// GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetAuthorizationEndpointOk() (*string, bool) {
	if o == nil || o.AuthorizationEndpoint == nil {
		return nil, false
	}
	return o.AuthorizationEndpoint, true
}

// HasAuthorizationEndpoint returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsRequest) HasAuthorizationEndpoint() bool {
	if o != nil && o.AuthorizationEndpoint != nil {
		return true
	}

	return false
}

// SetAuthorizationEndpoint gets a reference to the given string and assigns it to the AuthorizationEndpoint field.
func (o *IdentityOIDCSettingsRequest) SetAuthorizationEndpoint(v string) {
	o.AuthorizationEndpoint = &v
}

// GetTokenEndpoint returns the TokenEndpoint field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsRequest) GetTokenEndpoint() string {
	if o == nil || o.TokenEndpoint == nil {
		var ret string
		return ret
	}
	return *o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetTokenEndpointOk() (*string, bool) {
	if o == nil || o.TokenEndpoint == nil {
		return nil, false
	}
	return o.TokenEndpoint, true
}

// HasTokenEndpoint returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsRequest) HasTokenEndpoint() bool {
	if o != nil && o.TokenEndpoint != nil {
		return true
	}

	return false
}

// SetTokenEndpoint gets a reference to the given string and assigns it to the TokenEndpoint field.
func (o *IdentityOIDCSettingsRequest) SetTokenEndpoint(v string) {
	o.TokenEndpoint = &v
}

// GetUserinfoEndpoint returns the UserinfoEndpoint field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsRequest) GetUserinfoEndpoint() string {
	if o == nil || o.UserinfoEndpoint == nil {
		var ret string
		return ret
	}
	return *o.UserinfoEndpoint
}

// GetUserinfoEndpointOk returns a tuple with the UserinfoEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetUserinfoEndpointOk() (*string, bool) {
	if o == nil || o.UserinfoEndpoint == nil {
		return nil, false
	}
	return o.UserinfoEndpoint, true
}

// HasUserinfoEndpoint returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsRequest) HasUserinfoEndpoint() bool {
	if o != nil && o.UserinfoEndpoint != nil {
		return true
	}

	return false
}

// SetUserinfoEndpoint gets a reference to the given string and assigns it to the UserinfoEndpoint field.
func (o *IdentityOIDCSettingsRequest) SetUserinfoEndpoint(v string) {
	o.UserinfoEndpoint = &v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsRequest) GetJwksUri() string {
	if o == nil || o.JwksUri == nil {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetJwksUriOk() (*string, bool) {
	if o == nil || o.JwksUri == nil {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsRequest) HasJwksUri() bool {
	if o != nil && o.JwksUri != nil {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *IdentityOIDCSettingsRequest) SetJwksUri(v string) {
	o.JwksUri = &v
}

// GetOtpUrl returns the OtpUrl field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsRequest) GetOtpUrl() string {
	if o == nil || o.OtpUrl == nil {
		var ret string
		return ret
	}
	return *o.OtpUrl
}

// GetOtpUrlOk returns a tuple with the OtpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetOtpUrlOk() (*string, bool) {
	if o == nil || o.OtpUrl == nil {
		return nil, false
	}
	return o.OtpUrl, true
}

// HasOtpUrl returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsRequest) HasOtpUrl() bool {
	if o != nil && o.OtpUrl != nil {
		return true
	}

	return false
}

// SetOtpUrl gets a reference to the given string and assigns it to the OtpUrl field.
func (o *IdentityOIDCSettingsRequest) SetOtpUrl(v string) {
	o.OtpUrl = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsRequest) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsRequest) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *IdentityOIDCSettingsRequest) SetIssuer(v string) {
	o.Issuer = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *IdentityOIDCSettingsRequest) GetKeys() IdentityOIDCSettingsKeys {
	if o == nil || o.Keys == nil {
		var ret IdentityOIDCSettingsKeys
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityOIDCSettingsRequest) GetKeysOk() (*IdentityOIDCSettingsKeys, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *IdentityOIDCSettingsRequest) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given IdentityOIDCSettingsKeys and assigns it to the Keys field.
func (o *IdentityOIDCSettingsRequest) SetKeys(v IdentityOIDCSettingsKeys) {
	o.Keys = &v
}

func (o IdentityOIDCSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if o.Identifier != nil {
		toSerialize["identifier"] = o.Identifier
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.RedirectHostname != nil {
		toSerialize["redirect_hostname"] = o.RedirectHostname
	}
	if o.AuthorizationEndpoint != nil {
		toSerialize["authorization_endpoint"] = o.AuthorizationEndpoint
	}
	if o.TokenEndpoint != nil {
		toSerialize["token_endpoint"] = o.TokenEndpoint
	}
	if o.UserinfoEndpoint != nil {
		toSerialize["userinfo_endpoint"] = o.UserinfoEndpoint
	}
	if o.JwksUri != nil {
		toSerialize["jwks_uri"] = o.JwksUri
	}
	if o.OtpUrl != nil {
		toSerialize["otp_url"] = o.OtpUrl
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityOIDCSettingsRequest struct {
	value *IdentityOIDCSettingsRequest
	isSet bool
}

func (v NullableIdentityOIDCSettingsRequest) Get() *IdentityOIDCSettingsRequest {
	return v.value
}

func (v *NullableIdentityOIDCSettingsRequest) Set(val *IdentityOIDCSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityOIDCSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityOIDCSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityOIDCSettingsRequest(val *IdentityOIDCSettingsRequest) *NullableIdentityOIDCSettingsRequest {
	return &NullableIdentityOIDCSettingsRequest{value: val, isSet: true}
}

func (v NullableIdentityOIDCSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityOIDCSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

