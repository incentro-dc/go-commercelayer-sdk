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

// FixedPricePromotionCreate struct for FixedPricePromotionCreate
type FixedPricePromotionCreate struct {
	Data FixedPricePromotionCreateData `json:"data"`
}

// NewFixedPricePromotionCreate instantiates a new FixedPricePromotionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedPricePromotionCreate(data FixedPricePromotionCreateData) *FixedPricePromotionCreate {
	this := FixedPricePromotionCreate{}
	this.Data = data
	return &this
}

// NewFixedPricePromotionCreateWithDefaults instantiates a new FixedPricePromotionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedPricePromotionCreateWithDefaults() *FixedPricePromotionCreate {
	this := FixedPricePromotionCreate{}
	return &this
}

// GetData returns the Data field value
func (o *FixedPricePromotionCreate) GetData() FixedPricePromotionCreateData {
	if o == nil {
		var ret FixedPricePromotionCreateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FixedPricePromotionCreate) GetDataOk() (*FixedPricePromotionCreateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FixedPricePromotionCreate) SetData(v FixedPricePromotionCreateData) {
	o.Data = v
}

func (o FixedPricePromotionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableFixedPricePromotionCreate struct {
	value *FixedPricePromotionCreate
	isSet bool
}

func (v NullableFixedPricePromotionCreate) Get() *FixedPricePromotionCreate {
	return v.value
}

func (v *NullableFixedPricePromotionCreate) Set(val *FixedPricePromotionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedPricePromotionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedPricePromotionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedPricePromotionCreate(val *FixedPricePromotionCreate) *NullableFixedPricePromotionCreate {
	return &NullableFixedPricePromotionCreate{value: val, isSet: true}
}

func (v NullableFixedPricePromotionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedPricePromotionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
