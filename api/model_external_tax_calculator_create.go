/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ExternalTaxCalculatorCreate struct for ExternalTaxCalculatorCreate
type ExternalTaxCalculatorCreate struct {
	Data ExternalTaxCalculatorCreateData `json:"data"`
}

// NewExternalTaxCalculatorCreate instantiates a new ExternalTaxCalculatorCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalTaxCalculatorCreate(data ExternalTaxCalculatorCreateData) *ExternalTaxCalculatorCreate {
	this := ExternalTaxCalculatorCreate{}
	this.Data = data
	return &this
}

// NewExternalTaxCalculatorCreateWithDefaults instantiates a new ExternalTaxCalculatorCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalTaxCalculatorCreateWithDefaults() *ExternalTaxCalculatorCreate {
	this := ExternalTaxCalculatorCreate{}
	return &this
}

// GetData returns the Data field value
func (o *ExternalTaxCalculatorCreate) GetData() ExternalTaxCalculatorCreateData {
	if o == nil {
		var ret ExternalTaxCalculatorCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExternalTaxCalculatorCreate) GetDataOk() (*ExternalTaxCalculatorCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExternalTaxCalculatorCreate) SetData(v ExternalTaxCalculatorCreateData) {
	o.Data = v
}

func (o ExternalTaxCalculatorCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalTaxCalculatorCreate struct {
	value *ExternalTaxCalculatorCreate
	isSet bool
}

func (v NullableExternalTaxCalculatorCreate) Get() *ExternalTaxCalculatorCreate {
	return v.value
}

func (v *NullableExternalTaxCalculatorCreate) Set(val *ExternalTaxCalculatorCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalTaxCalculatorCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalTaxCalculatorCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalTaxCalculatorCreate(val *ExternalTaxCalculatorCreate) *NullableExternalTaxCalculatorCreate {
	return &NullableExternalTaxCalculatorCreate{value: val, isSet: true}
}

func (v NullableExternalTaxCalculatorCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalTaxCalculatorCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
