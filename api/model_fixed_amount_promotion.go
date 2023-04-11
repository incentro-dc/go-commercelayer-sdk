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

// FixedAmountPromotion struct for FixedAmountPromotion
type FixedAmountPromotion struct {
	Data *FixedAmountPromotionData `json:"data,omitempty"`
}

// NewFixedAmountPromotion instantiates a new FixedAmountPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedAmountPromotion() *FixedAmountPromotion {
	this := FixedAmountPromotion{}
	return &this
}

// NewFixedAmountPromotionWithDefaults instantiates a new FixedAmountPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedAmountPromotionWithDefaults() *FixedAmountPromotion {
	this := FixedAmountPromotion{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FixedAmountPromotion) GetData() FixedAmountPromotionData {
	if o == nil || o.Data == nil {
		var ret FixedAmountPromotionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedAmountPromotion) GetDataOk() (*FixedAmountPromotionData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FixedAmountPromotion) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given FixedAmountPromotionData and assigns it to the Data field.
func (o *FixedAmountPromotion) SetData(v FixedAmountPromotionData) {
	o.Data = &v
}

func (o FixedAmountPromotion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableFixedAmountPromotion struct {
	value *FixedAmountPromotion
	isSet bool
}

func (v NullableFixedAmountPromotion) Get() *FixedAmountPromotion {
	return v.value
}

func (v *NullableFixedAmountPromotion) Set(val *FixedAmountPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedAmountPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedAmountPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedAmountPromotion(val *FixedAmountPromotion) *NullableFixedAmountPromotion {
	return &NullableFixedAmountPromotion{value: val, isSet: true}
}

func (v NullableFixedAmountPromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedAmountPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
