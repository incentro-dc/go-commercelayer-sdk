/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// SkuListPromotionRuleUpdate struct for SkuListPromotionRuleUpdate
type SkuListPromotionRuleUpdate struct {
	Data SkuListPromotionRuleUpdateData `json:"data"`
}

// NewSkuListPromotionRuleUpdate instantiates a new SkuListPromotionRuleUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListPromotionRuleUpdate(data SkuListPromotionRuleUpdateData) *SkuListPromotionRuleUpdate {
	this := SkuListPromotionRuleUpdate{}
	this.Data = data
	return &this
}

// NewSkuListPromotionRuleUpdateWithDefaults instantiates a new SkuListPromotionRuleUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListPromotionRuleUpdateWithDefaults() *SkuListPromotionRuleUpdate {
	this := SkuListPromotionRuleUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *SkuListPromotionRuleUpdate) GetData() SkuListPromotionRuleUpdateData {
	if o == nil {
		var ret SkuListPromotionRuleUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleUpdate) GetDataOk() (*SkuListPromotionRuleUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SkuListPromotionRuleUpdate) SetData(v SkuListPromotionRuleUpdateData) {
	o.Data = v
}

func (o SkuListPromotionRuleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListPromotionRuleUpdate struct {
	value *SkuListPromotionRuleUpdate
	isSet bool
}

func (v NullableSkuListPromotionRuleUpdate) Get() *SkuListPromotionRuleUpdate {
	return v.value
}

func (v *NullableSkuListPromotionRuleUpdate) Set(val *SkuListPromotionRuleUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListPromotionRuleUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListPromotionRuleUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListPromotionRuleUpdate(val *SkuListPromotionRuleUpdate) *NullableSkuListPromotionRuleUpdate {
	return &NullableSkuListPromotionRuleUpdate{value: val, isSet: true}
}

func (v NullableSkuListPromotionRuleUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListPromotionRuleUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
