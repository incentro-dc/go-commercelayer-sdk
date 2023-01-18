/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes struct for PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes
type PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes struct {
	// The payment gateway's internal name.
	Name *string `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The gateway country code one of EU, US, or OC.
	CountryCode *string `json:"country_code,omitempty"`
	// The public key linked to your API credential.
	ApiKey *string `json:"api_key,omitempty"`
	// The gateway API key.
	ApiSecret *string `json:"api_secret,omitempty"`
}

// NewPATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes instantiates a new PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes() *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes {
	this := PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes{}
	return &this
}

// NewPATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributesWithDefaults instantiates a new PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributesWithDefaults() *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes {
	this := PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetApiSecret returns the ApiSecret field value if set, zero value otherwise.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetApiSecret() string {
	if o == nil || o.ApiSecret == nil {
		var ret string
		return ret
	}
	return *o.ApiSecret
}

// GetApiSecretOk returns a tuple with the ApiSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) GetApiSecretOk() (*string, bool) {
	if o == nil || o.ApiSecret == nil {
		return nil, false
	}
	return o.ApiSecret, true
}

// HasApiSecret returns a boolean if a field has been set.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) HasApiSecret() bool {
	if o != nil && o.ApiSecret != nil {
		return true
	}

	return false
}

// SetApiSecret gets a reference to the given string and assigns it to the ApiSecret field.
func (o *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) SetApiSecret(v string) {
	o.ApiSecret = &v
}

func (o PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	if o.ApiSecret != nil {
		toSerialize["api_secret"] = o.ApiSecret
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes struct {
	value *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) Get() *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) Set(val *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes(val *PATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) *NullablePATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes {
	return &NullablePATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHKlarnaGatewaysKlarnaGatewayId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
