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

// SkuListPromotionRule struct for SkuListPromotionRule
type SkuListPromotionRule struct {
	Data *SkuListPromotionRuleData `json:"data,omitempty"`
}

// NewSkuListPromotionRule instantiates a new SkuListPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListPromotionRule() *SkuListPromotionRule {
	this := SkuListPromotionRule{}
	return &this
}

// NewSkuListPromotionRuleWithDefaults instantiates a new SkuListPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListPromotionRuleWithDefaults() *SkuListPromotionRule {
	this := SkuListPromotionRule{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SkuListPromotionRule) GetData() SkuListPromotionRuleData {
	if o == nil || o.Data == nil {
		var ret SkuListPromotionRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRule) GetDataOk() (*SkuListPromotionRuleData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SkuListPromotionRule) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given SkuListPromotionRuleData and assigns it to the Data field.
func (o *SkuListPromotionRule) SetData(v SkuListPromotionRuleData) {
	o.Data = &v
}

func (o SkuListPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListPromotionRule struct {
	value *SkuListPromotionRule
	isSet bool
}

func (v NullableSkuListPromotionRule) Get() *SkuListPromotionRule {
	return v.value
}

func (v *NullableSkuListPromotionRule) Set(val *SkuListPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListPromotionRule(val *SkuListPromotionRule) *NullableSkuListPromotionRule {
	return &NullableSkuListPromotionRule{value: val, isSet: true}
}

func (v NullableSkuListPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
