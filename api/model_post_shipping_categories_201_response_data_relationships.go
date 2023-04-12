/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the POSTShippingCategories201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingCategories201ResponseDataRelationships{}

// POSTShippingCategories201ResponseDataRelationships struct for POSTShippingCategories201ResponseDataRelationships
type POSTShippingCategories201ResponseDataRelationships struct {
	Skus        *POSTBundles201ResponseDataRelationshipsSkus                `json:"skus,omitempty"`
	Attachments *POSTAvalaraAccounts201ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewPOSTShippingCategories201ResponseDataRelationships instantiates a new POSTShippingCategories201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingCategories201ResponseDataRelationships() *POSTShippingCategories201ResponseDataRelationships {
	this := POSTShippingCategories201ResponseDataRelationships{}
	return &this
}

// NewPOSTShippingCategories201ResponseDataRelationshipsWithDefaults instantiates a new POSTShippingCategories201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingCategories201ResponseDataRelationshipsWithDefaults() *POSTShippingCategories201ResponseDataRelationships {
	this := POSTShippingCategories201ResponseDataRelationships{}
	return &this
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *POSTShippingCategories201ResponseDataRelationships) GetSkus() POSTBundles201ResponseDataRelationshipsSkus {
	if o == nil || IsNil(o.Skus) {
		var ret POSTBundles201ResponseDataRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingCategories201ResponseDataRelationships) GetSkusOk() (*POSTBundles201ResponseDataRelationshipsSkus, bool) {
	if o == nil || IsNil(o.Skus) {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *POSTShippingCategories201ResponseDataRelationships) HasSkus() bool {
	if o != nil && !IsNil(o.Skus) {
		return true
	}

	return false
}

// SetSkus gets a reference to the given POSTBundles201ResponseDataRelationshipsSkus and assigns it to the Skus field.
func (o *POSTShippingCategories201ResponseDataRelationships) SetSkus(v POSTBundles201ResponseDataRelationshipsSkus) {
	o.Skus = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *POSTShippingCategories201ResponseDataRelationships) GetAttachments() POSTAvalaraAccounts201ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret POSTAvalaraAccounts201ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTShippingCategories201ResponseDataRelationships) GetAttachmentsOk() (*POSTAvalaraAccounts201ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *POSTShippingCategories201ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given POSTAvalaraAccounts201ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *POSTShippingCategories201ResponseDataRelationships) SetAttachments(v POSTAvalaraAccounts201ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o POSTShippingCategories201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingCategories201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Skus) {
		toSerialize["skus"] = o.Skus
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

type NullablePOSTShippingCategories201ResponseDataRelationships struct {
	value *POSTShippingCategories201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTShippingCategories201ResponseDataRelationships) Get() *POSTShippingCategories201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTShippingCategories201ResponseDataRelationships) Set(val *POSTShippingCategories201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingCategories201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingCategories201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingCategories201ResponseDataRelationships(val *POSTShippingCategories201ResponseDataRelationships) *NullablePOSTShippingCategories201ResponseDataRelationships {
	return &NullablePOSTShippingCategories201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTShippingCategories201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingCategories201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
