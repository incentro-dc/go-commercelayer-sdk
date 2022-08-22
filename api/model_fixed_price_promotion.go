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

// FixedPricePromotion struct for FixedPricePromotion
type FixedPricePromotion struct {
	Data FixedPricePromotionData `json:"data"`
}

// NewFixedPricePromotion instantiates a new FixedPricePromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedPricePromotion(data FixedPricePromotionData) *FixedPricePromotion {
	this := FixedPricePromotion{}
	this.Data = data
	return &this
}

// NewFixedPricePromotionWithDefaults instantiates a new FixedPricePromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedPricePromotionWithDefaults() *FixedPricePromotion {
	this := FixedPricePromotion{}
	return &this
}

// GetData returns the Data field value
func (o *FixedPricePromotion) GetData() FixedPricePromotionData {
	if o == nil {
		var ret FixedPricePromotionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FixedPricePromotion) GetDataOk() (*FixedPricePromotionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *FixedPricePromotion) SetData(v FixedPricePromotionData) {
	o.Data = v
}

func (o FixedPricePromotion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableFixedPricePromotion struct {
	value *FixedPricePromotion
	isSet bool
}

func (v NullableFixedPricePromotion) Get() *FixedPricePromotion {
	return v.value
}

func (v *NullableFixedPricePromotion) Set(val *FixedPricePromotion) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedPricePromotion) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedPricePromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedPricePromotion(val *FixedPricePromotion) *NullableFixedPricePromotion {
	return &NullableFixedPricePromotion{value: val, isSet: true}
}

func (v NullableFixedPricePromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedPricePromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
