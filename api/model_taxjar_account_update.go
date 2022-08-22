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

// TaxjarAccountUpdate struct for TaxjarAccountUpdate
type TaxjarAccountUpdate struct {
	Data TaxjarAccountUpdateData `json:"data"`
}

// NewTaxjarAccountUpdate instantiates a new TaxjarAccountUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxjarAccountUpdate(data TaxjarAccountUpdateData) *TaxjarAccountUpdate {
	this := TaxjarAccountUpdate{}
	this.Data = data
	return &this
}

// NewTaxjarAccountUpdateWithDefaults instantiates a new TaxjarAccountUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxjarAccountUpdateWithDefaults() *TaxjarAccountUpdate {
	this := TaxjarAccountUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *TaxjarAccountUpdate) GetData() TaxjarAccountUpdateData {
	if o == nil {
		var ret TaxjarAccountUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TaxjarAccountUpdate) GetDataOk() (*TaxjarAccountUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TaxjarAccountUpdate) SetData(v TaxjarAccountUpdateData) {
	o.Data = v
}

func (o TaxjarAccountUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTaxjarAccountUpdate struct {
	value *TaxjarAccountUpdate
	isSet bool
}

func (v NullableTaxjarAccountUpdate) Get() *TaxjarAccountUpdate {
	return v.value
}

func (v *NullableTaxjarAccountUpdate) Set(val *TaxjarAccountUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxjarAccountUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxjarAccountUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxjarAccountUpdate(val *TaxjarAccountUpdate) *NullableTaxjarAccountUpdate {
	return &NullableTaxjarAccountUpdate{value: val, isSet: true}
}

func (v NullableTaxjarAccountUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxjarAccountUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
