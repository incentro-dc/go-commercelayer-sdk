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

// CouponUpdateDataRelationships struct for CouponUpdateDataRelationships
type CouponUpdateDataRelationships struct {
	PromotionRule *CouponCreateDataRelationshipsPromotionRule `json:"promotion_rule,omitempty"`
}

// NewCouponUpdateDataRelationships instantiates a new CouponUpdateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponUpdateDataRelationships() *CouponUpdateDataRelationships {
	this := CouponUpdateDataRelationships{}
	return &this
}

// NewCouponUpdateDataRelationshipsWithDefaults instantiates a new CouponUpdateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponUpdateDataRelationshipsWithDefaults() *CouponUpdateDataRelationships {
	this := CouponUpdateDataRelationships{}
	return &this
}

// GetPromotionRule returns the PromotionRule field value if set, zero value otherwise.
func (o *CouponUpdateDataRelationships) GetPromotionRule() CouponCreateDataRelationshipsPromotionRule {
	if o == nil || o.PromotionRule == nil {
		var ret CouponCreateDataRelationshipsPromotionRule
		return ret
	}
	return *o.PromotionRule
}

// GetPromotionRuleOk returns a tuple with the PromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponUpdateDataRelationships) GetPromotionRuleOk() (*CouponCreateDataRelationshipsPromotionRule, bool) {
	if o == nil || o.PromotionRule == nil {
		return nil, false
	}
	return o.PromotionRule, true
}

// HasPromotionRule returns a boolean if a field has been set.
func (o *CouponUpdateDataRelationships) HasPromotionRule() bool {
	if o != nil && o.PromotionRule != nil {
		return true
	}

	return false
}

// SetPromotionRule gets a reference to the given CouponCreateDataRelationshipsPromotionRule and assigns it to the PromotionRule field.
func (o *CouponUpdateDataRelationships) SetPromotionRule(v CouponCreateDataRelationshipsPromotionRule) {
	o.PromotionRule = &v
}

func (o CouponUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PromotionRule != nil {
		toSerialize["promotion_rule"] = o.PromotionRule
	}
	return json.Marshal(toSerialize)
}

type NullableCouponUpdateDataRelationships struct {
	value *CouponUpdateDataRelationships
	isSet bool
}

func (v NullableCouponUpdateDataRelationships) Get() *CouponUpdateDataRelationships {
	return v.value
}

func (v *NullableCouponUpdateDataRelationships) Set(val *CouponUpdateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponUpdateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponUpdateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponUpdateDataRelationships(val *CouponUpdateDataRelationships) *NullableCouponUpdateDataRelationships {
	return &NullableCouponUpdateDataRelationships{value: val, isSet: true}
}

func (v NullableCouponUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponUpdateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
