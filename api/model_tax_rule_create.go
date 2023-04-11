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

// TaxRuleCreate struct for TaxRuleCreate
type TaxRuleCreate struct {
	Data TaxRuleCreateData `json:"data"`
}

// NewTaxRuleCreate instantiates a new TaxRuleCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRuleCreate(data TaxRuleCreateData) *TaxRuleCreate {
	this := TaxRuleCreate{}
	this.Data = data
	return &this
}

// NewTaxRuleCreateWithDefaults instantiates a new TaxRuleCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRuleCreateWithDefaults() *TaxRuleCreate {
	this := TaxRuleCreate{}
	return &this
}

// GetData returns the Data field value
func (o *TaxRuleCreate) GetData() TaxRuleCreateData {
	if o == nil {
		var ret TaxRuleCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TaxRuleCreate) GetDataOk() (*TaxRuleCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TaxRuleCreate) SetData(v TaxRuleCreateData) {
	o.Data = v
}

func (o TaxRuleCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTaxRuleCreate struct {
	value *TaxRuleCreate
	isSet bool
}

func (v NullableTaxRuleCreate) Get() *TaxRuleCreate {
	return v.value
}

func (v *NullableTaxRuleCreate) Set(val *TaxRuleCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRuleCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRuleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRuleCreate(val *TaxRuleCreate) *NullableTaxRuleCreate {
	return &NullableTaxRuleCreate{value: val, isSet: true}
}

func (v NullableTaxRuleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRuleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
