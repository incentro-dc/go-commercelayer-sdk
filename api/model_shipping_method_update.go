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

// ShippingMethodUpdate struct for ShippingMethodUpdate
type ShippingMethodUpdate struct {
	Data ShippingMethodUpdateData `json:"data"`
}

// NewShippingMethodUpdate instantiates a new ShippingMethodUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingMethodUpdate(data ShippingMethodUpdateData) *ShippingMethodUpdate {
	this := ShippingMethodUpdate{}
	this.Data = data
	return &this
}

// NewShippingMethodUpdateWithDefaults instantiates a new ShippingMethodUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingMethodUpdateWithDefaults() *ShippingMethodUpdate {
	this := ShippingMethodUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ShippingMethodUpdate) GetData() ShippingMethodUpdateData {
	if o == nil {
		var ret ShippingMethodUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ShippingMethodUpdate) GetDataOk() (*ShippingMethodUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ShippingMethodUpdate) SetData(v ShippingMethodUpdateData) {
	o.Data = v
}

func (o ShippingMethodUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableShippingMethodUpdate struct {
	value *ShippingMethodUpdate
	isSet bool
}

func (v NullableShippingMethodUpdate) Get() *ShippingMethodUpdate {
	return v.value
}

func (v *NullableShippingMethodUpdate) Set(val *ShippingMethodUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingMethodUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingMethodUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingMethodUpdate(val *ShippingMethodUpdate) *NullableShippingMethodUpdate {
	return &NullableShippingMethodUpdate{value: val, isSet: true}
}

func (v NullableShippingMethodUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingMethodUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
