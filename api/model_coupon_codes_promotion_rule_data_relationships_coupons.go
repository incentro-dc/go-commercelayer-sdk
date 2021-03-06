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

// CouponCodesPromotionRuleDataRelationshipsCoupons struct for CouponCodesPromotionRuleDataRelationshipsCoupons
type CouponCodesPromotionRuleDataRelationshipsCoupons struct {
	// The resource's type
	Type string `json:"type"`
	// The resource's id
	Id string `json:"id"`
}

// NewCouponCodesPromotionRuleDataRelationshipsCoupons instantiates a new CouponCodesPromotionRuleDataRelationshipsCoupons object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodesPromotionRuleDataRelationshipsCoupons(type_ string, id string) *CouponCodesPromotionRuleDataRelationshipsCoupons {
	this := CouponCodesPromotionRuleDataRelationshipsCoupons{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCouponCodesPromotionRuleDataRelationshipsCouponsWithDefaults instantiates a new CouponCodesPromotionRuleDataRelationshipsCoupons object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodesPromotionRuleDataRelationshipsCouponsWithDefaults() *CouponCodesPromotionRuleDataRelationshipsCoupons {
	this := CouponCodesPromotionRuleDataRelationshipsCoupons{}
	var type_ string = "coupons"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *CouponCodesPromotionRuleDataRelationshipsCoupons) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleDataRelationshipsCoupons) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CouponCodesPromotionRuleDataRelationshipsCoupons) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CouponCodesPromotionRuleDataRelationshipsCoupons) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleDataRelationshipsCoupons) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CouponCodesPromotionRuleDataRelationshipsCoupons) SetId(v string) {
	o.Id = v
}

func (o CouponCodesPromotionRuleDataRelationshipsCoupons) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCouponCodesPromotionRuleDataRelationshipsCoupons struct {
	value *CouponCodesPromotionRuleDataRelationshipsCoupons
	isSet bool
}

func (v NullableCouponCodesPromotionRuleDataRelationshipsCoupons) Get() *CouponCodesPromotionRuleDataRelationshipsCoupons {
	return v.value
}

func (v *NullableCouponCodesPromotionRuleDataRelationshipsCoupons) Set(val *CouponCodesPromotionRuleDataRelationshipsCoupons) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodesPromotionRuleDataRelationshipsCoupons) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodesPromotionRuleDataRelationshipsCoupons) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodesPromotionRuleDataRelationshipsCoupons(val *CouponCodesPromotionRuleDataRelationshipsCoupons) *NullableCouponCodesPromotionRuleDataRelationshipsCoupons {
	return &NullableCouponCodesPromotionRuleDataRelationshipsCoupons{value: val, isSet: true}
}

func (v NullableCouponCodesPromotionRuleDataRelationshipsCoupons) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodesPromotionRuleDataRelationshipsCoupons) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
