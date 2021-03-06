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

// CouponCodesPromotionRuleUpdate struct for CouponCodesPromotionRuleUpdate
type CouponCodesPromotionRuleUpdate struct {
	Data CouponCodesPromotionRuleUpdateData `json:"data"`
}

// NewCouponCodesPromotionRuleUpdate instantiates a new CouponCodesPromotionRuleUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCouponCodesPromotionRuleUpdate(data CouponCodesPromotionRuleUpdateData) *CouponCodesPromotionRuleUpdate {
	this := CouponCodesPromotionRuleUpdate{}
	this.Data = data
	return &this
}

// NewCouponCodesPromotionRuleUpdateWithDefaults instantiates a new CouponCodesPromotionRuleUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponCodesPromotionRuleUpdateWithDefaults() *CouponCodesPromotionRuleUpdate {
	this := CouponCodesPromotionRuleUpdate{}
	return &this
}

// GetData returns the Data field value
func (o *CouponCodesPromotionRuleUpdate) GetData() CouponCodesPromotionRuleUpdateData {
	if o == nil {
		var ret CouponCodesPromotionRuleUpdateData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CouponCodesPromotionRuleUpdate) GetDataOk() (*CouponCodesPromotionRuleUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CouponCodesPromotionRuleUpdate) SetData(v CouponCodesPromotionRuleUpdateData) {
	o.Data = v
}

func (o CouponCodesPromotionRuleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCouponCodesPromotionRuleUpdate struct {
	value *CouponCodesPromotionRuleUpdate
	isSet bool
}

func (v NullableCouponCodesPromotionRuleUpdate) Get() *CouponCodesPromotionRuleUpdate {
	return v.value
}

func (v *NullableCouponCodesPromotionRuleUpdate) Set(val *CouponCodesPromotionRuleUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCouponCodesPromotionRuleUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCouponCodesPromotionRuleUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCouponCodesPromotionRuleUpdate(val *CouponCodesPromotionRuleUpdate) *NullableCouponCodesPromotionRuleUpdate {
	return &NullableCouponCodesPromotionRuleUpdate{value: val, isSet: true}
}

func (v NullableCouponCodesPromotionRuleUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCouponCodesPromotionRuleUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
