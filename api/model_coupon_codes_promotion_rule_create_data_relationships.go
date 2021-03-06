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

// CouponCodesPromotionRuleCreateDataRelationships struct for CouponCodesPromotionRuleCreateDataRelationships
type CouponCodesPromotionRuleCreateDataRelationships struct {
	Promotion CouponCodesPromotionRuleDataRelationshipsPromotion `json:"promotion"`
	Coupons   *CouponCodesPromotionRuleDataRelationshipsCoupons  `json:"coupons,omitempty"`
}

// NewCouponCodesPromotionRuleCreateDataRelationships instantiates a new CouponCodesPromotionRuleCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodesPromotionRuleCreateDataRelationships(promotion CouponCodesPromotionRuleDataRelationshipsPromotion) *CouponCodesPromotionRuleCreateDataRelationships {
	this := CouponCodesPromotionRuleCreateDataRelationships{}
	this.Promotion = promotion
	return &this
}

// NewCouponCodesPromotionRuleCreateDataRelationshipsWithDefaults instantiates a new CouponCodesPromotionRuleCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodesPromotionRuleCreateDataRelationshipsWithDefaults() *CouponCodesPromotionRuleCreateDataRelationships {
	this := CouponCodesPromotionRuleCreateDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value
func (o *CouponCodesPromotionRuleCreateDataRelationships) GetPromotion() CouponCodesPromotionRuleDataRelationshipsPromotion {
	if o == nil {
		var ret CouponCodesPromotionRuleDataRelationshipsPromotion
		return ret
	}

	return o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleCreateDataRelationships) GetPromotionOk() (*CouponCodesPromotionRuleDataRelationshipsPromotion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Promotion, true
}

// SetPromotion sets field value
func (o *CouponCodesPromotionRuleCreateDataRelationships) SetPromotion(v CouponCodesPromotionRuleDataRelationshipsPromotion) {
	o.Promotion = v
}

// GetCoupons returns the Coupons field value if set, zero value otherwise.
func (o *CouponCodesPromotionRuleCreateDataRelationships) GetCoupons() CouponCodesPromotionRuleDataRelationshipsCoupons {
	if o == nil || o.Coupons == nil {
		var ret CouponCodesPromotionRuleDataRelationshipsCoupons
		return ret
	}
	return *o.Coupons
}

// GetCouponsOk returns a tuple with the Coupons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleCreateDataRelationships) GetCouponsOk() (*CouponCodesPromotionRuleDataRelationshipsCoupons, bool) {
	if o == nil || o.Coupons == nil {
		return nil, false
	}
	return o.Coupons, true
}

// HasCoupons returns a boolean if a field has been set.
func (o *CouponCodesPromotionRuleCreateDataRelationships) HasCoupons() bool {
	if o != nil && o.Coupons != nil {
		return true
	}

	return false
}

// SetCoupons gets a reference to the given CouponCodesPromotionRuleDataRelationshipsCoupons and assigns it to the Coupons field.
func (o *CouponCodesPromotionRuleCreateDataRelationships) SetCoupons(v CouponCodesPromotionRuleDataRelationshipsCoupons) {
	o.Coupons = &v
}

func (o CouponCodesPromotionRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["promotion"] = o.Promotion
	}
	if o.Coupons != nil {
		toSerialize["coupons"] = o.Coupons
	}
	return json.Marshal(toSerialize)
}

type NullableCouponCodesPromotionRuleCreateDataRelationships struct {
	value *CouponCodesPromotionRuleCreateDataRelationships
	isSet bool
}

func (v NullableCouponCodesPromotionRuleCreateDataRelationships) Get() *CouponCodesPromotionRuleCreateDataRelationships {
	return v.value
}

func (v *NullableCouponCodesPromotionRuleCreateDataRelationships) Set(val *CouponCodesPromotionRuleCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodesPromotionRuleCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodesPromotionRuleCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodesPromotionRuleCreateDataRelationships(val *CouponCodesPromotionRuleCreateDataRelationships) *NullableCouponCodesPromotionRuleCreateDataRelationships {
	return &NullableCouponCodesPromotionRuleCreateDataRelationships{value: val, isSet: true}
}

func (v NullableCouponCodesPromotionRuleCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodesPromotionRuleCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
