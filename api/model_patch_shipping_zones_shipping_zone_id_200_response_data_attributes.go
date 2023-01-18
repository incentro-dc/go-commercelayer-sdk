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

// PATCHShippingZonesShippingZoneId200ResponseDataAttributes struct for PATCHShippingZonesShippingZoneId200ResponseDataAttributes
type PATCHShippingZonesShippingZoneId200ResponseDataAttributes struct {
	// The shipping zone's internal name.
	Name *string `json:"name,omitempty"`
	// The regex that will be evaluated to match the shipping address country code.
	CountryCodeRegex *string `json:"country_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping address country code.
	NotCountryCodeRegex *string `json:"not_country_code_regex,omitempty"`
	// The regex that will be evaluated to match the shipping address state code.
	StateCodeRegex *string `json:"state_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping address state code.
	NotStateCodeRegex *string `json:"not_state_code_regex,omitempty"`
	// The regex that will be evaluated to match the shipping address zip code.
	ZipCodeRegex *string `json:"zip_code_regex,omitempty"`
	// The regex that will be evaluated as negative match for the shipping zip country code.
	NotZipCodeRegex *string `json:"not_zip_code_regex,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHShippingZonesShippingZoneId200ResponseDataAttributes instantiates a new PATCHShippingZonesShippingZoneId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHShippingZonesShippingZoneId200ResponseDataAttributes() *PATCHShippingZonesShippingZoneId200ResponseDataAttributes {
	this := PATCHShippingZonesShippingZoneId200ResponseDataAttributes{}
	return &this
}

// NewPATCHShippingZonesShippingZoneId200ResponseDataAttributesWithDefaults instantiates a new PATCHShippingZonesShippingZoneId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHShippingZonesShippingZoneId200ResponseDataAttributesWithDefaults() *PATCHShippingZonesShippingZoneId200ResponseDataAttributes {
	this := PATCHShippingZonesShippingZoneId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetCountryCodeRegex returns the CountryCodeRegex field value if set, zero value otherwise.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetCountryCodeRegex() string {
	if o == nil || o.CountryCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.CountryCodeRegex
}

// GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetCountryCodeRegexOk() (*string, bool) {
	if o == nil || o.CountryCodeRegex == nil {
		return nil, false
	}
	return o.CountryCodeRegex, true
}

// HasCountryCodeRegex returns a boolean if a field has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasCountryCodeRegex() bool {
	if o != nil && o.CountryCodeRegex != nil {
		return true
	}

	return false
}

// SetCountryCodeRegex gets a reference to the given string and assigns it to the CountryCodeRegex field.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetCountryCodeRegex(v string) {
	o.CountryCodeRegex = &v
}

// GetNotCountryCodeRegex returns the NotCountryCodeRegex field value if set, zero value otherwise.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotCountryCodeRegex() string {
	if o == nil || o.NotCountryCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.NotCountryCodeRegex
}

// GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotCountryCodeRegexOk() (*string, bool) {
	if o == nil || o.NotCountryCodeRegex == nil {
		return nil, false
	}
	return o.NotCountryCodeRegex, true
}

// HasNotCountryCodeRegex returns a boolean if a field has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasNotCountryCodeRegex() bool {
	if o != nil && o.NotCountryCodeRegex != nil {
		return true
	}

	return false
}

// SetNotCountryCodeRegex gets a reference to the given string and assigns it to the NotCountryCodeRegex field.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetNotCountryCodeRegex(v string) {
	o.NotCountryCodeRegex = &v
}

// GetStateCodeRegex returns the StateCodeRegex field value if set, zero value otherwise.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetStateCodeRegex() string {
	if o == nil || o.StateCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.StateCodeRegex
}

// GetStateCodeRegexOk returns a tuple with the StateCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetStateCodeRegexOk() (*string, bool) {
	if o == nil || o.StateCodeRegex == nil {
		return nil, false
	}
	return o.StateCodeRegex, true
}

// HasStateCodeRegex returns a boolean if a field has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasStateCodeRegex() bool {
	if o != nil && o.StateCodeRegex != nil {
		return true
	}

	return false
}

// SetStateCodeRegex gets a reference to the given string and assigns it to the StateCodeRegex field.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetStateCodeRegex(v string) {
	o.StateCodeRegex = &v
}

// GetNotStateCodeRegex returns the NotStateCodeRegex field value if set, zero value otherwise.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotStateCodeRegex() string {
	if o == nil || o.NotStateCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.NotStateCodeRegex
}

// GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotStateCodeRegexOk() (*string, bool) {
	if o == nil || o.NotStateCodeRegex == nil {
		return nil, false
	}
	return o.NotStateCodeRegex, true
}

// HasNotStateCodeRegex returns a boolean if a field has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasNotStateCodeRegex() bool {
	if o != nil && o.NotStateCodeRegex != nil {
		return true
	}

	return false
}

// SetNotStateCodeRegex gets a reference to the given string and assigns it to the NotStateCodeRegex field.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetNotStateCodeRegex(v string) {
	o.NotStateCodeRegex = &v
}

// GetZipCodeRegex returns the ZipCodeRegex field value if set, zero value otherwise.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetZipCodeRegex() string {
	if o == nil || o.ZipCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.ZipCodeRegex
}

// GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetZipCodeRegexOk() (*string, bool) {
	if o == nil || o.ZipCodeRegex == nil {
		return nil, false
	}
	return o.ZipCodeRegex, true
}

// HasZipCodeRegex returns a boolean if a field has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasZipCodeRegex() bool {
	if o != nil && o.ZipCodeRegex != nil {
		return true
	}

	return false
}

// SetZipCodeRegex gets a reference to the given string and assigns it to the ZipCodeRegex field.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetZipCodeRegex(v string) {
	o.ZipCodeRegex = &v
}

// GetNotZipCodeRegex returns the NotZipCodeRegex field value if set, zero value otherwise.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotZipCodeRegex() string {
	if o == nil || o.NotZipCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.NotZipCodeRegex
}

// GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetNotZipCodeRegexOk() (*string, bool) {
	if o == nil || o.NotZipCodeRegex == nil {
		return nil, false
	}
	return o.NotZipCodeRegex, true
}

// HasNotZipCodeRegex returns a boolean if a field has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasNotZipCodeRegex() bool {
	if o != nil && o.NotZipCodeRegex != nil {
		return true
	}

	return false
}

// SetNotZipCodeRegex gets a reference to the given string and assigns it to the NotZipCodeRegex field.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetNotZipCodeRegex(v string) {
	o.NotZipCodeRegex = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHShippingZonesShippingZoneId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CountryCodeRegex != nil {
		toSerialize["country_code_regex"] = o.CountryCodeRegex
	}
	if o.NotCountryCodeRegex != nil {
		toSerialize["not_country_code_regex"] = o.NotCountryCodeRegex
	}
	if o.StateCodeRegex != nil {
		toSerialize["state_code_regex"] = o.StateCodeRegex
	}
	if o.NotStateCodeRegex != nil {
		toSerialize["not_state_code_regex"] = o.NotStateCodeRegex
	}
	if o.ZipCodeRegex != nil {
		toSerialize["zip_code_regex"] = o.ZipCodeRegex
	}
	if o.NotZipCodeRegex != nil {
		toSerialize["not_zip_code_regex"] = o.NotZipCodeRegex
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
	return json.Marshal(toSerialize)
}

type NullablePATCHShippingZonesShippingZoneId200ResponseDataAttributes struct {
	value *PATCHShippingZonesShippingZoneId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHShippingZonesShippingZoneId200ResponseDataAttributes) Get() *PATCHShippingZonesShippingZoneId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHShippingZonesShippingZoneId200ResponseDataAttributes) Set(val *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHShippingZonesShippingZoneId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHShippingZonesShippingZoneId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHShippingZonesShippingZoneId200ResponseDataAttributes(val *PATCHShippingZonesShippingZoneId200ResponseDataAttributes) *NullablePATCHShippingZonesShippingZoneId200ResponseDataAttributes {
	return &NullablePATCHShippingZonesShippingZoneId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHShippingZonesShippingZoneId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHShippingZonesShippingZoneId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
