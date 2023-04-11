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

// GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion struct for GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion
type GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                 `json:"links,omitempty"`
	Data  *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotionData `json:"data,omitempty"`
}

// NewGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion instantiates a new GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion() *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion {
	this := GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion{}
	return &this
}

// NewGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotionWithDefaults instantiates a new GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotionWithDefaults() *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion {
	this := GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) GetData() GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotionData {
	if o == nil || o.Data == nil {
		var ret GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) GetDataOk() (*GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotionData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotionData and assigns it to the Data field.
func (o *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) SetData(v GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotionData) {
	o.Data = &v
}

func (o GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion struct {
	value *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion
	isSet bool
}

func (v NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) Get() *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion {
	return v.value
}

func (v *NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) Set(val *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion(val *GETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) *NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion {
	return &NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion{value: val, isSet: true}
}

func (v NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCouponCodesPromotionRules200ResponseDataInnerRelationshipsPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
