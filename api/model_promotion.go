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

// Promotion struct for Promotion
type Promotion struct {
	Data PromotionData `json:"data"`
}

// NewPromotion instantiates a new Promotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotion(data PromotionData) *Promotion {
	this := Promotion{}
	this.Data = data
	return &this
}

// NewPromotionWithDefaults instantiates a new Promotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionWithDefaults() *Promotion {
	this := Promotion{}
	return &this
}

// GetData returns the Data field value
func (o *Promotion) GetData() PromotionData {
	if o == nil {
		var ret PromotionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Promotion) GetDataOk() (*PromotionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Promotion) SetData(v PromotionData) {
	o.Data = v
}

func (o Promotion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePromotion struct {
	value *Promotion
	isSet bool
}

func (v NullablePromotion) Get() *Promotion {
	return v.value
}

func (v *NullablePromotion) Set(val *Promotion) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotion) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotion(val *Promotion) *NullablePromotion {
	return &NullablePromotion{value: val, isSet: true}
}

func (v NullablePromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}