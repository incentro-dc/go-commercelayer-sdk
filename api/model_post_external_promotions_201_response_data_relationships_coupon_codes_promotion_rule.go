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

// checks if the POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule{}

// POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule struct for POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule
type POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks                         `json:"links,omitempty"`
	Data  *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData `json:"data,omitempty"`
}

// NewPOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule instantiates a new POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule() *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule {
	this := POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule{}
	return &this
}

// NewPOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleWithDefaults instantiates a new POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleWithDefaults() *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule {
	this := POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) GetData() POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData {
	if o == nil || IsNil(o.Data) {
		var ret POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) GetDataOk() (*POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData and assigns it to the Data field.
func (o *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) SetData(v POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRuleData) {
	o.Data = &v
}

func (o POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule struct {
	value *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule
	isSet bool
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) Get() *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule {
	return v.value
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) Set(val *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule(val *POSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) *NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule {
	return &NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule{value: val, isSet: true}
}

func (v NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTExternalPromotions201ResponseDataRelationshipsCouponCodesPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
