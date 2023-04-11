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

// checks if the POSTCouponsRequestDataRelationshipsPromotionRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCouponsRequestDataRelationshipsPromotionRule{}

// POSTCouponsRequestDataRelationshipsPromotionRule struct for POSTCouponsRequestDataRelationshipsPromotionRule
type POSTCouponsRequestDataRelationshipsPromotionRule struct {
	Data POSTCouponsRequestDataRelationshipsPromotionRuleData `json:"data"`
}

// NewPOSTCouponsRequestDataRelationshipsPromotionRule instantiates a new POSTCouponsRequestDataRelationshipsPromotionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCouponsRequestDataRelationshipsPromotionRule(data POSTCouponsRequestDataRelationshipsPromotionRuleData) *POSTCouponsRequestDataRelationshipsPromotionRule {
	this := POSTCouponsRequestDataRelationshipsPromotionRule{}
	this.Data = data
	return &this
}

// NewPOSTCouponsRequestDataRelationshipsPromotionRuleWithDefaults instantiates a new POSTCouponsRequestDataRelationshipsPromotionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCouponsRequestDataRelationshipsPromotionRuleWithDefaults() *POSTCouponsRequestDataRelationshipsPromotionRule {
	this := POSTCouponsRequestDataRelationshipsPromotionRule{}
	return &this
}

// GetData returns the Data field value
func (o *POSTCouponsRequestDataRelationshipsPromotionRule) GetData() POSTCouponsRequestDataRelationshipsPromotionRuleData {
	if o == nil {
		var ret POSTCouponsRequestDataRelationshipsPromotionRuleData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *POSTCouponsRequestDataRelationshipsPromotionRule) GetDataOk() (*POSTCouponsRequestDataRelationshipsPromotionRuleData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *POSTCouponsRequestDataRelationshipsPromotionRule) SetData(v POSTCouponsRequestDataRelationshipsPromotionRuleData) {
	o.Data = v
}

func (o POSTCouponsRequestDataRelationshipsPromotionRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCouponsRequestDataRelationshipsPromotionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePOSTCouponsRequestDataRelationshipsPromotionRule struct {
	value *POSTCouponsRequestDataRelationshipsPromotionRule
	isSet bool
}

func (v NullablePOSTCouponsRequestDataRelationshipsPromotionRule) Get() *POSTCouponsRequestDataRelationshipsPromotionRule {
	return v.value
}

func (v *NullablePOSTCouponsRequestDataRelationshipsPromotionRule) Set(val *POSTCouponsRequestDataRelationshipsPromotionRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCouponsRequestDataRelationshipsPromotionRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCouponsRequestDataRelationshipsPromotionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCouponsRequestDataRelationshipsPromotionRule(val *POSTCouponsRequestDataRelationshipsPromotionRule) *NullablePOSTCouponsRequestDataRelationshipsPromotionRule {
	return &NullablePOSTCouponsRequestDataRelationshipsPromotionRule{value: val, isSet: true}
}

func (v NullablePOSTCouponsRequestDataRelationshipsPromotionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCouponsRequestDataRelationshipsPromotionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}