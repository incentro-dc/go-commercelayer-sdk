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

// TaxCategoryUpdate struct for TaxCategoryUpdate
type TaxCategoryUpdate struct {
	Data TaxCategoryUpdateData `json:"data"`
}

// NewTaxCategoryUpdate instantiates a new TaxCategoryUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxCategoryUpdate(data TaxCategoryUpdateData) *TaxCategoryUpdate {
	this := TaxCategoryUpdate{}
	this.Data = data
	return &this
}

// NewTaxCategoryUpdateWithDefaults instantiates a new TaxCategoryUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxCategoryUpdateWithDefaults() *TaxCategoryUpdate {
	this := TaxCategoryUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *TaxCategoryUpdate) GetData() TaxCategoryUpdateData {
	if o == nil {
		var ret TaxCategoryUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TaxCategoryUpdate) GetDataOk() (*TaxCategoryUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TaxCategoryUpdate) SetData(v TaxCategoryUpdateData) {
	o.Data = v
}

func (o TaxCategoryUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTaxCategoryUpdate struct {
	value *TaxCategoryUpdate
	isSet bool
}

func (v NullableTaxCategoryUpdate) Get() *TaxCategoryUpdate {
	return v.value
}

func (v *NullableTaxCategoryUpdate) Set(val *TaxCategoryUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxCategoryUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxCategoryUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxCategoryUpdate(val *TaxCategoryUpdate) *NullableTaxCategoryUpdate {
	return &NullableTaxCategoryUpdate{value: val, isSet: true}
}

func (v NullableTaxCategoryUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxCategoryUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
