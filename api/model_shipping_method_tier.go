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

// ShippingMethodTier struct for ShippingMethodTier
type ShippingMethodTier struct {
	Data *ShippingMethodTierData `json:"data,omitempty"`
}

// NewShippingMethodTier instantiates a new ShippingMethodTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodTier() *ShippingMethodTier {
	this := ShippingMethodTier{}
	return &this
}

// NewShippingMethodTierWithDefaults instantiates a new ShippingMethodTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodTierWithDefaults() *ShippingMethodTier {
	this := ShippingMethodTier{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShippingMethodTier) GetData() ShippingMethodTierData {
	if o == nil || o.Data == nil {
		var ret ShippingMethodTierData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingMethodTier) GetDataOk() (*ShippingMethodTierData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShippingMethodTier) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ShippingMethodTierData and assigns it to the Data field.
func (o *ShippingMethodTier) SetData(v ShippingMethodTierData) {
	o.Data = &v
}

func (o ShippingMethodTier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingMethodTier struct {
	value *ShippingMethodTier
	isSet bool
}

func (v NullableShippingMethodTier) Get() *ShippingMethodTier {
	return v.value
}

func (v *NullableShippingMethodTier) Set(val *ShippingMethodTier) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodTier) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodTier(val *ShippingMethodTier) *NullableShippingMethodTier {
	return &NullableShippingMethodTier{value: val, isSet: true}
}

func (v NullableShippingMethodTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
