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

// checks if the POSTOrderAmountPromotionRules201ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderAmountPromotionRules201ResponseDataRelationships{}

// POSTOrderAmountPromotionRules201ResponseDataRelationships struct for POSTOrderAmountPromotionRules201ResponseDataRelationships
type POSTOrderAmountPromotionRules201ResponseDataRelationships struct {
	Promotion *POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion `json:"promotion,omitempty"`
}

// NewPOSTOrderAmountPromotionRules201ResponseDataRelationships instantiates a new POSTOrderAmountPromotionRules201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderAmountPromotionRules201ResponseDataRelationships() *POSTOrderAmountPromotionRules201ResponseDataRelationships {
	this := POSTOrderAmountPromotionRules201ResponseDataRelationships{}
	return &this
}

// NewPOSTOrderAmountPromotionRules201ResponseDataRelationshipsWithDefaults instantiates a new POSTOrderAmountPromotionRules201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderAmountPromotionRules201ResponseDataRelationshipsWithDefaults() *POSTOrderAmountPromotionRules201ResponseDataRelationships {
	this := POSTOrderAmountPromotionRules201ResponseDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *POSTOrderAmountPromotionRules201ResponseDataRelationships) GetPromotion() POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion {
	if o == nil || IsNil(o.Promotion) {
		var ret POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrderAmountPromotionRules201ResponseDataRelationships) GetPromotionOk() (*POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion, bool) {
	if o == nil || IsNil(o.Promotion) {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *POSTOrderAmountPromotionRules201ResponseDataRelationships) HasPromotion() bool {
	if o != nil && !IsNil(o.Promotion) {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion and assigns it to the Promotion field.
func (o *POSTOrderAmountPromotionRules201ResponseDataRelationships) SetPromotion(v POSTCouponCodesPromotionRules201ResponseDataRelationshipsPromotion) {
	o.Promotion = &v
}

func (o POSTOrderAmountPromotionRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderAmountPromotionRules201ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Promotion) {
		toSerialize["promotion"] = o.Promotion
	}
	return toSerialize, nil
}

type NullablePOSTOrderAmountPromotionRules201ResponseDataRelationships struct {
	value *POSTOrderAmountPromotionRules201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTOrderAmountPromotionRules201ResponseDataRelationships) Get() *POSTOrderAmountPromotionRules201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTOrderAmountPromotionRules201ResponseDataRelationships) Set(val *POSTOrderAmountPromotionRules201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderAmountPromotionRules201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderAmountPromotionRules201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderAmountPromotionRules201ResponseDataRelationships(val *POSTOrderAmountPromotionRules201ResponseDataRelationships) *NullablePOSTOrderAmountPromotionRules201ResponseDataRelationships {
	return &NullablePOSTOrderAmountPromotionRules201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTOrderAmountPromotionRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderAmountPromotionRules201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
