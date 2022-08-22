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

// ManualTaxCalculatorUpdate struct for ManualTaxCalculatorUpdate
type ManualTaxCalculatorUpdate struct {
	Data ManualTaxCalculatorUpdateData `json:"data"`
}

// NewManualTaxCalculatorUpdate instantiates a new ManualTaxCalculatorUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManualTaxCalculatorUpdate(data ManualTaxCalculatorUpdateData) *ManualTaxCalculatorUpdate {
	this := ManualTaxCalculatorUpdate{}
	this.Data = data
	return &this
}

// NewManualTaxCalculatorUpdateWithDefaults instantiates a new ManualTaxCalculatorUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManualTaxCalculatorUpdateWithDefaults() *ManualTaxCalculatorUpdate {
	this := ManualTaxCalculatorUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ManualTaxCalculatorUpdate) GetData() ManualTaxCalculatorUpdateData {
	if o == nil {
		var ret ManualTaxCalculatorUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ManualTaxCalculatorUpdate) GetDataOk() (*ManualTaxCalculatorUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ManualTaxCalculatorUpdate) SetData(v ManualTaxCalculatorUpdateData) {
	o.Data = v
}

func (o ManualTaxCalculatorUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableManualTaxCalculatorUpdate struct {
	value *ManualTaxCalculatorUpdate
	isSet bool
}

func (v NullableManualTaxCalculatorUpdate) Get() *ManualTaxCalculatorUpdate {
	return v.value
}

func (v *NullableManualTaxCalculatorUpdate) Set(val *ManualTaxCalculatorUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableManualTaxCalculatorUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableManualTaxCalculatorUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManualTaxCalculatorUpdate(val *ManualTaxCalculatorUpdate) *NullableManualTaxCalculatorUpdate {
	return &NullableManualTaxCalculatorUpdate{value: val, isSet: true}
}

func (v NullableManualTaxCalculatorUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManualTaxCalculatorUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
