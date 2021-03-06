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

// ExternalPromotion struct for ExternalPromotion
type ExternalPromotion struct {
	Data ExternalPromotionData `json:"data"`
}

// NewExternalPromotion instantiates a new ExternalPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPromotion(data ExternalPromotionData) *ExternalPromotion {
	this := ExternalPromotion{}
	this.Data = data
	return &this
}

// NewExternalPromotionWithDefaults instantiates a new ExternalPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPromotionWithDefaults() *ExternalPromotion {
	this := ExternalPromotion{}
	return &this
}

// GetData returns the Data field value
func (o *ExternalPromotion) GetData() ExternalPromotionData {
	if o == nil {
		var ret ExternalPromotionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExternalPromotion) GetDataOk() (*ExternalPromotionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ExternalPromotion) SetData(v ExternalPromotionData) {
	o.Data = v
}

func (o ExternalPromotion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPromotion struct {
	value *ExternalPromotion
	isSet bool
}

func (v NullableExternalPromotion) Get() *ExternalPromotion {
	return v.value
}

func (v *NullableExternalPromotion) Set(val *ExternalPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPromotion(val *ExternalPromotion) *NullableExternalPromotion {
	return &NullableExternalPromotion{value: val, isSet: true}
}

func (v NullableExternalPromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
