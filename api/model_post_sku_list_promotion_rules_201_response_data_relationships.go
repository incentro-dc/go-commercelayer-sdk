/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// POSTSkuListPromotionRules201ResponseDataRelationships struct for POSTSkuListPromotionRules201ResponseDataRelationships
type POSTSkuListPromotionRules201ResponseDataRelationships struct {
	Promotion GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion `json:"promotion"`
	SkuList   *GETBundles200ResponseDataInnerRelationshipsSkuList                    `json:"sku_list,omitempty"`
}

// NewPOSTSkuListPromotionRules201ResponseDataRelationships instantiates a new POSTSkuListPromotionRules201ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSkuListPromotionRules201ResponseDataRelationships(promotion GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) *POSTSkuListPromotionRules201ResponseDataRelationships {
	this := POSTSkuListPromotionRules201ResponseDataRelationships{}
	this.Promotion = promotion
	return &this
}

// NewPOSTSkuListPromotionRules201ResponseDataRelationshipsWithDefaults instantiates a new POSTSkuListPromotionRules201ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSkuListPromotionRules201ResponseDataRelationshipsWithDefaults() *POSTSkuListPromotionRules201ResponseDataRelationships {
	this := POSTSkuListPromotionRules201ResponseDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetPromotion() GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion {
	if o == nil {
		var ret GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion
		return ret
	}

	return o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value
// and a boolean to check if the value has been set.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetPromotionOk() (*GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Promotion, true
}

// SetPromotion sets field value
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) SetPromotion(v GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) {
	o.Promotion = v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetSkuList() GETBundles200ResponseDataInnerRelationshipsSkuList {
	if o == nil || o.SkuList == nil {
		var ret GETBundles200ResponseDataInnerRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) GetSkuListOk() (*GETBundles200ResponseDataInnerRelationshipsSkuList, bool) {
	if o == nil || o.SkuList == nil {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) HasSkuList() bool {
	if o != nil && o.SkuList != nil {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given GETBundles200ResponseDataInnerRelationshipsSkuList and assigns it to the SkuList field.
func (o *POSTSkuListPromotionRules201ResponseDataRelationships) SetSkuList(v GETBundles200ResponseDataInnerRelationshipsSkuList) {
	o.SkuList = &v
}

func (o POSTSkuListPromotionRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["promotion"] = o.Promotion
	}
	if o.SkuList != nil {
		toSerialize["sku_list"] = o.SkuList
	}
	return json.Marshal(toSerialize)
}

type NullablePOSTSkuListPromotionRules201ResponseDataRelationships struct {
	value *POSTSkuListPromotionRules201ResponseDataRelationships
	isSet bool
}

func (v NullablePOSTSkuListPromotionRules201ResponseDataRelationships) Get() *POSTSkuListPromotionRules201ResponseDataRelationships {
	return v.value
}

func (v *NullablePOSTSkuListPromotionRules201ResponseDataRelationships) Set(val *POSTSkuListPromotionRules201ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSkuListPromotionRules201ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSkuListPromotionRules201ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSkuListPromotionRules201ResponseDataRelationships(val *POSTSkuListPromotionRules201ResponseDataRelationships) *NullablePOSTSkuListPromotionRules201ResponseDataRelationships {
	return &NullablePOSTSkuListPromotionRules201ResponseDataRelationships{value: val, isSet: true}
}

func (v NullablePOSTSkuListPromotionRules201ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSkuListPromotionRules201ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}