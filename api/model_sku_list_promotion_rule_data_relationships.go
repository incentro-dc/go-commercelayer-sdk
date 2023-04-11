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

// SkuListPromotionRuleDataRelationships struct for SkuListPromotionRuleDataRelationships
type SkuListPromotionRuleDataRelationships struct {
	Promotion *CouponCodesPromotionRuleDataRelationshipsPromotion `json:"promotion,omitempty"`
	SkuList   *BundleDataRelationshipsSkuList                     `json:"sku_list,omitempty"`
	Skus      *BundleDataRelationshipsSkus                        `json:"skus,omitempty"`
}

// NewSkuListPromotionRuleDataRelationships instantiates a new SkuListPromotionRuleDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuListPromotionRuleDataRelationships() *SkuListPromotionRuleDataRelationships {
	this := SkuListPromotionRuleDataRelationships{}
	return &this
}

// NewSkuListPromotionRuleDataRelationshipsWithDefaults instantiates a new SkuListPromotionRuleDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuListPromotionRuleDataRelationshipsWithDefaults() *SkuListPromotionRuleDataRelationships {
	this := SkuListPromotionRuleDataRelationships{}
	return &this
}

// GetPromotion returns the Promotion field value if set, zero value otherwise.
func (o *SkuListPromotionRuleDataRelationships) GetPromotion() CouponCodesPromotionRuleDataRelationshipsPromotion {
	if o == nil || o.Promotion == nil {
		var ret CouponCodesPromotionRuleDataRelationshipsPromotion
		return ret
	}
	return *o.Promotion
}

// GetPromotionOk returns a tuple with the Promotion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleDataRelationships) GetPromotionOk() (*CouponCodesPromotionRuleDataRelationshipsPromotion, bool) {
	if o == nil || o.Promotion == nil {
		return nil, false
	}
	return o.Promotion, true
}

// HasPromotion returns a boolean if a field has been set.
func (o *SkuListPromotionRuleDataRelationships) HasPromotion() bool {
	if o != nil && o.Promotion != nil {
		return true
	}

	return false
}

// SetPromotion gets a reference to the given CouponCodesPromotionRuleDataRelationshipsPromotion and assigns it to the Promotion field.
func (o *SkuListPromotionRuleDataRelationships) SetPromotion(v CouponCodesPromotionRuleDataRelationshipsPromotion) {
	o.Promotion = &v
}

// GetSkuList returns the SkuList field value if set, zero value otherwise.
func (o *SkuListPromotionRuleDataRelationships) GetSkuList() BundleDataRelationshipsSkuList {
	if o == nil || o.SkuList == nil {
		var ret BundleDataRelationshipsSkuList
		return ret
	}
	return *o.SkuList
}

// GetSkuListOk returns a tuple with the SkuList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleDataRelationships) GetSkuListOk() (*BundleDataRelationshipsSkuList, bool) {
	if o == nil || o.SkuList == nil {
		return nil, false
	}
	return o.SkuList, true
}

// HasSkuList returns a boolean if a field has been set.
func (o *SkuListPromotionRuleDataRelationships) HasSkuList() bool {
	if o != nil && o.SkuList != nil {
		return true
	}

	return false
}

// SetSkuList gets a reference to the given BundleDataRelationshipsSkuList and assigns it to the SkuList field.
func (o *SkuListPromotionRuleDataRelationships) SetSkuList(v BundleDataRelationshipsSkuList) {
	o.SkuList = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *SkuListPromotionRuleDataRelationships) GetSkus() BundleDataRelationshipsSkus {
	if o == nil || o.Skus == nil {
		var ret BundleDataRelationshipsSkus
		return ret
	}
	return *o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuListPromotionRuleDataRelationships) GetSkusOk() (*BundleDataRelationshipsSkus, bool) {
	if o == nil || o.Skus == nil {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *SkuListPromotionRuleDataRelationships) HasSkus() bool {
	if o != nil && o.Skus != nil {
		return true
	}

	return false
}

// SetSkus gets a reference to the given BundleDataRelationshipsSkus and assigns it to the Skus field.
func (o *SkuListPromotionRuleDataRelationships) SetSkus(v BundleDataRelationshipsSkus) {
	o.Skus = &v
}

func (o SkuListPromotionRuleDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Promotion != nil {
		toSerialize["promotion"] = o.Promotion
	}
	if o.SkuList != nil {
		toSerialize["sku_list"] = o.SkuList
	}
	if o.Skus != nil {
		toSerialize["skus"] = o.Skus
	}
	return json.Marshal(toSerialize)
}

type NullableSkuListPromotionRuleDataRelationships struct {
	value *SkuListPromotionRuleDataRelationships
	isSet bool
}

func (v NullableSkuListPromotionRuleDataRelationships) Get() *SkuListPromotionRuleDataRelationships {
	return v.value
}

func (v *NullableSkuListPromotionRuleDataRelationships) Set(val *SkuListPromotionRuleDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuListPromotionRuleDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuListPromotionRuleDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuListPromotionRuleDataRelationships(val *SkuListPromotionRuleDataRelationships) *NullableSkuListPromotionRuleDataRelationships {
	return &NullableSkuListPromotionRuleDataRelationships{value: val, isSet: true}
}

func (v NullableSkuListPromotionRuleDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuListPromotionRuleDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
