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

// checks if the CouponDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CouponDataRelationships{}

// CouponDataRelationships struct for CouponDataRelationships
type CouponDataRelationships struct {
	PromotionRule *CouponDataRelationshipsPromotionRule `json:"promotion_rule,omitempty"`
}

// NewCouponDataRelationships instantiates a new CouponDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponDataRelationships() *CouponDataRelationships {
	this := CouponDataRelationships{}
	return &this
}

// NewCouponDataRelationshipsWithDefaults instantiates a new CouponDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponDataRelationshipsWithDefaults() *CouponDataRelationships {
	this := CouponDataRelationships{}
	return &this
}

// GetPromotionRule returns the PromotionRule field value if set, zero value otherwise.
func (o *CouponDataRelationships) GetPromotionRule() CouponDataRelationshipsPromotionRule {
	if o == nil || IsNil(o.PromotionRule) {
		var ret CouponDataRelationshipsPromotionRule
		return ret
	}
	return *o.PromotionRule
}

// GetPromotionRuleOk returns a tuple with the PromotionRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponDataRelationships) GetPromotionRuleOk() (*CouponDataRelationshipsPromotionRule, bool) {
	if o == nil || IsNil(o.PromotionRule) {
		return nil, false
	}
	return o.PromotionRule, true
}

// HasPromotionRule returns a boolean if a field has been set.
func (o *CouponDataRelationships) HasPromotionRule() bool {
	if o != nil && !IsNil(o.PromotionRule) {
		return true
	}

	return false
}

// SetPromotionRule gets a reference to the given CouponDataRelationshipsPromotionRule and assigns it to the PromotionRule field.
func (o *CouponDataRelationships) SetPromotionRule(v CouponDataRelationshipsPromotionRule) {
	o.PromotionRule = &v
}

func (o CouponDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CouponDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PromotionRule) {
		toSerialize["promotion_rule"] = o.PromotionRule
	}
	return toSerialize, nil
}

type NullableCouponDataRelationships struct {
	value *CouponDataRelationships
	isSet bool
}

func (v NullableCouponDataRelationships) Get() *CouponDataRelationships {
	return v.value
}

func (v *NullableCouponDataRelationships) Set(val *CouponDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponDataRelationships(val *CouponDataRelationships) *NullableCouponDataRelationships {
	return &NullableCouponDataRelationships{value: val, isSet: true}
}

func (v NullableCouponDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
