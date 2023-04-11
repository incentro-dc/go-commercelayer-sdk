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

// GETOrderAmountPromotionRules200ResponseDataInnerRelationships struct for GETOrderAmountPromotionRules200ResponseDataInnerRelationships
type GETOrderAmountPromotionRules200ResponseDataInnerRelationships struct {
	Promotion *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion `json:"promotion,omitempty"`
}

// NewGETOrderAmountPromotionRules200ResponseDataInnerRelationships instantiates a new GETOrderAmountPromotionRules200ResponseDataInnerRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETOrderAmountPromotionRules200ResponseDataInnerRelationships() *GETOrderAmountPromotionRules200ResponseDataInnerRelationships {
	this := GETOrderAmountPromotionRules200ResponseDataInnerRelationships{}
	return &this
}

// NewGETOrderAmountPromotionRules200ResponseDataInnerRelationshipsWithDefaults instantiates a new GETOrderAmountPromotionRules200ResponseDataInnerRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETOrderAmountPromotionRules200ResponseDataInnerRelationshipsWithDefaults() *GETOrderAmountPromotionRules200ResponseDataInnerRelationships {
	this := GETOrderAmountPromotionRules200ResponseDataInnerRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *GETOrderAmountPromotionRules200ResponseDataInnerRelationships) GetPromotion() GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion {
	if o == nil || o.Promotion == nil {
		var ret GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETOrderAmountPromotionRules200ResponseDataInnerRelationships) GetPromotionOk() (*GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion, bool) {
	if o == nil || o.Promotion == nil {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *GETOrderAmountPromotionRules200ResponseDataInnerRelationships) HasPromotion() bool {
	if o != nil && o.Promotion != nil {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion and assigns it to the Promotion field.
func (o *GETOrderAmountPromotionRules200ResponseDataInnerRelationships) SetPromotion(v GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) {
	o.Promotion = &v
}

func (o GETOrderAmountPromotionRules200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Promotion != nil {
		toSerialize["promotion"] = o.Promotion
	}
	return json.Marshal(toSerialize)
}

type NullableGETOrderAmountPromotionRules200ResponseDataInnerRelationships struct {
	value *GETOrderAmountPromotionRules200ResponseDataInnerRelationships
	isSet bool
}

func (v NullableGETOrderAmountPromotionRules200ResponseDataInnerRelationships) Get() *GETOrderAmountPromotionRules200ResponseDataInnerRelationships {
	return v.value
}

func (v *NullableGETOrderAmountPromotionRules200ResponseDataInnerRelationships) Set(val *GETOrderAmountPromotionRules200ResponseDataInnerRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETOrderAmountPromotionRules200ResponseDataInnerRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETOrderAmountPromotionRules200ResponseDataInnerRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETOrderAmountPromotionRules200ResponseDataInnerRelationships(val *GETOrderAmountPromotionRules200ResponseDataInnerRelationships) *NullableGETOrderAmountPromotionRules200ResponseDataInnerRelationships {
	return &NullableGETOrderAmountPromotionRules200ResponseDataInnerRelationships{value: val, isSet: true}
}

func (v NullableGETOrderAmountPromotionRules200ResponseDataInnerRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETOrderAmountPromotionRules200ResponseDataInnerRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
