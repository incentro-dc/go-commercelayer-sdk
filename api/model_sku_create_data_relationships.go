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

// SkuCreateDataRelationships struct for SkuCreateDataRelationships
type SkuCreateDataRelationships struct {
	ShippingCategory ShipmentDataRelationshipsShippingCategory `json:"shipping_category"`
}

// NewSkuCreateDataRelationships instantiates a new SkuCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuCreateDataRelationships(shippingCategory ShipmentDataRelationshipsShippingCategory) *SkuCreateDataRelationships {
	this := SkuCreateDataRelationships{}
	this.ShippingCategory = shippingCategory
	return &this
}

// NewSkuCreateDataRelationshipsWithDefaults instantiates a new SkuCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuCreateDataRelationshipsWithDefaults() *SkuCreateDataRelationships {
	this := SkuCreateDataRelationships{}
	return &this
}

// GetShippingCategory returns the ShippingCategory field value
func (o *SkuCreateDataRelationships) GetShippingCategory() ShipmentDataRelationshipsShippingCategory {
	if o == nil {
		var ret ShipmentDataRelationshipsShippingCategory
		return ret
	}

	return o.ShippingCategory
}

// GetShippingCategoryOk returns a tuple with the ShippingCategory field value
// and a boolean to check if the value has been set.
func (o *SkuCreateDataRelationships) GetShippingCategoryOk() (*ShipmentDataRelationshipsShippingCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingCategory, true
}

// SetShippingCategory sets field value
func (o *SkuCreateDataRelationships) SetShippingCategory(v ShipmentDataRelationshipsShippingCategory) {
	o.ShippingCategory = v
}

func (o SkuCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["shipping_category"] = o.ShippingCategory
	}
	return json.Marshal(toSerialize)
}

type NullableSkuCreateDataRelationships struct {
	value *SkuCreateDataRelationships
	isSet bool
}

func (v NullableSkuCreateDataRelationships) Get() *SkuCreateDataRelationships {
	return v.value
}

func (v *NullableSkuCreateDataRelationships) Set(val *SkuCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuCreateDataRelationships(val *SkuCreateDataRelationships) *NullableSkuCreateDataRelationships {
	return &NullableSkuCreateDataRelationships{value: val, isSet: true}
}

func (v NullableSkuCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
