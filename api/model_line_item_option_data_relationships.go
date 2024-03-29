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

// checks if the LineItemOptionDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LineItemOptionDataRelationships{}

// LineItemOptionDataRelationships struct for LineItemOptionDataRelationships
type LineItemOptionDataRelationships struct {
	LineItem  *LineItemOptionDataRelationshipsLineItem  `json:"line_item,omitempty"`
	SkuOption *LineItemOptionDataRelationshipsSkuOption `json:"sku_option,omitempty"`
}

// NewLineItemOptionDataRelationships instantiates a new LineItemOptionDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionDataRelationships() *LineItemOptionDataRelationships {
	this := LineItemOptionDataRelationships{}
	return &this
}

// NewLineItemOptionDataRelationshipsWithDefaults instantiates a new LineItemOptionDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionDataRelationshipsWithDefaults() *LineItemOptionDataRelationships {
	this := LineItemOptionDataRelationships{}
	return &this
}

// GetLineItem returns the LineItem field value if set, zero value otherwise.
func (o *LineItemOptionDataRelationships) GetLineItem() LineItemOptionDataRelationshipsLineItem {
	if o == nil || IsNil(o.LineItem) {
		var ret LineItemOptionDataRelationshipsLineItem
		return ret
	}
	return *o.LineItem
}

// GetLineItemOk returns a tuple with the LineItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataRelationships) GetLineItemOk() (*LineItemOptionDataRelationshipsLineItem, bool) {
	if o == nil || IsNil(o.LineItem) {
		return nil, false
	}
	return o.LineItem, true
}

// HasLineItem returns a boolean if a field has been set.
func (o *LineItemOptionDataRelationships) HasLineItem() bool {
	if o != nil && !IsNil(o.LineItem) {
		return true
	}

	return false
}

// SetLineItem gets a reference to the given LineItemOptionDataRelationshipsLineItem and assigns it to the LineItem field.
func (o *LineItemOptionDataRelationships) SetLineItem(v LineItemOptionDataRelationshipsLineItem) {
	o.LineItem = &v
}

// GetSkuOption returns the SkuOption field value if set, zero value otherwise.
func (o *LineItemOptionDataRelationships) GetSkuOption() LineItemOptionDataRelationshipsSkuOption {
	if o == nil || IsNil(o.SkuOption) {
		var ret LineItemOptionDataRelationshipsSkuOption
		return ret
	}
	return *o.SkuOption
}

// GetSkuOptionOk returns a tuple with the SkuOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionDataRelationships) GetSkuOptionOk() (*LineItemOptionDataRelationshipsSkuOption, bool) {
	if o == nil || IsNil(o.SkuOption) {
		return nil, false
	}
	return o.SkuOption, true
}

// HasSkuOption returns a boolean if a field has been set.
func (o *LineItemOptionDataRelationships) HasSkuOption() bool {
	if o != nil && !IsNil(o.SkuOption) {
		return true
	}

	return false
}

// SetSkuOption gets a reference to the given LineItemOptionDataRelationshipsSkuOption and assigns it to the SkuOption field.
func (o *LineItemOptionDataRelationships) SetSkuOption(v LineItemOptionDataRelationshipsSkuOption) {
	o.SkuOption = &v
}

func (o LineItemOptionDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LineItemOptionDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LineItem) {
		toSerialize["line_item"] = o.LineItem
	}
	if !IsNil(o.SkuOption) {
		toSerialize["sku_option"] = o.SkuOption
	}
	return toSerialize, nil
}

type NullableLineItemOptionDataRelationships struct {
	value *LineItemOptionDataRelationships
	isSet bool
}

func (v NullableLineItemOptionDataRelationships) Get() *LineItemOptionDataRelationships {
	return v.value
}

func (v *NullableLineItemOptionDataRelationships) Set(val *LineItemOptionDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionDataRelationships(val *LineItemOptionDataRelationships) *NullableLineItemOptionDataRelationships {
	return &NullableLineItemOptionDataRelationships{value: val, isSet: true}
}

func (v NullableLineItemOptionDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
