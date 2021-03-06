/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GoogleGeocoderCreateDataAttributes struct for GoogleGeocoderCreateDataAttributes
type GoogleGeocoderCreateDataAttributes struct {
	// The geocoder's internal name
	Name string `json:"name"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The Google Map API key
	ApiKey string `json:"api_key"`
}

// NewGoogleGeocoderCreateDataAttributes instantiates a new GoogleGeocoderCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleGeocoderCreateDataAttributes(name string, apiKey string) *GoogleGeocoderCreateDataAttributes {
	this := GoogleGeocoderCreateDataAttributes{}
	this.Name = name
	this.ApiKey = apiKey
	return &this
}

// NewGoogleGeocoderCreateDataAttributesWithDefaults instantiates a new GoogleGeocoderCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleGeocoderCreateDataAttributesWithDefaults() *GoogleGeocoderCreateDataAttributes {
	this := GoogleGeocoderCreateDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *GoogleGeocoderCreateDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderCreateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GoogleGeocoderCreateDataAttributes) SetName(v string) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GoogleGeocoderCreateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderCreateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GoogleGeocoderCreateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GoogleGeocoderCreateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GoogleGeocoderCreateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderCreateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GoogleGeocoderCreateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GoogleGeocoderCreateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GoogleGeocoderCreateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderCreateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GoogleGeocoderCreateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GoogleGeocoderCreateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetApiKey returns the ApiKey field value
func (o *GoogleGeocoderCreateDataAttributes) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *GoogleGeocoderCreateDataAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *GoogleGeocoderCreateDataAttributes) SetApiKey(v string) {
	o.ApiKey = v
}

func (o GoogleGeocoderCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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
	if true {
		toSerialize["api_key"] = o.ApiKey
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleGeocoderCreateDataAttributes struct {
	value *GoogleGeocoderCreateDataAttributes
	isSet bool
}

func (v NullableGoogleGeocoderCreateDataAttributes) Get() *GoogleGeocoderCreateDataAttributes {
	return v.value
}

func (v *NullableGoogleGeocoderCreateDataAttributes) Set(val *GoogleGeocoderCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleGeocoderCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleGeocoderCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleGeocoderCreateDataAttributes(val *GoogleGeocoderCreateDataAttributes) *NullableGoogleGeocoderCreateDataAttributes {
	return &NullableGoogleGeocoderCreateDataAttributes{value: val, isSet: true}
}

func (v NullableGoogleGeocoderCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleGeocoderCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
