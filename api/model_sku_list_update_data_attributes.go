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

// SkuListUpdateDataAttributes struct for SkuListUpdateDataAttributes
type SkuListUpdateDataAttributes struct {
	// The SKU list's internal name.
	Name *string `json:"name,omitempty"`
	// An internal description of the SKU list.
	Description *string `json:"description,omitempty"`
	// The URL of an image that represents the SKU list.
	ImageUrl *string `json:"image_url,omitempty"`
	// Indicates if the SKU list is populated manually.
	Manual *bool `json:"manual,omitempty"`
	// The regex that will be evaluated to match the SKU codes.
	SkuCodeRegex *string `json:"sku_code_regex,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewSkuListUpdateDataAttributes instantiates a new SkuListUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListUpdateDataAttributes() *SkuListUpdateDataAttributes {
	this := SkuListUpdateDataAttributes{}
	return &this
}

// NewSkuListUpdateDataAttributesWithDefaults instantiates a new SkuListUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListUpdateDataAttributesWithDefaults() *SkuListUpdateDataAttributes {
	this := SkuListUpdateDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SkuListUpdateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListUpdateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SkuListUpdateDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SkuListUpdateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SkuListUpdateDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListUpdateDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SkuListUpdateDataAttributes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SkuListUpdateDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *SkuListUpdateDataAttributes) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListUpdateDataAttributes) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *SkuListUpdateDataAttributes) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *SkuListUpdateDataAttributes) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetManual returns the Manual field value if set, zero value otherwise.
func (o *SkuListUpdateDataAttributes) GetManual() bool {
	if o == nil || o.Manual == nil {
		var ret bool
		return ret
	}
	return *o.Manual
}

// GetManualOk returns a tuple with the Manual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListUpdateDataAttributes) GetManualOk() (*bool, bool) {
	if o == nil || o.Manual == nil {
		return nil, false
	}
	return o.Manual, true
}

// HasManual returns a boolean if a field has been set.
func (o *SkuListUpdateDataAttributes) HasManual() bool {
	if o != nil && o.Manual != nil {
		return true
	}

	return false
}

// SetManual gets a reference to the given bool and assigns it to the Manual field.
func (o *SkuListUpdateDataAttributes) SetManual(v bool) {
	o.Manual = &v
}

// GetSkuCodeRegex returns the SkuCodeRegex field value if set, zero value otherwise.
func (o *SkuListUpdateDataAttributes) GetSkuCodeRegex() string {
	if o == nil || o.SkuCodeRegex == nil {
		var ret string
		return ret
	}
	return *o.SkuCodeRegex
}

// GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListUpdateDataAttributes) GetSkuCodeRegexOk() (*string, bool) {
	if o == nil || o.SkuCodeRegex == nil {
		return nil, false
	}
	return o.SkuCodeRegex, true
}

// HasSkuCodeRegex returns a boolean if a field has been set.
func (o *SkuListUpdateDataAttributes) HasSkuCodeRegex() bool {
	if o != nil && o.SkuCodeRegex != nil {
		return true
	}

	return false
}

// SetSkuCodeRegex gets a reference to the given string and assigns it to the SkuCodeRegex field.
func (o *SkuListUpdateDataAttributes) SetSkuCodeRegex(v string) {
	o.SkuCodeRegex = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *SkuListUpdateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListUpdateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *SkuListUpdateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *SkuListUpdateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *SkuListUpdateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListUpdateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *SkuListUpdateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *SkuListUpdateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SkuListUpdateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListUpdateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SkuListUpdateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *SkuListUpdateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o SkuListUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.Manual != nil {
		toSerialize["manual"] = o.Manual
	}
	if o.SkuCodeRegex != nil {
		toSerialize["sku_code_regex"] = o.SkuCodeRegex
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

type NullableSkuListUpdateDataAttributes struct {
	value *SkuListUpdateDataAttributes
	isSet bool
}

func (v NullableSkuListUpdateDataAttributes) Get() *SkuListUpdateDataAttributes {
	return v.value
}

func (v *NullableSkuListUpdateDataAttributes) Set(val *SkuListUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListUpdateDataAttributes(val *SkuListUpdateDataAttributes) *NullableSkuListUpdateDataAttributes {
	return &NullableSkuListUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableSkuListUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
