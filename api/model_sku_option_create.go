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

// SkuOptionCreate struct for SkuOptionCreate
type SkuOptionCreate struct {
	Data SkuOptionCreateData `json:"data"`
}

// NewSkuOptionCreate instantiates a new SkuOptionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuOptionCreate(data SkuOptionCreateData) *SkuOptionCreate {
	this := SkuOptionCreate{}
	this.Data = data
	return &this
}

// NewSkuOptionCreateWithDefaults instantiates a new SkuOptionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuOptionCreateWithDefaults() *SkuOptionCreate {
	this := SkuOptionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *SkuOptionCreate) GetData() SkuOptionCreateData {
	if o == nil {
		var ret SkuOptionCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SkuOptionCreate) GetDataOk() (*SkuOptionCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SkuOptionCreate) SetData(v SkuOptionCreateData) {
	o.Data = v
}

func (o SkuOptionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSkuOptionCreate struct {
	value *SkuOptionCreate
	isSet bool
}

func (v NullableSkuOptionCreate) Get() *SkuOptionCreate {
	return v.value
}

func (v *NullableSkuOptionCreate) Set(val *SkuOptionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuOptionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuOptionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuOptionCreate(val *SkuOptionCreate) *NullableSkuOptionCreate {
	return &NullableSkuOptionCreate{value: val, isSet: true}
}

func (v NullableSkuOptionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuOptionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
