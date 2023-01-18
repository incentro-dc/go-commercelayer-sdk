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

// ShippingCategoryDataRelationships struct for ShippingCategoryDataRelationships
type ShippingCategoryDataRelationships struct {
	Skus        *BundleDataRelationshipsSkus                `json:"skus,omitempty"`
	Attachments *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewShippingCategoryDataRelationships instantiates a new ShippingCategoryDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingCategoryDataRelationships() *ShippingCategoryDataRelationships {
	this := ShippingCategoryDataRelationships{}
	return &this
}

// NewShippingCategoryDataRelationshipsWithDefaults instantiates a new ShippingCategoryDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingCategoryDataRelationshipsWithDefaults() *ShippingCategoryDataRelationships {
	this := ShippingCategoryDataRelationships{}
	return &this
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *ShippingCategoryDataRelationships) GetSkus() BundleDataRelationshipsSkus {
	if o == nil || o.Skus == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCategoryDataRelationships) GetSkusOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || o.Skus == nil {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *ShippingCategoryDataRelationships) HasSkus() bool {
	if o != nil && o.Skus != nil {
		return true
	}

	return false
}

// SetSkus gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Skus field.
func (o *ShippingCategoryDataRelationships) SetSkus(v BundleDataRelationshipsSkus) {
	o.Skus = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ShippingCategoryDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingCategoryDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ShippingCategoryDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ShippingCategoryDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o ShippingCategoryDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skus != nil {
		toSerialize["skus"] = o.Skus
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableShippingCategoryDataRelationships struct {
	value *ShippingCategoryDataRelationships
	isSet bool
}

func (v NullableShippingCategoryDataRelationships) Get() *ShippingCategoryDataRelationships {
	return v.value
}

func (v *NullableShippingCategoryDataRelationships) Set(val *ShippingCategoryDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingCategoryDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingCategoryDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingCategoryDataRelationships(val *ShippingCategoryDataRelationships) *NullableShippingCategoryDataRelationships {
	return &NullableShippingCategoryDataRelationships{value: val, isSet: true}
}

func (v NullableShippingCategoryDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingCategoryDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
