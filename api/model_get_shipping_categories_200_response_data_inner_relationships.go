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

// GETShippingCategories200ResponseDataInnerRelationships struct for GETShippingCategories200ResponseDataInnerRelationships
type GETShippingCategories200ResponseDataInnerRelationships struct {
	Skus        *GETBundles200ResponseDataInnerRelationshipsSkus                `json:"skus,omitempty"`
	Attachments *GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewGETShippingCategories200ResponseDataInnerRelationships instantiates a new GETShippingCategories200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETShippingCategories200ResponseDataInnerRelationships() *GETShippingCategories200ResponseDataInnerRelationships {
	this := GETShippingCategories200ResponseDataInnerRelationships{}
	return &this
}

// NewGETShippingCategories200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETShippingCategories200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETShippingCategories200ResponseDataInnerRelationshipsWithDefaults() *GETShippingCategories200ResponseDataInnerRelationships {
	this := GETShippingCategories200ResponseDataInnerRelationships{}
	return &this
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *GETShippingCategories200ResponseDataInnerRelationships) GetSkus() GETBundles200ResponseDataInnerRelationshipsSkus {
	if o == nil || o.Skus == nil {
		var ret GETBundles200ResponseDataInnerRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingCategories200ResponseDataInnerRelationships) GetSkusOk() (*GETBundles200ResponseDataInnerRelationshipsSkus, bool) {
	if o == nil || o.Skus == nil {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *GETShippingCategories200ResponseDataInnerRelationships) HasSkus() bool {
	if o != nil && o.Skus != nil {
		return true
	}

	return false
}

// SetSkus gets a reference to the given GETBundles200ResponseDataInnerRelationshipsSkus and assigns it to the Skus field.
func (o *GETShippingCategories200ResponseDataInnerRelationships) SetSkus(v GETBundles200ResponseDataInnerRelationshipsSkus) {
	o.Skus = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETShippingCategories200ResponseDataInnerRelationships) GetAttachments() GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETShippingCategories200ResponseDataInnerRelationships) GetAttachmentsOk() (*GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETShippingCategories200ResponseDataInnerRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETShippingCategories200ResponseDataInnerRelationships) SetAttachments(v GETAvalaraAccounts200ResponseDataInnerRelationshipsAttachments) {
	o.Attachments = &v
}

func (o GETShippingCategories200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skus != nil {
		toSerialize["skus"] = o.Skus
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableGETShippingCategories200ResponseDataInnerRelationships struct {
	value *GETShippingCategories200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETShippingCategories200ResponseDataInnerRelationships) Get() *GETShippingCategories200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETShippingCategories200ResponseDataInnerRelationships) Set(val *GETShippingCategories200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETShippingCategories200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETShippingCategories200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETShippingCategories200ResponseDataInnerRelationships(val *GETShippingCategories200ResponseDataInnerRelationships) *NullableGETShippingCategories200ResponseDataInnerRelationships {
	return &NullableGETShippingCategories200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETShippingCategories200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETShippingCategories200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
