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

// ShippingWeightTier struct for ShippingWeightTier
type ShippingWeightTier struct {
	Data *ShippingWeightTierData `json:"data,omitempty"`
}

// NewShippingWeightTier instantiates a new ShippingWeightTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingWeightTier() *ShippingWeightTier {
	this := ShippingWeightTier{}
	return &this
}

// NewShippingWeightTierWithDefaults instantiates a new ShippingWeightTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingWeightTierWithDefaults() *ShippingWeightTier {
	this := ShippingWeightTier{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShippingWeightTier) GetData() ShippingWeightTierData {
	if o == nil || o.Data == nil {
		var ret ShippingWeightTierData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingWeightTier) GetDataOk() (*ShippingWeightTierData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShippingWeightTier) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ShippingWeightTierData and assigns it to the Data field.
func (o *ShippingWeightTier) SetData(v ShippingWeightTierData) {
	o.Data = &v
}

func (o ShippingWeightTier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingWeightTier struct {
	value *ShippingWeightTier
	isSet bool
}

func (v NullableShippingWeightTier) Get() *ShippingWeightTier {
	return v.value
}

func (v *NullableShippingWeightTier) Set(val *ShippingWeightTier) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingWeightTier) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingWeightTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingWeightTier(val *ShippingWeightTier) *NullableShippingWeightTier {
	return &NullableShippingWeightTier{value: val, isSet: true}
}

func (v NullableShippingWeightTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingWeightTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
