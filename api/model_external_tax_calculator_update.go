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

// ExternalTaxCalculatorUpdate struct for ExternalTaxCalculatorUpdate
type ExternalTaxCalculatorUpdate struct {
	Data ExternalTaxCalculatorUpdateData `json:"data"`
}

// NewExternalTaxCalculatorUpdate instantiates a new ExternalTaxCalculatorUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTaxCalculatorUpdate(data ExternalTaxCalculatorUpdateData) *ExternalTaxCalculatorUpdate {
	this := ExternalTaxCalculatorUpdate{}
	this.Data = data
	return &this
}

// NewExternalTaxCalculatorUpdateWithDefaults instantiates a new ExternalTaxCalculatorUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTaxCalculatorUpdateWithDefaults() *ExternalTaxCalculatorUpdate {
	this := ExternalTaxCalculatorUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *ExternalTaxCalculatorUpdate) GetData() ExternalTaxCalculatorUpdateData {
	if o == nil {
		var ret ExternalTaxCalculatorUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorUpdate) GetDataOk() (*ExternalTaxCalculatorUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExternalTaxCalculatorUpdate) SetData(v ExternalTaxCalculatorUpdateData) {
	o.Data = v
}

func (o ExternalTaxCalculatorUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalTaxCalculatorUpdate struct {
	value *ExternalTaxCalculatorUpdate
	isSet bool
}

func (v NullableExternalTaxCalculatorUpdate) Get() *ExternalTaxCalculatorUpdate {
	return v.value
}

func (v *NullableExternalTaxCalculatorUpdate) Set(val *ExternalTaxCalculatorUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTaxCalculatorUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTaxCalculatorUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTaxCalculatorUpdate(val *ExternalTaxCalculatorUpdate) *NullableExternalTaxCalculatorUpdate {
	return &NullableExternalTaxCalculatorUpdate{value: val, isSet: true}
}

func (v NullableExternalTaxCalculatorUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTaxCalculatorUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
