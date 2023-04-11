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

// CouponCreateDataRelationshipsPromotionRule struct for CouponCreateDataRelationshipsPromotionRule
type CouponCreateDataRelationshipsPromotionRule struct {
	Data CouponDataRelationshipsPromotionRuleData `json:"data"`
}

// NewCouponCreateDataRelationshipsPromotionRule instantiates a new CouponCreateDataRelationshipsPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCreateDataRelationshipsPromotionRule(data CouponDataRelationshipsPromotionRuleData) *CouponCreateDataRelationshipsPromotionRule {
	this := CouponCreateDataRelationshipsPromotionRule{}
	this.Data = data
	return &this
}

// NewCouponCreateDataRelationshipsPromotionRuleWithDefaults instantiates a new CouponCreateDataRelationshipsPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCreateDataRelationshipsPromotionRuleWithDefaults() *CouponCreateDataRelationshipsPromotionRule {
	this := CouponCreateDataRelationshipsPromotionRule{}
	return &this
}

// GetData returns the Data field value
func (o *CouponCreateDataRelationshipsPromotionRule) GetData() CouponDataRelationshipsPromotionRuleData {
	if o == nil {
		var ret CouponDataRelationshipsPromotionRuleData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponCreateDataRelationshipsPromotionRule) GetDataOk() (*CouponDataRelationshipsPromotionRuleData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponCreateDataRelationshipsPromotionRule) SetData(v CouponDataRelationshipsPromotionRuleData) {
	o.Data = v
}

func (o CouponCreateDataRelationshipsPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCouponCreateDataRelationshipsPromotionRule struct {
	value *CouponCreateDataRelationshipsPromotionRule
	isSet bool
}

func (v NullableCouponCreateDataRelationshipsPromotionRule) Get() *CouponCreateDataRelationshipsPromotionRule {
	return v.value
}

func (v *NullableCouponCreateDataRelationshipsPromotionRule) Set(val *CouponCreateDataRelationshipsPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCreateDataRelationshipsPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCreateDataRelationshipsPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCreateDataRelationshipsPromotionRule(val *CouponCreateDataRelationshipsPromotionRule) *NullableCouponCreateDataRelationshipsPromotionRule {
	return &NullableCouponCreateDataRelationshipsPromotionRule{value: val, isSet: true}
}

func (v NullableCouponCreateDataRelationshipsPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCreateDataRelationshipsPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
