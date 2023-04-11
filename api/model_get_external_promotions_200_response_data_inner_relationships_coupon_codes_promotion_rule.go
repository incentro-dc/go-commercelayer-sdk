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

// GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule struct for GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule
type GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks                         `json:"links,omitempty"`
	Data  *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRuleData `json:"data,omitempty"`
}

// NewGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule instantiates a new GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule() *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule {
	this := GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule{}
	return &this
}

// NewGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRuleWithDefaults instantiates a new GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRuleWithDefaults() *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule {
	this := GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) GetData() GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRuleData {
	if o == nil || o.Data == nil {
		var ret GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) GetDataOk() (*GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRuleData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRuleData and assigns it to the Data field.
func (o *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) SetData(v GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRuleData) {
	o.Data = &v
}

func (o GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule struct {
	value *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule
	isSet bool
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) Get() *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule {
	return v.value
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) Set(val *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule(val *GETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) *NullableGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule {
	return &NullableGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule{value: val, isSet: true}
}

func (v NullableGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETExternalPromotions200ResponseDataInnerRelationshipsCouponCodesPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
