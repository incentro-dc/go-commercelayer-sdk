/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETCoupons200ResponseDataInnerRelationshipsPromotionRule struct for GETCoupons200ResponseDataInnerRelationshipsPromotionRule
type GETCoupons200ResponseDataInnerRelationshipsPromotionRule struct {
	Links *GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks   `json:"links,omitempty"`
	Data  *GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData `json:"data,omitempty"`
}

// NewGETCoupons200ResponseDataInnerRelationshipsPromotionRule instantiates a new GETCoupons200ResponseDataInnerRelationshipsPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETCoupons200ResponseDataInnerRelationshipsPromotionRule() *GETCoupons200ResponseDataInnerRelationshipsPromotionRule {
	this := GETCoupons200ResponseDataInnerRelationshipsPromotionRule{}
	return &this
}

// NewGETCoupons200ResponseDataInnerRelationshipsPromotionRuleWithDefaults instantiates a new GETCoupons200ResponseDataInnerRelationshipsPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETCoupons200ResponseDataInnerRelationshipsPromotionRuleWithDefaults() *GETCoupons200ResponseDataInnerRelationshipsPromotionRule {
	this := GETCoupons200ResponseDataInnerRelationshipsPromotionRule{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRule) GetLinks() GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRule) GetLinksOk() (*GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRule) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRule) SetLinks(v GETAddresses200ResponseDataInnerRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRule) GetData() GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData {
	if o == nil || o.Data == nil {
		var ret GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRule) GetDataOk() (*GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRule) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData and assigns it to the Data field.
func (o *GETCoupons200ResponseDataInnerRelationshipsPromotionRule) SetData(v GETCoupons200ResponseDataInnerRelationshipsPromotionRuleData) {
	o.Data = &v
}

func (o GETCoupons200ResponseDataInnerRelationshipsPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRule struct {
	value *GETCoupons200ResponseDataInnerRelationshipsPromotionRule
	isSet bool
}

func (v NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRule) Get() *GETCoupons200ResponseDataInnerRelationshipsPromotionRule {
	return v.value
}

func (v *NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRule) Set(val *GETCoupons200ResponseDataInnerRelationshipsPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCoupons200ResponseDataInnerRelationshipsPromotionRule(val *GETCoupons200ResponseDataInnerRelationshipsPromotionRule) *NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRule {
	return &NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRule{value: val, isSet: true}
}

func (v NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCoupons200ResponseDataInnerRelationshipsPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}