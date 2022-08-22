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

// TaxCalculator struct for TaxCalculator
type TaxCalculator struct {
	Data TaxCalculatorData `json:"data"`
}

// NewTaxCalculator instantiates a new TaxCalculator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCalculator(data TaxCalculatorData) *TaxCalculator {
	this := TaxCalculator{}
	this.Data = data
	return &this
}

// NewTaxCalculatorWithDefaults instantiates a new TaxCalculator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCalculatorWithDefaults() *TaxCalculator {
	this := TaxCalculator{}
	return &this
}

// GetData returns the Data field value
func (o *TaxCalculator) GetData() TaxCalculatorData {
	if o == nil {
		var ret TaxCalculatorData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TaxCalculator) GetDataOk() (*TaxCalculatorData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TaxCalculator) SetData(v TaxCalculatorData) {
	o.Data = v
}

func (o TaxCalculator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTaxCalculator struct {
	value *TaxCalculator
	isSet bool
}

func (v NullableTaxCalculator) Get() *TaxCalculator {
	return v.value
}

func (v *NullableTaxCalculator) Set(val *TaxCalculator) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCalculator) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCalculator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCalculator(val *TaxCalculator) *NullableTaxCalculator {
	return &NullableTaxCalculator{value: val, isSet: true}
}

func (v NullableTaxCalculator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCalculator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
