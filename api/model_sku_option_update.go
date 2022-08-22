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

// SkuOptionUpdate struct for SkuOptionUpdate
type SkuOptionUpdate struct {
	Data SkuOptionUpdateData `json:"data"`
}

// NewSkuOptionUpdate instantiates a new SkuOptionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuOptionUpdate(data SkuOptionUpdateData) *SkuOptionUpdate {
	this := SkuOptionUpdate{}
	this.Data = data
	return &this
}

// NewSkuOptionUpdateWithDefaults instantiates a new SkuOptionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuOptionUpdateWithDefaults() *SkuOptionUpdate {
	this := SkuOptionUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *SkuOptionUpdate) GetData() SkuOptionUpdateData {
	if o == nil {
		var ret SkuOptionUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SkuOptionUpdate) GetDataOk() (*SkuOptionUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SkuOptionUpdate) SetData(v SkuOptionUpdateData) {
	o.Data = v
}

func (o SkuOptionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSkuOptionUpdate struct {
	value *SkuOptionUpdate
	isSet bool
}

func (v NullableSkuOptionUpdate) Get() *SkuOptionUpdate {
	return v.value
}

func (v *NullableSkuOptionUpdate) Set(val *SkuOptionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuOptionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuOptionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuOptionUpdate(val *SkuOptionUpdate) *NullableSkuOptionUpdate {
	return &NullableSkuOptionUpdate{value: val, isSet: true}
}

func (v NullableSkuOptionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuOptionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
