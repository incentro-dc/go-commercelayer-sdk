/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Adjustment struct for Adjustment
type Adjustment struct {
	Data AdjustmentData `json:"data"`
}

// NewAdjustment instantiates a new Adjustment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustment(data AdjustmentData) *Adjustment {
	this := Adjustment{}
	this.Data = data
	return &this
}

// NewAdjustmentWithDefaults instantiates a new Adjustment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustmentWithDefaults() *Adjustment {
	this := Adjustment{}
	return &this
}

// GetData returns the Data field value
func (o *Adjustment) GetData() AdjustmentData {
	if o == nil {
		var ret AdjustmentData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Adjustment) GetDataOk() (*AdjustmentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Adjustment) SetData(v AdjustmentData) {
	o.Data = v
}

func (o Adjustment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAdjustment struct {
	value *Adjustment
	isSet bool
}

func (v NullableAdjustment) Get() *Adjustment {
	return v.value
}

func (v *NullableAdjustment) Set(val *Adjustment) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustment) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustment(val *Adjustment) *NullableAdjustment {
	return &NullableAdjustment{value: val, isSet: true}
}

func (v NullableAdjustment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
